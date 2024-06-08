
resource "random_string" "bucket_name" {
  lower = true
  upper =  false
  length   = 16
  special  = false
}

resource "aws_s3_bucket" "example" {
  bucket = random_string.bucket_name.result

  tags = {
    UserUuid  = "vars.user_uuid"
  }

}