package ews

import (
	"encoding/xml"
	"errors"
)

type DeleteItemRequest struct {
	XMLName                  struct{}            `xml:"m:DeleteItem"`
	DeleteType               string              `xml:"DeleteType,attr"`
	Items                    []DeleteItemMessage `xml:"m:ItemIds"`
	SendMeetingCancellations *string             `xml:"SendMeetingCancellations,attr,omitempty"`
}

type DeleteItemMessage struct {
	ItemId ItemId `xml:"t:ItemId"`
}

type deleteItemResponseEnvelop struct {
	XMLName struct{}               `xml:"Envelope"`
	Body    deleteItemResponseBody `xml:"Body"`
}
type deleteItemResponseBody struct {
	DeleteItemResponse DeleteItemResponse `xml:"DeleteItemResponse"`
}

type DeleteItemResponse struct {
	DeleteItemResponseMessages DeleteItemResponseMessages `xml:"ResponseMessages"`
}

type DeleteItemResponseMessages struct {
	DeleteItemResponseMessage DeleteItemResponseMessage `xml:"DeleteItemResponseMessage"`
}

type DeleteItemResponseMessage struct {
	Response
}

// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/deleteitem-operation
func DeleteItem(c Client, r *DeleteItemRequest) (*DeleteItemResponse, error) {

	xmlBytes, err := xml.MarshalIndent(r, "", "  ")
	if err != nil {
		return nil, err
	}

	bb, err := c.SendAndReceive(xmlBytes)
	if err != nil {
		return nil, err
	}

	var soapResp deleteItemResponseEnvelop
	err = xml.Unmarshal(bb, &soapResp)
	if err != nil {
		return nil, err
	}

	if soapResp.Body.DeleteItemResponse.DeleteItemResponseMessages.DeleteItemResponseMessage.ResponseClass != ResponseClassSuccess {
		return nil, errors.New(soapResp.Body.DeleteItemResponse.DeleteItemResponseMessages.DeleteItemResponseMessage.MessageText)
	}

	return &soapResp.Body.DeleteItemResponse, nil
}
