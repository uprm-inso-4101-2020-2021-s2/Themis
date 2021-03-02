package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

func GetGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "group",
		Short:                      "Group subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreateGroup())
	cmd.AddCommand(CmdSetGroupName())

	return cmd
}

func CmdCreateGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [name]",
		Short: "Creates a new group",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsName := string(args[0])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateGroup(clientCtx.GetFromAddress().String(), string(argsName))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdSetGroupName() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-name [group] [name]",
		Short: "Changes group's name",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsID := string(args[0])
			argsName := string(args[1])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetGroupName(clientCtx.GetFromAddress().String(), argsID, argsName)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
