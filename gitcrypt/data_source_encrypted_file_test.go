package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceEncryptedFile(t *testing.T) {

	resource.UnitTest(t, resource.TestCase{
		// PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceEncryptedFile,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("data.gitcrypt_encrypted_file.example", "file_path", regexp.MustCompile("^.*")),
					resource.TestCheckResourceAttr("data.gitcrypt_encrypted_file.example", "secrets.var1", "value1"),
					resource.TestCheckResourceAttr("data.gitcrypt_encrypted_file.example", "secrets.var2", "value2"),
					resource.TestCheckResourceAttr("data.gitcrypt_encrypted_file.example", "secrets.var3", "value3"),
				),
			},
		},
	})
}

const testAccDataSourceEncryptedFile = `
provider "gitcrypt" {
    gitcrypt_key_base64 = "AEdJVENSWVBUS0VZAAAAAgAAAAAAAAABAAAABAAAAAAAAAADAAAAIDJ6yMP6EdHmYJ2VyFa1LU1zitt4G4gJdD3O1/8L1ZZEAAAABQAAAEAtubx4wwVHvOAIuz/K7fvrtFFUBzsA2Dl4AGuyK3WGOd1v1HuDFW6tN65V4D3j+M4+0ly25+xYukN7Qdw6ZjDJAAAAAA=="
}

data "gitcrypt_encrypted_file" "example" {
  file_path = "./test-data/encrypted_vars.yml"
}
`
