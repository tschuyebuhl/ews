package ews

import (
	"encoding/xml"
	"errors"
)

type GetItemMessage struct {
	ItemId ItemId `xml:"t:ItemId"`
}

type GetItemRequest struct {
	XMLName   struct{}         `xml:"m:GetItem"`
	ItemShape GetItemShape     `xml:"m:ItemShape"`
	Items     []GetItemMessage `xml:"m:ItemIds"`
}

type GetItemShape struct {
	BaseShape          BaseShape   `xml:"t:BaseShape"`
	IncludeMimeContent BooleanType `xml:"t:IncludeMimeContent"`
}

type getItemResponseEnvelop struct {
	XMLName struct{}            `xml:"Envelope"`
	Body    getItemResponseBody `xml:"Body"`
}
type getItemResponseBody struct {
	GetItemResponse GetItemResponse `xml:"GetItemResponse"`
}

type GetItemResponse struct {
	GetItemResponseMessages GetItemResponseMessages `xml:"ResponseMessages"`
}

type GetItemResponseMessages struct {
	GetItemResponseMessage GetItemResponseMessage `xml:"GetItemResponseMessage"`
}

type GetItemResponseMessage struct {
	Response
	Items GetItems `xml:"Items"`
}

type GetItems struct {
	Message []ItemMessage `xml:"Message"`
}

type ItemMessage struct {
	ItemId                     ItemId      `xml:"ItemId"`
	Subject                    string      `xml:"Subject"`
	MimeContent                MimeContent `xml:"MimeContent"`
	Sensitivity                string      `xml:"Sensitivity"`
	Body                       EmailBody   `xml:"Body"`
	Attachments                Attachments `xml:"Attachments"`
	Size                       uint64      `xml:"Size"`
	DateTimeSent               Time        `xml:"DateTimeSent"`
	DateTimeCreated            Time        `xml:"DateTimeCreated"`
	HasAttachments             BooleanType `xml:"HasAttachments"`
	IsAssociated               BooleanType `xml:"IsAssociated"`
	ToRecipients               From        `xml:"ToRecipients"`
	IsReadReceiptRequested     BooleanType `xml:"IsReadReceiptRequested"`
	IsDeliveryReceiptRequested BooleanType `xml:"IsDeliveryReceiptRequested"`
	From                       From        `xml:"From"`
	IsRead                     BooleanType `xml:"IsRead"`
}

type GetItemEmailAddress struct {
	Name         string `xml:"Name"`
	EmailAddress string `xml:"EmailAddress"`
	RoutingType  string `xml:"RoutingType"`
	MailboxType  string `xml:"MailboxType"`
}

type From struct {
	Mailbox []GetItemEmailAddress `xml:"Mailbox"`
}

type Attachments struct {
	FileAttachment []FileAttachment `xml:"FileAttachment"`
}

type FileAttachment struct {
	AttachmentId     AttachmentId `xml:"AttachmentId"`
	Name             string       `xml:"Name"`
	ContentType      string       `xml:"ContentType"`
	Size             uint64       `xml:"Size"`
	LastModifiedTime Time         `xml:"LastModifiedTime"`
	IsInline         bool         `xml:"IsInline"`
	IsContactPhoto   bool         `xml:"IsContactPhoto"`
}

type AttachmentId struct {
	Id string `xml:"Id,attr"`
}

// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/getitem-operation-email-message
func GetItem(c Client, r *GetItemRequest) (*GetItemResponse, error) {

	xmlBytes, err := xml.MarshalIndent(r, "", "  ")
	if err != nil {
		return nil, err
	}

	bb, err := c.SendAndReceive(xmlBytes)
	if err != nil {
		return nil, err
	}

	var soapResp getItemResponseEnvelop
	err = xml.Unmarshal(bb, &soapResp)
	if err != nil {
		return nil, err
	}

	if soapResp.Body.GetItemResponse.GetItemResponseMessages.GetItemResponseMessage.ResponseClass != ResponseClassSuccess {
		return nil, errors.New(soapResp.Body.GetItemResponse.GetItemResponseMessages.GetItemResponseMessage.MessageText)
	}

	return &soapResp.Body.GetItemResponse, nil
}
