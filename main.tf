terraform {
  cloud {
    organization = "spaceships-moonlights"

    workspaces {
      name = "Terra-house-1"
    }
  }

  
  
  required_providers {
    random = {
      source = "hashicorp/random"
      version = "3.6.0"
    }    
    aws = {
      source = "hashicorp/aws"
      version = "5.44.0"
    }
  }
}

provider "aws" {
  # Configuration options
}
  

  
provider "random" {
  # Configuration options
}

resource "random_string" "bucket_name" {
  lower = true
  upper =  false
  length   = 16
  special  = false
}

resource "aws_s3_bucket" "example" {
  bucket = random_string.bucket_name.result

}

output "random_bucket_name_result" {
    value = random_string.bucket_name.result
  
}