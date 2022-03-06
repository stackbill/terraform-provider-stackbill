# Stackbill Provider

The Stackbill provider is used to interact with the resources supported by Stackbill. The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```
terraform {
  required_providers {
    stackbill = {
      version = "~> 1.0.0"
      source  = "stackbill.com/assistanz/stackbill"
    }
  }
}

provider "stackbill" {
  url    = "api URL"
  api_key    = "api key"
  secret_key = "api token"
}

```

## Argument Reference

The following arguments are supported:

- `url` - (Required) This is the Stackbill API URL. Alternatively, this can also be specified using environment variables ordered by precedence:
- `api_key` - (Required) This is the Stackbill API key. you can generate your API key in your account settings on the Stackbill dashboard
- `secret_key` - (Required) This is the Stackbill API secret token. you can generate your API secret token in your account settings on the Stackbill dashboard
