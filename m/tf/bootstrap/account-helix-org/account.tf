module "helix-org" {
    source = "../../mod/terraform-aws-defn-account"
    context = module.this.context

    providers = {
        aws = aws.helix-org
    }
}