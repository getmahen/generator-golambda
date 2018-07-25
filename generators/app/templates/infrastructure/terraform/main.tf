module "lambda_tags" {
  source = "modules/tags"
  description = "ADD HUMAN READABLE DATA HERE"
  role = "add_machine_readable_data_here"
  environment = "${var.environment}"
  project = "ordersettlement"
}

resource "aws_lambda_function" "ordersettlement" {
  s3_bucket         = "${data.aws_s3_bucket_object.ordersettlement_pkg.bucket}"
  s3_key            = "${data.aws_s3_bucket_object.ordersettlement_pkg.key}"
  s3_object_version = "${data.aws_s3_bucket_object.ordersettlement_pkg.version_id}"
  function_name     = "ordersettlement"
  runtime           = "go1.x"
  handler           = "ordersettlement"
  role              = "${aws_iam_role.ordersettlement_role.arn}"
  timeout           = 10

  vpc_config {
    security_group_ids = ["${data.aws_security_group.docker.id}"]
    subnet_ids         = ["${data.aws_subnet_ids.application.ids}"]
  }

  environment {
    variables = {
      ENVIRONMENT      = "${var.environment}"
      LOG_LEVEL        = "INFO"
      CONSUL_URL       = "https://consul.credo${var.environment}.dev"
    }
  }

  tags = "${module.lambda_tags.tags}"

}

resource "aws_iam_role" "ordersettlement_role" {
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

//example of using AWS managed policy MODIFY FOR YOUR USE CASE
resource "aws_iam_role_policy_attachment" "attachment" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaKinesisExecutionRole"
  role = "${aws_iam_role.ordersettlement_role.name}"
}

//example policy for vpc access MODIFY TO SUIT YOUR USE CASE
resource "aws_iam_policy" "ordersettlement_policy" {
  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "ec2:CreateNetworkInterface",
                "ec2:DescribeNetworkInterfaces",
                "ec2:DeleteNetworkInterface"
            ],
            "Resource": "*"
        }
    ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "vpc_policy_attachment" {
  policy_arn = "${aws_iam_policy.ordersettlement_policy.arn}"
  role = "${aws_iam_role.ordersettlement_role.name}"
}


resource "aws_cloudwatch_log_group" "lambda_log_group" {
  name              = "/aws/lambda/${aws_lambda_function.ordersettlement.function_name}"
  retention_in_days = 7
}

module "log_forwarder" {
  source = "modules/loggroup_sumo_forwarder"
  environment = "${var.environment}"
  forwarder_name = "${aws_lambda_function.ordersettlement.function_name}"
  log_group_name = "${aws_cloudwatch_log_group.lambda_log_group.name}"
  source_category = "${var.environment}/aws/lambda/${aws_lambda_function.ordersettlement.function_name}"
}