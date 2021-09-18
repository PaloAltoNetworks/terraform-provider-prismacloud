package prismacloud

import (
	"fmt"
	"log"

	"github.com/paloaltonetworks/prisma-cloud-go/timerange"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func totalSchema(desc string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeInt,
		Computed:    true,
		Description: fmt.Sprintf("Total number of %s", desc),
	}
}

/*
This function may need to be revisited..  Not happy with the "style" param
and it makes an assumption that the param this time range is being saved to
is `time_range`.
*/
func timeRangeSchema(style string) *schema.Schema {
	absolute_resource := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"start": {
				Type:        schema.TypeInt,
				Description: "Start time",
			},
			"end": {
				Type:        schema.TypeInt,
				Description: "End time",
			},
		},
	}

	relative_resource := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"amount": {
				Type:        schema.TypeInt,
				Description: "The time number",
			},
			"unit": {
				Type:        schema.TypeString,
				Description: "The time unit",
				ValidateFunc: validation.StringInSlice(
					[]string{
						timerange.Hour,
						timerange.Day,
						timerange.Week,
						timerange.Month,
						timerange.Year,
					},
					false,
				),
			},
			"relative_time_type": {
				Type:        schema.TypeString,
				Description: "Relative time type",
			},
		},
	}

	to_now_resource := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"unit": {
				Type:        schema.TypeString,
				Description: "The time unit",
				ValidateFunc: validation.StringInSlice(
					[]string{
						timerange.Login,
						timerange.Epoch,
						timerange.Day,
						timerange.Week,
						timerange.Month,
						timerange.Year,
					},
					false,
				),
			},
		},
	}

	model := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"absolute": {
				Type:        schema.TypeList,
				Description: "An absolute time range",
				MaxItems:    1,
				Elem:        absolute_resource,
			},
			"relative": {
				Type:        schema.TypeList,
				Description: "Relative time range",
				MaxItems:    1,
				Elem:        relative_resource,
			},
			"to_now": {
				Type:        schema.TypeList,
				Description: "From some time in the past to now",
				MaxItems:    1,
				Elem:        to_now_resource,
			},
		},
	}

	ans := &schema.Schema{
		Type:        schema.TypeList,
		Description: "The time range spec",
		MaxItems:    1,
		Elem:        model,
	}

	switch style {
	case "data_source_rql_historic_search":
		ans.Computed = true

		model.Schema["absolute"].Computed = true
		absolute_resource.Schema["start"].Computed = true
		absolute_resource.Schema["end"].Computed = true

		model.Schema["relative"].Computed = true
		relative_resource.Schema["amount"].Computed = true
		relative_resource.Schema["unit"].Computed = true
		relative_resource.Schema["unit"].ValidateFunc = nil
		relative_resource.Schema["relative_time_type"].Computed = true

		model.Schema["to_now"].Computed = true
		to_now_resource.Schema["unit"].Computed = true
		to_now_resource.Schema["unit"].ValidateFunc = nil
	case "resource_saved_search":
		ans.ForceNew = true
		fallthrough
	default:
		ans.Required = true

		model.Schema["absolute"].Optional = true
		absolute_resource.Schema["start"].Required = true
		absolute_resource.Schema["end"].Required = true

		model.Schema["relative"].Optional = true
		relative_resource.Schema["amount"].Required = true
		relative_resource.Schema["unit"].Required = true
		delete(relative_resource.Schema, "relative_time_type")

		model.Schema["to_now"].Optional = true
		to_now_resource.Schema["unit"].Required = true
	}

	switch style {
	case "resource_report", "data_source_report":
		model.Schema["absolute"].ConflictsWith = []string{
			"target.time_range.relative",
			"target.time_range.to_now",
		}
		model.Schema["relative"].ConflictsWith = []string{
			"target.time_range.absolute",
			"target.time_range.to_now",
		}
		model.Schema["to_now"].ConflictsWith = []string{
			"target.time_range.absolute",
			"target.time_range.relative",
		}

	default:
		model.Schema["absolute"].ConflictsWith = []string{
			"time_range.relative",
			"time_range.to_now",
		}
		model.Schema["relative"].ConflictsWith = []string{
			"time_range.absolute",
			"time_range.to_now",
		}
		model.Schema["to_now"].ConflictsWith = []string{
			"time_range.absolute",
			"time_range.relative",
		}
	}

	return ans
}

func flattenTimeRange(t timerange.TimeRange) []interface{} {
	val := map[string]interface{}{
		"absolute": nil,
		"relative": nil,
		"to_now":   nil,
	}

	if err := t.SetValue(); err != nil {
		log.Printf("[WARN] time range SetValue failed: %s", err)
	}

	switch v := t.Value.(type) {
	case timerange.Absolute:
		val["absolute"] = []interface{}{map[string]interface{}{
			"start": v.Start,
			"end":   v.End,
		}}
	case timerange.Relative:
		val["relative"] = []interface{}{map[string]interface{}{
			"amount":             v.Amount,
			"unit":               v.Unit,
			"relative_time_type": t.RelativeTimeType,
		}}
	case timerange.ToNow:
		val["to_now"] = []interface{}{map[string]interface{}{
			"unit": v.Unit,
		}}
	}

	return []interface{}{val}
}
