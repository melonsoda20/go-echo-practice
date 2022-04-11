package configs

type DBConfig struct {
	Database struct {
		Postgresql struct {
			Host     string `json:"host:`
			Port     int    `json:"port:`
			User     string `json:"user:`
			Password string `json:"password:`
			DBName   string `json:"dbname:`
		}
	}
}
