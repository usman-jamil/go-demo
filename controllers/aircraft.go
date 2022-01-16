package controllers

import (
	"context"
	"encoding/json"
	"enttest/ent"
	"enttest/models"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
	ctx           context.Context
	client        *ent.Client
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/aircrafts" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w, r)
		//case http.MethodPost:
		//uc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		//id, err := strconv.Atoi(matches[1])
		//if err != nil {
		//w.WriteHeader(http.StatusNotFound)
		//return
		//}
		switch r.Method {
		case http.MethodGet:
			uc.get(matches[1], w)
			//case http.MethodPut:
			//uc.put(id, w, r)
			//case http.MethodDelete:
			//uc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (uc *userController) getAll(w http.ResponseWriter, r *http.Request) {
	u, _ := models.GetAircrafts(uc.ctx, uc.client)

	encodeResponseAsJSON(u, w)
}

func stringToUUID(newUUID string) (*uuid.UUID, error) {
	id, err := uuid.Parse(newUUID)

	if err != nil {
		return nil, fmt.Errorf("failed to convert string to []byte: %w", err)
	}

	return &id, nil
}

func (uc *userController) get(id string, w http.ResponseWriter) {

	uuid, err := stringToUUID(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Invalid UUID"))
	}

	u, err := models.GetAircraftByID(uc.ctx, uc.client, *uuid)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	encodeResponseAsJSON(u, w)
}

/*
func (uc *userController) post(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}
	u, err = models.AddUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *userController) put(id int, w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}
	if id != u.ID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of submitted user must match ID in URL"))
		return
	}
	u, err = models.UpdateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *userController) delete(id int, w http.ResponseWriter) {
	err := models.RemoveUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}*/

func (uc *userController) parseRequest(r *http.Request) (models.Aircraft, error) {
	dec := json.NewDecoder(r.Body)
	var u models.Aircraft
	err := dec.Decode(&u)
	if err != nil {
		return models.Aircraft{}, err
	}
	return u, nil
}

func newUserController(dbContext context.Context, dbClient *ent.Client) *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/aircrafts/(\w{8}-\w{4}-\w{4}-\w{4}-\w{12})/?`),
		ctx:           dbContext,
		client:        dbClient,
	}
}
