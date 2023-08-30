package protocal

type DouyinUserLoginRequest struct {
	Username *string `protobuf:"bytes,1,req,name=username" json:"username,omitempty"` // 登录用户名
	Password *string `protobuf:"bytes,2,req,name=password" json:"password,omitempty"` // 登录密码
}

type DouyinUserLoginResponse struct {
	StatusCode *int32  `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty"`     // 返回状态描述
	UserId     *int64  `protobuf:"varint,3,req,name=user_id,json=userId" json:"user_id,omitempty"`             // 用户id
	Token      *string `protobuf:"bytes,4,req,name=token" json:"token,omitempty"`                              // 用户鉴权token
}
