---
description: Get a VM Autoscaling Template
---

# AutoscalingTemplateGet

## Usage

```text
ionosctl autoscaling template get [flags]
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

For `get` command:

```text
[g]
```

## Description

Use this command to retrieve details about a VM Autoscaling Template by using its ID.

Required values to run command:

* VM Autoscaling Template Id

## Options

```text
  -u, --api-url string       Override default host url (default "https://api.ionos.com")
      --cols strings         Set of columns to be printed on output 
                             Available columns: [TemplateId Name Location CpuFamily AvailabilityZone Cores Ram State] (default [TemplateId,Name,Location,CpuFamily,AvailabilityZone,Ram,State])
  -c, --config string        Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
  -f, --force                Force command to execute without user input
  -h, --help                 help for get
  -o, --output string        Desired output format [text|json] (default "text")
  -q, --quiet                Quiet output
  -i, --template-id string   The unique Template Id (required)
```

## Examples

```text
ionosctl autoscaling template get -i TEMPLATE_ID
```

