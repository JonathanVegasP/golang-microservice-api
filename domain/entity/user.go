package entity

type User struct {
	ID       uint64 `gorm:"primary_key;auto_increment"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
}
