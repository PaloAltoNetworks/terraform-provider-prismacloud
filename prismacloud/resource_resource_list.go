package prismacloud

import (
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	resource_list "github.com/paloaltonetworks/prisma-cloud-go/resource-list"
	"golang.org/x/net/context"
	"log"
)

func resourceResourceList() *schema.Resource {
	return &schema.Resource{
		CreateContext: createResourceList,
		ReadContext:   readResourceList,
		UpdateContext: updateResourceList,
		DeleteContext: deleteResourceList,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource list ID",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Resource list description",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Resource list name",
			},
			"resource_list_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Resource list type",
				ValidateFunc: validation.StringInSlice(
					[]string{
						resource_list.TypeTags,
						resource_list.TypeAzureResourceGroups,
						resource_list.TypeComputeAccessGroups,
					},
					false,
				),
			},
			"last_modified_ts": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Resource list last modified timestamp",
			},
			"last_modified_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource list last modified by",
			},
			"members": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Resource list members",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tags": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of key:value pairs of tag members",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"value": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "",
									},
								},
							},
						},
						"azure_resource_groups": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of Azure resource groups part of the resource list",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"compute_access_groups": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Members when resource list type = compute access group",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hosts": {
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Description: "List of pattern strings to define hosts in resource list",
										MinItems:    1,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"app_id": {
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Description: "List of pattern strings to define app_id in resource list",
										MinItems:    1,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"images": {
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Description: "List of pattern strings to define images in resource list",
										MinItems:    1,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"labels": {
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Description: "List of pattern strings to define labels in resource list",
										MinItems:    1,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"clusters": {
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Description: "List of pattern strings to define clusters in resource list",
										MinItems:    1,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"code_repos": {
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Description: "List of pattern strings to define code_repos in resource list",
										MinItems:    1,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"functions": {
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Description: "List of pattern strings to define functions in resource list",
										MinItems:    1,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"containers": {
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Description: "List of pattern strings to define containers in resource list",
										MinItems:    1,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"namespaces": {
										Type:        schema.TypeList,
										Optional:    true,
										Computed:    true,
										Description: "List of pattern strings to define namespaces in resource list",
										MinItems:    1,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
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

func deleteResourceList(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	id := d.Id()
	log.Printf("[INFO]: Deleting Resource List, Id:%+v\n", id)
	if err := resource_list.Delete(client, id); err != nil {
		if !errors.Is(err, pc.ObjectNotFoundError) {
			return diag.FromErr(err)
		}
	}
	d.SetId("")
	return nil
}

func updateResourceList(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var err error
	client := meta.(*pc.Client)
	o, err := parseResourceList(d)
	if err != nil {
		return diag.FromErr(err)
	}
	log.Printf("[INFO]: Updating Resource List, Id:%+v\n", d.Get("id"))
	if _, err = resource_list.Update(client, o, d.Get("id").(string)); err != nil {
		if errors.Is(err, pc.ObjectNotFoundError) {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}
	return readResourceList(ctx, d, meta)
}

func createResourceList(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var err error
	client := meta.(*pc.Client)
	o, err := parseResourceList(d)
	if err != nil {
		return diag.FromErr(err)
	}
	var listRes resource_list.ResourceList
	if listRes, err = resource_list.Create(client, o); err != nil {
		return diag.FromErr(err)
	}
	log.Printf("[INFO]: Resource List Created Successfully, Id:%+v\n", listRes.Id)
	d.SetId(listRes.Id)
	return readResourceList(ctx, d, meta)
}

func parseResourceList(d *schema.ResourceData) (resource_list.ResourceListRequest, error) {

	resourceListType := d.Get("resource_list_type").(string)
	rlReq := resource_list.ResourceListRequest{
		Name:             d.Get("name").(string),
		Description:      d.Get("description").(string),
		ResourceListType: resourceListType,
		Members:          nil,
	}

	if resourceListType == resource_list.TypeTags {
		members := ResourceDataInterfaceMap(d, "members")
		//condition to check if "tags" is defined
		if _, ok := members["tags"]; !ok || len(members["tags"].([]interface{})) == 0 {
			return resource_list.ResourceListRequest{}, errors.New("tags need to be defined in members")
		}
		tags := members["tags"].([]interface{})
		formattedMembers := make([]interface{}, 0, len(tags))
		for _, obj := range tags {
			//condition to check if key is defined within each key object
			if obj == nil || obj.(map[string]interface{})["key"] == nil {
				return resource_list.ResourceListRequest{}, errors.New("key needs to be defined for each tag")
			}
			tag := map[string]string{
				obj.(map[string]interface{})["key"].(string): obj.(map[string]interface{})["value"].(string),
			}
			formattedMembers = append(formattedMembers, tag)
		}
		rlReq.Members = formattedMembers
	} else if resourceListType == resource_list.TypeComputeAccessGroups {
		members := ResourceDataInterfaceMap(d, "members")
		//condition to check if "compute_access_groups" is defined
		if _, ok := members["compute_access_groups"]; !ok || len(members["compute_access_groups"].([]interface{})) == 0 {
			return resource_list.ResourceListRequest{}, errors.New("compute_access_groups need to be defined in members")
		}
		//condition to check if atleast one group within compute_access_groups is defined
		if members["compute_access_groups"].([]interface{})[0] == nil {
			return resource_list.ResourceListRequest{}, errors.New("compute_access_groups fields need to be defined")
		}
		memberComputeAccessGroups := members["compute_access_groups"].([]interface{})
		formattedMembers := make([]interface{}, 0)
		for _, obj := range memberComputeAccessGroups {
			formattedMember := map[string]interface{}{
				"hosts":      obj.(map[string]interface{})["hosts"],
				"appIDs":     obj.(map[string]interface{})["app_id"],
				"images":     obj.(map[string]interface{})["images"],
				"labels":     obj.(map[string]interface{})["labels"],
				"clusters":   obj.(map[string]interface{})["clusters"],
				"codeRepos":  obj.(map[string]interface{})["code_repos"],
				"functions":  obj.(map[string]interface{})["functions"],
				"containers": obj.(map[string]interface{})["containers"],
				"namespaces": obj.(map[string]interface{})["namespaces"],
			}
			formattedMembers = append(formattedMembers, formattedMember)
		}
		rlReq.Members = formattedMembers
	} else if resourceListType == resource_list.TypeAzureResourceGroups {
		members := ResourceDataInterfaceMap(d, "members")
		if _, ok := members["azure_resource_groups"]; !ok || len(members["azure_resource_groups"].([]interface{})) == 0 {
			return resource_list.ResourceListRequest{}, errors.New("azure_resource_groups need to be defined in members")
		}
		rlReq.Members = ResourceDataInterfaceMap(d, "members")["azure_resource_groups"].([]interface{})
	}

	return rlReq, nil
}

func readResourceList(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*pc.Client)
	id := d.Id()
	obj, err := resource_list.Get(client, id)
	if err != nil {
		if errors.Is(err, pc.ResourceListNotFoundError) {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}
	saveResourceList(d, obj)
	return nil
}
func saveResourceList(d *schema.ResourceData, o resource_list.ResourceList) {
	d.Set("name", o.Name)
	d.Set("description", o.Description)
	d.Set("resource_list_type", o.ResourceListType)
	d.Set("last_modified_ts", o.LastModifiedTs)
	d.Set("last_modified_by", o.LastModifiedBy)

	membersForAll := map[string]interface{}{
		"tags":                  []interface{}{},
		"azure_resource_groups": []interface{}{},
		"compute_access_groups": []interface{}{},
	}

	if o.ResourceListType == resource_list.TypeTags {
		if len(o.Members) != 0 {
			tagMembers := make([]map[string]string, 0, len(o.Members))
			for _, m := range o.Members {
				tagMember := make(map[string]string)
				for key, value := range m.(map[string]interface{}) {
					tagMember["key"] = key
					tagMember["value"] = value.(string)
				}
				tagMembers = append(tagMembers, tagMember)
			}
			membersForAll["tags"] = tagMembers
		}
	} else if o.ResourceListType == resource_list.TypeAzureResourceGroups {
		if len(o.Members) != 0 {
			membersForAll["azure_resource_groups"] = o.Members
		}
	} else if o.ResourceListType == resource_list.TypeComputeAccessGroups {
		memberComputeAccessGroups := make([]interface{}, 0, len(o.Members))
		if len(o.Members) != 0 {
			for _, member := range o.Members {
				memberMap := member.(map[string]interface{})
				memberComputeAccessGroups = append(memberComputeAccessGroups, map[string]interface{}{
					"hosts":      memberMap["hosts"],
					"app_id":     memberMap["appIDs"],
					"images":     memberMap["images"],
					"labels":     memberMap["labels"],
					"clusters":   memberMap["clusters"],
					"code_repos": memberMap["codeRepos"],
					"functions":  memberMap["functions"],
					"containers": memberMap["containers"],
					"namespaces": memberMap["namespaces"],
				})
			}
		}
		membersForAll["compute_access_groups"] = memberComputeAccessGroups
	}

	if err := d.Set("members", []interface{}{membersForAll}); err != nil {
		log.Printf("[WARN] Error setting 'members' for %s: %s", d.Id(), err)
	}
}
