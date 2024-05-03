package response

type Response struct {
	Data  any `json:"data,omitempty"`
	Error any `json:"error,omitempty"`
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

func Error(v any) Response {
	return Response{
		Error: v,
	}
}

func Data(v any) Response {
	return Response{
		Data: v,
	}
}
