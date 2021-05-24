package gitcrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"io/ioutil"
)

const (
	// git-crypt encrypted file consists of: ( header[10], IV(HMAC[12], blockID[4]), encryptedData[...] )
	gitCryptFileHeaderLength = 10
	gitCryptFileHMACLength   = 12

	// git-crypt key file consists of: ( header[32], AESInfoBlock[8], AES[32], HMACInfoBlock[8], HMAC[64], EndBlock[4] )
	gitCryptKeyInfoBlockLength = 8
	gitCryptKeyHeaderLength    = 32
	gitCryptKeyAESLength       = 32
	gitCryptKeyHMACLength      = 64
	gitCryptKeyEndBlockLength  = 4
	gitCryptKeyFileLength      = (gitCryptKeyHeaderLength + gitCryptKeyInfoBlockLength + gitCryptKeyHeaderLength +
		gitCryptKeyInfoBlockLength + gitCryptKeyHMACLength + gitCryptKeyEndBlockLength)
)

// KeyData is struct for keep AES and HMAC from git-crypt key
type KeyData struct {
	AES  []byte
	HMAC []byte
}

// LoadKey convert key base64 to KeyData struct
func LoadKey(keyFileDataBase64 string) (KeyData, error) {
	//convert base64 to []byte
	keyFileData, _ := base64.StdEncoding.DecodeString(keyFileDataBase64)
	// Git-crypt key size must be 148 bytes
	if len(keyFileData) != gitCryptKeyFileLength {
		return KeyData{}, errors.New("Error: invalid git-crypt keyfile lengh")
	}

	// Identificators for AES and HMAC keys.(constants from the git-crypt repo)
	const AESKeyIdentificator = 3
	const HMACKeyIdentificator = 5

	// need to cut out header:
	data := keyFileData[gitCryptKeyHeaderLength:]

	// create variables for AESkey and HMACkey
	AESKey := make([]byte, gitCryptKeyAESLength)
	HMACKey := make([]byte, gitCryptKeyHMACLength)

	// looking for info block and confirm AES and HMAC position and length
	for len(data) >= gitCryptKeyInfoBlockLength {
		blockID := (data[0] << 24) | (data[1] << 16) | (data[2] << 8) | data[3]
		blockLen := (data[4] << 24) | (data[5] << 16) | (data[6] << 8) | data[7]

		if int(blockID) == AESKeyIdentificator {
			AESKey = data[gitCryptKeyInfoBlockLength : gitCryptKeyInfoBlockLength+blockLen]
		} else if int(blockID) == HMACKeyIdentificator {
			HMACKey = data[gitCryptKeyInfoBlockLength : gitCryptKeyInfoBlockLength+blockLen]
		}
		// cut out the info block and parameter value and continue
		data = data[gitCryptKeyInfoBlockLength+blockLen:]
	}

	if len(AESKey) == 0 || len(HMACKey) == 0 {
		return KeyData{}, errors.New("Error: invalid git-crypt keyfile")
	}

	keyContent := KeyData{AES: AESKey, HMAC: HMACKey}
	return keyContent, nil
}

// UnlockFile return content of the file decrypted by AES from KeyData
func UnlockFile(filepath string, secretsKeys KeyData) ([]byte, error) {

	encryptedData, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gitCryptFileHMACLength)
	copy(nonce, encryptedData[gitCryptFileHeaderLength:(gitCryptFileHeaderLength+gitCryptFileHMACLength)])

	decryptedData, err := decryptData(encryptedData, secretsKeys.AES)
	if err != nil {
		return nil, err
	}

	if !validMAC(decryptedData, nonce, secretsKeys.HMAC) {
		return nil, errors.New("HMAC is not valid. The file could be corrupted")
	}

	return decryptedData, nil
}

// GetFileHMAC returns encrypted file HMAC
func GetFileHMAC(filePath string) (string, error) {
	encryptedData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	HMAC := make([]byte, gitCryptFileHMACLength)
	copy(HMAC, encryptedData[gitCryptFileHeaderLength:(gitCryptFileHeaderLength+gitCryptFileHMACLength)])

	return base64.StdEncoding.EncodeToString(HMAC), nil
}

func validMAC(message, messageMAC, key []byte) bool {

	mac := hmac.New(sha1.New, key)
	mac.Write(message)
	// log.Printf("[DEBUG] File HMAC: %v", messageMAC)

	expectedMAC := mac.Sum(nil)[:gitCryptFileHMACLength]
	// log.Printf("[DEBUG] Expected HMAC: %v", expectedMAC)

	return hmac.Equal(messageMAC, expectedMAC)
}

func decryptData(encryptedData, AESKey []byte) ([]byte, error) {
	// creating cipher with received AES Key
	blockCipher, err := aes.NewCipher(AESKey)
	if err != nil {
		return nil, err
	}

	// need to cut out git-crypt header from file:
	encryptedData = encryptedData[gitCryptFileHeaderLength:]

	// nonce is unique information about encrypted file (file SHA-1 HMAC)
	nonce := encryptedData[:gitCryptFileHMACLength]

	// need to cut out nonce from encryptedData
	encryptedData = encryptedData[gitCryptFileHMACLength:]

	// need to create IV for decryption
	iv := make([]byte, gitCryptFileHMACLength)

	// for every block need to use unique IV (nonce + blockID)
	// IV lengh is 16 bytes (nonce[12]+blockID[4])
	copy(iv, nonce)
	iv = append(iv, []byte{0, 0, 0, 0}...)

	// create stream with cipher and IV
	stream := cipher.NewCTR(blockCipher, iv)

	decFileContent := make([]byte, len(encryptedData))

	// decrypt file content
	stream.XORKeyStream(decFileContent, encryptedData)

	return decFileContent, nil
}
