package services

import (
	"errors"
	"tiktok/internal/repository/models"
	"tiktok/pkg/middleware"
)

const (
	MaxUsernameLength = 100
	MaxPasswordLength = 20
	MinPasswordLength = 8
)

type LoginResponse struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

func QueryUserLogin(username, password string) (*LoginResponse, error) {
	return NewQueryUserLoginFlow(username, password).Do()
}

func NewQueryUserLoginFlow(username, password string) *QueryUserLoginFlow {
	return &QueryUserLoginFlow{username: username, password: password}
}

type QueryUserLoginFlow struct {
	username string
	password string

	data   *LoginResponse
	userid int64
	token  string
}

func (q *QueryUserLoginFlow) Do() (*LoginResponse, error) {
	if err := q.checkNum(); err != nil {
		return nil, err
	}
	if err := q.prepareData(); err != nil {
		return nil, err
	}
	if err := q.packData(); err != nil {
		return nil, err
	}
	return q.data, nil
}

func (q *QueryUserLoginFlow) checkNum() error {
	if q.username == "" {
		return errors.New("用户名为空")
	}
	if len(q.username) > MaxUsernameLength {
		return errors.New("用户名长度超出限制")
	}
	if q.password == "" {
		return errors.New("密码为空")
	}
	return nil
}

func (q *QueryUserLoginFlow) prepareData() error {
	userLoginDAO := models.NewUserLoginDao()
	var login models.UserLogin
	err := userLoginDAO.QueryUserLogin(q.username, q.password, &login)
	if err != nil {
		return err
	}
	q.userid = login.UserInfoId

	token, err := middleware.ReleaseToken(login)
	if err != nil {
		return err
	}
	q.token = token
	return nil
}

func (q *QueryUserLoginFlow) packData() error {
	q.data = &LoginResponse{
		UserId: q.userid,
		Token:  q.token,
	}
	return nil
}
