package company

import (
	"fmt"
	"github.com/kalunik/companyInfo/internal/parse"
	"golang.org/x/net/context"
)

type CompanyGRPC struct {
	UnimplementedCompanyServer
}

func (*CompanyGRPC) GetInfo(ctx context.Context, req *Request) (*Response, error) {

	info := parse.ScrapeRusprofile(req.GetTaxId())
	fmt.Println(info)

	return &Response{
		TaxId:   info.TaxId,
		Kpp:     info.KPP,
		Name:    info.Name,
		CeoName: info.CeoName,
	}, nil
}
