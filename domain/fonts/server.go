package fonts

import (
	"github.com/chazari-x/ningyotsukai/config"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/chazari-x/ningyotsukai/domain/fonts/handler"
	"github.com/go-chi/chi/v5"
)

func StartServer(cfg *config.Fonts) error {
	h := handler.NewHandler(cfg)
	router := chi.NewRouter()
	router.Post("/status", h.Status)
	router.Post("/font/{font}", h.Font)
	router.Post("/fonts", h.Fonts)

	log.Trace("fonts start")
	return http.ListenAndServe(cfg.Host, router)
}
