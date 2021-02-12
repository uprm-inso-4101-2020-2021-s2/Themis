package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers Themis-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	// this line is used by starport scaffolding # 1
	r.HandleFunc("/Themis/poll", createPollHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/Themis/poll", extendPollDeadlineHandler(cliCtx)).Methods("PUT")
	r.HandleFunc("/Themis/poll", setPollDescHandler(cliCtx)).Methods("PUT")
	r.HandleFunc("/Themis/poll", listPollHandler(cliCtx, "Themis")).Methods("GET")
	r.HandleFunc("/Themis/poll", listGroupPollsHandler(cliCtx, "Themis")).Methods("GET")
	r.HandleFunc("/Themis/poll/{key}", getPollHandler(cliCtx, "Themis")).Methods("GET")

	r.HandleFunc("/Themis/account", createAccountHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/Themis/account", listAccountHandler(cliCtx, "Themis")).Methods("GET")
	r.HandleFunc("/THemis/account", listUserAccountsHandler(cliCtx, "Themis")).Methods("GET")
	r.HandleFunc("/THemis/account", listGroupAccountsHandler(cliCtx, "Themis")).Methods("GET")
	r.HandleFunc("/Themis/account/{key}", getAccountHandler(cliCtx, "Themis")).Methods("GET")

	r.HandleFunc("/Themis/group", createGroupHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/Themis/group", setGroupNameHandler(cliCtx)).Methods("PUT")
	r.HandleFunc("/Themis/group", listGroupHandler(cliCtx, "Themis")).Methods("GET")
	r.HandleFunc("/Themis/group/{key}", getGroupHandler(cliCtx, "Themis")).Methods("GET")
	//r.HandleFunc("/Themis/group", setGroupHandler(cliCtx)).Methods("PUT")
	//r.HandleFunc("/Themis/group", deleteGroupHandler(cliCtx)).Methods("DELETE")

}
