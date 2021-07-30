---
description: Delete an Autoscaling Template
---

# AutoscalingTemplateDelete

## Usage

```text
ionosctl autoscaling template delete [flags]
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

For `delete` command:

```text
[d]
```

## Description

Use this command to delete a specified Autoscaling Template from your account.

Required values to run command:

* Autoscaling Template Id

## Options

```text
  -u, --api-url string       Override default host url (default "https://api.ionos.com")
      --cols strings         Set of columns to be printed on output 
                             Available columns: [TemplateId Name Location CpuFamily AvailabilityZone Cores Ram State] (default [TemplateId,Name,Location,CpuFamily,AvailabilityZone,Ram,State])
  -c, --config string        Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
  -f, --force                Force command to execute without user input
  -h, --help                 help for delete
  -o, --output string        Desired output format [text|json] (default "text")
  -q, --quiet                Quiet output
  -i, --template-id string   The unique Template Id (required)
```

## Examples

```text
ionosctl autoscaling template delete -i TEMPLATE_ID
```

