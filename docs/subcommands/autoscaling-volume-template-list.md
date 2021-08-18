---
description: List Volume Templates from a VM Autoscaling Template
---

# AutoscalingVolumeTemplateList

## Usage

```text
ionosctl autoscaling volume-template list [flags]
```

## Aliases

For `autoscaling` command:

```text
[auto]
```

For `volume-template` command:

```text
[v]
```

For `list` command:

```text
[l ls]
```

## Description

Use this command to retrieve a complete list of Volume Templates from a specific VM Autoscaling Template provisioned under your account.

Required values to run command:

* VM Autoscaling Template Id

## Options

```text
  -u, --api-url string       Override default host url (default "https://api.ionos.com")
      --cols strings         Set of columns to be printed on output 
                             Available columns: [Name Size Type Image UserData] (default [Name,Size,Type,Image,UserData])
  -c, --config string        Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
  -f, --force                Force command to execute without user input
  -h, --help                 help for list
  -o, --output string        Desired output format [text|json] (default "text")
  -q, --quiet                Quiet output
  -i, --template-id string   The unique Template Id (required)
  -v, --verbose              see step by step process when running a command
```

## Examples

```text
ionosctl autoscaling volume-template list --template-id TEMPLATE_ID
```

