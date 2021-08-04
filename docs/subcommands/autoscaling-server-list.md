---
description: List Servers from a VM Autoscaling Group
---

# AutoscalingServerList

## Usage

```text
ionosctl autoscaling server list [flags]
```

## Aliases

For `autoscaling` command:

```text
[auto]
```

For `server` command:

```text
[svr]
```

For `list` command:

```text
[l ls]
```

## Description

Use this command to retrieve a complete list of Servers from a specified VM Autoscaling Group provisioned under your account.

Required values to run command:

* VM Autoscaling Group Id

## Options

```text
  -u, --api-url string    Override default host url (default "https://api.ionos.com")
      --cols strings      Set of columns to be printed on output 
                          Available columns: [ServerId DatacenterId Name] (default [ServerId,DatacenterId,Name])
  -c, --config string     Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
  -f, --force             Force command to execute without user input
      --group-id string   The unique Group Id (required)
  -h, --help              help for list
  -o, --output string     Desired output format [text|json] (default "text")
  -q, --quiet             Quiet output
```

## Examples

```text
ionosctl autoscaling server list --group-id GROUP_ID
```

