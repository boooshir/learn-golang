package objectrelationmapping

import (
	"time"

	"gorm.io/gorm"
)

type Name struct {
	FirstName  string `gorm:"column:first_name"`
	MiddleName string `gorm:"column:middle_name"`
	LastName   string `gorm:"colimn:last_name"`
}

type User struct {
	ID           int       `gorm:"primaryKey;column:id;<-:create"`
	Password     string    `gorm:"column:password"`
	Name         Name      `gorm:"embedded"`
	CreatedAt    string    `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt    string    `gotm:"column:updated_at;autoUpdateTime"`
	Wallet       Wallet    `gorm:"foreignKey:user_id;reference:id"`
	Addresses    []Address `gorm:"foreignKey:user_id;reference:id"`
	LikeProducts []Product `gorm:"many2many:user_like_product;foreignKey:id;joinForeignKey:user_id;reference:id;joinReferences:product_id"`
}

func (u *User) TableName() string {
	return "users"
}

// func (u *User) BeforeCreate(db *gorm) error {
// 	if u.ID == "" {
// 		u.ID = "user-"+time.Now().Format("2006010200000")
// 	}
// 	return nil
// }

type UserLog struct {
	ID        int    `gorm:"primaryKey;column:id;autoIncrement"`
	UserID    int    `gorm:"column:user_id"`
	Action    string `gorm:"column:action"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime:milli;"`
	UpdatedAt int64  `gotm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (u *UserLog) TableName() string {
	return "user_logs"
}

type Todo struct {
	gorm.Model
	ID          int64  `gorm:"primaryKey;column:id;autoIncrement"`
	UserID      int    `gorm:"column:user_id"`
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`
	// CreatedAt   time.Time      `gorm:"column:created_at;autoCreateTime;type:datetime"`
	// UpdatedAt   time.Time      `gorm:"column:updated_at;autoCreateTime;autoUpdateTime;type:datetime"`
	// DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:datetime"`
}

func (u *Todo) TableName() string {
	return "todos"
}

type Wallet struct {
	ID        string    `gorm:"primaryKey;column:id"`
	UserId    int       `gorm:"column:user_id"`
	Balance   int64     `gorm:"column:balance"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;type:datetime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime;type:datetime"`
	User      *User     `gorm:"foreignKey:user_id;reference:id"`
}

func (w *Wallet) TableName() string {
	return "wallets"
}

type Address struct {
	ID        int       `gorm:"primaryKey;column:id"`
	UserId    int       `gorm:"column:user_id"`
	Address   string    `gorm:"column:address"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;type:datetime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime;type:datetime"`
	User      User      `gorm:"foreignKey:user_id;reference:id"`
}

func (a *Address) TableName() string {
	return "addresses"
}

type Product struct {
	ID           string    `gorm:"primaryKey;column:id"`
	Name         string    `gorm:"column:name"`
	Price        int       `gorm:"column:price"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime;type:datetime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime;type:datetime"`
	LikedByUsers []User    `gorm:"many2many:user_like_product;foreignKey:id;joinForeignKey:product_id;reference:id;joinReferences:user_id"`
}

func (p *Product) TableName() string {
	return "products"
}

type Guestbook struct {
	ID        int       `gorm:"primary_key;column:id;autoIncrement"`
	Name      string    `gorm:"column:name"`
	Email     string    `gorm:"column:email"`
	Message   string    `gorm:"column:message"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;type:datetime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime;type:datetime"`
}

func (g *Guestbook) TableName() string {
	return "guest_books"
}
