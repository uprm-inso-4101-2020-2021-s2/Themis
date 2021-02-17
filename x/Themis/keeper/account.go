package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
	"strconv"
)

// GetAccountCount get the total number of account
func (k Keeper) GetAccountCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountCountKey))
	byteKey := types.KeyPrefix(types.AccountCountKey)
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

// SetAccountCount set the total number of account
func (k Keeper) SetAccountCount(ctx sdk.Context, count int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountCountKey))
	byteKey := types.KeyPrefix(types.AccountCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateAccount creates a account with a new id and update the count with a related pointer
func (k Keeper) CreateAccount(ctx sdk.Context, msg types.MsgAddAccountVouchers) {
	// Create the account
	count := k.GetAccountCount(ctx)
	var account = types.Account{
		User:     msg.User,
		Id:       k.NewKey(msg.Group, msg.User),
		Group:    msg.Group,
		Vouchers: msg.Vouchers,
	}

	var accountPtr = types.AccountPTR{
		Id: account.Id,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountKey))
	storePtr := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountPtrKey))
	key := types.KeyPrefix(types.AccountKey + account.Id)
	keyPtr := types.KeyPrefix(types.AccountPtrKey + k.NewKey(msg.User, msg.Group))
	value := k.cdc.MustMarshalBinaryBare(&account)
	valuePtr := k.cdc.MustMarshalBinaryBare(&accountPtr)
	store.Set(key, value)
	storePtr.Set(keyPtr, valuePtr)

	// Update account count
	k.SetAccountCount(ctx, count+1)
}

// SetAccountVoucher Edits the account vouchers
func (k Keeper) SetAccountVoucher(ctx sdk.Context, key string, qty int64) {
	account := k.GetAccount(ctx, key)
	account.Vouchers += qty
	k.SetAccount(ctx, account)
}

// SetAccount set a specific account in the store
func (k Keeper) SetAccount(ctx sdk.Context, account types.Account) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountKey))
	b := k.cdc.MustMarshalBinaryBare(&account)
	store.Set(types.KeyPrefix(types.AccountKey+account.Id), b)
}

// GetAccount returns a account from its id
func (k Keeper) GetAccount(ctx sdk.Context, key string) types.Account {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountKey))
	var account types.Account
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.AccountKey+key)), &account)
	return account
}

// GetAccountId returns account ID from its user and group
func (k Keeper) GetAccountId(ctx sdk.Context, user string, group string) string {
	return k.GetAccount(ctx, k.NewKey(group, user)).Id
}

// HasAccount checks if the account exists
func (k Keeper) HasAccount(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountKey))
	return store.Has(types.KeyPrefix(types.AccountKey + id))
}

// AccountExistsInGroup checks if the account exists
func (k Keeper) AccountExistsInGroup(ctx sdk.Context, user string, group string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountKey))
	return store.Has(types.KeyPrefix(types.AccountKey + k.NewKey(group, user)))
}

// GetAccountOwner returns the creator of the account
func (k Keeper) GetAccountOwner(ctx sdk.Context, key string) string {
	return k.GetAccount(ctx, key).User
}

// GetAccountVouchers returns the total vouchers remaining
func (k Keeper) GetAccountVouchers(ctx sdk.Context, key string) int64 {
	return k.GetAccount(ctx, key).Vouchers
}

// EditAccountVouchers adds the given amount to the vouchers
func (k Keeper) EditAccountVouchers(ctx sdk.Context, key string, v int64) {
	account := k.GetAccount(ctx, key)
	account.Vouchers += v
	k.SetAccount(ctx, account)
}

// DeleteAccount deletes a account
func (k Keeper) DeleteAccount(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountKey))
	store.Delete(types.KeyPrefix(types.AccountKey + key))
}

// GetAllAccount returns all account
func (k Keeper) GetAllAccount(ctx sdk.Context) (msgs []types.Account) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.AccountKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Account
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}

// GetAllUserAccount returns all account
func (k Keeper) GetAllUserAccount(ctx sdk.Context, user string) (msgs []types.Account) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountPtrKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.AccountPtrKey+user))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var ptr types.AccountPTR
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &ptr)
		msg := k.GetAccount(ctx, ptr.Id)
		msgs = append(msgs, msg)
	}

	return
}

// GetAllGroupAccount returns all account
func (k Keeper) GetAllGroupAccount(ctx sdk.Context, group string) (msgs []types.Account) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AccountKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.AccountKey+group))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Account
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
