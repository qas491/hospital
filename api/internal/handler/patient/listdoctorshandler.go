package patient

import (
	"net/http"

	"github.com/qas491/hospital/api/internal/logic/patient"
	"github.com/qas491/hospital/api/internal/svc"
	"github.com/qas491/hospital/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListDoctorsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListDoctorsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := patient.NewListDoctorsLogic(r.Context(), svcCtx)
		resp, err := l.ListDoctors(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
