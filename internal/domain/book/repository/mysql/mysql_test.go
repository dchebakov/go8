package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	//_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null/v8"

	"github.com/gmhafiz/go8/cmd/extmigrate/migrate"
	"github.com/gmhafiz/go8/internal/domain/book"
	"github.com/gmhafiz/go8/internal/models"
	"github.com/gmhafiz/go8/internal/utility/filter"
)

//go:generate mockgen -package mock -source ../../repository.go -destination=../../mock/mock_repository.go

var (
	repo book.Repository
)

func TestMain(m *testing.M) {
	// must go back to project's root path to get to the .env and ./database/migrations/ folder
	to := "../../../../../"
	err := os.Chdir(to)
	if err != nil {
		log.Fatalln(err)
	}
	path, _ := os.Getwd()
	fmt.Println(path)
	err = godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}
	//cfg := configs.DockerTestCfg()

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("mysql", "5.7", []string{"MYSQL_ROOT_PASSWORD=secret"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err = pool.Retry(func() error {
		db, err := sqlx.Open("mysql", fmt.Sprintf("root:secret@(localhost:%s)/mysql",
			resource.GetPort("3306/tcp")))
		if err != nil {
			return err
		}
		repo = New(db)
		return db.Ping()
	}); err != nil {
		log.Fatalf("could not connect to docker: %s", err.Error())
	}

	code := m.Run()

	os.Exit(code)
}

func TestBookRepository_Create(t *testing.T) {
	migrate.Start()

	dt := "2020-01-01T15:04:05Z"
	timeWant, err := time.Parse(time.RFC3339, dt)
	if err != nil {
		t.Fatal(err)
	}
	bookTest := &models.Book{
		Title:         "test1",
		PublishedDate: timeWant,
		Description:   "test1",
		ImageURL: null.String{
			String: "http://example.com/image.png",
			Valid:  true,
		},
	}

	bookID, err := repo.Create(context.Background(), bookTest)

	assert.NoError(t, err)
	assert.NotEqual(t, 0, bookID)

	migrate.Down()
}

func TestRepository_Find(t *testing.T) {
	migrate.Start()

	dt := "2020-01-01T15:04:05Z"
	timeWant, err := time.Parse(time.RFC3339, dt)
	if err != nil {
		t.Fatal(err)
	}
	bookWant := &models.Book{
		Title:         "test2",
		PublishedDate: timeWant,
		Description:   "test2",
	}
	bookID, err := repo.Create(context.Background(), bookWant)
	if err != nil {
		t.Fatal(err)
	}

	bookGot, err := repo.Read(context.Background(), bookID)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, bookGot.Title, bookWant.Title)
	assert.Equal(t, bookGot.Description, bookWant.Description)
	assert.Equal(t, bookGot.PublishedDate.String(), bookWant.PublishedDate.String())

	migrate.Down()
}

func TestRepository_List(t *testing.T) {
	migrate.Start()

	ctx := context.Background()

	f := &book.Filter{
		Base: filter.Filter{
			Page:   1,
			Size:   10,
			Search: false,
		},
		Title:         "",
		Description:   "",
		PublishedDate: "",
	}
	books, err := repo.List(ctx, f)

	assert.NoError(t, err)
	assert.Len(t, books, 0)

	migrate.Down()
}

func TestRepository_Update(t *testing.T) {
	migrate.Start()

	ctx := context.Background()
	dt := "2020-01-01T15:04:05Z"
	timeWant, err := time.Parse(time.RFC3339, dt)
	if err != nil {
		assert.NoError(t, err)
	}

	want := &models.Book{
		BookID:        1,
		Title:         "updated title 1",
		PublishedDate: timeWant,
		ImageURL: null.String{
			String: "http://example.com/image2.png",
			Valid:  true,
		},
		Description: "updated description",
	}

	err = repo.Update(ctx, want)

	assert.NoError(t, err)

	migrate.Down()
}

func TestRepository_Delete(t *testing.T) {
	migrate.Start()

	ctx := context.Background()

	err := repo.Delete(ctx, 1)

	assert.NoError(t, err)

	got, err := repo.Read(ctx, 1)

	assert.Nil(t, got)
	assert.Error(t, err, sql.ErrNoRows)

	migrate.Down()
}

func TestRepository_Search(t *testing.T) {
	migrate.Start()

	ctx := context.Background()

	re := book.Filter{
		Base:          filter.Filter{},
		Title:         "test2",
		Description:   "",
		PublishedDate: "",
	}

	got, err := repo.Search(ctx, &re)

	assert.NoError(t, err)
	assert.Len(t, got, 0)

	migrate.Down()
}
