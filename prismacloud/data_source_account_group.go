package prismacloud

import (
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/cloud/account/group"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAccountGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAccountGroupRead,

		Schema: map[string]*schema.Schema{
			// Input.
			"group_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Account group ID",
				AtLeastOneOf: []string{"group_id", "name"},
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Name of the group",
				AtLeastOneOf: []string{"group_id", "name"},
			},

			// Output.
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
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

func dataSourceAccountGroupRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)

	var err error
	id := d.Get("group_id").(string)

	if id == "" {
		name := d.Get("name").(string)
		id, err = group.Identify(client, name)
		if err != nil {
			if err == pc.ObjectNotFoundError {
				d.SetId("")
				return nil
			}
			return err
		}
	}

	obj, err := group.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return err
	}

	d.SetId(obj.Id)
	saveAccountGroup(d, obj)

	return nil
}
