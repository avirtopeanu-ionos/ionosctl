---
description: Create a Kubernetes Cluster
---

# K8sClusterCreate

## Usage

```text
ionosctl k8s cluster create [flags]
```

## Aliases

For `cluster` command:

```text
[c]
```

For `create` command:

```text
[c]
```

## Description

Use this command to create a new Managed Kubernetes Cluster. Regarding the name for the Kubernetes Cluster, the limit is 63 characters following the rule to begin and end with an alphanumeric character with dashes, underscores, dots, and alphanumerics between. Regarding the Kubernetes Version for the Cluster, if not set via flag, it will be used the default one: `ionosctl k8s version get`.

You can wait for the Cluster to be in "ACTIVE" state using `--wait-for-state` flag together with `--timeout` option.

## Options

```text
      --api-subnets strings   Access to the K8s API server is restricted to these CIDRs. Cluster-internal traffic is not affected by this restriction. If no allowlist is specified, access is not restricted. If an IP without subnet mask is provided, the default value will be used: 32 for IPv4 and 128 for IPv6
  -u, --api-url string        Override default host url (default "https://api.ionos.com")
      --cols strings          Set of columns to be printed on output 
                              Available columns: [ClusterId Name K8sVersion State MaintenanceWindow AvailableUpgradeVersions ViableNodePoolVersions S3Bucket ApiSubnetAllowList] (default [ClusterId,Name,K8sVersion,State,MaintenanceWindow])
  -c, --config string         Configuration file used for authentication (default "$XDG_CONFIG_HOME/ionosctl/config.json")
  -D, --depth int32           Controls the detail depth of the response objects. Max depth is 10.
  -f, --force                 Force command to execute without user input
  -h, --help                  Print usage
      --k8s-version string    The K8s version for the Cluster. If not set, the default one will be used
  -n, --name string           The name for the K8s Cluster (default "UnnamedCluster")
  -o, --output string         Desired output format [text|json] (default "text")
  -q, --quiet                 Quiet output
      --s3bucket string       S3 Bucket name configured for K8s usage
  -t, --timeout int           Timeout option for waiting for Cluster/Request [seconds] (default 600)
  -v, --verbose               Print step-by-step process when running command
  -w, --wait-for-request      Wait for the Request for Cluster creation to be executed
  -W, --wait-for-state        Wait for the new Cluster to be in ACTIVE state
```

## Examples

```text
ionosctl k8s cluster create --name NAME
```

