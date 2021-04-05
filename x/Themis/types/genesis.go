package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		VoteList:    []*Vote{},
		PollList:    []*Poll{},
		GroupList:   []*Group{},
		AccountList: []*Account{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated ID in vote
	voteIdMap := make(map[uint64]bool)

	for _, elem := range gs.VoteList {
		if _, ok := voteIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for vote")
		}
		voteIdMap[elem.Id] = true
	}
	// Check for duplicated ID in poll
	pollIdMap := make(map[uint64]bool)

	for _, elem := range gs.PollList {
		if _, ok := pollIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for poll")
		}
		pollIdMap[elem.Id] = true
	}
	// Check for duplicated ID in group
	groupIdMap := make(map[uint64]bool)

	for _, elem := range gs.GroupList {
		if _, ok := groupIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for group")
		}
		groupIdMap[elem.Id] = true
	}
	// Check for duplicated ID in account
	accountIdMap := make(map[uint64]bool)

	for _, elem := range gs.AccountList {
		if _, ok := accountIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for account")
		}
		accountIdMap[elem.Id] = true
	}

	return nil
}
