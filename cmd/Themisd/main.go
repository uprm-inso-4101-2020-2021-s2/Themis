package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/app"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/cmd/Themisd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
