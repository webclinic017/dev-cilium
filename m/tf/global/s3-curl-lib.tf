module "s3-curl-lib" {
  source     = "../terraform-aws-s3-bucket"
  context    = module.this.context
  attributes = [ "curl-lib" ]

  providers = {
    aws = aws.curl-lib
  }

  acl                = "private"
  user_enabled       = false
  versioning_enabled = false

  lifecycle_configuration_rules = local.lifecycle_configuration_rules

  privileged_principal_arns    = local.privileged_principal_arns
  privileged_principal_actions = local.privileged_principal_actions
}