module "jianghu-tahoe" {
    source = "../../terraform-aws-defn-account"
    context = module.this.context

    providers = {
        aws = aws.jianghu-tahoe
    }
}