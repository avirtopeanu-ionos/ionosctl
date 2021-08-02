---
description: List Actions from an Autoscaling Group
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

Use this command to retrieve a complete list of Actions from an Autoscaling Group provisioned under your account.

Required values to run command:

* Autoscaling Group Id

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
```

## Examples

```text
ionosctl autoscaling action list --group-id GROUP_ID
```

