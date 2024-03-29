package m

import (
	"github.com/defn/dev/m/command"
)

input: organization: [NAME=string]: command.#AwsOrganization & {
	name: NAME
}

input: kubernetes: [NAME=string]: command.#KubernetesCluster & {
	name: NAME
}

input: organization: {
	defn: {
		region: "us-east-2"
	}
	gyre: {
		region:   "us-east-2"
		accounts: ops_accounts
	}
	curl: {
		region:   "us-west-2"
		accounts: env_accounts
	}
	coil: {
		region:   "us-east-1"
		accounts: env_accounts
	}
	helix: {
		region:   "us-east-2"
		accounts: full_accounts
		domain:   "defn.sh"
	}
	spiral: {
		region:   "us-west-2"
		accounts: full_accounts
	}
}

meh: {
	vault: {
		region: "us-east-2"
	}
	circus: {
		region: "us-west-2"
	}
	chamber: {
		region: "us-west-2"
	}
	whoa: {
		region: "us-west-2"
	}
	imma: {
		region: "us-west-2"
	}
	immanent: {
		region: "us-west-2"
	}
	jianghu: {
		region: "us-west-2"
	}
	fogg: {
		region: "us-west-2"
	}
}

input: kubernetes: {
	"gyre-1": {
		region: "us-east-2"
		nodegroup: "gyre-1-general": {
			instance_types: ["r5d.large", "r5ad.large", "r5dn.large"]
			az: {
				"us-east-2a": network: "10.249.0.0/22"
				"us-east-2b": network: "10.249.4.0/22"
				"us-east-2c": network: "10.249.8.0/22"
			}
		}
		vpc: {
			cidrs: [
				"10.249.0.0/19",
				// available: 10.249.12.0/22"
			]
		}
	}
}
