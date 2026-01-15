terraform {
  required_version = ">= 1.13.0"
  
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

component "example" {
  source = "./example"
  
  inputs = {
    environment = var.environment
  }
}

variable "environment" {
  type        = string
  description = "Environment name"
}
