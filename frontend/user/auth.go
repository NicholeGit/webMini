package user

import (
	"crypto/subtle"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/NicholeGit/webMini/frontend"
	"github.com/NicholeGit/webMini/model"
)

// 鉴权, 判断是否登录, 是否状态正常, 正常 ok == true, 否则为 false
func auth(w http.ResponseWriter, r *http.Request, urlValues url.Values) (session *Session, userData *model.UserData, ok bool) {
	accessToken := urlValues.Get("access_token")
	if accessToken == "" {
		resp := frontend.NewError(frontend.ErrCodeUnauthenticated, "您没有登录")
		json.NewEncoder(w).Encode(resp)
		return
	}
	cookie, err := r.Cookie(sidCookieName)
	if err != nil {
		resp := frontend.NewError(frontend.ErrCodeUnauthenticated, "您没有登录")
		json.NewEncoder(w).Encode(resp)
		return
	}
	session, err = getSession(cookie.Value)
	if err != nil {
		if err != frontend.ErrNotFound {
			resp := frontend.NewError(frontend.ErrCodeInternalServerError, "服务器出错")
			json.NewEncoder(w).Encode(resp)
			return
		}
		resp := frontend.NewError(frontend.ErrCodeUnauthenticated, "您没有登录, 或者您是非法用户")
		json.NewEncoder(w).Encode(resp)
		return
	}
	if len(accessToken) != len(session.AccessToken) {
		deleteSession(session.SessionId)

		resp := frontend.NewError(frontend.ErrCodeUnauthenticated, "您没有登录, 或者您是非法用户")
		json.NewEncoder(w).Encode(resp)
		return
	}
	if subtle.ConstantTimeCompare([]byte(accessToken), session.AccessToken) != 1 {
		deleteSession(session.SessionId)

		resp := frontend.NewError(frontend.ErrCodeUnauthenticated, "您没有登录, 或者您是非法用户")
		json.NewEncoder(w).Encode(resp)
		return
	}
	userData, err = frontend.ModelProxy.UserGet(session.UserID)
	if err != nil {
		if err == model.ErrNotFound {
			deleteSession(session.SessionId)
		}
		resp := frontend.NewError(frontend.ErrCodeInternalServerError, "服务器出错, 或者您已经被删除")
		json.NewEncoder(w).Encode(resp)
		return
	}

	ok = true
	return
}
