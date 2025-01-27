package main

import "testing"

func TestCleanInput(t *testing.T) {

    cases := []struct {
        input string
        expected []string
    }{
        {
        input: "hello world",
        expected: []string{"hello", "world"},
    },
    {
        input: "CHARMANDER Bulbasaur PIkaChu",
        expected: []string{"charmander","bulbasaur","pikachu"},
    },
}


    for _, c := range cases {
        actual := cleanInput(c.input)


        if len(actual) != len(c.expected) {
            t.Errorf("got %d words, want %d words", len(actual), len(c.expected))
            continue
        }


        for i := range actual {
            if actual[i] != c.expected[i] {
                t.Errorf("go %q at position %d, want %q", actual[i],i, c.expected[i])
            }
        }
     }


}
