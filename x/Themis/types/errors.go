package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/Themis module sentinel errors
var (
	ErrNoFunds           = sdkerrors.Register(ModuleName, 1100, "No funds available")
	ErrInvalidGroup      = sdkerrors.Register(ModuleName, 1101, "Invalid group ID")
	ErrNameEmpty         = sdkerrors.Register(ModuleName, 1102, "Group name empty")
	ErrNoPermission      = sdkerrors.Register(ModuleName, 1103, "Not group owner")
	ErrVoteNumberInvalid = sdkerrors.Register(ModuleName, 1104, "Votes must be 0 or more")
	ErrInvalidPoll       = sdkerrors.Register(ModuleName, 1105, "Invalid poll ID")
	ErrInvalidPollTitle  = sdkerrors.Register(ModuleName, 1106, "Invalid poll title")
	ErrInvalidPollDate   = sdkerrors.Register(ModuleName, 1107, "Invalid poll deadline")
	ErrInvalidPollDesc   = sdkerrors.Register(ModuleName, 1108, "Invalid poll description")
	ErrInvalidPollOptns  = sdkerrors.Register(ModuleName, 1109, "Invalid poll options")
)
