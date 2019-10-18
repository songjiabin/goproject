package service

import (
	"fmt"
	"myapiserver/model"
	"myapiserver/util"

	"sync"
)

//查询所有的userList
func Listuser(name string, offset, limit int) ([]*model.UserInfo, uint, error) {

	userInfos := make([]*model.UserInfo, 0)
	users, count, err := model.ListUser(name, offset, limit)
	if err != nil {
		return nil, count, err
	}

	ids := []uint{}
	for _, user := range users {
		ids = append(ids, user.Id)
	}

	wg := sync.WaitGroup{}

	userList := model.UserList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint]*model.UserInfo, len(users)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	for _, u := range users {
		wg.Add(1)
		go func(u *model.UserModel) {
			defer wg.Done()
			shortId, err := util.GenShortId()
			if err != nil {
				//写入错误
				errChan <- err
				return
			}
			//添加锁
			userList.Lock.Lock()
			defer userList.Lock.Unlock()
			userList.IdMap[u.Id] = &model.UserInfo{
				Id:        u.Id,
				Username:  u.Username,
				SayHello:  fmt.Sprintf("Hello %s", shortId),
				Password:  u.Password,
				CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
			}
		}(u)

	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, count, err
	}

	for _, id := range ids {
		userInfos = append(userInfos, userList.IdMap[id])
	}

	return userInfos, count, nil
}
