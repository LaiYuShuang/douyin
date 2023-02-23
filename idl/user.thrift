namespace go user

struct BaseResp {
  1:i32 status_code
  2:string status_message
}

struct douyin_user_register_request {
  1:required string username
  2:required string password
}

struct douyin_user_register_response {
  1:BaseResp base_resp
  2:required i64 user_id
  3:required string token
}

struct douyin_user_login_request {
  1:required string username
  2:required string password
}

struct douyin_user_login_response {
  1:BaseResp base_resp
  2:required i64 user_id
  3:required string token
}

struct douyin_user_request {
  1:required i64 user_id
  2:required string token
}

struct douyin_user_response {
  1:BaseResp base_resp
  2:required User user
}

struct User {
  1:required i64 id
  2:required string name
}

service UserService {
      douyin_user_register_response CreateUser(1:douyin_user_register_request req)
      douyin_user_login_response CheckUser(1:douyin_user_login_request req)
      douyin_user_response QueryCurUser(1:douyin_user_request req)
}
