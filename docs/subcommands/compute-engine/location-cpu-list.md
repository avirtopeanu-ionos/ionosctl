---
description: List available CPU Architecture from a Location
---

# LocationCpuList

## Usage

```text
ionosctl location cpu list [flags]
```

## Aliases

For `location` command:

```text
[loc]
```

For `list` command:

```text
[l ls]
```

## Description

Use this command to get information about available CPU Architectures from a specific Location.

Required values to run command:

* Location Id

## Options

```text
  -u, --api-url string       Override default host url (default "https://api.ionos.com")
      --cols strings         Set of columns to be printed on output 
                             Available columns: [CpuFamily MaxCores MaxRam Vendor] (default [CpuFamily,MaxCores,MaxRam,Vendor])
  -c, --config string        Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
  -D, --depth int32          Controls the detail depth of the response objects. Max depth is 10. (default 1)
  -f, --force                Force command to execute without user input
  -h, --help                 Print usage
      --location-id string   The unique Location Id (required)
      --no-headers           When using text output, don't print headers
  -o, --output string        Desired output format [text|json] (default "text")
  -q, --quiet                Quiet output
  -v, --verbose              Print step-by-step process when running command
```

## Examples

```text
ionosctl location cpu list --location-id LOCATION_ID
```

