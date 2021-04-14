package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

// GetQueryVoteCmd returns the cli query commands for this module
func GetQueryVoteCmd(queryRoute string) *cobra.Command {
	// Group Themis queries under a subcommand
	cmd := &cobra.Command{
		Use:                        "vote",
		Short:                      "Manages vote queries",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1

	cmd.AddCommand(CmdListVote())
	cmd.AddCommand(CmdShowVote())
	cmd.AddCommand(CmdListVoteWithGroup())
	cmd.AddCommand(CmdListVoteWithPoll())
	cmd.AddCommand(CmdListVoteWithUser())

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

func CmdShowVote() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show [id]",
		Short: "shows a vote",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetVoteRequest{
				Id: id,
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

func CmdListVoteWithGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "group [group] [poll]",
		Short: "list all vote in that group and optionally poll",
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
				params := &types.QueryVoteWithGroup{
					Group:      groupStr,
					Pagination: pageReq,
				}

				res, err := queryClient.VoteWithGroup(context.Background(), params)
				if err != nil {
					return err
				}

				return clientCtx.PrintProto(res)

			} else {
				pollStr, err := strconv.ParseUint(args[1], 10, 64)
				if err != nil {
					return err
				}

				params := &types.QueryVoteWithGroupAndPoll{
					Group:      groupStr,
					Poll:       pollStr,
					Pagination: pageReq,
				}

				res, err := queryClient.VoteWithGroupAndPoll(context.Background(), params)
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

func CmdListVoteWithUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user [user] [poll]",
		Short: "list all vote by that user and optionally poll",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			if len(args) == 1 {
				params := &types.QueryVoteWithUser{
					User:       args[0],
					Pagination: pageReq,
				}

				res, err := queryClient.VoteWithUser(context.Background(), params)
				if err != nil {
					return err
				}

				return clientCtx.PrintProto(res)

			} else {
				pollStr, err := strconv.ParseUint(args[1], 10, 64)
				if err != nil {
					return err
				}

				params := &types.QueryVoteWithUserAndPoll{
					User:       args[0],
					Poll:       pollStr,
					Pagination: pageReq,
				}

				res, err := queryClient.VoteWithUserAndPoll(context.Background(), params)
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

func CmdListVoteWithPoll() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "poll [poll] [vote]",
		Short: "list all vote in that poll and optionally per vote category",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			pollStr, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			if len(args) == 1 {
				params := &types.QueryVoteWithPoll{
					Poll:       pollStr,
					Pagination: pageReq,
				}

				res, err := queryClient.VoteWithPoll(context.Background(), params)
				if err != nil {
					return err
				}

				return clientCtx.PrintProto(res)

			} else {

				params := &types.QueryVoteWithPollAndVote{
					Poll:       pollStr,
					Vote:       args[1],
					Pagination: pageReq,
				}

				res, err := queryClient.VoteWithPollAndVote(context.Background(), params)
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
