package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/data-security/datapattern"
	"golang.org/x/net/context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDataPattern() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDataPatternRead,

		Schema: map[string]*schema.Schema{
			// Input.
			"pattern_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Pattern ID",
				AtLeastOneOf: []string{"name", "pattern_id"},
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Pattern name",
				AtLeastOneOf: []string{"name", "pattern_id"},
			},

			// Output.
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Pattern description",
			},
			"mode": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Mode - predefined or custom",
			},
			"detection_technique": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Detection technique",
			},
			"entity": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Entity value",
			},
			"grammar": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Grammar value",
			},
			"parent_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Parent ID",
			},
			"proximity_keywords": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "List of proximity keywords",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"regexes": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "List of regexes",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"regex": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Regex",
						},
						"weight": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Weight",
						},
					},
				},
			},
			"root_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Root type",
			},
			"s3_path": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "S3 path",
			},
			"created_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Created by",
			},
			"updated_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Updated by",
			},
			"updated_at": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Last updated at",
			},
			"is_editable": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Is editable",
			},
		},
	}
}

func dataSourceDataPatternRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var err error
	client := meta.(*pc.Client)

	id := d.Get("pattern_id").(string)

	if id == "" {
		name := d.Get("name").(string)
		id, err = datapattern.Identify(client, name)
		if err != nil {
			if err == pc.ObjectNotFoundError {
				d.SetId("")
				return nil
			}
			return diag.FromErr(err)
		}
	}

	obj, err := datapattern.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	d.SetId(id)
	saveDataPattern(d, obj)

	return nil
}
