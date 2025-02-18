# Table: aws_securityhub_hubs

This table shows data for AWS Security Hub Hubs.

https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_DescribeHub.html

The composite primary key for this table is (**account_id**, **region**, **hub_arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|tags|`json`|
|auto_enable_controls|`bool`|
|control_finding_generator|`utf8`|
|hub_arn (PK)|`utf8`|
|subscribed_at|`timestamp[us, tz=UTC]`|