---
description: Get a VM Autoscaling Group
---

# AutoscalingGroupGet

## Usage

```text
ionosctl autoscaling group get [flags]
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

For `get` command:

```text
[g]
```

## Description

Use this command to retrieve details about a VM Autoscaling Group by using its ID.

Required values to run command:

* VM Autoscaling Group Id

## Options

```text
  -u, --api-url string    Override default host url (default "https://api.ionos.com")
      --cols strings      Set of columns to be printed on output 
                          Available columns: [GroupId Name DatacenterId Location TemplateId MaxReplicaCount MinReplicaCount TargetReplicaCount Metric Range Unit ScaleInThreshold ScaleInAmount ScaleInAmountType ScaleInCoolDownPeriod ScaleOutThreshold ScaleOutAmount ScaleOutAmountType ScaleOutCoolDownPeriod State] (default [GroupId,Name,DatacenterId,Location,TemplateId,TargetReplicaCount,State])
  -c, --config string     Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
  -f, --force             Force command to execute without user input
  -i, --group-id string   The unique Group Id (required)
  -h, --help              help for get
  -o, --output string     Desired output format [text|json] (default "text")
  -q, --quiet             Quiet output
  -v, --verbose           see step by step process when running a command
```

## Examples

```text
ionosctl autoscaling group get -i GROUP_ID
```

