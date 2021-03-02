package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

func GetPollCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "poll",
		Short:                      "Poll subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreatePoll())
	cmd.AddCommand(CmdSetPollDesc())
	cmd.AddCommand(CmdExtendPollDeadline())

	return cmd
}

func CmdCreatePoll() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [group] [title] [description] [deadline] [options]",
		Short: "Creates a new poll",
		Args:  cobra.MinimumNArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsGroup := string(args[0])
			argsTitle := string(args[1])
			argsDescription := string(args[2])
			argsDeadline, _ := strconv.ParseInt(args[3], 10, 64)
			argsOptions := args[4:]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreatePoll(clientCtx.GetFromAddress().String(), string(argsGroup), string(argsTitle), string(argsDescription), argsOptions, argsDeadline)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdSetPollDesc() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-description [poll] [description]",
		Short: "Changes the poll's description",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsId := string(args[0])
			argsDescription := string(args[1])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetPollDesc(clientCtx.GetFromAddress().String(), argsId, argsDescription)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdExtendPollDeadline() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "extend-deadline [poll] [deadline]",
		Short: "Changes the poll's deadline to one greater than the current.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsId := string(args[0])
			argsDeadline, _ := strconv.ParseInt(args[1], 10, 64)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgExtendPollDeadline(clientCtx.GetFromAddress().String(), argsId, argsDeadline)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
