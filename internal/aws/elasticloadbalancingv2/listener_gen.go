// Code generated by generators/resource/main.go; DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("aws_elasticloadbalancingv2_listener", listenerResourceType)
}

// listenerResourceType returns the Terraform aws_elasticloadbalancingv2_listener resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::ElasticLoadBalancingV2::Listener resource type.
func listenerResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]schema.Attribute{
		"alpn_policy": {
			// Property: AlpnPolicy
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Optional: true,
		},
		"certificates": {
			// Property: Certificates
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "CertificateArn": {
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			// Ordered set.
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"certificate_arn": {
						// Property: CertificateArn
						Type:     types.StringType,
						Optional: true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Optional: true,
		},
		"default_actions": {
			// Property: DefaultActions
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "AuthenticateCognitoConfig": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "AuthenticationRequestExtraParams": {
			//             "patternProperties": {
			//               "": {
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "OnUnauthenticatedRequest": {
			//             "type": "string"
			//           },
			//           "Scope": {
			//             "type": "string"
			//           },
			//           "SessionCookieName": {
			//             "type": "string"
			//           },
			//           "SessionTimeout": {
			//             "type": "string"
			//           },
			//           "UserPoolArn": {
			//             "type": "string"
			//           },
			//           "UserPoolClientId": {
			//             "type": "string"
			//           },
			//           "UserPoolDomain": {
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "UserPoolClientId",
			//           "UserPoolDomain",
			//           "UserPoolArn"
			//         ],
			//         "type": "object"
			//       },
			//       "AuthenticateOidcConfig": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "AuthenticationRequestExtraParams": {
			//             "patternProperties": {
			//               "": {
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "AuthorizationEndpoint": {
			//             "type": "string"
			//           },
			//           "ClientId": {
			//             "type": "string"
			//           },
			//           "ClientSecret": {
			//             "type": "string"
			//           },
			//           "Issuer": {
			//             "type": "string"
			//           },
			//           "OnUnauthenticatedRequest": {
			//             "type": "string"
			//           },
			//           "Scope": {
			//             "type": "string"
			//           },
			//           "SessionCookieName": {
			//             "type": "string"
			//           },
			//           "SessionTimeout": {
			//             "type": "string"
			//           },
			//           "TokenEndpoint": {
			//             "type": "string"
			//           },
			//           "UserInfoEndpoint": {
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "TokenEndpoint",
			//           "Issuer",
			//           "ClientSecret",
			//           "UserInfoEndpoint",
			//           "ClientId",
			//           "AuthorizationEndpoint"
			//         ],
			//         "type": "object"
			//       },
			//       "FixedResponseConfig": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "ContentType": {
			//             "type": "string"
			//           },
			//           "MessageBody": {
			//             "type": "string"
			//           },
			//           "StatusCode": {
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "StatusCode"
			//         ],
			//         "type": "object"
			//       },
			//       "ForwardConfig": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "TargetGroupStickinessConfig": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "DurationSeconds": {
			//                 "type": "integer"
			//               },
			//               "Enabled": {
			//                 "type": "boolean"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "TargetGroups": {
			//             "items": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "TargetGroupArn": {
			//                   "type": "string"
			//                 },
			//                 "Weight": {
			//                   "type": "integer"
			//                 }
			//               },
			//               "type": "object"
			//             },
			//             "type": "array",
			//             "uniqueItems": true
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "Order": {
			//         "type": "integer"
			//       },
			//       "RedirectConfig": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Host": {
			//             "type": "string"
			//           },
			//           "Path": {
			//             "type": "string"
			//           },
			//           "Port": {
			//             "type": "string"
			//           },
			//           "Protocol": {
			//             "type": "string"
			//           },
			//           "Query": {
			//             "type": "string"
			//           },
			//           "StatusCode": {
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "StatusCode"
			//         ],
			//         "type": "object"
			//       },
			//       "TargetGroupArn": {
			//         "type": "string"
			//       },
			//       "Type": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Type"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			// Ordered set.
			Attributes: schema.ListNestedAttributes(
				map[string]schema.Attribute{
					"authenticate_cognito_config": {
						// Property: AuthenticateCognitoConfig
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"authentication_request_extra_params": {
									// Property: AuthenticationRequestExtraParams
									// Pattern: ""
									Type:     types.MapType{ElemType: types.StringType},
									Optional: true,
								},
								"on_unauthenticated_request": {
									// Property: OnUnauthenticatedRequest
									Type:     types.StringType,
									Optional: true,
								},
								"scope": {
									// Property: Scope
									Type:     types.StringType,
									Optional: true,
								},
								"session_cookie_name": {
									// Property: SessionCookieName
									Type:     types.StringType,
									Optional: true,
								},
								"session_timeout": {
									// Property: SessionTimeout
									Type:     types.StringType,
									Optional: true,
								},
								"user_pool_arn": {
									// Property: UserPoolArn
									Type:     types.StringType,
									Required: true,
								},
								"user_pool_client_id": {
									// Property: UserPoolClientId
									Type:     types.StringType,
									Required: true,
								},
								"user_pool_domain": {
									// Property: UserPoolDomain
									Type:     types.StringType,
									Required: true,
								},
							},
						),
						Optional: true,
					},
					"authenticate_oidc_config": {
						// Property: AuthenticateOidcConfig
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"authentication_request_extra_params": {
									// Property: AuthenticationRequestExtraParams
									// Pattern: ""
									Type:     types.MapType{ElemType: types.StringType},
									Optional: true,
								},
								"authorization_endpoint": {
									// Property: AuthorizationEndpoint
									Type:     types.StringType,
									Required: true,
								},
								"client_id": {
									// Property: ClientId
									Type:     types.StringType,
									Required: true,
								},
								"client_secret": {
									// Property: ClientSecret
									Type:     types.StringType,
									Required: true,
								},
								"issuer": {
									// Property: Issuer
									Type:     types.StringType,
									Required: true,
								},
								"on_unauthenticated_request": {
									// Property: OnUnauthenticatedRequest
									Type:     types.StringType,
									Optional: true,
								},
								"scope": {
									// Property: Scope
									Type:     types.StringType,
									Optional: true,
								},
								"session_cookie_name": {
									// Property: SessionCookieName
									Type:     types.StringType,
									Optional: true,
								},
								"session_timeout": {
									// Property: SessionTimeout
									Type:     types.StringType,
									Optional: true,
								},
								"token_endpoint": {
									// Property: TokenEndpoint
									Type:     types.StringType,
									Required: true,
								},
								"user_info_endpoint": {
									// Property: UserInfoEndpoint
									Type:     types.StringType,
									Required: true,
								},
							},
						),
						Optional: true,
					},
					"fixed_response_config": {
						// Property: FixedResponseConfig
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"content_type": {
									// Property: ContentType
									Type:     types.StringType,
									Optional: true,
								},
								"message_body": {
									// Property: MessageBody
									Type:     types.StringType,
									Optional: true,
								},
								"status_code": {
									// Property: StatusCode
									Type:     types.StringType,
									Required: true,
								},
							},
						),
						Optional: true,
					},
					"forward_config": {
						// Property: ForwardConfig
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"target_group_stickiness_config": {
									// Property: TargetGroupStickinessConfig
									Attributes: schema.SingleNestedAttributes(
										map[string]schema.Attribute{
											"duration_seconds": {
												// Property: DurationSeconds
												Type:     types.NumberType,
												Optional: true,
											},
											"enabled": {
												// Property: Enabled
												Type:     types.BoolType,
												Optional: true,
											},
										},
									),
									Optional: true,
								},
								"target_groups": {
									// Property: TargetGroups
									// Ordered set.
									Attributes: schema.ListNestedAttributes(
										map[string]schema.Attribute{
											"target_group_arn": {
												// Property: TargetGroupArn
												Type:     types.StringType,
												Optional: true,
											},
											"weight": {
												// Property: Weight
												Type:     types.NumberType,
												Optional: true,
											},
										},
										schema.ListNestedAttributesOptions{},
									),
									Optional: true,
								},
							},
						),
						Optional: true,
					},
					"order": {
						// Property: Order
						Type:     types.NumberType,
						Optional: true,
					},
					"redirect_config": {
						// Property: RedirectConfig
						Attributes: schema.SingleNestedAttributes(
							map[string]schema.Attribute{
								"host": {
									// Property: Host
									Type:     types.StringType,
									Optional: true,
								},
								"path": {
									// Property: Path
									Type:     types.StringType,
									Optional: true,
								},
								"port": {
									// Property: Port
									Type:     types.StringType,
									Optional: true,
								},
								"protocol": {
									// Property: Protocol
									Type:     types.StringType,
									Optional: true,
								},
								"query": {
									// Property: Query
									Type:     types.StringType,
									Optional: true,
								},
								"status_code": {
									// Property: StatusCode
									Type:     types.StringType,
									Required: true,
								},
							},
						),
						Optional: true,
					},
					"target_group_arn": {
						// Property: TargetGroupArn
						Type:     types.StringType,
						Optional: true,
					},
					"type": {
						// Property: Type
						Type:     types.StringType,
						Required: true,
					},
				},
				schema.ListNestedAttributesOptions{},
			),
			Required: true,
		},
		"listener_arn": {
			// Property: ListenerArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"load_balancer_arn": {
			// Property: LoadBalancerArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			// LoadBalancerArn is a force-new attribute.
		},
		"port": {
			// Property: Port
			// CloudFormation resource type schema:
			// {
			//   "type": "integer"
			// }
			Type:     types.NumberType,
			Optional: true,
		},
		"protocol": {
			// Property: Protocol
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
		},
		"ssl_policy": {
			// Property: SslPolicy
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Optional: true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = schema.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::ElasticLoadBalancingV2::Listener",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::ElasticLoadBalancingV2::Listener").WithTerraformTypeName("aws_elasticloadbalancingv2_listener").WithTerraformSchema(schema)

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/DefaultActions/*/AuthenticateOidcConfig/ClientSecret",
		"/properties/DefaultActions/*/ForwardConfig",
		"/properties/DefaultActions/*/TargetGroupArn",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "aws_elasticloadbalancingv2_listener", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}