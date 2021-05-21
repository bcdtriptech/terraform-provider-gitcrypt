---
layout: "gitcrypt"
page_title: "gitcrypt: gitcrypt_encrypted_file"
description: |-
  Read and decrypt the content of the encrypted with git-crypt file.
---

# gitcrypt_encrypted_file

Use this data source to read and decrypt file encrypted with git-crypt.
For decryption file will be used key provided to `provider`.   

The data source can parse simple file which contain `key: value` pairs.
Example of the encrypted file content:
```
var1: value1
var2: value2
var3: value3
```
## Example Usage

```hcl
# Read encrypte file:
data "gitcrypt_encrypted_file" "example" {
  file_path = "./test-data/encrypted_vars.yml"
}

# Use value from encrypted file (2 formats available)
resource "db_instance" "default {
  ...
  password = data.gitcrypt_encrypted_file.example.secrets.db_password
}

resource "aws_ssm_parameter" "db_password" {
  name   = "db_password"
  type   = "SecureString"
  value  = data.gitcrypt_encrypted_file.example.secrets["db_password"]
}
```

## Argument Reference

 * `file_path` - (Required) Path to the encrypted file.

## Attributes Reference

 * `file_path` - Path to the encrypted file.
 * `secrets`   - Map of the secrets from encrypted file. Use one of the following syntaxes to access a specific secret:
 ```
 value = data.gitcrypt_encrypted_file.example.secrets.var1
 value = data.gitcrypt_encrypted_file.example.secrets["var1"]
 ```
