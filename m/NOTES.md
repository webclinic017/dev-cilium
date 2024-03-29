- https://buildkite.com/docs/agent/v3/configuration
- https://buildkite.com/docs/pipelines/environment-variables

# misc

```
docker run --privileged --rm tonistiigi/binfmt --install arm64

mkdir -p /tmp/etc/ssh
nix run .#ssh-keygen -- -A -f /tmp
nix run .#sshd -- -f ~/etc/sshd_config -o ListenAddress=$(nix run .#tailscale -- ip -4)

gpg --armor --export | ssh "$ssh_host" gpg --import
gpg --export-ownertrust | ssh "$ssh_host" gpg --import-ownertrust
ln -nfs ln -nfs /home/ubuntu/.gnupg/S.gpg-agent /home/ubuntu/.gnupg/S.gpg-agent.extra /run/user/1000/gnupg/ /run/user/1000/gnupg/

nix run .#ssh -- \
    -o StrictHostKeyChecking=no \
    -o UserKnownHostsFile=/dev/null \
    -o ServerAliveInterval=1 \
    -o ServerAliveCountMax=10 \
    -o ConnectTimeout=1 \
    -o ConnectionAttempts=5 \
    -o PasswordAuthentication=no \
    -o StreamLocalBindUnlink=yes \
    -o RemoteForward="/home/ubuntu/.gnupg/S.gpg-agent $HOME/.gnupg/S.gpg-agent.extra" \
    -o RemoteForward="/home/ubuntu/.gnupg/S.gpg-agent.extra $HOME/.gnupg/S.gpg-agent.extra" \
    -v -p 2222 -A ubuntu@$ssh_host \
        bash -c '"ln -nfs $SSH_AUTH_SOCK $HOME/.ssh/S.ssh-agent; ssh-add -L; echo connected; exec sleep infinity"'

```

```
nix profile install nixpkgs#nodejs-18_x
PATH=$HOME/.nix-profile/bin:$PATH
npm install -u @devcontainers/cli
devcontainer build --workspace-folder .

rm -f go/bin/derper
go install tailscale.com/cmd/derper@latest

cd .acme.sh/dir
ln -nfs the.crt derp.defn.run.crt
ln -nfs the.key derp.defn.run.key

go/bin/derper --hostname=derp.defn.run --certmode=manual --certdir=/Users/defn/.acme.sh/\*.defn.run -a=:3340 --stun=true --http-port=8080 --verify-clients=false -c=$HOME/etc/derper.conf

cd etc/openvn
easyrsa init-pki
easyrsa gen-dh
easyrsa build-ca
easyrsa build-server-full server
openssl rsa -in pki/private/server.key  -out pki/private/server2.key
sudo openvpn server.conf

acme.sh --register-account -m iam@defn.sh
acme.sh --issue --domain '*.defn.run' --dns --yes-I-know-dns-manual-mode-enough-go-ahead-please
```
