package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/trusted-alert-ip"
	"golang.org/x/net/context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTrustedAlertIp() *schema.Resource {
	return &schema.Resource{
		CreateContext: createTrustedAlertIp,
		ReadContext:   readTrustedAlertIp,
		UpdateContext: updateTrustedAlertIp,
		DeleteContext: deleteTrustedAlertIp,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"uuid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Trusted alert ip ID",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Trusted alert ip name",
			},
			"cidrs": {
				Type:        schema.TypeSet, //data source of both working fine with list
				Optional:    true,
				Description: "CIDRs",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cidr": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "CIDR",
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Description",
						},
						"uuid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "UUID",
						},
						"created_on": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Created on",
						},
					},
				},
			},
			"cidr_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Associated cloud account type",
			},
		},
	}
}

func parseTrustedAlertIp(d *schema.ResourceData, id string) trustedalertip.TrustedAlertIP {
	ans := trustedalertip.TrustedAlertIP{
		UUID:      id,
		Name:      d.Get("name").(string),
		CidrCount: d.Get("cidr_count").(int),
	}
	cidrs := d.Get("cidrs").(*schema.Set).List()
	ans.CIDRS = make([]trustedalertip.CIDRS, 0, len(cidrs))
	for _, tlex := range cidrs {
		tle := tlex.(map[string]interface{})
		ans.CIDRS = append(ans.CIDRS, trustedalertip.CIDRS{
			CIDR:        tle["cidr"].(string),
			UUID:        tle["uuid"].(string),
			Description: tle["description"].(string),
			CreatedOn:   tle["created_on"].(int),
		})
	}
	return ans
}

func saveTrustedAlertIp(d *schema.ResourceData, obj trustedalertip.TrustedAlertIP) {
	d.Set("uuid", obj.UUID)
	d.Set("name", obj.Name)
	d.Set("cidr_count", obj.CidrCount)

	cidrs := make([]interface{}, 0, len(obj.CIDRS))
	if len(obj.CIDRS) != 0 {
		for _, cidr := range obj.CIDRS {
			cidrs = append(cidrs, map[string]interface{}{
				"cidr":        cidr.CIDR,
				"created_on":  cidr.CreatedOn,
				"description": cidr.Description,
				"uuid":        cidr.UUID,
			})
		}
	}
	if err := d.Set("cidrs", cidrs); err != nil {
		log.Printf("[WARN] Error setting 'cidrs' for %q: %s", d.Id(), err)
	}
}

func createTrustedAlertIp(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	obj := parseTrustedAlertIp(d, "")

	if _, err := trustedalertip.Create(client, obj); err != nil {
		return diag.FromErr(err)
	}

	PollApiUntilSuccess(func() error {
		_, err := trustedalertip.Identify(client, obj.Name)
		return err
	})

	id, err := trustedalertip.Identify(client, obj.Name)
	if err != nil {
		return diag.FromErr(err)
	}

	PollApiUntilSuccess(func() error {
		_, err := trustedalertip.Get(client, id)
		return err
	})

	d.SetId(id)
	return readTrustedAlertIp(ctx, d, meta)
}

func readTrustedAlertIp(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	id := d.Id()

	obj, err := trustedalertip.Get(client, id)
	if err != nil {
		if err == pc.AccountGroupNotFoundError { //changehere
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	saveTrustedAlertIp(d, obj)

	return nil
}

func updateTrustedAlertIp(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	obj := parseTrustedAlertIp(d, d.Id())

	if _, err := trustedalertip.Update(client, obj); err != nil {
		return diag.FromErr(err)
	}

	return readTrustedAlertIp(ctx, d, meta)
}

func deleteTrustedAlertIp(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	id := d.Id()

	if err := trustedalertip.Delete(client, id); err != nil {
		if err != pc.ObjectNotFoundError {
			return diag.FromErr(err)
		}
	}

	d.SetId("")
	return nil
}
