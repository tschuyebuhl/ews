package ewsutil

import (
	"github.com/mhewedy/ews"
)

func GetAttachment(c ews.Client, attachmentId ews.AttachmentId) (*ews.Attachments, error) {

	req := &ews.GetAttachmentRequest{
		Items: []ews.AttachmentItemMessage{{AttachmentId: attachmentId}},
	}

	resp, err := ews.GetAttachment(c, req)

	if err != nil {
		return nil, err
	}

	return &resp.GetAttachmentResponseMessages.GetAttachmentResponseMessage.Attachments, nil
}
