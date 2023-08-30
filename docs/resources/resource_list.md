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

* `name` - (required) Name of the resource list.
* `resource_list_type` - (required) Name of the resource list.
* `description` - (optional) Description of the resource list.
* `members` - (required) Associated resource list members as defined [below](#members).

## Attribute Reference

* `last_modified_by` - Last modified by.
* `last_modified_ts` - (int) Last modified timestamp.

### Members

Each member has the following attributes:

* `tags` - Associated resource list tags as defined [below](#tags)
* `azure_resource_groups` - Consists of a list of Azure Resource Groups IDs associated with the resource list.
* `compute_access_groups` - Associated resource list Compute Access Groups as defined [below](#compute-access-groups)

#### Tags

Each Tag has the following attributes:

* `key` - (required) Key of the tag.
* `value` - (optional) Value of the tag.

#### Compute Access Groups

Specifies the filters to define the scope of what is accessible within each type of resource. By default, each field is populated with a wildcard to match all objects of a specific type.

Each Compute Access Groups object has the following attributes:

* `app_id` - (optional) app id
* `clusters` - (optional) clusters
* `code_repos` - (optional) code repos
* `containers` - (optional) containers
* `functions` - (optional) functions
* `hosts` - (optional) hosts
* `images` - (optional) images
* `labels` - (optional) labels
* `namespaces` - (optional) namespaces