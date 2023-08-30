package protocal

type DouyinCommentListRequest struct {
	Token   *string `protobuf:"bytes,1,req,name=token" json:"token,omitempty"`                     // 用户鉴权token
	VideoId *int64  `protobuf:"varint,2,req,name=video_id,json=videoId" json:"video_id,omitempty"` // 视频id
}

type DouyinCommentListResponse struct {
	StatusCode  *int32     `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg   *string    `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty"`       // 返回状态描述
	CommentList []*Comment `protobuf:"bytes,3,rep,name=comment_list,json=commentList" json:"comment_list,omitempty"` // 评论列表
}

type Comment struct {
	Id         *int64  `protobuf:"varint,1,req,name=id" json:"id,omitempty"`                                  // 视频评论id
	User       *User   `protobuf:"bytes,2,req,name=user" json:"user,omitempty"`                               // 评论用户信息
	Content    *string `protobuf:"bytes,3,req,name=content" json:"content,omitempty"`                         // 评论内容
	CreateDate *string `protobuf:"bytes,4,req,name=create_date,json=createDate" json:"create_date,omitempty"` // 评论发布日期，格式 mm-dd
}
