---
layout: "gitcrypt"
page_title: "Provider: gitcrypt"
description: |-
  The gitcrypt provider is used to read files encrypted with git-crypt.
---

# gitcrypt Provider

The gitcrypt provider is used to read files encrypted with git-crypt.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the gitcrypt Provider
provider "gitcrypt" {
}

# Read encrypte file:
data "gitcrypt_encrypted_file" "example" {
  file_path = "./test-data/encrypted_vars.yml"
}
```

## Argument Reference

The following arguments are supported in the `provider` block:

* `gitcrypt_key_base64` - (Required) A git-crypt key for repository.
When not provided or made available via the `GIT_CRYPT_KEY_BASE64` or `KEY_BASE64` environment variable.

!> It is strongly NOT recommended to set parameter `gitcrypt_key_base64` in the open. It is dangerous and not secure.
Anyone who knows the value of argument `gitcrypt_key_base64` can decrypt secret files.

-> To understand where to get the value for argument `gitcrypt_key_base64`, see [this page](https://github.com/bcdtriptech/terraform-provider-gitcrypt#encrypted-file-format).
