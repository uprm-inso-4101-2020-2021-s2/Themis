package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

// GetAccountCount get the total number of Account
func (k Keeper) GetAccountCount(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.AccountCountPrefix)
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

// SetAccountCount set the total number of Account
func (k Keeper) SetAccountCount(ctx sdk.Context, count int64) {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.AccountCountPrefix)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateAccount creates a Account
func (k Keeper) CreateAccount(ctx sdk.Context, msg types.MsgAccountAddVotes) {
	// Create the Account
	count := k.GetAccountCount(ctx)
	var account = types.Account{
		Owner:  msg.Owner,
		ID:     buildKey(msg.Owner.String(), string(msg.Group)),
		Group:  msg.Group,
		Amount: msg.Amount,
	}

	var accountPrt = types.AccountPTR{
		ID: account.ID,
	}
	//

	store := ctx.KVStore(k.storeKey)
	key := []byte(types.AccountPrefix + string(account.ID))
	value := k.cdc.MustMarshalBinaryLengthPrefixed(account)
	store.Set(key, value)

	// Reserved for listing by group instead of user
	ptrKey := []byte(types.AccountPtrPrefix + string(msg.Group) + "-" + msg.Owner.String())
	ptrValue := k.cdc.MustMarshalBinaryLengthPrefixed(accountPrt)
	store.Set(ptrKey, ptrValue)

	// Update account count
	k.SetAccountCount(ctx, count+1)
}

// GetAccount returns the account information
func (k Keeper) GetAccount(ctx sdk.Context, key string) (types.Account, error) {
	store := ctx.KVStore(k.storeKey)
	var account types.Account
	byteKey := []byte(types.AccountPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &account)
	if err != nil {
		return account, err
	}
	return account, nil
}

// SetAccount sets a account
func (k Keeper) SetAccount(ctx sdk.Context, account types.Account) {
	accountKey := account.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(account)
	key := []byte(types.AccountPrefix + accountKey)
	store.Set(key, bz)
}

// DeleteAccount deletes a account
func (k Keeper) DeleteAccount(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.AccountPrefix + key))
}

//
// Functions used by querier
//

func listAccount(ctx sdk.Context, k Keeper) ([]byte, error) {
	var accountList []types.Account
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.AccountPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var account types.Account
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &account)
		accountList = append(accountList, account)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, accountList)
	return res, nil
}

// Queries all accounts associated with that user
func listUserAccounts(ctx sdk.Context, path []string, k Keeper) ([]byte, error) {
	user := path[0]
	var accountList []types.Account
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.AccountPrefix+user))
	for ; iterator.Valid(); iterator.Next() {
		var account types.Account
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &account)
		accountList = append(accountList, account)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, accountList)
	return res, nil
}

// Queries all accounts associated with that group
func listGroupAccounts(ctx sdk.Context, path []string, k Keeper) ([]byte, error) {
	group := path[0]
	var accountList []types.Account
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.AccountPtrPrefix+group))
	for ; iterator.Valid(); iterator.Next() {
		var accountPtr types.AccountPTR
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &accountPtr)
		// Get account associated to ptr
		account, _ := k.GetAccount(ctx, accountPtr.ID)
		accountList = append(accountList, account)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, accountList)
	return res, nil
}

func getAccount(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	account, err := k.GetAccount(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, account)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetAccountOwner(ctx sdk.Context, key string) sdk.AccAddress {
	account, err := k.GetAccount(ctx, key)
	if err != nil {
		return nil
	}
	return account.Owner
}

// Check if the key exists in the store
func (k Keeper) AccountExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.AccountPrefix + key))
}

// Check if the account exists in the store
func (k Keeper) UserGroupAccountExists(ctx sdk.Context, user string, group string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.AccountPrefix + buildKey(user, group)))
}

// Custom Functions

// Sets account as used
func (k Keeper) UseAccount(ctx sdk.Context, user string, group string) {
	account, _ := k.GetAccount(ctx, buildKey(user, group))
	account.Amount -= 1
	k.SetAccount(ctx, account)
}

func (k Keeper) AddToAccount(ctx sdk.Context, user string, group string, votes int) {
	account, _ := k.GetAccount(ctx, buildKey(user, group))
	account.Amount += votes
	k.SetAccount(ctx, account)
}
