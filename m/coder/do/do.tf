locals {
  coder_name = "coder-${data.coder_workspace.me.owner}-${data.coder_workspace.me.name}"
}

resource "digitalocean_volume" "home_volume" {
  region                   = data.coder_parameter.region.value
  name                     = "coder-${data.coder_workspace.me.id}-home"
  size                     = data.coder_parameter.home_volume_size.value
  initial_filesystem_type  = "ext4"
  initial_filesystem_label = "coder-home"

  lifecycle {
    ignore_changes = all
  }
}

resource "digitalocean_volume" "nix_volume" {
  region                   = data.coder_parameter.region.value
  name                     = "coder-${data.coder_workspace.me.id}-nix"
  size                     = data.coder_parameter.nix_volume_size.value
  initial_filesystem_type  = "ext4"
  initial_filesystem_label = "coder-nix"

  lifecycle {
    ignore_changes = all
  }
}

resource "digitalocean_droplet" "workspace" {
  count = data.coder_workspace.me.start_count

  name = "coder-${data.coder_workspace.me.owner}-${data.coder_workspace.me.name}"

  region = data.coder_parameter.region.value

  image = data.coder_parameter.droplet_image.value
  size  = data.coder_parameter.droplet_size.value

  volume_ids = [
    digitalocean_volume.home_volume.id,
    digitalocean_volume.nix_volume.id
  ]

  user_data = templatefile("cloud-config.yaml.tftpl", {
    home_volume_label = digitalocean_volume.home_volume.initial_filesystem_label
    nix_volume_label  = digitalocean_volume.nix_volume.initial_filesystem_label
    init_script       = base64encode(coder_agent.main.init_script)
    coder_agent_token = coder_agent.main.token
  })
}