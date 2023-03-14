package common

import "main.go/service"

type CommonGroup struct {
	ChartApi
}

var commonChartService = service.ServiceGroupApp.CommonServiceGroup.CommonChartService
