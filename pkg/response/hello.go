package response

import "net/http"

var HelloWorld = StructResponse{
	HttpStatus: http.StatusOK,
	StateCode:  "BM_REQUEST_SF_200",
	MessageEn:  "Backend Management request found.",
	MessageBn:  "Backend Management request found.",
}
