package gtk

import (
	"testing"
)

func init() {
	Init(nil)
}

func TestTextIterForwardSearch(t *testing.T) {
	table, err := TextTagTableNew()
	if err != nil {
		t.Error("could not create tag table")
	}
	buf, err := TextBufferNew(table)
	if err != nil {
		t.Error("could not create text buffer")
	}
	buf.SetText("This is some text to search")

	start := buf.GetStartIter()
	flags := VISIBLE_ONLY
	expected := "text"

	var limit *TextIter

	// test with no limit set
	matchStart, matchEnd, found := start.ForwardSearch(expected, flags, limit)
	if !found {
		t.Error("(no limit) forward search failed")
	}
	actual, err := buf.GetText(matchStart, matchEnd, false)
	if err != nil {
		t.Error("(no limit) could not get text")
	}
	if actual != expected {
		t.Errorf("(no limit) Expected '%s'; Got '%s'", expected, actual)
	}

	// test with limit set
	limit = buf.GetEndIter()
	matchStart, matchEnd, found = start.ForwardSearch(expected, flags, limit)
	if !found {
		t.Error("(limit) forward search failed")
	}
	actual, err = buf.GetText(matchStart, matchEnd, false)
	if err != nil {
		t.Error("(limit) could not get text")
	}
	if actual != expected {
		t.Errorf("(limit) Expected '%s'; Got '%s'", expected, actual)
	}

	// test for text not found
	expected = "abc"
	matchStart, matchEnd, found = start.ForwardSearch(expected, flags, limit)
	if found {
		t.Error("(not found) forward search false match found")
	}
}

func TestTextIterBackwardSearch(t *testing.T) {
	table, err := TextTagTableNew()
	if err != nil {
		t.Error("could not create tag table")
	}
	buf, err := TextBufferNew(table)
	if err != nil {
		t.Error("could not create text buffer")
	}
	buf.SetText("This is some text to search")

	start := buf.GetEndIter()
	flags := VISIBLE_ONLY
	expected := "text"

	var limit *TextIter

	// test with no limit set
	matchStart, matchEnd, found := start.BackwardSearch(expected, flags, limit)
	if !found {
		t.Error("(no limit) backward search failed")
	}
	actual, err := buf.GetText(matchStart, matchEnd, false)
	if err != nil {
		t.Error("(no limit) could not get text")
	}
	if actual != expected {
		t.Errorf("(no limit) Expected '%s'; Got '%s'", expected, actual)
	}

	// test with limit set
	limit = buf.GetStartIter()
	matchStart, matchEnd, found = start.BackwardSearch(expected, flags, limit)
	if !found {
		t.Error("(limit) backward search failed")
	}
	actual, err = buf.GetText(matchStart, matchEnd, false)
	if err != nil {
		t.Error("(limit) could not get text")
	}
	if actual != expected {
		t.Errorf("(limit) Expected '%s'; Got '%s'", expected, actual)
	}

	// test for text not found
	expected = "abc"
	matchStart, matchEnd, found = start.BackwardSearch(expected, flags, limit)
	if found {
		t.Error("(not found) backward search false match found")
	}
}
