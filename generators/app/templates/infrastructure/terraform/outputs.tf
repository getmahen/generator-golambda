output "lambda_arn" {
  value = "${aws_lambda_function.ordersettlement.arn}"
}

output "lambda_function_name" {
  value = "${aws_lambda_function.ordersettlement.function_name}"
}

output "lambda_function_version_metadata" {
  value = "${data.aws_s3_bucket_object.ordersettlement_pkg.metadata}"
}