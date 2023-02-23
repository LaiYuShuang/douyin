package constants

const (
	MySQLDefaultDSN   = "root:0915@tcp(localhost:3306)/douyin?charset=utf8&parseTime=True&loc=Local"
	UserTableName     = "users"
	VideoTableName    = "videos"
	FavoriteTableName = "favorite"

	EtcdAddress         = "127.0.0.1:2379"
	UserServiceName     = "user"
	UserServiceAddr     = "127.0.0.1:8889"
	VideoServiceName    = "video"
	VideoServiceAddr    = "127.0.0.1:8890"
	FavoriteServiceName = "favorite"
	FavoriteServiceAddr = "127.0.0.1:8891"

	VideoResourceAddr  = "http://202.199.13.136:8080"
	VideoSavePath      = "../../cmd/api/resource/videos"
	VideoCoverSavePath = "../../cmd/api/resource/cover"
	VideoUrlPath       = "resource/videos"
	VideoCoverUrlPath  = "resource/cover"

	EmptyUserId           = -1
	EmptyToken            = ""
	CPURateLimit  float64 = 80.0
	MaxTime               = 9223372036854775807
	VideoLimitNum         = 30

	//jwt token
	Secret = "this is a secret"
	Issuer = "douyin"
	Expire = 7200
)
