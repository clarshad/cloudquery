# Table: aws_waf_web_acls

This table shows data for WAF Web ACLs.

https://docs.aws.amazon.com/waf/latest/APIReference/API_waf_WebACLSummary.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|default_action|`json`|
|rules|`json`|
|web_acl_id|`utf8`|
|metric_name|`utf8`|
|name|`utf8`|
|web_acl_arn|`utf8`|
|logging_configuration|`json`|