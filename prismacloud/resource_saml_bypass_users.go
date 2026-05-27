package prismacloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/user/saml"
)

func resourceSamlBypassUsers() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSamlBypassUsersCreate,
		ReadContext:   resourceSamlBypassUsersRead,
		UpdateContext: resourceSamlBypassUsersUpdate,
		DeleteContext: resourceSamlBypassUsersDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"usernames": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "List of usernames that are allowed to bypass SAML authentication.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
		Description: "Manage the list of users that are allowed to bypass SAML authentication.",
	}
}

func resourceSamlBypassUsersCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*pc.Client)
	log.Printf("[DEBUG] Starting resourceSamlBypassUsersCreate")
	usernames := convertStringSet(d.Get("usernames"))

	log.Printf("[DEBUG] resourceSamlBypassUsersCreate Updated users list to send to API: %v", usernames)
	err := saml.UpdateBypassUsers(c, usernames)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("saml_bypass_users")
	return resourceSamlBypassUsersRead(ctx, d, meta)
}

func resourceSamlBypassUsersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*pc.Client)
	log.Printf("[DEBUG] Starting resourceSamlBypassUsersRead")
	usernames, err := saml.GetBypassUsers(c)
	log.Printf("[DEBUG] resourceSamlBypassUsersRead user list from API: %v", usernames)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("usernames", usernames); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceSamlBypassUsersUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*pc.Client)
	log.Printf("[DEBUG] Starting resourceSamlBypassUsersUpdate")
	usernames := convertStringSet(d.Get("usernames"))
	log.Printf("[DEBUG] resourceSamlBypassUsersUpdate Updated users list to send to API: %v", usernames)
	err := saml.UpdateBypassUsers(c, usernames)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceSamlBypassUsersRead(ctx, d, meta)
}

func resourceSamlBypassUsersDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*pc.Client)
	log.Printf("[DEBUG] Starting resourceSamlBypassUsersDelete")

	// To "delete" we set an empty list of usernames
	err := saml.UpdateBypassUsers(c, []string{})
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// convertStringSet converts a *schema.Set to a []string
func convertStringSet(set interface{}) []string {
	result := make([]string, 0)
	for _, v := range set.(*schema.Set).List() {
		result = append(result, v.(string))
	}
	return result
}
