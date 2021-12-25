terraform {
  required_providers {
    yandex = {
      source  = "yandex-cloud/yandex"
      version = "0.61.0"
    }
  }
}

provider "yandex" {
  token     = "AQAAAAA241WeAATuwfdGrzE74krGtR1QsxCCc0c"
  cloud_id  = "b1gku58k1k8tvqul8qbq"
  folder_id = "b1g82tojtne9luuk16s5"
  zone      = "ru-central1-a"
}
