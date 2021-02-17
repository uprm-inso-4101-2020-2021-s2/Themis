package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

func CmdCreateVote() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-vote [poll] [option]",
		Short: "Creates a new vote",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsPoll := string(args[0])
			argsOption, _ := strconv.ParseInt(args[1], 10, 32)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateVote(clientCtx.GetFromAddress().String(), argsPoll, int32(argsOption))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
