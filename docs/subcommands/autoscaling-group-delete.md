---
description: Delete an Autoscaling Group
---

# AutoscalingGroupDelete

## Usage

```text
ionosctl autoscaling group delete [flags]
```

## Aliases

For `autoscaling` command:

```text
[auto]
```

For `group` command:

```text
[g]
```

For `delete` command:

```text
[d]
```

## Description

Use this command to delete a specified Autoscaling Group from your account.

Required values to run command:

* Autoscaling Group Id

## Options

```text
  -u, --api-url string    Override default host url (default "https://api.ionos.com")
      --cols strings      Set of columns to be printed on output 
                          Available columns: [GroupId Name DatacenterId Location TemplateId MaxReplicaCount MinReplicaCount TargetReplicaCount Metric Range Unit ScaleInThreshold ScaleInAmount ScaleInAmountType ScaleInCoolDownPeriod ScaleOutThreshold ScaleOutAmount ScaleOutAmountType ScaleOutCoolDownPeriod State] (default [GroupId,Name,DatacenterId,Location,TemplateId,TargetReplicaCount,State])
  -c, --config string     Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
  -f, --force             Force command to execute without user input
  -i, --group-id string   The unique Group Id (required)
  -h, --help              help for delete
  -o, --output string     Desired output format [text|json] (default "text")
  -q, --quiet             Quiet output
```

## Examples

```text
ionosctl autoscaling group delete -i TEMPLATE_ID
```

