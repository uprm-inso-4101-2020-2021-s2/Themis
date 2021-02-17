package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	// this line is used by starport scaffolding # 1
)

const (
	MethodGet = "GET"
)

// RegisterRoutes registers Themis-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
	r.HandleFunc("/Themis/votes/{id}", getVoteHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/Themis/votes/{user}", listUserVoteHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/Themis/votes/{poll}", listPollVoteHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/Themis/votes", listVoteHandler(clientCtx)).Methods("GET")

	r.HandleFunc("/Themis/polls/{id}", getPollHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/Themis/polls", listPollHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/Themis/polls/{group}", listGroupPollHandler(clientCtx)).Methods("GET")

	r.HandleFunc("/Themis/accounts/{id}", getAccountHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/Themis/accounts/{group}", listGroupAccountsHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/Themis/accounts/{user}", listUserAccountsHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/Themis/accounts", listAccountHandler(clientCtx)).Methods("GET")

	r.HandleFunc("/Themis/groups/{id}", getGroupHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/Themis/groups", listGroupHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
	r.HandleFunc("/Themis/votes", createVoteHandler(clientCtx)).Methods("POST")

	r.HandleFunc("/Themis/polls", createPollHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/Themis/polls", setPollDescHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/Themis/polls", extendPollDeadlineHandler(clientCtx)).Methods("POST")

	r.HandleFunc("/Themis/accounts", createSetAccountVoucherHandler(clientCtx)).Methods("POST")

	r.HandleFunc("/Themis/groups", createGroupHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/Themis/groups", setGroupNameHandler(clientCtx)).Methods("POST")

}
