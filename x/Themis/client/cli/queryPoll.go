package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

// GetQueryPollCmd returns the cli query commands for this module
func GetQueryPollCmd(queryRoute string) *cobra.Command {
	// Group Themis queries under a subcommand
	cmd := &cobra.Command{
		Use:                        "poll",
		Short:                      "Manages poll queries",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1

	cmd.AddCommand(CmdListPoll())
	cmd.AddCommand(CmdShowPoll())
	cmd.AddCommand(CmdListPollsWithName())
	cmd.AddCommand(CmdListPollsInGroup())

	return cmd
}

func CmdListPoll() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list all poll",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllPollRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.PollAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowPoll() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show [id]",
		Short: "shows a poll",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetPollRequest{
				Id: id,
			}

			res, err := queryClient.Poll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdListPollsWithName() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "name [name]",
		Short: "list all poll with that name",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllPollsWithNameRequest{
				Name:       args[0],
				Pagination: pageReq,
			}

			res, err := queryClient.PollsWithName(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdListPollsInGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "group [group] [name]",
		Short: "list all poll in that group and optionally name",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			groupStr, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			if len(args) == 1 {
				params := &types.QueryAllPollsInGroupRequest{
					Group:      groupStr,
					Pagination: pageReq,
				}

				res, err := queryClient.PollsInGroup(context.Background(), params)
				if err != nil {
					return err
				}

				return clientCtx.PrintProto(res)
			} else {
				params := &types.QueryAllPollsInGroupWithNameRequest{
					Group:      groupStr,
					Name:       args[1],
					Pagination: pageReq,
				}

				res, err := queryClient.PollsInGroupWithName(context.Background(), params)
				if err != nil {
					return err
				}

				return clientCtx.PrintProto(res)
			}
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
