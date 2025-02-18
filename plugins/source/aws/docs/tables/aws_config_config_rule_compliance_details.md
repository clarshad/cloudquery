# Table: aws_config_config_rule_compliance_details

This table shows data for Config Config Rule Compliance Details.

https://docs.aws.amazon.com/config/latest/APIReference/API_EvaluationResult.html

The composite primary key for this table is (**config_rule_arn**, **resource_evaluation_id**).

## Relations

This table depends on [aws_config_config_rules](aws_config_config_rules.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|config_rule_arn (PK)|`utf8`|
|resource_evaluation_id (PK)|`utf8`|
|config_rule_name|`utf8`|
|annotation|`utf8`|
|compliance_type|`utf8`|
|config_rule_invoked_time|`timestamp[us, tz=UTC]`|
|evaluation_result_identifier|`json`|
|result_recorded_time|`timestamp[us, tz=UTC]`|
|result_token|`utf8`|