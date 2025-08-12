package requests

type RequestOutput struct {
	Payload map[string][]error
}

func (r *RequestOutput) HasErrors() bool {
	for _, errors := range r.Payload {
		if len(errors) > 0 {
			return true
		}
	}
	return false
}
