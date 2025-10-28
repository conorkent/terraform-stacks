terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

variable "region" {
  description = "AWS region"
  type        = string
}

provider "aws" {
  region = var.region
}

variable "instance_id" {}

output "message" {
  value = "App layer deployed on instance: ${var.instance_id}"
}
