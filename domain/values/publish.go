package values

import "time"

type Publish struct {
	Date *time.Time `json:"date"`
}

func (p Publish) IsValid() bool {
	if p.Date != nil {
		return !time.Now().Before(*p.Date)
	}
	return true
}
