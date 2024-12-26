package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Person struct{
	Name  string `json:"name"`
	Age   int     `json:"age"`
	Role  string   `json:"role"`
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	fmt.Print(ping)
	

	jsonResult,err := json.Marshal(Person{
			Name: "Abdelatty Badwy",
			Age : 20,
			Role: "Software Enginner",
		})
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	err = client.Set(context.Background(), "Data", jsonResult, 0).Err()
	if err != nil {
		fmt.Print("Faild to set value to redis")
		return
	}

	val, err := client.Get(context.Background(), "Data").Result()
	if err != nil {
		fmt.Printf("faild to get value from redis")
	}
	fmt.Printf(val)
}
