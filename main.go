package main

import (
	"cash/cache"
	"fmt"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42)
	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId = cache.Get("userId")
	fastString := cache.Get("fastString")

	fmt.Println(userId)
	fmt.Println(fastString)
}
