package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

// GetQueryGroupCmd returns the cli query commands for this module
func GetQueryGroupCmd(queryRoute string) *cobra.Command {
	// Group Themis queries under a subcommand
	cmd := &cobra.Command{
		Use:                        "group",
		Short:                      "Manages group queries",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1

	cmd.AddCommand(CmdListGroup())
	cmd.AddCommand(CmdShowGroup())
	cmd.AddCommand(CmdListGroupAddress())
	cmd.AddCommand(CmdListGroupWithNames())

	return cmd
}

func CmdListGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list all group",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllGroupRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.GroupAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show [id]",
		Short: "shows a group",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetGroupRequest{
				Id: id,
			}

			res, err := queryClient.Group(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdListGroupAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "address [addr]",
		Short: "list all groups under that address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetGroupAddressRequest{
				Addr:       args[0],
				Pagination: pageReq,
			}

			res, err := queryClient.GroupAddress(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdListGroupWithNames() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "name [name]",
		Short: "list all group under that name",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllGroupWithNamesRequest{
				Name:       args[0],
				Pagination: pageReq,
			}

			res, err := queryClient.GroupWithNames(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
