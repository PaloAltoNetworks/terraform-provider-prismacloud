package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"golang.org/x/net/context"
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/data-security/datapattern"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDataPattern() *schema.Resource {
	return &schema.Resource{
		CreateContext: createDataPattern,
		ReadContext:   readDataPattern,
		UpdateContext: updateDataPattern,
		DeleteContext: deleteDataPattern,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"pattern_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Pattern ID",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Pattern name",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Pattern description",
			},
			"mode": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Mode - predefined or custom",
			},
			"detection_technique": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Detection technique",
				Default:     "regex",
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
				Optional:    true,
				Description: "List of proximity keywords",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"regexes": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "List of regexes",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"regex": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Regex",
						},
						"weight": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Weight",
							Default:     1,
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

func parseDataPattern(d *schema.ResourceData, id string) datapattern.Pattern {
	ans := datapattern.Pattern{
		Id:                 id,
		Name:               d.Get("name").(string),
		Description:        d.Get("description").(string),
		DetectionTechnique: d.Get("detection_technique").(string),
		ProximityKeywords:  SetToStringSlice(d.Get("proximity_keywords").(*schema.Set)),
	}
	regexes := d.Get("regexes").(*schema.Set).List()
	ans.Regexes = make([]datapattern.RegexInfo, 0, len(regexes))
	for _, regx := range regexes {
		reg := regx.(map[string]interface{})
		ans.Regexes = append(ans.Regexes, datapattern.RegexInfo{
			Regex:  reg["regex"].(string),
			Weight: reg["weight"].(int),
		})
	}

	return ans
}

func saveDataPattern(d *schema.ResourceData, obj datapattern.Pattern) {
	d.Set("pattern_id", obj.Id)
	d.Set("name", obj.Name)
	d.Set("description", obj.Description)
	d.Set("mode", obj.Mode)
	d.Set("detection_technique", obj.DetectionTechnique)
	d.Set("entity", obj.Entity)
	d.Set("grammar", obj.Grammar)
	d.Set("parent_id", obj.ParentId)
	d.Set("root_type", obj.RootType)
	d.Set("s3_path", obj.S3Path)
	d.Set("created_by", obj.CreatedBy)
	d.Set("updated_by", obj.UpdatedBy)
	d.Set("updated_at", obj.UpdatedAt)
	d.Set("is_editable", obj.IsEditable)

	if err := d.Set("proximity_keywords", StringSliceToSet(obj.ProximityKeywords)); err != nil {
		log.Printf("[WARN] Error setting 'proximity_keywords' for %q: %s", d.Id(), err)
	}

	if len(obj.Regexes) == 0 {
		d.Set("regexes", nil)
		return
	}

	regexes := make([]interface{}, 0, len(obj.Regexes))
	for _, reg := range obj.Regexes {
		regexes = append(regexes, map[string]interface{}{
			"regex":  reg.Regex,
			"weight": reg.Weight,
		})
	}
	if err := d.Set("regexes", regexes); err != nil {
		log.Printf("[WARN] Error setting 'regexes' for %q: %s", d.Id(), err)
	}
}

func createDataPattern(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	obj := parseDataPattern(d, "")

	if err := datapattern.Create(client, obj); err != nil {
		return diag.FromErr(err)
	}

	PollApiUntilSuccess(func() error {
		_, err := datapattern.Identify(client, obj.Name)
		return err
	})

	id, err := datapattern.Identify(client, obj.Name)
	if err != nil {
		return diag.FromErr(err)
	}

	PollApiUntilSuccess(func() error {
		_, err := datapattern.Get(client, id)
		return err
	})

	d.SetId(id)
	return readDataPattern(ctx, d, meta)
}

func readDataPattern(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	id := d.Id()

	obj, err := datapattern.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	saveDataPattern(d, obj)

	return nil
}

func updateDataPattern(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	id := d.Id()
	obj := parseDataPattern(d, id)

	if err := datapattern.Update(client, obj); err != nil {
		return diag.FromErr(err)
	}

	return readDataPattern(ctx, d, meta)
}

func deleteDataPattern(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	id := d.Id()

	err := datapattern.Delete(client, id)
	if err != nil {
		if err != pc.ObjectNotFoundError {
			return diag.FromErr(err)
		}
	}

	d.SetId("")
	return nil
}
