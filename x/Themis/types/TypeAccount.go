package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Account struct {
	ID     string         `json:"id" yaml:"id"`
	Owner  sdk.AccAddress `json:"owner" yaml:"owner"`
	Group  string         `json:"group" yaml:"group"`
	Amount int            `json:"amount" yaml:"amount"`
}

type AccountPTR struct {
	ID string `json:"id" yaml:"id"`
}
