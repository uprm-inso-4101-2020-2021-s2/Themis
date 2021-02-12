package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Poll struct {
	Creator     sdk.AccAddress `json:"creator" yaml:"creator"`
	ID          string         `json:"id" yaml:"id"`
	Group       string         `json:"group" yaml:"group"`
	Title       string         `json:"title" yaml:"title"`
	Description string         `json:"description" yaml:"description"`
	Options     []string       `json:"options" yaml:"options"`
	Deadline    int            `json:"deadline" yaml:"deadline"`
}
