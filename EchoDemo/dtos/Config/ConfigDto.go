package config

//CORS CORS域名配置
type CORS struct {
	Domains []string `yaml:"cors_domains"`
}

//DataBase 数据库配置
type DataBase struct {
	Sqlite string `yaml:"sqlite_db"`
}

//Salt 密码盐
type Salt struct {
	Salt string `yaml:"salt"`
}

//Config 系统配置
type Config struct {
	CORS     `yaml:"cors,inline"`
	DataBase `yaml:"database,inline"`
	Salt     `yaml:"salt,inline"`
}
