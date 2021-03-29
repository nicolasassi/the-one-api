package values

import (
	"testing"
	"time"
)

func TestPublish_IsValid(t *testing.T) {
	type fields struct {
		Date *time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"default", fields{Date: func() *time.Time {
			t := time.Now().Add(-10 * time.Hour)
			return &t
		}()}, true},
		{"publish_date_after_now", fields{Date: func() *time.Time {
			t := time.Now().Add(10 * time.Hour)
			return &t
		}()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Publish{Date: tt.fields.Date}
			if got := p.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v. want %v", got, tt.want)
			}
		})
	}
}
