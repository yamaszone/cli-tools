package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseCredentials(t *testing.T) {
	credentialsFile1 := "./fixtures/credentials1"
	testcases := []struct {
		profile string
		field   string
		value   string
	}{
		// "Profile"	"Field"		"Expected Field Value"
		{"", "key", "AKAIISDAIK5LJWV264SAGX7XNA"},
		{"", "secret", "dxt7e3d3fCSf/NftN8Cys5L4duQ2KPQwN0n+oW"},
		{"default", "", ""},
		{"default", "key", "AKAIISDAIK5LJWV264SAGX7XNA"},
		{"default", "secret", "dxt7e3d3fCSf/NftN8Cys5L4duQ2KPQwN0n+oW"},
		{"profile1", "key", "testkeyvalue"},
		{"profile1", "secret", "testsecretvalue"},
		{"profile1", "both", "testkeyvalue\ntestsecretvalue"},
		{"profile1", "", ""},
	}

	for _, testcase := range testcases {
		fieldValue := ParseCredentials(credentialsFile1, testcase.profile, testcase.field)
		if fieldValue != testcase.value {
			t.Errorf("Actual: %s, Expected: %s.", fieldValue, testcase.value)
		}
	}

	credentialsFile2 := "./fixtures/credentials2"
	assert.Panics(t, func() { ParseCredentials(credentialsFile2, "default", "key") })

	credentialsFile3 := "./fixtures/credentials3"
	assert.Panics(t, func() { ParseCredentials(credentialsFile3, "default", "key") })
}

func TestParseCredentialsWithBadFile(t *testing.T) {
	credentialsFile2 := "./fixtures/credentials2"
	assert.Panics(t, func() { ParseCredentials(credentialsFile2, "default", "key") })
	assert.Panics(t, func() { ParseCredentials(credentialsFile2, "bad-profile", "secret") })
	assert.Panics(t, func() { ParseCredentials(credentialsFile2, "bad-profile2", "secret") })
}

func TestParseCredentialsWithNonExistentFile(t *testing.T) {
	credentialsFile3 := "./fixtures/credentials3"
	assert.Panics(t, func() { ParseCredentials(credentialsFile3, "default", "key") })
}
