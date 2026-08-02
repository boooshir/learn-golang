package objectrelationmapping

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

func OpenConnection() *gorm.DB {
	dsn := "dev.db?_loc=auto&_datetime_format=limit"
	dialect := sqlite.Open(dsn)
	db, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetConnMaxLifetime(30 * time.Minute)
	sqlDb.SetConnMaxLifetime(time.Minute * 5)
	return db
}

var db = OpenConnection()

func TestOpenConnection(t *testing.T) {
	assert.NotNil(t, db)
}

func TestExecuteSQL(t *testing.T) {
	err := db.Exec("insert into sample(id, name) values (?,?)", "1", "boshir").Error
	assert.Nil(t, err)
	err = db.Exec("insert into sample(id, name) values (?,?)", "2", "budi").Error
	assert.Nil(t, err)
	err = db.Exec("insert into sample(id, name) values (?,?)", "3", "joko").Error
	assert.Nil(t, err)
	err = db.Exec("insert into sample(id, name) values (?,?)", "4", "rully").Error
	assert.Nil(t, err)

}

type Sample struct {
	Id   string `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func TestRawSQL(t *testing.T) {
	var sample Sample

	err := db.Raw("select id, name from sample where id = ?", "1").Scan(&sample).Error
	assert.Nil(t, err)
	assert.Equal(t, "1", sample.Id)

	var samples []Sample
	err = db.Raw("select id, name from sample").Scan(&samples).Error
	assert.Nil(t, err)
	assert.Equal(t, 4, len(samples))
}

func TestSqlRows(t *testing.T) {
	rows, err := db.Raw("select id, name from sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	var samples []Sample
	for rows.Next() {
		var id string
		var name string

		err := rows.Scan(&id, &name)
		assert.Nil(t, err)

		samples = append(samples, Sample{
			Id:   id,
			Name: name,
		})
	}

	assert.Equal(t, 4, len(samples))
}

func TestScanRows(t *testing.T) {
	var samples []Sample
	rows, err := db.Raw("select id, name from sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	for rows.Next() {
		err := db.ScanRows(rows, &samples)
		assert.Nil(t, err)
	}

	assert.Equal(t, 4, len(samples))
}

func TestCreateUser(t *testing.T) {
	user := User{
		Password: "rahasia",
		Name: Name{
			FirstName:  "boshir",
			MiddleName: "yusuf",
			LastName:   "ahmad",
		},
	}

	response := db.Create(&user)
	assert.Nil(t, response.Error)
	assert.Equal(t, int64(1), response.RowsAffected)
}

func TestBatchInsert(t *testing.T) {
	var users []User

	for i := range 8 {
		users = append(users, User{
			Name: Name{
				FirstName: "user " + strconv.Itoa(i),
			},
			Password: "rahasia",
		})
	}

	response := db.Create(&users)
	assert.Nil(t, response.Error)
	assert.Equal(t, int64(8), response.RowsAffected)
}

func TestTransaction(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&User{
			Password: "rahasia",
			Name: Name{
				FirstName: "User 10",
			},
		}).Error
		if err != nil {
			return err
		}
		err = tx.Create(&User{
			Password: "rahasia",
			Name: Name{
				FirstName: "User 11",
			},
		}).Error
		if err != nil {
			return err
		}
		err = tx.Create(&User{
			Password: "rahasia",
			Name: Name{
				FirstName: "User 12",
			},
		}).Error
		if err != nil {
			return err
		}

		return nil
	})

	assert.Nil(t, err)
}

func TestManualTransaction(t *testing.T) {
	tx := db.Begin()
	defer tx.Rollback()

	err := tx.Create(&User{Password: "rahasia", Name: Name{FirstName: "User 13"}}).Error
	assert.Nil(t, err)

	err = tx.Create(&User{Password: "rahasia", Name: Name{FirstName: "User 14"}}).Error
	assert.Nil(t, err)

	if err == nil {
		tx.Commit()
	}
}

func TestQuerySingleObject(t *testing.T) {
	user := User{}
	result := db.First(&user)
	assert.Nil(t, result.Error)
	assert.Equal(t, 1, user.ID)

	user = User{}
	result = db.Last(&user)
	assert.Nil(t, result.Error)
	assert.Equal(t, 15, user.ID)

}

func TestQueryInlineCondition(t *testing.T) {
	user := User{}
	// SELECT * FROM `users` WHERE id = 5 ORDER BY `users`.`id` LIMIT 1
	// result := db.First(&user, "id = ?", 5)
	// SELECT * FROM `users` WHERE id = 5 LIMIT 1
	result := db.Take(&user, "id = ?", 5)
	assert.Nil(t, result.Error)
	assert.Equal(t, 5, user.ID)
}

func TestQueryAllObject(t *testing.T) {
	var users []User

	result := db.Find(&users, "id in ?", []int{1, 3, 4, 5})
	assert.Nil(t, result.Error)
	assert.Equal(t, 4, len(users))
}

func TestQueryCondition(t *testing.T) {
	var users []User

	result := db.Where("first_name like ?", "%User%").Where("password = ?", "rahasia").Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 13, len(users))
}

func TestOrOperator(t *testing.T) {
	var users []User
	result := db.Where("first_name like ?", "%User%").Or("password = ?", "rahasia").Find(&users)

	assert.Nil(t, result.Error)
	assert.Equal(t, 14, len(users))
}

func TestNotOperator(t *testing.T) {
	var users []User

	result := db.Not("first_name like ? ", "%User%").Where("password = ?", "rahasia").Find(&users)

	assert.Nil(t, result.Error)
	assert.Equal(t, 1, len(users))
}

func TestSelectFields(t *testing.T) {
	var users []User
	result := db.Select("id", "first_name").Find(&users)
	assert.Nil(t, result.Error)

	for _, user := range users {
		assert.NotNil(t, user.ID)
		assert.NotEqual(t, "", user.Name.FirstName)
	}

	assert.Equal(t, 14, len(users))
}

func TestStructCondition(t *testing.T) {
	userCondition := User{
		Name: Name{
			FirstName: "user 5",
		},
	}

	var users []User
	result := db.Where(userCondition).Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 1, len(users))
}

func TestMapCondition(t *testing.T) {
	mapCondition := map[string]any{
		"middle_name": "",
	}
	var users []User
	result := db.Where(mapCondition).Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 13, len(users))
}

func TestOrderLimitOffset(t *testing.T) {
	var users []User
	result := db.Order("id asc, first_name asc").Limit(5).Offset(5).Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 5, len(users))
	assert.Equal(t, 7, users[0].ID)
}

type UserResponse struct {
	ID        string
	FirstName string
	LastName  string
}

func TestQueryNonModel(t *testing.T) {
	var users []UserResponse
	result := db.Model(&User{}).Select("id", "first_name", "last_name").Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 14, len(users))
}

func TestUpdate(t *testing.T) {
	user := User{}
	result := db.First(&user, "id = ?", 1)
	assert.Nil(t, result.Error)

	user.Name.FirstName = "budi"
	user.Name.MiddleName = ""
	user.Name.LastName = "hanzo"
	user.Password = "secret"
	result = db.Save(&user)
	assert.Nil(t, result.Error)
}

func TestSelectedColumn(t *testing.T) {
	result := db.Model(&User{}).Where("id = ?", 1).Updates(map[string]any{
		"middle_name": "jiro",
		"last_name":   "kayaba",
	})
	assert.Nil(t, result.Error)

	result = db.Model(&User{}).Where("id = ?", 1).Update("password", "diubahlagi")
	assert.Nil(t, result.Error)

	result = db.Where("id = ?", 1).Updates(User{
		Name: Name{
			FirstName: "obito",
			LastName:  "uchiha",
		},
	})
	assert.Nil(t, result.Error)
}

func TestAutoIncrement(t *testing.T) {
	for range 10 {
		userLog := UserLog{
			UserID: 1,
			Action: "test action",
		}
		result := db.Create(&userLog)
		assert.Nil(t, result.Error)
		assert.NotEqual(t, 0, userLog.ID)
		fmt.Println(userLog.ID)
	}
}

func TestSaveOrUpdate(t *testing.T) {
	userLog := UserLog{
		UserID: 1,
		Action: "new action",
	}

	result := db.Save(&userLog)
	assert.Nil(t, result.Error)

	userLog.UserID = 2
	result = db.Save(&userLog)
	assert.Nil(t, result.Error)
}

//	func TestConflict(t *testing.T) {
//		user := User{
//			ID: 88,
//			Name: Name{
//				FirstName: "User 88",
//			},
//		}
//
//		result := db.Clauses(clause.OnConflict{
//			UpdateAll: true,
//		}).Create(&user)
//
//		assert.Nil(t, result.Error)
//	}
func TestDelete(t *testing.T) {
	user := User{}
	result := db.Take(&user, "id = ?", 17)
	assert.Nil(t, result.Error)
	result = db.Delete(&user)
	assert.Nil(t, result.Error)

	result = db.Delete(&User{}, "id = ?", 16)
	assert.Nil(t, result.Error)

	result = db.Where("id = ?", 18).Delete(&User{})
	assert.Nil(t, result.Error)
}

func TestSoftDelete(t *testing.T) {
	todo := Todo{
		UserID:      1,
		Title:       "Todo i",
		Description: "isi todo 1",
	}

	result := db.Create(&todo)
	assert.Nil(t, result.Error)

	// attempt to softdelete
	result = db.Delete(&todo)
	assert.Nil(t, result.Error)
	assert.NotNil(t, todo.DeletedAt)

	var todos []Todo
	result = db.Find(&todos)

	assert.Nil(t, result.Error)
	assert.Equal(t, 0, len(todos))
}

func TestUnscope(t *testing.T) {
	var todo Todo

	result := db.Unscoped().Where("id = ?", 11).Find(&todo)
	assert.Nil(t, result.Error)
	fmt.Println(todo)

	result = db.Unscoped().Delete(&todo)
	assert.Nil(t, result.Error)

	var todos []Todo
	result = db.Unscoped().Find(&todos)
	assert.Nil(t, result.Error)
	assert.Equal(t, 0, len(todos))
}

func TestLock(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		var user User
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&user, "id = ?", 1).Error
		assert.Nil(t, err)
		if err != nil {
			return err
		}

		user.Name.FirstName = "Sakura"
		user.Name.LastName = "hasahi"
		return tx.Save(&user).Error
	})
	assert.Nil(t, err)
}
func TestCreateWallet(t *testing.T) {
	wallet := Wallet{
		ID:      "1",
		UserId:  1,
		Balance: 1000000,
	}
	err := db.Create(&wallet).Error
	assert.Nil(t, err)
}
func TestRetrieveRelation(t *testing.T) {
	var user User
	err := db.Model(&User{}).Preload("Wallet").First(&user).Error
	assert.Nil(t, err)

	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "1", user.Wallet.ID)
}

func TestRetrieveJoin(t *testing.T) {
	var user User
	err := db.Model(&User{}).Joins("Wallet").First(&user, "users.id = ?", 1).Error
	assert.Nil(t, err)

	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "1", user.Wallet.ID)
}

func TestAutoCreateUpdate(t *testing.T) {
	user := User{
		Password: "rahasia",
		Name: Name{
			FirstName: "User 16",
		},
		CreatedAt: time.Now().String(),
		Wallet: Wallet{
			ID:      "20",
			Balance: 1000000,
		},
	}

	err := db.Create(&user).Error
	assert.Nil(t, err)
}

func TestSkipAutoCreateUpdate(t *testing.T) {
	user := User{
		Password: "rahasia",
		Name: Name{
			FirstName: "User 17",
		},
		CreatedAt: time.Now().String(),
		Wallet: Wallet{
			ID:      "21",
			Balance: 1000000,
		},
	}

	err := db.Omit(clause.Associations).Create(&user).Error
	assert.Nil(t, err)
}

func TestUserAndAddresses(t *testing.T) {
	user := User{
		Password: "rahasia",
		Name: Name{
			FirstName: "sibondol",
		},
		CreatedAt: time.Now().String(),
		Wallet: Wallet{
			ID:      "50",
			Balance: 1000000,
		},
		Addresses: []Address{
			{
				Address: "jalan a",
			},
			{
				Address: "jalan b",
			},
		},
	}
	err := db.Create(&user).Error
	assert.Nil(t, err)
}

func TestPreloadJoinOneToMany(t *testing.T) {
	var user User
	err := db.Model(&User{}).Preload("Addresses").Joins("Wallet").Find(&user, "users.id = ?", 21).Error
	assert.Nil(t, err)
}

func TestBelongsTo(t *testing.T) {
	fmt.Println("Preload")
	var addresses []Address
	err := db.Model(&Address{}).Preload("User").Find(&addresses).Error
	assert.Nil(t, err)

	fmt.Println("joins")
	addresses = []Address{}
	err = db.Model(&Address{}).Joins("User").Find(&addresses).Error
	assert.Nil(t, err)
}

func TestBelongsToOneToOne(t *testing.T) {
	fmt.Println("preload")
	var wallets []Wallet
	err := db.Preload("User").Find(&wallets).Error
	assert.Nil(t, err)

	fmt.Println("Joins")
	wallets = []Wallet{}
	err = db.Joins("User").Find(&wallets).Error
	assert.Nil(t, err)
}

func TestCreateManyToMany(t *testing.T) {
	product := Product{
		ID:    "P001",
		Name:  "example product",
		Price: 20000,
	}
	err := db.Save(&product).Error
	assert.Nil(t, err)

	err = db.Table("user_like_product").Create(map[string]any{
		"user_id":    1,
		"product_id": "P001",
	}).Error
	assert.Nil(t, err)

	err = db.Table("user_like_product").Create(map[string]any{
		"user_id":    3,
		"product_id": "P001",
	}).Error
	assert.Nil(t, err)
}

func TestPreloadManyToMany(t *testing.T) {
	var product Product
	err := db.Preload("LikedByUsers").First(&product, "id = ?", "P001").Error
	assert.Nil(t, err)
	assert.Equal(t, 2, len(product.LikedByUsers))
}

func TestPreloadManyToManyUser(t *testing.T) {
	var user User
	err := db.Preload("LikeProducts").Take(&user, "id = ?", 1).Error
	assert.Nil(t, err)
	assert.Equal(t, 1, len(user.LikeProducts))
}

func TestAssociationFind(t *testing.T) {
	var product Product
	err := db.First(&product, "id = ?", "P001").Error
	assert.Nil(t, err)

	var users []User
	err = db.Model(&product).Where("users.first_name LIKE ?", "user%").Association("LikedByUsers").Find(&users)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
}

func TestAssociationAdd(t *testing.T) {
	var user User
	err := db.First(&user, "id = ?", 4).Error
	assert.Nil(t, err)

	var product Product
	err = db.First(&product, "id = ?", "P001").Error
	assert.Nil(t, err)

	err = db.Model(&product).Association("LikedByUsers").Append(&user)
	assert.Nil(t, err)
}

func TestAssociationReplace(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		var user User
		err := tx.Take(&user, "id = ?", 1).Error
		assert.Nil(t, err)
		fmt.Println(user.ID)
		wallet := Wallet{
			ID:      "01",
			UserId:  user.ID,
			Balance: 1000000,
		}
		err = tx.Model(&user).Association("Wallet").Replace(&wallet)
		return err
	})
	assert.Nil(t, err)
}

func TestAssociationDelete(t *testing.T) {
	var user User
	err := db.First(&user, "id = ?", 3).Error
	assert.Nil(t, err)

	var product Product
	err = db.First(&product, "id = ?", "P001").Error
	assert.Nil(t, err)

	err = db.Model(&product).Association("LikedByUsers").Delete(&user)
	assert.Nil(t, err)
}

func TestAssociationClear(t *testing.T) {
	var product Product
	err := db.Take(&product, "id = ?", "P001").Error
	assert.Nil(t, err)

	err = db.Model(&product).Association("LikedByUsers").Clear()
	assert.Nil(t, err)
}

func TestPreloadingWithCondition(t *testing.T) {
	var user User
	// take user with balance above 100000
	err := db.Preload("Wallet", "balance > ?", 1000000).First(&user, "id = ?", 1).Error
	assert.Nil(t, err)
}

func TestNestedPeloading(t *testing.T) {
	var wallet Wallet
	err := db.Preload("User.Addresses").Find(&wallet, "id = ?", "50").Error
	assert.Nil(t, err)

	fmt.Println(wallet)
}
func TestPreloadAll(t *testing.T) {
	var user User
	err := db.Preload(clause.Associations).First(&user, "id = ?", 1).Error
	assert.Nil(t, err)
}

func TestJoinQuery(t *testing.T) {
	var users []User
	err := db.Joins("join wallets on wallets.user_id = users.id").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 3, len(users))

	users = []User{}
	err = db.Joins("Wallet").Find(&users).Error //left join
	assert.Nil(t, err)
	assert.Equal(t, 17, len(users))
}

func TestJoinQueryCondition(t *testing.T) {
	var users []User
	err := db.Joins("join wallets on wallets.user_id = users.id AND wallets.balance > ?", 500000).Find(&users).Error
	assert.Nil(t, err)

	users = []User{}
	err = db.Joins("Wallet").Where("Wallet.balance > ?", 500000).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 3, len(users))
}

func TestCount(t *testing.T) {
	var count int64
	err := db.Model(&User{}).Joins("Wallet").Where("Wallet.balance > ?", 500000).Count(&count).Error
	assert.Nil(t, err)
	assert.Equal(t, int64(3), count)
}

type AggregationResult struct {
	TotalBalance int64
	MinBalance   int64
	MaxBalance   int64
	AvgBalance   float64
}

func TestAggregation(t *testing.T) {
	var result AggregationResult
	err := db.Model(&Wallet{}).Select("sum(balance) as total_balance", "min(balance) as min_balance", "max(balance) as max_balance", "avg(balance) as avg_balance").Take(&result).Error
	assert.Nil(t, err)

	assert.Equal(t, int64(3000000), result.TotalBalance)
	assert.Equal(t, int64(1000000), result.MinBalance)
	assert.Equal(t, int64(1000000), result.MaxBalance)
	assert.Equal(t, float64(1000000), result.AvgBalance)
}

// failed
func TestGroupByHaving(t *testing.T) {
	var results []AggregationResult
	err := db.Model(&Wallet{}).Select("sum(balance) as total_balance", "min(balance) as min_balance", "max(balance) as max_balance", "avg(balance) as avg_balance").Joins("join users on users.id = wallets.user_id").Group("users.id").Having("sum(balance) > ?", 100000).Take(&results).Error
	assert.Nil(t, err)
	assert.Equal(t, 0, len(results))
}

func TestContext(t *testing.T) {
	ctx := context.Background()

	var users []User
	err := db.WithContext(ctx).Find(&users).Error
	assert.Nil(t, err)
}

func BrokeWalletBalance(db *gorm.DB) *gorm.DB {
	return db.Where("balance = ? ", 0)
}

func SultanWalletBalance(db *gorm.DB) *gorm.DB {
	return db.Where("balance > ?", 10000000)
}

func TestScope(t *testing.T) {
	var wallets []Wallet
	err := db.Scopes(BrokeWalletBalance).Find(&wallets).Error
	assert.Nil(t, err)

	wallets = []Wallet{}
	err = db.Scopes(SultanWalletBalance).Find(&wallets).Error
	assert.Nil(t, err)
}

func TestMigrator(t *testing.T) {
	err := db.Migrator().AutoMigrate(&Guestbook{})
	assert.Nil(t, err)
}
