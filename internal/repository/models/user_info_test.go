package models

import (
	"fmt"
	"os"
	"sync"
	"testing"
)

var (
	userCache    = make(map[int]*UserInfo)
	userCacheMux sync.Mutex
)

func TestMain(m *testing.M) {
	InitDB()
	code := m.Run()
	os.Exit(code)
}

func getUserInfoFromCache(userID int) (*UserInfo, bool) {
	userCacheMux.Lock()
	defer userCacheMux.Unlock()

	user, exists := userCache[userID]
	return user, exists
}

func addUserToCache(userID int, user *UserInfo) {
	userCacheMux.Lock()
	defer userCacheMux.Unlock()

	userCache[userID] = user
}

func Test_QueryUserInfoById(t *testing.T) {
	userID := 6
	user, exists := getUserInfoFromCache(userID)

	if !exists {
		user = &UserInfo{}
		err := NewUserInfoDAO().QueryUserInfoById(int64(userID), user)
		if err != nil {
			t.Fatalf("QueryUserInfoById failed: %v", err)
		}

		addUserToCache(userID, user)
	}

	fmt.Printf("user name: %s\n", user.Name)
}

func Test_IsUserExistById(t *testing.T) {
	userID := 9
	user, exists := getUserInfoFromCache(userID)

	if !exists {
		user = &UserInfo{}
		err := NewUserInfoDAO().QueryUserInfoById(int64(userID), user)
		if err != nil {
			t.Fatalf("QueryUserInfoById failed: %v", err)
		}

		addUserToCache(userID, user)
	}

	if user == nil {
		t.Fatalf("User does not exist")
	}

	fmt.Printf("User exists: %v\n", user != nil)
}
