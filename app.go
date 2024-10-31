// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mpesaflow

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/MpesaFlow/mpesaflow-go/internal/apijson"
	"github.com/MpesaFlow/mpesaflow-go/internal/apiquery"
	"github.com/MpesaFlow/mpesaflow-go/internal/param"
	"github.com/MpesaFlow/mpesaflow-go/internal/requestconfig"
	"github.com/MpesaFlow/mpesaflow-go/option"
)

// AppService contains methods and other services that help with interacting with
// the mpesaflow API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppService] method instead.
type AppService struct {
	Options []option.RequestOption
	APIKeys *AppAPIKeyService
}

// NewAppService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAppService(opts ...option.RequestOption) (r *AppService) {
	r = &AppService{}
	r.Options = opts
	r.APIKeys = NewAppAPIKeyService(opts...)
	return
}

// Create a new application
func (r *AppService) New(ctx context.Context, body AppNewParams, opts ...option.RequestOption) (res *AppNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "apps/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all applications
func (r *AppService) List(ctx context.Context, query AppListParams, opts ...option.RequestOption) (res *AppListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "apps/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete an application
func (r *AppService) Delete(ctx context.Context, appID string, opts ...option.RequestOption) (res *AppDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if appID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	path := fmt.Sprintf("apps/%s", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AppNewResponse struct {
	ApplicationID string             `json:"applicationId"`
	Message       string             `json:"message"`
	JSON          appNewResponseJSON `json:"-"`
}

// appNewResponseJSON contains the JSON metadata for the struct [AppNewResponse]
type appNewResponseJSON struct {
	ApplicationID apijson.Field
	Message       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AppNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appNewResponseJSON) RawJSON() string {
	return r.raw
}

type AppListResponse struct {
	Data []AppListResponseData `json:"data"`
	JSON appListResponseJSON   `json:"-"`
}

// appListResponseJSON contains the JSON metadata for the struct [AppListResponse]
type appListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appListResponseJSON) RawJSON() string {
	return r.raw
}

type AppListResponseData struct {
	ID          string                  `json:"id,required"`
	Name        string                  `json:"name,required"`
	Description string                  `json:"description"`
	JSON        appListResponseDataJSON `json:"-"`
}

// appListResponseDataJSON contains the JSON metadata for the struct
// [AppListResponseData]
type appListResponseDataJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appListResponseDataJSON) RawJSON() string {
	return r.raw
}

type AppDeleteResponse struct {
	Message string                `json:"message"`
	JSON    appDeleteResponseJSON `json:"-"`
}

// appDeleteResponseJSON contains the JSON metadata for the struct
// [AppDeleteResponse]
type appDeleteResponseJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AppNewParams struct {
	Description param.Field[string] `json:"description"`
	Name        param.Field[string] `json:"name"`
}

func (r AppNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AppListParams struct {
	// Cursor for the previous page
	EndingBefore param.Field[string] `query:"ending_before"`
	// Number of items to return
	Limit param.Field[int64] `query:"limit"`
	// Cursor for the next page
	StartingAfter param.Field[string] `query:"starting_after"`
}

// URLQuery serializes [AppListParams]'s query parameters as `url.Values`.
func (r AppListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
