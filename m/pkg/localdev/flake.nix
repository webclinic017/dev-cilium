{
  inputs = {
    tilt.url = github:defn/dev/pkg-tilt-0.33.4-1?dir=m/pkg/tilt;
    gh.url = github:defn/dev/pkg-gh-2.32.1-1?dir=m/pkg/gh;
    earthly.url = github:defn/dev/pkg-earthly-0.7.15-1?dir=m/pkg/earthly;
    oras.url = github:defn/dev/pkg-oras-1.0.1-1?dir=m/pkg/oras;
    buildkite.url = github:defn/dev/pkg-buildkite-3.50.4-1?dir=m/pkg/buildkite;
    bk.url = github:defn/dev/pkg-bk-2.0.0-11?dir=m/pkg/bk;
    buildevents.url = github:defn/dev/pkg-buildevents-0.15.0-1?dir=m/pkg/buildevents;
    honeyvent.url = github:defn/dev/pkg-honeyvent-1.1.3-10?dir=m/pkg/honeyvent;
    honeymarker.url = github:defn/dev/pkg-honeymarker-0.2.10-10?dir=m/pkg/honeymarker;
    honeytail.url = github:defn/dev/pkg-honeytail-1.8.3-8?dir=m/pkg/honeytail;
    hugo.url = github:defn/dev/pkg-hugo-0.0.5?dir=m/pkg/hugo;
    temporal.url = github:defn/dev/pkg-temporal-0.10.5-1?dir=m/pkg/temporal;
  };

  outputs = inputs: inputs.tilt.inputs.pkg.main rec {
    src = ./.;

    defaultPackage = ctx: ctx.wrap.nullBuilder {
      propagatedBuildInputs =
        let
          flakeInputs = with ctx.pkgs; [
            inputs.tilt.defaultPackage.${ctx.system}
            inputs.gh.defaultPackage.${ctx.system}
            inputs.earthly.defaultPackage.${ctx.system}
            inputs.oras.defaultPackage.${ctx.system}
            inputs.buildkite.defaultPackage.${ctx.system}
            inputs.bk.defaultPackage.${ctx.system}
            inputs.buildevents.defaultPackage.${ctx.system}
            inputs.honeyvent.defaultPackage.${ctx.system}
            inputs.honeymarker.defaultPackage.${ctx.system}
            inputs.honeytail.defaultPackage.${ctx.system}
            inputs.hugo.defaultPackage.${ctx.system}
            inputs.temporal.defaultPackage.${ctx.system}
            skopeo
          ];
        in
        flakeInputs;
    };
  };
}
