package utils

import "encoding/json"

func JsonStatus(status string) []byte {
    resp, _ := json.Marshal(struct {
        Status string `json:"status"`
    }{
        Status: status,
    })

    return resp
}
