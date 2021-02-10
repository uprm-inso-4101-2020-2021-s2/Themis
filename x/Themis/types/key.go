package types

const (
	// ModuleName is the name of the module
	ModuleName = "Themis"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querier msgs
	QuerierRoute = ModuleName
)

const (
	GroupPrefix      = "group-value-"
	GroupCountPrefix = "group-count-"
)

const (
	AccountPrefix      = "account-value-"
	AccountPtrPrefix   = "account-prt-value"
	AccountCountPrefix = "account-account-count-"
)
