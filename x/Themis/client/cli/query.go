package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group Themis queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1

	cmd.AddCommand(CmdListPoll())
	cmd.AddCommand(CmdListGroupPoll())
	cmd.AddCommand(CmdShowPoll())

	cmd.AddCommand(CmdListAccount())
	cmd.AddCommand(CmdListUserAccount())
	cmd.AddCommand(CmdListGroupAccount())
	cmd.AddCommand(CmdShowAccount())

	cmd.AddCommand(CmdListGroup())
	cmd.AddCommand(CmdShowGroup())

	return cmd
}
