---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_b2bi_transformer Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::B2BI::Transformer Resource Type
---

# awscc_b2bi_transformer (Resource)

Definition of AWS::B2BI::Transformer Resource Type



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `edi_type` (Attributes) (see [below for nested schema](#nestedatt--edi_type))
- `file_format` (String)
- `mapping_template` (String)
- `name` (String)
- `status` (String)

### Optional

- `modified_at` (String)
- `sample_document` (String)
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `created_at` (String)
- `id` (String) Uniquely identifies the resource.
- `transformer_arn` (String)
- `transformer_id` (String)

<a id="nestedatt--edi_type"></a>
### Nested Schema for `edi_type`

Optional:

- `x12_details` (Attributes) (see [below for nested schema](#nestedatt--edi_type--x12_details))

<a id="nestedatt--edi_type--x12_details"></a>
### Nested Schema for `edi_type.x12_details`

Optional:

- `transaction_set` (String)
- `version` (String)



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_b2bi_transformer.example <resource ID>
```
