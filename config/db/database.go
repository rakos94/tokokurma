package db

// DatabaseList is struct for list of database
type DatabaseList struct {
	Nds struct {
		Mysql Database
	}
	NdsTob struct {
		Mysql Database
	}
	WebBrinets struct {
		Mysql Database
	}
	Rabbit struct {
		RabbitMq Database
	}
	NdsParam struct {
		Mysql Database
	}
	NdsLn struct {
		Mysql Database
	}
	Redis Database
}

// Database is struct for Database conf
type Database struct {
	Host     string
	Port     string
	Username string
	Password string
	Dbname   string
	Adapter  string
}
