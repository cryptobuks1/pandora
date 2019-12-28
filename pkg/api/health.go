package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HealthController struct{}

func NewHealthController() *HealthController { return &HealthController{} }

func (*HealthController) LiveH(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}
