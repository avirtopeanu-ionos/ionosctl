---
description: List Autoscaling Templates
---

# AutoscalingTemplateList

## Usage

```text
ionosctl autoscaling template list [flags]
```

## Aliases

For `autoscaling` command:

```text
[auto]
```

For `template` command:

```text
[t]
```

For `list` command:

```text
[l ls]
```

## Description

Use this command to retrieve a complete list of Autoscaling Templates provisioned under your account.

## Options

```text
  -u, --api-url string   Override default host url (default "https://api.ionos.com")
      --cols strings     Set of columns to be printed on output 
                         Available columns: [TemplateId Name Location CpuFamily AvailabilityZone Cores Ram State] (default [TemplateId,Name,Location,CpuFamily,AvailabilityZone,Ram,State])
  -c, --config string    Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
  -f, --force            Force command to execute without user input
  -h, --help             help for list
  -o, --output string    Desired output format [text|json] (default "text")
  -q, --quiet            Quiet output
```

## Examples

```text
ionosctl autoscaling template list
```

