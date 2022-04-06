package models

type DbSettings struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	Hostname string `json:"Hostname"`
	Dbname   string `json:"Dbname"`
}

type RedisSettings struct {
	Address  string `json:"Address"`
	Password string `json:"Password"`
	Port     string `json:"Port"`
}

type Config struct {
	Key           string        `json:"Key"`
	DbSettings    DbSettings    `json:"DbSettings"`
	RedisSettings RedisSettings `json:"RedisSettings"`
	HostName      string        `json:"HostName"`
}
