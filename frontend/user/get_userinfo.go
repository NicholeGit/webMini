package user

import (
	"encoding/json"
	"net/http"

	"github.com/NicholeGit/webMini/frontend"
	"github.com/NicholeGit/webMini/model"
	"github.com/NicholeGit/webMini/util"
)

func init() {
	util.HandleFunc("/user/get_userinfo", GetUserInfoHandler)
}

// 获取自己信息, GET.
func GetUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		resp := frontend.NewError(frontend.ErrCodeBadRequest, "解析请求出错")
		json.NewEncoder(w).Encode(resp)
		return
	}

	// 鉴权
	_, userData, ok := auth(w, r, r.Form)
	if !ok {
		return
	}
	resp := struct {
		frontend.Error
		*model.UserData `json:",omitempty"`
	}{
		UserData: userData,
	}
	json.NewEncoder(w).Encode(&resp)
	return
}
