# Terraform provider gitcrypt

This is a simple terraform provider which can decrypt files encrypted with [git-crypt](https://github.com/AGWA/git-crypt).

# Actions status

[![release](https://github.com/bcdtriptech/terraform-provider-gitcrypt/actions/workflows/release.yml/badge.svg)](https://github.com/bcdtriptech/terraform-provider-gitcrypt/actions/workflows/release.yml)  [![Tests](https://github.com/bcdtriptech/terraform-provider-gitcrypt/actions/workflows/test.yml/badge.svg?branch=main&event=push)](https://github.com/bcdtriptech/terraform-provider-gitcrypt/actions/workflows/test.yml)

# What is this provider for?

In any infrastructure, there is a set of secrets that need to be stored somewhere. If you define your infrastructure as a code, that is good to be able to keep secrets in code too. However, they need to be secured properly. One way to achieve that is to encrypt your secrets in VCS, e.g. with [git-crypt](https://github.com/AGWA/git-crypt).

In order to provide terraform with access to encrypted content you can use this provider to decrypt and parse secret files on the fly. It is especially useful in cases when you have no control over terraform execution environment to decrypt files with git-crypt CLI, for example if you use [Terraform Enterprise](https://app.terraform.io).

# Documentation

[Terraform docs](https://registry.terraform.io/providers/bcdtriptech/gitcrypt/latest/docs)

# Encrypted file format

gitcrypt terraform provider can parse simple file which contains `key: value` pairs like `var1: value1`.

You can see [ENCRYPTED](gitcrypt/test-data/encrypted_vars.yml) and [DECRYPTED](gitcrypt/test-data/decrypted_vars.yml) files example.  

# Contributing

If you're having trouble using gitcrypt provider, create a [Github issue](https://github.com/bcdtriptech/terraform-provider-gitcrypt/issues) or open a [pull request](https://github.com/bcdtriptech/terraform-provider-gitcrypt/pulls).
