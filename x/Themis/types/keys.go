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
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	GroupKey      = "Group-value-"
	GroupCountKey = "Group-count-"
)

const (
	AccountKey      = "Account-value-"
	AccountPtrKey   = "Account-ptr-value-"
	AccountCountKey = "Account-count-"
)

const (
	PollKey      = "Poll-value-"
	PollCountKey = "Poll-count-"
)

const (
	VoteKey      = "Vote-value-"
	VotePtrKey   = "Vote-ptr-value-"
	VoteCountKey = "Vote-count-"
)
