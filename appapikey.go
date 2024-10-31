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

// AppAPIKeyService contains methods and other services that help with interacting
// with the mpesaflow API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppAPIKeyService] method instead.
type AppAPIKeyService struct {
	Options []option.RequestOption
}

// NewAppAPIKeyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAppAPIKeyService(opts ...option.RequestOption) (r *AppAPIKeyService) {
	r = &AppAPIKeyService{}
	r.Options = opts
	return
}

// Create a new API key for an application
func (r *AppAPIKeyService) New(ctx context.Context, appID string, body AppAPIKeyNewParams, opts ...option.RequestOption) (res *AppAPIKeyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if appID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	path := fmt.Sprintf("apps/%s/api-keys/create", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all API keys for an application
func (r *AppAPIKeyService) List(ctx context.Context, appID string, query AppAPIKeyListParams, opts ...option.RequestOption) (res *AppAPIKeyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if appID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	path := fmt.Sprintf("apps/%s/api-keys/list", appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete an API key
func (r *AppAPIKeyService) Delete(ctx context.Context, appID string, apiKeyID string, opts ...option.RequestOption) (res *AppAPIKeyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if appID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	if apiKeyID == "" {
		err = errors.New("missing required apiKeyId parameter")
		return
	}
	path := fmt.Sprintf("apps/%s/api-keys/%s", appID, apiKeyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AppAPIKeyNewResponse struct {
	APIKeyID string                   `json:"apiKeyId"`
	Message  string                   `json:"message"`
	JSON     appAPIKeyNewResponseJSON `json:"-"`
}

// appAPIKeyNewResponseJSON contains the JSON metadata for the struct
// [AppAPIKeyNewResponse]
type appAPIKeyNewResponseJSON struct {
	APIKeyID    apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppAPIKeyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appAPIKeyNewResponseJSON) RawJSON() string {
	return r.raw
}

type AppAPIKeyListResponse struct {
	Data []AppAPIKeyListResponseData `json:"data"`
	JSON appAPIKeyListResponseJSON   `json:"-"`
}

// appAPIKeyListResponseJSON contains the JSON metadata for the struct
// [AppAPIKeyListResponse]
type appAPIKeyListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppAPIKeyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appAPIKeyListResponseJSON) RawJSON() string {
	return r.raw
}

type AppAPIKeyListResponseData struct {
	ID            string                        `json:"id,required"`
	ApplicationID string                        `json:"applicationId,required"`
	KeyName       string                        `json:"keyName,required"`
	JSON          appAPIKeyListResponseDataJSON `json:"-"`
}

// appAPIKeyListResponseDataJSON contains the JSON metadata for the struct
// [AppAPIKeyListResponseData]
type appAPIKeyListResponseDataJSON struct {
	ID            apijson.Field
	ApplicationID apijson.Field
	KeyName       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AppAPIKeyListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appAPIKeyListResponseDataJSON) RawJSON() string {
	return r.raw
}

type AppAPIKeyDeleteResponse struct {
	Message string                      `json:"message"`
	JSON    appAPIKeyDeleteResponseJSON `json:"-"`
}

// appAPIKeyDeleteResponseJSON contains the JSON metadata for the struct
// [AppAPIKeyDeleteResponse]
type appAPIKeyDeleteResponseJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppAPIKeyDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appAPIKeyDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AppAPIKeyNewParams struct {
	KeyName param.Field[string] `json:"keyName"`
}

func (r AppAPIKeyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AppAPIKeyListParams struct {
	// Cursor for the previous page
	EndingBefore param.Field[string] `query:"ending_before"`
	// Number of items to return
	Limit param.Field[int64] `query:"limit"`
	// Cursor for the next page
	StartingAfter param.Field[string] `query:"starting_after"`
}

// URLQuery serializes [AppAPIKeyListParams]'s query parameters as `url.Values`.
func (r AppAPIKeyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
