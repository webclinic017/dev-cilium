module "jianghu-klamath" {
    source = "../../mod/terraform-aws-defn-account"
    context = module.this.context

    providers = {
        aws = aws.jianghu-klamath
    }
}