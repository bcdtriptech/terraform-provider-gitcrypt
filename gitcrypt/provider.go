package provider

import (
	"context"

	gc "github.com/bcdtriptech/terraform-provider-gitcrypt/gitcrypt/internal/gitcrypt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	schema.DescriptionKind = schema.StringMarkdown
}

// New git-crypt provider
func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"gitcrypt_key_base64": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
					DefaultFunc: schema.MultiEnvDefaultFunc([]string{
						"GIT_CRYPT_KEY_BASE64",
						"KEY_BASE64",
					}, nil),
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"gitcrypt_encrypted_file": dataSourceEncryptedFile(),
			},
			ResourcesMap: map[string]*schema.Resource{},
		}
		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func configure(terraformVersion string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		// convert git-crypt key and get AES and HMAC
		keyBase64 := d.Get("gitcrypt_key_base64").(string)
		gitcryptKeyContent, err := gc.LoadKey(keyBase64)
		if err != nil {
			return &gc.KeyData{}, diag.FromErr(err)
		}
		return &gitcryptKeyContent, nil
	}
}
