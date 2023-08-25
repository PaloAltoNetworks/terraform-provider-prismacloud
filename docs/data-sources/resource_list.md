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

* `id` - (string) ID of the resource list. (required)

## Attribute Reference

* `name` - (string) List of cloud account IDs.
* `description` - (string) Description of the resource list.
* `resource_list_type` - (string) Type of resource list.
* `last_modified_by` - (string) Last modified by.
* `last_modified_ts` - (int) Last modified timestamp.
* `members` - Associated resource list members as defined [below](#members).

### Members

Each member has the following attributes:

* `tags` - Associated resource list tags as defined [below](#tags)
* `azure_resource_groups` - Associated resource list Azure Resource Groups as defined [below](#azure-resource-groups)
* `compute_access_groups` - Associated resource list Compute Access Groups as defined [below](#compute-access-groups)

#### Tags

Each Tag has the following attributes:

* `key` - Key of the tag.
* `value` - Value of the tag.

#### Azure Resource Groups

Consists of a list of Azure Resource Groups IDs (string) associated with the resource list.

#### Compute Access Groups

Each Tag has the following attributes:

* `app_id` - List of patterns (string) to be matched to the app id.
* `clusters` - List of patterns (string) to be matched to the clusters.
* `code_repos` - List of patterns (string) to be matched to the code_repos.
* `containers` - List of patterns (string) to be matched to the containers.
* `functions` - List of patterns (string) to be matched to the functions.
* `hosts` - List of patterns (string) to be matched to the hosts.
* `images` - List of patterns (string) to be matched to the images.
* `labels` - List of patterns (string) to be matched to the labels.
* `namespaces` - List of patterns (string) to be matched to the namespaces.

