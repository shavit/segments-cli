package stream

import (
	"encoding/xml"
	"sort"
	"time"

	segTime "github.com/shavit/segments-cli/time"
)

// Silences holds a silence list
type Silences struct {
	XMLName xml.Name   `xml:"silences"`
	Nodes   []*Silence `xml:"silence"`
}

// Decodesilences initialize silences 
func DecodeSilences(b []byte) (s *Silences, err error) {
	xml.Unmarshal(b, &s)
	for _, item := range s.Nodes {
		item.Init()
	}

	return s, err
}

// Sort sorts the silences by their start period
func (s *Silences) Sort() {
	sort.Sort(s)
}

// Len implements sort.Interface for Silences
func (s *Silences) Len() int {
	return len(s.Nodes)
}

// Swap implements sort.Interface for Silences
func (s *Silences) Swap(i, j int) {
	s.Nodes[i], s.Nodes[j] = s.Nodes[j], s.Nodes[i]
}

// Less implements sort.Interface for Silences
func (s *Silences) Less(i, j int) bool {
	return s.Nodes[i].from_ < s.Nodes[j].from_
}

type Silence struct {
	XmlName xml.Name `xml:"silence"`
	From    string   `xml:"from,attr"`
	Until   string   `xml:"until,attr"`
	from_   time.Duration
	until_  time.Duration
	Period  time.Duration
}

// Init is used to calculate durations
func (s *Silence) Init() (err error) {
	s.from_, err = segTime.FromIso8601(s.From)
	if err != nil {
		return
	}

	s.until_, err = segTime.FromIso8601(s.Until)
	if err != nil {
		return
	}

	s.Period = s.until_ - s.from_

	return err
}

// Split is used to split long segments
func (s *Silence) Split(every int) (nodes []*Silence) {
	ms := time.Duration(every) * time.Millisecond
	n := int(s.Period / ms)
	from := s.from_
	until := from + ms

	for i := 0; i < n; i++ {
		node := &Silence{
			from_:  from,
			until_: until,
			From:   segTime.ToIso8601(from),
			Until:  segTime.ToIso8601(until),
			Period: until - from,
		}
		nodes = append(nodes, node)
		from = until
		until += ms
	}

	return nodes
}
