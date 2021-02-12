package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrNoFunds            = sdkerrors.Register(ModuleName, 1, "No funds available")
	ErrInvalidGroup       = sdkerrors.Register(ModuleName, 2, "Invalid group ID")
	ErrNameEmpty          = sdkerrors.Register(ModuleName, 3, "Group name empty")
	ErrNoPermission       = sdkerrors.Register(ModuleName, 4, "Not group owner")
	ErrVoteNumberInvalid  = sdkerrors.Register(ModuleName, 5, "Votes must be 0 or more")
	ErrInvalidPoll        = sdkerrors.Register(ModuleName, 6, "Invalid poll ID")
	ErrTitleEmpty         = sdkerrors.Register(ModuleName, 7, "Poll title is emmpty")
	ErrDescLimitPassed    = sdkerrors.Register(ModuleName, 8, "Description passed the letter limit")
	ErrOptionsEmpty       = sdkerrors.Register(ModuleName, 9, "Poll options are empty")
	ErrDeadlineAlreadyMet = sdkerrors.Register(ModuleName, 10, "Poll deadline already met")
	ErrDeadlineInvalid    = sdkerrors.Register(ModuleName, 11, "Poll deadline invalid")
)
