package startrack

import "errors"

type createLabelGroup struct {
	Group      string `json:"group"`
	Layout     string `json:"layout"`
	Branded    bool   `json:"branded"`
	LeftOffset int    `json:"left_offset"`
	TopOffset  int    `json:"top_offset"`
}

type createLabelShipment struct {
	ShipmentID string `json:"shipment_id"`
}

type createLabelPreferences struct {
	Type   string             `json:"type"`
	Groups []createLabelGroup `json:"groups"`
}

type createLabelRequest struct {
	Preferences []createLabelPreferences `json:"preferences"`
	Shipments   []createLabelShipment    `json:"shipments"`
}

type CreateLabelLabel struct {
	RequestID   string   `json:"request_id"`
	Status      string   `json:"status"`
	RequestDate string   `json:"request_date"`
	ShipmentIds []string `json:"shipment_ids"`
}

type CreateLabelResponse struct {
	Message string             `json:"message"`
	Code    string             `json:"code"`
	Labels  []CreateLabelLabel `json:"labels"`
}

func (c *Client) CreateLabel() (*CreateLabelResponse, error) {
	req := createLabelRequest{
		Preferences: []createLabelPreferences{
			{
				Type: "PRINT",
				Groups: []createLabelGroup{
					{
						Group:      "Parcel Post",
						Layout:     "A4-1pp",
						Branded:    true,
						LeftOffset: 0,
						TopOffset:  0,
					}, {
						Group:      "Express Post",
						Layout:     "A4-1pp",
						Branded:    false,
						LeftOffset: 0,
						TopOffset:  0,
					},
				},
			},
		},
		Shipments: []createLabelShipment{
			{
				//ShipmentID: "YBGsEAOgfAgAAAFIRL1786Eu",
				ShipmentID: "jklasdjklsdfjkasdjklasd",
			},
		},
	}

	resp := CreateLabelResponse{}
	if err := c.post(req, "/shipping/v1/labels", &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) GetLabel(id string) error {
	return errors.New("lol wut")
}
