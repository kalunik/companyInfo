package rest

import (
	"bytes"
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

func (e *requestError) ErrorByCode() string {
	switch e.errorCode {
	case wrongLenTaxId:
		return "Requested taxId should have length between 10 and 12"
	case invalidTaxId:
		return "Requested taxId not found. Is it correct?"
	}
	return ""
}

func (e *requestError) ErrorJson() []byte {
	jsonResponse := &bytes.Buffer{}
	jsonResponse.WriteString(`{"errorCode": "`)
	jsonResponse.WriteString(strconv.Itoa(e.errorCode))
	jsonResponse.WriteString(`", "errorDetails" : "`)
	jsonResponse.WriteString(e.ErrorByCode())
	jsonResponse.WriteString(`"}`)
	return jsonResponse.Bytes()
}
