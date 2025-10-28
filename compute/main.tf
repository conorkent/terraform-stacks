terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.region
}

variable "vpc_id" {
  description = "ID of the VPC where instances will be deployed"
  type        = string
}

variable "region" {
  description = "AWS region to use for the provider"
  type        = string
  default     = "us-west-2"
}

resource "aws_instance" "web" {
  ami           = "ami-0c55b159cbfafe1f0" # Amazon Linux
  instance_type = "t3.micro"
  subnet_id     = null # you’ll refine this later
  tags = { Name = "stack-demo-instance" }
}

output "instance_id" {
  value = aws_instance.web.id
}
