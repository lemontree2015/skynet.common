package api

import (
	"github.com/lemontree2015/skynet"
	"github.com/lemontree2015/skynet.common/session_client"
)

/////////////////
// Server APIs
/////////////////

// 备注:
// 通过判断返回的*skynet.ServiceKey是否为nil来判断是否删除了OldServiceKey
func DelSessionWithSessionId(account string, sessionId uint64) (*skynet.ServiceKey, uint64, error) {
	return session_client.DelSessionWithSessionId(account, sessionId)
}

/////////////////
// Session APIs
/////////////////

// 备注:
// 通过判断返回的*skynet.ServiceKey是否为nil来判断是否有对应的Session存在
func TouchSession(account string) (*skynet.ServiceKey, uint64, error) {
	return session_client.TouchSession(account)
}

// 备注:
// 通过判断返回的*skynet.ServiceKey是否为nil来判断是否有OldServiceKey
func SetSession(account string, serviceKey *skynet.ServiceKey, remoteSessionId uint64) (*skynet.ServiceKey, uint64, error) {
	return session_client.SetSession(account, serviceKey, remoteSessionId)
}
