package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"golang.org/x/net/context"
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-go"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/paloaltonetworks/prisma-cloud-go/anomalySettings/anomalyTrustedList"
	"strconv"
)

func resourceAnomalyTrustedList() *schema.Resource {
	return &schema.Resource{
		CreateContext: createAnomalyTrustedList,
		ReadContext:   readAnomalyTrustedList,
		UpdateContext: updateAnomalyTrustedList,
		DeleteContext: deleteAnomalyTrustedList,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"atl_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Anomaly Trusted List ID",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Anomaly Trusted List name",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Reason for trusted listing",
			},
			"trusted_list_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Anomaly Trusted List type",
				ValidateFunc: validation.StringInSlice(
					[]string{
						"ip",
						"resource",
						"image",
						"tag",
						"service",
						"port",
						"subject",
						"domain",
						"protocol",
					},
					false,
				),
			},
			"account_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "any",
				Description: "Anomaly Trusted List account id.",
			},
			"applicable_policies": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "Applicable Policies",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"vpc": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "any",
				Description: "VPC",
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
			"trusted_list_entries": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "List of network anomalies in the trusted list",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tag_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Tag key",
						},
						"tag_value": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Tag value",
						},
						"image_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Image ID",
						},
						"ip_cidr": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Ip CIDR",
						},
						"port": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Port",
						},
						"resource_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							Description: "Resource ID",
						},
						"service": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Service",
						},
						"subject": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Subject",
						},
						"domain": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Domain",
						},
						"protocol": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Protocol",
						},
					},
				},
			},
		},
	}
}

func parseAnomalyTrustedList(d *schema.ResourceData, Id string) anomalyTrustedList.AnomalyTrustedList {
	ans := anomalyTrustedList.AnomalyTrustedList{
		Atl_Id:             d.Get("atl_id").(int), //***
		Name:               d.Get("name").(string),
		Description:        d.Get("description").(string),
		TrustedListType:    d.Get("trusted_list_type").(string),
		AccountId:          d.Get("account_id").(string),
		VPC:                d.Get("vpc").(string),
		ApplicablePolicies: SetToStringSlice(d.Get("applicable_policies").(*schema.Set)),
	}
	trustedListEntries := d.Get("trusted_list_entries").(*schema.Set).List()
	ans.TrustedListEntries = make([]anomalyTrustedList.TrustedListEntry, 0, len(trustedListEntries))
	for _, tlex := range trustedListEntries {
		tle := tlex.(map[string]interface{})
		ans.TrustedListEntries = append(ans.TrustedListEntries, anomalyTrustedList.TrustedListEntry{
			TagKey:     tle["tag_key"].(string),
			TagValue:   tle["tag_value"].(string),
			ImageID:    tle["image_id"].(string),
			IpCIDR:     tle["ip_cidr"].(string),
			Port:       tle["port"].(string),
			ResourceID: tle["resource_id"].(string),
			Service:    tle["service"].(string),
			Subject:    tle["subject"].(string),
			Domain:     tle["domain"].(string),
			Protocol:   tle["protocol"].(string),
		})
	}
	return ans
}

func saveAnomalyTrustedList(d *schema.ResourceData, o anomalyTrustedList.AnomalyTrustedList) {
	d.Set("atl_id", o.Atl_Id)
	d.Set("name", o.Name)
	d.Set("description", o.Description)
	d.Set("trusted_list_type", o.TrustedListType)
	d.Set("account_id", o.AccountId)
	d.Set("vpc", o.VPC)
	d.Set("created_by", o.CreatedBy)
	d.Set("created_on", o.CreatedOn)

	if err := d.Set("applicable_policies", StringSliceToSet(o.ApplicablePolicies)); err != nil {
		log.Printf("[WARN] Error setting 'applicable_policies' for %q: %s", d.Id(), err)
	}

	trustedListEntries := make([]interface{}, 0, len(o.TrustedListEntries))
	if len(o.TrustedListEntries) != 0 {
		for _, tle := range o.TrustedListEntries {
			trustedListEntries = append(trustedListEntries, map[string]interface{}{
				"tag_value":   tle.TagValue,
				"image_id":    tle.ImageID,
				"ip_cidr":     tle.IpCIDR,
				"port":        tle.Port,
				"tag_key":     tle.TagKey,
				"resource_id": tle.ResourceID,
				"service":     tle.Service,
				"subject":     tle.Subject,
				"domain":      tle.Domain,
				"protocol":    tle.Protocol,
			})
		}
	}
	if err := d.Set("trusted_list_entries", trustedListEntries); err != nil {
		log.Printf("[WARN] Error setting 'trusted_list_entries' for %q: %s", d.Id(), err)
	}

}

func createAnomalyTrustedList(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	o := parseAnomalyTrustedList(d, "")

	var res int
	res, err := anomalyTrustedList.Create(client, o)
	if err != nil {
		return diag.FromErr(err)
	}

	PollApiUntilSuccess(func() error {
		_, err := anomalyTrustedList.Identify(client, strconv.Itoa(res))
		return err
	})

	id, err := anomalyTrustedList.Identify(client, strconv.Itoa(res))
	if err != nil {
		return diag.FromErr(err)
	}

	PollApiUntilSuccess(func() error {
		_, err := anomalyTrustedList.Get(client, id)
		return err
	})

	d.SetId(id)
	return readAnomalyTrustedList(ctx, d, meta)
}

func readAnomalyTrustedList(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	id := d.Id()

	ans, err := anomalyTrustedList.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	saveAnomalyTrustedList(d, ans)
	return nil
}

func updateAnomalyTrustedList(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	id := d.Id()
	o := parseAnomalyTrustedList(d, id)

	if _, err := anomalyTrustedList.Update(client, o); err != nil {
		return diag.FromErr(err)
	}
	return readAnomalyTrustedList(ctx, d, meta)
}

func deleteAnomalyTrustedList(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	csId := d.Id()

	err := anomalyTrustedList.Delete(client, csId)
	if err != nil {
		if err != pc.ObjectNotFoundError {
			return diag.FromErr(err)
		}
	}

	d.SetId("")
	return nil
}
