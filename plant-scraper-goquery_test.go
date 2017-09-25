package main

import (
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestFindSingleDetailsLink(t *testing.T) {
	doc, err := goquery.NewDocument("https://www.rhs.org.uk/plants/bulbs")
	if err != nil {
		t.Fatalf("not able to create doc %v", err)
	}
	got := postScrape(doc)
	want := []string{
		"/Plants/132792/i-Eranthis-hyemalis-i-Orange-Glow/Details",
		"/Plants/82514/i-Hyacinthus-orientalis-i-White-Pearl/Details",
		"/Plants/11346/i-Narcissus-asturiensis-i-(13)/Details",
		"/Plants/859/i-Allium-schoenoprasum-i/Details",
		"/Plants/351604/i-Iris-i-Palm-Springs-(Reticulata)/Details",
		"/Plants/57234/i-Narcissus-i-Trevithian-(7)/Details",
		"/Plants/62581/i-Lilium-i-Grand-Cru-(Ia-b)/Details",
		"/Plants/58177/i-Narcissus-i-Ben-Hee-(2)/Details",
		"/Plants/40852/i-Hymenocallis-i-×-i-macrostephana-i/Details",
		"/Plants/269262/i-Narcissus-i-Derringer-(7)/Details"}

	if len(got) != len(want) {
		t.Errorf("wanted %d results but got %d", len(want), len(got))
	}
	for i := 0; i < len(want); i++ {
		if got[i] != want[i] {
			t.Errorf("wanted %v but got %v at %d", got[i], want[i], i)
		}
	}
}
