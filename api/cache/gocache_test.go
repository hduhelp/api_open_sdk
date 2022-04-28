package cache

import (
	"context"
	"fmt"
	"testing"
	"time"
)

type BookQuery struct {
	Slug string
}

type Book struct {
	ID   int
	Name string
	Slug string
}

func TestCache(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()
	SetUp(DefaultMemCache())
	key := BookQuery{Slug: "my-test-amazing-book"}
	value := Book{ID: 1, Name: "My test amazing book", Slug: "my-test-amazing-book"}

	err := marshal.Set(ctx, key, value, nil)
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
	d, err := marshal.Get(ctx, key, new(Book))
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
}
