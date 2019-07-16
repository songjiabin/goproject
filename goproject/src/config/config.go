package config

//redis配置文件信息
var RedisConfig = map[string]string{
	"name":    "redis",
	"type":    "tcp",
	"address": "127.0.0.1:6379",
	"auth":    "123456",
}

//mysql配置文件信息
var MySqlConfig = map[string]string{
	"mysql":     "mysql",
	"name":      "root",
	"password":  "123456",
	"type":      "tcp",
	"host":      "3306",
	"allconfig": "root:123456@tcp(localhost:3306)/mydb",
}
