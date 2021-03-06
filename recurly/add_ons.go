package recurly

import (
	"encoding/xml"
	"fmt"
)

type (
	// AddOnsService handles communication with the add ons related methods
	// of the recurly API.
	AddOnsService struct {
		client *Client
	}

	// AddOn represents an individual add on linked to a plan.
	AddOn struct {
		XMLName                     xml.Name   `xml:"add_on"`
		Code                        string     `xml:"add_on_code,omitempty"`
		Name                        string     `xml:"name,omitempty"`
		DefaultQuantity             NullInt    `xml:"default_quantity,omitempty"`
		DisplayQuantityOnHostedPage NullBool   `xml:"display_quantity_on_hosted_page,omitempty"`
		TaxCode                     string     `xml:"tax_code,omitempty"`
		UnitAmountInCents           UnitAmount `xml:"unit_amount_in_cents,omitempty"`
		AccountingCode              string     `xml:"accounting_code,omitempty"`
		CreatedAt                   NullTime   `xml:"created_at,omitempty"`
	}
)

// List returns a list of add ons for a plan.
// https://docs.recurly.com/api/plans/add-ons#list-addons
func (service AddOnsService) List(planCode string, params Params) (*Response, []AddOn, error) {
	action := fmt.Sprintf("plans/%s/add_ons", planCode)
	req, err := service.client.newRequest("GET", action, params, nil)
	if err != nil {
		return nil, nil, err
	}

	var p struct {
		XMLName xml.Name `xml:"add_ons"`
		AddOns  []AddOn  `xml:"add_on"`
	}
	res, err := service.client.do(req, &p)

	return res, p.AddOns, err
}

// Get returns information about an add on.
// https://docs.recurly.com/api/plans/add-ons#lookup-addon
func (service AddOnsService) Get(planCode string, code string) (*Response, AddOn, error) {
	action := fmt.Sprintf("plans/%s/add_ons/%s", planCode, code)
	req, err := service.client.newRequest("GET", action, nil, nil)
	if err != nil {
		return nil, AddOn{}, err
	}

	var a AddOn
	res, err := service.client.do(req, &a)

	return res, a, err
}

// Create adds an add on to a plan.
// https://docs.recurly.com/api/plans/add-ons#create-addon
func (service AddOnsService) Create(planCode string, a AddOn) (*Response, AddOn, error) {
	action := fmt.Sprintf("plans/%s/add_ons", planCode)
	req, err := service.client.newRequest("POST", action, nil, a)
	if err != nil {
		return nil, AddOn{}, err
	}

	var dest AddOn
	res, err := service.client.do(req, &dest)

	return res, dest, err
}

// Update will update the pricing information or description for an add-on.
// Subscriptions who have already subscribed to the add-on will not receive the new pricing.
// https://docs.recurly.com/api/plans/add-ons#update-addon
func (service AddOnsService) Update(planCode string, code string, a AddOn) (*Response, AddOn, error) {
	action := fmt.Sprintf("plans/%s/add_ons/%s", planCode, code)
	req, err := service.client.newRequest("PUT", action, nil, a)
	if err != nil {
		return nil, AddOn{}, err
	}

	var dest AddOn
	res, err := service.client.do(req, &dest)

	return res, dest, err
}

// Delete will remove an add on from a plan.
// https://docs.recurly.com/api/plans/add-ons#delete-addon
func (service AddOnsService) Delete(planCode string, code string) (*Response, error) {
	action := fmt.Sprintf("plans/%s/add_ons/%s", planCode, code)
	req, err := service.client.newRequest("DELETE", action, nil, nil)
	if err != nil {
		return nil, err
	}

	return service.client.do(req, nil)
}
