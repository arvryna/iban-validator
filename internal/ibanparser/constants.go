package ibanparser

const IbanMinSize = 5
const IbanMaxSize = 34

type IBANStructure struct {
	Name       string
	CheckDigit int
	Length     int
}

// Key: Country code
var SupportedIbanCountries = map[string]IBANStructure{
	"SE": IBANStructure{Name: "Sweden", Length: 24, CheckDigit: 45},
	"FR": IBANStructure{Name: "France", Length: 27, CheckDigit: 14},
	"NL": IBANStructure{Name: "Netherlands", Length: 18, CheckDigit: 91},
	"IT": IBANStructure{Name: "Italy", Length: 27, CheckDigit: 60},
	"BR": IBANStructure{Name: "Brazil", Length: 29, CheckDigit: 97},
}
