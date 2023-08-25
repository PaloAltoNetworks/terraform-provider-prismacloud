---
page_title: "Prisma Cloud: prismacloud_resource_list"
---

# prismacloud_resource_list

Create resource list information by id.

## Example Usage (with tags)

```hcl
resource "prismacloud_resource_list" "example" {
  name = "name"
  resource_list_type = "TAG"
  members {
    tags {
      key = "key1"
      value = "value1"
    }
    tags {
      key = "key2"
      value = "value2"
    }
  }
}
```

## Example Usage (with azure resource groups)

```hcl
resource "prismacloud_resource_list" "example" {
  name = "name"
  resource_list_type = "RESOURCE_GROUP"
  members {
    azure_resource_groups = ["resource-groups-1", "resource-group-2"]
  }
}
```

## Example Usage (with compute access groups)

```hcl
resource "prismacloud_resource_list" "example" {
  name = "name"
  resource_list_type = "COMPUTE_ACCESS_GROUP"
  members {
    compute_access_groups {
      hosts = ["*"]
      app_id = ["*"]
      images = ["*"]
      labels = ["*"]
      clusters = ["*"]
      code_repos = ["*"]
      functions = ["*"]
      containers = ["*"]
      namespaces = ["*"]
    }
  }
}
```

## Argument Reference

* `name` - (string) Name of the resource list. (required)
* `resource_list_type` - (string) Name of the resource list. (required)
* `description` - (string) Description of the resource list.
* `members` - Associated resource list members as defined [below](#members). (required)

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

