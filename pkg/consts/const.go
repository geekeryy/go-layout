package consts

const (
	TraceName = "trace_id" // trace名
)

// apollo 配置
const (
	ApolloAppID= "go-layout"
	ApolloCluster="default"
	ApolloNamespace="application"
)


// 环境变量名
const (
	AppName      = "APP_NAME"
	AppVersion   = "APP_VERSION"
	AppEnv       = "APP_ENV"
	ApolloUrl    = "APOLLO_URL"
	ApolloSecret = "APOLLO_ACCESS_KEY_SECRET"
)

var EnvMap = map[string]string{
	AppName:      "go-layout",
	AppVersion:   "v1.0",
	AppEnv:       "dev",
	ApolloUrl:    "http://apollo.dev.jiangyang.me",
	ApolloSecret: "",
}