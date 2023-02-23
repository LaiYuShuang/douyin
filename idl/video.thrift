namespace go user

struct BaseResp {
  1:required i32 status_code  // 状态码，0-成功，其他值-失败
  2:string status_message // 返回状态描述
}

struct douyin_publish_action_request {
  1:required string token // 用户鉴权token
  2:required binary data // 视频数据
  3:required string title  // 视频标题
}

struct douyin_publish_action_response {
  1:BaseResp base_resp
}

struct douyin_publish_list_request {
  1:required i64 user_id  // 用户id
  2:required string token // 用户鉴权token
}

struct douyin_publish_list_response {
  1:BaseResp base_resp
  2:list<Video> video_list  // 用户发布的视频列表
}

struct douyin_feed_request {
  1:optional i64 latest_time // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
  2:optional string token  // 可选参数，登录用户设置
}

struct douyin_feed_response {
  1:BaseResp base_resp
  2:list<Video> video_list  // 视频列表
  4:optional i64 next_time // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

struct Video {
  1:required i64 id // 视频唯一标识
  2:required User author  // 视频作者信息
  3:required string play_url  // 视频播放地址
  4:required string cover_url  // 视频封面地址
  5:required i64 favorite_count // 视频的点赞总数
  6:required i64 comment_count // 视频的评论总数
  7:required bool is_favorite // true-已点赞，false-未点赞
  8:required string title // 视频标题
}

struct User {
  1:required i64 id  // 用户id
  2:required string name  // 用户名称
}

service VideoService{
    douyin_publish_action_response PublishAction(1:douyin_publish_action_request req)
    douyin_publish_list_response PublishList(1:douyin_publish_list_request  req)
    douyin_feed_response  Feed(1:douyin_feed_request  req)
}