resource "aws_s3_bucket" "bucket" {
    bucket = "kvetalexbucket-${count.index}-${terraform.workspace}"
    acl    = "private"
    versioning {
      enabled = true
    }
    tags = {
      Name = "Bucket ${count.index}"
      Environment = terraform.workspace
    }
    count = "${local.workspace["instance_count"]}"
}

resource "aws_dynamodb_table" "terraform_state_locking_dynamodb" {
  name = "terraform-state-locking"
  hash_key = "LockID"
  read_capacity = 20
  write_capacity = 20

  attribute {
    name = "LockID"
    type = "S"
  }

  tags = {
    Name = "Terraform State File Locking"
  }
}

resource "aws_vpc" "main" {
  cidr_block = "10.0.0.0/16"

  tags = {
    Name = "my-terraform-aws-vpc"
  }
}

terraform {
  backend "s3" {
    bucket = "kvetalexbucket-default"
    key    = "workspaces/dev/terraform_dev.tfstate"
    region = "us-east-2"
  }
}

resource "aws_instance" "my-comp" {
    ami = "${local.workspace["ami"]}"
    instance_type = "${local.workspace["instance_type"]}"
    count = "${local.workspace["instance_count"]}"
    tags = {
      Name = "my-comp-${count.index}"
    }
    lifecycle {
      create_before_destroy = true
    }
}

locals {
  env = {
    default = {
      instance_type  = "t2.micro"
      ami            = "ami-0ff8a91507f77f867"
      instance_count = 1
      }
    stage = {
      instance_type  = "t2.micro"
      ami            = "ami-0c55b159cbfafe1f0"
      instance_count = 1
    }
    prod = {
      instance_type  = "t2.micro"
      ami            = "ami-002068ed284fb165b"
      instance_count = 2
    }
  }
  environmentvars = "${contains(keys(local.env),
terraform.workspace) ? terraform.workspace : "default"}"
  workspace       = "${merge(local.env["default"],
local.env[local.environmentvars])}"
}

locals {
  backets_ids = toset([
    "e1",
    "e2",
  ])
}

resource "aws_s3_bucket" "bucket_each" {
  for_each = local.backets_ids
  bucket = "kvetalexbucket-${each.key}-${terraform.workspace}"
  acl    = "private"
  tags = {
    Name        = "Bucket ${each.key}"
    Environment = terraform.workspace
  }
}
