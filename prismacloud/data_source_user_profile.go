package prismacloud

import (
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/user/profile"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUserProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUserProfileRead,

		Schema: map[string]*schema.Schema{
			// Input.
			"profile_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Profile ID",
			},

			// Output.
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
			"first_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "First name",
			},
			"last_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Last name",
			},
			"display_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Display name",
			},
			"email": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Email ID",
			},
			"access_keys_allowed": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Access keys allowed",
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
			"access_keys_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Access key count",
			},
			"roles": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Model for User Profile Role Detail",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User Role ID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User Role Name",
						},
						"only_allow_ci_access": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "true = Allow only CI Access for Build and Deploy security roles",
						},
						"only_allow_compute_access": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "true = Allow only Compute Access for reduced system admin roles",
						},
						"only_allow_read_access": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "true = Allow only read access",
						},
						"role_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User Role Type",
						},
					},
				},
			},
		},
	}
}

func dataSourceUserProfileRead(d *schema.ResourceData, meta interface{}) error {
	var err error
	client := meta.(*pc.Client)

	id := d.Get("profile_id").(string)

	obj, err := profile.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return err
	}

	d.SetId(id)
	saveUserProfile(d, obj)

	return nil
}
