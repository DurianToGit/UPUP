package experiments

import "testing"

func TestValueReceiver(t *testing.T) {
	user := User{
		ID:   1,
		Name: "Tom",
	}

	user.RenameByValue("Value Jerry")
	if user.Name != "Tom" {
		t.Fatalf("expected name=Tom, but got name=%s", user.Name)
	}
}

func TestPointerReceiver(t *testing.T) {
	user := User{
		ID:   1,
		Name: "Tom",
	}
	user.RenameByPointer("Pointer Jerry")
	if user.Name != "Pointer Jerry" {
		t.Fatalf("expected name=Pointer Jerry, but got name=%s", user.Name)
	}
}
