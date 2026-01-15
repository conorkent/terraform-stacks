terraform {
  required_version = ">= 1.8.0"
  
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

variable "vpc_id" {
  description = "ID of the VPC where instances will be deployed"
  type        = string
}

data "aws_availability_zones" "available" {
  state = "available"
}

resource "aws_subnet" "main" {
  vpc_id            = var.vpc_id
  cidr_block        = "10.0.1.0/24"
  availability_zone = data.aws_availability_zones.available.names[0]
  tags = { Name = "stack-demo-subnet" }
}

data "aws_ami" "amazon_linux" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "name"
    values = ["amzn2-ami-hvm-*-x86_64-gp2"]
  }
}

resource "aws_instance" "web" {
  ami           = data.aws_ami.amazon_linux.id
  instance_type = "t3.micro"
  subnet_id     = aws_subnet.main.id
  tags = { Name = "stack-demo-instance" }
}

output "instance_id" {
  value = aws_instance.web.id
}
