package frontend

import (
	"errors"

	"github.com/NicholeGit/webMini/model"

	"github.com/chanxuehong/session"
)

//"github.com/chanxuehong/session"

var (
	SessionStorage *session.Storage
	ModelProxy     model.Model

	ErrNotFound = errors.New("item not found")
)

func init() {

	SessionStorage = session.New(20*60, 60*60)
	//	if err := ModelProxy.MysqlDB.Open("mysql", "root:123456@tcp(localhost:3306)/weixinfenxiangv2?parseTime=true&loc=Asia%2FShanghai&timeout=5s&charset=utf8&collation=utf8_general_ci"); err != nil {
	//		panic(err)
	//	}
}

//func GetIntParameter(urlValues url.Values, key string) (i int64, err error) {
//	value := urlValues.Get(key)
//	if value == "" {
//		err = errors.New(key + " 为空")
//		return
//	}
//	i, err = strconv.ParseInt(value, 10, 64)
//	if err != nil {
//		err = errors.New(key + " 参数不正确")
//	}
//	return
//}
