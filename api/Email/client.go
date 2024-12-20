package email

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type OpenObserverOptions struct {
	OrgID      string
	StreamName string
	User       string
	Password   string
	Http_api   string
}

type OpenObserverClient struct {
	client     *http.Client
	orgID      string
	streamName string
	user       string
	password   string
	baseURL    string
}

func NewOpenObserverClient(opts OpenObserverOptions) *OpenObserverClient {
	url := fmt.Sprintf("%s/%s/_search", opts.Http_api, opts.OrgID)

	return &OpenObserverClient{
		client:     &http.Client{},
		orgID:      opts.OrgID,
		streamName: opts.StreamName,
		user:       opts.User,
		password:   opts.Password,
		baseURL:    url,
	}
}

func (ooc OpenObserverClient) SearchOpenObserver(sql string, params Params) ([]Hit, error) {

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"end_time":   params.EndTime,
			"from":       params.From,
			"size":       params.Size,
			"sql":        sql,
			"start_time": params.StartTime,
		},
	}

	body, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", ooc.baseURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(ooc.user, ooc.password)
	resp, err := ooc.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	resbody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var responseOpenObserver OpenObserverResponse

	err = json.Unmarshal(resbody, &responseOpenObserver)
	if err != nil {
		return nil, err
	}

	return responseOpenObserver.Hits, nil
}
