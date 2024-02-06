package ews

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func Test_marshal_FindItemContains(t *testing.T) {

	req := &FindItemRequest{
		Traversal: "Shallow",
		ItemShape: ItemShape{BaseShape: BaseShapeIdOnly, AdditionalProperties: AdditionalProperties{
			FieldURI: []FieldURI{
				{FieldURI: "item:Subject"},
			},
		}},
		IndexedPageItemView: &IndexedPageItemView{
			MaxEntriesReturned: 5,
			Offset:             0,
			BasePoint:          BasePointBeginning,
		},
		Restriction:     &Restriction{},
		ParentFolderIds: ParentFolderIds{DistinguishedFolderId: DistinguishedFolderId{Id: "inbox"}},
	}

	req.Restriction.IsEqualTo = nil
	req.Restriction.Contains = &Contains{BaseFiltering: BaseFiltering{
		AdditionalProperties: AdditionalProperties{
			FieldURI: []FieldURI{
				{FieldURI: "item:Subject"},
			},
		},
	},
		Constant: []Constant{
			{Value: "Test Subject"},
		},
		ContainmentMode:       "Substring",
		ContainmentComparison: "IgnoreCase",
	}

	xmlBytes, err := xml.MarshalIndent(req, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, `<m:FindItem Traversal="Shallow">
  <m:ItemShape>
    <t:BaseShape>IdOnly</t:BaseShape>
    <t:AdditionalProperties>
      <t:FieldURI FieldURI="item:Subject"></t:FieldURI>
    </t:AdditionalProperties>
  </m:ItemShape>
  <m:IndexedPageItemView MaxEntriesReturned="5" Offset="0" BasePoint="Beginning"></m:IndexedPageItemView>
  <m:ParentFolderIds>
    <t:DistinguishedFolderId Id="inbox"></t:DistinguishedFolderId>
  </m:ParentFolderIds>
  <m:Restriction>
    <t:Contains ContainmentMode="Substring" ContainmentComparison="IgnoreCase">
      <t:FieldURI FieldURI="item:Subject"></t:FieldURI>
      <t:Constant Value="Test Subject"></t:Constant>
    </t:Contains>
  </m:Restriction>
</m:FindItem>`, string(xmlBytes))
}

func Test_marshal_FindItemIsEqualTo(t *testing.T) {

	req := &FindItemRequest{
		Traversal: "Shallow",
		ItemShape: ItemShape{BaseShape: BaseShapeIdOnly, AdditionalProperties: AdditionalProperties{
			FieldURI: []FieldURI{
				{FieldURI: "item:Subject"},
			},
		}},
		IndexedPageItemView: &IndexedPageItemView{
			MaxEntriesReturned: 5,
			Offset:             0,
			BasePoint:          BasePointBeginning,
		},
		Restriction:     &Restriction{},
		ParentFolderIds: ParentFolderIds{DistinguishedFolderId: DistinguishedFolderId{Id: "inbox"}},
	}

	req.Restriction.IsEqualTo = &IsEqualTo{BaseFiltering: BaseFiltering{
		AdditionalProperties: AdditionalProperties{
			FieldURI: []FieldURI{
				{FieldURI: "item:Subject"},
			},
		},
	},
		FieldURIOrConstant: FieldURIOrConstant{
			Constant: []Constant{
				{Value: "Test Subject"},
			},
		}}
	req.Restriction.Contains = nil

	xmlBytes, err := xml.MarshalIndent(req, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, `<m:FindItem Traversal="Shallow">
  <m:ItemShape>
    <t:BaseShape>IdOnly</t:BaseShape>
    <t:AdditionalProperties>
      <t:FieldURI FieldURI="item:Subject"></t:FieldURI>
    </t:AdditionalProperties>
  </m:ItemShape>
  <m:IndexedPageItemView MaxEntriesReturned="5" Offset="0" BasePoint="Beginning"></m:IndexedPageItemView>
  <m:ParentFolderIds>
    <t:DistinguishedFolderId Id="inbox"></t:DistinguishedFolderId>
  </m:ParentFolderIds>
  <m:Restriction>
    <t:IsEqualTo>
      <t:FieldURI FieldURI="item:Subject"></t:FieldURI>
      <t:FieldURIOrConstant>
        <t:Constant Value="Test Subject"></t:Constant>
      </t:FieldURIOrConstant>
    </t:IsEqualTo>
  </m:Restriction>
</m:FindItem>`, string(xmlBytes))
}

func Test_Marshal_CalendarItems(t *testing.T) {
	req := FindItemRequest{
		Traversal: "Shallow",
		ItemShape: ItemShape{
			BaseShape:            BaseShapeIdOnly,
			AdditionalProperties: AdditionalProperties{FieldURI: []FieldURI{{"item:Subject"}, {"calendar:Start"}, {"calendar:End"}}},
		},
		ParentFolderIds: ParentFolderIds{
			DistinguishedFolderId: DistinguishedFolderId{
				Id: "calendar",
			},
		},
	}
	xmlBytes, err := xml.MarshalIndent(req, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	t.Log(string(xmlBytes))
}
