package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/uprm-inso-4101-2020-2021-s2/Themis/x/Themis/types"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type createVoucherRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string       `json:"creator"`
	Group   string       `json:"group"`
	Owner   string       `json:"owner"`
	Amount  int          `json:"amount" yaml:"amount"`
}

func createAccountHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createVoucherRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		creator, err := sdk.AccAddressFromBech32(req.Creator)
		owner, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		parsedGroup := req.Group
		parsedAmount := req.Amount

		msg := types.NewMsgAccountAddVotes(
			creator,
			owner,
			parsedGroup,
			parsedAmount,
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
