package handlers

import (
	"encoding/json"
	"github.com/aws/aws-lamda-go/events"
)

func apiResponse(status int, body interface{})(*event.APIGateProxyRequest, error){
	resp:=events.APIGateProxyRequest{Headers:map[sting]string ["Content-type":"application/json"]}
	resp.status=status

	stringBody,_:=json.Marsha(body)
	resp.Body=sting(stringBody)
	return &resp, nil 
}