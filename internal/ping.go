package internal

// 用户的 unamePing 是否重复
//msg, err := entity.CheckUnamePing(req, dsUser)
//if err != nil {
//	return dto.UserAccount{}, err
//}
//if msg != "" {
//	return dto.UserAccount{Msg: msg}, nil
//}

//func (entity *CreateUserService) CheckUnamePing(req *auth.AddUserReq, dsUser *data.UserData) (msg string, err error) {
//	name := []rune(req.UnamePing)
//	n := len(name) - 1
//	for n > 0 {
//		if name[n] < 48 || name[n] > 57 {
//			break
//		}
//		n--
//	}
//	unamePing := string(name[:n+1])
//
//	nu := 0
//	str := string(name[n+1:])
//	if str != "" {
//		nu, err = strconv.Atoi(str)
//		if err != nil {
//			entity.LogWarnf("string convert int fail: [str: %s] [error: %s]", str, err.Error())
//			return "", components.ErrorSystemError
//		}
//	}
//
//	userList, err := dsUser.GetUserListByUnamePing(req.SchoolId, unamePing)
//	if err != nil {
//		return "", components.ErrorSystemError
//	}
//	if len(userList) == 0 {
//		return "", nil
//	}
//
//	var userUnamePingList []int
//	for _, user := range userList {
//		if !PrefixMatch(unamePing, user.UnamePing) {
//			continue
//		}
//		tmp := []rune(user.UnamePing)
//		i := len(tmp) - 1
//		for i > 0 {
//			if tmp[i] < 48 || tmp[i] > 57 {
//				break
//			}
//			i--
//		}
//		s := string(tmp[i+1:])
//		if s == "" {
//			userUnamePingList = append(userUnamePingList, 0)
//			continue
//		}
//		atoi, err := strconv.Atoi(s)
//		if err != nil {
//			return "", components.ErrorSystemError
//		}
//		userUnamePingList = append(userUnamePingList, atoi)
//	}
//	l := len(userUnamePingList)
//	for _, v := range userUnamePingList {
//		if nu != v {
//			continue
//		}
//		for j := 0; j < l-1; j++ {
//			if userUnamePingList[j]+1 == userUnamePingList[j+1] {
//				continue
//			}
//			suf := ""
//			if l >= 10 {
//				suf = strconv.Itoa(userUnamePingList[j] + 1)
//			} else {
//				suf = "0" + strconv.Itoa(userUnamePingList[j]+1)
//			}
//			return "姓名全拼重复，推荐使用：" + unamePing + suf, nil
//		}
//		suf := ""
//		if l >= 10 {
//			suf = strconv.Itoa(l)
//		} else {
//			suf = "0" + strconv.Itoa(l)
//		}
//		return "姓名全拼重复，推荐使用：" + unamePing + suf, nil
//	}
//
//	return "", nil
//}
