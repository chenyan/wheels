// BEGIN: 8f7e6d5c1b3a
package middlewares

import (
	"testing"
)

func TestGenJWTToken(t *testing.T) {
	userID := int64(123)
	extra := map[string]any{"foo": "bar"}

	token, err := GenJWTToken(userID, extra)
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
