package enotify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	iftttHookURI = "https://maker.ifttt.com/trigger/%s/with/key/%s"
)

type ifttt struct {
	event string
	key   string
}

type hookReq struct {
	V1 string `json:"value1"`
	V2 string `json:"value2"`
	V3 string `json:"value3"`
}

// NewIFTTT sets up a new IFTTT webhook object.
func NewIFTTT(event string, key string) *ifttt {
	return &ifttt{
		event: event,
		key:   key,
	}
}

// Send sends a webhook to IFTTT with the three specified values.
func (i *ifttt) Send(v1 string, v2 string, v3 string) bool {
	reqStruct := hookReq{
		V1: v1,
		V2: v2,
		V3: v3,
	}

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(reqStruct); err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf(iftttHookURI, i.event, i.key), buf)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}
