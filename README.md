# Terraform Custom Provider - Stackbill

The goal of this project is to create Stackbill resources via terraform

## Description

This would help us to create the stackbill resources via Terraform.

## Requirements
- [Terraform](https://www.terraform.io/downloads.html) 1.1.x
- [Go](https://golang.org/doc/install) 1.17 (to build the provider plugin)

[![N|Stackbill](https://www.stackbill.com/wp-content/uploads/2017/11/stackbill-logo-white.png)](https://www.stackbill.com/)

## Usage
- Clone this repository
- Run install.sh file

## Examples
Provider configuration

```
# This is required for Terraform 0.13+
terraform {
  required_providers {
    stackbill = {
      source = "stackbill/stackbill"
      version = "0.1.0"
    }
  }
}

provider "stackbill" {
  api_key    = "Place your api key here"
  secret_key = "Place your secret key here"
  url = "Place your api url here"
}
```

# Resource example - Instance Creation
```
resource "stackbill_instance" "my-server" {
  compute_offering_uuid = "2f7b1b02-078f-42fb-aece-aa21bfa89e1a"
  cpu_core              = "3"
  # disk_size             = 0
  hypervisor_name       = "string"
  # memory                = "0"
  name                  = "AzTestingVmFour"
  network_uuid          = "0c7cc302-abe5-4687-8351-26aba55b3d0d"
  # security_group_name   = "string"
  # ssh_key_name          = "string"
  # storage_offering_uuid = "string"
  template_uuid         = "e39172b4-715b-44ac-ada2-7f97f733eecc"
  zone_uuid             = "74b12720-73ce-49b6-857f-48cdac6dcd3f"
}
# - Optional values
```

# Dataset example - Compute offering list
```
data "stackbill_compute_offering_list" "all" {
  zone_uuid = "74b12720-73ce-49b6-857f-48cdac6dcd3f"
  #uuid = "c674ac49-32cd-4aae-96f8-25458bded6ad"
}

# Display compute offering list
output "all_compute_offerings" {
    value = data.stackbill_compute_offering_list.all.computeofferings
}
# - Optional values
```