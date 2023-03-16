package models

// models user
type User struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	Name       string         `json:"name"`
	Email      string         `json:"email" gorm:"unique"`
	Password   []byte         `json:"-"`
	Notes      []Note         `json:"-"`
}

// models notes
type Note struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `json:"title"`
	Category  string         `json:"category"`
	Details   string         `json:"details"`
	UserID    uint           `json:"user_id"`
}

func (User) TableName() string {
	return "users"
}

func (Note) TableName() string {
	return "notes"
}
