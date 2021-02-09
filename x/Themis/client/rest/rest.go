package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers Themis-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	// this line is used by starport scaffolding # 1
	r.HandleFunc("/Themis/voucher", createVoucherHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/Themis/voucher", listVoucherHandler(cliCtx, "Themis")).Methods("GET")
	r.HandleFunc("/Themis/voucher/{key}", getVoucherHandler(cliCtx, "Themis")).Methods("GET")

	r.HandleFunc("/Themis/group", createGroupHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/Themis/group", listGroupHandler(cliCtx, "Themis")).Methods("GET")
	r.HandleFunc("/Themis/group/{key}", getGroupHandler(cliCtx, "Themis")).Methods("GET")
	//r.HandleFunc("/Themis/group", setGroupHandler(cliCtx)).Methods("PUT")
	//r.HandleFunc("/Themis/group", deleteGroupHandler(cliCtx)).Methods("DELETE")

}
