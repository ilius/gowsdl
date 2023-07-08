// Code generated by gowsdl DO NOT EDIT.

package gen

import (
	"context"
	"encoding/xml"
	"github.com/ilius/gowsdl/soap"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type AnyType struct {
	InnerXML string `xml:",innerxml"`
}

type AnyURI string

type NCName string

type GetInfo struct {
	XMLName xml.Name `xml:"http://www.mnb.hu/webservices/ GetInfo"`

	// comment

	Id string `xml:"Id,omitempty" json:"Id,omitempty"`
}

type GetInfoResponse struct {
	XMLName xml.Name `xml:"http://www.mnb.hu/webservices/ GetInfoResponse"`

	// this is a comment
	GetInfoResult string `xml:"GetInfoResult,omitempty" json:"GetInfoResult,omitempty"`
}

type ElementWithLocalSimpleType string

const (

	// First enum value
	ElementWithLocalSimpleTypeEnum1 ElementWithLocalSimpleType = "enum1"

	// Second enum value
	ElementWithLocalSimpleTypeEnum2 ElementWithLocalSimpleType = "enum2"
)

type StartDate soap.XSDDateTime

func (xdt StartDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return soap.XSDDateTime(xdt).MarshalXML(e, start)
}

func (xdt *StartDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	return (*soap.XSDDateTime)(xdt).UnmarshalXML(d, start)
}

type ResponseStatus struct {
	Status []struct {
		Value string `xml:",chardata" json:"-,"`

		Code string `xml:"http://www.mnb.hu/webservices/ code,attr,omitempty" json:"code,omitempty"`
	} `xml:"status,omitempty" json:"status,omitempty"`

	ResponseCode string `xml:"http://www.mnb.hu/webservices/ responseCode,attr,omitempty" json:"responseCode,omitempty"`
}

type MNBArfolyamServiceType interface {
	GetInfoSoap(request *GetInfo) (*GetInfoResponse, error)

	GetInfoSoapContext(ctx context.Context, request *GetInfo) (*GetInfoResponse, error)
}

type mNBArfolyamServiceType struct {
	client *soap.Client
}

func NewMNBArfolyamServiceType(client *soap.Client) MNBArfolyamServiceType {
	return &mNBArfolyamServiceType{
		client: client,
	}
}

func (service *mNBArfolyamServiceType) GetInfoSoapContext(ctx context.Context, request *GetInfo) (*GetInfoResponse, error) {
	response := new(GetInfoResponse)
	err := service.client.CallContext(ctx, "''", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *mNBArfolyamServiceType) GetInfoSoap(request *GetInfo) (*GetInfoResponse, error) {
	return service.GetInfoSoapContext(
		context.Background(),
		request,
	)
}
