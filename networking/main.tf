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

resource "aws_vpc" "main" {
  cidr_block = "10.0.0.0/16"
  tags = { Name = "stack-demo-vpc" }
}

output "vpc_id" {
  value = aws_vpc.main.id
}

variable "region" {
  description = "AWS region to create resources in"
  type        = string
  default     = "us-east-1"
}
