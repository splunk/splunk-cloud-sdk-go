package util

import "time"

const (
	// TestToken is the auth token used by stubby server
	TestToken = "eyJraWQiOiJTVGR3WXFCVnFQZlpkeXNNUTQyOElEWTQ5VzRZQzN5MzR2ajNxSl9LRjlvIiwiYWxnIjoiUlMyNTYifQ.eyJ2ZXIiOjEsImp0aSI6IkFULkRaRHJwZkwzZFBQRnMzXzc2b2ZERlB3WUFLTGo1QzgzNmptczQ0Rmh0RlEiLCJpc3MiOiJodHRwczovL3NwbHVuay1jaWFtLm9rdGEuY29tL29hdXRoMi9kZWZhdWx0IiwiYXVkIjoiYXBpOi8vZGVmYXVsdCIsImlhdCI6MTUyNDAwNTEyOCwiZXhwIjoxNTI0MDA4NzI4LCJjaWQiOiIwb2FwYmcyem1MYW1wV2daNDJwNiIsInVpZCI6IjAwdTEwYW14MWxJck9ZbGtnMnA3Iiwic2NwIjpbInByb2ZpbGUiLCJvcGVuaWQiLCJlbWFpbCJdLCJzdWIiOiJkbmd1eWVuQHNwbHVuay5jb20ifQ.DQI7rj_3q9GvRKp-LgJ130p64bhEOnjAOFIUzu7VN7chFH645Pu5IJ--ucxTyOIO2JbY3pd5LoX7AFe7XMGvYWeHkreg3-JucDwMMdKIiwIEcjcTL0i4KpDjtaIkaE6KEgtPDR5vDlrhEM0capK2LQnceCC3RY8Md_BGLNuotHuwPCw0OnXVdoqQkfkyIoCt8ncow0XpUS6hnWEuBSqew-I6nxrw8Z6v8tg-zmx-2r6QeiQJBcLCkOz7U2ViC3hmCARch37uMcc8lRSGGn0eq8dcl3Bfo66U88vzb4moOJe40cCPhjoXPLFuUlzgM5AlvXdIhvkd4i9u2so2CrCNpQ"
	// TestHost is the localhost used for adhoc testing
	TestHost = "https://localhost:8089"
	// TestStubbyHost is the stubby host
	TestStubbyHost = "ssc-sdk-shared-stubby:8882"
	// TestStubbySchme is the stubby scheme
	TestStubbySchme = "http"
	// TestTimeOut is the client timeout used in tests
	TestTimeOut = time.Second * 10
	// TestTenantID is the tenant id used by stubby tests
	TestTenantID = "TEST_TENANT"
)
