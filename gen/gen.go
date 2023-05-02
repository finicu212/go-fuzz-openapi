package gen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"log"
	"net/http"
	"strings"
	"testing"
)

/*
 * Automatically generated file...
 * TODO: Figure out an elegant way to generate this before running `loadtest`.
 * TODO: i.e. running fuzzctl should also update this file based on the spec
 *
 */

func init() {
}

type Customer struct {
	Username string
	Address  []Address `fakesize:"2"`
	Id       int32
}

// FakeCustomer returns a faked struct of Customer type
func FakeCustomer() Customer {
	var f Customer
	gofakeit.Struct(&f)
	return f
}

type Order struct {
	Shipdate string
	Status   string
	Complete bool
	Id       int32
	Petid    int32
	Quantity int32
}

// FakeOrder returns a faked struct of Order type
func FakeOrder() Order {
	var f Order
	gofakeit.Struct(&f)
	return f
}

type Pet struct {
	Name      string
	Photourls []string `fakesize:"2"`
	Status    string
	Tags      []Tag `fakesize:"2"`
	Category  Category
	Id        int32
}

// FakePet returns a faked struct of Pet type
func FakePet() Pet {
	var f Pet
	gofakeit.Struct(&f)
	return f
}

type Tag struct {
	Id   int32
	Name string
}

// FakeTag returns a faked struct of Tag type
func FakeTag() Tag {
	var f Tag
	gofakeit.Struct(&f)
	return f
}

type User struct {
	Phone      string
	Userstatus int32
	Username   string
	Email      string
	Firstname  string
	Id         int32
	Lastname   string
	Password   string
}

// FakeUser returns a faked struct of User type
func FakeUser() User {
	var f User
	gofakeit.Struct(&f)
	return f
}

type Address struct {
	City   string
	State  string
	Street string
	Zip    string
}

// FakeAddress returns a faked struct of Address type
func FakeAddress() Address {
	var f Address
	gofakeit.Struct(&f)
	return f
}

type ApiResponse struct {
	Code    int32
	Message string
	Type    string
}

// FakeApiResponse returns a faked struct of ApiResponse type
func FakeApiResponse() ApiResponse {
	var f ApiResponse
	gofakeit.Struct(&f)
	return f
}

type Category struct {
	Id   int32
	Name string
}

// FakeCategory returns a faked struct of Category type
func FakeCategory() Category {
	var f Category
	gofakeit.Struct(&f)
	return f
}

// GetEmptyStructByName returns a zero-value struct for the given name. For example, "pet" returns an empty Pet{} object.
func GetEmptyStructByName(name string) interface{} {
	switch strings.ToLower(name) {
	case "customer":
		return Customer{}
	case "order":
		return Order{}
	case "pet":
		return Pet{}
	case "tag":
		return Tag{}
	case "user":
		return User{}
	case "address":
		return Address{}
	case "apiresponse":
		return ApiResponse{}
	case "category":
		return Category{}

	}
	return nil
}

// GetFakedStructByName returns a struct filled with faked values for the given name.
// We can't apply gofakeit by name using this func, since we require type casting from interface{} to the concrete type. We don't know the concrete types ahead of time
// (i.e. gofakeit(&interface{}) makes no modifications to the empty interface)
func GetFakedStructByName(name string) interface{} {
	switch strings.ToLower(name) {
	case "customer":
		return FakeCustomer()
	case "order":
		return FakeOrder()
	case "pet":
		return FakePet()
	case "tag":
		return FakeTag()
	case "user":
		return FakeUser()
	case "address":
		return FakeAddress()
	case "apiresponse":
		return FakeApiResponse()
	case "category":
		return FakeCategory()

	}
	return nil
}

func encodeAndSend(v any, method, endpoint string) (*http.Response, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(v)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(method, fmt.Sprintf("https://petstore3.swagger.io/api/v3/%s", endpoint), &buf)
	if err != nil {
		log.Fatalf("Failed creating request with data: %+v", buf)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(req)
}

// DELETE /store/order/
func Fuzz_deleteOrder_OnlyRequired(f *testing.F) {
	f.Skip()

}

// GET /store/order/
func Fuzz_getOrderById_OnlyRequired(f *testing.F) {
	f.Skip()

	// Response Body: Order
}

// POST /user
func Fuzz_createUser_OnlyRequired(f *testing.F) {
	f.Skip()
	// Request Body: User

}

// POST /user/createWithList
func Fuzz_createUsersWithListInput_OnlyRequired(f *testing.F) {
	f.Skip()
	// Request Body: .
	// Response Body: User
}

// GET /user/login
func Fuzz_loginUser_OnlyRequired(f *testing.F) {
	f.Skip()

	// Response Body: .
}

// POST /pet
func Fuzz_addPet_OnlyRequired(f *testing.F) {
	f.Skip()
	// Request Body: Pet
	// Response Body: Pet
}

// PUT /pet
func Fuzz_updatePet_OnlyRequired(f *testing.F) {
	f.Skip()
	// Request Body: Pet
	// Response Body: Pet
}

// GET /pet/findByStatus
func Fuzz_findPetsByStatus_OnlyRequired(f *testing.F) {
	f.Skip()

	// Response Body: .
}

// GET /pet/findByTags
func Fuzz_findPetsByTags_OnlyRequired(f *testing.F) {
	f.Skip()

	// Response Body: .
}

// POST /store/order
func Fuzz_placeOrder_OnlyRequired(f *testing.F) {
	f.Skip()
	// Request Body: Order
	// Response Body: Order
}

// DELETE /user/
func Fuzz_deleteUser_OnlyRequired(f *testing.F) {
	f.Skip()

}

// GET /user/
func Fuzz_getUserByName_OnlyRequired(f *testing.F) {
	f.Skip()

	// Response Body: User
}

// PUT /user/
func Fuzz_updateUser_OnlyRequired(f *testing.F) {
	f.Skip()
	// Request Body: User

}

// POST /pet/
func Fuzz_updatePetWithForm_OnlyRequired(f *testing.F) {
	f.Skip()

}

// DELETE /pet/
func Fuzz_deletePet_OnlyRequired(f *testing.F) {
	f.Skip()

}

// GET /pet/
func Fuzz_getPetById_OnlyRequired(f *testing.F) {
	f.Skip()

	// Response Body: Pet
}

// POST /pet/
func Fuzz_uploadFile_OnlyRequired(f *testing.F) {
	f.Skip()

	// Response Body: ApiResponse
}

// GET /store/inventory
func Fuzz_getInventory_OnlyRequired(f *testing.F) {
	f.Skip()

	// Response Body: .
}

// GET /user/logout
func Fuzz_logoutUser_OnlyRequired(f *testing.F) {
	f.Skip()

}
