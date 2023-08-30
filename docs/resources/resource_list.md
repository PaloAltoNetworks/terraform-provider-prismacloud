---
page_title: "Prisma Cloud: prismacloud_resource_list"
---

# prismacloud_resource_list

Manage a resource list.

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

* `name` - (Required) Name of the resource list.
* `resource_list_type` - (Required) Name of the resource list.
* `description` - (Optional) Description of the resource list.
* `members` - (Required) Associated resource list members as defined [below](#members).

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

* `key` - (Required) Key of the tag.
* `value` - (Optional) Value of the tag.

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