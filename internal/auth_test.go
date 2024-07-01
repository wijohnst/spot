package internal

import (
	"testing"
	"wijohnst/spot/internal/pkg/test_utils"
)

func TestAuthGetToken(t *testing.T) {
	sut := "Auth - GetToken()"

	//Test 0
	desc := "Should return an Auth token"
	expected := "Foo"
	auth := Auth{}
	actual := ""

	auth.Init()

	test_utils.Assert(sut, desc, 0, t, expected, actual, nil)
}
