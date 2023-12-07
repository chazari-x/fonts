package handler

import (
	"encoding/json"
	"fmt"
	"github.com/chazari-x/ningyotsukai/config"
	"github.com/chazari-x/ningyotsukai/domain/fonts/model"
	"github.com/go-chi/chi/v5"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type Controller struct {
	cfg *config.Fonts
}

func NewHandler(cfg *config.Fonts) *Controller {
	return &Controller{cfg: cfg}
}

func (c *Controller) Status(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	files, err := ioutil.ReadDir("fonts")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	marshal, err := json.Marshal(model.Status{
		Num: len(files),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(marshal); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (c *Controller) Font(w http.ResponseWriter, r *http.Request) {
	fileName := chi.URLParam(r, "font")

	file, err := os.Open(fmt.Sprintf("fonts/%s.ttf", fileName))
	if err != nil {
		file, err = os.Open(fmt.Sprintf("fonts/%s.TTF", fileName))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}

	defer func() {
		_ = file.Close()
	}()

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.ttf", fileName))

	_, err = io.Copy(w, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (c *Controller) Fonts(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	files, err := ioutil.ReadDir("fonts")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var Files []model.File
	for _, file := range files {
		Files = append(Files, model.File{Name: file.Name()[:len(file.Name())-4]})
	}

	marshal, err := json.Marshal(Files)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err = w.Write(marshal); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
