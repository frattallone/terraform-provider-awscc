---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_servicecatalog_service_action Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::ServiceCatalog::ServiceAction
---

# awscc_servicecatalog_service_action (Data Source)

Data Source schema for AWS::ServiceCatalog::ServiceAction



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **accept_language** (String)
- **definition** (Attributes List) (see [below for nested schema](#nestedatt--definition))
- **definition_type** (String)
- **description** (String)
- **name** (String)

<a id="nestedatt--definition"></a>
### Nested Schema for `definition`

Read-Only:

- **key** (String)
- **value** (String)

