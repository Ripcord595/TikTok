package protocal

type DouyinFeedRequest struct {
	LatestTime *int64  `protobuf:"varint,1,opt,name=latest_time,json=latestTime" json:"latest_time,omitempty"` // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token      *string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`                              // 可选参数，登录用户设置
}

type DouyinFeedResponse struct {
	StatusCode *int32   `protobuf:"varint,1,req,name=status_code,json=statusCode" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg" json:"status_msg,omitempty"`     // 返回状态描述
	VideoList  []*Video `protobuf:"bytes,3,rep,name=video_list,json=videoList" json:"video_list,omitempty"`     // 视频列表
	NextTime   *int64   `protobuf:"varint,4,opt,name=next_time,json=nextTime" json:"next_time,omitempty"`       // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

//type Video struct {
//	Id            *int64  `protobuf:"varint,1,req,name=id" json:"id,omitempty"`                                            // 视频唯一标识
//	Author        *User   `protobuf:"bytes,2,req,name=author" json:"author,omitempty"`                                     // 视频作者信息
//	PlayUrl       *string `protobuf:"bytes,3,req,name=play_url,json=playUrl" json:"play_url,omitempty"`                    // 视频播放地址
//	CoverUrl      *string `protobuf:"bytes,4,req,name=cover_url,json=coverUrl" json:"cover_url,omitempty"`                 // 视频封面地址
//	FavoriteCount *int64  `protobuf:"varint,5,req,name=favorite_count,json=favoriteCount" json:"favorite_count,omitempty"` // 视频的点赞总数
//	CommentCount  *int64  `protobuf:"varint,6,req,name=comment_count,json=commentCount" json:"comment_count,omitempty"`    // 视频的评论总数
//	IsFavorite    *bool   `protobuf:"varint,7,req,name=is_favorite,json=isFavorite" json:"is_favorite,omitempty"`          // true-已点赞，false-未点赞
//	Title         *string `protobuf:"bytes,8,req,name=title" json:"title,omitempty"`                                       // 视频标题
//}
