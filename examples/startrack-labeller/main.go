package main

import (
	"fmt"
	"log"
	"os"

	"github.com/porty/startrack"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "key, secret and account number required")
		os.Exit(1)
	}
	c := startrack.New(os.Args[1], os.Args[2], os.Args[3])
	c.BaseURL = startrack.TestbedBetaBaseURL

	resp, err := c.CreateShipment(startrack.CreateShipmentRequest{
		Shipments: []startrack.ShipmentRequest{
			{
				ShipmentReference:  "asdasdasd",
				CustomerReference1: "jsfksjdkassd",
				From: startrack.From{
					Name:     "Bob",
					Lines:    []string{"40 Stewart St"},
					Suburb:   "Richmond",
					State:    "VIC",
					Postcode: "3121",
				},
				To: startrack.To{
					Name:     "Bill",
					Lines:    []string{"41 Stewart St"},
					Suburb:   "Richmond",
					State:    "VIC",
					Postcode: "3121",
				},
				Items: []startrack.Item{
					{
						ItemReference: "jhasdasdasdasdjk",
						ProductID:     "T28S", // ???
						Length:        "10",
						Width:         "10",
						Height:        "10",
						Weight:        "1.0",
						PackagingType: "CTN", // ???
					},
				},
			},
		},
	})
	if err != nil {
		log.Print("Error creating shipment: " + err.Error())
		os.Exit(1)
	}

	log.Printf("%#v", resp)
}
