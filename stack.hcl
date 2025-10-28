stack "three-tier" {
  description = "Demo 3-tier application managed with HCP Terraform Stacks"
  source = "./"
  variables = {
    region = "us-east-1"
  }
}

workspace "networking" {
  description = "VPC and subnets"
  source = "./networking"
}

workspace "compute" {
  description = "EC2 instances"
  source = "./compute"
  depends_on = ["networking"]
}

workspace "app" {
  description = "App layer that depends on compute"
  source = "./app"
  depends_on = ["compute"]
}
