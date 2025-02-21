# Table: aws_elasticbeanstalk_environments

This table shows data for AWS Elastic Beanstalk Environments.

https://docs.aws.amazon.com/elasticbeanstalk/latest/APIReference/API_EnvironmentDescription.html

The composite primary key for this table is (**account_id**, **id**).

## Relations

The following tables depend on aws_elasticbeanstalk_environments:
  - [aws_elasticbeanstalk_configuration_options](aws_elasticbeanstalk_configuration_options.md)
  - [aws_elasticbeanstalk_configuration_settings](aws_elasticbeanstalk_configuration_settings.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|arn|`utf8`|
|region|`utf8`|
|tags|`json`|
|id (PK)|`utf8`|
|listeners|`json`|
|abortable_operation_in_progress|`bool`|
|application_name|`utf8`|
|cname|`utf8`|
|date_created|`timestamp[us, tz=UTC]`|
|date_updated|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|endpoint_url|`utf8`|
|environment_arn|`utf8`|
|environment_id|`utf8`|
|environment_links|`json`|
|environment_name|`utf8`|
|health|`utf8`|
|health_status|`utf8`|
|operations_role|`utf8`|
|platform_arn|`utf8`|
|resources|`json`|
|solution_stack_name|`utf8`|
|status|`utf8`|
|template_name|`utf8`|
|tier|`json`|
|version_label|`utf8`|