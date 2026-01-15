terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

variable "instance_id" {
  type = string
}

output "message" {
  value = "App layer deployed on instance: ${var.instance_id}"
}
