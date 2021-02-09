package cli

import (
	"bufio"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

func _GetAddress(addr string) sdk.AccAddress {
	ret, _ := sdk.AccAddressFromBech32(addr)
	return ret
}

func GetCmdCreateVoucher(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-voucher [group] [receiver]",
		Short: "Creates a new voucher",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsGroup := string(args[0])
			argsOwner := string(args[1])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateVoucher(cliCtx.GetFromAddress(), _GetAddress(argsOwner), string(argsGroup))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
