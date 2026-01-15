terraform {
  required_version = ">= 1.9.0"
  
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

variable "region" {
  type    = string
  default = "us-east-1"
}

component "example" {
  source = "./example"
  
  providers = {
    aws = provider.aws
  }
}

provider "aws" {
  region = var.region
}
