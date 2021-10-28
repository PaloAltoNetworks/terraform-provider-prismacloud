package prismacloud

import (
	"log"

	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"github.com/paloaltonetworks/prisma-cloud-go/rql/search"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRqlSearch() *schema.Resource {
	return &schema.Resource{
		Create: createUpdateRqlSearch,
		Read:   readRqlSearch,
		Update: createUpdateRqlSearch,
		Delete: deleteRqlSearch,
		Schema: map[string]*schema.Schema{
			// Input.
			"search_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The search type",
				Default:     "config",
				ValidateFunc: validation.StringInSlice([]string{
					"config",
					"network",
					"event",
					"iam",
				}, false),
				ForceNew: true,
			},
			"query": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The RQL search to perform",
				ForceNew:    true,
			},
			"time_range": timeRangeSchema("resource_rql_search"),
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Limit results",
				Default:     10,
			},
			// Output.
			"group_by": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Group by",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"search_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The search ID",
			},
			"cloud_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The cloud type",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The search name",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The description",
			},
			"saved": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Is search saved",
			},
			"config_data": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of config data structs",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"state_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The state ID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL",
						},
					},
				},
			},
			"event_data": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of event data structs",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account",
						},
						"region_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Region ID",
						},
						"region_api_identifier": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Region API identifier",
						},
					},
				},
			},
			"network_data": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of network data structs",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account",
						},
						"region_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Region ID",
						},
						"account_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "account_name",
						},
					},
				},
			},
			"iam_data": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of iam data structs",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accessed_resources_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Accessed resource count",
						},
						"dest_cloud_account": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Destination cloud account",
						},
						"dest_cloud_region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Destination cloud region",
						},
						"dest_cloud_resource_rrn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Destination cloud resource RRN",
						},
						"dest_cloud_service_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Destination cloud service name",
						},
						"dest_cloud_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Destination cloud type",
						},
						"dest_resource_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Destination cloud resource id",
						},
						"dest_resource_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Destination cloud resource name",
						},
						"dest_resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Destination cloud resource type",
						},
						"effective_action_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Effective action name",
						},
						"granted_by_cloud_entity_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Granted by cloud entity id",
						},
						"granted_by_cloud_entity_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Granted by cloud entity name",
						},
						"granted_by_cloud_entity_rrn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Granted by cloud entity rrn",
						},
						"granted_by_cloud_entity_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Granted by cloud entity type",
						},
						"granted_by_cloud_policy_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Granted by cloud policy id",
						},
						"granted_by_cloud_policy_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Granted by cloud policy name",
						},
						"granted_by_cloud_policy_rrn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Granted by cloud policy rrn",
						},
						"granted_by_cloud_policy_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Granted by cloud policy type",
						},
						"granted_by_cloud_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Granted by cloud type",
						},
						"message_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Message id",
						},
						"is_wild_card_dest_cloud_resource_name": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Is destination cloud resource name a wildcard",
						},
						"last_access_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Last access date",
						},
						"source_cloud_account": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Source cloud account",
						},
						"source_cloud_region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Source cloud region",
						},
						"source_cloud_resource_rrn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Source cloud resource rrn",
						},
						"source_cloud_service_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Source cloud service name",
						},
						"source_cloud_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Source cloud type",
						},
						"source_idp_domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Source IDP domain",
						},
						"source_idp_email": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Source IDP email",
						},
						"source_idp_group": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Source IDP group",
						},
						"source_idp_rrn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Source IDP rrn",
						},
						"source_idp_service": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Source IDP service",
						},
						"source_idp_user_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Source IDP user name",
						},
						"source_public": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Is source public",
						},
						"source_resource_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Source cloud resource id",
						},
						"source_resource_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Source cloud resource name",
						},
						"source_resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Source cloud resource type",
						},
						"exceptions": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Permission exception list",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"message_code": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Message code",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func createUpdateRqlSearch(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	query := d.Get("query").(string)
	limit := d.Get("limit").(int)
	searchType := d.Get("search_type").(string)
	tr := ParseTimeRange(ResourceDataInterfaceMap(d, "time_range"))
	var id string

	if d.Id() != "" {
		return readRqlSearch(d, meta)
	}

	switch searchType {
	case "config":
		req := search.ConfigRequest{
			Query:     query,
			Limit:     limit,
			TimeRange: tr,
		}

		resp, err := search.ConfigSearch(client, req)
		if err != nil {
			return err
		}

		PollApiUntilSuccess(func() error {
			r := search.ConfigRequest{
				Id:        resp.Id,
				Query:     query,
				Limit:     limit,
				TimeRange: tr,
			}
			_, err := search.ConfigSearch(client, r)
			return err
		})

		id = buildRqlSearchId(searchType, query, resp.Id)
	case "network":
		req := search.NetworkRequest{
			Query:     query,
			Limit:     limit,
			TimeRange: tr,
		}

		resp, err := search.NetworkSearch(client, req)
		if err != nil {
			return err
		}

		PollApiUntilSuccess(func() error {
			r := search.NetworkRequest{
				Id:        resp.Id,
				Query:     query,
				Limit:     limit,
				TimeRange: tr,
			}
			_, err := search.NetworkSearch(client, r)
			return err
		})

		id = buildRqlSearchId(searchType, query, resp.Id)
	case "event":
		req := search.EventRequest{
			Query:     query,
			Limit:     limit,
			TimeRange: tr,
		}

		resp, err := search.EventSearch(client, req)
		if err != nil {
			return err
		}

		PollApiUntilSuccess(func() error {
			r := search.EventRequest{
				Id:        resp.Id,
				Query:     query,
				Limit:     limit,
				TimeRange: tr,
			}
			_, err := search.EventSearch(client, r)
			return err
		})

		id = buildRqlSearchId(searchType, query, resp.Id)
	case "iam":
		req := search.IamRequest{
			Query: query,
			Limit: limit,
		}

		resp, err := search.IamSearch(client, req)
		if err != nil {
			return err
		}

		PollApiUntilSuccess(func() error {
			r := search.IamRequest{
				Id:    resp.Id,
				Query: query,
				Limit: limit,
			}
			_, err := search.IamSearch(client, r)
			return err
		})

		id = buildRqlSearchId(searchType, query, resp.Id)
	}

	d.SetId(id)

	return readRqlSearch(d, meta)
}

func readRqlSearch(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pc.Client)
	searchType, query, searchId := parseRqlSearchId(d.Id())
	limit := d.Get("limit").(int)
	tr := ParseTimeRange(ResourceDataInterfaceMap(d, "time_range"))

	switch searchType {
	case "config":
		req := search.ConfigRequest{
			Id:        searchId,
			Query:     query,
			Limit:     limit,
			TimeRange: tr,
		}

		resp, err := search.ConfigSearch(client, req)
		if err != nil {
			return err
		}

		if err = d.Set("group_by", resp.GroupBy); err != nil {
			log.Printf("[WARN] Error setting 'group_by' for %q: %s", d.Id(), err)
		}
		d.Set("search_id", resp.Id)
		d.Set("cloud_type", resp.CloudType)
		d.Set("name", resp.Name)
		d.Set("description", resp.Description)
		d.Set("event_data", nil)
		d.Set("network_data", nil)
		d.Set("iam_data", nil)

		if len(resp.Data.Items) == 0 {
			d.Set("config_data", nil)
		} else {
			list := make([]interface{}, 0, len(resp.Data.Items))
			for _, x := range resp.Data.Items {
				list = append(list, map[string]interface{}{
					"state_id": x.StateId,
					"name":     x.Name,
					"url":      x.Url,
				})
			}

			if err = d.Set("config_data", list); err != nil {
				log.Printf("[WARN] Error setting 'config_data' for %q: %s", d.Id(), err)
			}
		}
	case "network":
		req := search.NetworkRequest{
			Id:        searchId,
			Query:     query,
			Limit:     limit,
			TimeRange: tr,
		}

		resp, err := search.NetworkSearch(client, req)
		if err != nil {
			return err
		}

		if err = d.Set("group_by", resp.GroupBy); err != nil {
			log.Printf("[WARN] Error setting 'group_by' for %q: %s", d.Id(), err)
		}
		d.Set("search_id", resp.Id)
		d.Set("cloud_type", resp.CloudType)
		d.Set("name", resp.Name)
		d.Set("description", resp.Description)
		d.Set("event_data", nil)
		d.Set("config_data", nil)
		d.Set("iam_data", nil)

		if len(resp.Data.Items) == 0 {
			d.Set("network_data", nil)
		} else {
			list := make([]interface{}, 0, len(resp.Data.Items))
			for _, x := range resp.Data.Items {
				list = append(list, map[string]interface{}{
					"account":      x.Account,
					"region_id":    x.RegionId,
					"account_name": x.AccountName,
				})
			}

			if err = d.Set("network_data", list); err != nil {
				log.Printf("[WARN] Error setting 'network_data' for %q: %s", d.Id(), err)
			}
		}
	case "event":
		req := search.EventRequest{
			Id:        searchId,
			Query:     query,
			Limit:     limit,
			TimeRange: tr,
		}
		resp, err := search.EventSearch(client, req)
		if err != nil {
			return err
		}
		if err = d.Set("group_by", resp.GroupBy); err != nil {
			log.Printf("[WARN] Error setting 'group_by' for %q: %s", d.Id(), err)
		}
		d.Set("search_id", resp.Id)
		d.Set("cloud_type", resp.CloudType)
		d.Set("name", resp.Name)
		d.Set("description", resp.Description)
		d.Set("config_data", nil)
		d.Set("network_data", nil)
		d.Set("iam_data", nil)

		if len(resp.Data.Items) == 0 {
			d.Set("event_data", nil)
		} else {
			list := make([]interface{}, 0, len(resp.Data.Items))
			for _, x := range resp.Data.Items {
				list = append(list, map[string]interface{}{
					"account":               x.Account,
					"region_id":             x.RegionId,
					"region_api_identifier": x.RegionApiIdentifier,
				})
			}
			if err = d.Set("event_data", list); err != nil {
				log.Printf("[WARN] Error setting 'event_data' for %q: %s", d.Id(), err)
			}
		}
	case "iam":
		req := search.IamRequest{
			Id:    searchId,
			Query: query,
			Limit: limit,
		}

		resp, err := search.IamSearch(client, req)
		if err != nil {
			return err
		}

		d.Set("search_id", resp.Id)
		d.Set("name", resp.Name)
		d.Set("description", resp.Description)
		d.Set("saved", resp.Saved)
		d.Set("config_data", nil)
		d.Set("network_data", nil)
		d.Set("event_data", nil)

		tr := flattenTimeRange(resp.TimeRange)
		if err = d.Set("time_range", tr); err != nil {
			log.Printf("[WARN] Error setting 'time_range' for %q: %s", d.Id(), err)
		}
		if len(resp.Data.Items) == 0 {
			d.Set("iam_data", nil)
		} else {
			list := make([]interface{}, 0, len(resp.Data.Items))
			for _, x := range resp.Data.Items {
				list = append(list, map[string]interface{}{
					"accessed_resources_count":              x.AccessedResourcesCount,
					"dest_cloud_account":                    x.DestCloudAccount,
					"dest_cloud_region":                     x.DestCloudRegion,
					"dest_cloud_resource_rrn":               x.DestCloudResourceRrn,
					"dest_cloud_service_name":               x.DestCloudServiceName,
					"dest_cloud_type":                       x.DestCloudType,
					"dest_resource_id":                      x.DestResourceId,
					"dest_resource_name":                    x.DestResourceName,
					"dest_resource_type":                    x.DestResourceType,
					"effective_action_name":                 x.EffectiveActionName,
					"granted_by_cloud_entity_id":            x.GrantedByCloudEntityId,
					"granted_by_cloud_entity_name":          x.GrantedByCloudEntityName,
					"granted_by_cloud_entity_rrn":           x.GrantedByCloudEntityRrn,
					"granted_by_cloud_entity_type":          x.GrantedByCloudEntityType,
					"granted_by_cloud_policy_id":            x.GrantedByCloudPolicyId,
					"granted_by_cloud_policy_name":          x.GrantedByCloudPolicyName,
					"granted_by_cloud_policy_rrn":           x.GrantedByCloudPolicyRrn,
					"granted_by_cloud_policy_type":          x.GrantedByCloudPolicyType,
					"granted_by_cloud_type":                 x.GrantedByCloudType,
					"message_id":                            x.MessageId,
					"is_wild_card_dest_cloud_resource_name": x.IsWildCardDestCloudResourceName,
					"last_access_date":                      x.LastAccessDate,
					"source_cloud_account":                  x.SourceCloudAccount,
					"source_cloud_region":                   x.SourceCloudRegion,
					"source_cloud_resource_rrn":             x.SourceCloudResourceRrn,
					"source_cloud_service_name":             x.SourceCloudServiceName,
					"source_cloud_type":                     x.SourceCloudType,
					"source_idp_domain":                     x.SourceIdpDomain,
					"source_idp_email":                      x.SourceIdpEmail,
					"source_idp_group":                      x.SourceIdpGroup,
					"source_idp_rrn":                        x.SourceIdpRrn,
					"source_idp_service":                    x.SourceIdpService,
					"source_idp_user_name":                  x.SourceIdpRrn,
					"source_public":                         x.SourcePublic,
					"source_resource_id":                    x.SourceResourceId,
					"source_resource_name":                  x.SourceResourceName,
					"source_resource_type":                  x.SourceResourceType,
				})

				excList := make([]interface{}, 0, len(x.Exceptions))
				for _, exc := range x.Exceptions {
					excList = append(excList, map[string]interface{}{
						"message_code": exc.MessageCode,
					})
				}
				list = append(list, map[string]interface{}{
					"exceptions": excList,
				})
			}

			if err = d.Set("iam_data", list); err != nil {
				log.Printf("[WARN] Error setting 'iam_data' for %q: %s", d.Id(), err)
			}

		}
	}
	return nil
}
func deleteRqlSearch(d *schema.ResourceData, meta interface{}) error {
	// There is no way to delete a search, so this is a no-op.
	return nil
}

// Id functions.
func buildRqlSearchId(a, b, c string) string {
	res := Base64Encode([]interface{}{a, b, c})
	return res
}

func parseRqlSearchId(v string) (string, string, string) {
	t := Base64Decode(v)
	return t[0], t[1], t[2]
}
