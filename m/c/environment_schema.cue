package c

import (
	"encoding/yaml"
	"strings"
)

#TransformEnvToAnyResource: {
	from: {
		#Input

		type: string
	}

	to: #KustomizeHelm & {
		_in: from

		helm: {
			release: "bootstrap"
			name:    "any-resource"
			version: "0.1.0"
			repo:    "https://kiwigrid.github.io"
			values: {
				anyResources: {
					for _app_name, _app in bootstrap[_in.name].apps {
						"\(_app_name)": yaml.Marshal(_app.application)
					}
				}
			}
		}
	}
}

#TransformEnvToBootstrapMachine: {
	from: {
		#Input

		type: string
		bootstrap: [string]: [int, ...string]
	}

	to: #BootstrapMachine
}

#BootstrapMachine: ctx={
	_in: #TransformEnvToBootstrapMachine.from

	machine_name: string | *_in.name
	machine_type: string | *_in.type

	apps: [string]: #BootstrapApp
	apps: {
		for _app_name, _app in _in.bootstrap {
			"\(_app_name)": #BootstrapApp & {
				machine_type:     ctx.machine_type
				machine_name:     ctx.machine_name
				app_name:         _app_name
				app_wave:         _app[0]
				app_namespace:    _app[1]
				app_sync_options: _app[2:]
			}
		}
	}
}

#BootstrapApp: {
	machine_type:     string
	machine_name:     string
	app_name:         string
	app_namespace:    string
	app_wave:         int
	app_sync_options: [...string] | *[]

	application: {
		apiVersion: "argoproj.io/v1alpha1"
		kind:       "Application"

		metadata: {
			namespace: "argocd"
			if app_name =~ "^\(machine_type)-\(machine_name)-" {
				name: "\(app_name)"
			}
			if app_name !~ "^\(machine_type)-\(machine_name)-" {
				name: "\(machine_type)-\(machine_name)-\(app_name)"
			}
			annotations: "argocd.argoproj.io/sync-wave": "\(app_wave)"
		}

		spec: {
			project: "default"

			destination: {
				name: "\(machine_type)-\(machine_name)"
				if app_namespace != "" {
					namespace: app_namespace
				}
			}

			source: {
				repoURL:        "https://github.com/defn/dev"
				targetRevision: "main"
				if strings.HasPrefix(app_name, "\(machine_type)-\(machine_name)") {
					path: "m/k/r/\(app_name)"
				}
				if !strings.HasPrefix(app_name, "\(machine_type)-\(machine_name)") {
					path: "m/k/r/\(machine_type)-\(machine_name)-\(app_name)"
				}
			}

			syncPolicy: {
				automated: {
					prune:    true
					selfHeal: true
				}
				syncOptions: app_sync_options
			}
		}
	}
}

// Machine
#Machine: {
	type: string
	name: string

	destination: string | *"\(type)-\(name)"

	bootstrap: [string]: [int, ...string]
	env: #EnvApp
	env: {
		// ex: k/k3d-dgd
		// ex: k/vcluster-dfd-vc0
		spec: source: path: "m/k/r/\(type)-\(name)"

		spec: "destination": "name": destination
	}
}

#EnvApp: {
	apiVersion: "argoproj.io/v1alpha1"
	kind:       "Application"

	metadata: {
		namespace: "argocd"
		name:      string
	}

	spec: {
		project: "default"

		destination: name:       string
		destination: namespace?: string
		source: {
			repoURL:        "https://github.com/defn/dev"
			targetRevision: "main"
			path:           string
		}

		syncPolicy: automated: {
			prune:    bool | *true
			selfHeal: bool | *true
		}
	}
}

// K3D Machine
#TransformK3D: {
	from: {
		#Input
		bootstrap: [string]: [int, ...string]
	}

	to: #K3D
}

#K3D: ctx={
	_in: #TransformK3D.from

	#Machine

	type:      "k3d"
	name:      _in.name
	bootstrap: _in.bootstrap

	// ex: k3d-dfd
	env: metadata: name: "\(type)-\(ctx.name)"
}

// VCluster Machine
#TransformVCluster: {
	from: {
		#Input
		bootstrap: [string]: [int, ...string]
		instance_types: [...string]
		parent: #K3D
	}

	to: #VCluster
}

#VCluster: ctx={
	_in: #TransformVCluster.from

	#Machine

	type:           "vcluster"
	name:           _in.name
	destination:    "\(parent.type)-\(parent.name)"
	bootstrap:      _in.bootstrap
	instance_types: _in.instance_types
	parent:         #K3D & _in.parent

	instance_types: [...string] | *["t3.medium", "t3a.medium"]

	// ex: vcluster-k3d-dfd-vc1
	env: metadata: name: "\(type)-\(ctx.name)"
}
