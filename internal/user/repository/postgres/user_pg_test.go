// internal/user/repository/postgres
package postgres_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"time"

	"github.com/joho/godotenv"
	"github.com/rezalaal/coral/internal/db"
	"github.com/rezalaal/coral/internal/user/models"
	"github.com/rezalaal/coral/internal/user/repository/postgres"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
    rootPath, _ := filepath.Abs(filepath.Join("..", "..", "..", "..", ".env"))
	err := godotenv.Load(rootPath)
    if err != nil {
        fmt.Println("⚠️  Failed to load .env:", err)
    } else {
        fmt.Println("✅ .env loaded")
    }

    code := m.Run()
    os.Exit(code)
}



func TestUserPG_CreateAndGet(t *testing.T) {
	database, err := db.Connect()
	assert.NoError(t, err)
	defer database.Close()

	repo := postgres.NewUserPG(database)

	user := &models.User{
		Name:         "کاربر تستی",
		Mobile:       fmt.Sprintf("091200000%04d", time.Now().UnixNano()%10000), // موبایل یکتا
		PasswordHash: "hashed_password",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err = repo.Create(user)
	assert.NoError(t, err)
	assert.NotZero(t, user.ID)

	defer func() {
		_, err := database.Exec("DELETE FROM users WHERE id = $1", user.ID)
		assert.NoError(t, err)
	}()

	userFromDB, err := repo.GetByID(user.ID)
	assert.NoError(t, err)
	assert.NotNil(t, userFromDB)

	assert.Equal(t, user.Name, userFromDB.Name)
	assert.Equal(t, user.Mobile, userFromDB.Mobile)
	assert.Equal(t, user.PasswordHash, userFromDB.PasswordHash)
}

func TestUserPG_Create(t *testing.T) {
    database, err := db.Connect()
    assert.NoError(t, err)
    defer database.Close()

    userRepo := postgres.NewUserPG(database)
    
    user := &models.User{
        Name:         "Test User",
        Mobile:       fmt.Sprintf("091200000%04d", time.Now().UnixNano()%10000),
        PasswordHash: "hashedpassword",
        CreatedAt:    time.Now(),
        UpdatedAt:    time.Now(),
    }

    err = userRepo.Create(user)
    assert.NoError(t, err)
    assert.NotZero(t, user.ID)

    // حذف رکورد ساخته شده
    defer func() {
        _, err := database.Exec("DELETE FROM users WHERE id = $1", user.ID)
        assert.NoError(t, err)
    }()
}


func TestUserPG_List(t *testing.T) {
    database, err := db.Connect()
    assert.NoError(t, err)
    defer database.Close()

    userRepo := postgres.NewUserPG(database)

    // ساخت دو کاربر جدید با شماره موبایل یکتا
    usersToCreate := []*models.User{
        {
            Name:         "User One",
            Mobile:       fmt.Sprintf("091200000%04d", time.Now().UnixNano()%10000),
            PasswordHash: "hashedpassword1",
            CreatedAt:    time.Now(),
            UpdatedAt:    time.Now(),
        },
        {
            Name:         "User Two",
            Mobile:       fmt.Sprintf("091200000%04d", time.Now().UnixNano()%10000 + 1),
            PasswordHash: "hashedpassword2",
            CreatedAt:    time.Now(),
            UpdatedAt:    time.Now(),
        },
    }

    for _, u := range usersToCreate {
        err := userRepo.Create(u)
        assert.NoError(t, err)
        assert.NotZero(t, u.ID)
        defer func(id int64) {
            _, err := database.Exec("DELETE FROM users WHERE id = $1", id)
            assert.NoError(t, err)
        }(u.ID)
    }

    // اجرای متد List
    users, err := userRepo.List()
    assert.NoError(t, err)

    // بررسی اینکه کاربران ساخته شده در لیست وجود دارند
    foundCount := 0
    for _, u := range users {
        for _, created := range usersToCreate {
            if u.ID == created.ID {
                foundCount++
            }
        }
    }

    assert.Equal(t, len(usersToCreate), foundCount, "تمام کاربران ساخته شده باید در لیست باشند")
}


func TestUserPG_GetByMobile(t *testing.T) {
    database, err := db.Connect()
    assert.NoError(t, err)
    defer database.Close()

    userRepo := postgres.NewUserPG(database)

    mobile := fmt.Sprintf("091200000%04d", time.Now().UnixNano()%10000)

    user := &models.User{
        Name:         "Test Mobile User",
        Mobile:       mobile,
        PasswordHash: "hashedpassword",
        CreatedAt:    time.Now(),
        UpdatedAt:    time.Now(),
    }

    // ایجاد کاربر
    err = userRepo.Create(user)
    assert.NoError(t, err)
    assert.NotZero(t, user.ID)

    defer func(id int64) {
        _, err := database.Exec("DELETE FROM users WHERE id = $1", id)
        assert.NoError(t, err)
    }(user.ID)

    // فراخوانی GetByMobile
    userFromDB, err := userRepo.GetByMobile(mobile)
    assert.NoError(t, err)
    assert.NotNil(t, userFromDB)
    assert.Equal(t, user.ID, userFromDB.ID)
    assert.Equal(t, user.Name, userFromDB.Name)
    assert.Equal(t, user.Mobile, userFromDB.Mobile)
}

func TestUserPG_Delete(t *testing.T) {
    database, err := db.Connect()
    assert.NoError(t, err)
    defer database.Close()

    userRepo := postgres.NewUserPG(database)

    // یک کاربر جدید ایجاد می‌کنیم
    user := &models.User{
        Name:         "User To Delete",
        Mobile:       fmt.Sprintf("091200000%04d", time.Now().UnixNano()%10000),
        PasswordHash: "hashedpassword",
        CreatedAt:    time.Now(),
        UpdatedAt:    time.Now(),
    }

    err = userRepo.Create(user)
    assert.NoError(t, err)
    assert.NotZero(t, user.ID)

    // حذف کاربر
    err = userRepo.Delete(user.ID)
    assert.NoError(t, err)

    // تلاش برای دریافت کاربر حذف شده باید نتیجه nil بدهد
    deletedUser, err := userRepo.GetByID(user.ID)
    assert.NoError(t, err)
    assert.Nil(t, deletedUser)
}

