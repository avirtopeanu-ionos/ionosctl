---
description: Get a Server from an Autoscaling Group
---

# AutoscalingServerGet

## Usage

```text
ionosctl autoscaling server get [flags]
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

For `get` command:

```text
[g]
```

## Description

Use this command to retrieve details about an Server from a specific Autoscaling Group by using its ID.

Required values to run command:

* Autoscaling Group Id

* Server Id

## Options

```text
  -u, --api-url string     Override default host url (default "https://api.ionos.com")
      --cols strings       Set of columns to be printed on output 
                           Available columns: [ServerId DatacenterId Name] (default [ServerId,DatacenterId,Name])
  -c, --config string      Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
  -f, --force              Force command to execute without user input
      --group-id string    The unique Group Id (required)
  -h, --help               help for get
  -o, --output string      Desired output format [text|json] (default "text")
  -q, --quiet              Quiet output
  -i, --server-id string   The unique Server Id (required)
```

## Examples

```text
ionosctl autoscaling server get --group-id GROUP_ID --server-id SERVER_ID
```

