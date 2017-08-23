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
	text := "Hi th1s a t3st Matthew 12:12 and this Asdf 12:12 is another Matthew 12:12 test Exodus 1:100. test 1 Kings 1:1 and Exodus 1:1-10"

	// act
	result := b.parseText(text)

	// assert
	if len(result) != 4 {
		t.Fatal("Expected length of 4 for result from parseText(). Got: ", len(result))
	}
	if result[0] != "Matthew 12:12" {
		t.Fatal("Expected <Matthew 12:12> from parseText(). Got: ", result[0])
	}
	if result[1] != "Exodus 1:100" {
		t.Fatal("Expected <Exodus 1:100> from parseText(). Got: ", result[1])
	}
	if result[2] != "1 Kings 1:1" {
		t.Fatal("Expected <1 Kings 1:1> from parseText(). Got: ", result[1])
	}
	if result[3] != "Exodus 1:1-10" {
		t.Fatal("Expected <Exodus 1:1-10> from parseText(). Got: ", result[1])
	}
}
