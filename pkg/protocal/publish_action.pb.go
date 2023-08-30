package protocal

type DouyinPublishActionRequest struct {
	Token *string `protobuf:"bytes,1,req,name=token" json:"token,omitempty"` // 用户鉴权token
	Data  []byte  `protobuf:"bytes,2,req,name=data" json:"data,omitempty"`   // 视频数据
	Title *string `protobuf:"bytes,3,req,name=title" json:"title,omitempty"` // 视频标题
}

type DouyinPublishActionResponse struct {
	StatusCode *int32  `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty"`     // 返回状态描述
}
