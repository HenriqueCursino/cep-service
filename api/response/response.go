package response

type Response struct {
	Data  any `json:"data,omitempty"`
	Error any `json:"error,omitempty"`
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
