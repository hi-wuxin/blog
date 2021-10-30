package config

//mysql 配置
type Mysql struct {
	Addr     string
	Port     int
	Name     string
	Password string
	Database string
}

//redis 配置
type Redis struct {
	Addr string
	Port int
}
type OSS struct {
	Ak string
	Sk string
}
