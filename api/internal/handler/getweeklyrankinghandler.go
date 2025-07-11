package handler

import (
	"net/http"

	"hospital/api/internal/logic"
	"hospital/api/internal/svc"
	"hospital/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetWeeklyRankingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetWeeklyRankingReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		// 设置默认值
		if req.Limit <= 0 {
			req.Limit = 10
		}

		l := logic.NewGetWeeklyRankingLogic(r.Context(), svcCtx)
		resp, err := l.GetWeeklyRanking(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
