variable "vpc_id" {
  description = "ID of the VPC where instances will be deployed"
  type        = string
}

resource "aws_instance" "web" {
  ami           = "ami-0c55b159cbfafe1f0" # Amazon Linux
  instance_type = "t3.micro"
  subnet_id     = null # refine this later
  tags = { Name = "stack-demo-instance" }
}

output "instance_id" {
  value = aws_instance.web.id
}
