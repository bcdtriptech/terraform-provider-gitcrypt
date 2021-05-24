package provider

import (
	"context"

	gc "github.com/bcdtriptech/terraform-provider-gitcrypt/gitcrypt/internal/gitcrypt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gopkg.in/yaml.v3"
)

func dataSourceEncryptedFile() *schema.Resource {
	return &schema.Resource{
		// Data source for read and decrypt file encrypted by git-crypt
		Description: "Data source for read and decrypt file encrypted by git-crypt.",
		ReadContext: dataSourceEncryptedFileRead,
		Schema: map[string]*schema.Schema{
			"file_path": {
				// Path to the file encrypted by git-crypt
				Description: "Path to the file encrypted by git-crypt",
				Type:        schema.TypeString,
				Required:    true,
			},
			"secrets": &schema.Schema{
				// Variables from ecrypted file after decryption
				Description: "Variables from ecrypted file after decryption",
				Type:        schema.TypeMap,
				Computed:    true,
				Sensitive:   true,
			},
		},
	}
}

func dataSourceEncryptedFileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	gitcryptKey := meta.(*gc.KeyData)

	filePath := d.Get("file_path").(string)

	decryptedContent, err := gc.UnlockFile(filePath, *gitcryptKey)
	if err != nil {
		return diag.FromErr(err)
	}

	secretsMap := make(map[string]string, 0)

	err = yaml.Unmarshal(decryptedContent, &secretsMap)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("secrets", secretsMap)
	if err != nil {
		return diag.FromErr(err)
	}

	fileHMAC, err := gc.GetFileHMAC(filePath)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(fileHMAC)

	return nil
}
