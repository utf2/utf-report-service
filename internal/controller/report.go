package controller

import transfer "github.com/utf2/utf-report-service/internal/controller/model"

type ReportControllerAPI interface {
	SendReportMessage(transfer.FormReportRequest) transfer.FormReportResponse
}
