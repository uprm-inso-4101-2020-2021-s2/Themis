package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

// GetPollCmd returns the commands for this module
func GetPollCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "poll",
		Short:                      "Manages poll txs",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreatePoll())
	cmd.AddCommand(CmdUpdatePoll())
	cmd.AddCommand(CmdDeletePoll())

	return cmd
}

func CmdCreatePoll() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [group] [name] [description] [deadline] [options...]",
		Short: "Creates a new poll",
		Args:  cobra.MinimumNArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsGroup, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			argsName := args[1]
			argsDescription := args[2]
			argsDeadline, err := strconv.ParseUint(args[3], 10, 64)
			if err != nil {
				return err
			}

			argsOptions := args[4:]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreatePoll(clientCtx.GetFromAddress().String(), argsName, argsGroup, argsOptions, argsDescription, argsDeadline)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdatePoll() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [id] [description] [deadline]",
		Short: "Update a poll",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsDescription := string(args[1])
			argsDeadline, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdatePoll(clientCtx.GetFromAddress().String(), id, argsDescription, argsDeadline)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeletePoll() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [id]",
		Short: "Delete a poll by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeletePoll(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
