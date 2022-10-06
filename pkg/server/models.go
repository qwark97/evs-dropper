package server

type errResp struct {
	Status int    `json:"status,omitempty"`
	Reason string `json:"reason,omitempty"`
}
