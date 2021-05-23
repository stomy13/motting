package internal

import (
	"encoding/json"
	"net/http"

	"github.com/MasatoTokuse/motting/motting/error_response"
	"github.com/MasatoTokuse/motting/motting/infrastracture/persistence/mysql"
)

func Dispatch(err error, response http.ResponseWriter) {
	if err == nil {
		return
	}
	switch err.(type) {
	case *UnmarshalJsonError:
		response.WriteHeader(400)
		encoder := json.NewEncoder(response)
		encoder.Encode(err)
	case *error_response.ValidateErrors:
		response.WriteHeader(400)
		encoder := json.NewEncoder(response)
		encoder.Encode(err)
	case *mysql.DBError:
		response.WriteHeader(500)
		encoder := json.NewEncoder(response)
		encoder.Encode(err)
	default:
		response.WriteHeader(500)
		response.Write([]byte(`{"message":"internal server error."}`))
	}
}
