package goWHOIS

type Req struct {
	Object string
}

func NewReq(object string) *Req {
	return &Req{Object: object}
}

func (req *Req) Raw() (string, error) {
	nameServer, _ := req.getAuthoritativeNameServer()
	authoritativeResponse, _ := req.whoisRequest(nameServer, "")
	return authoritativeResponse, nil
}
