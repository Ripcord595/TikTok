package protocal

type DouyinCommentActionRequest struct {
	Token       *string `protobuf:"bytes,1,req,name=token" json:"token,omitempty"`                                // 用户鉴权token
	VideoId     *int64  `protobuf:"varint,2,req,name=video_id,json=videoId" json:"video_id,omitempty"`            // 视频id
	ActionType  *int32  `protobuf:"varint,3,req,name=action_type,json=actionType" json:"action_type,omitempty"`   // 1-发布评论，2-删除评论
	CommentText *string `protobuf:"bytes,4,opt,name=comment_text,json=commentText" json:"comment_text,omitempty"` // 用户填写的评论内容，在action_type=1的时候使用
	CommentId   *int64  `protobuf:"varint,5,opt,name=comment_id,json=commentId" json:"comment_id,omitempty"`      // 要删除的评论id，在action_type=2的时候使用
}

type DouyinCommentActionResponse struct {
	StatusCode *int32   `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty"`     // 返回状态描述
	Comment    *Comment `protobuf:"bytes,3,opt,name=comment" json:"comment,omitempty"`                          // 评论成功返回评论内容，不需要重新拉取整个列表
}
