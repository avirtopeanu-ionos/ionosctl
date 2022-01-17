package dbaas_postgres

const (
	ArgIdShort              = "i"
	ArgClusterId            = "cluster-id"
	ArgStartTime            = "start-time"
	ArgStartTimeShort       = "s"
	ArgEndTime              = "end-time"
	ArgEndTimeShort         = "e"
	ArgLimit                = "limit"
	ArgLimitShort           = "l"
	ArgVersion              = "version"
	ArgVersionShort         = "V"
	ArgInstances            = "instances"
	ArgInstancesShort       = "I"
	ArgSyncMode             = "sync"
	ArgSyncModeShort        = "S"
	ArgCores                = "cores"
	ArgRam                  = "ram"
	ArgStorageSize          = "storage-size"
	ArgStorageType          = "storage-type"
	ArgDatacenterId         = "datacenter-id"
	ArgDatacenterIdShort    = "D"
	ArgBackupId             = "backup-id"
	ArgBackupIdShort        = "b"
	ArgRecoveryTime         = "recovery-time"
	ArgRecoveryTimeShort    = "R"
	ArgCidr                 = "cidr"
	ArgCidrShort            = "C"
	ArgLanId                = "lan-id"
	ArgLanIdShort           = "L"
	ArgLocation             = "location-id"
	ArgName                 = "name"
	ArgNameShort            = "n"
	ArgDbUsername           = "db-username"
	ArgDbUsernameShort      = "U"
	ArgDbPassword           = "db-password"
	ArgDbPasswordShort      = "P"
	ArgMaintenanceTime      = "maintenance-time"
	ArgMaintenanceTimeShort = "T"
	ArgMaintenanceDay       = "maintenance-day"
	ArgMaintenanceDayShort  = "d"
	ArgRemoveConnection     = "remove-connection"
)

const (
	ClusterId             = "The unique ID of the Cluster"
	BackupId              = "The unique ID of the Backup"
	DefaultClusterTimeout = int(1200)
)
