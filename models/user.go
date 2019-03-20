package models

import(
	"challenge/db"
	"log"
)


type User struct{
	Name string `json:"name"`
	Value string `json:"value"`
}
type Users struct{
	element []User
}


func (h User) GetByName(name string) (*User, error){
	db := db.GetDB()
	value, ok := db[name]
	log.Println("......")
	log.Println(value)
	user := new(User)
	user.Name = name
	user.Value = value
	if(ok){
		return user, nil
	} else{
		return nil, nil
	}

}

func(h User) GetAll()(*map[string]string, error){
	db := db.GetDB()
	
	return &db, nil
}