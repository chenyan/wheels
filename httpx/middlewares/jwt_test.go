// BEGIN: 8f7e6d5c1b3a
package middlewares

import (
	"testing"
)

func TestGenJWTToken(t *testing.T) {
	userID := uint64(123)
	userName := "tom"
	extra := map[string]any{"foo": "bar"}

	token, err := GenJWTToken(userID, userName, extra)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	claims, err := ParseJWT(token, Secret)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if claims.UserID != userID {
		t.Errorf("expected user ID %d, got %d", userID, claims.UserID)
	}

	if claims.Extra["foo"] != extra["foo"] {
		t.Errorf("expected extra data %v, got %v", extra, claims.Extra)
	}
}
