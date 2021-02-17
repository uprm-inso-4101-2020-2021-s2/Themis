package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
	"strconv"
)

// GetVoteCount get the total number of vote
func (k Keeper) GetVoteCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteCountKey))
	byteKey := types.KeyPrefix(types.VoteCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetVoteCount set the total number of vote
func (k Keeper) SetVoteCount(ctx sdk.Context, count int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteCountKey))
	byteKey := types.KeyPrefix(types.VoteCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateVote creates a vote with a new id and update the count
func (k Keeper) CreateVote(ctx sdk.Context, msg types.MsgCreateVote) {
	// Create the vote
	count := k.GetVoteCount(ctx)

	// Get group from poll
	// Get account from creator and group
	group := k.GetPollGroup(ctx, msg.Poll)
	acc := k.GetAccountId(ctx, msg.Creator, group)

	// Vote id are composed of pollID-accountID and the Ptrs are the opposite
	var vote = types.Vote{
		Id:      k.NewKey(msg.Poll, k.NewKey(acc, strconv.FormatInt(count, 10))),
		Account: acc,
		Poll:    msg.Poll,
		Option:  msg.Option,
	}

	var votePtr = types.VotePtr{
		Id: vote.Id,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteKey))
	storePtr := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VotePtrKey))
	key := types.KeyPrefix(types.VoteKey + vote.Id)
	keyPtr := types.KeyPrefix(types.VotePtrKey + k.NewKey(acc, k.NewKey(msg.Poll, strconv.FormatInt(count, 10))))
	value := k.cdc.MustMarshalBinaryBare(&vote)
	valuePtr := k.cdc.MustMarshalBinaryBare(&votePtr)
	store.Set(key, value)
	storePtr.Set(keyPtr, valuePtr)

	// Update vote count
	k.SetVoteCount(ctx, count+1)
}

// SetVote set a specific vote in the store
func (k Keeper) SetVote(ctx sdk.Context, vote types.Vote) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteKey))
	b := k.cdc.MustMarshalBinaryBare(&vote)
	store.Set(types.KeyPrefix(types.VoteKey+vote.Id), b)
}

// GetVote returns a vote from its id
func (k Keeper) GetVote(ctx sdk.Context, key string) types.Vote {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteKey))
	var vote types.Vote
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.VoteKey+key)), &vote)
	return vote
}

// HasVote checks if the vote exists
func (k Keeper) HasVote(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteKey))
	return store.Has(types.KeyPrefix(types.VoteKey + id))
}

// GetVoteAccount returns the creator of the vote
func (k Keeper) GetVoteAccount(ctx sdk.Context, key string) string {
	return k.GetVote(ctx, key).Account
}

// GetVoteOwner returns the creator of the vote
func (k Keeper) GetVoteOwner(ctx sdk.Context, key string) string {
	return k.GetAccountOwner(ctx, k.GetVoteAccount(ctx, key))
}

// DeleteVote deletes a vote
func (k Keeper) DeleteVote(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteKey))
	store.Delete(types.KeyPrefix(types.VoteKey + key))
}

// GetAllVote returns all vote
func (k Keeper) GetAllVote(ctx sdk.Context) (msgs []types.Vote) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.VoteKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Vote
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}

// GetAllPollVote returns votes in a poll
func (k Keeper) GetAllPollVote(ctx sdk.Context, poll string) (msgs []types.Vote) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VoteKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.VoteKey+poll))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Vote
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}

// GetAllAccountVote returns votes in an account
func (k Keeper) GetAllAccountVote(ctx sdk.Context, account string) (msgs []types.Vote) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VotePtrKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.VotePtrKey+account))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var ptr types.VotePtr
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &ptr)
		msg := k.GetVote(ctx, ptr.Id)
		msgs = append(msgs, msg)
	}

	return
}
