package cli

import (
	"bufio"
	"strconv"

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

func GetCmdAddVote(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "give-vote [group] [receiver] [vote amount]",
		Short: "Adds vote to account, if not created then create one",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsGroup := string(args[0])
			argsOwner := string(args[1])
			argsVotes, _ := strconv.Atoi(string(args[2]))

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgAccountAddVotes(cliCtx.GetFromAddress(), _GetAddress(argsOwner), string(argsGroup), argsVotes)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
