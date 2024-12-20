package email

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
)

// type EmailHandlerImpl interface {
// 	GetEmailById(w http.ResponseWriter, r *http.Request)
// }

type EmailHandler struct {
	openObserverClient *OpenObserverClient
}

func NewEmailHandler(ooc OpenObserverClient) EmailHandler {
	return EmailHandler{
		openObserverClient: &ooc,
	}
}

type Response struct {
	Count   int         `json:"count"`
	Message string      `json:"message,omitempty"`
	Results interface{} `json:"results,omitempty"`
}

type Params struct {
	StartTime int    `json:"start_time,omitempty"`
	EndTime   int    `json:"end_time,omitempty"`
	From      int    `json:"from,omitempty"`
	Size      int    `json:"size,omitempty"`
	Subject   string `json:"subject,omitempty"`
	To        string `json:"to,omitempty"`
	FromEmail string `json:"from_email,omitempty"`
}

func (e EmailHandler) GetEmailById(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&Response{
			Message: "id required",
		})
		return
	}

	queryParams := r.URL.Query()
	streamLog := queryParams.Get("stream_log")
	startTimeStr := queryParams.Get("start_time")
	endTimeStr := queryParams.Get("end_time")

	if streamLog == "" || endTimeStr == "" || startTimeStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&Response{
			Message: "params required",
		})
		return
	}

	startTime, err := strconv.Atoi(startTimeStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&Response{
			Message: "error parser field",
		})
		return
	}

	endTime, err := strconv.Atoi(endTimeStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&Response{
			Message: "error parser field",
		})
		return
	}

	params := Params{
		From:      0,
		Size:      1,
		StartTime: startTime,
		EndTime:   endTime,
	}

	sql := fmt.Sprintf(`
	SELECT to, body, subject, date, cc, %s.from, message_id 
	FROM %s 
	WHERE message_id = '%s'
	`, streamLog, streamLog, id)

	emails, err := e.openObserverClient.SearchOpenObserver(sql, params)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&Response{
			Message: "error service",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&Response{
		Message: "success",
		Results: emails[0],
	})
}

func (e EmailHandler) GetAllEmails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	queryParams := r.URL.Query()
	streamLog := queryParams.Get("stream_log")
	startTimeStr := queryParams.Get("start_time")
	endTimeStr := queryParams.Get("end_time")
	sizeStr := queryParams.Get("size")
	fromStr := queryParams.Get("from")
	subject := queryParams.Get("subject")
	to := queryParams.Get("to")
	fromEmail := queryParams.Get("fromEmail")

	if streamLog == "" || endTimeStr == "" || startTimeStr == "" || sizeStr == "" || fromStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&Response{
			Message: "params required",
		})
		return
	}

	from, err := parseField(fromStr, w)
	if err != nil {
		return
	}

	size, err := parseField(sizeStr, w)
	if err != nil {
		return
	}

	startTime, err := parseField(startTimeStr, w)
	if err != nil {
		return
	}

	endTime, err := parseField(endTimeStr, w)
	if err != nil {
		return
	}

	params := Params{
		From:      from,
		Size:      size,
		StartTime: startTime,
		EndTime:   endTime,
	}

	sql := fmt.Sprintf(`
	SELECT to, body, subject, date, cc, %s.from, message_id 
	FROM %s 
	`, streamLog, streamLog)

	countSql := fmt.Sprintf(`SELECT COUNT(*) AS count FROM %s`, streamLog)

	conditions := []string{}

	if subject != "" {
		conditions = append(conditions, fmt.Sprintf("str_match_ignore_case(subject, '%s')", subject))
	}

	if to != "" {
		conditions = append(conditions, fmt.Sprintf("str_match_ignore_case(to, '%s')", to))
	}

	if fromEmail != "" {
		conditions = append(conditions, fmt.Sprintf("str_match_ignore_case(from, '%s')", fromEmail))
	}

	if len(conditions) > 0 {
		sql += " WHERE " + strings.Join(conditions, " AND ")
		countSql += " WHERE " + strings.Join(conditions, " AND ")
	}

	emails, err := e.openObserverClient.SearchOpenObserver(sql, params)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&Response{
			Message: "error service",
		})
		return
	}

	params.From = 0

	countHit, err := e.openObserverClient.SearchOpenObserver(countSql, params)

	fmt.Println("COUNT", countHit[0].Count)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&Response{
			Message: "error service",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&Response{
		Count:   countHit[0].Count,
		Results: emails,
	})
}

func parseField(field string, w http.ResponseWriter) (int, error) {
	fieldParser, err := strconv.Atoi(field)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&Response{
			Message: fmt.Sprintf("Invalid value for field: %s", field),
		})
		return 0, fmt.Errorf("invalid value for field: %s", field)
	}

	return fieldParser, nil
}
