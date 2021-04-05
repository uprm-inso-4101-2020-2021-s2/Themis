package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
	"strconv"
)

// GetAccountCount get the total number of account
func (k Keeper) GetAccountCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountCountKey))
	byteKey := types.KeyPrefix(types.AccountCountKey)
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

// SetAccountCount set the total number of account
func (k Keeper) SetAccountCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountCountKey))
	byteKey := types.KeyPrefix(types.AccountCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendAccount appends a account in the store with a new id and update the count
func (k Keeper) AppendAccount(
	ctx sdk.Context,
	creator string,
	name string,
) uint64 {
	// Create the account
	groupsMap := make(map[uint64]uint64)

	count := k.GetAccountCount(ctx)
	var account = types.Account{
		Creator: creator,
		Id:      count,
		Name:    name,
		Groups:  groupsMap,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountKey))
	value := k.cdc.MustMarshalBinaryBare(&account)
	store.Set(GetAccountIDBytes(account.Id), value)

	// Store by name
	countStr := strconv.FormatUint(count, 10)
	prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountNameKey)).Set(types.GetStringBytes(name+"-"+countStr), value)
	prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountAddrKey)).Set(types.GetStringBytes(creator+"-"+countStr), value)

	// Update account count
	k.SetAccountCount(ctx, count+1)

	return count
}

// SetAccount set a specific account in the store
func (k Keeper) SetAccount(ctx sdk.Context, account types.Account) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountKey))
	b := k.cdc.MustMarshalBinaryBare(&account)
	store.Set(GetAccountIDBytes(account.Id), b)

	idStr := strconv.FormatUint(account.Id, 10)
	// Extra account key stores
	// TODO: delete old Name and Addr entry
	prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountNameKey)).Set(types.GetStringBytes(account.Name+"-"+idStr), b)
	prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountAddrKey)).Set(types.GetStringBytes(account.Creator+"-"+idStr), b)
}

// ChangeAccountName changes the account's name
func (k Keeper) ChangeAccountName(ctx sdk.Context, id uint64, name string) {
	acc := k.GetAccount(ctx, id)
	acc.Name = name
	k.SetAccount(ctx, acc)
}

// AddAccountGroup adds another group to the account
func (k Keeper) AddAccountGroup(ctx sdk.Context, id uint64, group uint64, date uint64) {
	acc := k.GetAccount(ctx, id)
	acc.Groups[group] = date
	k.SetAccount(ctx, acc)
}

// AccountInGroup checks if account already in group
func (k Keeper) AccountInGroup(ctx sdk.Context, id uint64, group uint64) bool {
	acc := k.GetAccount(ctx, id)
	if _, ok := acc.Groups[group]; ok {
		return true
	}
	return false
}

// GetAccountId gets account id from address
func (k Keeper) GetAccountId(ctx sdk.Context, addr string) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountAddrKey))
	var account types.Account
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.GetStringBytes(addr)), &account)
	return account.Id
}

// GetAccount returns a account from its id
func (k Keeper) GetAccount(ctx sdk.Context, id uint64) types.Account {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountKey))
	var account types.Account
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetAccountIDBytes(id)), &account)
	return account
}

// HasAccount checks if the account exists in the store
func (k Keeper) HasAccount(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountKey))
	return store.Has(GetAccountIDBytes(id))
}

// HasAccountAddr checks if the account exists in the store
func (k Keeper) HasAccountAddr(ctx sdk.Context, addr string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountNameKey))
	return store.Has(types.GetStringBytes(addr))
}

// GetAccountOwner returns the creator of the account
func (k Keeper) GetAccountOwner(ctx sdk.Context, id uint64) string {
	return k.GetAccount(ctx, id).Creator
}

// RemoveAccount removes a account from the store
func (k Keeper) RemoveAccount(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountKey))
	store.Delete(GetAccountIDBytes(id))
}

// GetAllAccount returns all account
func (k Keeper) GetAllAccount(ctx sdk.Context) (list []types.Account) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Account
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAccountIDBytes returns the byte representation of the ID
func GetAccountIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetAccountIDFromBytes returns ID in uint64 format from a byte array
func GetAccountIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
