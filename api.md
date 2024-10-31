# Apps

Response Types:

- <a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go">mpesaflow</a>.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#AppNewResponse">AppNewResponse</a>
- <a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go">mpesaflow</a>.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#AppListResponse">AppListResponse</a>
- <a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go">mpesaflow</a>.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#AppDeleteResponse">AppDeleteResponse</a>

Methods:

- <code title="post /apps/create">client.Apps.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#AppService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go">mpesaflow</a>.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#AppNewParams">AppNewParams</a>) (<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go">mpesaflow</a>.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#AppNewResponse">AppNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /apps/list">client.Apps.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#AppService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go">mpesaflow</a>.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#AppListParams">AppListParams</a>) (<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go">mpesaflow</a>.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#AppListResponse">AppListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /apps/{appId}">client.Apps.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#AppService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, appID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go">mpesaflow</a>.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#AppDeleteResponse">AppDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## APIKeys

Response Types:

- <a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go">mpesaflow</a>.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#AppAPIKeyNewResponse">AppAPIKeyNewResponse</a>
- <a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go">mpesaflow</a>.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#AppAPIKeyListResponse">AppAPIKeyListResponse</a>
- <a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go">mpesaflow</a>.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#AppAPIKeyDeleteResponse">AppAPIKeyDeleteResponse</a>

Methods:

- <code title="post /apps/{appId}/api-keys/create">client.Apps.APIKeys.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#AppAPIKeyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, appID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go">mpesaflow</a>.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#AppAPIKeyNewParams">AppAPIKeyNewParams</a>) (<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go">mpesaflow</a>.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#AppAPIKeyNewResponse">AppAPIKeyNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /apps/{appId}/api-keys/list">client.Apps.APIKeys.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#AppAPIKeyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, appID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go">mpesaflow</a>.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#AppAPIKeyListParams">AppAPIKeyListParams</a>) (<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go">mpesaflow</a>.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#AppAPIKeyListResponse">AppAPIKeyListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /apps/{appId}/api-keys/{apiKeyId}">client.Apps.APIKeys.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#AppAPIKeyService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, appID <a href="https://pkg.go.dev/builtin#string">string</a>, apiKeyID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go">mpesaflow</a>.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#AppAPIKeyDeleteResponse">AppAPIKeyDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Transactions

Response Types:

- <a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go">mpesaflow</a>.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#Transaction">Transaction</a>
- <a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go">mpesaflow</a>.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#TransactionNewResponse">TransactionNewResponse</a>
- <a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go">mpesaflow</a>.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#TransactionListResponse">TransactionListResponse</a>

Methods:

- <code title="post /transactions/create">client.Transactions.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#TransactionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go">mpesaflow</a>.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#TransactionNewParams">TransactionNewParams</a>) (<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go">mpesaflow</a>.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#TransactionNewResponse">TransactionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /transactions/{transactionId}">client.Transactions.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#TransactionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go">mpesaflow</a>.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#Transaction">Transaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /transactions/list">client.Transactions.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#TransactionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go">mpesaflow</a>.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#TransactionListParams">TransactionListParams</a>) (<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go">mpesaflow</a>.<a href="https://pkg.go.dev/github.com/MpesaFlow/mpesaflow-go#TransactionListResponse">TransactionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
