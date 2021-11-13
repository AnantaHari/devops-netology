packer {
  required_plugins {
    docker = {
      version = ">= 0.0.7"
      source  = "github.com/hashicorp/docker"
    }
  }
}

source "docker" "redhat8" {
  image  = "redhat/ubi8:latest"
  commit = true
}

build {
  name = "learn-packer"
  sources = [
    "source.docker.redhat8"
  ]
}
