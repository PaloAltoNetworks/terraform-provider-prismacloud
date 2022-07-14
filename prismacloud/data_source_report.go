package prismacloud

import (
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/report"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceReport() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceReportRead,

		Schema: map[string]*schema.Schema{
			// Input.
			"report_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Report ID",
				AtLeastOneOf: []string{"name", "report_id"},
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Report name",
				AtLeastOneOf: []string{"name", "report_id"},
			},

			// Output.
			"report_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Report type",
			},
			"cloud_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cloud type",
			},
			"compliance_standard_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Compliance Standard ID",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Report status",
			},
			"created_on": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Created on",
			},
			"created_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Created by",
			},
			"last_modified_on": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Last modified on",
			},
			"last_modified_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Last modified by",
			},
			"next_schedule": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Next schedule",
			},
			"last_scheduled": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Last scheduled",
			},
			"total_instance_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Total instance count",
			},
			"target": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Model for report target",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_groups": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "List of cloud account groups",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"accounts": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "List of cloud accounts",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"regions": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "List of regions",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"compliance_standard_ids": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "List of compliance standard IDs",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"resource_groups": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "List of resource groups",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"notify_to": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "List of email addresses to receive notification",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"compression_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Business unit detailed report compression enabled",
						},
						"download_now": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "True = download now",
						},
						"schedule_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Report scheduling enabled",
						},
						"schedule": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Recurring report schedule in RRULE format",
						},
						"notification_template_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Notification template id",
						},
						"time_range": timeRangeSchema("data_source_report"),
					},
				},
			},
			"counts": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Model for compliance aggregate count",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"failed": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Failed",
						},
						"high_severity_failed": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of high-severity failures",
						},
						"low_severity_failed": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of low-severity failures",
						},
						"medium_severity_failed": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of medium-severity failures",
						},
						"passed": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Passed",
						},
						"total": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Total",
						},
					},
				},
			},
		},
	}
}

func dataSourceReportRead(d *schema.ResourceData, meta interface{}) error {
	var err error
	client := meta.(*pc.Client)

	id := d.Get("report_id").(string)

	if id == "" {
		name := d.Get("name").(string)
		id, err = report.Identify(client, name)
		if err != nil {
			if err == pc.ObjectNotFoundError {
				d.SetId("")
				return nil
			}
			return err
		}
	}

	obj, err := report.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return err
	}

	d.SetId(id)
	saveReport(d, obj)

	return nil
}
