package prismacloud

import (
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/user/profile"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceUserProfiles() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUserProfilesRead,

		Schema: map[string]*schema.Schema{
			// Output.
			"total": totalSchema("all integrations"),
			"listing": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of all integrations",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"profile_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Profile ID",
						},
						"account_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account Type",
						},
						"username": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User email or service account name",
						},
						"display_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Display name",
						},
						"default_role_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Default Role ID",
						},
						"role_ids": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "List of Role IDs",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"time_zone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Time zone (e.g. America/Los_Angeles)",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enabled",
						},
						"last_login_ts": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Last login time",
						},
						"last_modified_by": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Last modified by",
						},
						"last_modified_ts": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Last modified timestamp",
						},
					},
				},
			},
		},
	}
}

func dataSourceUserProfilesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)

	items, err := profile.List(client)
	if err != nil {
		return err
	}

	d.SetId("user profiles")
	d.Set("total", len(items))

	listing := make([]interface{}, 0, len(items))
	for _, o := range items {
		listing = append(listing, map[string]interface{}{
			"profile_id":       o.Username,
			"account_type":     o.AccountType,
			"username":         o.Username,
			"display_name":     o.DisplayName,
			"default_role_id":  o.DefaultRoleId,
			"role_ids":         o.RoleIds,
			"time_zone":        o.TimeZone,
			"enabled":          o.Enabled,
			"last_login_ts":    o.LastLoginTs,
			"last_modified_by": o.LastModifiedBy,
			"last_modified_ts": o.LastModifiedTs,
		})
	}

	if err := d.Set("listing", listing); err != nil {
		log.Printf("[WARN] Error setting 'listing' field for %q: %s", d.Id(), err)
	}

	return nil
}
