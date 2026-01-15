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

component "networking" {
  source = "./networking"
  
  providers = {
    aws = provider.aws.main
  }
}

component "compute" {
  source = "./compute"
  
  providers = {
    aws = provider.aws.main
  }
  
  inputs = {
    vpc_id = component.networking.vpc_id
  }
}

component "app" {
  source = "./app"
  
  providers = {
    aws = provider.aws.main
  }
  
  inputs = {
    instance_id = component.compute.instance_id
  }
}

provider "aws" "main" {
  config {
    region = var.region
  }
}
