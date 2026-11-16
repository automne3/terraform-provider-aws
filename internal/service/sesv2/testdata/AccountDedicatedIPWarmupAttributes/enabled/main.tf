# Copyright IBM Corp. 2014, 2026
# SPDX-License-Identifier: MPL-2.0

resource "aws_sesv2_account_dedicated_ip_warmup_attributes" "test" {
  auto_warmup_enabled = true
}
