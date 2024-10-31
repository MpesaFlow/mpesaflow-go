// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mpesaflow

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/mpesaflow-go/internal/apijson"
	"github.com/stainless-sdks/mpesaflow-go/internal/apiquery"
	"github.com/stainless-sdks/mpesaflow-go/internal/param"
	"github.com/stainless-sdks/mpesaflow-go/internal/requestconfig"
	"github.com/stainless-sdks/mpesaflow-go/option"
)

// TransactionService contains methods and other services that help with
// interacting with the mpesaflow API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionService] method instead.
type TransactionService struct {
	Options []option.RequestOption
}

// NewTransactionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransactionService(opts ...option.RequestOption) (r *TransactionService) {
	r = &TransactionService{}
	r.Options = opts
	return
}

// Create a new transaction
func (r *TransactionService) New(ctx context.Context, body TransactionNewParams, opts ...option.RequestOption) (res *TransactionNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "transactions/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get transaction details
func (r *TransactionService) Get(ctx context.Context, transactionID string, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = append(r.Options[:], opts...)
	if transactionID == "" {
		err = errors.New("missing required transactionId parameter")
		return
	}
	path := fmt.Sprintf("transactions/%s", transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all transactions
func (r *TransactionService) List(ctx context.Context, query TransactionListParams, opts ...option.RequestOption) (res *TransactionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "transactions/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Transaction struct {
	ID               string            `json:"id,required"`
	AccountReference string            `json:"accountReference,required"`
	Amount           float64           `json:"amount,required"`
	PhoneNumber      string            `json:"phoneNumber,required"`
	TransactionDesc  string            `json:"transactionDesc,required"`
	TransactionID    string            `json:"transactionId,required"`
	DateCreated      time.Time         `json:"date_created" format:"date-time"`
	ResultDesc       string            `json:"resultDesc"`
	Status           TransactionStatus `json:"status"`
	JSON             transactionJSON   `json:"-"`
}

// transactionJSON contains the JSON metadata for the struct [Transaction]
type transactionJSON struct {
	ID               apijson.Field
	AccountReference apijson.Field
	Amount           apijson.Field
	PhoneNumber      apijson.Field
	TransactionDesc  apijson.Field
	TransactionID    apijson.Field
	DateCreated      apijson.Field
	ResultDesc       apijson.Field
	Status           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Transaction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionJSON) RawJSON() string {
	return r.raw
}

type TransactionStatus string

const (
	TransactionStatusPending   TransactionStatus = "pending"
	TransactionStatusCompleted TransactionStatus = "completed"
	TransactionStatusFailed    TransactionStatus = "failed"
)

func (r TransactionStatus) IsKnown() bool {
	switch r {
	case TransactionStatusPending, TransactionStatusCompleted, TransactionStatusFailed:
		return true
	}
	return false
}

type TransactionNewResponse struct {
	Message       string                     `json:"message"`
	TransactionID string                     `json:"transactionId"`
	JSON          transactionNewResponseJSON `json:"-"`
}

// transactionNewResponseJSON contains the JSON metadata for the struct
// [TransactionNewResponse]
type transactionNewResponseJSON struct {
	Message       apijson.Field
	TransactionID apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TransactionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionNewResponseJSON) RawJSON() string {
	return r.raw
}

type TransactionListResponse struct {
	Data []Transaction               `json:"data"`
	JSON transactionListResponseJSON `json:"-"`
}

// transactionListResponseJSON contains the JSON metadata for the struct
// [TransactionListResponse]
type transactionListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransactionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transactionListResponseJSON) RawJSON() string {
	return r.raw
}

type TransactionNewParams struct {
	AccountReference param.Field[string]  `json:"accountReference"`
	Amount           param.Field[float64] `json:"amount"`
	PhoneNumber      param.Field[string]  `json:"phoneNumber"`
	TransactionDesc  param.Field[string]  `json:"transactionDesc"`
}

func (r TransactionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TransactionListParams struct {
	AppID param.Field[string] `query:"appId,required"`
	// Cursor for the previous page
	EndingBefore param.Field[string] `query:"ending_before"`
	// Number of items to return
	Limit param.Field[int64] `query:"limit"`
	// Cursor for the next page
	StartingAfter param.Field[string] `query:"starting_after"`
}

// URLQuery serializes [TransactionListParams]'s query parameters as `url.Values`.
func (r TransactionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
