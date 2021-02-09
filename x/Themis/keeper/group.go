package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

// GetGroupCount get the total number of group
func (k Keeper) GetGroupCount(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.GroupCountPrefix)
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

// SetGroupCount set the total number of group
func (k Keeper) SetGroupCount(ctx sdk.Context, count int64) {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.GroupCountPrefix)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateGroup creates a group
// Initialized vouchers to 0
func (k Keeper) CreateGroup(ctx sdk.Context, msg types.MsgCreateGroup) {
	// Create the group
	count := k.GetGroupCount(ctx)
	var group = types.Group{
		Creator:  msg.Creator,
		ID:       strconv.FormatInt(count, 10),
		Name:     msg.Name,
		Vouchers: 0,
	}

	store := ctx.KVStore(k.storeKey)
	key := []byte(types.GroupPrefix + group.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(group)
	store.Set(key, value)

	// Update group count
	k.SetGroupCount(ctx, count+1)
}

// GetGroup returns the group information
func (k Keeper) GetGroup(ctx sdk.Context, key string) (types.Group, error) {
	store := ctx.KVStore(k.storeKey)
	var group types.Group
	byteKey := []byte(types.GroupPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &group)
	if err != nil {
		return group, err
	}
	return group, nil
}

// SetGroup sets a group
func (k Keeper) SetGroup(ctx sdk.Context, group types.Group) {
	groupKey := group.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(group)
	key := []byte(types.GroupPrefix + groupKey)
	store.Set(key, bz)
}

// DeleteGroup deletes a group
func (k Keeper) DeleteGroup(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.GroupPrefix + key))
}

//
// Functions used by querier
//

func listGroup(ctx sdk.Context, k Keeper) ([]byte, error) {
	var groupList []types.Group
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.GroupPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var group types.Group
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &group)
		groupList = append(groupList, group)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, groupList)
	return res, nil
}

func getGroup(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	group, err := k.GetGroup(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, group)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetGroupOwner(ctx sdk.Context, key string) sdk.AccAddress {
	group, err := k.GetGroup(ctx, key)
	if err != nil {
		return nil
	}
	return group.Creator
}

// Check if the key exists in the store
func (k Keeper) GroupExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.GroupPrefix + key))
}

// Custom functions

// Changes the group name
func (k Keeper) SetGroupName(ctx sdk.Context, key string, name string) {
	group, _ := k.GetGroup(ctx, key)
	group.Name = name
	k.SetGroup(ctx, group)
}

// Adds funds to the group
//func (k Keeper) AddGroupFunds(ctx sdk.Context, key string, newFunds float64) {
//	group, _ := k.GetGroup(ctx, key)
//	group.Funds += newFunds
//	k.SetGroup(ctx, group)
//}

// Adds  one voucher
func (k Keeper) AddGroupVoucher(ctx sdk.Context, key string) {
	group, _ := k.GetGroup(ctx, key)
	group.Vouchers += 1
	k.SetGroup(ctx, group)
}
