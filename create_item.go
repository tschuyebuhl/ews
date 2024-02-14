package ews

import (
	"encoding/xml"
	"errors"
	"time"
)

type CreateItem struct {
	XMLName                struct{}          `xml:"m:CreateItem"`
	MessageDisposition     string            `xml:"MessageDisposition,attr"`
	SendMeetingInvitations string            `xml:"SendMeetingInvitations,attr"`
	SavedItemFolderId      SavedItemFolderId `xml:"m:SavedItemFolderId"`
	Items                  Items             `xml:"m:Items"`
}

type Items struct {
	Message      []Message      `xml:"t:Message"`
	CalendarItem []CalendarItem `xml:"t:CalendarItem"`
}

type SavedItemFolderId struct {
	DistinguishedFolderId DistinguishedFolderId `xml:"t:DistinguishedFolderId"`
}

type Message struct {
	ItemClass    string             `xml:"t:ItemClass"`
	Subject      string             `xml:"t:Subject"`
	Body         Body               `xml:"t:Body"`
	Sender       OneMailbox         `xml:"t:Sender"`
	ToRecipients XMailbox           `xml:"t:ToRecipients"`
	Attachments  *CreateAttachments `xml:"t:Attachments"`
}

type CreateAttachments struct {
	CreateFileAttachment []CreateFileAttachment `xml:"t:FileAttachment"`
}

type CreateFileAttachment struct {
	Name           string `xml:"t:Name"`
	IsInline       bool   `xml:"t:IsInline"`
	IsContactPhoto bool   `xml:"t:IsContactPhoto"`
	Content        string `xml:"t:Content"`
}

type CalendarItem struct {
	Subject                    string      `xml:"t:Subject"`
	Body                       Body        `xml:"t:Body"`
	ReminderIsSet              bool        `xml:"t:ReminderIsSet"`
	ReminderMinutesBeforeStart int         `xml:"t:ReminderMinutesBeforeStart"`
	Start                      time.Time   `xml:"t:Start"`
	End                        time.Time   `xml:"t:End"`
	IsAllDayEvent              bool        `xml:"t:IsAllDayEvent"`
	LegacyFreeBusyStatus       string      `xml:"t:LegacyFreeBusyStatus"`
	Location                   string      `xml:"t:Location"`
	RequiredAttendees          []Attendees `xml:"t:RequiredAttendees"`
	OptionalAttendees          []Attendees `xml:"t:OptionalAttendees"`
	Resources                  []Attendees `xml:"t:Resources"`
}

type Body struct {
	BodyType string `xml:"BodyType,attr"`
	Body     []byte `xml:",chardata"`
}

type OneMailbox struct {
	Mailbox Mailbox `xml:"t:Mailbox"`
}

type XMailbox struct {
	Mailbox []Mailbox `xml:"t:Mailbox"`
}

type Mailbox struct {
	EmailAddress string `xml:"EmailAddress"`
	Name         string `xml:"Name"`
	RoutingType  string `xml:"RoutingType"`
	MailboxType  string `xml:"MailboxType"`
}

type Attendee struct {
	Mailbox Mailbox `xml:"Mailbox"`
}
type Organizer struct {
	Mailbox Mailbox `xml:"Mailbox"`
}

type Attendees struct {
	Attendee []Attendee `xml:"Attendee"`
}

type createItemResponseBodyEnvelop struct {
	XMLName struct{}               `xml:"Envelope"`
	Body    createItemResponseBody `xml:"Body"`
}
type createItemResponseBody struct {
	CreateItemResponse CreateItemResponse `xml:"CreateItemResponse"`
}

type CreateItemResponse struct {
	ResponseMessages ResponseMessages `xml:"ResponseMessages"`
}

type ResponseMessages struct {
	CreateItemResponseMessage Response `xml:"CreateItemResponseMessage"`
}

// CreateMessageItem
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/createitem-operation-email-message
func CreateMessageItem(c Client, m ...Message) error {

	item := &CreateItem{
		MessageDisposition: "SendAndSaveCopy",
		SavedItemFolderId:  SavedItemFolderId{DistinguishedFolderId{Id: "sentitems"}},
	}
	item.Items.Message = append(item.Items.Message, m...)

	xmlBytes, err := xml.MarshalIndent(item, "", "  ")
	if err != nil {
		return err
	}

	bb, err := c.SendAndReceive(xmlBytes)
	if err != nil {
		return err
	}

	if err := checkCreateItemResponseForErrors(bb); err != nil {
		return err
	}

	return nil
}

// CreateCalendarItem
// https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/createitem-operation-calendar-item
func CreateCalendarItem(c Client, ci ...CalendarItem) error {

	item := &CreateItem{
		SendMeetingInvitations: "SendToAllAndSaveCopy",
		SavedItemFolderId:      SavedItemFolderId{DistinguishedFolderId{Id: "calendar"}},
	}
	item.Items.CalendarItem = append(item.Items.CalendarItem, ci...)

	xmlBytes, err := xml.MarshalIndent(item, "", "  ")
	if err != nil {
		return err
	}

	bb, err := c.SendAndReceive(xmlBytes)
	if err != nil {
		return err
	}

	if err := checkCreateItemResponseForErrors(bb); err != nil {
		return err
	}

	return nil
}

func checkCreateItemResponseForErrors(bb []byte) error {
	var soapResp createItemResponseBodyEnvelop
	if err := xml.Unmarshal(bb, &soapResp); err != nil {
		return err
	}

	resp := soapResp.Body.CreateItemResponse.ResponseMessages.CreateItemResponseMessage
	if resp.ResponseClass == ResponseClassError {
		return errors.New(resp.MessageText)
	}
	return nil
}
