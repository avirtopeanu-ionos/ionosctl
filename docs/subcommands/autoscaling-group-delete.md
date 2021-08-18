---
description: Delete a VM Autoscaling Group
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

Use this command to delete a specified VM Autoscaling Group from your account.

Required values to run command:

* VM Autoscaling Group Id

## Options

```text
  -u, --api-url string    Override default host url (default "https://api.ionos.com")
  -c, --config string     Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
  -f, --force             Force command to execute without user input
  -i, --group-id string   The unique Group Id (required)
  -h, --help              help for delete
  -o, --output string     Desired output format [text|json] (default "text")
  -q, --quiet             Quiet output
  -v, --verbose           see step by step process when running a command
```

## Examples

```text
ionosctl autoscaling group delete -i GROUP_ID
```

