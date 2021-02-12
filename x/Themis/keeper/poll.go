package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

// GetPollCount get the total number of poll
func (k Keeper) GetPollCount(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.PollCountPrefix)
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
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.PollCountPrefix)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreatePoll creates a poll
func (k Keeper) CreatePoll(ctx sdk.Context, msg types.MsgCreatePoll) {
	// Create the poll
	count := k.GetPollCount(ctx)
	var poll = types.Poll{
		Creator:     msg.Creator,
		ID:          buildKey(msg.Group, strconv.FormatInt(count, 10)),
		Group:       msg.Group,
		Title:       msg.Title,
		Description: msg.Description,
		Options:     msg.Options,
		Deadline:    msg.Deadline,
	}

	store := ctx.KVStore(k.storeKey)
	key := []byte(types.PollPrefix + poll.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(poll)
	store.Set(key, value)

	// Update poll count
	k.SetPollCount(ctx, count+1)
}

// GetPoll returns the poll information
func (k Keeper) GetPoll(ctx sdk.Context, key string) (types.Poll, error) {
	store := ctx.KVStore(k.storeKey)
	var poll types.Poll
	byteKey := []byte(types.PollPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &poll)
	if err != nil {
		return poll, err
	}
	return poll, nil
}

// SetPoll sets a poll
func (k Keeper) SetPoll(ctx sdk.Context, poll types.Poll) {
	pollKey := poll.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(poll)
	key := []byte(types.PollPrefix + pollKey)
	store.Set(key, bz)
}

// DeletePoll deletes a poll
func (k Keeper) DeletePoll(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.PollPrefix + key))
}

func (k Keeper) SetPollDeadline(ctx sdk.Context, key string, deadline int) {
	poll, _ := k.GetPoll(ctx, key)
	poll.Deadline = deadline
	k.SetPoll(ctx, poll)
}

func (k Keeper) SetPollDescription(ctx sdk.Context, key string, desc string) {
	poll, _ := k.GetPoll(ctx, key)
	poll.Description = desc
	k.SetPoll(ctx, poll)
}

//
// Functions used by querier
//

func listPoll(ctx sdk.Context, k Keeper) ([]byte, error) {
	var pollList []types.Poll
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.PollPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var poll types.Poll
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &poll)
		pollList = append(pollList, poll)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, pollList)
	return res, nil
}

func listGroupPolls(ctx sdk.Context, path []string, k Keeper) ([]byte, error) {
	group := path[0]

	var pollList []types.Poll
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.PollPrefix+group))
	for ; iterator.Valid(); iterator.Next() {
		var poll types.Poll
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &poll)
		pollList = append(pollList, poll)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, pollList)
	return res, nil
}

func getPoll(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	poll, err := k.GetPoll(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, poll)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetPollOwner(ctx sdk.Context, key string) sdk.AccAddress {
	poll, err := k.GetPoll(ctx, key)
	if err != nil {
		return nil
	}
	return poll.Creator
}

func (k Keeper) GetPollGroup(ctx sdk.Context, key string) string {
	poll, err := k.GetPoll(ctx, key)
	if err != nil {
		return ""
	}
	return poll.Group
}

func (k Keeper) GetPollDeadline(ctx sdk.Context, key string) int {
	poll, err := k.GetPoll(ctx, key)
	if err != nil {
		return -1
	}
	return poll.Deadline
}

// Check if the key exists in the store
func (k Keeper) PollExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.PollPrefix + key))
}
