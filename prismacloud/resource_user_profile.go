package prismacloud

import (
	"encoding/json"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/user/profile"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserProfile() *schema.Resource {
	return &schema.Resource{
		Create: createUserProfile,
		Read:   readUserProfile,
		Update: updateUserProfile,
		Delete: deleteUserProfile,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"profile_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Profile ID",
			},
			"account_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Account Type",
				Default:     profile.TypeUserAccount,
				ValidateFunc: validation.StringInSlice(
					[]string{
						profile.TypeUserAccount,
						profile.TypeServiceAccount,
					},
					false,
				),
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "User email or service account name",
			},
			"first_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "First name",
			},
			"last_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Last name",
			},
			"display_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Display name",
			},
			"email": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Email ID",
			},
			"access_keys_allowed": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Access keys allowed",
			},
			"access_key_expiration": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Access key expiration timestamp in milliseconds",
			},
			"access_key_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Access key name",
			},
			"access_key_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Access key ID",
			},
			"secret_key": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Access key secret",
			},
			"default_role_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Default Role ID",
			},
			"enable_key_expiration": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Enable access key expiration",
			},
			"role_ids": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "List of Role IDs",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"time_zone": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Time zone (e.g. America/Los_Angeles)",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
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
				Description: "Model for User Profile Roles Details",
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

func parseUserProfile(d *schema.ResourceData, id string) profile.Profile {
	roleIds := d.Get("role_ids")

	return profile.Profile{
		Id:                  id,
		AccountType:         d.Get("account_type").(string),
		Username:            d.Get("username").(string),
		FirstName:           d.Get("first_name").(string),
		LastName:            d.Get("last_name").(string),
		Email:               d.Get("email").(string),
		AccessKeysAllowed:   d.Get("access_keys_allowed").(bool),
		AccessKeyExpiration: d.Get("access_key_expiration").(int),
		AccessKeyName:       d.Get("access_key_name").(string),
		DefaultRoleId:       d.Get("default_role_id").(string),
		EnableKeyExpiration: d.Get("enable_key_expiration").(bool),
		RoleIds:             SetToStringSlice(roleIds.(*schema.Set)),
		TimeZone:            d.Get("time_zone").(string),
		Enabled:             d.Get("enabled").(bool),
	}
}

func saveUserProfile(d *schema.ResourceData, o profile.Profile) {

	d.Set("profile_id", o.Username)
	d.Set("account_type", o.AccountType)
	d.Set("last_name", o.LastName)
	d.Set("username", o.Username)
	d.Set("display_name", o.DisplayName)
	d.Set("access_keys_allowed", o.AccessKeysAllowed)
	d.Set("default_role_id", o.DefaultRoleId)
	d.Set("role_ids", o.RoleIds)
	d.Set("time_zone", o.TimeZone)
	d.Set("enabled", o.Enabled)
	d.Set("last_login_ts", o.LastLoginTs)
	d.Set("last_modified_by", o.LastModifiedBy)
	d.Set("last_modified_ts", o.LastModifiedTs)
	d.Set("access_keys_count", o.AccessKeysCount)

	if o.AccountType == profile.TypeUserAccount {
		d.Set("email", o.Email)
		d.Set("first_name", o.FirstName)
	}

	if len(o.Roles) == 0 {
		d.Set("roles", nil)
		return
	}

	roleList := make([]interface{}, 0, len(o.Roles))
	for _, role := range o.Roles {
		roleList = append(roleList, map[string]interface{}{
			"role_id":                   role.RoleId,
			"name":                      role.Name,
			"only_allow_ci_access":      role.OnlyAllowCIAccess,
			"only_allow_compute_access": role.OnlyAllowComputeAccess,
			"only_allow_read_access":    role.OnlyAllowReadAccess,
			"role_type":                 role.RoleType,
		})
	}
	if err := d.Set("roles", roleList); err != nil {
		log.Printf("[WARN] Error setting 'roles' for %q: %s", d.Id(), err)
	}
}

func createUserProfile(d *schema.ResourceData, meta interface{}) error {
	var err error
	client := meta.(*pc.Client)
	o := parseUserProfile(d, "")

	id := o.Username
	var keyResponse []byte
	if keyResponse, err = profile.Create(client, o); err != nil {
		return err
	}
	var accessKeyResponse profile.AccessKeyResponse
	json.Unmarshal(keyResponse, &accessKeyResponse)
	PollApiUntilSuccess(func() error {
		_, err := profile.Get(client, id)
		return err
	})

	d.SetId(id)
	d.Set("access_key_id", accessKeyResponse.AccessKeyId)
	d.Set("secret_key", accessKeyResponse.SecretKey)
	return readUserProfile(d, meta)
}

func readUserProfile(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	id := d.Id()

	o, err := profile.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return err
	}

	saveUserProfile(d, o)
	return nil
}

func updateUserProfile(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	id := d.Id()
	o := parseUserProfile(d, id)

	if _, err := profile.Update(client, o); err != nil {
		return err
	}

	return readUserProfile(d, meta)
}

func deleteUserProfile(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	id := d.Id()
	accountType := d.Get("account_type").(string)
	err := profile.Delete(client, id, accountType)
	if err != nil {
		if err != pc.ObjectNotFoundError {
			return err
		}
	}

	d.SetId("")
	return nil
}
