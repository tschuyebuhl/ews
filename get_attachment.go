package ews

import (
	"encoding/xml"
	"errors"
)

type GetAttachmentRequest struct {
	XMLName struct{}                `xml:"m:GetAttachment"`
	Items   []AttachmentItemMessage `xml:"m:AttachmentIds"`
}

type AttachmentItemMessage struct {
	AttachmentId AttachmentId `xml:"t:AttachmentId"`
}

type getAttachmentResponseEnvelop struct {
	XMLName struct{}                  `xml:"Envelope"`
	Body    getAttachmentResponseBody `xml:"Body"`
}
type getAttachmentResponseBody struct {
	GetAttachmentResponse GetAttachmentResponse `xml:"GetAttachmentResponse"`
}

type GetAttachmentResponse struct {
	GetAttachmentResponseMessages GetAttachmentResponseMessages `xml:"ResponseMessages"`
}

type GetAttachmentResponseMessages struct {
	GetAttachmentResponseMessage GetAttachmentResponseMessage `xml:"GetAttachmentResponseMessage"`
}

type GetAttachmentResponseMessage struct {
	Response
	Attachments Attachments `xml:"Attachments"`
}

// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/getattachment-operation
func GetAttachment(c Client, r *GetAttachmentRequest) (*GetAttachmentResponse, error) {

	xmlBytes, err := xml.MarshalIndent(r, "", "  ")
	if err != nil {
		return nil, err
	}

	bb, err := c.SendAndReceive(xmlBytes)
	if err != nil {
		return nil, err
	}

	var soapResp getAttachmentResponseEnvelop
	err = xml.Unmarshal(bb, &soapResp)
	if err != nil {
		return nil, err
	}

	if soapResp.Body.GetAttachmentResponse.GetAttachmentResponseMessages.GetAttachmentResponseMessage.ResponseClass != ResponseClassSuccess {
		return nil, errors.New(soapResp.Body.GetAttachmentResponse.GetAttachmentResponseMessages.GetAttachmentResponseMessage.MessageText)
	}

	return &soapResp.Body.GetAttachmentResponse, nil
}
