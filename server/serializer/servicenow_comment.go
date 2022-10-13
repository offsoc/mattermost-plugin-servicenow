package serializer

import (
	"encoding/json"
	"errors"
	"io"
	"strings"
)

type ServiceNowComment struct {
	CommentsAndWorkNotes string `json:"comments_and_work_notes"`
}

type ServiceNowCommentPayload struct {
	Comments string `json:"comments"`
}

type ServiceNowCommentsResult struct {
	Result *ServiceNowComment `json:"result"`
}

func ServiceNowCommentPayloadFromJSON(data io.Reader) (*ServiceNowCommentPayload, error) {
	var scp *ServiceNowCommentPayload
	if err := json.NewDecoder(data).Decode(&scp); err != nil {
		return nil, err
	}

	return scp, nil
}

func (s *ServiceNowCommentPayload) Validate() error {
	s.Comments = strings.TrimSpace(s.Comments)
	if s.Comments == "" {
		return errors.New("comment value cannot be empty")
	}

	return nil
}
