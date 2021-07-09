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
