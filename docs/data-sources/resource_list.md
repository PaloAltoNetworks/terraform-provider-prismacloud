---
page_title: "Prisma Cloud: prismacloud_resource_list"
---

# prismacloud_resource_list

Retrieves resource list information by id.

## Example Usage

```hcl
data "prismacloud_resource_list" "example" {
  id = "resource list id"
}
```

## Argument Reference

You must specify:

* `id` - ID of the resource list.

## Attribute Reference

* `name` - Name of the resource list.
* `description` - Description of the resource list.
* `resource_list_type` - Type of resource list.
* `last_modified_by` - Last modified by.
* `last_modified_ts` - (int) Last modified timestamp.
* `members` - Associated resource list members as defined [below](#members).

### Members

Each member has the following attributes:

* `tags` - Associated resource list tags as defined [below](#tags)
* `azure_resource_groups` - Consists of a list of Azure Resource Groups IDs associated with the resource list.
* `compute_access_groups` - Associated resource list Compute Access Groups as defined [below](#compute-access-groups)

#### Tags

Each Tag has the following attributes:

* `key` - Key of the tag.
* `value` - Value of the tag.

#### Compute Access Groups

Specifies the filters to define the scope of what is accessible within each type of resource. By default, each field is populated with a wildcard to match all objects of a specific type.

Each Compute Access Groups object has the following attributes:

* `app_id` - (Optional) App id
* `clusters` - (Optional) Clusters
* `code_repos` - (Optional) Code repos
* `containers` - (Optional) Containers
* `functions` - (Optional) Functions
* `hosts` - (Optional) Hosts
* `images` - (Optional) Images
* `labels` - (Optional) Labels
* `namespaces` - (Optional) Namespaces