package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

// GetQueryAccountCmd returns the cli query commands for this module
func GetQueryAccountCmd(queryRoute string) *cobra.Command {
	// Group Themis queries under a subcommand
	cmd := &cobra.Command{
		Use:                        "account",
		Short:                      "Manages account queries",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1

	cmd.AddCommand(CmdListAccount())
	cmd.AddCommand(CmdShowAccount())
	cmd.AddCommand(CmdShowAccountAddr())
	cmd.AddCommand(CmdShowAccountWithNames())

	return cmd
}

func CmdListAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list all account",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllAccountRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.AccountAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show [id]",
		Short: "shows a account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetAccountRequest{
				Id: id,
			}

			res, err := queryClient.Account(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowAccountAddr() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "address [addr]",
		Short: "shows a account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			addr := args[0]

			params := &types.QueryGetAccountAddressRequest{
				Addr: addr,
			}

			res, err := queryClient.AccountAddress(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowAccountWithNames() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "name [name]",
		Short: "list all account with that name",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			name := args[0]

			params := &types.QueryAllAccountWithNamesRequest{
				Name:       name,
				Pagination: pageReq,
			}

			res, err := queryClient.AccountWithNames(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
