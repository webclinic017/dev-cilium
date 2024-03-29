{
  inputs = {
    kubectl.url = github:defn/dev/pkg-kubectl-1.26.7-4?dir=m/pkg/kubectl;
    k3d.url = github:defn/dev/pkg-k3d-5.5.2-1?dir=m/pkg/k3d;
    k9s.url = github:defn/dev/pkg-k9s-0.27.4-4?dir=m/pkg/k9s;
    helm.url = github:defn/dev/pkg-helm-3.12.3-1?dir=m/pkg/helm;
    kustomize.url = github:defn/dev/pkg-kustomize-5.0.3-4?dir=m/pkg/kustomize;
    stern.url = github:defn/dev/pkg-stern-1.25.0-9?dir=m/pkg/stern;
    argoworkflows.url = github:defn/dev/pkg-argoworkflows-3.4.9-1?dir=m/pkg/argoworkflows;
    argocd.url = github:defn/dev/pkg-argocd-2.8.0-1?dir=m/pkg/argocd;
    kn.url = github:defn/dev/pkg-kn-1.11.0-1?dir=m/pkg/kn;
    vcluster.url = github:defn/dev/pkg-vcluster-0.15.5-1?dir=m/pkg/vcluster;
    kubevirt.url = github:defn/dev/pkg-kubevirt-1.0.0-2?dir=m/pkg/kubevirt;
    cilium.url = github:defn/dev/pkg-cilium-0.15.5-1?dir=m/pkg/cilium;
    hubble.url = github:defn/dev/pkg-hubble-0.12.0-2?dir=m/pkg/hubble;
    tfo.url = github:defn/dev/pkg-tfo-1.3.0-2?dir=m/pkg/tfo;
  };

  outputs = inputs: inputs.kubectl.inputs.pkg.main rec {
    src = ./.;

    defaultPackage = ctx: ctx.wrap.nullBuilder {
      propagatedBuildInputs =
        let
          flakeInputs = [
            inputs.kubectl.defaultPackage.${ctx.system}
            inputs.k3d.defaultPackage.${ctx.system}
            inputs.k9s.defaultPackage.${ctx.system}
            inputs.helm.defaultPackage.${ctx.system}
            inputs.kustomize.defaultPackage.${ctx.system}
            inputs.stern.defaultPackage.${ctx.system}
            inputs.argoworkflows.defaultPackage.${ctx.system}
            inputs.argocd.defaultPackage.${ctx.system}
            inputs.kn.defaultPackage.${ctx.system}
            inputs.vcluster.defaultPackage.${ctx.system}
            inputs.kubevirt.defaultPackage.${ctx.system}
            inputs.cilium.defaultPackage.${ctx.system}
            inputs.hubble.defaultPackage.${ctx.system}
            inputs.tfo.defaultPackage.${ctx.system}
          ];
        in
        flakeInputs;
    };
  };
}
