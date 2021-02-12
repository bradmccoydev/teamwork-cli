package teamwork

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type CreateTimeEntryOps struct {
	// Description of time entry
	Description string `json:"description"`
	// ID of the user for the entry
	PersonID string `json:"person-id"`
	// Start Date of the time entry in YYYYMMDD format
	Date string `json:"date"`
	// Start Time of the time entry in HH:MM:SS format
	Time string `json:"time"`
	// Hours logged for the time entry
	Hours string `json:"hours"`
	// Minutes logged for the time entry
	Minutes string `json:"minutes"`
	// billable flag
	// Valid Input: true, false
	IsBillable string `json:"isbillable,omitempty"`
	// task associated with this entry
	TaskID string `json:"task-id,omitempty"`
}

// CreateTimeEntryResponse captures the response returned from a create time entry action
type CreateTimeEntryResponse struct {
	ID     string `json:"timeLogId"`
	Status string `json:"STATUS"`
}

func (conn *Connection) CreateTimeEntryForProject(projectID string, ops *CreateTimeEntryOps) (*CreateTimeEntryResponse, error) {
	jsonBody, err := json.Marshal(struct {
		TimeEntry *CreateTimeEntryOps `json:"time-entry"`
	}{TimeEntry: ops})
	// fmt.Println(string(jsonBody))
	createResponse := &CreateTimeEntryResponse{}
	method := "POST"
	url := fmt.Sprintf("%sprojects/%s/time_entries.json", conn.Account.Url, projectID)
	reader, _, err := request(conn.ApiToken, method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	// data, _ := ioutil.ReadAll(reader)
	// fmt.Printf(string(data))
	// getHeaders(headers, pages)
	defer reader.Close()

	err = json.NewDecoder(reader).Decode(&struct {
		*CreateTimeEntryResponse `json:"time-entry"`
	}{createResponse})
	if err != nil {
		return nil, err
	}

	return createResponse, nil
}
