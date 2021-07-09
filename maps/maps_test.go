package maps

import "testing"

func TestSearch(t *testing.T) {

	assertString := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %v | want %v", got, want)
		}
	}

	assertError := func(t testing.TB, got, want error) {
		t.Helper()
		if got != want {
			t.Errorf("got error %q | want error %q", got, want)
		}
	}

	t.Run("known word", func(t *testing.T) {
		dic := Dictionary{"test": "this is just a test"}
		got, _ := dic.Search("test")
		want := "this is just a test"

		assertString(t, got, want)

	})

	t.Run("unkown word", func(t *testing.T) {
		dic := Dictionary{"test": "this is another test"}
		_, err := dic.Search("unknownkey")

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, err, ErrNotFound)
	})

	t.Run("add new word", func(t *testing.T) {
		dic := Dictionary{}
		dic.Add("name", "name is a definition of a thing")
		want := "name is a definition of a thing"
		got, _ := dic.Search("name")
		assertString(t, got, want)
	})

}
