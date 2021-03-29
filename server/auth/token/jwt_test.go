package token

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA5Q55KWmVy6klvZ9Kq4Wj1lFYy2U6ESzI+BJBQCoql8/Cbwzt
1wIA+fNS4wKATmcgJoaGBP/xDmJYPv9jCUsQf98TLhK/QhxZpSm0pUiT4U+A/+oL
q48OYkLUVgb3Dpus22GDxPw0JATVR1f78Vig1NDjfq4M3Jyya/+vI+N5sKytoN0O
44ykq89IMHpy0F/7FBUTQqLmMUdGj8iiJoxsCkUOw3/vu39eH4QWxafVV77+EQFS
1LMMtkGPedRWR1btB4Obqv4lfCtJgt0zkBCHZYZ5xItdRBukiNdxPjZQfCwXPoKo
VDfGPChkT2zaPGTORhAdKdAvXXj3T3QpMAhJ7QIDAQABAoIBAQDAewkV48QdAGbW
kaUNvZ+P2YvkX4wMIrY+wBhU8xoXuF6LMzIqG2T2paJhYRiwybgap7C3QtwuWjYN
uk5e4NVcnvhfHyHZTDmIsSM8QWEDVOvyIZrs76oGqGIGaJPzkp4PPEKepdCq8+TF
rwnoJbIBJAgHy/rC6GqGlrV2UxFPhDRm2MJrVBcE3QoVnx9ehD04ovHfZs8OKAdU
//Pa1bkDLRylDlrH4oYAyeRabYc84V/uBqiSUGTrAGdXilqaLst5k8A6D59FaIKf
7y9UVFabLp1oQAW+N+tid8QbMLeqFsc7xKmJQXwlcXj8tC+XNhVrRouBNIc/+153
LjNuYOuBAoGBAP6PXofKhthnQezg3GoBqhoamcd6wEac8Be7IYXRQlu99B2WMzlj
6LXj7qQcC0J2DQ6taBEtS6uqCl6RXkjlDuKyHaH7xkJtI5pJ2dHPLQ1PnXNUlirv
WP8IkuxQ8lqxr7J/bNK5kmjtbqhK0EViWORAgqOoy2379q2d7675xuhXAoGBAOZa
LBOl7wc0sPzEfhqrJGuf7tOUTYIHg9tlNGH5BPN5Nd7L0Rnm2RWMlDvm9mlU79SK
1pXhJtyuVELLq4SpMRtAk6usw75eRynnHhYF5FGNoVybQZblDzVefJFVcWPkAP+o
kZIBrdFzGyxoXLXSSP3yrhgJRiqZ34/RSpBkRgVbAoGBAJ6Mx5EGOLTCo5IZf89a
2ArINK3FK8sPZo8xU7DYZF9l0Y04BxGgw8m+fN2kRfMQOhPBN8czEiWKlEjQfAwk
9v+FduV0PskS1xD7HHsMcgAPbX7u2VVGnITEX1ZnGHwandcsfKhHQwStlbWmW1BL
8BfCglvS/1myjLMYYrf2BMijAoGAf2uJziSPoBDJhki66M2ai94gIfF7Hl5B1e6Z
l0hEMNnvrppXxFgp5P2qUahkzccqEwvpv2aYNKJKAdjfu7ZLb6O9C825SWilPMyX
m98OeP47MBrBoVJ7oy3tAgedfX/1/XWlcHp1R9LASURBZGI59h9zpeAbbB1JySgg
ctbE1NkCgYAmaquvLPkIersSznQeM2A45dCxZVISNU+NR/aes8gzq3o1lvR9nnm9
qfudt0WJevI0l1E8DjPackp9D82XZnanDJzZhm6/oTNEmAKYuojZ45m6KAywwTM/
JXc5LgKRzIfT4KqymiIk6CBISRB4hy4fMSyau4z4K/3932RC7igpnw==
-----END RSA PRIVATE KEY-----`

func TestGenerateToken(t *testing.T) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		t.Fatalf("cannot parse private key: %v", err)
	}
	g := NewJWTTokenGen("coolcar/auth", key)
	g.nowFunc = func() time.Time {
		return time.Unix(1516239022, 0)
	}

	tkn, err := g.GenerateToken("051N090w3Y2Q5W2M7b3w3ffYlv4N090c", 2*time.Hour)
	if err != nil {
		t.Errorf("cannot generate token: %v", err)
	}
	want := `eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiMDUxTjA5MHczWTJRNVcyTTdiM3czZmZZbHY0TjA5MGMifQ.i9AvYdh1gEUQeXvZxAH03Jrlts9Wn95x_0g70jrakrnix4_wXhWJdZBM5DqxEyNRerm__mscF81_HhCnJ_kYxrVvOt07h0ZDfCfU61zeee7mx4KFX-dv6pT8P2tn4x7ggjjpVSdZi6-pripGxSW7VticQvsNKUE-wfcgSLrSTRHfk3ZgK1sIb_WWiL4E9Wmxhvk0OmAQEVfI4zkVogNdv7E4cwjezbjb9Be8DnO7ytBBiPXC7m6VE-nJwFbJWxg1grOHOzMCF2aWFrsmOCFKjxaI6KTWFsMDfUmsegU8A7tq92fFun-aBJKqB-NvOd9updl3AxoUmRdTOWUt58EwVg`
	if tkn != want {
		t.Errorf("wrong token generated. want: %q\n got: %q", want, tkn)
	}

}
