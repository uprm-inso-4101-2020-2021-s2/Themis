package Themis

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/keeper"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		case *types.MsgCreatePoll:
			return handleMsgCreatePoll(ctx, k, msg)

		case *types.MsgSetPollDesc:
			return handleMsgSetPollDesc(ctx, k, msg)

		case *types.MsgExtendPollDeadline:
			return handleMsgExtendPollDeadline(ctx, k, msg)

		case *types.MsgAddAccountVouchers:
			return handleMsgAddAccountVouchers(ctx, k, msg)

		case *types.MsgCreateGroup:
			return handleMsgCreateGroup(ctx, k, msg)

		case *types.MsgSetGroupName:
			return handleMsgSetGroupName(ctx, k, msg)

		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
