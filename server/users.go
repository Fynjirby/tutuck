package main

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"username"`
	Key  string `json:"key"`
}

type UserStore struct {
	Users []User `json:"users"`
}

var (
	userStore = UserStore{Users: []User{}}
	userLock  sync.Mutex
)

func loadUsers() {
	data, err := ioutil.ReadFile("users.json")
	if err == nil {
		_ = json.Unmarshal(data, &userStore)
	}
}

func saveUsers() {
	data, _ := json.MarshalIndent(userStore, "", "  ")
	_ = ioutil.WriteFile("users.json", data, 0644)
}

func nextUID() int {
	max := 0
	for _, u := range userStore.Users {
		if u.ID > max {
			max = u.ID
		}
	}
	return max + 1
}
