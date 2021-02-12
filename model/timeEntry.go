package model

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
