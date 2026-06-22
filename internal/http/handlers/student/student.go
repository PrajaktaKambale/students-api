package student

import (
	"log/slog"
	"net/http"
	"github.com/PrajaktaKambale/students-api/internal/types"
	"github.com/PrajaktaKambale/students-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {

		slog.Info("student handler called")
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student
	err :=	json.NewDecoder(r.Body).Decode(&student)
if errors.Is(err,io.EOF){
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("request body empty")))
			return
}
if err != nil{
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
			return
}

//request validation
validate:=validator.New()
if err := validate.Struct(student); err != nil {
	validateErrs := err.(validator.ValidationErrors)
    response.WriteJSON(w, http.StatusBadRequest, response.ValidationError(validateErrs))
    return
}


		
		response.WriteJSON(w, http.StatusCreated, response.GeneralSuccess("student created successfully"))
	}
}
