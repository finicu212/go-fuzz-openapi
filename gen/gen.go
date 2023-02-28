package gen

type Order struct {
	Petid    int32
	Quantity int32
	Shipdate string
	Status   string
	Complete interface{} //TODO: Handle others
	Id       int32
}

type Pet struct {
	Category  Category
	Id        int32
	Name      string
	Photourls []string
	Status    string
	Tags      []Tag
}

type Tag struct {
	Id   int32
	Name string
}

type User struct {
	Id         int32
	Lastname   string
	Password   string
	Phone      string
	Userstatus int32
	Username   string
	Email      string
	Firstname  string
}

type Address struct {
	Street string
	Zip    string
	City   string
	State  string
}

type ApiResponse struct {
	Code    int32
	Message string
	Type    string
}

type Category struct {
	Id   int32
	Name string
}

type Customer struct {
	Address  []Address
	Id       int32
	Username string
}
