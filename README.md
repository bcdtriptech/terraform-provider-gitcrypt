# Terraform provider gitcrypt

This is a simple terraform provider which can decrypt files encrypted with [git-crypt](https://github.com/AGWA/git-crypt).

# Actions status

[![release](https://github.com/bcdtriptech/terraform-provider-gitcrypt/actions/workflows/release.yml/badge.svg)](https://github.com/bcdtriptech/terraform-provider-gitcrypt/actions/workflows/release.yml)  [![Tests](https://github.com/bcdtriptech/terraform-provider-gitcrypt/actions/workflows/test.yml/badge.svg?branch=main&event=push)](https://github.com/bcdtriptech/terraform-provider-gitcrypt/actions/workflows/test.yml)

# Using gitcrypt provider

To use this provider you need get git-crypt key of you repository in base64 format:
```
  $ cd ./your-repository/
  $ git-crypt unlock
  $ base64 -i .git/git-crypt/keys/default
```
The example of output:
```
  AEdJVENSWVBUS0VZAAAAAgAAAAAAAAABAAAABAAAAAAAAAADAAAAIDJ6yMP6EdHmYJ2VyFa1LU1zitt4G4gJdD3O1/8L1ZZEAAAABQAAAEAtubx4wwVHvOAIuz/K7fvrtFFUBzsA2Dl4AGuyK3WGOd1v1HuDFW6tN65V4D3j+M4+0ly25+xYukN7Qdw6ZjDJAAAAAA==
```
Then set this value as environment variable `GIT_CRYPT_KEY_BASE64` or `KEY_BASE64` on machine where you will init gitcrypt terraform provider. If you use Terraform Enterprise you can create environment variable for your workspace.

You also can set it as parameter `gitcrypt_key_base64` in provider section like this:
```
  provider "gitcrypt" {
      gitcrypt_key_base64 = "AEdJVENSWVBUS0VZAAAAAgAAAAAAAAABAAAABAAAAAAAAAADAAAAIDJ6yMP6EdHmYJ2VyFa1LU1zitt4G4gJdD3O1/8L1ZZEAAAABQAAAEAtubx4wwVHvOAIuz/K7fvrtFFUBzsA2Dl4AGuyK3WGOd1v1HuDFW6tN65V4D3j+M4+0ly25+xYukN7Qdw6ZjDJAAAAAA=="
  }
```
WARNING! This method is NOT secure and NOT recommended because everyone who know your `gitcrypt_key_base64` can decrypt you secret files!

# Encrypted file format

gitcrypt terraform provider can parse simple file which contain `key: value` pairs like `var1: value1`.

You can see [ENCRYPTED](gitcrypt/test-data/encrypted_vars.yml) and [DECRYPTED](gitcrypt/test-data/decrypted_vars.yml) files example.  

# Contributing

If you're having trouble using gitcrypt provider, create a [Github issue](https://github.com/bcdtriptech/terraform-provider-gitcrypt/issues) or open a [pull request](https://github.com/bcdtriptech/terraform-provider-gitcrypt/pulls).
