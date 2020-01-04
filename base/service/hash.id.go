package service

import (
	"github.com/speps/go-hashids"
	"log"
)

var hashID *hashids.HashID

func Init(salt string) {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = 30
	h, err := hashids.NewWithData(hd)
	if err != nil {
		log.Fatal(err)
	}
	hashID = h
}

func EncodeID(id int) (string, error) {
	return hashID.Encode([]int{id})
}

func DecodeID(id string) (int, error) {
	ids, err := hashID.DecodeWithError(id)
	if err != nil {
		return 0, err
	}
	return ids[0], nil
}
