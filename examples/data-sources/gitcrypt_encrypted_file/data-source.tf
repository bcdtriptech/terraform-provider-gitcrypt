terraform {
  required_providers {
    gitcrypt = {
      source = "bcdtriptech/gitcrypt"
      version = "1.0.0"
    }
  }
}

provider "gitcrypt" {
    gitcrypt_key_base64 = "AEdJVENSWVBUS0VZAAAAAgAAAAAAAAABAAAABAAAAAAAAAADAAAAIDJ6yMP6EdHmYJ2VyFa1LU1zitt4G4gJdD3O1/8L1ZZEAAAABQAAAEAtubx4wwVHvOAIuz/K7fvrtFFUBzsA2Dl4AGuyK3WGOd1v1HuDFW6tN65V4D3j+M4+0ly25+xYukN7Qdw6ZjDJAAAAAA=="
    # this key can be set as an ENV variable GIT_CRYPT_KEY_BASE64 or KEY_BASE64
}

data "gitcrypt_encrypted_file" "some_file" {
  file_path = "./encrypted_vars.yml"
}
