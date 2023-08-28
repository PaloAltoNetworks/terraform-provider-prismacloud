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

* `last_modified_by` - (string) Last modified by.
* `last_modified_ts` - (int) Last modified timestamp.

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

* `app_id` - List of filters to define the scope of what app_ids are accessible. By default, it is populated with a wildcard to match all objects of the specific type.
* `clusters` - List of filters to define the scope of what clusters are accessible. By default, it is populated with a wildcard to match all objects of the specific type.
* `code_repos` - List of filters to define the scope of what code_repos are accessible. By default, it is populated with a wildcard to match all objects of the specific type.
* `containers` - List of filters to define the scope of what containers are accessible. By default, it is populated with a wildcard to match all objects of the specific type.
* `functions` - List of filters to define the scope of what functions are accessible. By default, it is populated with a wildcard to match all objects of the specific type.
* `hosts` - List of filters to define the scope of what hosts are accessible. By default, it is populated with a wildcard to match all objects of the specific type.
* `images` - List of filters to define the scope of what images are accessible. By default, it is populated with a wildcard to match all objects of the specific type.
* `labels` - List of filters to define the scope of what labels are accessible. By default, it is populated with a wildcard to match all objects of the specific type.
* `namespaces` - List of filters to define the scope of what namespaces are accessible. By default, it is populated with a wildcard to match all objects of the specific type.