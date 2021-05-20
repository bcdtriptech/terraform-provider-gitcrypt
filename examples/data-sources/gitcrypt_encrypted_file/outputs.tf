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
