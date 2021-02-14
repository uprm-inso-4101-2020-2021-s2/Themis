package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
	"strconv"
)

// GetGroupCount get the total number of group
func (k Keeper) GetGroupCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GroupCountKey))
	byteKey := types.KeyPrefix(types.GroupCountKey)
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
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GroupCountKey))
	byteKey := types.KeyPrefix(types.GroupCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateGroup creates a group with a new id and update the count
func (k Keeper) CreateGroup(ctx sdk.Context, msg types.MsgCreateGroup) {
	// Create the group
	count := k.GetGroupCount(ctx)
	var group = types.Group{
		Creator: msg.Creator,
		Id:      strconv.FormatInt(count, 10),
		Name:    msg.Name,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GroupKey))
	key := types.KeyPrefix(types.GroupKey + group.Id)
	value := k.cdc.MustMarshalBinaryBare(&group)
	store.Set(key, value)

	// Update group count
	k.SetGroupCount(ctx, count+1)
}

// SetGroup set a specific group in the store
func (k Keeper) SetGroup(ctx sdk.Context, group types.Group) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GroupKey))
	b := k.cdc.MustMarshalBinaryBare(&group)
	store.Set(types.KeyPrefix(types.GroupKey+group.Id), b)
}

// Changes the groups name
func (k Keeper) SetGroupName(ctx sdk.Context, key string, name string) {
	group := k.GetGroup(ctx, key)
	group.Name = name
	k.SetGroup(ctx, group)
}

// GetGroup returns a group from its id
func (k Keeper) GetGroup(ctx sdk.Context, key string) types.Group {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GroupKey))
	var group types.Group
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.GroupKey+key)), &group)
	return group
}

// HasGroup checks if the group exists
func (k Keeper) HasGroup(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GroupKey))
	return store.Has(types.KeyPrefix(types.GroupKey + id))
}

// GetGroupOwner returns the creator of the group
func (k Keeper) GetGroupOwner(ctx sdk.Context, key string) string {
	return k.GetGroup(ctx, key).Creator
}

// DeleteGroup deletes a group
func (k Keeper) DeleteGroup(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GroupKey))
	store.Delete(types.KeyPrefix(types.GroupKey + key))
}

// GetAllGroup returns all group
func (k Keeper) GetAllGroup(ctx sdk.Context) (msgs []types.Group) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GroupKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.GroupKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Group
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
