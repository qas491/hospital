package patient

import (
	"net/http"

	"github.com/qas491/hospital/api/internal/logic/patient"
	"github.com/qas491/hospital/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetAppointmentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := patient.NewGetAppointmentLogic(r.Context(), svcCtx)
		resp, err := l.GetAppointment()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
