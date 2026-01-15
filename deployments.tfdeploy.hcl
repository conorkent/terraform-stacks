deployment "dev" {
}

deployment "prod" {
  inputs = {
    region = "us-west-2"
  }
}
