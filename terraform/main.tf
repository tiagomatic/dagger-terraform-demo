// A trivial resource example
resource "null_resource" "example" {
  provisioner "local-exec" {
    command = "echo Hello from an example null_resource!"
  }
}