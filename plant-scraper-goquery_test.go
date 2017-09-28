package main

import (
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestFindSingleDetailsLink(t *testing.T) {
	doc := getDoc(t, "https://www.rhs.org.uk/plants/bulbs")
	got := findPlantURLs(doc)
	want := []string{
		"https://www.rhs.org.uk/Plants/132792/i-Eranthis-hyemalis-i-Orange-Glow/Details",
		"https://www.rhs.org.uk/Plants/82514/i-Hyacinthus-orientalis-i-White-Pearl/Details",
		"https://www.rhs.org.uk/Plants/11346/i-Narcissus-asturiensis-i-(13)/Details",
		"https://www.rhs.org.uk/Plants/859/i-Allium-schoenoprasum-i/Details",
		"https://www.rhs.org.uk/Plants/351604/i-Iris-i-Palm-Springs-(Reticulata)/Details",
		"https://www.rhs.org.uk/Plants/57234/i-Narcissus-i-Trevithian-(7)/Details",
		"https://www.rhs.org.uk/Plants/62581/i-Lilium-i-Grand-Cru-(Ia-b)/Details",
		"https://www.rhs.org.uk/Plants/58177/i-Narcissus-i-Ben-Hee-(2)/Details",
		"https://www.rhs.org.uk/Plants/40852/i-Hymenocallis-i-×-i-macrostephana-i/Details",
		"https://www.rhs.org.uk/Plants/269262/i-Narcissus-i-Derringer-(7)/Details"}

	if len(got) != len(want) {
		t.Errorf("wanted %d results but got %d", len(want), len(got))
	}
	for i := 0; i < len(want); i++ {
		if got[i] != want[i] {
			t.Errorf("wanted %v but got %v at %d", got[i], want[i], i)
		}
	}
}

func TestFind(t *testing.T) {
	doc := getDoc(t, "https://www.rhs.org.uk/Plants/132792/i-Eranthis-hyemalis-i-Orange-Glow/Details")
	got := findPlantName(doc)
	want := "Eranthis hyemalis 'Orange Glow'"

	if got != want {
		t.Errorf("wanted %v but got %v", want, got)
	}
}

func getDoc(t *testing.T, URL string) *goquery.Document {
	doc, err := goquery.NewDocument(URL)
	if err != nil {
		t.Fatalf("not able to create doc %v", err)
	}
	return doc
}
