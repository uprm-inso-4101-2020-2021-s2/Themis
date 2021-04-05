package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/Themis module sentinel errors
var (
	ErrInvalidName  = sdkerrors.Register(ModuleName, 1100, "Name is invalid")
	ErrInvalidDate  = sdkerrors.Register(ModuleName, 1101, "Date is invalid")
	ErrInvalidVotes = sdkerrors.Register(ModuleName, 1102, "Votes are invalid")
)
