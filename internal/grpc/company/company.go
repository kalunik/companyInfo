package company

import (
	"fmt"
	gen "github.com/kalunik/companyInfo/internal/grpc/proto/generated"
	"github.com/kalunik/companyInfo/internal/parse"
	"golang.org/x/net/context"
)

type CompanyGRPC struct {
	gen.UnimplementedCompanyServer
}

func (*CompanyGRPC) GetInfo(ctx context.Context, req *gen.Request) (*gen.Response, error) {
	info := parse.ScrapeRusprofile(req.GetTaxId())
	fmt.Println(info)

	return &gen.Response{
		TaxId:   info.TaxId,
		Kpp:     info.KPP,
		Name:    info.Name,
		CeoName: info.CeoName,
	}, nil
}
