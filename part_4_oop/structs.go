package part4

import "time"

//Budget budget data type
type Budget struct {
	CampaignID int
	Balance    float64
	Expires    time.Time
}

//ExtendBalance extend balance field by extra
func (b *Budget) ExtendBalance(extra float64) {
	(*b).Balance += extra
}
