package handlers

import (
	"car-images-api/files"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"net/http"
	"path/filepath"
)

type Files struct {
	log   hclog.Logger
	store files.Storage
}

func NewFiles(s files.Storage, l hclog.Logger) *Files {
	return &Files{l, s}
}

func (f *Files) HandlePostFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fileName := vars["filename"]

	f.log.Info("handle POST file", "id", id, "filename", fileName)
	f.saveFile(id, fileName, w, r)
}

func (f *Files) invalidURI(uri string, rw http.ResponseWriter) {
	f.log.Error("Invalid PATH", "path", uri)
	http.Error(rw, "Invalid file path. Should be in the format: /[id]/[filepaht]", http.StatusBadRequest)
}

func (f *Files) saveFile(id, path string, rw http.ResponseWriter, r *http.Request) {
	f.log.Info("Save file for car", "id", id, "path", path)

	fp := filepath.Join(id, path)
	err := f.store.Save(fp, r.Body)
	if err != nil {
		f.log.Error("Unable to save file", "error", err)
		http.Error(rw, "Unable to save file", http.StatusInternalServerError)
	}
}
