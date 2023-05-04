package rest

import (
	"bytes"
	"errors"
	"strconv"
)

const (
	wrongLenTaxId int = iota
	invalidTaxId
)

type requestError struct {
	errorCode int
	Err       error
}

func (e requestError) Error() string {
	result := &bytes.Buffer{}
	result.WriteString(e.Err.Error())
	return result.String()
}

func (e *requestError) setError() {
	switch e.errorCode {
	case wrongLenTaxId:
		e.Err = errors.New("Requested taxId should have length between 10 and 12")
	case invalidTaxId:
		e.Err = errors.New("Requested taxId not found. Is it correct?")
	}
}

func (e *requestError) ErrorJson() []byte {
	jsonResponse := &bytes.Buffer{}
	jsonResponse.WriteString(`{"errorCode": "`)
	jsonResponse.WriteString(strconv.Itoa(e.errorCode))
	jsonResponse.WriteString(`", "errorDetails" : "`)
	e.setError()
	jsonResponse.WriteString(e.Error())
	jsonResponse.WriteString(`"}`)
	return jsonResponse.Bytes()
}
