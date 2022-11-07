package part4

import "time"

//Budget budget data type
type Budget struct {
	CampaignID int
	Balance    float64
	Expires    time.Time
}
