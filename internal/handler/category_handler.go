package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/rodrigofrumento/usersGo/internal/dto"
	"github.com/rodrigofrumento/usersGo/internal/handler/httperr"
	"github.com/rodrigofrumento/usersGo/internal/handler/validation"
)

// Create category
//
//	@Summary		Create new category
//	@Description	Endpoint for create category
//	@Tags			category
//	@Accept			json
//	@Produce		json
//	@Param			body	body	dto.CreateCategoryDto	true	"Create category dto"	true
//	@Success		200
//	@Failure		400	{object}	httperr.RestErr
//	@Failure		500	{object}	httperr.RestErr
//	@Router			/category [post]
func (h *handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateCategoryDto

	if r.Body == http.NoBody {
		slog.Error("body is empty", slog.String("package", "categoryhandler"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("body is required")
		json.NewEncoder(w).Encode(msg)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("error to decode body", "err", err, slog.String("package", "categoryhandler"))
		w.WriteHeader(http.StatusBadRequest)
		msg := httperr.NewBadRequestError("error to decode body")
		json.NewEncoder(w).Encode(msg)
		return
	}
	httpErr := validation.ValidateHttpData(req)
	if httpErr != nil {
		slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "categoryhandler"))
		w.WriteHeader(httpErr.Code)
		json.NewEncoder(w).Encode(httpErr)
		return
	}
	err = h.categoryService.CreateCategory(r.Context(), req)
	if err != nil {
		slog.Error(fmt.Sprintf("error to create category: %v", err), slog.String("package", "categoryhandler"))
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
}
