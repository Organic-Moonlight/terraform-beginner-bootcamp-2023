terraform {
  #cloud {
  #  organization = "spaceships-moonlights"

  #  workspaces {
  #    name = "Terra-house-1"
  #  }
  #}
  
  
  required_providers {    
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