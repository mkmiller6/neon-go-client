package neon_test

import (
	"testing"

	"github.com/mkmiller6/neon-go-client"
	"github.com/mkmiller6/neon-go-client/client"
)

func TestBackendCall(t *testing.T) {
	backend := neon.GetBackendWithConfig()

	nc := client.New("", "", backend)

	t.Run("get an account", func(t *testing.T) {
		acct, _ := nc.Accounts.GetByID(1743)

		got := acct.IndividualAccount.PrimaryContact.FirstName
		want := "Matthew"

		assertEqual(t, got, want)

		got = acct.IndividualAccount.AccountCurrentMembershipStatus
		want = "Inactive"

		assertEqual(t, got, want)
	})

	// t.Run("search for accounts", func(t *testing.T) {
	// 	accts := nc.Accounts.Search(params)

	// 	got := accts[0].IndividualAccount.PrimaryContact.FirstName
	// 	want := "Matthew"

	// 	assertEqual(t, got, want)
	// })

}

func assertEqual(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
