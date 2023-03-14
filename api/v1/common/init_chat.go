package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"main.go/model/common"
	"main.go/model/common/request"
	"main.go/model/common/response"
	"sync"
	"time"
)

// 读写锁
var userChatLocker sync.RWMutex
var userChatMap map[int]*UserNode = make(map[int]*UserNode, 0)

// 读写锁
var adminChatLocker sync.RWMutex
var adminChatMap map[int]*AdminNode = make(map[int]*AdminNode, 0)
var adminMapNum = 0

type ChartNode interface {
	CloseChat(int)          //关闭连接
	SendProc(int, int, int) //发送消息
	RecvProc()              //接受消息
	Lock()                  //加锁
	UnLock()                //解锁
}

// 管理员聊天是绑定
type AdminNode struct {
	Uid       int
	AdminConn *websocket.Conn
	DataQueue chan *request.ChartMsg
	Locker    chan int
}

// 初始化管理员连接
func InitAdminChat(aid, uid int, aconn *websocket.Conn) ChartNode {
	//放入map中
	node := &AdminNode{
		Uid:       uid,
		AdminConn: aconn,
		DataQueue: make(chan *request.ChartMsg, 50),
		Locker:    make(chan int),
	}
	adminChatLocker.Lock()
	adminChatMap[aid] = node
	adminMapNum++
	adminChatLocker.Unlock()
	return node
}

// 返回在线管理员conn
func GetAdminNode(uid int) ([]*AdminNode, error) {
	remap := make([]*AdminNode, 0)
	sum := 0
	for _, v := range adminChatMap {
		if v.AdminConn != nil && uid == v.Uid {
			remap = append(remap, v)
			sum++
		}
	}
	if sum == 0 {
		return nil, errors.New("没有找到管理员在线")
	}
	return remap, nil
}

// 管理员下线
func (u *AdminNode) CloseChat(aid int) {
	adminChatLocker.Lock()
	u.AdminConn.Close()
	adminMapNum--
	delete(adminChatMap, aid)
	adminChatLocker.Unlock()
}

// 发送消息fSendProc
func (u *AdminNode) SendProc(uid, suid, isAdmin int) {
	for {
		select {
		case msg := <-u.DataQueue:
			chartMsg := &response.ResHistoryList{
				Chart: common.Charts{
					MsgId:       0,
					SendId:      suid,
					GetId:       uid,
					IsSendAdmin: isAdmin,
					SendTime:    int(time.Now().Unix()),
					MsgType:     msg.MsgType,
					Msg:         msg.Msg,
				},
				AdminName: commonChartService.GetAdminName(suid),
			}
			datamsg, _ := json.Marshal(chartMsg)
			if ugnode, err := GetUserNode(uid); err != nil {
				//用户不在线，进行处理
				commonChartService.AddUUnRead(uid)
				fmt.Println("用户不在线")
			} else {
				ugnode.UserConn.WriteMessage(websocket.TextMessage, datamsg)
			}

			if adminChatGroup, err := GetAdminNode(uid); err != nil {
				//没有管理员连接，进行其他操作
				commonChartService.AddAUnRead(uid)
				fmt.Println("没有管理员在线")
			} else {
				for _, v := range adminChatGroup {
					err := v.AdminConn.WriteMessage(websocket.TextMessage, datamsg)
					if err != nil {
						continue
					}
				}
			}
			//对消息进行持久化处理
			commonChartService.AddMsg(msg, uid, suid, isAdmin)
		}
	}
}

// 接收消息
func (u *AdminNode) RecvProc() {
	for {
		_, data, err := u.AdminConn.ReadMessage()
		if err != nil {
			fmt.Println("接受消息错误", err)
			break
		}
		msg := &request.ChartMsg{}
		json.Unmarshal(data, msg)
		u.DataQueue <- msg
	}
	u.Lock()
}
func (u *AdminNode) Lock() {
	u.Locker <- 1
}
func (u *AdminNode) UnLock() {
	<-u.Locker
}

// 用户聊天绑定
type UserNode struct {
	UserConn  *websocket.Conn
	DataQueue chan *request.ChartMsg
	Locker    chan int
}

// 初始化用户连接
func InitUserChat(uid int, uconn *websocket.Conn) ChartNode {
	//已存在就先删除掉
	userChatLocker.Lock()
	if uba, ok := userChatMap[uid]; ok {
		uba.UserConn.Close()
		delete(userChatMap, uid)
	}
	//放入map中
	node := &UserNode{
		UserConn:  uconn,
		DataQueue: make(chan *request.ChartMsg, 50),
		Locker:    make(chan int),
	}
	userChatMap[uid] = node
	userChatLocker.Unlock()
	return node
}

// 返回用户conn
func GetUserNode(uid int) (*UserNode, error) {
	if ucm, ok := userChatMap[uid]; ok {
		return ucm, nil
	}
	return nil, errors.New("没有找到用户连接")
}

// 用户下线
func (u *UserNode) CloseChat(uid int) {
	userChatLocker.Lock()
	u.UserConn.Close()
	delete(userChatMap, uid)
	userChatLocker.Unlock()
}

// 发送消息fSendProc
func (u *UserNode) SendProc(uid, suid, isAdmin int) {
	for {
		select {
		case msg := <-u.DataQueue:
			chartMsg := &response.ResHistoryList{
				Chart: common.Charts{
					MsgId:       0,
					SendId:      suid,
					GetId:       uid,
					IsSendAdmin: isAdmin,
					SendTime:    int(time.Now().Unix()),
					MsgType:     msg.MsgType,
					Msg:         msg.Msg,
				},
				AdminName: "",
			}
			datamsg, _ := json.Marshal(chartMsg)
			ugnode, err := GetUserNode(uid)
			if err != nil {
				//用户不在线，进行处理
				commonChartService.AddUUnRead(uid)
				fmt.Println("用户不在线")
			}
			//else {
			//	ugnode.UserConn.WriteMessage(websocket.TextMessage, datamsg)
			//}

			if adminChatGroup, err := GetAdminNode(uid); err != nil {
				//没有管理员连接，进行其他操作
				commonChartService.AddAUnRead(uid)
				fmt.Println("没有管理员在线")
				liveMsg := commonChartService.GetLivingMsg()
				nowtime := time.Now()
				liveFlag := false
				if liveMsg.EntTime > liveMsg.StartTime && nowtime.Hour() < liveMsg.EntTime && nowtime.Hour() > liveMsg.StartTime {
					liveFlag = true
				}
				if liveMsg.EntTime < liveMsg.StartTime && (nowtime.Hour() > liveMsg.StartTime || (nowtime.Hour() < liveMsg.EntTime && nowtime.Hour() >= 0)) {
					liveFlag = true
				}
				if liveMsg.IsUse != 0 {
					liveFlag = false
				}
				if liveFlag {
					chartMsg.Chart.Msg = liveMsg.LiveMsg
					chartMsg.Chart.IsSendAdmin = 1
					liveMsgjson, _ := json.Marshal(chartMsg)
					ugnode.UserConn.WriteMessage(websocket.TextMessage, liveMsgjson)
				}

			} else {
				for _, v := range adminChatGroup {
					err := v.AdminConn.WriteMessage(websocket.TextMessage, datamsg)
					if err != nil {
						continue
					}
				}
			}
			//对消息进行持久化处理
			commonChartService.AddMsg(msg, uid, suid, isAdmin)
		}
	}
}

// 接收消息
func (u *UserNode) RecvProc() {
	for {
		_, data, err := u.UserConn.ReadMessage()
		if err != nil {
			fmt.Println("接受消息错误", err)
			break
		}
		msg := &request.ChartMsg{}
		json.Unmarshal(data, msg)
		u.DataQueue <- msg
	}
	u.Lock()
}
func (u *UserNode) Lock() {
	u.Locker <- 1
}
func (u *UserNode) UnLock() {
	<-u.Locker
}
