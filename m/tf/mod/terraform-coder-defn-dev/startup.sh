  #!/usr/bin/env bash

  set -ex

  sudo dd if=/dev/zero of=/root/swap bs=1M count=4096
  sudo chmod 0600 /root/swap
  sudo mkswap /root/swap
  sudo swapon /root/swap || true

  sudo install -d -m 0700 -o ubuntu -g ubuntu /run/user/1000 /run/user/1000/gnupg
  sudo install -d -m 0700 -o ubuntu -g ubuntu /nix /nix

  if [[ ! -d "/nix/home/.git/." ]]; then
    ssh -o StrictHostKeyChecking=no git@github.com true || true
    git clone http://github.com/defn/dev /nix/home
    pushd /nix/home
    git reset --hard
    popd
  else
    pushd /nix/home
    git pull
    popd
  fi

  sudo rm -rf "$HOME"
  sudo ln -nfs /nix/home "$HOME"
  ssh -o StrictHostKeyChecking=no git@github.com true || true

  cd
  source .bash_profile
  make install
  
  cd m
  setsid ~/bin/nix/tilt up

  exit 0
