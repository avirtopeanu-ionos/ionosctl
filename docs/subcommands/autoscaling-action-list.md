---
description: List Actions from a VM Autoscaling Group
---

# AutoscalingActionList

## Usage

```text
ionosctl autoscaling action list [flags]
```

## Aliases

For `autoscaling` command:

```text
[auto]
```

For `action` command:

```text
[a]
```

For `list` command:

```text
[l ls]
```

## Description

Use this command to retrieve a complete list of Actions from a VM Autoscaling Group provisioned under your account.

Use flags to retrieve a list of Actions:

* sorting by type, using `ionosctl autoscaling action list --group-id GROUP_ID --type ACTION_TYPE`
* sorting by status, using `ionosctl autoscaling action list --group-id GROUP_ID --status ACTION_STATUS`

Required values to run command:

* VM Autoscaling Group Id

## Options

```text
  -u, --api-url string    Override default host url (default "https://api.ionos.com")
      --cols strings      Set of columns to be printed on output 
                          Available columns: [ActionId ActionStatus ActionType TargetReplicaCount] (default [ActionId,ActionStatus,ActionType,TargetReplicaCount])
  -c, --config string     Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
  -f, --force             Force command to execute without user input
      --group-id string   The unique Group Id (required)
  -h, --help              help for list
  -o, --output string     Desired output format [text|json] (default "text")
  -q, --quiet             Quiet output
  -s, --status string     Sort Actions based on VM Autoscaling Action Status
  -t, --type string       Sort Actions based on VM Autoscaling Action Type
  -v, --verbose           see step by step process when running a command
```

## Examples

```text
ionosctl autoscaling action list --group-id GROUP_ID

ionosctl autoscaling action list --group-id GROUP_ID --status ACTION_STATUS --type ACTION_TYPE
```

