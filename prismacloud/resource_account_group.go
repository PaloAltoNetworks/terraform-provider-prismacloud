package prismacloud

import (
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/cloud/account/group"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAccountGroup() *schema.Resource {
	return &schema.Resource{
		Create: createAccountGroup,
		Read:   readAccountGroup,
		Update: updateAccountGroup,
		Delete: deleteAccountGroup,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the group",
			},
			"group_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Account group ID",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description",
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
			"account_ids": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Cloud account IDs",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			/*
				"accounts": {
					Type:        schema.TypeList,
					Computed:    true,
					Description: "Associated cloud accounts",
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"account_id": {
								Type:        schema.TypeString,
								Computed:    true,
								Description: "Associated cloud account ID",
							},
							"name": {
								Type:        schema.TypeString,
								Computed:    true,
								Description: "Associated cloud account name",
							},
							"account_type": {
								Type:        schema.TypeString,
								Computed:    true,
								Description: "Associated cloud account type",
							},
						},
					},
				},
				"alert_rules": {
					Type:        schema.TypeList,
					Computed:    true,
					Description: "Singly associated alert rules which cannot exist in the system without the account group",
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"alert_id": {
								Type:        schema.TypeString,
								Computed:    true,
								Description: "The alert ID",
							},
							"name": {
								Type:        schema.TypeString,
								Computed:    true,
								Description: "Alert name",
							},
						},
					},
				},
			*/
		},
	}
}

func parseAccountGroup(d *schema.ResourceData, id string) group.Group {
	return group.Group{
		Id:          id,
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		AccountIds:  []string{},
	}
}

func saveAccountGroup(d *schema.ResourceData, obj group.Group) {
	var err error

	d.Set("group_id", obj.Id)
	d.Set("name", obj.Name)
	d.Set("description", obj.Description)
	d.Set("last_modified_by", obj.LastModifiedBy)
	d.Set("last_modified_ts", obj.LastModifiedTs)
	if err = d.Set("account_ids", obj.AccountIds); err != nil {
		log.Printf("[WARN] Error setting 'account_ids' field for %q: %s", d.Id(), err)
	}

	/*
	   // Neither accounts nor alert rules is returned from the API when querying
	   // a specific group ID (today).
	   aInfo := make([]interface{}, 0, len(obj.Accounts))
	   for _, ai := range obj.Accounts {
	       aInfo = append(aInfo, map[string]interface{}{
	           "account_id":   ai.Id,
	           "name":         ai.Name,
	           "account_type": ai.Type,
	       })
	   }
	   if err = d.Set("accounts", aInfo); err != nil {
	       log.Printf("[WARN] Error setting 'accounts' field for %q: %s", d.Id(), err)
	   }

	   arInfo := make([]interface{}, 0, len(obj.AlertRules))
	   for _, ari := range obj.AlertRules {
	       arInfo = append(arInfo, map[string]interface{}{
	           "alert_id": ari.Id,
	           "name":     ari.Name,
	       })
	   }
	   if err = d.Set("alert_rules", arInfo); err != nil {
	       log.Printf("[WARN] Error setting 'alert_rules' field for %q: %s", d.Id(), err)
	   }
	*/
}

func createAccountGroup(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	obj := parseAccountGroup(d, "")

	if err := group.Create(client, obj); err != nil {
		return err
	}

	id, err := group.Identify(client, obj.Name)
	if err != nil {
		return err
	}

	d.SetId(id)
	return readAccountGroup(d, meta)
}

func readAccountGroup(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	id := d.Id()

	obj, err := group.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return err
	}

	saveAccountGroup(d, obj)

	return nil
}

func updateAccountGroup(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	obj := parseAccountGroup(d, d.Id())

	if err := group.UpdateUsingLiveAccountIds(client, obj); err != nil {
		return err
	}

	return readAccountGroup(d, meta)
}

func deleteAccountGroup(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	id := d.Id()

	err := group.Delete(client, id)
	if err != nil {
		if err != pc.ObjectNotFoundError {
			return err
		}
	}

	d.SetId("")
	return nil
}
