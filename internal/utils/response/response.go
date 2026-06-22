package response
import (
	"encoding/json"
	"net/http"
	"strings"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Response struct{
	Status string `json:"status"`
	Error string `json:"error"`
}
const (
	StatusOK = "OK"
	StatusError = "ERROR"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response{
	return Response{Status: StatusError, Error: err.Error()}
}

func GeneralSuccess(msg string) Response{
	return Response{Status: StatusOK, Error: msg}
}

func ValidationError(errs validator.ValidationErrors) Response{
	var errMsgs []string
	for _, err := range errs {
		switch err.ActualTag(){
			case "required":
				errMsgs = append(errMsgs, fmt.Sprintf("%s is required", err.Field()))
			case "email":
				errMsgs = append(errMsgs, fmt.Sprintf("%s must be an email", err.Field()))
			case "numeric":
				errMsgs = append(errMsgs, fmt.Sprintf("%s must be a number", err.Field()))
			default:
				errMsgs = append(errMsgs, fmt.Sprintf("%s is invalid", err.Field()))
		}
	}
	return Response{Status: StatusError, Error: strings.Join(errMsgs, ", ")}
}