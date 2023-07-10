package gowsdl

var serverTmpl = `

var WSDLUndefinedError = errors.New("Server was unable to process request. --> Object reference not set to an instance of an object.")

type SOAPEnvelopeRequest struct {
	XMLName xml.Name ` + "`" + `xml:"http://www.w3.org/2003/05/soap-envelope Envelope"` + "`" + `
	Body SOAPBodyRequest
}

type SOAPBodyRequest struct {
	XMLName xml.Name ` + "`" + `xml:"http://www.w3.org/2003/05/soap-envelope Body"` + "`" + `
	{{range .}}
		{{range .Operations}}
				{{$requestType := findType .Input.Message | replaceReservedWords | makePublic}} ` + `
  				{{$requestType}} *{{$requestType}} ` + "`" + `xml:,omitempty` + "`" + `
		{{end}}
	{{end}}
}

type SOAPEnvelopeResponse struct { ` + `
	XMLName    xml.Name` + "`" + `xml:"soap:Envelope"` + "`" + `
	PrefixSoap string  ` + "`" + `xml:"xmlns:soap,attr"` + "`" + `
	PrefixXsi  string  ` + "`" + `xml:"xmlns:xsi,attr"` + "`" + `
	PrefixXsd  string  ` + "`" + `xml:"xmlns:xsd,attr"` + "`" + `

	Body SOAPBodyResponse
}

func NewSOAPEnvelopResponse() *SOAPEnvelopeResponse {
	return &SOAPEnvelopeResponse{
		PrefixSoap: "http://www.w3.org/2003/05/soap-envelope",
		PrefixXsd:  "http://www.w3.org/2001/XMLSchema",
		PrefixXsi:  "http://www.w3.org/2001/XMLSchema-instance",
	}
}

type Fault struct { ` + `
	XMLName xml.Name ` + "`" + `xml:"SOAP-ENV:Fault"` + "`" + `
	Space   string   ` + "`" + `xml:"xmlns:SOAP-ENV,omitempty,attr"` + "`" + `

	Code   string    ` + "`" + `xml:"faultcode,omitempty"` + "`" + `
	String string    ` + "`" + `xml:"faultstring,omitempty"` + "`" + `
	Actor  string 	 ` + "`" + `xml:"faultactor,omitempty"` + "`" + `
	Detail string    ` + "`" + `xml:"detail,omitempty"` + "`" + `
}


type SOAPBodyResponse struct { ` + `
	XMLName xml.Name   ` + "`" + `xml:"soap:Body"` + "`" + `
	Fault   *Fault ` + "`" + `xml:",omitempty"` + "`" + `
{{range .}}
	{{range .Operations}}
		{{$responseType := findType .Output.Message | replaceReservedWords | makePublic}}
		{{$requestType := findType .Input.Message | replaceReservedWords | makePublic}} ` + `
			{{$requestType}} *{{$responseType}} ` + "`" + `xml:",omitempty"` + "`" + `
	{{end}}
{{end}}

}

{{range .}}
	{{range .Operations}}
		{{$responseType := findType .Output.Message | replaceReservedWords | makePublic}}
		{{$requestType := findType .Input.Message | replaceReservedWords | makePublic}}
		{{$requestTypeSource := findType .Input.Message | replaceReservedWords }}
func (service *SOAPEnvelopeRequest) {{$requestType}}Func(request *{{$requestType}}) (*{{$responseType}}, error) {
    return &{{$responseType}}{}, nil
}
	{{end}}
{{end}}


func (service *SOAPEnvelopeRequest) call(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/xml; charset=utf-8")

	if r.Method == http.MethodGet {
		w.Write([]byte(wsdl))
		return
	}

	resp := NewSOAPEnvelopResponse()
	defer func() {
		if r := recover(); r != nil {
			resp.Body.Fault = &Fault{}
			resp.Body.Fault.Space = "http://www.w3.org/2003/05/soap-envelope"
			resp.Body.Fault.Code = "soap:Server"
			resp.Body.Fault.Detail = fmt.Sprintf("%v", r)
			resp.Body.Fault.String = fmt.Sprintf("%v", r)
		}
		xml.NewEncoder(w).Encode(resp)
	}()

	err := xml.NewDecoder(r.Body).Decode(service)
	if err != nil {
		panic(err)
	}

	switch {
	{{range .}}
	{{range .Operations}}
		{{$requestType := findType .Input.Message | replaceReservedWords | makePublic}} ` + `
	case service.Body.{{$requestType}} != nil:
		resRes, err := service.{{$requestType}}Func(service.Body.{{$requestType}})
		if err != nil {
			panic(err)
		}
		resp.Body.{{$requestType}} = resRes
	{{end}}
	{{end}}
	default:
		panic(WSDLUndefinedError)
	}

}

func Endpoint(w http.ResponseWriter, r *http.Request) {
	request := SOAPEnvelopeRequest{}
	request.call(w, r)
}

`
