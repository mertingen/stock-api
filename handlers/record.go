package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mertingen/stock-api/models"
	"github.com/mertingen/stock-api/services"
	"go.mongodb.org/mongo-driver/bson"
)

type Record struct {
	recordService services.Record
}

func InitRecord(recordService services.Record) Record {
	return Record{recordService: recordService}
}

type recordsReq struct {
	Startdate string `json:"startDate" validate:"required"`
	EndDate   string `json:"endDate" validate:"required"`
	MinCount  int16  `json:"minCount" validate:"required"`
	MaxCount  int16  `json:"maxCount" validate:"required"`
}

type recordsResp struct {
	Code    int8            `json:"code"`
	Msg     string          `json:"msg"`
	Records []models.Record `json:"records"`
}

func (rec *Record) FetchAll(w http.ResponseWriter, r *http.Request) {
	// Set response content type and response struct
	w.Header().Add("Content-Type", "application/json")
	resp := recordsResp{}

	// Accept only HTTP POST verb
	switch r.Method {
	case "POST":
		// Accept only 'application/json' conent-type
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			resp = recordsResp{
				Code:    -2,
				Msg:     "Error: Content type is not valid",
				Records: []models.Record{},
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resp)
			return
		}

		// Set request body into recordsReq struct
		recordsReq := recordsReq{}
		body, _ := io.ReadAll(r.Body)

		err := json.Unmarshal(body, &recordsReq)
		if err != nil {
			resp = recordsResp{
				Code:    -3,
				Msg:     fmt.Sprintf("Error: %s", err.Error()),
				Records: []models.Record{},
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resp)
			return
		}

		// Validate the HTTP Body fields according to the recordsReq struct
		validate := validator.New()
		err = validate.Struct(recordsReq)
		if err != nil {
			resp = recordsResp{
				Code:    -3,
				Msg:     fmt.Sprintf("Error: %s", err.Error()),
				Records: []models.Record{},
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resp)
			return
		}

		// Checks date formats for YYYY-MM-DD
		if !rec.recordService.ValidateDate(recordsReq.Startdate) || !rec.recordService.ValidateDate(recordsReq.EndDate) {
			resp = recordsResp{
				Code:    -4,
				Msg:     "Error: Date formats are invalid",
				Records: []models.Record{},
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resp)
			return
		}

		// Set the filter to query DB side
		filter := bson.M{
			"createdAt": bson.M{
				"$gte": recordsReq.Startdate,
				"$lte": recordsReq.EndDate,
			},
			"totalCount": bson.M{
				"$gte": recordsReq.MinCount,
				"$lte": recordsReq.MaxCount,
			},
		}
		records, err := rec.recordService.FetchAll(filter)
		if err != nil {
			resp = recordsResp{
				Code:    -1,
				Msg:     fmt.Sprintf("Error: %s", err.Error()),
				Records: []models.Record{},
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resp)
			return
		}

		resp = recordsResp{
			Code:    0,
			Msg:     "Success",
			Records: records,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
		return

	default:
		resp = recordsResp{
			Code:    -10,
			Msg:     "Error: HTTP verb is invalid",
			Records: []models.Record{},
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
		return
	}
}
