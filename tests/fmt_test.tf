resource "aws_kms_key" "a" {}

resource "aws_kms_alias" "bad_format" {
  name  = "alias/my-key-alias"
  target_key_id = aws_kms_key.a.key_id
}
