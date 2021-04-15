
variable "region" {
  default = "us-south"
}

provider "ibm" {
  region  = var.region
}
