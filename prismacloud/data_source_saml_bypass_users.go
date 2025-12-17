package prismacloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/user/saml"
)

func dataSourceSamlBypassUsers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSamlBypassUsersRead,
		Schema: map[string]*schema.Schema{
			"usernames": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of usernames that are allowed to bypass SAML authentication.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
		Description: "Retrieve the list of users that are allowed to bypass SAML authentication.",
	}
}

func dataSourceSamlBypassUsersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*pc.Client)
	log.Printf("[DEBUG] Starting dataSourceSamlBypassUsersRead")
	usernames, err := saml.GetBypassUsers(c)
	log.Printf("[DEBUG] dataSourceSamlBypassUsersRead - Got usernames: %v", usernames)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("usernames", usernames); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("saml_bypass_users")
	return nil
}
