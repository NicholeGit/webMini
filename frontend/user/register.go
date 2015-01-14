package user

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"io"
	"net/http"

	"github.com/NicholeGit/webMini/frontend"

	"github.com/NicholeGit/webMini/util"
	"github.com/chanxuehong/util/random"
)

func init() {
	util.HandleFunc("/user/register", registerHandler)
}

// 创建用户, POST.
// 提供 username, password
func registerHandler(w http.ResponseWriter, r *http.Request) {
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

	salt := random.NewRawRandom()
	Hash := hmac.New(sha1.New, salt)
	io.WriteString(Hash, password)
	HashSum := Hash.Sum(nil)

	if err := frontend.ModelProxy.CreateUser(username, HashSum, salt); err != nil {
		resp := frontend.NewError(frontend.ErrCodeInternalServerError, "服务器出错")
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := frontend.NewError(frontend.ErrCodeOK, "")
	json.NewEncoder(w).Encode(resp)
	return
}
