module "fogg-hub" {
    source = "../../terraform-aws-defn-account"
    context = module.this.context

    providers = {
        aws = aws.fogg-hub
    }
}