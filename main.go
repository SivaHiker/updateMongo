package main

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"bytes"
)

func main() {

	var counter int

	session, err := mgo.Dial("10.15.0.145")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("userlist").C("mqttuserdata")
	fmt.Println(c.Name)

	result1 := []UserInfo{}
	err = c.Find(bson.M{"flag": true}).All(&result1)

	for _, v := range result1 {
		fmt.Println(v.ID)
		_ = c.Update(v, bson.M{"$set": bson.M{"flag": false}})
		counter = counter +1
	}
	fmt.Println("Total Fetched Documents ",len(result1))
	fmt.Println("Total Updated Documents ",counter)
}

type UserInfo struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	UserData UserData `json:"UserData"`
	Flag bool `json:"flag"`
}

type UserData struct {
	Msisdn string `json:"msisdn"`
	Token  string `json:"token"`
	UID    string `json:"uid"`
}