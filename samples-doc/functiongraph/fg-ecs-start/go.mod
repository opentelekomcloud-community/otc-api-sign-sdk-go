module sample/fg/startecs

go 1.17

require (
	github.com/opentelekomcloud-community/otc-api-sign-sdk-go/core v0.0.0-00010101000000-000000000000
	github.com/opentelekomcloud-community/otc-functiongraph-go-runtime v0.0.0-20251022103732-4f9a08818827
)

replace (
	github.com/opentelekomcloud-community/otc-api-sign-sdk-go/core => ../../../core
)
