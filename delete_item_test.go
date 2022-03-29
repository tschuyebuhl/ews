package ews

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func Test_marshal_DeleteItem(t *testing.T) {

	req := &DeleteItemRequest{
		DeleteType: DeleteTypeHardDelete,
		Items: []DeleteItemMessage{{ItemId: ItemId{
			Id:        "TestId",
			ChangeKey: "TestChangeKey",
		}}},
	}

	xmlBytes, err := xml.MarshalIndent(req, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, `<m:DeleteItem DeleteType="HardDelete">
  <m:ItemIds>
    <t:ItemId Id="TestId" ChangeKey="TestChangeKey"></t:ItemId>
  </m:ItemIds>
</m:DeleteItem>`, string(xmlBytes))
}
