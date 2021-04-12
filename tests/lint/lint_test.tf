resource "aws_instance" "invalid_type" {
  instance_type = "t2.skynet"
}
