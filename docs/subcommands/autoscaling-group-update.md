---
description: Update a VM Autoscaling Group
---

# AutoscalingGroupUpdate

## Usage

```text
ionosctl autoscaling group update [flags]
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

For `update` command:

```text
[up]
```

## Description

Update a VM Autoscaling Group. 

Required values to run command:

* VM Autoscaling Group Id

## Options

```text
  -u, --api-url string                         Override default host url (default "https://api.ionos.com")
      --cols strings                           Set of columns to be printed on output 
                                               Available columns: [GroupId Name DatacenterId Location TemplateId MaxReplicaCount MinReplicaCount TargetReplicaCount Metric Range Unit ScaleInThreshold ScaleInAmount ScaleInAmountType ScaleInCoolDownPeriod ScaleOutThreshold ScaleOutAmount ScaleOutAmountType ScaleOutCoolDownPeriod State] (default [GroupId,Name,DatacenterId,Location,TemplateId,TargetReplicaCount,State])
  -c, --config string                          Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
  -f, --force                                  Force command to execute without user input
  -i, --group-id string                        The unique Group Id (required)
  -h, --help                                   help for update
      --max-replica-count targetReplicaCount   Maximum replica count value for targetReplicaCount. Will be enforced for both automatic and manual changes. Mininum: 0; Maximum: 200
  -m, --metric string                          [Group Policy] The Metric that should trigger Scaling Actions. The values of the Metric are checked in fixed intervals
      --min-replica-count targetReplicaCount   Minimum replica count value for targetReplicaCount. Will be enforced for both automatic and manual changes. Mininum: 0; Maximum: 200
  -n, --name string                            User-defined name for the VM Autoscaling Group
  -o, --output string                          Desired output format [text|json] (default "text")
  -q, --quiet                                  Quiet output
  -r, --range string                           [Group Policy] Defines the range of time from which samples will be aggregated
      --scale-in-amount float32                [Group Policy][Scale In Action] Amount of VMs (in percentage or absolute value) to be removed in a Scale In Action
      --scale-in-amount-type string            [Group Policy][Scale In Action] The type for the given amount
      --scale-in-cooldown string               [Group Policy][Scale In Action] Cool Down Period
      --scale-in-threshold metric              [Group Policy][Scale In Action] A lower threshold on the value of metric
      --scale-out-amount float32               [Group Policy][Scale Out Action] Amount of VMs (in percentage or absolute value) to be added in a Scale Out Action
      --scale-out-amount-type string           [Group Policy][Scale Out Action] The type for the given amount
      --scale-out-cooldown string              [Group Policy][Scale Out Action] Cool Down Period
      --scale-out-threshold metric             [Group Policy][Scale Out Action] An upper threshold on the value of metric (default 40)
      --template-id string                     The unique Template Id
      --unit string                            [Group Policy] Unit of the applied Metric
```

## Examples

```text
ionosctl autoscaling group update -i GROUP_ID --name GROUP_NAME
```

