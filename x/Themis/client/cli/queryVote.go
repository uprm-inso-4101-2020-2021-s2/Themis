package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

func GetQueryVoteCmd(queryRoute string) *cobra.Command {
	// Group Themis queries under a subcommand
	cmd := &cobra.Command{
		Use:                        "vote",
		Short:                      "Querying commands for votes",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdListVote())
	cmd.AddCommand(CmdListUserVote())
	cmd.AddCommand(CmdListPollVote())
	cmd.AddCommand(CmdShowVote())

	return cmd
}

func CmdListVote() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list all vote",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllVoteRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.VoteAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

//TODO: both UserVote and PollVote not showing up in cmd
func CmdListUserVote() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user [user]",
		Short: "list all votes from a user",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllUserVoteRequest{
				User:       args[0],
				Pagination: pageReq,
			}

			res, err := queryClient.UserVoteAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdListPollVote() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "poll [poll]",
		Short: "list all votes in a poll",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllPollVoteRequest{
				Poll:       args[0],
				Pagination: pageReq,
			}

			res, err := queryClient.PollVoteAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowVote() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show [vote]",
		Short: "shows a vote",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetVoteRequest{
				Id: args[0],
			}

			res, err := queryClient.Vote(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
