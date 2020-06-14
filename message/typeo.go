package message

import (
	"time"
)

type loginQueue struct {
	uuid string
	sid  string
}

type tokenReturned struct {
	token string
}

type onlineStatus struct {
	online bool
	sid    int
}
type chanOnLineStatus chan<- onlineStatus

type checkOnlineStatus struct {
	online bool
	sid    int
	token  string
}
type chanCheckOnlineStatus chan<- checkOnlineStatus

type getUnreadNumbers struct {
	sid   int
	token string
}
type chanGetUnreadNumbers chan<- getUnreadNumbers

type unreadNumbers struct {
	badge int
}
type chanUnreadNumbers chan<- unreadNumbers

type unreadMessage struct {
	messageContent string
	messageGen     int
	messageDate    time.Time
	messageRead    bool
}
type checkReadMessage struct {
	messageGen int
	token      string
}
type messagePerPerson map[string]unreadMessage

//["王册全",{"你好","33223344","Time",true}]

type messageUnread map[string]messagePerPerson

//["idUnread",["王册全",{"你好","33223344","Time"}]
