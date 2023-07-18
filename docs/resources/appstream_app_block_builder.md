---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_appstream_app_block_builder Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::AppStream::AppBlockBuilder.
---

# awscc_appstream_app_block_builder (Resource)

Resource Type definition for AWS::AppStream::AppBlockBuilder.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `instance_type` (String)
- `name` (String)
- `platform` (String)
- `vpc_config` (Attributes) (see [below for nested schema](#nestedatt--vpc_config))

### Optional

- `access_endpoints` (Attributes Set) (see [below for nested schema](#nestedatt--access_endpoints))
- `app_block_arns` (Set of String)
- `description` (String)
- `display_name` (String)
- `enable_default_internet_access` (Boolean)
- `iam_role_arn` (String)
- `tags` (Attributes Set) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String)
- `created_time` (String)
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--vpc_config"></a>
### Nested Schema for `vpc_config`

Optional:

- `security_group_ids` (List of String)
- `subnet_ids` (List of String)


<a id="nestedatt--access_endpoints"></a>
### Nested Schema for `access_endpoints`

Required:

- `endpoint_type` (String)
- `vpce_id` (String)


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_appstream_app_block_builder.example <resource ID>
```