# Copyright IBM Corp. 2014, 2026
# SPDX-License-Identifier: MPL-2.0

resource "aws_sesv2_account_dedicated_ip_warmup_attributes" "test" {
  region = var.region

  auto_warmup_enabled = false
}


variable "region" {
  description = "Region to deploy resource in"
  type        = string
  nullable    = false
}
