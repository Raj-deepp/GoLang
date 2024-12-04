package main

import (
	"errors"
	// "fmt"
)

// Maps
func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	userMap := make(map[string]user)
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("Invalid Sizes")
	}
	for i := 0; i < len(names); i++ {
		name := names[i]
		phoneNumber := phoneNumbers[i]
		userMap[name] = user{
			name:        name,
			phoneNumber: phoneNumber,
		}
	}
	return userMap, nil
}

type user struct {
	name        string
	phoneNumber int
} //------------------------------------------------------------------------------------------------

//Mutations
// Insert an element
// m[key] = elem

// Get an element
// elem = m[key]

// Delete an element
// delete(m, key)

// Check if a key exists
// elem, ok := m[key]

func deleteIfNecessary(users map[string]user2, name string) (deleted bool, err error) {
	existingUser, ok := users[name]
	if !ok {
		return false, errors.New("Not Found")
	}
	if existingUser.scheduledForDeletion {
		delete(users, name)
		return true, nil
	}
	return false, nil
}

type user2 struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

//-------------------------------------------------------------------------------------------------

//Count Instances

func getCounts(messagedUsers []string, validUsers map[string]int) {
	for _, username := range messagedUsers {
		if _, exists := validUsers[username]; exists {
			validUsers[username]++
		}
	}
}

//-------------------------------------------------------------------------------------------------

//Nested Maps

func getNameCounts(names []string) map[rune]map[string]int {
	counts := make(map[rune]map[string]int)
	for _, name := range names {
		if name == "" {
			continue
		}
		firstChar := rune(name[0])
		_, ok := counts[firstChar]
		if !ok {
			counts[firstChar] = make(map[string]int)
		}
		counts[firstChar][name]++
	}
	return counts
}

func main() {

}
