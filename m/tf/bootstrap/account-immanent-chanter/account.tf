module "immanent-chanter" {
    source = "../../mod/terraform-aws-defn-account"
    context = module.this.context

    providers = {
        aws = aws.immanent-chanter
    }
}