package problem41

import (
	"fmt"
	"sort"
	"strings"
)

// List contains items.
type List struct {
	items []string
}

// NewList returns new list with items.
func NewList(items []string) *List {
	return &List{items}
}

// Pop removes element by given index.
// If index is -1 removes list item.
func (l *List) Pop(i int) {
	if i == -1 {
		l.items = l.items[:len(l.items)-1]
		return
	}
	l.items = append(l.items[:i], l.items[i+1:]...)
}

// Append inserts new element.
func (l *List) Append(el string) {
	l.items = append(l.items, el)
}

// Extend appends list of items.
func (l *List) Extend(els []string) {
	l.items = append(l.items, els...)
}

// Insert appends element at the given index.
func (l *List) Insert(i int, el string) {
	right := l.items[i:]
	l.items = append(l.items[:i], el)
	l.items = append(l.items, right...)
}

// Enum returns list's items.
func (l *List) Enum() []string {
	return l.items
}

// Flight holds origin and destination point.
type Flight struct {
	Origin      string
	Destination string
}

func (f Flight) String() string {
	return fmt.Sprintf("('%s', '%s')", f.Origin, f.Destination)
}

func visit(flightMap map[string]*List, totalFlights int, currentItinerary []string) [][]string {
	// If our itinerary uses up all the flights, we're done here.
	if len(currentItinerary) == totalFlights+1 {
		return [][]string{currentItinerary}
	}

	lastStop := currentItinerary[len(currentItinerary)-1]
	// If we haven't used all the flights yet but we have no way
	// of getting out of this airport, then we're stuck. Backtrack out.
	if _, ok := flightMap[lastStop]; !ok {
		return [][]string{}
	}

	// Otherwise, let's try all the options out of the current stop recursively.
	// We temporarily take them out of the mapping once we use them.
	potentialItineraries := [][]string{}
	for i, flight := range flightMap[lastStop].Enum() {
		flightMap[lastStop].Pop(i)
		currentItinerary = append(currentItinerary, flight)
		potentialItineraries = append(potentialItineraries, visit(flightMap, totalFlights, currentItinerary)...)
		flightMap[lastStop].Insert(i, flight)
		currentItinerary = currentItinerary[:len(currentItinerary)-1]
	}
	return potentialItineraries
}

func itinerary(flights []Flight, start string) []string {
	flightMap := make(map[string]*List)
	for _, flight := range flights {
		flightMap[flight.Origin] = NewList([]string{flight.Destination})
	}
	validItineraries := visit(flightMap, len(flights), []string{start})
	if len(validItineraries) > 0 {
		sort.SliceStable(validItineraries, func(i, j int) bool {
			left := strings.Join(validItineraries[i], "")
			right := strings.Join(validItineraries[j], "")
			return left < right
		})
		return validItineraries[0]
	}
	return []string{}
}
