package buymeacoffee

import (
	"encoding/json"
)

type Event struct {
	Type     string          `json:"type"`
	LiveMode bool            `json:"live_mode"`
	Data     json.RawMessage `json:"data"`
}

type MembershipStartedData struct {
	SupporterEmail     string `json:"supporter_email"`
	CurrentPeriodStart int64  `json:"current_period_start"`
}

type MembershipCancelledData struct {
	SupporterEmail   string `json:"supporter_email"`
	CurrentPeriodEnd int64  `json:"current_period_end"`
}
