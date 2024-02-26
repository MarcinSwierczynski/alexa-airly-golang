output "alexa-airly-golang-arn" {
  description = "The ARN of alexa-airly-golang Lambda"
  value       = module.lambda_function_existing_package_local.lambda_function_arn
}