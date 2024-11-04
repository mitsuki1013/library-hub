package action

import (
	"encoding/json"
	"io"
	"library-hub/app/adapter/api/schema"
	"library-hub/app/usecase"
	"net/http"
)

type CreateCompanyAction struct {
	uc usecase.CreateCompanyUseCase
}

func NewCreateCompanyAction(uc usecase.CreateCompanyUseCase) CreateCompanyAction {
	return CreateCompanyAction{
		uc: uc,
	}
}

func (a CreateCompanyAction) CreateCompany(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var requestData schema.Company
	if err = json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, "Error parsing request json", http.StatusBadRequest)
		return
	}

	output, err := a.uc.Execute(usecase.CreateCompanyInput{
		ID:   requestData.ID,
		Name: requestData.Name,
	})
	if err != nil {
		http.Error(w, "Error executing use case", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(schema.Company{
		ID:   output.ID,
		Name: output.Name,
	}); err != nil {
		http.Error(w, "Error encoding response json", http.StatusInternalServerError)
		return
	}
}
