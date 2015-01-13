package user

import (
	"errors"

	"github.com/chanxuehong/session"

	"github.com/NicholeGit/webMini/frontend"
)

const sidCookieName = "sidweb"

type Session struct {
	SessionId   string
	UserID      int64
	AccessToken []byte
}

// 根据 sid 获取 Session, 如果找不到返回 frontend.ErrNotFound.
func getSession(sid string) (ss *Session, err error) {
	v, err := frontend.SessionStorage.Get(sid)
	if err != nil {
		if err != session.ErrNotFound {
			return
		}
		err = frontend.ErrNotFound
		return
	}

	ss, ok := v.(*Session)
	if !ok {
		err = errors.New("服务器出错, Session 类型不匹配")
		return
	}
	return
}

func addSession(ss *Session) (err error) {
	return frontend.SessionStorage.Add(ss.SessionId, ss)
}

func setSession(ss *Session) (err error) {
	return frontend.SessionStorage.Set(ss.SessionId, ss)
}

func deleteSession(sid string) (err error) {
	return frontend.SessionStorage.Delete(sid)
}
