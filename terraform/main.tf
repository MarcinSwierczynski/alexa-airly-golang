terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}

module "lambda_function_existing_package_local" {
  source = "terraform-aws-modules/lambda/aws"

  function_name = "alexa-airly-golang-tf"
  description   = "Golang based lambda for Ailry Alexa skill"
  handler       = "bootstrap"
  runtime       = "provided.al2023"
  architectures = ["arm64"]

  create_package         = false
  local_existing_package = "../alexa-airly-golang.zip"

  tags = {
    Terraform = "yes"
  }
}

resource "aws_lambda_permission" "with_alexa" {
  statement_id  = "AllowExecutionFromAlexa"
  action        = "lambda:InvokeFunction"
  function_name = module.lambda_function_existing_package_local.lambda_function_name
  principal     = "alexa-appkit.amazon.com"
}
