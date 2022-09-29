package types

type SpeciesType int

const (
	Zooplankton SpeciesType = iota
	Phytoplankton
	Nekton
	Benthos
	Bacteria
)

var SpeciesTypes = [5]string{"zooplankton", "phytoplankton", "nekton", "benthos", "bacteria"}

func (s SpeciesType) String() string {
	return SpeciesTypes[s]
}

func (s SpeciesType) EnumIndex() int {
	return int(s)
}

type SpeciesTypeListResponse struct {
	BaseResponse
	Data []string `json:"data"`
}
