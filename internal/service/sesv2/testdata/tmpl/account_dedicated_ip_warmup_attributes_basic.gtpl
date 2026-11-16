resource "aws_sesv2_account_dedicated_ip_warmup_attributes" "test" {
{{- template "region" }}
  auto_warmup_enabled = false
}
