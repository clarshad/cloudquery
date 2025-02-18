# Table: aws_iam_user_attached_policies

This table shows data for IAM User Attached Policies.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_AttachedPolicy.html

The composite primary key for this table is (**account_id**, **user_arn**, **policy_name**).

## Relations

This table depends on [aws_iam_users](aws_iam_users.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|user_arn (PK)|`utf8`|
|policy_name (PK)|`utf8`|
|user_id|`utf8`|
|policy_arn|`utf8`|