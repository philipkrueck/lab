package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is what we want"}

	t.Run("known word", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "this is what we want"

		assertError(t, err, nil)
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assertError(t, err, ErrWordNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dict.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dict, word, definition)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dict := Dictionary{word: definition}
		err := dict.Add(word, "new test")

		assertError(t, err, ErrWordAlreadyExists)
		assertDefinition(t, dict, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "definition"
		dict := Dictionary{word: definition}

		newDefinition := "new definition"
		err := dict.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dict, word, newDefinition)
	})
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}

		err := dict.Update("new word", "definition")

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		dict := Dictionary{word: "definition"}

		err := dict.Delete(word)
		assertError(t, err, nil)

		_, err = dict.Search(word)

		assertError(t, err, ErrWordNotFound)
	})
	t.Run("non-existing word", func(t *testing.T) {
		word := "test"
		dict := Dictionary{}

		err := dict.Delete(word)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertDefinition(t testing.TB, dict Dictionary, word, definition string) {
	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("should find added word", err)
	}

	assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}
