package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/cloud/account"
	"github.com/paloaltonetworks/prisma-cloud-go/report"
	"log"
)

func resourceReport() *schema.Resource {
	return &schema.Resource{
		Create: createReport,
		Read:   readReport,
		Update: updateReport,
		Delete: deleteReport,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"report_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Report ID",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Report Name",
			},
			"report_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Report type",
			},
			"cloud_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Cloud type (Required parameter for Compliance report and Cloud Security Assessment report)",
				ValidateFunc: validation.StringInSlice(
					[]string{
						account.TypeAws,
						account.TypeAzure,
						account.TypeGcp,
						account.TypeAlibaba,
						"all",
					},
					false,
				),
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
				Required:    true,
				MinItems:    1,
				MaxItems:    1,
				Description: "Model for report target",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_groups": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "List of cloud account groups",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"accounts": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "List of cloud accounts",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"regions": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "List of regions",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"compliance_standard_ids": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "List of compliance standard IDs",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"resource_groups": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "List of resource groups",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"notify_to": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "List of email addresses to receive notification",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"compression_enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Business unit detailed report compression enabled",
						},
						"download_now": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "True = download now",
						},
						"schedule_enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Report scheduling enabled",
						},
						"schedule": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Recurring report schedule in RRULE format",
						},
						"notification_template_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Notification template id",
						},
						"time_range": timeRangeSchema("resource_report"),
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

func parseReport(d *schema.ResourceData, id string) report.Report {
	tgt := ResourceDataInterfaceMap(d, "target")

	var accountGroups []string
	if accg := tgt["account_groups"]; accg != nil {
		accountGroups = SetToStringSlice(accg.(*schema.Set))
	}

	var accounts []string
	if accs := tgt["accounts"]; accs != nil {
		accounts = SetToStringSlice(accs.(*schema.Set))
	}

	var regions []string
	if rg := tgt["regions"]; rg != nil {
		regions = SetToStringSlice(rg.(*schema.Set))
	}

	var complianceStdIds []string
	if cpl := tgt["compliance_standard_ids"]; cpl != nil {
		complianceStdIds = SetToStringSlice(cpl.(*schema.Set))
	}

	var notifyTo []string
	if noti := tgt["notify_to"]; noti != nil {
		notifyTo = SetToStringSlice(noti.(*schema.Set))
	}

	var resourceGroups []string
	if res := tgt["resource_groups"]; res != nil {
		resourceGroups = SetToStringSlice(res.(*schema.Set))
	}

	time_range := tgt["time_range"].([]interface{})
	tr := ParseTimeRange(time_range[0].(map[string]interface{}))

	ans := report.Report{
		Id:        id,
		Name:      d.Get("name").(string),
		Type:      d.Get("report_type").(string),
		CloudType: d.Get("cloud_type").(string),

		Target: report.Target{
			AccountGroups:          accountGroups,
			Accounts:               accounts,
			Regions:                regions,
			ComplianceStandardIds:  complianceStdIds,
			NotifyTo:               notifyTo,
			ResourceGroups:         resourceGroups,
			TimeRange:              tr,
			CompressionEnabled:     tgt["compression_enabled"].(bool),
			DownloadNow:            tgt["download_now"].(bool),
			Schedule:               tgt["schedule"].(string),
			ScheduleEnabled:        tgt["schedule_enabled"].(bool),
			NotificationTemplateId: tgt["notification_template_id"].(string),
		},
	}

	return ans
}

func saveReport(d *schema.ResourceData, obj report.Report) {
	d.Set("report_id", obj.Id)
	d.Set("name", obj.Name)
	d.Set("report_type", obj.Type)
	d.Set("cloud_type", obj.CloudType)
	d.Set("compliance_standard_id", obj.ComplianceStandardId)
	d.Set("status", obj.Status)
	d.Set("created_on", obj.CreatedOn)
	d.Set("created_by", obj.CreatedBy)
	d.Set("last_modified_on", obj.LastModifiedOn)
	d.Set("last_modified_by", obj.LastModifiedBy)
	d.Set("next_schedule", obj.NextSchedule)
	d.Set("last_scheduled", obj.LastScheduled)
	d.Set("total_instance_count", obj.TotalInstanceCount)

	trgt := ResourceDataInterfaceMap(d, "target")
	tr := trgt["time_range"].([]interface{})

	// Target.
	tgt := map[string]interface{}{
		"account_groups":           obj.Target.AccountGroups,
		"accounts":                 obj.Target.Accounts,
		"regions":                  obj.Target.Regions,
		"compliance_standard_ids":  obj.Target.ComplianceStandardIds,
		"compression_enabled":      obj.Target.CompressionEnabled,
		"download_now":             obj.Target.DownloadNow,
		"notify_to":                obj.Target.NotifyTo,
		"resource_groups":          obj.Target.ResourceGroups,
		"schedule":                 obj.Target.Schedule,
		"schedule_enabled":         obj.Target.ScheduleEnabled,
		"notification_template_id": obj.Target.NotificationTemplateId,
		"time_range":               tr,
	}

	if err := d.Set("target", []interface{}{tgt}); err != nil {
		log.Printf("[WARN] Error setting 'target' for %q: %s", d.Id(), err)
	}

	// Counts.
	cnts := map[string]interface{}{
		"failed":                 obj.Counts.Failed,
		"high_severity_failed":   obj.Counts.HighSeverityFailed,
		"low_severity_failed":    obj.Counts.LowSeverityFailed,
		"medium_severity_failed": obj.Counts.MediumSeverityFailed,
		"passed":                 obj.Counts.Passed,
		"total":                  obj.Counts.Total,
	}

	if err := d.Set("counts", []interface{}{cnts}); err != nil {
		log.Printf("[WARN] Error setting 'counts' for %q: %s", d.Id(), err)
	}
}

func createReport(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	obj := parseReport(d, "")

	if err := report.Create(client, obj); err != nil {
		return err
	}

	PollApiUntilSuccess(func() error {
		_, err := report.Identify(client, obj.Name)
		return err
	})

	id, err := report.Identify(client, obj.Name)
	if err != nil {
		return err
	}

	PollApiUntilSuccess(func() error {
		_, err := report.Get(client, id)
		return err
	})

	d.SetId(id)
	return readReport(d, meta)
}

func readReport(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	id := d.Id()

	obj, err := report.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return err
	}

	saveReport(d, obj)

	return nil
}

func updateReport(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	id := d.Id()
	obj := parseReport(d, id)

	if err := report.Update(client, obj); err != nil {
		return err
	}

	return readReport(d, meta)
}

func deleteReport(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	id := d.Id()

	err := report.Delete(client, id)
	if err != nil {
		if err != pc.ObjectNotFoundError {
			return err
		}
	}

	d.SetId("")
	return nil
}
