package user

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/subtle"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/NicholeGit/webMini/frontend"
	"github.com/NicholeGit/webMini/model"

	"github.com/NicholeGit/webMini/util"
	"github.com/chanxuehong/util/random"
)

func init() {
	util.HandleFunc("/user/login", LoginHandler)
}

// 登录验证, POST.
// 需提供 username 和 password 两个参数.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Host)
	if r.Method != "POST" {
		return
	}

	if err := r.ParseForm(); err != nil {
		resp := frontend.NewError(frontend.ErrCodeBadRequest, "解析请求出错")
		json.NewEncoder(w).Encode(resp)
		return
	}

	// 判断输入参数
	username := r.FormValue("username")
	if username == "" {
		resp := frontend.NewError(frontend.ErrCodeBadRequest, "username 为空")
		json.NewEncoder(w).Encode(resp)
		return
	}
	password := r.FormValue("password")
	if password == "" {
		resp := frontend.NewError(frontend.ErrCodeBadRequest, "password 为空")
		json.NewEncoder(w).Encode(resp)
		return
	}

	useData, err := frontend.ModelProxy.GetLoginData(username)
	if err != nil {
		if err != model.ErrNotFound {
			resp := frontend.NewError(frontend.ErrCodeInternalServerError, "服务器出错")
			json.NewEncoder(w).Encode(resp)
			return
		}
		resp := frontend.NewError(frontend.ErrCodeAuthenticateFailed, "错误的用户名和密码")
		json.NewEncoder(w).Encode(resp)
		return
	}

	// 验证密码
	Hash := hmac.New(sha1.New, useData.Salt)
	io.WriteString(Hash, password)
	HashSum := Hash.Sum(nil)

	if len(HashSum) != len(useData.PassWord) {
		resp := frontend.NewError(frontend.ErrCodeAuthenticateFailed, "错误的用户名和密码")
		json.NewEncoder(w).Encode(resp)
		return
	}
	if subtle.ConstantTimeCompare(HashSum, useData.PassWord) != 1 {
		resp := frontend.NewError(frontend.ErrCodeAuthenticateFailed, "错误的用户名和密码")
		json.NewEncoder(w).Encode(resp)
		return
	}

	// 验证成功, 创建session
	session := Session{
		SessionId:   string(random.NewSessionId()),
		UserID:      useData.Id,
		AccessToken: random.NewToken(),
	}

	if err := addSession(&session); err != nil {
		resp := frontend.NewError(frontend.ErrCodeInternalServerError, "服务器出错")
		json.NewEncoder(w).Encode(resp)
		return
	}

	cookie := http.Cookie{
		Name:     sidCookieName,
		Value:    session.SessionId,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	resp := struct {
		frontend.Error
		*model.UserData `json:",omitempty"`
		AccessToken     string `json:"access_token"`
	}{
		UserData:    useData,
		AccessToken: string(session.AccessToken),
	}
	json.NewEncoder(w).Encode(&resp)
	return
}
