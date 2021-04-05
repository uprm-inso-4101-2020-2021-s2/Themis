package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
	"strconv"
)

// GetGroupCount get the total number of group
func (k Keeper) GetGroupCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GroupCountKey))
	byteKey := types.KeyPrefix(types.GroupCountKey)
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

// SetGroupCount set the total number of group
func (k Keeper) SetGroupCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GroupCountKey))
	byteKey := types.KeyPrefix(types.GroupCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendGroup appends a group in the store with a new id and update the count
func (k Keeper) AppendGroup(
	ctx sdk.Context,
	name string,
	owner string,
) uint64 {
	// Create the group
	count := k.GetGroupCount(ctx)
	var group = types.Group{
		Id:    count,
		Name:  name,
		Owner: owner,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GroupKey))
	value := k.cdc.MustMarshalBinaryBare(&group)
	store.Set(GetGroupIDBytes(group.Id), value)

	// Store by name
	countStr := strconv.FormatUint(count, 10)
	prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GroupNameKey)).Set(types.GetStringBytes(name+"-"+countStr), value)
	prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GroupAddrKey)).Set(types.GetStringBytes(owner+"-"+countStr), value)

	// Update group count
	k.SetGroupCount(ctx, count+1)

	return count
}

// SetGroup set a specific group in the store
func (k Keeper) SetGroup(ctx sdk.Context, group types.Group) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GroupKey))
	b := k.cdc.MustMarshalBinaryBare(&group)
	store.Set(GetGroupIDBytes(group.Id), b)

	idStr := strconv.FormatUint(group.Id, 10)
	// Extra account key stores
	// TODO: delete old Name and Addr entry
	prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GroupNameKey)).Set(types.GetStringBytes(group.Name+"-"+idStr), b)
	prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GroupAddrKey)).Set(types.GetStringBytes(group.Owner+"-"+idStr), b)
}

// ChangeGroupName changes the group's name
func (k Keeper) ChangeGroupName(ctx sdk.Context, id uint64, name string) {
	group := k.GetGroup(ctx, id)
	group.Name = name
	k.SetGroup(ctx, group)
}

// ChangeGroupOwner changes the group's name
func (k Keeper) ChangeGroupOwner(ctx sdk.Context, id uint64, owner string) {
	group := k.GetGroup(ctx, id)
	group.Owner = owner
	k.SetGroup(ctx, group)
}

// GetGroup returns a group from its id
func (k Keeper) GetGroup(ctx sdk.Context, id uint64) types.Group {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GroupKey))
	var group types.Group
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetGroupIDBytes(id)), &group)
	return group
}

// HasGroup checks if the group exists in the store
func (k Keeper) HasGroup(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GroupKey))
	return store.Has(GetGroupIDBytes(id))
}

// HasGroupAddr checks if the group exists in the store
func (k Keeper) HasGroupAddr(ctx sdk.Context, addr string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GroupAddrKey))
	return store.Has(types.GetStringBytes(addr))
}

// GetGroupOwner returns the creator of the group
func (k Keeper) GetGroupOwner(ctx sdk.Context, id uint64) string {
	return k.GetGroup(ctx, id).Owner
}

// RemoveGroup removes a group from the store
func (k Keeper) RemoveGroup(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GroupKey))
	store.Delete(GetGroupIDBytes(id))
}

// GetAllGroup returns all group
func (k Keeper) GetAllGroup(ctx sdk.Context) (list []types.Group) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GroupKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Group
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetGroupIDBytes returns the byte representation of the ID
func GetGroupIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetGroupIDFromBytes returns ID in uint64 format from a byte array
func GetGroupIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
