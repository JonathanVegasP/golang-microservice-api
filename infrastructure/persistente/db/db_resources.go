package db

func NewDBResources() *DBResources {
	return &DBResources{
		User:   "",
		Pass:   "",
		DBName: "",
	}
}

type DBResources struct {
	User   string
	Pass   string
	DBName string
}
