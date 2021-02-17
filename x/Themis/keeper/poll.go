package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
	"strconv"
)

// GetPollCount get the total number of poll
func (k Keeper) GetPollCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollCountKey))
	byteKey := types.KeyPrefix(types.PollCountKey)
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

// SetPollCount set the total number of poll
func (k Keeper) SetPollCount(ctx sdk.Context, count int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollCountKey))
	byteKey := types.KeyPrefix(types.PollCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreatePoll creates a poll with a new id and update the count
func (k Keeper) CreatePoll(ctx sdk.Context, msg types.MsgCreatePoll) {
	// Create the poll
	count := k.GetPollCount(ctx)
	var poll = types.Poll{
		Id:          k.NewKey(msg.Group, strconv.FormatInt(count, 10)),
		Group:       msg.Group,
		Title:       msg.Title,
		Description: msg.Description,
		Options:     msg.Options,
		Deadline:    msg.Deadline,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollKey))
	key := types.KeyPrefix(types.PollKey + poll.Id)
	value := k.cdc.MustMarshalBinaryBare(&poll)
	store.Set(key, value)

	// Update poll count
	k.SetPollCount(ctx, count+1)
}

// SetPoll set a specific poll in the store
func (k Keeper) SetPoll(ctx sdk.Context, poll types.Poll) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollKey))
	b := k.cdc.MustMarshalBinaryBare(&poll)
	store.Set(types.KeyPrefix(types.PollKey+poll.Id), b)
}

// ExtendPollDeadline extends the poll's deadline
func (k Keeper) ExtendPollDeadline(ctx sdk.Context, id string, deadline int64) {
	poll := k.GetPoll(ctx, id)
	poll.Deadline = deadline
	k.SetPoll(ctx, poll)
}

// SetPollDescription sets the poll's description
func (k Keeper) SetPollDescription(ctx sdk.Context, id string, desc string) {
	poll := k.GetPoll(ctx, id)
	poll.Description = desc
	k.SetPoll(ctx, poll)
}

// GetPoll returns a poll from its id
func (k Keeper) GetPoll(ctx sdk.Context, key string) types.Poll {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollKey))
	var poll types.Poll
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.PollKey+key)), &poll)
	return poll
}

// HasPoll checks if the poll exists
func (k Keeper) HasPoll(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollKey))
	return store.Has(types.KeyPrefix(types.PollKey + id))
}

// GetPollOwner returns the creator of the poll
func (k Keeper) GetPollOwner(ctx sdk.Context, key string) string {
	return k.GetGroupOwner(ctx, k.GetPollGroup(ctx, key))
}

// GetPollGroup returns the group the poll belongs to
func (k Keeper) GetPollGroup(ctx sdk.Context, key string) string {
	return k.GetPoll(ctx, key).Group
}

//GetPollOptions return the options for the poll
func (k Keeper) GetPollOptions(ctx sdk.Context, key string) []string {
	return k.GetPoll(ctx, key).Options
}

//GetPollDeadline return the deadline for the poll
func (k Keeper) GetPollDeadline(ctx sdk.Context, key string) int64 {
	return k.GetPoll(ctx, key).Deadline
}

// DeletePoll deletes a poll
func (k Keeper) DeletePoll(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollKey))
	store.Delete(types.KeyPrefix(types.PollKey + key))
}

// GetAllPoll returns all poll
func (k Keeper) GetAllPoll(ctx sdk.Context) (msgs []types.Poll) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.PollKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Poll
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}

// GetAllPoll returns all poll
func (k Keeper) GetAllGroupPoll(ctx sdk.Context, group string) (msgs []types.Poll) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PollKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.PollKey+group))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Poll
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
