package troll

import (
	"fmt"
	"strings"
	"testing"
)

func TestTrollMessageWithNoMatch(t *testing.T) {
	// arrange
	message := "Hiya... this has nothing for the troll to key onto."

	// act
	result := getTrollMessage(message)

	// assert
	if result != "" {
		t.Fatalf("Expected empty troll message. Got %s\n", result)
	}
}

func TestTrollMessageWithAMatch(t *testing.T) {
	// arrange
	message := "Hiya... this is kind of like a dead mouse."
	expected := "like a dead mouse."

	// act
	result := getTrollMessage(message)

	// assert
	if !strings.HasSuffix(result, expected) {
		t.Fatalf("Expected suffix <%s>. Got %s\n", expected, result)
	}
}

func TestTrollMessageWithALongMessage(t *testing.T) {
	// arrange
	message := "Hiya... this is kind of like a dead mouse.whadaya think?"
	expected := "like a dead mouse."

	// act
	result := getTrollMessage(message)

	// assert
	if !strings.HasSuffix(result, expected) {
		t.Fatalf("Expected suffix <%s>. Got %s\n", expected, result)
	}
}

func TestTrollMessageWithALineBreak(t *testing.T) {
	// arrange
	message := `Hiya... this is kind of like a dead mouse
and a line break
whadaya think?`
	expected := "like a dead mouse"

	// act
	result := getTrollMessage(message)

	// assert
	fmt.Println(result)
	if !strings.HasSuffix(result, expected) {
		t.Fatalf("Expected suffix <%s>. Got %s\n", expected, result)
	}
}

func TestTrollMessageWithMoreThanOneMatch(t *testing.T) {
	// arrange
	message := `Hiya... this is kind of like a dead mouse
and a line break
whadaya think? I think it is kind of like a nightmare.`
	expected := "like a nightmare."

	// act
	result := getTrollMessage(message)

	// assert
	if !strings.HasSuffix(result, expected) {
		t.Fatalf("Expected suffix <%s>. Got %s\n", expected, result)
	}
}
