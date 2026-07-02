// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package demo6

import (
	"net/http"

	"demo6/internal/logic/demo6"
	"demo6/internal/svc"
	"demo6/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 测试api
func TestapiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TestApiReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := demo6.NewTestapiLogic(r.Context(), svcCtx)
		resp, err := l.Testapi(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
