package dictionary

import "testing"

func TestSearch(t *testing.T) {
	disctionary := Dictionary{"chromatography": "The process of making food from sunlight"}

	t.Run("known word", func(t *testing.T) {
		got, _ := disctionary.Search("chromatography")
		want := "The process of making food from sunlight"
		assertString(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := disctionary.Search("love")
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "passion"
		def := "a strong feeling of enthusiasm"
		err := dictionary.Add(word, def)
		if err != nil {
			t.Fatal("unexpected error")
		}
		assertDefinition(t, dictionary, word, def)
	})

	t.Run("adding existing word", func(t *testing.T) {
		word := "passion"
		def := "a strong feeling of enthusiasm"
		dictionary := Dictionary{word: def}
		err := dictionary.Add(word, def)
		assertError(t, err, ErrWordExists)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("updating existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)
		assertError(t, err, nil)

		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("updating non-existent word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete existing key", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Delete(word)
		assertError(t, err, nil)
		_, err = dictionary.Search(word)
		assertError(t, err, ErrNotFound)
	})
	t.Run("delete non-existent key", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		err := dictionary.Delete(word)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, err error, want error) {
	t.Helper()
	if err != want {
		t.Errorf("got %q, want %q", err, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, def string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertString(t, got, def)
}
