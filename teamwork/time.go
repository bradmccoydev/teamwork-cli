package teamwork

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/bradmccoydev/teamwork-cli/model"
)

// CreateTimeEntryForProject - Creates time Entry
func CreateTimeEntryForProject(conn *Connection, projectID string, ops *model.CreateTimeEntryOps) (*model.CreateTimeEntryResponse, error) {
	jsonBody, err := json.Marshal(struct {
		TimeEntry *model.CreateTimeEntryOps `json:"time-entry"`
	}{TimeEntry: ops})
	fmt.Println(string(jsonBody))
	createResponse := &model.CreateTimeEntryResponse{}
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
		*model.CreateTimeEntryResponse `json:"time-entry"`
	}{createResponse})
	if err != nil {
		return nil, err
	}

	return createResponse, nil
}
