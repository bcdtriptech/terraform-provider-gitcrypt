package gitcrypt

import (
	"reflect"
	"testing"
)

const encryptedFilePath = "../../test-data/encrypted_vars.yml"

func TestLoadKey(t *testing.T) {
	cases := []struct {
		keyFileDataBase64 string
		expectedResult    KeyData
		errorIsExpected   bool
	}{
		{
			keyFileDataBase64: string("AEdJVENSWVBUS0VZAAAAAgAAAAAAAAABAAAABAAAAAAAAAADAAAAIDJ6yMP6EdHmYJ2VyFa1LU1zitt4G4gJdD3O1/8L1ZZEAAAABQAAAEAtubx4wwVHvOAIuz/K7fvrtFFUBzsA2Dl4AGuyK3WGOd1v1HuDFW6tN65V4D3j+M4+0ly25+xYukN7Qdw6ZjDJAAAAAA=="),
			expectedResult: KeyData{
				AES: []byte{50, 122, 200, 195, 250, 17, 209, 230, 96, 157,
					149, 200, 86, 181, 45, 77, 115, 138, 219, 120, 27, 136,
					9, 116, 61, 206, 215, 255, 11, 213, 150, 68},
				HMAC: []byte{45, 185, 188, 120, 195, 5, 71, 188, 224, 8,
					187, 63, 202, 237, 251, 235, 180, 81, 84, 7, 59, 0, 216,
					57, 120, 0, 107, 178, 43, 117, 134, 57, 221, 111, 212,
					123, 131, 21, 110, 173, 55, 174, 85, 224, 61, 227, 248,
					206, 62, 210, 92, 182, 231, 236, 88, 186, 67, 123, 65,
					220, 58, 102, 48, 201},
			},
			errorIsExpected: false,
		},
		{
			keyFileDataBase64: string("SOMEWRONGKEYAAAAAgAAAAAAAAABAAAABAAAAAAAAAADAAAAIDJ6yMP6EdHmYJ2VyFa1LU1zitt4G4gJdD3O1/8L1ZZEAAAABQAAAEAtubx4wwVHvOAIuz/K7fvrtFFUBzsA2Dl4AGuyK3WGOd1v1HuDFW6tN65V4D3j+M4+0ly25+xYukN7Qdw6ZjDJAAAAAA=="),
			expectedResult:    KeyData{},
			errorIsExpected:   true,
		},
	}

	for tcIdx, tc := range cases {
		actualResult, err := LoadKey(tc.keyFileDataBase64)
		if err != nil && !tc.errorIsExpected {
			t.Fatalf("Test case #%d\nUnexpected error appeared: %s", tcIdx+1, err)
		}
		if err == nil && tc.errorIsExpected {
			t.Fatalf("Test case #%d\nExpected error didn't appear.", tcIdx+1)
		}
		if !reflect.DeepEqual(actualResult, tc.expectedResult) {
			t.Fatalf("Test case #%d\nWrong!\nexpected: %s\nactual: %s", tcIdx+1, tc.expectedResult, actualResult)
		}
	}
}

func TestDecryptData(t *testing.T) {
	cases := []struct {
		encryptedData   []byte
		AESKey          []byte
		expectedResult  []byte
		errorIsExpected bool
	}{
		{ // right encrypted data and right key
			encryptedData: []byte{0, 71, 73, 84, 67, 82, 89, 80, 84,
				0, 121, 62, 40, 129, 210, 59, 2, 229, 70, 128, 139, 8, 114,
				35, 106, 94, 37, 79, 238, 86, 189, 124, 57, 4},
			AESKey: []byte{50, 122, 200, 195, 250, 17, 209, 230,
				96, 157, 149, 200, 86, 181, 45, 77, 115, 138, 219, 120, 27,
				136, 9, 116, 61, 206, 215, 255, 11, 213, 150, 68},
			expectedResult:  []byte{100, 117, 109, 109, 121, 32, 118, 97, 108, 117, 101, 10},
			errorIsExpected: false,
		},
		{ // right encrypted data and wrong key
			encryptedData: []byte{0, 71, 73, 84, 67, 82, 89, 80, 84, 0, 121,
				62, 40, 129, 210, 59, 2, 229, 70, 128, 139, 8, 114, 35, 106, 94,
				37, 79, 238, 86, 189, 124, 57, 4},
			AESKey:          []byte{27, 136, 9, 116, 61, 206, 215, 255, 11, 213, 150, 68},
			expectedResult:  nil,
			errorIsExpected: true,
		},
	}

	for tcIdx, tc := range cases {
		actualResult, err := decryptData(tc.encryptedData, tc.AESKey)
		if err != nil && !tc.errorIsExpected {
			t.Fatalf("Test case #%d\nUnexpected error appeared: %s", tcIdx+1, err)
		}
		if err == nil && tc.errorIsExpected {
			t.Fatalf("Test case #%d\nExpected error didn't appear.", tcIdx+1)
		}
		if !reflect.DeepEqual(actualResult, tc.expectedResult) {
			t.Fatalf("Test case #%d\nWrong!\nexpected: %s\nactual: %s", tcIdx+1, tc.expectedResult, actualResult)
		}
	}
}

func TestUnlockFile(t *testing.T) {
	cases := []struct {
		filePath        string
		secretKey       KeyData
		expectedResult  []byte
		errorIsExpected bool
	}{
		{ // right secretKey
			filePath: encryptedFilePath,
			secretKey: KeyData{
				AES: []byte{50, 122, 200, 195, 250, 17, 209, 230, 96, 157,
					149, 200, 86, 181, 45, 77, 115, 138, 219, 120, 27, 136,
					9, 116, 61, 206, 215, 255, 11, 213, 150, 68},
				HMAC: []byte{45, 185, 188, 120, 195, 5, 71, 188, 224, 8,
					187, 63, 202, 237, 251, 235, 180, 81, 84, 7, 59, 0, 216,
					57, 120, 0, 107, 178, 43, 117, 134, 57, 221, 111, 212,
					123, 131, 21, 110, 173, 55, 174, 85, 224, 61, 227, 248,
					206, 62, 210, 92, 182, 231, 236, 88, 186, 67, 123, 65,
					220, 58, 102, 48, 201},
			},
			expectedResult: []byte{118, 97, 114, 49, 58, 32, 118, 97,
				108, 117, 101, 49, 10, 118, 97, 114, 50, 58, 32, 118, 97,
				108, 117, 101, 50, 10, 118, 97, 114, 51, 58, 32, 118, 97,
				108, 117, 101, 51, 10},
			errorIsExpected: false,
		},
		{ // wrong secretKey
			filePath: encryptedFilePath,
			secretKey: KeyData{
				AES: []byte{50, 122, 200, 195, 250, 17, 209, 230, 96, 157,
					108, 117, 101, 49, 10, 149, 200, 86, 181, 45, 77, 115,
					138, 219, 120, 27, 136},
				HMAC: []byte{45, 185, 188, 120, 195, 5, 71, 188, 224, 8,
					187, 63, 202, 237, 251, 235, 180, 81, 84, 7, 59, 0, 216,
					57, 120, 0, 107, 178, 43, 117, 134, 57, 221, 111, 212,
					220, 58, 102, 48, 201},
			},
			expectedResult:  nil,
			errorIsExpected: true,
		},
	}

	for tcIdx, tc := range cases {
		actualResult, err := UnlockFile(tc.filePath, tc.secretKey)
		if err != nil && !tc.errorIsExpected {
			t.Fatalf("Test case #%d\nUnexpected error appeared: %s", tcIdx+1, err)
		}
		if err == nil && tc.errorIsExpected {
			t.Fatalf("Test case #%d\nExpected error didn't appear.", tcIdx+1)
		}
		if !reflect.DeepEqual(actualResult, tc.expectedResult) {
			t.Fatalf("Test case #%d\nWrong!\nexpected:\n %v\nactual:\n %v", tcIdx+1, tc.expectedResult, string(actualResult))
		}
	}
}

func TestGetFileHMAC(t *testing.T) {
	cases := []struct {
		filePath        string
		expectedResult  string
		errorIsExpected bool
	}{
		{
			filePath:        encryptedFilePath,
			expectedResult:  "wvJBqVVqBQrjlHYF",
			errorIsExpected: false,
		},
		{
			filePath:        "notExistingEncryptedFilePath",
			expectedResult:  "",
			errorIsExpected: true,
		},
	}

	for tcIdx, tc := range cases {
		actualResult, err := GetFileHMAC(tc.filePath)
		if err != nil && !tc.errorIsExpected {
			t.Fatalf("Test case #%d\nUnexpected error appeared: %s", tcIdx+1, err)
		}
		if err == nil && tc.errorIsExpected {
			t.Fatalf("Test case #%d\nExpected error didn't appear.", tcIdx+1)
		}
		if !reflect.DeepEqual(actualResult, tc.expectedResult) {
			t.Fatalf("Test case #%d\nWrong!\nexpected:\n %v\nactual:\n %v", tcIdx+1, tc.expectedResult, string(actualResult))
		}
	}
}
