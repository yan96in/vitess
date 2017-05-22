package ldapauthserver

import (
	"fmt"
	"testing"

	"gopkg.in/ldap.v2"
)

type MockLdapClient struct{}

func (mlc *MockLdapClient) Connect(network string, config *ServerConfig) error { return nil }
func (mlc *MockLdapClient) Close()                                             {}
func (mlc *MockLdapClient) Bind(username, password string) error {
	if username != "testuser" || password != "testpass" {
		return fmt.Errorf("invalid credentials: %s, %s", username, password)
	}
	return nil
}
func (mlc *MockLdapClient) Search(searchRequest *ldap.SearchRequest) (*ldap.SearchResult, error) {
	return &ldap.SearchResult{}, nil
}

func TestValidateClearText(t *testing.T) {
	asl := &AuthServerLdap{
		Client:        &MockLdapClient{},
		User:          "testuser",
		Password:      "testpass",
		UserDnPattern: "%s",
	}
	_, err := asl.validate("testuser", "testpass")
	if err != nil {
		t.Fatalf("AuthServerLdap failed to validate valid credentials. Got: %v", err)
	}

	_, err = asl.validate("invaliduser", "invalidpass")
	if err == nil {
		t.Fatalf("AuthServerLdap validated invalid credentials.")
	}
}