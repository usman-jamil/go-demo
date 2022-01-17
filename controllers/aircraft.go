package controllers

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"enttest/ent"
	"enttest/models"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

type aircraftController struct {
	aircraftPagedPattern        *regexp.Regexp
	aircraftIDPattern           *regexp.Regexp
	aircraftOrderByPattern      *regexp.Regexp
	aircraftTypePattern         *regexp.Regexp
	aircraftRegistrationPattern *regexp.Regexp
	aircraftExportCSVPattern    *regexp.Regexp
	ctx                         context.Context
	client                      *ent.Client
}

func (uc aircraftController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/aircrafts" && r.URL.RawQuery == "" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else if matches := uc.aircraftExportCSVPattern.FindStringSubmatch(r.URL.Path); len(matches) > 0 {
		switch r.Method {
		case http.MethodGet:
			w.Header().Set("Content-Type", "text/csv")
			w.Header().Set("Content-Disposition", "attachment;filename=TheCSVFileName.csv")
			w.Header().Set("Transfer-Encoding", "chunked")
			wr := csv.NewWriter(w)
			aircrafts, _ := models.GetAircrafts(uc.ctx, uc.client)

			wr.Write([]string{
				"Aircraft Id",
				"Company Id",
				"Current Flight Hours",
				"Current Cycles",
				"Aircraft Registration",
				"Base Airport Code",
				"Manufacturer",
				"Manufacturer Designator",
				"Common Designation",
				"Common Name",
				"Pilots Required To Fly)",
				"Default Values",
				"Maximum Values",
				"Current Landings)",
				"Fuel Details",
				"Oil Details",
			})

			for _, value := range aircrafts {
				wr.Write([]string{
					(*value).AircraftId.String(),
					(*value).CompanyId.String(),
					fmt.Sprintf("%f", (*value).CurrentFlightHours),
					strconv.Itoa((*value).CurrentCycles),
					(*value).AircraftRegistration,
					(*value).BaseAirportCode,
					(*value).Manufacturer,
					(*value).ManufacturerDesignator,
					(*value).CommonDesignation,
					(*value).CommonName,
					strconv.Itoa((*value).PilotsRequiredToFly),
					(*value).DefaultValues,
					(*value).MaximumValues,
					strconv.Itoa((*value).CurrentLandings),
					(*value).FuelDetails,
					(*value).OilDetails,
				})
			}

			wr.Flush()
			//uc.getAllCsv(*wr, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else if matches := uc.aircraftIDPattern.FindStringSubmatch(r.URL.Path); len(matches) > 0 {
		switch r.Method {
		case http.MethodGet:
			uc.get(matches[1], w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else if matches := uc.aircraftPagedPattern.FindStringSubmatch(strings.ToLower(r.URL.RawQuery)); r.URL.Path == "/aircrafts" && len(matches) > 0 {
		switch r.Method {
		case http.MethodGet:
			values := r.URL.Query()
			limit := getQueryStringValue(values, "limit")
			offset := getQueryStringValue(values, "offset")
			iLimit, err := strconv.Atoi(limit)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("limit should be a valid integer"))
				return
			}

			iOffset, err := strconv.Atoi(offset)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("offset should be a valid integer"))
				return
			}

			uc.getAllPaged(w, r, iLimit, iOffset)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else if matches := uc.aircraftOrderByPattern.FindStringSubmatch(strings.ToLower(r.URL.RawQuery)); r.URL.Path == "/aircrafts" && len(matches) > 0 {
		switch r.Method {
		case http.MethodGet:
			values := r.URL.Query()
			orderBy := getQueryStringValue(values, "orderby")
			direction := getQueryStringValue(values, "direction")

			isAscending := (direction == "" || direction == "asc") || !(direction != "" && direction == "desc")

			if orderBy == "flighthours" {
				if isAscending {
					uc.getAllSortByFlightHoursAscending(w, r)
				} else {
					uc.getAllSortByFlightHoursDescending(w, r)
				}
			} else {
				w.WriteHeader(http.StatusNotImplemented)
			}
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else if matches := uc.aircraftTypePattern.FindStringSubmatch(strings.ToLower(r.URL.RawQuery)); r.URL.Path == "/aircrafts" && len(matches) > 0 {
		switch r.Method {
		case http.MethodGet:
			values := r.URL.Query()
			designation := getQueryStringValue(values, "designation")

			uc.getAllByCommonDesignation(w, r, designation)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else if matches := uc.aircraftRegistrationPattern.FindStringSubmatch(strings.ToLower(r.URL.RawQuery)); r.URL.Path == "/aircrafts" && len(matches) > 0 {
		switch r.Method {
		case http.MethodGet:
			values := r.URL.Query()
			registration := getQueryStringValue(values, "registration")

			uc.getAllByCommonRegistration(w, r, registration)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func getQueryStringValue(values url.Values, key string) string {
	var value string

	for k, _ := range values {
		if strings.ToLower(k) == key {
			value = strings.ToLower(values.Get(k))
		}
	}

	return value
}

func (uc *aircraftController) getAll(w http.ResponseWriter, r *http.Request) {
	u, _ := models.GetAircrafts(uc.ctx, uc.client)

	encodeResponseAsJSON(u, w)
}

func (uc *aircraftController) getAllPaged(w http.ResponseWriter, r *http.Request, limit int, offset int) {
	u, _ := models.GetAircraftsPaged(uc.ctx, uc.client, limit, offset)

	encodeResponseAsJSON(u, w)
}

func (uc *aircraftController) getAllSortByFlightHoursAscending(w http.ResponseWriter, r *http.Request) {
	u, _ := models.GetAircraftsByFlightHoursAscending(uc.ctx, uc.client)

	encodeResponseAsJSON(u, w)
}

func (uc *aircraftController) getAllSortByFlightHoursDescending(w http.ResponseWriter, r *http.Request) {
	u, _ := models.GetAircraftsByFlightHoursDescending(uc.ctx, uc.client)

	encodeResponseAsJSON(u, w)
}

func (uc *aircraftController) getAllByCommonDesignation(w http.ResponseWriter, r *http.Request, designation string) {
	u, _ := models.GetAircraftsByType(uc.ctx, uc.client, designation)

	encodeResponseAsJSON(u, w)
}

func (uc *aircraftController) getAllByCommonRegistration(w http.ResponseWriter, r *http.Request, registration string) {
	u, _ := models.GetAircraftsByRegistration(uc.ctx, uc.client, registration)

	encodeResponseAsJSON(u, w)
}

func stringToUUID(newUUID string) (*uuid.UUID, error) {
	id, err := uuid.Parse(newUUID)

	if err != nil {
		return nil, fmt.Errorf("failed to convert string to []byte: %w", err)
	}

	return &id, nil
}

func (uc *aircraftController) get(id string, w http.ResponseWriter) {

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

func (uc *aircraftController) parseRequest(r *http.Request) (models.Aircraft, error) {
	dec := json.NewDecoder(r.Body)
	var u models.Aircraft
	err := dec.Decode(&u)
	if err != nil {
		return models.Aircraft{}, err
	}
	return u, nil
}

func newAircraftController(dbContext context.Context, dbClient *ent.Client) *aircraftController {
	return &aircraftController{
		aircraftPagedPattern:        regexp.MustCompile(`^limit=\d+&offset=\d+/?`),
		aircraftIDPattern:           regexp.MustCompile(`^/aircrafts/(\w{8}-\w{4}-\w{4}-\w{4}-\w{12})/?`),
		aircraftOrderByPattern:      regexp.MustCompile(`^orderby=\w+(&direction=(asc|desc))?/?`),
		aircraftTypePattern:         regexp.MustCompile(`^designation=[\w\d-]*/?`),
		aircraftRegistrationPattern: regexp.MustCompile(`^registration=[\w\d-]*/?`),
		aircraftExportCSVPattern:    regexp.MustCompile(`^/aircrafts/exportcsv/?`),
		ctx:                         dbContext,
		client:                      dbClient,
	}
}
