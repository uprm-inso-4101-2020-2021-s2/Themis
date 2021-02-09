package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Voucher struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
	Owner   sdk.AccAddress `json:"owner" yaml:"owner"`
	Group   string         `json:"group" yaml:"group"`
	Used    int            `json:"used" yaml:"used"`
}
