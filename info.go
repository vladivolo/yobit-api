package api

import ()

type Info struct {
	Success    uint8                             `json:"success"`
	ServerTime uint64                            `json:"server_time"`
	Pairs      map[string]map[string]interface{} `json:"pairs"`
	Error      string                            `json:"error"`
}
