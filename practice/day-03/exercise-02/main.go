package main

func main() {
}

type User struct {
	Name       string
	Department string
}

func groupByDepartment(users []User) map[string][]string {
	mmap := make(map[string][]string)

	if len(users) == 0 {
		return mmap
	}

	for _, u := range users {
		mmap[u.Department] = append(mmap[u.Department], u.Name)
	}

	return mmap
}
