package types

const (
	// ModuleName defines the module name
	ModuleName = "Themis"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"

	// Special State variables
	MaxNameSize    = 20
	MaxDescSize    = 200
	MaxOptions     = 20
	MaxOptionsSize = 40
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// GetStringBytes Turns strings to Bytes
func GetStringBytes(str string) []byte {
	bz := KeyPrefix(str)
	return bz
}

// GetBytesString Turns bytes into a string
func GetBytesString(bz []byte) string {
	return string(bz)
}

const (
	AccountKey      = "Account-value-"
	AccountNameKey  = "Account-name-"
	AccountAddrKey  = "Account-addr-"
	AccountCountKey = "Account-count-"
)

const (
	GroupKey      = "Group-value-"
	GroupNameKey  = "Group-name-"
	GroupAddrKey  = "Group-addr-"
	GroupCountKey = "Group-count-"
)

const (
	PollKey      = "Poll-value-"
	PollNameKey  = "Poll-name-"
	PollGroupKey = "Poll-group-"
	PollCountKey = "Poll-count-"
)

const (
	VoteKey       = "Vote-value-"
	VoteGroupKey  = "Vote-group-"
	VoteUserKey   = "Vote-user-"
	VoteOptionKey = "Vote-option-"
	VoteCountKey  = "Vote-count-"
)
