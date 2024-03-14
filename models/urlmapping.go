package models

import "github.com/kamva/mgm/v3"

type URLMapping struct {
	mgm.DefaultModel `bson:",inline"`
	OriginalUrl      string `json:"originalUrl" bson:"originalUrl"`
	ShortenedUrl     string `json:"shortenedUrl" bson:"shortenedUrl"`
}
