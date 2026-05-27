package prismacloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/user/saml"
)

func resourceSamlBypassSubsetUsers() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSamlBypassSubsetUsersCreate,
		ReadContext:   resourceSamlBypassSubsetUsersRead,
		UpdateContext: resourceSamlBypassSubsetUsersUpdate,
		DeleteContext: resourceSamlBypassSubsetUsersDelete,
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
		Description: "Manage a subset of users that are allowed to bypass SAML authentication. This resource adds or removes only the specified users while preserving other users that may be managed outside of Terraform.",
	}
}

func resourceSamlBypassSubsetUsersCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*pc.Client)
	log.Printf("[DEBUG] Starting resourceSamlBypassSubsetUsersCreate")
	usernames := convertStringSet(d.Get("usernames"))

	// Get the current list of users from the API
	currentUsers, err := saml.GetBypassUsers(c)
	if err != nil {
		return diag.FromErr(err)
	}

	// Reconcile the current users with the new users
	updatedUsers := reconcileUsers(currentUsers, usernames, []string{})

	// Update the API with the updated list
	log.Printf("[DEBUG] resourceSamlBypassSubsetUsersCreate Updated users list to send to API: %v", updatedUsers)
	err = saml.UpdateBypassUsers(c, updatedUsers)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("saml_bypass_subset_users")
	return resourceSamlBypassSubsetUsersRead(ctx, d, meta)
}

func resourceSamlBypassSubsetUsersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*pc.Client)
	log.Printf("[DEBUG] Starting resourceSamlBypassSubsetUsersRead")
	// Get the current list of users from the API
	currentUsers, err := saml.GetBypassUsers(c)
	if err != nil {
		return diag.FromErr(err)
	}

	// Get the list of users managed by Terraform
	terraformUsers := convertStringSet(d.Get("usernames"))

	// Save the list of users managed by Terraform
	saveSamlBypassSubsetUsers(d, terraformUsers)

	// Output for debugging
	log.Printf("[DEBUG] resourceSamlBypassSubsetUsersRead Current users from API: %v", currentUsers)
	log.Printf("[DEBUG] resourceSamlBypassSubsetUsersRead Terraform-managed users: %v", terraformUsers)

	return nil
}

func resourceSamlBypassSubsetUsersUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*pc.Client)
	log.Printf("[DEBUG] Starting resourceSamlBypassSubsetUsersUpdate")
	// Get the current list of users from the API
	currentUsers, err := saml.GetBypassUsers(c)
	if err != nil {
		return diag.FromErr(err)
	}

	// Get the previous and current list of users managed by Terraform
	oldUsers, newUsers := d.GetChange("usernames")
	oldUsersList := convertStringSet(oldUsers)
	newUsersList := convertStringSet(newUsers)

	// Reconcile the users
	updatedUsers := reconcileUsers(currentUsers, newUsersList, oldUsersList)

	// Update the API with the updated list
	log.Printf("[DEBUG] resourceSamlBypassSubsetUsersUpdate - Updated users to send to API: %v", updatedUsers)
	err = saml.UpdateBypassUsers(c, updatedUsers)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceSamlBypassSubsetUsersRead(ctx, d, meta)
}

func resourceSamlBypassSubsetUsersDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*pc.Client)

	log.Printf("[DEBUG] Starting resourceSamlBypassSubsetUsersDelete")
	// Get the current list of users from the API
	currentUsers, err := saml.GetBypassUsers(c)
	if err != nil {
		return diag.FromErr(err)
	}

	// Get the list of users managed by Terraform
	terraformUsers := convertStringSet(d.Get("usernames"))

	// Remove the Terraform-managed users from the current list
	updatedUsers := reconcileUsers(currentUsers, []string{}, terraformUsers)
	// Update the API with the updated list
	log.Printf("[DEBUG] resourceSamlBypassSubsetUsersDelete Updated users to send to API: %v", updatedUsers)
	err = saml.UpdateBypassUsers(c, updatedUsers)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// reconcileUsers reconciles the current users with the new users managed by Terraform
func reconcileUsers(currentUsers []string, newUsers []string, oldUsers []string) []string {
	// Add new users that are not already in the current list
	for _, user := range newUsers {
		if !contains(currentUsers, user) {
			currentUsers = append(currentUsers, user)
		}
	}

	// Remove users that are no longer in the new list but were in the old list
	for _, user := range oldUsers {
		if !contains(newUsers, user) && contains(currentUsers, user) {
			currentUsers = remove(currentUsers, user)
		}
	}
	return currentUsers
}

func saveSamlBypassSubsetUsers(d *schema.ResourceData, usernames []string) error {
	// Save the list of users managed by Terraform
	if err := d.Set("usernames", usernames); err != nil {
		return err
	}
	return nil
}

// contains checks if a string is present in a slice
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// remove removes a string from a slice
func remove(slice []string, item string) []string {
	for i, s := range slice {
		if s == item {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
