package session

import (
	"github.com/gofrs/uuid"
	"gofun/def"
	"sync"
	"time"
)

var sessionMap sync.Map

func init() {
	sessionMap = sync.Map{}
}

// 获取毫秒级时间戳
func NowInMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

func GenerateNewSession(userName string) string {
	sid := uuid.Must(uuid.NewV4()).String()

	ttl := NowInMilli() + 1000*60*30 //30min

	ss := &def.SimpleSession{
		UserName: userName,
		TTL:      ttl,
	}

	// todo :插入数据库
	sessionMap.LoadOrStore(sid, ss)
	return sid
}

// 程序重启后调用
func LoadSessionFromDB() {
	// todo：从数据库中取出session导入
}

// delete
func DeleteExpireSession(sid string) {
	// todo :数据库删除
	sessionMap.Delete(sid)

}

// SessionIsExpire:判7断session是否过期
func SessionIsExpire(sid string) (string, bool) {
	v, ok := sessionMap.Load(sid)
	if ok {
		if v.(*def.SimpleSession).TTL < NowInMilli() {
			DeleteExpireSession(sid)
			return "", true

		}
		return v.(*def.SimpleSession).UserName, false
	}
	return "", true
}
