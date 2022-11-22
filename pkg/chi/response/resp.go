package response

import (
	"encoding/json"
	"net/http"
)

// JSONResponse
type JSONResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

const MarshalErr = `{"code":500","msg":"parsing error"}`

type H map[string]interface{}

// JSON
func JSON(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write(NewJSONResponse(nil))
}

// JSONData
func JSONData(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Write(NewJSONResponse(data))
}

// JSONList
func JSONList(w http.ResponseWriter, list interface{}, total int64) {
	w.WriteHeader(http.StatusOK)
	w.Write(NewJSONResponse(H{"list": list, "total": total}))

}

func JSONBadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write(NewJSONErrResponse(err))
}

func JSONServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(NewJSONErrResponse(err))
}

func NewJSONResponse(data interface{}) []byte {
	response := JSONResponse{Code: 0, Msg: "ok", Data: data}
	resp, err := json.Marshal(response)
	if err != nil {
		return []byte(MarshalErr)
	} else {
		return resp
	}
}

func NewJSONErrResponse(err error) []byte {
	response := JSONResponse{Code: 100, Msg: err.Error()}
	resp, err := json.Marshal(response)
	if err != nil {
		return []byte(MarshalErr)
	} else {
		return resp
	}
}
