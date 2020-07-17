// Code generated by goa v3.2.0, DO NOT EDIT.
//
// discount HTTP server encoders and decoders
//
// Command:
// $ goa gen fruitshop/design

package server

import (
	"context"
	discountviews "fruitshop/gen/discount/views"
	"net/http"

	goahttp "goa.design/goa/v3/http"
)

// EncodeGetResponse returns an encoder for responses returned by the discount
// get endpoint.
func EncodeGetResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(discountviews.DiscountManagementCollection)
		enc := encoder(ctx, w)
		body := NewDiscountManagementResponseCollection(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetRequest returns a decoder for requests sent to the discount get
// endpoint.
func DecodeGetRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			userID string

			params = mux.Vars(r)
		)
		userID = params["userId"]
		payload := NewGetPayload(userID)

		return payload, nil
	}
}

// marshalDiscountviewsDiscountManagementViewToDiscountManagementResponse
// builds a value of type *DiscountManagementResponse from a value of type
// *discountviews.DiscountManagementView.
func marshalDiscountviewsDiscountManagementViewToDiscountManagementResponse(v *discountviews.DiscountManagementView) *DiscountManagementResponse {
	res := &DiscountManagementResponse{
		UserID: *v.UserID,
		Name:   v.Name,
		Status: v.Status,
	}

	return res
}
