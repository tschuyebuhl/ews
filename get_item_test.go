package ews

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func Test_marshal_GetEmail(t *testing.T) {

	req := &GetItemRequest{
		ItemShape: GetItemShape{
			BaseShape:          BaseShapeDefault,
			IncludeMimeContent: true,
		},
		Items: []GetItemMessage{{ItemId: ItemId{
			Id:        "TestId",
			ChangeKey: "TestChangeKey",
		}}},
	}

	xmlBytes, err := xml.MarshalIndent(req, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, `<m:GetItem>
  <m:ItemShape>
    <t:BaseShape>Default</t:BaseShape>
    <t:IncludeMimeContent>true</t:IncludeMimeContent>
  </m:ItemShape>
  <m:ItemIds>
    <t:ItemId Id="TestId" ChangeKey="TestChangeKey"></t:ItemId>
  </m:ItemIds>
</m:GetItem>`, string(xmlBytes))
}
