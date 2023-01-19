// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package logs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_logs_destination", destinationDataSource)
}

// destinationDataSource returns the Terraform awscc_logs_destination data source.
// This Terraform data source corresponds to the CloudFormation AWS::Logs::Destination resource.
func destinationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: DestinationName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the destination resource",
		//	  "maxLength": 512,
		//	  "minLength": 1,
		//	  "pattern": "^[^:*]{1,512}$",
		//	  "type": "string"
		//	}
		"destination_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the destination resource",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DestinationPolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An IAM policy document that governs which AWS accounts can create subscription filters against this destination.",
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"destination_policy": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An IAM policy document that governs which AWS accounts can create subscription filters against this destination.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of an IAM role that permits CloudWatch Logs to send data to the specified AWS resource",
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of an IAM role that permits CloudWatch Logs to send data to the specified AWS resource",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TargetArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the physical target where the log events are delivered (for example, a Kinesis stream)",
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"target_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the physical target where the log events are delivered (for example, a Kinesis stream)",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Logs::Destination",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Logs::Destination").WithTerraformTypeName("awscc_logs_destination")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                "Arn",
		"destination_name":   "DestinationName",
		"destination_policy": "DestinationPolicy",
		"role_arn":           "RoleArn",
		"target_arn":         "TargetArn",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}