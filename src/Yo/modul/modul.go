package yo

// 后期考虑配置
type UserData struct {
	Id         int    // user id
	Name       string // user name
	FriendList []int  // 好友列表
	MsgList    []int  // 消息列表
}

type ContactInfo struct {
	Id        int
	Type      int // contact type
	LstUserId []int
}

const (
	EMsgInfo_Unread = iota
	EMsgInfo_Read
)

type MsgInfo struct {
	Id       int
	Type     int
	SenderId int
	GeterId  int
	Data     string
}
