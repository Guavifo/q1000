package bibleverse

import (
	"testing"
)

func TestEmptyText(t *testing.T) {
	// arrange
	b := NewBehavior()
	text := ""

	// act
	result := b.parseText(text)

	// assert
	if len(result) > 0 {
		t.Fatal("Expect 0 length result from parseText(). Got: ", len(result))
	}
}
func TestVerseText(t *testing.T) {
	// arrange
	b := NewBehavior()
	text := "Hi th1s a t3st Matthew 12:12 and this Asdf 12:12 is another Matthew 12:12 test Exodus 1:100.test"

	// act
	result := b.parseText(text)

	// assert
	if len(result) != 2 {
		t.Fatal("Expected length of 2 for result from parseText(). Got: ", len(result))
	}
	if result[0] != "Matthew 12:12" {
		t.Fatal("Expected <Matthew 12:12> from parseText(). Got: ", result[0])
	}
	if result[1] != "Exodus 1:100" {
		t.Fatal("Expected <Exodus 1:100> from parseText(). Got: ", result[1])
	}
}
