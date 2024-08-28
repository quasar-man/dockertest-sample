package repository_test

 import (
	"testing"
	"log"
	"fmt"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"gorm.io/gorm"
  "gorm.io/driver/mysql"

	"github.com/quasar-man/dockertest-sample/repository"
	"github.com/quasar-man/dockertest-sample/entity"
 )

 var db *gorm.DB

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{	
		Repository: "mysql",
		Tag: "8.0",
		Env: []string{"MYSQL_ROOT_PASSWORD=secret", "MYSQL_DATABASE=dockertest_db"},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// コンテナの削除をスケジュール
	defer func() {
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}()

	// コンテナの準備が整うまで待機
	if err := pool.Retry(func() error {
		dsn := fmt.Sprintf("root:secret@(localhost:%s)/dockertest_db?charset=utf8mb4&parseTime=True&loc=Local", resource.GetPort("3306/tcp"))
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return err
		}
		return db.AutoMigrate(&entity.User{})
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// 事前にデータを作成
	createUser()

	m.Run()
}

func TestFindAll(t *testing.T) {
	userRepository := repository.NewUserRepository(db)
	users, err := userRepository.FindAll()
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if len(*users) != 4 {
		t.Errorf("Expected: 4, but got: %d", len(*users))
	}
}

func TestFindByID(t *testing.T) {
	userRepository := repository.NewUserRepository(db)
	user, err := userRepository.FindByID(1)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if user.ID != 1 {
		t.Errorf("Expected: 1, but got: %d", user.ID)
	}
}

func TestFindByEmail(t *testing.T) {
	userRepository := repository.NewUserRepository(db)
	user, err := userRepository.FindByEmail("test3@samplemail.com")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if user.Email != "test3@samplemail.com" {
		t.Errorf("Expected: test3@samplemail.com but got: %s", user.Email)
	}
}

func createUser() {
	users := []*entity.User{
		{ ID: 1, Name: "test user1", Email: "test1@samplemail.com", Password: "password1" },
		{ ID: 2, Name: "test user2", Email: "test2@samplemail.com", Password: "password2" },
		{ ID: 3, Name: "test user3", Email: "test3@samplemail.com", Password: "password3" },
		{ ID: 4, Name: "test user4", Email: "test4@samplemail.com", Password: "password4" },
	}

	db.Create(users)
}
