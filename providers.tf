provider "google" {
  project = var.project_id
  region  = var.region
  zone    = var.zone
}

terraform {
  backend "gcs" {
    bucket = "tf-state-prod-edemdevsecops-1"
    prefix = "terraform/state"
  }
}