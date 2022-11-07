package workflows

#EarthlyJob: {
	"runs-on": "ubuntu-latest"

	env: FORCE_COLOR: 1

	steps: [...{...}]

	needs?: [...string] | string
}

#EarthlySteps: #DockerLoginSteps + [{
	name: "Download latest earthly"
	run: """
		sudo /bin/sh -c 'wget -q https://github.com/earthly/earthly/releases/download/v0.6.28/earthly-linux-amd64 -O /usr/local/bin/earthly && chmod +x /usr/local/bin/earthly'
		"""
}, {
	name: "Set up QEMU"
	id:   "qemu"
	uses: "docker/setup-qemu-action@v1"
	with: {
		image:     "tonistiigi/binfmt:latest"
		platforms: "all"
	}
}]

#DockerLoginSteps: [{
	name: "Checkout code"
	uses: "actions/checkout@v3"
	with: "fetch-depth": 0
}, {
	name: "Put back the git branch into git (Earthly uses it for tagging)"
	run: """
		branch=""
		if [ -n "$GITHUB_HEAD_REF" ]; then
		  branch="$GITHUB_HEAD_REF"
		  echo "BRANCH=PR-${{github.event.pull_request.number}}" >> $GITHUB_ENV
		  echo "TAG=PR-${{github.event.pull_request.number}}" >> $GITHUB_ENV
		else
		  branch="${GITHUB_REF_NAME}"
		  echo "BRANCH=${branch}" >> $GITHUB_ENV
		  echo "TAG=latest" >> $GITHUB_ENV
		fi

		git checkout -b "$branch" || true

		if [ "$GITHUB_REF_TYPE" = "tag" ]; then
		  echo "BRANCH=${GITHUB_REF_NAME}" >> $GITHUB_ENV
		  echo "TAG=${GITHUB_REF_NAME}" >> $GITHUB_ENV
		fi

		git config --unset http.https://github.com/.extraheader
		"""
}, {
	name: "Login to Packages Container registry"
	uses: "docker/login-action@v1"
	with: {
		registry: "ghcr.io"
		username: "not-used"
		password: "${{ secrets.GITHUB_TOKEN }}"
	}
}]
