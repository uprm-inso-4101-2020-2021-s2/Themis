package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

// GetVoucherCount get the total number of voucher
func (k Keeper) GetVoucherCount(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.VoucherCountPrefix)
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

// SetVoucherCount set the total number of voucher
func (k Keeper) SetVoucherCount(ctx sdk.Context, count int64) {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.VoucherCountPrefix)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateVoucher creates a voucher
func (k Keeper) CreateVoucher(ctx sdk.Context, msg types.MsgCreateVoucher) {
	// Create the voucher
	count := k.GetVoucherCount(ctx)
	var voucher = types.Voucher{
		Creator: msg.Creator,
		Owner:   msg.Owner,
		ID:      msg.Owner.String() + "-" + string(msg.Group) + "-" + strconv.FormatInt(count, 10),
		Group:   msg.Group,
		Used:    0,
	}
	//

	store := ctx.KVStore(k.storeKey)
	key := []byte(types.VoucherPrefix + string(voucher.ID))
	value := k.cdc.MustMarshalBinaryLengthPrefixed(voucher)
	store.Set(key, value)

	// Update voucher count
	k.SetVoucherCount(ctx, count+1)
}

// GetVoucher returns the voucher information
func (k Keeper) GetVoucher(ctx sdk.Context, key string) (types.Voucher, error) {
	store := ctx.KVStore(k.storeKey)
	var voucher types.Voucher
	byteKey := []byte(types.VoucherPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &voucher)
	if err != nil {
		return voucher, err
	}
	return voucher, nil
}

// SetVoucher sets a voucher
func (k Keeper) SetVoucher(ctx sdk.Context, voucher types.Voucher) {
	voucherKey := voucher.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(voucher)
	key := []byte(types.VoucherPrefix + voucherKey)
	store.Set(key, bz)
}

// DeleteVoucher deletes a voucher
func (k Keeper) DeleteVoucher(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.VoucherPrefix + key))
}

//
// Functions used by querier
//

func listVoucher(ctx sdk.Context, k Keeper) ([]byte, error) {
	var voucherList []types.Voucher
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.VoucherPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var voucher types.Voucher
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &voucher)
		voucherList = append(voucherList, voucher)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, voucherList)
	return res, nil
}

// Queries all vouchers associated with that user and group
func listUserVouchers(ctx sdk.Context, path []string, k Keeper) ([]byte, error) {
	userGroup := path[0]
	var voucherList []types.Voucher
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.VoucherPrefix+userGroup))
	for ; iterator.Valid(); iterator.Next() {
		var voucher types.Voucher
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &voucher)
		voucherList = append(voucherList, voucher)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, voucherList)
	return res, nil
}

func getVoucher(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	voucher, err := k.GetVoucher(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, voucher)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetVoucherOwner(ctx sdk.Context, key string) sdk.AccAddress {
	voucher, err := k.GetVoucher(ctx, key)
	if err != nil {
		return nil
	}
	return voucher.Owner
}

// Check if the key exists in the store
func (k Keeper) VoucherExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.VoucherPrefix + key))
}

// Custom Functions

// Sets voucher as used
func (k Keeper) UseVoucher(ctx sdk.Context, key string) {
	voucher, _ := k.GetVoucher(ctx, key)
	voucher.Used = 1
	k.SetVoucher(ctx, voucher)
}
