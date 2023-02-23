namespace go favorite

struct BaseResp {
  1:i32 status_code      // 状态码，0-成功，其他值-失败
  2:string status_message // 返回状态描述
}

struct douyin_favorite_action_request {
  1:required string token // 用户鉴权token
  2:required i64 video_id // 视频id
  3:required i32 action_type  // 1-点赞，2-取消点赞
}

struct douyin_favorite_action_response {
    1:BaseResp base_rasp
}

struct douyin_favorite_list_request {
  1:required i64 user_id  // 用户id
  2:required string token // 用户鉴权token
}

struct douyin_favorite_list_response {
  1:BaseResp base_rape
  2:list<Video> video_list  // 用户点赞视频列表
}

struct Video {
  1:required i64 id  // 视频唯一标识
  2:required User author // 视频作者信息
  3:required string play_url // 视频播放地址
  4:required string cover_url // 视频封面地址
  5:required i64 favorite_count  // 视频的点赞总数
  6:required i64 comment_count // 视频的评论总数
  7:required bool is_favorite  // true-已点赞，false-未点赞
  8:required string title // 视频标题
}

struct User {
  1:required i64 id // 用户id
  2:required string name // 用户名称
}

service FavoriteService {
    douyin_favorite_action_response FavoriteAction(1:douyin_favorite_action_request req)
    douyin_favorite_list_response FavoriteList(1:douyin_favorite_list_request req)
}