package models

type Manager struct {
	ManagerID int64 `json:"manager_id"`
	IsFree    bool  `json:"is_free"`
	ClientID  int64 `json:"client_id"`
}

type Request struct {
	ClientID int64 `json:"client_id"`
}

type EndRequest struct {
	ManagerID int64 `json:"manager_id"`
}
