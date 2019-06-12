package main

import (
	"errors"

	"github.com/asdine/storm"
	"gopkg.in/mgo.v2/bson"
)

//user holds data for single user
type User struct {
	ID   bson.ObjectId `json:"id" storm:"id"`
	Name string        `json:"name"`
	Role string        `json:"role`
}

const (
	dbPath = "users.db"
)

//errors
var (
	ErrRecordInvalid = errors.New("record is invalid")
)

//All returns all users from the database
//Provide the information from the database
func All() ([]User, error) {
	//open connection
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	//close the db connection to avoid dataleaks
	defer db.Close()

	//take all users
	users := []User{}
	err = db.All(&users)
	if err != nil {
		return nil, err
	}
	//if no error return users and nil
	return users, nil
}

//One returns a single user record from the database
func One(id bson.ObjectId) (*User, error) {
	//open connection
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	//close the db connection to avoid dataleaks
	defer db.Close()

	u := new(User)            //new keyword return the pointer to the structure
	err = db.One("ID", id, u) //(Field cointaining unique id, value of that field, pointer to structure where it is stored)
	if err != nil {
		return nil, err
	}
	//if no error return users and nil
	return u, nil
}

//Delete removes a given recoed froma  database
func Delete(id bson.ObjectId) error {
	//open connection
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	//close the db connection to avoid dataleaks
	defer db.Close()

	u := new(User)            //new keyword return the pointer to the structure
	err = db.One("ID", id, u) //(Field cointaining unique id, value of that field, pointer to structure where it is stored)
	if err != nil {
		return err
	}
	//this will delete the given record provided in struct
	return db.DeleteStruct(u)
}

//Creating and updating a record
func (u *User) Save() error {

	if err := u.validate(); err != nil { // this is like first condition will be checked and then error fucntion will be checked for validation
		return err
	}
	//open connection
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	//close the db connection to avoid dataleaks
	defer db.Close()
	return db.Save(u)
}

//validate makes sure that the record contains valid data
func (u *User) validate() error {
	if u.Name == "" {
		return ErrRecordInvalid
	}
	return nil
}
func main() {

}
