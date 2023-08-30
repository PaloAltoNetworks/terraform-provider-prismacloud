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

Each compute access group has the following attributes:

* `app_id` - List of filters to define the scope of what app_ids are accessible. By default, it is populated with a wildcard to match all objects of the specific type.
* `clusters` - List of filters to define the scope of what clusters are accessible. By default, it is populated with a wildcard to match all objects of the specific type.
* `code_repos` - List of filters to define the scope of what code_repos are accessible. By default, it is populated with a wildcard to match all objects of the specific type.
* `containers` - List of filters to define the scope of what containers are accessible. By default, it is populated with a wildcard to match all objects of the specific type.
* `functions` - List of filters to define the scope of what functions are accessible. By default, it is populated with a wildcard to match all objects of the specific type.
* `hosts` - List of filters to define the scope of what hosts are accessible. By default, it is populated with a wildcard to match all objects of the specific type.
* `images` - List of filters to define the scope of what images are accessible. By default, it is populated with a wildcard to match all objects of the specific type.
* `labels` - List of filters to define the scope of what labels are accessible. By default, it is populated with a wildcard to match all objects of the specific type.
* `namespaces` - List of filters to define the scope of what namespaces are accessible. By default, it is populated with a wildcard to match all objects of the specific type.

