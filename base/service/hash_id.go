package service

import (
	"github.com/speps/go-hashids"
	"log"
)

type IDService interface {
	EncodeID(id int) (string, error)
	DecodeID(id string) (int, error)
}

type idService struct {
	hashID *hashids.HashID
}

func NewIDService(salt string) IDService {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = 30
	h, err := hashids.NewWithData(hd)
	if err != nil {
		log.Fatal(err)
	}
	return &idService{hashID: h}
}

func (s idService) EncodeID(id int) (string, error) {
	return s.hashID.Encode([]int{id})
}

func (s idService) DecodeID(id string) (int, error) {
	ids, err := s.hashID.DecodeWithError(id)
	if err != nil {
		return 0, err
	}
	return ids[0], nil
}
