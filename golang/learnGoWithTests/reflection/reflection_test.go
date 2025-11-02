package reflection

import (
	"slices"
	"testing"
)

type Profile struct {
	Age        int
	Profession string
}

type Person struct {
	Name    string
	Profile Profile
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one field",
			struct {
				Name string
			}{"Phil"},
			[]string{"Phil"},
		},
		{
			"Struct with two fields",
			struct {
				Country string
				Capital string
			}{"Germany", "Berlin"},
			[]string{"Germany", "Berlin"},
		},
		{
			"Struct mixed types",
			struct {
				Name string
				Age  int
			}{"Phil", 27},
			[]string{"Phil"},
		},
		{
			"Struct nested fields",
			Person{
				"Phil",
				Profile{27, "Software Engineer"},
			},
			[]string{"Phil", "Software Engineer"},
		},
		{
			"Pointers to things",
			&Person{
				"Phil",
				Profile{27, "Software Engineer"},
			},
			[]string{"Phil", "Software Engineer"},
		},
		{
			"Slices",
			[]Profile{{27, "Phil"}, {34, "Pete"}},
			[]string{"Phil", "Pete"},
		},
		{
			"Arrays",
			[2]Profile{{27, "Phil"}, {34, "Pete"}},
			[]string{"Phil", "Pete"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			got := []string{}

			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !slices.Equal(got, test.ExpectedCalls) {
				t.Errorf("got: %s, want: %s", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with Maps", func(t *testing.T) {
		mapA := map[string]string{
			"foo": "bar",
			"baz": "boz",
		}

		var got []string
		Walk(mapA, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "bar")
		assertContains(t, got, "boz")
	})

	t.Run("with channels", func(t *testing.T) {
		chanA := make(chan string)

		go func() {
			chanA <- "bar"
			chanA <- "boz"
			close(chanA)
		}()

		var got []string
		want := []string{"bar", "boz"}

		Walk(chanA, func(input string) {
			got = append(got, input)
		})

		if !slices.Equal(got, want) {
			t.Errorf("expected %v, but received %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, string) {
			return Profile{27, "Philip"}, "bar"
		}

		var got []string
		want := []string{"Philip", "bar"}

		Walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !slices.Equal(got, want) {
			t.Errorf("expected %v, but received %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false

	for _, x := range haystack {
		if needle == x {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain '%s' but it didn't", haystack, needle)
	}
}
