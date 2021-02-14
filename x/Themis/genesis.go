package Themis

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/keeper"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the poll
	for _, elem := range genState.PollList {
		k.SetPoll(ctx, *elem)
	}

	// Set poll count
	k.SetPollCount(ctx, int64(len(genState.PollList)))

	// Set all the account
	for _, elem := range genState.AccountList {
		k.SetAccount(ctx, *elem)
	}

	// Set account count
	k.SetAccountCount(ctx, int64(len(genState.AccountList)))

	// Set all the group
	for _, elem := range genState.GroupList {
		k.SetGroup(ctx, *elem)
	}

	// Set group count
	k.SetGroupCount(ctx, int64(len(genState.GroupList)))

}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all poll
	pollList := k.GetAllPoll(ctx)
	for _, elem := range pollList {
		elem := elem
		genesis.PollList = append(genesis.PollList, &elem)
	}

	// Get all account
	accountList := k.GetAllAccount(ctx)
	for _, elem := range accountList {
		elem := elem
		genesis.AccountList = append(genesis.AccountList, &elem)
	}

	// Get all group
	groupList := k.GetAllGroup(ctx)
	for _, elem := range groupList {
		elem := elem
		genesis.GroupList = append(genesis.GroupList, &elem)
	}

	return genesis
}
