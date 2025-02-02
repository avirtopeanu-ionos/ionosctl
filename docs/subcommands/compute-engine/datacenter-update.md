---
description: Update a Data Center
---

# DatacenterUpdate

## Usage

```text
ionosctl datacenter update [flags]
```

## Aliases

For `datacenter` command:

```text
[d dc vdc]
```

For `update` command:

```text
[u up]
```

## Description

Use this command to change a Virtual Data Center's name, description.

You can wait for the Request to be executed using `--wait-for-request` option.

Required values to run command:

* Data Center Id

## Options

```text
  -u, --api-url string         Override default host url (default "https://api.ionos.com")
      --cols strings           Set of columns to be printed on output 
                               Available columns: [DatacenterId Name Location State Description Version Features CpuFamily SecAuthProtection] (default [DatacenterId,Name,Location,CpuFamily,State])
  -c, --config string          Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
  -i, --datacenter-id string   The unique Data Center Id (required)
  -D, --depth int32            Controls the detail depth of the response objects. Max depth is 10.
  -d, --description string     Description of the Data Center
  -f, --force                  Force command to execute without user input
  -h, --help                   Print usage
  -n, --name string            Name of the Data Center
  -o, --output string          Desired output format [text|json] (default "text")
  -q, --quiet                  Quiet output
  -t, --timeout int            Timeout option for Request for Data Center update [seconds] (default 60)
  -v, --verbose                Print step-by-step process when running command
  -w, --wait-for-request       Wait for the Request for Data Center update to be executed
```

## Examples

```text
ionosctl datacenter update --datacenter-id DATACENTER_ID --description DESCRIPTION --cols "DatacenterId,Description"
```

