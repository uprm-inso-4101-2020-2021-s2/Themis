package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

func GetCmdListVoucher(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-voucher",
		Short: "list all voucher",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListVoucher, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not list Voucher\n%s\n", err.Error())
				return nil
			}
			var out []types.Voucher
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdGetVoucher(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-voucher [key]",
		Short: "Query a voucher by key",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			key := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryGetVoucher, key), nil)
			if err != nil {
				fmt.Printf("could not resolve voucher %s \n%s\n", key, err.Error())

				return nil
			}

			var out types.Voucher
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdListUserVouchers(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-user-voucher [group] [user]",
		Short: "List all vouchers associated to user",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			group := args[0]
			user := args[1]

			// TODO: cleaner way to make a query instead of single variable methods
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryListUserVoucher, user+"-"+group), nil)
			if err != nil {
				fmt.Printf("could not resolve voucher with group id %s and user %s \n%s\n", group, user, err.Error())

				return nil
			}

			var out types.Voucher
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
