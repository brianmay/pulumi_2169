# Pulumi Test case for #2169

Simple test case for https://github.com/pulumi/pulumi-terraform-bridge/issues/2169

terraform files in root.

Pulumni files in test directory.

To reproduce:

```
$ vim main.tf   # update name of bucket, will probably conflict.
$ terraform apply
data.aws_iam_policy_document.assume_role: Reading...
data.aws_iam_policy_document.assume_role: Read complete after 0s [id=2690255455]
aws_s3_bucket.test: Refreshing state... [id=12345678-test]
aws_iam_role.test_update: Refreshing state... [id=test_update]
data.aws_iam_policy_document.test_s3: Reading...
data.aws_iam_policy_document.test_s3: Read complete after 0s [id=315358420]

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # aws_iam_policy.test_s3 will be created
  + resource "aws_iam_policy" "test_s3" {
      + arn              = (known after apply)
      + attachment_count = (known after apply)
      + description      = "IAM policy for accessing test s3 bucket"
      + id               = (known after apply)
      + name             = "test_s3"
      + name_prefix      = (known after apply)
      + path             = "/"
      + policy           = jsonencode(
            {
              + Statement = [
                  + {
                      + Action   = [
                          + "s3:PutObject",
                          + "s3:List*",
                          + "s3:Get*",
                          + "s3:DeleteObject",
                        ]
                      + Effect   = "Allow"
                      + Resource = [
                          + "arn:aws:s3:::12345678-test/*",
                          + "arn:aws:s3:::12345678-test",
                        ]
                    },
                ]
              + Version   = "2012-10-17"
            }
        )
      + policy_id        = (known after apply)
      + tags_all         = (known after apply)
    }

  # aws_iam_role_policy_attachment.test_s3 will be created
  + resource "aws_iam_role_policy_attachment" "test_s3" {
      + id         = (known after apply)
      + policy_arn = (known after apply)
      + role       = "test_update"
    }

Plan: 2 to add, 0 to change, 0 to destroy.

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes

aws_iam_policy.test_s3: Creating...
aws_iam_policy.test_s3: Creation complete after 2s [id=arn:aws:iam::136983895179:policy/test_s3]
aws_iam_role_policy_attachment.test_s3: Creating...
aws_iam_role_policy_attachment.test_s3: Creation complete after 0s [id=test_update-20240710003534050900000001]

Apply complete! Resources: 2 added, 0 changed, 0 destroyed.

bash-5.2$ cd test
bash-5.2$ pulumi stack init test
Created stack 'test'
Enter your passphrase to protect config/secrets:
Re-enter your passphrase to confirm:
bash-5.2$ pulumi import --from terraform ../terraform.tfstate
Enter your passphrase to unlock config/secrets
    (set PULUMI_CONFIG_PASSPHRASE or PULUMI_CONFIG_PASSPHRASE_FILE to remember):
Enter your passphrase to unlock config/secrets
Previewing import (test):
     Type                             Name         Plan       Info
 +   pulumi:pulumi:Stack              simple-test  create     2 warnings
 =   ├─ aws:iam:Role                  test_update  import
 =   ├─ aws:s3:BucketV2               test         import     4 warnings
 =   ├─ aws:iam:Policy                test_s3      import
 =   └─ aws:iam:RolePolicyAttachment  test_s3      import

Diagnostics:
  pulumi:pulumi:Stack (simple-test):
    warning: using pulumi-language-go from $PATH at /nix/store/j1kjljhp66ia1ml9ssj47mglm25v09pf-pulumi-language-go-3.122.0/bin/pulumi-language-go
    warning: using pulumi-language-go from $PATH at /nix/store/j1kjljhp66ia1ml9ssj47mglm25v09pf-pulumi-language-go-3.122.0/bin/pulumi-language-go

  aws:s3:BucketV2 (test):
    warning: urn:pulumi:test::simple::aws:s3/bucketV2:BucketV2::test verification warning: Use the aws_s3_bucket_versioning resource instead
    warning: urn:pulumi:test::simple::aws:s3/bucketV2:BucketV2::test verification warning: Use the aws_s3_bucket_server_side_encryption_configuration resource instead
    warning: urn:pulumi:test::simple::aws:s3/bucketV2:BucketV2::test verification warning: Use the aws_s3_bucket_acl resource instead
    warning: urn:pulumi:test::simple::aws:s3/bucketV2:BucketV2::test verification warning: Use the aws_s3_bucket_request_payment_configuration resource instead

Resources:
    + 1 to create
    = 4 to import
    5 changes

Do you want to perform this import? yes
Importing (test):
     Type                             Name         Status            Info
 +   pulumi:pulumi:Stack              simple-test  created (8s)      2 warnings
 =   ├─ aws:iam:Role                  test_update  imported (3s)
 =   ├─ aws:iam:Policy                test_s3      imported (3s)
 =   ├─ aws:s3:BucketV2               test         imported (5s)     4 warnings
 =   └─ aws:iam:RolePolicyAttachment  test_s3      imported (3s)

Diagnostics:
  pulumi:pulumi:Stack (simple-test):
    warning: using pulumi-language-go from $PATH at /nix/store/j1kjljhp66ia1ml9ssj47mglm25v09pf-pulumi-language-go-3.122.0/bin/pulumi-language-go
    warning: using pulumi-language-go from $PATH at /nix/store/j1kjljhp66ia1ml9ssj47mglm25v09pf-pulumi-language-go-3.122.0/bin/pulumi-language-go

  aws:s3:BucketV2 (test):
    warning: urn:pulumi:test::simple::aws:s3/bucketV2:BucketV2::test verification warning: Use the aws_s3_bucket_versioning resource instead
    warning: urn:pulumi:test::simple::aws:s3/bucketV2:BucketV2::test verification warning: Use the aws_s3_bucket_server_side_encryption_configuration resource instead
    warning: urn:pulumi:test::simple::aws:s3/bucketV2:BucketV2::test verification warning: Use the aws_s3_bucket_request_payment_configuration resource instead
    warning: urn:pulumi:test::simple::aws:s3/bucketV2:BucketV2::test verification warning: Use the aws_s3_bucket_acl resource instead

Resources:
    + 1 created
    = 4 imported
    5 changes

Duration: 9s

error: anonymous.pp:0,0-10,2: "test_s3" already declared; "test_s3" already declared
```

Note that preview appears to show non-existant changes, maybe because of
that error:

```
bash-5.2$ pulumi preview
Enter your passphrase to unlock config/secrets
    (set PULUMI_CONFIG_PASSPHRASE or PULUMI_CONFIG_PASSPHRASE_FILE to remember):
Enter your passphrase to unlock config/secrets
Previewing update (test):
     Type                             Name         Plan     Info
     pulumi:pulumi:Stack              simple-test           2 warnings
     ├─ aws:s3:BucketV2               test                  [diff: +forceDestroy-grants,requestPayer,serverSideEncryptionConfigurations,versionings~__defaults,protect]
     ├─ aws:iam:Role                  test_update           [diff: +forceDetachPolicies,maxSessionDuration,path-managedPolicyArns~__defaults,assumeRolePolicy,protect]
     ├─ aws:iam:Policy                test_s3               [diff: +path~policy,protect]
     └─ aws:iam:RolePolicyAttachment  test_s3               [diff: ~protect]

Diagnostics:
  pulumi:pulumi:Stack (simple-test):
    warning: using pulumi-language-go from $PATH at /nix/store/j1kjljhp66ia1ml9ssj47mglm25v09pf-pulumi-language-go-3.122.0/bin/pulumi-language-go
    warning: using pulumi-language-go from $PATH at /nix/store/j1kjljhp66ia1ml9ssj47mglm25v09pf-pulumi-language-go-3.122.0/bin/pulumi-language-go

Resources:
    5 unchanged
```
