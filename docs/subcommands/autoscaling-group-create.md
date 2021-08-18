---
description: Create a VM Autoscaling Group
---

# AutoscalingGroupCreate

## Usage

```text
ionosctl autoscaling group create [flags]
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

For `create` command:

```text
[c]
```

## Description

Create a VM Autoscaling Group. 

Regarding some of the VM Autoscaling Group Properties, please see more details:

* [Group][DatacenterId]The Datacenter Id property represents VMs for this VM Autoscaling Group will be created in this Virtual Datacenter. Please note, that it have the same location as the template.
* [Group][TargetReplicaCount] Depending on the scaling policy, the target number of VMs will be adjusted automatically. VMs will be created or destroyed automatically in order to adjust the actual number of VMs to this number. This value can be set only at Group creation time, subsequent change via update (PUT) request is not possible
* [Group Policy][ScaleInThreshold] Scale In Threshold is a lower threshold on the value of `metric`. Will be used with `less than` (<) operator. Exceeding this will start a Scale-In Action as specified by the `scaleInAction` property. The value must have a higher minimum delta to the `scaleOutThreshold` depending on the `metric` to avoid competitive actions at the same time
* [Group Policy][ScaleOutThreshold] An upper threshold on the value of `metric`.  Will be used with `greater than` (>) operator. Exceeding this will start a Scale-Out Action as specified by the `scaleOutAction` property. The value must have a lower minimum delta to the `scaleInThreshold` depending on the `metric` to avoid competitive actions at the same time
* [Group Policy Action][Amount] When `amountType == ABSOLUTE`, amount parameter is the number of VMs added or removed in one step. When `amountType == PERCENTAGE`, amount parameter is a percentage value, which will be applied to the Autoscaling Group's current `targetReplicaCount` in order to derive the number of VMs that will be added or removed in one step. There will always be at least one VM added or removed
* [Group Policy Action][CoolDownPeriod] Minimum time to pass after this Scaling Action has started, until the next Scaling Action will be started. Additionally, if a Scaling Action is currently in progress, no second Scaling Action will be started for the same Autoscaling Group. Instead, the Metric will be re-evaluated after the current Scaling Action completed (either successful or with failures)

Required values to run command:

* VM Autoscaling Template Id
* Datacenter Id

## Options

```text
  -u, --api-url string                         Override default host url (default "https://api.ionos.com")
      --cols strings                           Set of columns to be printed on output 
                                               Available columns: [GroupId Name DatacenterId Location TemplateId MaxReplicaCount MinReplicaCount TargetReplicaCount Metric Range Unit ScaleInThreshold ScaleInAmount ScaleInAmountType ScaleInCoolDownPeriod ScaleOutThreshold ScaleOutAmount ScaleOutAmountType ScaleOutCoolDownPeriod State] (default [GroupId,Name,DatacenterId,Location,TemplateId,TargetReplicaCount,State])
  -c, --config string                          Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
      --datacenter-id string                   The unique Data Center Id (required)
  -f, --force                                  Force command to execute without user input
  -h, --help                                   help for create
      --max-replica-count targetReplicaCount   Maximum replica count value for targetReplicaCount. Will be enforced for both automatic and manual changes. Mininum: 0; Maximum: 200 (default 5)
  -m, --metric string                          [Group Policy] The Metric that should trigger Scaling Actions. The values of the Metric are checked in fixed intervals (default "INSTANCE_CPU_UTILIZATION_AVERAGE")
      --min-replica-count targetReplicaCount   Minimum replica count value for targetReplicaCount. Will be enforced for both automatic and manual changes. Mininum: 0; Maximum: 200 (default 1)
  -n, --name string                            User-defined name for the VM Autoscaling Group (default "Unnamed VM Autoscaling Group")
  -o, --output string                          Desired output format [text|json] (default "text")
  -q, --quiet                                  Quiet output
  -r, --range string                           [Group Policy] Defines the range of time from which samples will be aggregated (default "PT24H")
      --scale-in-amount float32                [Group Policy][Scale In Action] Amount of VMs (in percentage or absolute value) to be removed in a Scale In Action (default 1)
      --scale-in-amount-type string            [Group Policy][Scale In Action] The type for the given amount (default "ABSOLUTE")
      --scale-in-cooldown string               [Group Policy][Scale In Action] Cool Down Period (default "5m")
      --scale-in-threshold metric              [Group Policy][Scale In Action] A lower threshold on the value of metric (default 33)
      --scale-out-amount float32               [Group Policy][Scale Out Action] Amount of VMs (in percentage or absolute value) to be added in a Scale Out Action (default 1)
      --scale-out-amount-type string           [Group Policy][Scale Out Action] The type for the given amount (default "ABSOLUTE")
      --scale-out-cooldown string              [Group Policy][Scale Out Action] Cool Down Period (default "5m")
      --scale-out-threshold metric             [Group Policy][Scale Out Action] An upper threshold on the value of metric (default 77)
  -t, --target-replica-count int               The target number of VMs in this Group. Minimum: 0; Maximum: 200 (default 1)
  -i, --template-id string                     The unique Template Id (required)
      --unit string                            [Group Policy] Unit of the applied Metric (default "PER_HOUR")
  -v, --verbose                                see step by step process when running a command
```

## Examples

```text
ionosctl autoscaling group create --datacenter-id DATACENTER_ID --template-id TEMPLATE_ID
```

