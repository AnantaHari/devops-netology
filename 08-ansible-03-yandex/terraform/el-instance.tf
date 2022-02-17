resource "yandex_compute_instance" "el-instance" {
  name                      = "el-instance"
  zone                      = "ru-central1-a"
  hostname                  = "el-instance.netology.cloud"
  allow_stopping_for_update = true

  resources {
    cores  = 2
    memory = 4
  }

  boot_disk {
    initialize_params {
      image_id = var.centos-7-base
      name     = "root-el-instance"
      type     = "network-nvme"
      size     = "50"
    }
  }

  network_interface {
    subnet_id = yandex_vpc_subnet.default.id
    nat       = true
  }

  metadata = {
    ssh-keys = "centos:${file("~/.ssh/id_rsa.pub")}"
  }
}
