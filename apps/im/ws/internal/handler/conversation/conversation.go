package conversation

import (
	"context"
	"easy-chat/apps/im/ws/internal/logic"
	"easy-chat/apps/im/ws/internal/svc"
	"easy-chat/apps/im/ws/websocket"
	"easy-chat/apps/im/ws/ws"
	"easy-chat/pkg/constants"
	"easy-chat/pkg/wuid"
	"github.com/mitchellh/mapstructure"
	v1 "k8s.io/api/autoscaling/v1"
	"time"
)

// Chat 单独定义会话是因为聊天本身是需要在会话中去完成的，所以不在用户中去创建会话
func Chat(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(srv *websocket.Server, conn *websocket.Conn, msg *websocket.Message) {
		// todo: 私聊
		var data ws.Chat
		// 对map[string]interface{}类型进行转换
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			srv.Send(websocket.NewErrMessage(err), conn)
			return
		}

		switch data.ChatType {
		case constants.SingleChatType:
			err := logic.NewConversation(context.Background(), srv, svc).SingleChat(&data, conn.Uid)
			if err != nil {
				srv.Send(websocket.NewErrMessage(err), conn)
				return
			}

			srv.SendByUserId(websocket.NewMessage(conn.Uid, ws.Chat{
				ConversationId: data.ConversationId,
				ChatType:       data.ChatType,
				SendId:         conn.Uid,
				RecvId:         data.RecvId,
				SendTime:       time.Now().UnixMilli(),
				Msg:            data.Msg,
			}), data.RecvId)
		}
	}
}

// MarkRead 处理 WebSocket 消息，标记消息为已读。
//
// 该函数返回一个 websocket.HandlerFunc 处理函数，用于接收并处理标记消息为已读的请求。
// 它将 WebSocket 消息解码为 ws.MarkRead 结构体，并将其传递给消息读取传输客户端进行处理。
// 如果解码或消息处理失败，将通过 WebSocket 向客户端发送错误信息。
//
// 参数:
//   - svc: 包含服务上下文的 *svc.ServiceContext，用于访问消息读取传输客户端。
//
// 返回:
//   - websocket.HandlerFunc: 处理 WebSocket 消息的处理函数。
func MarkRead(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(srv *websocket.Server, conn *websocket.Conn, msg *websocket.Message) {
		// todo: 已读未读处理
		var data ws.MarkRead
		// 解码 WebSocket 消息数据为 ws.MarkRead 结构体
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			// 如果解码失败，发送错误信息到客户端
			err := srv.Send(websocket.NewErrMessage(err), conn)
			if err != nil {
				srv.Errorf("error message send error: %v", err)
			}
			return
		}

		// 将标记已读的请求发送到消息读取传输客户端
		err := svc.MsgReadTransferClient.Push(&mq.MsgMarkRead{
			ChatType:       data.ChatType,
			ConversationId: data.ConversationId,
			SendId:         conn.Uid,
			RecvId:         data.RecvId,
			MsgIds:         data.MsgIds,
		})
		if err != nil {
			// 如果消息处理失败，发送错误信息到客户端
			err := srv.Send(websocket.NewErrMessage(err), conn)
			if err != nil {
				srv.Errorf("error message send error: %v", err)
			}
			return
		}
	}
}
