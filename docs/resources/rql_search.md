---
page_title: "Prisma Cloud: prismacloud_rql_search"
---

# prismacloud_rql_search

Manage a RQL search on the Prisma Cloud Platform.

NOTE:  Prisma Cloud does not currently support deleting RQL searches, so
`terraform destroy` is a noop.

## Example Usage

```hcl
resource "prismacloud_rql_search" "example" {
    search_type = "config"
    query = "config from cloud.resource where api.name = 'aws-ec2-describe-instances'"
    time_range {
        relative {
            unit = "hour"
            amount = 24
        }
    }
}
```

## Argument Reference

The following arguments are supported:

* `search_type` - (Required) The search type. Valid values are `config`
  (default) `event`, `network` and `iam`.
* `query` - (Required) The RQL query.
* `limit` - (int) Limit rules (default: `10`).
* `time_range` - (Required for config, event and network RQL search) The RQL time range spec, as defined [below](#time-range).

### Time Range

Only one of these can be defined:

* `absolute` - An absolute time range spec, as defined [below](#absolute-time-range).
* `relative` - A relative time range spec, as defined [below](#relative-time-range).
* `to_now` - A "To Now" time range spec, as defined [below](#to-now-time-range).

### Absolute Time Range

* `start` - (Required, int) Start time.
* `end` - (Required, int) End time.

### Relative Time Range

* `amount` - (Required, int) The time number.
* `unit` - (Required) The time unit.

### To Now Time Range

* `unit` - (Required) The time unit.

## Attribute Reference

The following attributes are supported:

* `search_id` - The search ID returned from a successful RQL query.
* `group_by` - (list) Group by.
* `cloud_type` - The cloud type.
* `name` - Name.
* `description` - Description.
* `config_data` - (for `search_type="config"`, list) List of config_data specs, as defined below.
* `event_data` - (For `search_type="event"`, list) List of event_data specs, as defined below.
* `network_data` - (For `search_type="network"`, list) List of network_data specs, as defined below.
* `iam_data` - (For `search_type="iam"`, list) List of iam_data specs, as defined below.

`config_data` supports the following attributes:

* `state_id` - The state ID.
* `name` - Name.
* `url` - The URL.

`event_data` supports the following attributes:

* `account` - Account.
* `region_id` - (int) Region ID.
* `region_api_identifier` - Region API identifier.

`network_data` supports the following attributes:

* `account` - Account.
* `region_id` - (int) Region ID.
* `account_name` - Account Name.

`iam_data` supports the following attributes:

* `accessed_resources_count` - (int) Accessed resource count.
* `dest_cloud_account` - Destination cloud account.
* `dest_cloud_region` - Destination cloud region.
* `dest_cloud_resource_rrn` - Destination cloud resource RRN.
* `dest_cloud_service_name` - Destination cloud service name.
* `dest_cloud_type` - Destination cloud type.
* `dest_resource_id` - Destination cloud resource id.
* `dest_resource_name` - Destination cloud resource name.
* `dest_resource_type` - Destination cloud resource type.
* `effective_action_name` - Effective action name.
* `granted_by_cloud_entity_id` - Granted by cloud entity id.
* `granted_by_cloud_entity_name` - Granted by cloud entity name.
* `granted_by_cloud_entity_rrn` - Granted by cloud entity rrn.
* `granted_by_cloud_entity_type` - Granted by cloud entity type.
* `granted_by_cloud_policy_id` - Granted by cloud policy id.
* `granted_by_cloud_policy_name` - Granted by cloud policy name.
* `granted_by_cloud_policy_rrn` - Granted by cloud policy rrn.
* `granted_by_cloud_policy_type` - Granted by cloud policy type.
* `granted_by_cloud_type` - Granted by cloud type.
* `message_id` - Message id.
* `is_wild_card_dest_cloud_resource_name` - (bool) Is destination cloud resource name a wildcard.
* `last_access_date` - Last access date.
* `source_cloud_account` - Source cloud account.
* `source_cloud_region` - Source cloud region.
* `source_cloud_resource_rrn` - Source cloud resource rrn.
* `source_cloud_service_name` - Source cloud service name.
* `source_cloud_type` - Source cloud type.
* `source_idp_domain` - Source IDP domain.
* `source_idp_email` - Source IDP email.
* `source_idp_group` - Source IDP group.
* `source_idp_rrn` - Source IDP rrn.
* `source_idp_service` - Source IDP service.
* `source_idp_user_name` - Source IDP user name.
* `source_public` - (bool) Is source public.
* `source_resource_id` - Source cloud resource id.
* `source_resource_type` - Source cloud resource type.
* `exceptions` - (list) Permission exception list, as defined below.

`exceptions` supports the following attributes:

* `message_code` - Message code.