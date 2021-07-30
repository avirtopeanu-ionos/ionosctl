---
description: List Autoscaling Groups
---

# AutoscalingGroupList

## Usage

```text
ionosctl autoscaling group list [flags]
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

For `list` command:

```text
[l ls]
```

## Description

Use this command to retrieve a complete list of Autoscaling Groups provisioned under your account.

## Options

```text
  -u, --api-url string   Override default host url (default "https://api.ionos.com")
      --cols strings     Set of columns to be printed on output 
                         Available columns: [GroupId Name DatacenterId Location TemplateId MaxReplicaCount MinReplicaCount TargetReplicaCount Metric Range Unit ScaleInThreshold ScaleInAmount ScaleInAmountType ScaleInCoolDownPeriod ScaleOutThreshold ScaleOutAmount ScaleOutAmountType ScaleOutCoolDownPeriod State] (default [GroupId,Name,DatacenterId,Location,TemplateId,TargetReplicaCount,State])
  -c, --config string    Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
  -f, --force            Force command to execute without user input
  -h, --help             help for list
  -o, --output string    Desired output format [text|json] (default "text")
  -q, --quiet            Quiet output
```

## Examples

```text
ionosctl autoscaling group list
```

