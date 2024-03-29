package server_client

import (
	"fmt"
	"github.com/lemontree2015/skynet"
	"github.com/lemontree2015/skynet.common/gproto"
	"github.com/lemontree2015/skynet/client"
)

////////////////////
// RPC UserCount
////////////////////

type RPCGIMServerUserCountIn struct {
}

type RPCGIMServerUserCountOut struct {
	Code  string
	Count int
}

// 查询在线用户到目标GIM Server[N] Service
func UserCount(serviceKey *skynet.ServiceKey) (error, int) {
	// 构造参数
	in := &RPCGIMServerUserCountIn{}
	out := &RPCGIMServerUserCountOut{}

	// RPC请求
	if err := client.GetClient(serviceKey.Name, serviceKey.Version).Send("UserCount", in, out); err == nil && out != nil {
		if out.Code == "success" {
			return nil, out.Count
		} else {
			return fmt.Errorf(out.Code), 0
		}
	} else {
		return err, 0
	}
}

/////////////////////////
// RPC GetAndDelMessage
/////////////////////////

type RPCGIMServerGetAndDelMessageIn struct {
	Account string
}

type RPCGIMServerGetAndDelMessageOut struct {
	Code    string
	Message map[string]*gproto.GProtoMessageNotify
}

// 查询在线用户到目标GIM Server[N] Service
func GetAndDelMessage(serviceKey *skynet.ServiceKey, account string) (error, map[string]*gproto.GProtoMessageNotify) {
	// 构造参数
	in := &RPCGIMServerGetAndDelMessageIn{
		Account: account,
	}
	out := &RPCGIMServerGetAndDelMessageOut{}

	// RPC请求
	if err := client.GetClient(serviceKey.Name, serviceKey.Version).Send("GetAndDelMessage", in, out); err == nil && out != nil {
		if out.Code == "success" {
			return nil, out.Message
		} else {
			return fmt.Errorf(out.Code), nil
		}
	} else {
		return err, nil
	}
}

////////////////////
// RPC KickOffNotify
////////////////////

type RPCGIMServerKickoffNotifyIn struct {
	Account       string
	SessionId     uint64
	KickOffNotify *gproto.GProtoKickOffNotify
}

type RPCGIMServerKickoffNotifyOut struct {
	Code string
}

// Route一条KickoffNotify消息到目标GIM Server[N] Service
func KickoffNotify(serviceKey *skynet.ServiceKey, account string, sessionId uint64, gProtoKickOffNotify *gproto.GProtoKickOffNotify) error {
	// 构造参数
	in := &RPCGIMServerKickoffNotifyIn{
		Account:       account,
		SessionId:     sessionId,
		KickOffNotify: gProtoKickOffNotify,
	}
	out := &RPCGIMServerKickoffNotifyOut{}

	// RPC请求
	if err := client.GetClient(serviceKey.Name, serviceKey.Version).Send("KickoffNotify", in, out); err == nil && out != nil {
		if out.Code == "success" {
			return nil
		} else {
			return fmt.Errorf(out.Code)
		}
	} else {
		return err
	}
}

///////////////////////
// RPC MessageResponse
///////////////////////

type RPCGIMServerMessageResponseIn struct {
	Account         string
	SessionId       uint64
	MessageResponse *gproto.GProtoMessageResponse
}

type RPCGIMServerMessageResponseOut struct {
	Code string
}

// Route一条MessageResponse消息到目标GIM Server[N] Service
func MessageResponse(serviceKey *skynet.ServiceKey, account string, sessionId uint64, gProtoMessageResponse *gproto.GProtoMessageResponse) error {
	// 构造参数
	in := &RPCGIMServerMessageResponseIn{
		Account:         account,
		SessionId:       sessionId,
		MessageResponse: gProtoMessageResponse,
	}
	out := &RPCGIMServerMessageResponseOut{}

	// RPC请求
	if err := client.GetClient(serviceKey.Name, serviceKey.Version).Send("MessageResponse", in, out); err == nil && out != nil {
		if out.Code == "success" {
			return nil
		} else {
			return fmt.Errorf(out.Code)
		}
	} else {
		return err
	}
}

////////////////////
// RPC MessageNotify
////////////////////

type RPCGIMServerMessageNotifyIn struct {
	Account       string
	SessionId     uint64
	MessageNotify *gproto.GProtoMessageNotify
}

type RPCGIMServerMessageNotifyOut struct {
	Code string
}

// Route一条MessageNotify消息到目标GIM Server[N] Service
func MessageNotify(serviceKey *skynet.ServiceKey, account string, sessionId uint64, gProtoMessageNotify *gproto.GProtoMessageNotify) error {
	// 构造参数
	in := &RPCGIMServerMessageNotifyIn{
		Account:       account,
		SessionId:     sessionId,
		MessageNotify: gProtoMessageNotify,
	}
	out := &RPCGIMServerMessageNotifyOut{}

	// RPC请求
	if err := client.GetClient(serviceKey.Name, serviceKey.Version).Send("MessageNotify", in, out); err == nil && out != nil {
		if out.Code == "success" {
			return nil
		} else {
			return fmt.Errorf(out.Code)
		}
	} else {
		return err
	}
}

///////////////////////
// RPC AddFriendNotify
///////////////////////

type RPCGIMServerAddFriendNotifyIn struct {
	Account         string
	SessionId       uint64
	AddFriendNotify *gproto.GProtoAddFriendNotify
}

type RPCGIMServerAddFriendNotifyOut struct {
	Code string
}

// Route一条AddFriendNotify消息到目标GIM Server[N] Service
func AddFriendNotify(serviceKey *skynet.ServiceKey, account string, sessionId uint64, gProtoAddFriendNotify *gproto.GProtoAddFriendNotify) error {
	// 构造参数
	in := &RPCGIMServerAddFriendNotifyIn{
		Account:         account,
		SessionId:       sessionId,
		AddFriendNotify: gProtoAddFriendNotify,
	}
	out := &RPCGIMServerAddFriendNotifyOut{}

	// RPC请求
	if err := client.GetClient(serviceKey.Name, serviceKey.Version).Send("AddFriendNotify", in, out); err == nil && out != nil {
		if out.Code == "success" {
			return nil
		} else {
			return fmt.Errorf(out.Code)
		}
	} else {
		return err
	}
}

//////////////////////////
// RPC AnswerFriendNotify
//////////////////////////
type RPCGIMServerAnswerFriendNotifyIn struct {
	Account            string
	SessionId          uint64
	AnswerFriendNotify *gproto.GProtoAnswerFriendNotify
}

type RPCGIMServerAnswerFriendNotifyOut struct {
	Code string
}

// Route一条AnswerFriendNotify消息到目标GIM Server[N] Service
func AnswerFriendNotify(serviceKey *skynet.ServiceKey, account string, sessionId uint64, gProtoAnswerFriendNotify *gproto.GProtoAnswerFriendNotify) error {
	// 构造参数
	in := &RPCGIMServerAnswerFriendNotifyIn{
		Account:            account,
		SessionId:          sessionId,
		AnswerFriendNotify: gProtoAnswerFriendNotify,
	}
	out := &RPCGIMServerAnswerFriendNotifyOut{}

	// RPC请求
	if err := client.GetClient(serviceKey.Name, serviceKey.Version).Send("AnswerFriendNotify", in, out); err == nil && out != nil {
		if out.Code == "success" {
			return nil
		} else {
			return fmt.Errorf(out.Code)
		}
	} else {
		return err
	}
}

//////////////////////
// RPC DelFriendNotify
//////////////////////
type RPCGIMServerDelFriendNotifyIn struct {
	Account         string
	SessionId       uint64
	DelFriendNotify *gproto.GProtoDelFriendNotify
}

type RPCGIMServerDelFriendNotifyOut struct {
	Code string
}

// Route一条DelFriendNotify消息到目标GIM Server[N] Service
func DelFriendNotify(serviceKey *skynet.ServiceKey, account string, sessionId uint64, gProtoDelFriendNotify *gproto.GProtoDelFriendNotify) error {
	// 构造参数
	in := &RPCGIMServerDelFriendNotifyIn{
		Account:         account,
		SessionId:       sessionId,
		DelFriendNotify: gProtoDelFriendNotify,
	}
	out := &RPCGIMServerDelFriendNotifyOut{}

	// RPC请求
	if err := client.GetClient(serviceKey.Name, serviceKey.Version).Send("DelFriendNotify", in, out); err == nil && out != nil {
		if out.Code == "success" {
			return nil
		} else {
			return fmt.Errorf(out.Code)
		}
	} else {
		return err
	}
}

///////////////////////////////
// RPC GetChatRoomInfoResponse
///////////////////////////////
type RPCGIMServerGetChatRoomInfoResponseIn struct {
	Account                 string
	SessionId               uint64
	GetChatRoomInfoResponse *gproto.GProtoGetChatRoomInfoResponse
}

type RPCGIMServerGetChatRoomInfoResponseOut struct {
	Code string
}

// Route一条GetChatRoomInfoResponse消息到目标GIM Server[N] Service
func GetChatRoomInfoResponse(serviceKey *skynet.ServiceKey, account string, sessionId uint64, gProtoGetChatRoomInfoResponse *gproto.GProtoGetChatRoomInfoResponse) error {
	// 构造参数
	in := &RPCGIMServerGetChatRoomInfoResponseIn{
		Account:                 account,
		SessionId:               sessionId,
		GetChatRoomInfoResponse: gProtoGetChatRoomInfoResponse,
	}
	out := &RPCGIMServerGetChatRoomInfoResponseOut{}

	// RPC请求
	if err := client.GetClient(serviceKey.Name, serviceKey.Version).Send("GetChatRoomInfoResponse", in, out); err == nil && out != nil {
		if out.Code == "success" {
			return nil
		} else {
			return fmt.Errorf(out.Code)
		}
	} else {
		return err
	}
}

////////////////////////////
// RPC EnterChatRoomResponse
////////////////////////////
type RPCGIMServerEnterChatRoomResponseIn struct {
	Account               string
	SessionId             uint64
	EnterChatRoomResponse *gproto.GProtoEnterChatRoomResponse
}

type RPCGIMServerEnterChatRoomResponseOut struct {
	Code string
}

// Route一条EnterChatRoomResponse消息到目标GIM Server[N] Service
func EnterChatRoomResponse(serviceKey *skynet.ServiceKey, account string, sessionId uint64, gProtoEnterChatRoomResponse *gproto.GProtoEnterChatRoomResponse) error {
	// 构造参数
	in := &RPCGIMServerEnterChatRoomResponseIn{
		Account:               account,
		SessionId:             sessionId,
		EnterChatRoomResponse: gProtoEnterChatRoomResponse,
	}
	out := &RPCGIMServerEnterChatRoomResponseOut{}

	// RPC请求
	if err := client.GetClient(serviceKey.Name, serviceKey.Version).Send("EnterChatRoomResponse", in, out); err == nil && out != nil {
		if out.Code == "success" {
			return nil
		} else {
			return fmt.Errorf(out.Code)
		}
	} else {
		return err
	}
}

////////////////////////////
// RPC EnterChatRoomNotify
////////////////////////////
type RPCGIMServerEnterChatRoomNotifyIn struct {
	Account             string
	SessionId           uint64
	EnterChatRoomNotify *gproto.GProtoEnterChatRoomNotify
}

type RPCGIMServerEnterChatRoomNotifyOut struct {
	Code string
}

// Route一条EnterChatRoomNotify消息到目标GIM Server[N] Service
func EnterChatRoomNotify(serviceKey *skynet.ServiceKey, account string, sessionId uint64, gProtoEnterChatRoomNotify *gproto.GProtoEnterChatRoomNotify) error {
	// 构造参数
	in := &RPCGIMServerEnterChatRoomNotifyIn{
		Account:             account,
		SessionId:           sessionId,
		EnterChatRoomNotify: gProtoEnterChatRoomNotify,
	}
	out := &RPCGIMServerEnterChatRoomNotifyOut{}

	// RPC请求
	if err := client.GetClient(serviceKey.Name, serviceKey.Version).Send("EnterChatRoomNotify", in, out); err == nil && out != nil {
		if out.Code == "success" {
			return nil
		} else {
			return fmt.Errorf(out.Code)
		}
	} else {
		return err
	}
}

////////////////////////////
// RPC LeaveChatRoomResponse
////////////////////////////
type RPCGIMServerLeaveChatRoomResponseIn struct {
	Account               string
	SessionId             uint64
	LeaveChatRoomResponse *gproto.GProtoLeaveChatRoomResponse
}

type RPCGIMServerLeaveChatRoomResponseOut struct {
	Code string
}

// Route一条LeaveChatRoomResponse消息到目标GIM Server[N] Service
func LeaveChatRoomResponse(serviceKey *skynet.ServiceKey, account string, sessionId uint64, gProtoLeaveChatRoomResponse *gproto.GProtoLeaveChatRoomResponse) error {
	// 构造参数
	in := &RPCGIMServerLeaveChatRoomResponseIn{
		Account:               account,
		SessionId:             sessionId,
		LeaveChatRoomResponse: gProtoLeaveChatRoomResponse,
	}
	out := &RPCGIMServerLeaveChatRoomResponseOut{}

	// RPC请求
	if err := client.GetClient(serviceKey.Name, serviceKey.Version).Send("LeaveChatRoomResponse", in, out); err == nil && out != nil {
		if out.Code == "success" {
			return nil
		} else {
			return fmt.Errorf(out.Code)
		}
	} else {
		return err
	}
}

////////////////////////////
// RPC LeaveChatRoomNotify
////////////////////////////
type RPCGIMServerLeaveChatRoomNotifyIn struct {
	Account             string
	SessionId           uint64
	LeaveChatRoomNotify *gproto.GProtoLeaveChatRoomNotify
}

type RPCGIMServerLeaveChatRoomNotifyOut struct {
	Code string
}

// Route一条LeaveChatRoomNotify消息到目标GIM Server[N] Service
func LeaveChatRoomNotify(serviceKey *skynet.ServiceKey, account string, sessionId uint64, gProtoLeaveChatRoomNotify *gproto.GProtoLeaveChatRoomNotify) error {
	// 构造参数
	in := &RPCGIMServerLeaveChatRoomNotifyIn{
		Account:             account,
		SessionId:           sessionId,
		LeaveChatRoomNotify: gProtoLeaveChatRoomNotify,
	}
	out := &RPCGIMServerLeaveChatRoomNotifyOut{}

	// RPC请求
	if err := client.GetClient(serviceKey.Name, serviceKey.Version).Send("LeaveChatRoomNotify", in, out); err == nil && out != nil {
		if out.Code == "success" {
			return nil
		} else {
			return fmt.Errorf(out.Code)
		}
	} else {
		return err
	}
}

////////////////////////////
// RPC ChatRoomKickoffNotify
////////////////////////////
type RPCGIMServerChatRoomKickoffNotifyIn struct {
	Account               string
	SessionId             uint64
	ChatRoomKickoffNotify *gproto.GProtoChatRoomKickoffNotify
}

type RPCGIMServerChatRoomKickoffNotifyOut struct {
	Code string
}

// Route一条ChatRoomKickoffNotify消息到目标GIM Server[N] Service
func ChatRoomKickoffNotify(serviceKey *skynet.ServiceKey, account string, sessionId uint64, gProtoChatRoomKickoffNotify *gproto.GProtoChatRoomKickoffNotify) error {
	// 构造参数
	in := &RPCGIMServerChatRoomKickoffNotifyIn{
		Account:               account,
		SessionId:             sessionId,
		ChatRoomKickoffNotify: gProtoChatRoomKickoffNotify,
	}
	out := &RPCGIMServerChatRoomKickoffNotifyOut{}

	// RPC请求
	if err := client.GetClient(serviceKey.Name, serviceKey.Version).Send("ChatRoomKickoffNotify", in, out); err == nil && out != nil {
		if out.Code == "success" {
			return nil
		} else {
			return fmt.Errorf(out.Code)
		}
	} else {
		return err
	}
}

////////////////////////////
// RPC EnterGroupNotify
////////////////////////////
type RPCGIMServerEnterGroupNotifyIn struct {
	Account          string
	SessionId        uint64
	EnterGroupNotify *gproto.GProtoEnterGroupNotify
}

type RPCGIMServerEnterGroupNotifyOut struct {
	Code string
}

// Route一条EnterGroupNotify消息到目标GIM Server[N] Service
func EnterGroupNotify(serviceKey *skynet.ServiceKey, account string, sessionId uint64, gProtoEnterGroupNotify *gproto.GProtoEnterGroupNotify) error {
	// 构造参数
	in := &RPCGIMServerEnterGroupNotifyIn{
		Account:          account,
		SessionId:        sessionId,
		EnterGroupNotify: gProtoEnterGroupNotify,
	}
	out := &RPCGIMServerEnterGroupNotifyOut{}

	// RPC请求
	if err := client.GetClient(serviceKey.Name, serviceKey.Version).Send("EnterGroupNotify", in, out); err == nil && out != nil {
		if out.Code == "success" {
			return nil
		} else {
			return fmt.Errorf(out.Code)
		}
	} else {
		return err
	}
}

////////////////////////////
// RPC LeaveGroupNotify
////////////////////////////
type RPCGIMServerLeaveGroupNotifyIn struct {
	Account          string
	SessionId        uint64
	LeaveGroupNotify *gproto.GProtoLeaveGroupNotify
}

type RPCGIMServerLeaveGroupNotifyOut struct {
	Code string
}

// Route一条LeaveGroupNotify消息到目标GIM Server[N] Service
func LeaveGroupNotify(serviceKey *skynet.ServiceKey, account string, sessionId uint64, gProtoLeaveGroupNotify *gproto.GProtoLeaveGroupNotify) error {
	// 构造参数
	in := &RPCGIMServerLeaveGroupNotifyIn{
		Account:          account,
		SessionId:        sessionId,
		LeaveGroupNotify: gProtoLeaveGroupNotify,
	}
	out := &RPCGIMServerLeaveGroupNotifyOut{}

	// RPC请求
	if err := client.GetClient(serviceKey.Name, serviceKey.Version).Send("LeaveGroupNotify", in, out); err == nil && out != nil {
		if out.Code == "success" {
			return nil
		} else {
			return fmt.Errorf(out.Code)
		}
	} else {
		return err
	}
}

////////////////////////////
// RPC GroupKickoffNotify
////////////////////////////
type RPCGIMServerGroupKickoffNotifyIn struct {
	Account            string
	SessionId          uint64
	GroupKickoffNotify *gproto.GProtoGroupKickoffNotify
}

type RPCGIMServerGroupKickoffNotifyOut struct {
	Code string
}

// Route一条GroupKickoffNotify消息到目标GIM Server[N] Service
func GroupKickoffNotify(serviceKey *skynet.ServiceKey, account string, sessionId uint64, gProtoGroupKickoffNotify *gproto.GProtoGroupKickoffNotify) error {
	// 构造参数
	in := &RPCGIMServerGroupKickoffNotifyIn{
		Account:            account,
		SessionId:          sessionId,
		GroupKickoffNotify: gProtoGroupKickoffNotify,
	}
	out := &RPCGIMServerGroupKickoffNotifyOut{}

	// RPC请求
	if err := client.GetClient(serviceKey.Name, serviceKey.Version).Send("GroupKickoffNotify", in, out); err == nil && out != nil {
		if out.Code == "success" {
			return nil
		} else {
			return fmt.Errorf(out.Code)
		}
	} else {
		return err
	}
}

///////////////////////
// RPC AddBlacklistNotify
///////////////////////

type RPCGIMServerAddBlacklistNotifyIn struct {
	Account            string
	SessionId          uint64
	AddBlacklistNotify *gproto.GProtoAddBlacklistNotify
}

type RPCGIMServerAddBlacklistNotifyOut struct {
	Code string
}

// Route一条AddBlacklistNotify消息到目标GIM Server[N] Service
func AddBlacklistNotify(serviceKey *skynet.ServiceKey, account string, sessionId uint64, gProtoAddBlacklistNotify *gproto.GProtoAddBlacklistNotify) error {
	// 构造参数
	in := &RPCGIMServerAddBlacklistNotifyIn{
		Account:            account,
		SessionId:          sessionId,
		AddBlacklistNotify: gProtoAddBlacklistNotify,
	}
	out := &RPCGIMServerAddBlacklistNotifyOut{}

	// RPC请求
	if err := client.GetClient(serviceKey.Name, serviceKey.Version).Send("AddBlacklistNotify", in, out); err == nil && out != nil {
		if out.Code == "success" {
			return nil
		} else {
			return fmt.Errorf(out.Code)
		}
	} else {
		return err
	}
}

//////////////////////
// RPC DelBlacklistNotify
//////////////////////
type RPCGIMServerDelBlacklistNotifyIn struct {
	Account            string
	SessionId          uint64
	DelBlacklistNotify *gproto.GProtoDelBlacklistNotify
}

type RPCGIMServerDelBlacklistNotifyOut struct {
	Code string
}

// Route一条DelBlacklistNotify消息到目标GIM Server[N] Service
func DelBlacklistNotify(serviceKey *skynet.ServiceKey, account string, sessionId uint64, gProtoDelBlacklistNotify *gproto.GProtoDelBlacklistNotify) error {
	// 构造参数
	in := &RPCGIMServerDelBlacklistNotifyIn{
		Account:            account,
		SessionId:          sessionId,
		DelBlacklistNotify: gProtoDelBlacklistNotify,
	}
	out := &RPCGIMServerDelBlacklistNotifyOut{}

	// RPC请求
	if err := client.GetClient(serviceKey.Name, serviceKey.Version).Send("DelBlacklistNotify", in, out); err == nil && out != nil {
		if out.Code == "success" {
			return nil
		} else {
			return fmt.Errorf(out.Code)
		}
	} else {
		return err
	}
}
