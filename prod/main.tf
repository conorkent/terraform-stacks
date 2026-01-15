terraform {
  required_version = ">= 1.13.0"
  
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "~> 3.0"
    }
  }
}

variable "region" {
  type        = string
  description = "AWS region"
  default     = "us-east-1"
}

variable "environment" {
  type        = string
  description = "Environment name (dev, prod, etc.)"
  default     = "prod"
}

provider "aws" {
  region = var.region
}

resource "aws_s3_bucket" "test" {
  bucket = "terraform-stacks-${var.environment}-${random_id.bucket_suffix.hex}"
  
  tags = {
    Name        = "Terraform Stacks ${title(var.environment)}"
    Environment = var.environment
    ManagedBy   = "terraform-stacks"
  }
}

resource "random_id" "bucket_suffix" {
  byte_length = 4
}

output "bucket_name" {
  value = aws_s3_bucket.test.bucket
}

output "bucket_arn" {
  value = aws_s3_bucket.test.arn
}
