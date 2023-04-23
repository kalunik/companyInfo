package internal

import "sync"

type CompanyInfo struct {
	TaxId   int
	KPP     int
	Name    string
	CeoName string
}

type CompanyInfoStore struct {
	sync.RWMutex
	info map[int]CompanyInfo

	infoId int
}

func New() *CompanyInfoStore {
	cis := &CompanyInfoStore{}
	cis.info = make(map[int]CompanyInfo)
	cis.infoId = 0
	return cis
}

func AddNewInfo() {
}
