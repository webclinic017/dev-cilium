module "chamber-f" {
    source = "../../terraform-aws-defn-account"
    context = module.this.context

    providers = {
        aws = aws.chamber-f
    }
}