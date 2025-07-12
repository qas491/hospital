package patient

import (
	"net/http"

	"github.com/qas491/hospital/api/internal/logic/patient"
	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CancelAppointmentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CancelAppointmentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := patient.NewCancelAppointmentLogic(r.Context(), svcCtx)
		resp, err := l.CancelAppointment(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
