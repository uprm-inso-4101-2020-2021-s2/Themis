package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Group struct {
	Creator  sdk.AccAddress `json:"creator" yaml:"creator"`
	ID       string         `json:"id" yaml:"id"`
	Name     string         `json:"name" yaml:"name"`
	Accounts int            `json:"account" yaml:"accounts"`
}
