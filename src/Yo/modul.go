package yo

// 后期考虑配置
type UserData struct {
	Id   int    // user id
	Name string // user name
}

type ContactInfo struct {
    Id        int
	Type      int // contact type
	LstUserId []int
}

type MsgInfo struct {
    Id        int
	Type      int
	LstUserId []int
	Data      string
}

type Server struct {
	// 一组用户列表
	LstUser []*UserData
	// 一组用户关系列表
	LstContact []*ContactInfo
	// 一组用户消息列表
	LstMsg []*MsgInfo
}

/*
函数名：根据用户名字获得用户

参数：注册用户名

返回值：u，当 u 为 nil 得时候，表示没有找到 UserData
*/
func (s *Server) GetUserByName(username string) (u *UserData) {
	for _, user := range s.UserData {
		if user.Name == username {
			return user
		}
	}

	return nil
}

/*
函数名：注册用户

参数：注册用户名称

返回值：err u

err : 错误信息
"success" 创建成功
"nameexist" 重名

u : 一个用户实例，创建不成功时，其值为 Null
*/
func (s *Server) RegistUser(username string) (err string, u *UserData) {
	if s.GetUserByName(username) != nil {
		return "nameexist", nil
	}

	user := make(UserData)
	user.Id = len(s.UserData) + 1
	user.Name = username

	s.UserData = append(s.UserData, user)
	return "succes", user
}

/*
函数名：增加新好友

参数：id username

id：提起加好友请求者

username：被邀请好友对象

返回值：err

err : 错误信息
"success":增加成功
"iduserempty":加好友请求者为空
"idnameissameone":自己加自己
"nameuserempty":被请求者为空
"alreadyfriend":两人已经是好友
*/
func (s *Server) AddFriend(id int, username string) (err string) {
	if len(s.UserData) <= id {
		return "iduserempty"
	}

	user1 := s.UserData[id]
	if user1 == nil {
		return "iduserempty"
	}

	user2 := s.GetUserByName(username)
	if user2 == nil {
		return "nameuserempty"
	}

    if user1.Id == user2.Id {
        return "idnameissameone"
    }

	for _, contactinfo := range s.LstContact{
        var bHasUser1, bHasUser2 bool
        bHasUser1 = false;
        bHasUser2 = false;

        for _, userId := range contactinfo.LstUserId {
            if bHasUser1 == false and userId == user1.Id {
                bHasUser1 = true;
            }

            if bHasUser2 == false and userId == user2.Id {
                bHasUser2 = true;
            }

            if bHasUser1 && bHasUser2 {
                return "alreadyfriend"
            }
        }
	}

    newcontact = make(ContactInfo)
    newcontact.Id = len(s.LstContact)
    newcontact.LstUserId = append(newcontact.LstUserId, user1.Id)
    newcontact.LstUserId = append(newcontact.LstUserId, user2.Id)

    s.LstContact = append(s.LstContact, newcontact)
    return "success"
}

/*
函数名：列出所有好友

参数：user id

返回值：err lstContact

err:错误信息
success: 取出成功
emptyuser: 传入id没有对应的 user

lstContact:所有的好友
*/
func (s *Server) GetFriendList(id int)(err string, lstContact []UserData) {
    user := s.LstUser[id]
    if user == nil {
        return "emptyuser", lstContact
    }

    var LstId []int
    for i, contactinfo := range s.LstContact {
        for _, userId := range contactinfo.LstUserId {
            if userId == user.Id {
                LstId = append(LstId, i)
                break
            }
        }
    }

    for _, index := range LstId {
        for _, userId := range s.LstContact[index].LstUserId {
            if userId != user.Id {
                lstContact = append(lstContact, &s.LstUser[userId])
            }
        }
    }

    return "success", lstContact
}
