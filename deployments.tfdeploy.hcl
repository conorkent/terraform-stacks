deployment "dev" {
  inputs = {
    region = "us-east-1"
  }
}

deployment "prod" {
  inputs = {
    region = "us-west-2"
  }
}
