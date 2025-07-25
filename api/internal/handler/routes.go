// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	doctor "github.com/qas491/hospital/api/internal/handler/doctor"
	patient "github.com/qas491/hospital/api/internal/handler/patient"
	"github.com/qas491/hospital/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/care-history",
				Handler: doctor.CreateCareHistoryHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/care-history",
				Handler: doctor.GetCareHistoryListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/care-history/:ch_id",
				Handler: doctor.GetCareHistoryDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/doctors/:doctor_id/performance",
				Handler: doctor.GetDoctorPerformanceHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/medicines",
				Handler: doctor.GetMedicinesListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/medicines/:medicines_id",
				Handler: doctor.GetMedicinesDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/medicines/select",
				Handler: doctor.SelectMedicinesHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/prescriptions",
				Handler: doctor.CreatePrescriptionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/prescriptions",
				Handler: doctor.GetPrescriptionListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/prescriptions/:co_id",
				Handler: doctor.GetPrescriptionDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/prescriptions/:co_id/review",
				Handler: doctor.ReviewPrescriptionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/rankings/weekly",
				Handler: doctor.GetWeeklyRankingHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/appointments",
				Handler: patient.MakeAppointmentHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/appointments/:appointment_id",
				Handler: patient.GetAppointmentHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/appointments/:appointment_id/cancel",
				Handler: patient.CancelAppointmentHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/departments",
				Handler: patient.ListDepartmentsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/departments/:department_id/doctors",
				Handler: patient.ListDoctorsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/doctors/:doctor_id/timeslots",
				Handler: patient.ListTimeSlotsHandler(serverCtx),
			},
		},
	)
}
