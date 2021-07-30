---
description: Get an Action from an Autoscaling Group
---

# AutoscalingActionGet

## Usage

```text
ionosctl autoscaling action get [flags]
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

For `get` command:

```text
[g]
```

## Description

Use this command to retrieve details about an Action from an Autoscaling Group by using its ID.

Required values to run command:

* Autoscaling Group Id
* Action Id

## Options

```text
  -i, --action-id string   The unique Action Id (required)
  -u, --api-url string     Override default host url (default "https://api.ionos.com")
      --cols strings       Set of columns to be printed on output 
                           Available columns: [ActionId ActionStatus ActionType TargetReplicaCount] (default [ActionId,ActionStatus,ActionType,TargetReplicaCount])
  -c, --config string      Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
  -f, --force              Force command to execute without user input
      --group-id string    The unique Group Id (required)
  -h, --help               help for get
  -o, --output string      Desired output format [text|json] (default "text")
  -q, --quiet              Quiet output
```

