package ews

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func Test_marshal_GetAttachment(t *testing.T) {

	req := &GetAttachmentRequest{
		Items: []AttachmentItemMessage{{AttachmentId: AttachmentId{Id: "TestId"}}},
	}

	xmlBytes, err := xml.MarshalIndent(req, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, `<m:GetAttachment>
  <m:AttachmentIds>
    <t:AttachmentId Id="TestId"></t:AttachmentId>
  </m:AttachmentIds>
</m:GetAttachment>`, string(xmlBytes))
}
