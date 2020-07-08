package handlers

import (
	"car-images-api/files"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
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
	f.saveFile(id, fileName, w, r.Body)
}

// This func will ne responsable for handling multipart request
func (f *Files) HandlePostMultipart(w http.ResponseWriter, r *http.Request) {
	// set the max size to be allocated in memory, the rest will be dump in temp files
	if err := r.ParseMultipartForm(128 * 1024); err != nil {
		// if something went wrong, inform the user and log the error
		f.log.Error("Bad request", "error", err)
		http.Error(w, "Expected multi-part form data", http.StatusBadRequest)
		return
	}
	id, idErr := strconv.Atoi(r.FormValue("id"))
	if idErr != nil {
		f.log.Error("Bad request - id is not a valid integer", "Error", idErr)
		http.Error(w, "Id is not a valid integer", http.StatusBadRequest)
		return
	}

	f.log.Debug("Form value", "id", id)
	file, mfh, err := r.FormFile("file")
	if err != nil {
		f.log.Error("Bad request - reading file multi-part", "Error", err)
		http.Error(w, "Expected multi-part form data", http.StatusBadRequest)
		return
	}

	f.saveFile(r.FormValue("id"), mfh.Filename, w, file)
}

func (f *Files) invalidURI(uri string, rw http.ResponseWriter) {
	f.log.Error("Invalid PATH", "path", uri)
	http.Error(rw, "Invalid file path. Should be in the format: /[id]/[filepaht]", http.StatusBadRequest)
}

func (f *Files) saveFile(id, path string, rw http.ResponseWriter, r io.ReadCloser  ) {
	f.log.Info("Save file for car", "id", id, "path", path)

	fp := filepath.Join(id, path)
	err := f.store.Save(fp, r)
	if err != nil {
		f.log.Error("Unable to save file", "error", err)
		http.Error(rw, "Unable to save file", http.StatusInternalServerError)
	}
}
