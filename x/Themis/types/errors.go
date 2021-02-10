package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalid           = sdkerrors.Register(ModuleName, 1, "custom error message")
	ErrNoFunds           = sdkerrors.Register(ModuleName, 2, "No funds available")
	ErrInvalidGroup      = sdkerrors.Register(ModuleName, 4, "Invalid group ID")
	ErrNameEmpty         = sdkerrors.Register(ModuleName, 5, "Group name empty")
	ErrNoPermission      = sdkerrors.Register(ModuleName, 6, "Not group owner")
	ErrVoteNumberInvalid = sdkerrors.Register(ModuleName, 7, "Votes must be 0 or more")
)
