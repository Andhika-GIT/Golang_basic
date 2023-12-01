package json

type User struct {
	Firstname string
	Lastname  string
	Skills    []string
	Addresses []Address
}

type Customer struct {
	Firstname string
	Lastname  string
	Id        int64
}

type Address struct {
	Street  string
	Country string
}
