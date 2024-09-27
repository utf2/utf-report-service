package transfer

import "github.com/google/uuid"

type FormReportRequest struct {
	FormID uuid.UUID
}

// Response returned immediately.
// Report will be generated when message will be received from Kafka.
type FormReportResponse struct {
	MessageSent bool
}
