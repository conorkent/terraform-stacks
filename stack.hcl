terraform {
  required_version = ">= 1.13.0"
  
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

variable "region" {
  type        = string
  description = "AWS region for resources"
}

component "example" {
  source = "./example"
  
  inputs = {
    region = var.region
  }
}
