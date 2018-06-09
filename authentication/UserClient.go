package main

import (
	"net/http"
	"github.com/DiTo04/spexflix/common/codecs"
)

type userClient struct {
	serviceAdress string
}

func (c *userClient) GetUserId(email string) (int64, error) {
	rsp, err := http.Get(c.serviceAdress + "/id?email=" + email)
	if err != nil {
		return -1, err
	}
	defer rsp.Body.Close()
	var id int64
	err = codecs.JSON.Decode(rsp.Body, &id)
	if err != nil {
		return -1, err
	}
	return id, nil
}
