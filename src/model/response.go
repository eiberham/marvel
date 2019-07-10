package model

import "encoding/json"

type Response struct {
	Code      		int             	`json:"code"`
	Status    		string          	`json:"status"`
	Copyright 		string 				`json:"copyright"`
	AttributionText string 				`json:"attributionText"`
	AttributionHTML string 				`json:"attributionHTML"`
	Data 			json.RawMessage 	`json:"data"`
	Etag 			string  			`json:"etag"`
}
