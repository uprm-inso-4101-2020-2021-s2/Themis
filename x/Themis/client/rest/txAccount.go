package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type setAccountVoucherRequest struct {
	BaseReq  rest.BaseReq `json:"base_req"`
	Creator  string       `json:"creator"`
	User     string       `json:"user"`
	Group    string       `json:"group"`
	Vouchers string       `json:"vouchers"`
}

func createSetAccountVoucherHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req setAccountVoucherRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		parsedGroup := req.Group
		parsedUser := req.User
		parsedVouchers, _ := strconv.ParseInt(req.Vouchers, 10, 64)

		msg := types.NewMsgAddAccountVouchers(
			req.Creator,
			parsedUser,
			parsedGroup,
			parsedVouchers,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
