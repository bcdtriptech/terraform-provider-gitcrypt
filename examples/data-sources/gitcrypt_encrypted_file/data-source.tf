terraform {
  required_providers {
    gitcrypt = {
      source = "bcdtriptech/gitcrypt"
      version = "0.0.1"
    }
  }
}

# init provider
provider "gitcrypt" {
    gitcrypt_key_base64 = "AEdJVENSWVBUS0VZAAAAAgAAAAAAAAABAAAABAAAAAAAAAADAAAAIDJ6yMP6EdHmYJ2VyFa1LU1zitt4G4gJdD3O1/8L1ZZEAAAABQAAAEAtubx4wwVHvOAIuz/K7fvrtFFUBzsA2Dl4AGuyK3WGOd1v1HuDFW6tN65V4D3j+M4+0ly25+xYukN7Qdw6ZjDJAAAAAA=="
    # this key can be set as an ENV variable GIT_CRYPT_KEY_BASE64 or KEY_BASE64
}

# decrypt and parse git-crypt encrypted file
data "gitcrypt_encrypted_file" "some_file" {
  file_path = "./encrypted_vars.yml"
}

# use secrets to define your own resources
# outputs are used as an example in order not to import any other providers
output "var1" {
  value = data.gitcrypt_encrypted_file.some_file.secrets.var1
  # or can be used in this format:
  # value = data.git_crypt_encrypted_file.yml_variables.vars["var1"]
}

output "all_vars" {
  value = data.gitcrypt_encrypted_file.some_file.secrets
  # or can be used in this format:
  # value = data.git_crypt_encrypted_file.yml_variables.secrets[*]
}
