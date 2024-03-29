{
  inputs = {
    terraform.url = github:defn/dev/pkg-terraform-1.5.5-1?dir=m/pkg/terraform;
    terraformdocs.url = github:defn/dev/pkg-terraformdocs-0.16.0-1?dir=m/pkg/terraformdocs;
    packer.url = github:defn/dev/pkg-packer-1.9.2-1?dir=m/pkg/packer;
    step.url = github:defn/dev/pkg-step-0.24.4-4?dir=m/pkg/step;
    awscli.url = github:defn/dev/pkg-awscli-2.13.9-1?dir=m/pkg/awscli;
    flyctl.url = github:defn/dev/pkg-flyctl-0.1.77-1?dir=m/pkg/flyctl;
    chamber.url = github:defn/dev/pkg-chamber-2.13.3-1?dir=m/pkg/chamber;
  };

  outputs = inputs: inputs.terraform.inputs.pkg.main rec {
    src = ./.;

    defaultPackage = ctx: ctx.wrap.bashBuilder {
      inherit src;

      propagatedBuildInputs =
        let
          flakeInputs = [
            inputs.terraform.defaultPackage.${ctx.system}
            inputs.terraformdocs.defaultPackage.${ctx.system}
            inputs.packer.defaultPackage.${ctx.system}
            inputs.step.defaultPackage.${ctx.system}
            inputs.awscli.defaultPackage.${ctx.system}
            inputs.flyctl.defaultPackage.${ctx.system}
            inputs.chamber.defaultPackage.${ctx.system}
          ];
        in
        flakeInputs;

      installPhase = ''
        mkdir -p $out/bin
        cp -a $src/bin/* $out/bin/
        chmod 755 $out/bin/tf $out/bin/tf-*
      '';
    };
  };
}