package generate

import "time"

// Models
type User struct {
	ID        uint      `gorm:"primary_key"json:"id"`
	Name      string    `gorm:"column:name"json:"name"`
	Email     string    `gorm:"column:email"json:"email"`
	Password  string    `gorm:"column:password"json:"-"`
	CreatedAt time.Time `gorm:"column:created_at"json:"createdAt"`
}

// The table name for the user model
func (User) TableName() string {
	return "users"
}
