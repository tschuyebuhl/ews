package ews

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func Test_marshal_CalendarItem(t *testing.T) {

	attendee := make([]Attendee, 0)
	attendee = append(attendee,
		Attendee{Mailbox: Mailbox{EmailAddress: "User1@example.com"}},
		Attendee{Mailbox: Mailbox{EmailAddress: "User2@example.com"}},
	)
	attendees := make([]Attendees, 0)
	attendees = append(attendees, Attendees{Attendee: attendee})

	start, _ := time.Parse(time.RFC3339, "2006-11-02T14:00:00Z")
	end, _ := time.Parse(time.RFC3339, "2006-11-02T15:00:00Z")

	citem := &CalendarItem{
		Subject: "Planning Meeting",
		Body: Body{
			BodyType: "Text",
			Body:     []byte("Plan the agenda for next week's meeting."),
		},
		ReminderIsSet:              true,
		ReminderMinutesBeforeStart: 60,
		Start:                      start,
		End:                        end,
		IsAllDayEvent:              false,
		LegacyFreeBusyStatus:       "Busy",
		Location:                   "Conference Room 721",
		RequiredAttendees:          attendees,
	}

	xmlBytes, err := xml.MarshalIndent(citem, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, `<CalendarItem>
  <t:Subject>Planning Meeting</t:Subject>
  <t:Body BodyType="Text">Plan the agenda for next week&#39;s meeting.</t:Body>
  <t:ReminderIsSet>true</t:ReminderIsSet>
  <t:ReminderMinutesBeforeStart>60</t:ReminderMinutesBeforeStart>
  <t:Start>2006-11-02T14:00:00Z</t:Start>
  <t:End>2006-11-02T15:00:00Z</t:End>
  <t:IsAllDayEvent>false</t:IsAllDayEvent>
  <t:LegacyFreeBusyStatus>Busy</t:LegacyFreeBusyStatus>
  <t:Location>Conference Room 721</t:Location>
  <t:RequiredAttendees>
    <t:Attendee>
      <t:Mailbox>
        <t:EmailAddress>User1@example.com</t:EmailAddress>
      </t:Mailbox>
    </t:Attendee>
    <t:Attendee>
      <t:Mailbox>
        <t:EmailAddress>User2@example.com</t:EmailAddress>
      </t:Mailbox>
    </t:Attendee>
  </t:RequiredAttendees>
</CalendarItem>`, string(xmlBytes))
}

func Test_marshal_SendMail(t *testing.T) {

	m := Message{
		ItemClass: "IPM.Note",
		Subject:   "test",
		Body: Body{
			Body:     []byte("Test"),
			BodyType: "Text",
		},
		Sender: OneMailbox{
			Mailbox: Mailbox{
				EmailAddress: "User1@example.com",
			},
		},
	}

	mb := make([]Mailbox, 1)

	mb[0].EmailAddress = "User1@example.com"

	m.ToRecipients.Mailbox = append(m.ToRecipients.Mailbox, mb...)

	citem := &CreateItem{
		MessageDisposition: "SendAndSaveCopy",
		SavedItemFolderId:  SavedItemFolderId{DistinguishedFolderId{Id: "sentitems"}},
	}
	citem.Items.Message = append(citem.Items.Message, m)
	xmlBytes, err := xml.MarshalIndent(citem, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, `<m:CreateItem MessageDisposition="SendAndSaveCopy" SendMeetingInvitations="">
  <m:SavedItemFolderId>
    <t:DistinguishedFolderId Id="sentitems"></t:DistinguishedFolderId>
  </m:SavedItemFolderId>
  <m:Items>
    <t:Message>
      <t:ItemClass>IPM.Note</t:ItemClass>
      <t:Subject>test</t:Subject>
      <t:Body BodyType="Text">Test</t:Body>
      <t:Sender>
        <t:Mailbox>
          <t:EmailAddress>User1@example.com</t:EmailAddress>
        </t:Mailbox>
      </t:Sender>
      <t:ToRecipients>
        <t:Mailbox>
          <t:EmailAddress>User1@example.com</t:EmailAddress>
        </t:Mailbox>
      </t:ToRecipients>
    </t:Message>
  </m:Items>
</m:CreateItem>`, string(xmlBytes))
}

func Test_marshal_SendMailWithAttachment(t *testing.T) {

	m := Message{
		ItemClass: "IPM.Note",
		Subject:   "test",
		Body: Body{
			Body:     []byte("Test"),
			BodyType: "Text",
		},
		Attachments: &CreateAttachments{CreateFileAttachment: []CreateFileAttachment{{
			Name:           "Test.txt",
			IsInline:       false,
			IsContactPhoto: false,
			Content:        "VGVzdCB0ZXh0",
		}}},
		Sender: OneMailbox{
			Mailbox: Mailbox{
				EmailAddress: "User1@example.com",
			},
		},
	}

	mb := make([]Mailbox, 1)

	mb[0].EmailAddress = "User1@example.com"

	m.ToRecipients.Mailbox = append(m.ToRecipients.Mailbox, mb...)

	citem := &CreateItem{
		MessageDisposition: "SendAndSaveCopy",
		SavedItemFolderId:  SavedItemFolderId{DistinguishedFolderId{Id: "sentitems"}},
	}
	citem.Items.Message = append(citem.Items.Message, m)
	xmlBytes, err := xml.MarshalIndent(citem, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, `<m:CreateItem MessageDisposition="SendAndSaveCopy" SendMeetingInvitations="">
  <m:SavedItemFolderId>
    <t:DistinguishedFolderId Id="sentitems"></t:DistinguishedFolderId>
  </m:SavedItemFolderId>
  <m:Items>
    <t:Message>
      <t:ItemClass>IPM.Note</t:ItemClass>
      <t:Subject>test</t:Subject>
      <t:Body BodyType="Text">Test</t:Body>
      <t:Sender>
        <t:Mailbox>
          <t:EmailAddress>User1@example.com</t:EmailAddress>
        </t:Mailbox>
      </t:Sender>
      <t:ToRecipients>
        <t:Mailbox>
          <t:EmailAddress>User1@example.com</t:EmailAddress>
        </t:Mailbox>
      </t:ToRecipients>
      <t:Attachments>
        <t:FileAttachment>
          <t:Name>Test.txt</t:Name>
          <t:IsInline>false</t:IsInline>
          <t:IsContactPhoto>false</t:IsContactPhoto>
          <t:Content>VGVzdCB0ZXh0</t:Content>
        </t:FileAttachment>
      </t:Attachments>
    </t:Message>
  </m:Items>
</m:CreateItem>`, string(xmlBytes))
}
