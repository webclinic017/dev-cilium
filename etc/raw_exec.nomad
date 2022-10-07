job "raw_exec" {
  datacenters = ["dev"]

  group "shell" {
    task "server" {
      driver = "raw_exec"

      config {
        command = "sleep"
        args = [
          "86400"
        ]
      }

      template {
        data = <<EOF
meh
EOF
        destination = "secrets/secrets.txt"
      }
    }
  }
}
