variable "instance_id" {
  type = string
}

output "message" {
  value = "App layer deployed on instance: ${var.instance_id}"
}
