package startrack

import (
	"time"
)

type From struct {
	Name     string   `json:"name"`
	Lines    []string `json:"lines"`
	Suburb   string   `json:"suburb"`
	State    string   `json:"state"`
	Postcode string   `json:"postcode"`
	Phone    string   `json:"phone"`
	Email    string   `json:"email"`
}

type To struct {
	Name         string   `json:"name"`
	BusinessName string   `json:"business_name"`
	Lines        []string `json:"lines"`
	Suburb       string   `json:"suburb"`
	State        string   `json:"state"`
	Postcode     string   `json:"postcode"`
	Phone        string   `json:"phone"`
	Email        string   `json:"email"`
}

type Item struct {
	ItemReference        string `json:"item_reference"`
	ProductID            string `json:"product_id"`
	Length               string `json:"length"`
	Height               string `json:"height"`
	Width                string `json:"width"`
	Weight               string `json:"weight"`
	AuthorityToLeave     bool   `json:"authority_to_leave"`
	AllowPartialDelivery bool   `json:"allow_partial_delivery"`
	PackagingType        string `json:"packaging_type"`
	// Features             struct {
	// 	TRANSITCOVER struct {
	// 		Attributes struct {
	// 			CoverAmount int `json:"cover_amount"`
	// 		} `json:"attributes"`
	// 	} `json:"TRANSIT_COVER"`
	// } `json:"features,omitempty"`
}

type ShipmentRequest struct {
	ShipmentReference    string `json:"shipment_reference"`
	CustomerReference1   string `json:"customer_reference_1"`
	CustomerReference2   string `json:"customer_reference_2"`
	EmailTrackingEnabled bool   `json:"email_tracking_enabled"`
	From                 From   `json:"from"`
	To                   To     `json:"to"`
	Items                []Item `json:"items"`
}

type CreateShipmentRequest struct {
	Shipments []ShipmentRequest `json:"shipments"`
}

type CreateShipmentResponse struct {
	Shipments []ShipmentResponse `json:"shipments"`
}

type ShipmentResponse struct {
	ShipmentID           string    `json:"shipment_id"`
	ShipmentReference    string    `json:"shipment_reference"`
	ShipmentCreationDate time.Time `json:"shipment_creation_date"`
	CustomerReference1   string    `json:"customer_reference_1"`
	CustomerReference2   string    `json:"customer_reference_2"`
	SenderReferences     []string  `json:"sender_references"`
	EmailTrackingEnabled bool      `json:"email_tracking_enabled"`
	Items                []struct {
		ItemID          string `json:"item_id"`
		ItemReference   string `json:"item_reference"`
		TrackingDetails struct {
			ArticleID     string `json:"article_id"`
			ConsignmentID string `json:"consignment_id"`
		} `json:"tracking_details"`
		ProductID   string `json:"product_id"`
		ItemSummary struct {
			TotalCost int     `json:"total_cost"`
			TotalGst  float64 `json:"total_gst"`
			Status    string  `json:"status"`
		} `json:"item_summary"`
	} `json:"items"`
	ShipmentSummary struct {
		TotalCost       int     `json:"total_cost"`
		TotalGst        float64 `json:"total_gst"`
		Status          string  `json:"status"`
		NumberOfItems   int     `json:"number_of_items"`
		TrackingSummary struct {
			Initiated int `json:"Initiated"`
		} `json:"tracking_summary"`
	} `json:"shipment_summary"`
}

func (c *Client) CreateShipment(req CreateShipmentRequest) (*CreateShipmentResponse, error) {
	resp := CreateShipmentResponse{}
	if err := c.post(req, "/shipping/v1/shipments", &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
