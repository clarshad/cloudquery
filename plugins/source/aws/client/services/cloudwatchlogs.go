// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
)

//go:generate mockgen -package=mocks -destination=../mocks/cloudwatchlogs.go -source=cloudwatchlogs.go CloudwatchlogsClient
type CloudwatchlogsClient interface {
	DescribeAccountPolicies(context.Context, *cloudwatchlogs.DescribeAccountPoliciesInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeAccountPoliciesOutput, error)
	DescribeDeliveries(context.Context, *cloudwatchlogs.DescribeDeliveriesInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeDeliveriesOutput, error)
	DescribeDeliveryDestinations(context.Context, *cloudwatchlogs.DescribeDeliveryDestinationsInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeDeliveryDestinationsOutput, error)
	DescribeDeliverySources(context.Context, *cloudwatchlogs.DescribeDeliverySourcesInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeDeliverySourcesOutput, error)
	DescribeDestinations(context.Context, *cloudwatchlogs.DescribeDestinationsInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeDestinationsOutput, error)
	DescribeExportTasks(context.Context, *cloudwatchlogs.DescribeExportTasksInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeExportTasksOutput, error)
	DescribeLogGroups(context.Context, *cloudwatchlogs.DescribeLogGroupsInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeLogGroupsOutput, error)
	DescribeLogStreams(context.Context, *cloudwatchlogs.DescribeLogStreamsInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeLogStreamsOutput, error)
	DescribeMetricFilters(context.Context, *cloudwatchlogs.DescribeMetricFiltersInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeMetricFiltersOutput, error)
	DescribeQueries(context.Context, *cloudwatchlogs.DescribeQueriesInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeQueriesOutput, error)
	DescribeQueryDefinitions(context.Context, *cloudwatchlogs.DescribeQueryDefinitionsInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeQueryDefinitionsOutput, error)
	DescribeResourcePolicies(context.Context, *cloudwatchlogs.DescribeResourcePoliciesInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeResourcePoliciesOutput, error)
	DescribeSubscriptionFilters(context.Context, *cloudwatchlogs.DescribeSubscriptionFiltersInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeSubscriptionFiltersOutput, error)
	GetDataProtectionPolicy(context.Context, *cloudwatchlogs.GetDataProtectionPolicyInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.GetDataProtectionPolicyOutput, error)
	GetDelivery(context.Context, *cloudwatchlogs.GetDeliveryInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.GetDeliveryOutput, error)
	GetDeliveryDestination(context.Context, *cloudwatchlogs.GetDeliveryDestinationInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.GetDeliveryDestinationOutput, error)
	GetDeliveryDestinationPolicy(context.Context, *cloudwatchlogs.GetDeliveryDestinationPolicyInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.GetDeliveryDestinationPolicyOutput, error)
	GetDeliverySource(context.Context, *cloudwatchlogs.GetDeliverySourceInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.GetDeliverySourceOutput, error)
	GetLogEvents(context.Context, *cloudwatchlogs.GetLogEventsInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.GetLogEventsOutput, error)
	GetLogGroupFields(context.Context, *cloudwatchlogs.GetLogGroupFieldsInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.GetLogGroupFieldsOutput, error)
	GetLogRecord(context.Context, *cloudwatchlogs.GetLogRecordInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.GetLogRecordOutput, error)
	GetQueryResults(context.Context, *cloudwatchlogs.GetQueryResultsInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.GetQueryResultsOutput, error)
	ListTagsForResource(context.Context, *cloudwatchlogs.ListTagsForResourceInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.ListTagsForResourceOutput, error)
	ListTagsLogGroup(context.Context, *cloudwatchlogs.ListTagsLogGroupInput, ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.ListTagsLogGroupOutput, error)
}
