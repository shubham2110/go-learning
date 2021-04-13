package main

import "testing"


func TestHello(t *testing.T){

	assertCorrectMessage := func(t testing.TB, got , want string){
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got , want)
		}
	}
	
	t.Run("saying hello to people", func(t *testing.T){
		got := Hello("Shubham", "")
		want := "Hello, Shubham"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to world", func(t *testing.T){
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
 
    t.Run("in Spanish", func(t *testing.T) {
        got := Hello("Elodie", "Spanish")
        want := "Hola, Elodie"
        assertCorrectMessage(t, got, want)
    })
	
	t.Run("In French", func(t *testing.T){
		got := Hello("Shubham", "French")
		want := "Bonjour, Shubham"
		assertCorrectMessage(t, got, want)
	})
}



/*
func TestHello(t *testing.T){
t.Run("saying hello to people", func(t *testing.T) {
        got := Hello("Chris")
        want := "Hello, Chris"
		
		
        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    })

t.Run("saying hello to world", func(t *testing.T){
	got := Hello("")
	want := "Hello, World"
	
	if got != want {
	
		t.Errorf("got %q want %q", got , want)
	
	}
})
}
*/

/*
func TestHello(t *testing.T) {
    got := Hello("Shubham")
    want := "Hello, Shubham"
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
	got = Hello("Rahul")
    want = "Hello, Rahul"
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
	
}
*/
/*
func TestHello(t *testing.T) {
	got := Hello("Shubham")
	want := "Hello, Shubham"
	
	if got != want {
		t.Errorf("got %q want %q", got, want)
	
	}
	
	got = Hello("")
	want = "Hello, World"
	
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
*/
