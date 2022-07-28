package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"golang.org/x/net/context"
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/compliance/standard"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceComplianceStandard() *schema.Resource {
	return &schema.Resource{
		CreateContext: createComplianceStandard,
		ReadContext:   readComplianceStandard,
		UpdateContext: updateComplianceStandard,
		DeleteContext: deleteComplianceStandard,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"cs_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Compliance standard ID",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description",
			},
			"created_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Created by",
			},
			"created_on": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Created on",
			},
			"last_modified_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Last modified by",
			},
			"last_modified_on": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Last modified on",
			},
			"system_default": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "System default",
			},
			"policies_assigned_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of assigned policies",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Compliance standard name",
			},
			"cloud_types": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Cloud type (determined based on policies assigned)",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func parseComplianceStandard(d *schema.ResourceData, csId string) standard.Standard {
	return standard.Standard{
		Id:          csId,
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}
}

func saveComplianceStandard(d *schema.ResourceData, o standard.Standard) {
	var err error

	d.Set("name", o.Name)
	d.Set("cs_id", o.Id)
	d.Set("description", o.Description)
	d.Set("created_by", o.CreatedBy)
	d.Set("created_on", o.CreatedOn)
	d.Set("last_modified_by", o.LastModifiedBy)
	d.Set("last_modified_on", o.LastModifiedOn)
	d.Set("system_default", o.SystemDefault)
	d.Set("policies_assigned_count", o.PoliciesAssignedCount)
	if err = d.Set("cloud_types", o.CloudTypes); err != nil {
		log.Printf("[WARN] Error setting 'cloud_types' for %q: %s", d.Id(), err)
	}
}

func createComplianceStandard(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	o := parseComplianceStandard(d, "")

	if err := standard.Create(client, o); err != nil {
		return diag.FromErr(err)
	}

	PollApiUntilSuccess(func() error {
		_, err := standard.Identify(client, o.Name)
		return err
	})

	csId, err := standard.Identify(client, o.Name)
	if err != nil {
		return diag.FromErr(err)
	}

	PollApiUntilSuccess(func() error {
		_, err := standard.Get(client, csId)
		return err
	})

	d.SetId(csId)
	return readComplianceStandard(ctx, d, meta)
}

func readComplianceStandard(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	csId := d.Id()

	o, err := standard.Get(client, csId)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	saveComplianceStandard(d, o)

	return nil
}

func updateComplianceStandard(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	csId := d.Id()
	o := parseComplianceStandard(d, csId)

	if err := standard.Update(client, o); err != nil {
		return diag.FromErr(err)
	}

	return readComplianceStandard(ctx, d, meta)
}

func deleteComplianceStandard(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	csId := d.Id()

	err := standard.Delete(client, csId)
	if err != nil {
		if err != pc.ObjectNotFoundError {
			return diag.FromErr(err)
		}
	}

	d.SetId("")
	return nil
}
