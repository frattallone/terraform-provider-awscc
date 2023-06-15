---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ec2_subnet_cidr_block Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::EC2::SubnetCidrBlock resource creates association between subnet and IPv6 CIDR
---

# awscc_ec2_subnet_cidr_block (Resource)

The AWS::EC2::SubnetCidrBlock resource creates association between subnet and IPv6 CIDR



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `ipv_6_cidr_block` (String) The IPv6 network range for the subnet, in CIDR notation. The subnet size must use a /64 prefix length
- `subnet_id` (String) The ID of the subnet

### Read-Only

- `id` (String) Information about the IPv6 association.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_ec2_subnet_cidr_block.example <resource ID>
```