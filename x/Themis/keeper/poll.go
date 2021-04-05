package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
	"strconv"
)

// GetPollCount get the total number of poll
func (k Keeper) GetPollCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollCountKey))
	byteKey := types.KeyPrefix(types.PollCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to iint64
		panic("cannot decode count")
	}

	return count
}

// SetPollCount set the total number of poll
func (k Keeper) SetPollCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollCountKey))
	byteKey := types.KeyPrefix(types.PollCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendPoll appends a poll in the store with a new id and update the count
func (k Keeper) AppendPoll(
	ctx sdk.Context,
	creator string,
	name string,
	group uint64,
	votes []string,
	description string,
	deadline uint64,
) uint64 {
	votesMap := make(map[string]uint64)

	for _, vote := range votes {
		votesMap[vote] = 0
	}

	// Create the poll
	count := k.GetPollCount(ctx)
	var poll = types.Poll{
		Creator:     creator,
		Id:          count,
		Name:        name,
		Group:       group,
		Votes:       votesMap,
		Description: description,
		Deadline:    deadline,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollKey))
	value := k.cdc.MustMarshalBinaryBare(&poll)
	store.Set(GetPollIDBytes(poll.Id), value)

	// Store by name
	// TODO: add a key to sort by deadline
	countStr := strconv.FormatUint(count, 10)
	groupStr := strconv.FormatUint(group, 10)
	prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollNameKey)).Set(types.GetStringBytes(name+"-"+countStr), value)
	prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollGroupKey)).Set(types.GetStringBytes(groupStr+"-"+name+"-"+countStr), value)

	// Update poll count
	k.SetPollCount(ctx, count+1)

	return count
}

// SetPoll set a specific poll in the store
func (k Keeper) SetPoll(ctx sdk.Context, poll types.Poll) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollKey))
	b := k.cdc.MustMarshalBinaryBare(&poll)
	store.Set(GetPollIDBytes(poll.Id), b)

	idStr := strconv.FormatUint(poll.Id, 10)
	groupStr := strconv.FormatUint(poll.Group, 10)
	// Extra account key stores
	prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollNameKey)).Set(types.GetStringBytes(poll.Name+"-"+idStr), b)
	prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollGroupKey)).Set(types.GetStringBytes(groupStr+"-"+poll.Name+"-"+idStr), b)
}

// AddPollVote sums a vote to the poll keyval
func (k Keeper) AddPollVote(ctx sdk.Context, id uint64, vote string) {
	poll := k.GetPoll(ctx, id)
	poll.Votes[vote]++
	k.SetPoll(ctx, poll)
}

// RemovePollVote removes one vote to the pollkeyval
func (k Keeper) RemovePollVote(ctx sdk.Context, id uint64, vote string) {
	poll := k.GetPoll(ctx, id)
	poll.Votes[vote]--
	k.SetPoll(ctx, poll)
}

// PollVoteExists checks if vote exists
func (k Keeper) PollVoteExists(ctx sdk.Context, id uint64, vote string) bool {
	poll := k.GetPoll(ctx, id)
	if _, ok := poll.Votes[vote]; ok {
		return true
	}
	return false
}

// GetPoll returns a poll from its id
func (k Keeper) GetPoll(ctx sdk.Context, id uint64) types.Poll {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollKey))
	var poll types.Poll
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetPollIDBytes(id)), &poll)
	return poll
}

// HasPoll checks if the poll exists in the store
func (k Keeper) HasPoll(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollKey))
	return store.Has(GetPollIDBytes(id))
}

// GetPollOwner returns the creator of the poll
func (k Keeper) GetPollOwner(ctx sdk.Context, id uint64) string {
	return k.GetPoll(ctx, id).Creator
}

// RemovePoll removes a poll from the store
func (k Keeper) RemovePoll(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollKey))
	store.Delete(GetPollIDBytes(id))
}

// GetAllPoll returns all poll
func (k Keeper) GetAllPoll(ctx sdk.Context) (list []types.Poll) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Poll
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetPollIDBytes returns the byte representation of the ID
func GetPollIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetPollIDFromBytes returns ID in uint64 format from a byte array
func GetPollIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
