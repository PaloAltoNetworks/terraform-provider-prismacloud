package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"golang.org/x/net/context"
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/user/profile"
	"github.com/paloaltonetworks/prisma-cloud-go/user/role"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUserRole() *schema.Resource {
	return &schema.Resource{
		CreateContext: createUserRole,
		ReadContext:   readUserRole,
		UpdateContext: updateUserRole,
		DeleteContext: deleteUserRole,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the role",
			},
			"role_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Role UUID",
			},
			"delete_associated_users": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Delete any associated users on role deletion. This is use useful when SSO is enabled",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description",
			},
			"role_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "User role type",
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
			"account_group_ids": {
				// TODO: Is this ordered or unordered?
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Accessible account group IDs",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"resource_list_ids": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Resource list IDs",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"code_repository_ids": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Code repository IDs",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"associated_users": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Associated application users which cannot exist in the system without the user role",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"restrict_dismissal_access": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Restrict dismissal access",
			},
			"account_groups": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Associated account groups",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The group ID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Group name",
						},
					},
				},
			},
			"additional_attributes": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "Additional Parameters",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"only_allow_ci_access": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Allows only CI Access",
						},
						"only_allow_compute_access": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Give access to only compute tab and access key tab",
						},
						"only_allow_read_access": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Allow only read access",
						},
						"has_defender_permissions": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Has defender Permissions",
						},
					},
				},
			},
		},
	}
}

func parseUserRole(d *schema.ResourceData) *role.Role {
	aspec := d.Get("additional_attributes").([]interface{})

	ans := &role.Role{
		Name:                    d.Get("name").(string),
		Description:             d.Get("description").(string),
		RoleType:                d.Get("role_type").(string),
		AccountGroupIds:         SetToStringSlice(d.Get("account_group_ids").(*schema.Set)),
		ResourceListIds:         SetToStringSlice(d.Get("resource_list_ids").(*schema.Set)),
		CodeRepositoryIds:       SetToStringSlice(d.Get("code_repository_ids").(*schema.Set)),
		RestrictDismissalAccess: d.Get("restrict_dismissal_access").(bool),
	}

	if len(aspec) > 0 {
		if aspecs := aspec[0].(map[string]interface{}); len(aspecs) > 0 {
			ans.AdditionalAttributes.OnlyAllowReadAccess = aspecs["only_allow_read_access"].(bool)
			ans.AdditionalAttributes.OnlyAllowComputeAccess = aspecs["only_allow_compute_access"].(bool)
			ans.AdditionalAttributes.OnlyAllowCIAccess = aspecs["only_allow_ci_access"].(bool)
			ans.AdditionalAttributes.HasDefenderPermissions = aspecs["has_defender_permissions"].(bool)
		}
	}

	return ans
}

func saveUserRole(d *schema.ResourceData, obj role.Role) {
	var err error

	d.Set("name", obj.Name)
	d.Set("role_id", obj.Id)
	d.Set("description", obj.Description)
	d.Set("role_type", obj.RoleType)
	d.Set("last_modified_by", obj.LastModifiedBy)
	d.Set("last_modified_ts", obj.LastModifiedTs)

	add_attr := map[string]interface{}{
		"only_allow_ci_access":      obj.AdditionalAttributes.OnlyAllowCIAccess,
		"only_allow_compute_access": obj.AdditionalAttributes.OnlyAllowComputeAccess,
		"only_allow_read_access":    obj.AdditionalAttributes.OnlyAllowReadAccess,
		"has_defender_permissions":  obj.AdditionalAttributes.HasDefenderPermissions,
	}
	if err := d.Set("additional_attributes", []interface{}{add_attr}); err != nil {
		log.Printf("[WARN] Error setting 'rule' for %q: %s", d.Id(), err)
	}
	if err = d.Set("account_group_ids", StringSliceToSet(obj.AccountGroupIds)); err != nil {
		log.Printf("[WARN] Error setting 'account_group_ids' field for %q: %s", d.Id(), err)
	}
	if err = d.Set("resource_list_ids", StringSliceToSet(obj.ResourceListIds)); err != nil {
		log.Printf("[WARN] Error setting 'resource_list_ids' field for %q: %s", d.Id(), err)
	}
	if err = d.Set("code_repository_ids", StringSliceToSet(obj.CodeRepositoryIds)); err != nil {
		log.Printf("[WARN] Error setting 'code_repository_ids' field for %q: %s", d.Id(), err)
	}
	if err = d.Set("associated_users", StringSliceToSet(obj.AssociatedUsers)); err != nil {
		log.Printf("[WARN] Error setting 'associated_users' field for %q: %s", d.Id(), err)
	}
	d.Set("restrict_dismissal_access", obj.RestrictDismissalAccess)

	agInfo := make([]interface{}, 0, len(obj.AccountGroups))
	for _, ag := range obj.AccountGroups {
		agInfo = append(agInfo, map[string]interface{}{
			"group_id": ag.Id,
			"name":     ag.Name,
		})
	}
	if err = d.Set("account_groups", agInfo); err != nil {
		log.Printf("[WARN] Error setting 'account_groups' field for %q: %s", d.Id(), err)
	}
}

func createUserRole(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	obj := parseUserRole(d)

	if err := role.Create(client, *obj); err != nil {
		return diag.FromErr(err)
	}

	PollApiUntilSuccess(func() error {
		_, err := role.Identify(client, obj.Name)
		return err
	})

	id, err := role.Identify(client, obj.Name)
	if err != nil {
		return diag.FromErr(err)
	}

	PollApiUntilSuccess(func() error {
		_, err := role.Get(client, id)
		return err
	})

	d.SetId(id)
	return readUserRole(ctx, d, meta)
}

func readUserRole(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	id := d.Id()

	obj, err := role.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	saveUserRole(d, obj)

	return nil
}

func updateUserRole(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	obj := parseUserRole(d)
	obj.Id = d.Id()

	if err := role.Update(client, *obj); err != nil {
		return diag.FromErr(err)
	}

	return readUserRole(ctx, d, meta)
}

func deleteUserRole(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	id := d.Id()

	delete_associated_users := d.Get("delete_associated_users").(bool)
	if delete_associated_users {
		associated_users := SetToStringSlice(d.Get("associated_users").(*schema.Set))
		for _, user := range associated_users {
			log.Printf("[DEBUG] Purging user %s", user)
			if err := profile.Delete(client, user, profile.TypeUserAccount); err != nil {
				return diag.FromErr(err)
			}
		}
	}

	err := role.Delete(client, id)
	if err != nil {
		if err != pc.ObjectNotFoundError {
			return diag.FromErr(err)
		}
	}

	d.SetId("")
	return nil
}
