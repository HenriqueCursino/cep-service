package response

type ResponseError struct {
	Message string `json:"message"`
}

type GetAddressByCepResponse struct {
	Street       string `json:"rua"`
	Neighborhood string `json:"bairro"`
	City         string `json:"cidade"`
	State        string `json:"estado"`
}

func (r GetAddressByCepResponse) Empty() bool {
	return r.Street == "" && r.Neighborhood == "" && r.City == "" && r.State == ""
}
