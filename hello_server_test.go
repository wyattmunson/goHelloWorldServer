package main
import (

    "time"
    "testing"
)


func TestGreetingSpecificJohn(t *testing.T) {
	greeting := CreateGreeting("John")
	if greeting != "Hello, John\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, John\n")
	}
}

func TestGreetingSpecificJohn1(t *testing.T) {
	time.Sleep(1 * time.Second)
	greeting := CreateGreeting("John1")
	if greeting != "Hello, John1\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, John\n")
	}
}


func TestGreetingSpecificJohn2(t *testing.T) {
	time.Sleep(2 * time.Second)
	greeting := CreateGreeting("John2")
	if greeting != "Hello, John2\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, John\n")
	}
}


func TestGreetingSpecificJohn3(t *testing.T) {
	time.Sleep(3 * time.Second)
	greeting := CreateGreeting("John3")
	if greeting != "Hello, John\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, John\n")
	}
}

func TestGreetingSpecificJohn4(t *testing.T) {
	time.Sleep(4 * time.Second)
	greeting := CreateGreeting("John4")
	if greeting != "Hello, John4\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, John\n")
	}
}

func TestGreetingSpecificJohn5(t *testing.T) {
	time.Sleep(5 * time.Second)
	greeting := CreateGreeting("John5")
	if greeting != "Hello, John5\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, John\n")
	}
}

func TestGreetingSpecificJohn6(t *testing.T) {
	time.Sleep(6 * time.Second)
	greeting := CreateGreeting("John6")
	if greeting != "Hello, John6\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, John\n")
	}
}


func TestGreetingSpecificJohn7(t *testing.T) {
	time.Sleep(7 * time.Second)
	greeting := CreateGreeting("John7")
	if greeting != "Hello, John7\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, John\n")
	}
}

func TestGreetingSpecificDemo(t *testing.T) {
	greeting := CreateGreeting("Demo")
	if greeting != "Hello, Demo\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, Demo\n")
	}
}

/* func TestShowFailure(t *testing.T) {
	greeting := CreateGreeting("Demo1")
	if greeting != "Hello, Demo\n" {
		t.Errorf("Intentional failure. got: %s, want: %s.", greeting, "Hello, Demo\n")
	}
} */



func TestGreetingDefault(t *testing.T) {
	greeting := CreateGreeting("")
	if greeting != "Hello, Guest\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, Guest\n")
	}
}
 


