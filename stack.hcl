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
  type = string
}

component "example" {
  source = "./example"
  
  providers = {
    aws = provider.aws
  }
  
  inputs = {
    region = var.region
  }
}

provider "aws" {
  region = var.region
}
