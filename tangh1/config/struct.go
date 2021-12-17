package config

type structa struct {
	NoNetWork bool `json:"NoNetWork"`
	Debug     bool `json:"Debug"`
	LogLevel  int  `json:"LogLevel"`
	Port      int  `json:"Port"`
}
