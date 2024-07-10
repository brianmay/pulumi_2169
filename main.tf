provider "aws" {
  region  = "us-east-2"
  profile = "sfp"
}

resource "aws_s3_bucket" "test" {
  bucket = "12345678-test"
}

data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }

    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role" "test_update" {
  name               = "test_update"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

data "aws_iam_policy_document" "test_s3" {
  statement {
    effect = "Allow"

    actions = [
      "s3:List*",
      "s3:Get*",
      "s3:PutObject",
      "s3:DeleteObject"
    ]

    resources = [
      "${aws_s3_bucket.test.arn}/*",
      "${aws_s3_bucket.test.arn}"
    ]
  }
}

resource "aws_iam_policy" "test_s3" {
  name        = "test_s3"
  path        = "/"
  description = "IAM policy for accessing test s3 bucket"
  policy      = data.aws_iam_policy_document.test_s3.json
}

resource "aws_iam_role_policy_attachment" "test_s3" {
  role       = aws_iam_role.test_update.name
  policy_arn = aws_iam_policy.test_s3.arn
}
