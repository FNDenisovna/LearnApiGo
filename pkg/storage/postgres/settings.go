package postgres

type Settings struct {
	Host     string
	Port     uint16
	Database string
	User     string
	Password string
	SSLMode  string
}
