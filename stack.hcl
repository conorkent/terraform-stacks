required_providers {
  aws = {
    source  = "hashicorp/aws"
    version = "~> 5.0"
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
}

component "app" {
  source = "./app"
  
  providers = {
    aws = provider.aws.main
  }
}

provider "aws" "main" {
  config {
    region = var.region
  }
}
