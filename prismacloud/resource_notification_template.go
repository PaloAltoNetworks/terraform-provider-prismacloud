package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	notification_template "github.com/paloaltonetworks/prisma-cloud-go/notification-template"
	"golang.org/x/net/context"
	"log"
	"reflect"
)

func resourceNotificationTemplate() *schema.Resource {
	return &schema.Resource{
		CreateContext: createNotificationTemplate,
		ReadContext:   readNotificationTemplate,
		UpdateContext: updateNotificationTemplate,
		DeleteContext: deleteNotificationTemplate,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "id",
			},
			"integration_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "integrationId",
			},
			"created_ts": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "createdTs",
			},
			"last_modified_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "lastModifiedBy",
			},
			"last_modified_ts": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "lastModifiedTs",
			},
			"integration_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "integrationName",
			},
			"customer_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "customerId",
			},
			"integration_type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice(
					[]string{
						notification_template.Email,
						notification_template.ServiceNow,
						notification_template.Jira,
					},
					false,
				),
				Description: "integrationType",
			},
			"created_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "createdBy",
			},
			"module": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "module",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "name",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "enabled",
			},
			"template_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "templateType",
			},
			"template_config": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "List of template_config",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"basic_config": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: getConfigSchema(),
							},
						},
						"open": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: getConfigSchema(),
							},
						},
						"resolved": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: getConfigSchema(),
							},
						},
						"dismissed": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: getConfigSchema(),
							},
						},
						"snoozed": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: getConfigSchema(),
							},
						},
					},
				},
			},
		},
	}
}

func deleteNotificationTemplate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	id := d.Id()
	log.Printf("[INFO]: Deleting Notification Template, Id:%+v\n", id)
	if err := notification_template.Delete(client, id); err != nil {
		if err != pc.ObjectNotFoundError {
			return diag.FromErr(err)
		}
	}
	d.SetId("")
	return nil
}

func updateNotificationTemplate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var err error
	client := meta.(*pc.Client)
	_, o := parseNotificationTemplate(d)
	var _ notification_template.NotificationTemplate
	log.Printf("[INFO]: Updating Notification Template, Id:%+v\n", d.Get("id"))
	if _, err = notification_template.Update(client, o, d.Get("id").(string)); err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}
	return readNotificationTemplate(ctx, d, meta)
	return nil
}

func createNotificationTemplate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var err error
	client := meta.(*pc.Client)
	_, o := parseNotificationTemplate(d)
	var templateRes notification_template.NotificationTemplate
	if templateRes, err = notification_template.Create(client, o); err != nil {
		return diag.FromErr(err)
	}
	log.Printf("[INFO]: Notification Created Successfully, Id:%+v\n", templateRes.Id)
	d.SetId(templateRes.Id)
	return readNotificationTemplate(ctx, d, meta)
	return nil
}

func parseNotificationTemplate(d *schema.ResourceData) (string, notification_template.NotificationTemplateRequest) {

	templateConfigStruct := notification_template.TemplateConfigStruct{}
	templateConfigMap := d.Get("template_config").([]interface{})
	templateConfigMap2 := templateConfigMap[0].(map[string]interface{})
	for templateName, configs := range templateConfigMap2 {
		configsSlice := make([]notification_template.Config, len(configs.([]interface{})))
		for i, config := range configs.([]interface{}) {
			var optionsSlice []notification_template.Option
			if (config.(map[string]interface{})["options"].([]interface{})) != nil {
				for _, option := range config.(map[string]interface{})["options"].([]interface{}) {
					if option != nil {
						tempOption := notification_template.Option{
							Name: option.(map[string]interface{})["name"].(string),
							Key:  option.(map[string]interface{})["key"].(string),
							Id:   option.(map[string]interface{})["id"].(string),
						}
						optionsSlice = append(optionsSlice, tempOption)
					}
				}
			}
			configsSlice[i] = notification_template.Config{
				FieldName:      config.(map[string]interface{})["field_name"].(string),
				Options:        ternaryOperator(len(optionsSlice) > 0, optionsSlice, make([]notification_template.Option, 0)).([]notification_template.Option),
				DisplayName:    config.(map[string]interface{})["display_name"].(string),
				Type:           config.(map[string]interface{})["type"].(string),
				Value:          config.(map[string]interface{})["value"].(string),
				RedlockMapping: config.(map[string]interface{})["redlock_mapping"].(bool),
				Required:       config.(map[string]interface{})["required"].(bool),
				TypeaheadUri:   config.(map[string]interface{})["type_ahead_uri"].(string),
				MaxLength:      config.(map[string]interface{})["max_length"].(int),
			}
		}
		switch configType := templateName; configType {
		case "basic_config":
			templateConfigStruct.BasicConfig = configsSlice
		case "resolved":
			templateConfigStruct.Resolved = configsSlice
		case "dismissed":
			templateConfigStruct.Dismissed = configsSlice
		case "open":
			templateConfigStruct.Open = configsSlice
		case "snoozed":
			templateConfigStruct.Snoozed = configsSlice
		default:
			log.Printf("[WARN]: State mapping not found for: %+v\n, Valid States are: [basic_config, resolved, dismissed , open, snoozed]", configType)
		}
	}

	ntReq := notification_template.NotificationTemplateRequest{
		IntegrationId:   d.Get("integration_id").(string),
		IntegrationType: d.Get("integration_type").(string),
		Name:            d.Get("name").(string),
		Enabled:         d.Get("enabled").(bool),
		TemplateType:    d.Get("template_type").(string),
		TemplateConfig:  templateConfigStruct,
	}
	return ntReq.Name, ntReq
}

func readNotificationTemplate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	id := d.Id()
	obj, err := notification_template.Get(client, id)
	if err != nil {
		if err == pc.ObjectNotFoundError {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}
	saveNotificationTemplate(d, obj)
	return nil
}
func saveNotificationTemplate(d *schema.ResourceData, o notification_template.NotificationTemplate) {
	d.Set("integration_id", o.IntegrationId)
	d.Set("created_ts", o.CreatedTs)
	d.Set("last_modified_by", o.LastModifiedBy)
	d.Set("integration_type", o.IntegrationType)
	d.Set("last_modified_ts", o.LastModifiedTs)
	d.Set("name", o.Name)
	d.Set("created_by", o.CreatedBy)
	d.Set("integration_name", o.IntegrationName)
	d.Set("customer_id", o.CustomerId)
	d.Set("module", o.Module)
	d.Set("enabled", o.Enabled)
	d.Set("template_type", o.TemplateType)

	tempConfigMap := map[string]interface{}{}
	values := reflect.ValueOf(o.TemplateConfig)
	typesOf := values.Type()
	for i := 0; i < values.NumField(); i++ {
		value := values.Field(i).Interface()
		fieldName := typesOf.Field(i).Name
		statusConfigList := value.([]notification_template.Config)
		if len(statusConfigList) > 0 {
			var statusConfigSlice []interface{}

			for _, element := range statusConfigList {
				options := element.Options
				var optionsList []map[string]interface{}
				if len(options) > 0 {
					for _, option := range options {
						optionMap := map[string]interface{}{
							"name": option.Name,
							"key":  option.Key,
							"id":   option.Id,
						}
						optionsList = append(optionsList, optionMap)
					}
				}
				configProps := map[string]interface{}{
					"display_name":    element.DisplayName,
					"field_name":      element.FieldName,
					"type":            element.Type,
					"value":           element.Value,
					"options":         optionsList,
					"max_length":      element.MaxLength,
					"required":        element.Required,
					"type_ahead_uri":  element.TypeaheadUri,
					"redlock_mapping": element.RedlockMapping,
					"alias_field":     element.AliasField,
				}
				statusConfigSlice = append(statusConfigSlice, configProps)
			}
			switch conf := fieldName; conf {
			case "BasicConfig":
				tempConfigMap["basic_config"] = statusConfigSlice
			case "Open":
				tempConfigMap["open"] = statusConfigSlice
			case "Resolved":
				tempConfigMap["resolved"] = statusConfigSlice
			case "Dismissed":
				tempConfigMap["dismissed"] = statusConfigSlice
			case "Snoozed":
				tempConfigMap["snoozed"] = statusConfigSlice
			default:
				log.Printf("[WARN]: State mapping not found for: %+v\n, Valid States are: [BasicConfig, Open Resolved , Snoozed, Dismissed]", conf)
			}
		}

	}
	templateConfigList := []interface{}{tempConfigMap}
	if err := d.Set("template_config", templateConfigList); err != nil {
		log.Printf("[ERROR]: Error setting 'template_config' %+v\n", err)
	}
}
func ternaryOperator(flag bool, obj1 interface{}, obj2 interface{}) interface{} {
	if flag {
		return obj1
	}
	return obj2
}
func getConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"field_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "fieldName",
		},
		"display_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "displayName",
		},
		"redlock_mapping": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "redlockMapping",
		},
		"required": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "required",
		},
		"type_ahead_uri": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "type Ahead URI",
		},
		"type": {
			Type:     schema.TypeString,
			Optional: true,
			ValidateFunc: validation.StringInSlice(
				[]string{
					notification_template.ListType,
					notification_template.TextType,
					notification_template.ArrayType,
					notification_template.BoolType,
					notification_template.IntegerType,
				},
				false,
			),
			Description: "type",
		},
		"value": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "value",
		},
		"alias_field": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "AliasField",
		},
		"max_length": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "MaxLength",
		},
		"options": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "options",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "name",
					},
					"key": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "key",
					},
					"id": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "id",
					},
				},
			},
		},
	}
}
