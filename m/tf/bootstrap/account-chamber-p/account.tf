module "chamber-p" {
    source = "../../mod/terraform-aws-defn-account"
    context = module.this.context

    providers = {
        aws = aws.chamber-p
    }
}