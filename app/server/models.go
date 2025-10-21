package server

type MessageData struct {
	TypeEvent   string `json:"type_event"`
	TypeService string `json:"type_service"`
	ClientIP    string `json:"client_ip"`
}
