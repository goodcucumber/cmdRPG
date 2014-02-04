package Map

import "../Event"

type Map struct {
	Title  string
	Map    []string
	Events []*Event.Event
}
