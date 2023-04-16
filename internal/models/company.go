package models

type Company struct {
	UlCount int           `json:"ul_count"`
	Ul      []CompanyInfo `json:"ul"`
	IpCount int           `json:"ip_count"`
	Success bool          `json:"success"`
	Code    int           `json:"code"`
	Message string        `json:"message"`
}

type CompanyInfo struct {
	Name              string      `json:"name"`
	RawName           string      `json:"raw_name"`
	ManyCeo           int         `json:"many_ceo"`
	Link              string      `json:"link"`
	Ogrn              string      `json:"ogrn"`
	RawOgrn           string      `json:"raw_ogrn"`
	Inn               string      `json:"inn"`
	Region            string      `json:"region"`
	Address           string      `json:"address"`
	Inactive          int         `json:"inactive"`
	StatusExtended    int         `json:"status_extended"`
	CeoName           string      `json:"ceo_name"`
	CeoType           string      `json:"ceo_type"`
	SnippetString     string      `json:"snippet_string"`
	SnippetType       string      `json:"snippet_type"`
	StatusCode        interface{} `json:"status_code"`
	SvprekrulDate     interface{} `json:"svprekrul_date"`
	MainOkvedId       string      `json:"main_okved_id"`
	OkvedDescr        string      `json:"okved_descr"`
	AuthorizedCapital string      `json:"authorized_capital"`
	RegDate           string      `json:"reg_date"`
	Okpo              interface{} `json:"okpo"`
	Url               string      `json:"url"`
	AciId             string      `json:"aci_id"`
}
