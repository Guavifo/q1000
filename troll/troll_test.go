package troll

import (
	"fmt"
	"strings"
	"testing"
)

func TestTrollMessageWithNoMatch(t *testing.T) {
	// arrange
	b := NewBehavior()
	message := "Hiya... this has nothing for the troll to key onto."

	// act
	result := b.getTrollMessage(message)

	// assert
	if result != "" {
		t.Fatalf("Expected empty troll message. Got %s\n", result)
	}
}

func TestTrollMessageWithAMatch(t *testing.T) {
	// arrange
	b := NewBehavior()
	message := "Hiya... this is kind of like a dead mouse."
	expected := "like a dead mouse."

	// act
	result := b.getTrollMessage(message)

	// assert
	if !strings.HasSuffix(result, expected) {
		t.Fatalf("Expected suffix <%s>. Got %s\n", expected, result)
	}
}

func TestTrollMessageWithALongMessage(t *testing.T) {
	// arrange
	b := NewBehavior()
	message := "Hiya... this is kind of like a dead mouse.whadaya think?"
	expected := "like a dead mouse."

	// act
	result := b.getTrollMessage(message)

	// assert
	if !strings.HasSuffix(result, expected) {
		t.Fatalf("Expected suffix <%s>. Got %s\n", expected, result)
	}
}

func TestTrollMessageWithALineBreak(t *testing.T) {
	// arrange
	b := NewBehavior()
	message := `Hiya... this is kind of like a dead mouse
and a line break
whadaya think?`
	expected := "like a dead mouse"

	// act
	result := b.getTrollMessage(message)

	// assert
	fmt.Println(result)
	if !strings.HasSuffix(result, expected) {
		t.Fatalf("Expected suffix <%s>. Got %s\n", expected, result)
	}
}

func TestTrollMessageWithMoreThanOneMatch(t *testing.T) {
	// arrange
	b := NewBehavior()
	message := `Hiya... this is kind of like a dead mouse
and a line break
whadaya think? I think it is kind of like a nightmare.`
	expected := "like a nightmare."

	// act
	result := b.getTrollMessage(message)

	// assert
	if !strings.HasSuffix(result, expected) {
		t.Fatalf("Expected suffix <%s>. Got %s\n", expected, result)
	}
}

func TestTrollMessageWithConcatenation(t *testing.T) {
	// arrange
	b := NewBehavior()
	message := "Hiya... it's kind of like a dead mouse"
	expected := "like a dead mouse"

	// act
	result := b.getTrollMessage(message)

	// assert
	if !strings.HasSuffix(result, expected) {
		t.Fatalf("Expected suffix <%s>. Got %s\n", expected, result)
	}
}

func TestTrollMessageWithConcatenationError(t *testing.T) {
	// arrange
	b := NewBehavior()
	message := "Hiya... its kind of like a dead mouse"
	expected := "like a dead mouse"

	// act
	result := b.getTrollMessage(message)

	// assert
	if !strings.HasSuffix(result, expected) {
		t.Fatalf("Expected suffix <%s>. Got %s\n", expected, result)
	}
}

func TestTrollMessageWithoutRepeatingIntro(t *testing.T) {
	// arrange
	b := NewBehavior()
	message := "thats kind of a big deal"
	expected := "Your face is kind of a big deal"

	// act
	result := b.getTrollMessage(message)

	// assert
	if result != expected {
		t.Fatalf("Expected: <%s>. Got: <%s>\n", expected, result)
	}
}

func TestTrollMessageWithoutConcatWithoutRepeatingIntro(t *testing.T) {
	// arrange
	b := NewBehavior()
	message := "it is kind of a big deal"
	expected := "Your face is kind of a big deal"

	// act
	result := b.getTrollMessage(message)

	// assert
	if result != expected {
		t.Fatalf("Expected: <%s>. Got: <%s>\n", expected, result)
	}
}

func TestIrishMessage(t *testing.T) {
	// arrange
	b := NewBehavior()
	message := `Hiya... this is kind of like a dead mouse
and a line break
whadaya think? I think it is kind of like a nightmare. Hammock`

	// act
	result := b.getIrishMessage(message)

	// assert
	if result == "" {
		t.Fatalf("Expected message. Got empty string\n")
	}
}

func TestIrishMessageWithNoMatch(t *testing.T) {
	// arrange
	b := NewBehavior()
	message := `Hiya... this is kind of like a dead mouse
and a line break
whadaya think? I think it is kind of like a nightmare. Shamrock`

	// act
	result := b.getIrishMessage(message)

	// assert
	if result != "" {
		t.Fatalf("Expected empty string. Got message\n")
	}
}
