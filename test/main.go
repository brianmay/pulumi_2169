package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		test, err := s3.NewBucketV2(ctx, "test", &s3.BucketV2Args{
			Bucket: pulumi.String("12345678-test"),
		})
		if err != nil {
			return err
		}
		assumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"lambda.amazonaws.com",
							},
						},
					},
					Actions: []string{
						"sts:AssumeRole",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		testUpdate, err := iam.NewRole(ctx, "test_update", &iam.RoleArgs{
			Name:             pulumi.String("test_update"),
			AssumeRolePolicy: pulumi.String(assumeRole.Json),
		})
		if err != nil {
			return err
		}
		testS3 := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Effect: pulumi.String("Allow"),
					Actions: pulumi.StringArray{
						pulumi.String("s3:List*"),
						pulumi.String("s3:Get*"),
						pulumi.String("s3:PutObject"),
						pulumi.String("s3:DeleteObject"),
					},
					Resources: pulumi.StringArray{
						test.Arn.ApplyT(func(arn string) (string, error) {
							return fmt.Sprintf("%v/*", arn), nil
						}).(pulumi.StringOutput),
						test.Arn,
					},
				},
			},
		}, nil)
		testS3Policy, err := iam.NewPolicy(ctx, "test_s3", &iam.PolicyArgs{
			Name:        pulumi.String("test_s3"),
			Path:        pulumi.String("/"),
			Description: pulumi.String("IAM policy for accessing test s3 bucket"),
			Policy: testS3.ApplyT(func(testS3 iam.GetPolicyDocumentResult) (*string, error) {
				return &testS3.Json, nil
			}).(pulumi.StringPtrOutput),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "test_s3", &iam.RolePolicyAttachmentArgs{
			Role:      testUpdate.Name,
			PolicyArn: testS3Policy.Arn,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
