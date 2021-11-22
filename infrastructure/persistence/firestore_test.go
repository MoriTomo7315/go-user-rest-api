package persistence

import (
	"testing"
)

func TestGetUserById(t *testing.T) {
	fc := NewFirestoreClient()
	user, err := fc.GetUserById("tx2ECMJfiFHe5rPhvZrO")
	if err != nil {
		t.Fatal("firestore error occurs")
	}
	expext := "testuser"
	if user.Name != expext {
		t.Error("\nActual： ", user.Name, "\nExpected： ", expext)
	}

	t.Log("Test done")
}
