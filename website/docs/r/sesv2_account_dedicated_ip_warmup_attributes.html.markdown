---
subcategory: "SESv2 (Simple Email V2)"
layout: "aws"
page_title: "AWS: aws_sesv2_account_dedicated_ip_warmup_attributes"
description: |-
  Manages AWS SESv2 (Simple Email V2) account-level dedicated IP warmup attributes.
---

# Resource: aws_sesv2_account_dedicated_ip_warmup_attributes

Manages AWS SESv2 (Simple Email V2) account-level dedicated IP warmup attributes. These control whether Amazon SES automatically warms up dedicated IP addresses that you lease for your account.

~> Destroying this resource only removes it from Terraform state. The account's automatic warmup setting is left unchanged.

## Example Usage

```terraform
resource "aws_sesv2_account_dedicated_ip_warmup_attributes" "example" {
  auto_warmup_enabled = false
}
```

## Argument Reference

This resource supports the following arguments:

* `auto_warmup_enabled` - (Required) Whether automatic warmup is enabled for dedicated IP addresses associated with the account.
* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).

## Attribute Reference

This resource exports no additional attributes.

## Import

In Terraform v1.12.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `identity` attribute. For example:

```terraform
import {
  to = aws_sesv2_account_dedicated_ip_warmup_attributes.example
  identity = {
    region = "us-west-2"
  }
}

resource "aws_sesv2_account_dedicated_ip_warmup_attributes" "example" {
  ### Configuration omitted for brevity ###
}
```

### Identity Schema

#### Required

* `region` (String) Region where this resource is managed.

#### Optional

* `account_id` (String) AWS Account where this resource is managed.

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import account-level dedicated IP warmup attributes using `region`. For example:

```terraform
import {
  to = aws_sesv2_account_dedicated_ip_warmup_attributes.example
  id = "us-west-2"
}
```

Using `terraform import`, import account-level dedicated IP warmup attributes using `region`. For example:

```console
% terraform import aws_sesv2_account_dedicated_ip_warmup_attributes.example us-west-2
```
