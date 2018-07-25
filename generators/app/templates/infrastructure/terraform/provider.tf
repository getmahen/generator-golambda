provider "aws" {
  region  = "us-east-2"
  profile = "credo-auth"
  assume_role {
    role_arn = "arn:aws:iam::${lookup(var.env_to_acct_id, var.environment)}:role/terraform"
  }
}

variable "env_to_acct_id" {
  type  = "map"
  default = {
    dev     = "674346455231"
    qa      = "772404289823"
    prod    = "465292320167"
  }
}