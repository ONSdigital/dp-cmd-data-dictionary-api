package api

import (
	"encoding/json"
	"net/http"

	"github.com/ONSdigital/dp-code-list-api/models"
	"github.com/ONSdigital/log.go/log"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func (c *CodeListAPI) getCodes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	edition := vars["edition"]
	data := log.Data{"codelist_id": id, "edition": edition}

	log.Event(ctx, "getCodes endpoint: attempting to get edition codes", log.INFO, data)

	dbCodes, err := c.store.GetCodes(ctx, id, edition)
	if err != nil {
		log.Event(ctx, "error getting codes", log.ERROR, log.Error(errors.WithMessage(err, "getCodes endpoint: store.GetCode returned an error")), data)
		handleError(ctx, w, err, data)
		return
	}
	codes := models.NewCodeResults(dbCodes)

	for _, item := range codes.Items {
		if err := item.UpdateLinks(c.apiURL, id, edition); err != nil {
			log.Event(ctx, "error updating links", log.ERROR, log.Error(errors.WithMessage(err, "getCodes endpoint: links could not be created")))
			http.Error(w, internalServerErr, http.StatusInternalServerError)
			return
		}
	}

	count := len(codes.Items)
	codes.Count = count
	codes.Limit = count
	codes.TotalCount = count

	b, err := json.Marshal(codes)
	if err != nil {
		log.Event(ctx, "marshal error", log.ERROR, log.Error(errors.WithMessage(err, "getCodes endpoint: failed to marshal codes to json bytes")), data)
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}

	if err := c.writeBody(w, b); err != nil {
		log.Event(ctx, "error writting body", log.ERROR, log.Error(errors.WithMessage(err, "getCodes endpoint: failed to write bytes to response")))
		return
	}

	log.Event(ctx, "getCodes endpoint: request successful", log.INFO, data)
}

func (c *CodeListAPI) getCode(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	edition := vars["edition"]
	code := vars["code"]
	data := log.Data{"codelist_id": id, "edition": edition, "code": code}

	log.Event(ctx, "getCode: attempting to get code list code", log.INFO, data)

	dbCode, err := c.store.GetCode(ctx, id, edition, code)
	if err != nil {
		log.Event(ctx, "error getting code", log.ERROR, log.Error(errors.WithMessage(err, "getCode endpoint: store.GetCode returned an error")), data)
		handleError(ctx, w, err, data)
		return
	}
	apiCode := models.NewCode(dbCode)

	if err := apiCode.UpdateLinks(c.apiURL, id, edition); err != nil {
		log.Event(ctx, "error updating links", log.ERROR, log.Error(errors.WithMessage(err, "getCode endpoint: links could not be created")))
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(apiCode)
	if err != nil {
		log.Event(ctx, "marshal error", log.ERROR, log.Error(errors.WithMessage(err, "getCode endpoint: error attempting to marshal result to JSON")), data)
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}

	if err := c.writeBody(w, b); err != nil {
		log.Event(ctx, "error writting body", log.ERROR, log.Error(errors.WithMessage(err, "getCode endpoint: failed to write bytes to response")))
		return
	}

	log.Event(ctx, "getCode endpoint: request successful", log.INFO, data)
}
