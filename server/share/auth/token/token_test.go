package token

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA5Q55KWmVy6klvZ9Kq4Wj
1lFYy2U6ESzI+BJBQCoql8/Cbwzt1wIA+fNS4wKATmcgJoaGBP/xDmJYPv9jCUsQ
f98TLhK/QhxZpSm0pUiT4U+A/+oLq48OYkLUVgb3Dpus22GDxPw0JATVR1f78Vig
1NDjfq4M3Jyya/+vI+N5sKytoN0O44ykq89IMHpy0F/7FBUTQqLmMUdGj8iiJoxs
CkUOw3/vu39eH4QWxafVV77+EQFS1LMMtkGPedRWR1btB4Obqv4lfCtJgt0zkBCH
ZYZ5xItdRBukiNdxPjZQfCwXPoKoVDfGPChkT2zaPGTORhAdKdAvXXj3T3QpMAhJ
7QIDAQAB
-----END PUBLIC KEY-----`

func TestVerify(t *testing.T) {
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		t.Fatalf("cannot parse public key: %v", err)
	}
	v := &JWTTokenVerifier{
		PublicKey: pubKey,
	}

	cases := []struct {
		name    string
		tkn     string
		now     time.Time
		want    string
		wantErr bool
	}{
		{
			name: "valid_token",
			tkn:  "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiMDUxTjA5MHczWTJRNVcyTTdiM3czZmZZbHY0TjA5MGMifQ.i9AvYdh1gEUQeXvZxAH03Jrlts9Wn95x_0g70jrakrnix4_wXhWJdZBM5DqxEyNRerm__mscF81_HhCnJ_kYxrVvOt07h0ZDfCfU61zeee7mx4KFX-dv6pT8P2tn4x7ggjjpVSdZi6-pripGxSW7VticQvsNKUE-wfcgSLrSTRHfk3ZgK1sIb_WWiL4E9Wmxhvk0OmAQEVfI4zkVogNdv7E4cwjezbjb9Be8DnO7ytBBiPXC7m6VE-nJwFbJWxg1grOHOzMCF2aWFrsmOCFKjxaI6KTWFsMDfUmsegU8A7tq92fFun-aBJKqB-NvOd9updl3AxoUmRdTOWUt58EwVg",
			now:  time.Unix(1516239122, 0),
			want: "051N090w3Y2Q5W2M7b3w3ffYlv4N090c",
		},
		{
			name:    "token_expired",
			tkn:     "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiMDUxTjA5MHczWTJRNVcyTTdiM3czZmZZbHY0TjA5MGMifQ.i9AvYdh1gEUQeXvZxAH03Jrlts9Wn95x_0g70jrakrnix4_wXhWJdZBM5DqxEyNRerm__mscF81_HhCnJ_kYxrVvOt07h0ZDfCfU61zeee7mx4KFX-dv6pT8P2tn4x7ggjjpVSdZi6-pripGxSW7VticQvsNKUE-wfcgSLrSTRHfk3ZgK1sIb_WWiL4E9Wmxhvk0OmAQEVfI4zkVogNdv7E4cwjezbjb9Be8DnO7ytBBiPXC7m6VE-nJwFbJWxg1grOHOzMCF2aWFrsmOCFKjxaI6KTWFsMDfUmsegU8A7tq92fFun-aBJKqB-NvOd9updl3AxoUmRdTOWUt58EwVg",
			now:     time.Unix(1526239122, 0),
			wantErr: true,
		},
		{
			name:    "bad_token",
			tkn:     "`bad_token",
			now:     time.Unix(1516239122, 0),
			wantErr: true,
		},
		{
			name:    "wrong_signature",
			tkn:     "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiMDUxTjA5MHczWTJRNVcyTTdiM3czZmZZbHY0TjA5MWMifQ.i9AvYdh1gEUQeXvZxAH03Jrlts9Wn95x_0g70jrakrnix4_wXhWJdZBM5DqxEyNRerm__mscF81_HhCnJ_kYxrVvOt07h0ZDfCfU61zeee7mx4KFX-dv6pT8P2tn4x7ggjjpVSdZi6-pripGxSW7VticQvsNKUE-wfcgSLrSTRHfk3ZgK1sIb_WWiL4E9Wmxhvk0OmAQEVfI4zkVogNdv7E4cwjezbjb9Be8DnO7ytBBiPXC7m6VE-nJwFbJWxg1grOHOzMCF2aWFrsmOCFKjxaI6KTWFsMDfUmsegU8A7tq92fFun-aBJKqB-NvOd9updl3AxoUmRdTOWUt58EwVg",
			now:     time.Unix(1516239122, 0),
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			jwt.TimeFunc = func() time.Time {
				return c.now
			}
			accountID, err := v.Verify(c.tkn)
			if !c.wantErr && err != nil {
				t.Errorf("verification failed: %v", err)
			}
			if c.wantErr && err == nil {
				t.Errorf("want error; got no  error.")
			}
			if accountID != c.want {
				t.Errorf("wrong account id . want: %q, got: %q", c.want, accountID)
			}
		})
	}
}
