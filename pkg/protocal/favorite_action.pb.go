package protocal

type DouyinFavoriteActionRequest struct {
	Token      *string `protobuf:"bytes,1,req,name=token" json:"token,omitempty"`                              // 用户鉴权token
	VideoId    *int64  `protobuf:"varint,2,req,name=video_id,json=videoId" json:"video_id,omitempty"`          // 视频id
	ActionType *int32  `protobuf:"varint,3,req,name=action_type,json=actionType" json:"action_type,omitempty"` // 1-点赞，2-取消点赞
}
