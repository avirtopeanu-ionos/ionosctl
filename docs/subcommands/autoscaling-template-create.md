---
description: Create a VM Autoscaling Template
---

# AutoscalingTemplateCreate

## Usage

```text
ionosctl autoscaling template create [flags]
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

For `create` command:

```text
[c]
```

## Description

Use this command to create a VM Autoscaling Template. The VM Autoscaling Template contains information for the VMs. You can specify the name, location, availability zone, cores, cpu family for the VMs.

Regarding the Ram size, it must be specified in multiples of 256 MB with a minimum of 256 MB; however, if you set ramHotPlug to TRUE then you must use a minimum of 1024 MB. If you set the RAM size more than 240GB, then ramHotPlug will be set to FALSE and can not be set to TRUE unless RAM size not set to less than 240GB.

Right now, the VM Autoscaling Template supports only one Template Volume. Important: the volume created will NOT be deleted on SCALE IN type of Autoscaling Actions. If you want to create a Volume Template, you need to provide Image Id. If you want to see the Volume Template properties, use `ionosctl autoscaling volume-template list` command.

Also, the VM Autoscaling Template supports multiple NIC Templates. To create a VM Autoscaling Template with multiple NIC Templates use `--lan-ids "LAN_ID1,LAN_ID2"` and `--template-nics "NAME1,NAME2"` options. It is recommended to use both options. If you want to see the NIC Templates properties, use `ionosctl autoscaling nic-template list` command.

## Options

```text
  -u, --api-url string             Override default host url (default "https://api.ionos.com")
  -z, --availability-zone string   Zone where the VMs created using this VM Autoscaling Template (default "AUTO")
      --cols strings               Set of columns to be printed on output 
                                   Available columns: [TemplateId Name Location CpuFamily AvailabilityZone Cores Ram State] (default [TemplateId,Name,Location,CpuFamily,AvailabilityZone,Ram,State])
  -c, --config string              Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
      --cores int                  The total number of cores for the VMs. Minimum: 1 (default 1)
      --cpu-family string          CPU family for the VMs created using the VM Autoscaling Template. If null, the VM will be created with the default CPU family from the assigned location
  -f, --force                      Force command to execute without user input
  -h, --help                       help for create
      --image-id string            Image installed on the Volume. Only Id of the Image is supported currently. Required flag when creating a Volume Template
      --lan-ids ints               Lan Ids for the NIC Templates. Minimum value for Lan Id: 1 (default [1])
  -l, --location string            Location for the VM Autoscaling Template (default "de/txl")
  -n, --name string                Name of the VM Autoscaling Template (default "Unnamed VM Autoscaling Template")
  -o, --output string              Desired output format [text|json] (default "text")
  -p, --password string            Image password for the Volume Template (default "abcde1234")
  -q, --quiet                      Quiet output
      --ram string                 The amount of memory for the VMs. Size must be specified in multiples of 256. e.g. --ram 2048 or --ram 2048MB (default "2048")
      --size string                User-defined size for this template volume in GB. e.g.: --size 10 or --size 10GB. (default "10")
      --ssh-keys strings           SSH Keys that have access to the Volume
      --template-nics strings      Names for the NIC Templates (default [Unnamed VM Autoscaling NIC Template])
      --template-volume string     Name of the Volume Template (default "Unnamed VM Autoscaling Template Volume")
      --type string                Type of the Volume (default "HDD")
      --user-data string           User-Data (Cloud Init) for the Volume Template
```

## Examples

```text
ionosctl autoscaling template create

ionosctl autoscaling template create --image-id IMAGE_ID
```

