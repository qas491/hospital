package svc

import (
	"github.com/qas491/hospital/api/internal/config"
	"github.com/qas491/hospital/doctor_srv/doctor"
	"github.com/qas491/hospital/patient_srv/patient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	DoctorRpc  doctor.DoctorServiceClient
	PatientRpc patient.MedicalServiceClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		DoctorRpc:  doctor.NewDoctorServiceClient(zrpc.MustNewClient(c.DoctorRpcConf).Conn()),
		PatientRpc: patient.NewMedicalServiceClient(zrpc.MustNewClient(c.PatientRpcConf).Conn()),
	}
}
