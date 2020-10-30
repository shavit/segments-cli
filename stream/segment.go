package stream

import (
	"fmt"
	"time"

	segTime "github.com/shavit/segments-cli/time"
)

type Segments struct {
	Segments []*Segment `json:"segments"`
}

// CreateChapters creates chapters from silences
//
// tMill - The silence in ms between chapters
// maxSegMill - Maximum segment length in ms
// parSegMill - Spilt duration in ms, for each segment.
//
// If segement is longer than permitted, it will be divided into number
//  of segments, each with maximum duration from the `parSegMill` value.
func CreateChapters(silences *Silences, tMill, maxSegMill, parSegMill int) (segments *Segments) {
	// validate the segments
	makeSegments(silences.Nodes, maxSegMill, parSegMill)
	silences.Sort()

	// create chapters
	segments = createChapters(silences.Nodes, tMill)

	return segments
}

func makeSegments(nodes []*Silence, maxSeg, div int) {
	var max time.Duration = time.Duration(maxSeg) * time.Millisecond
	var splitted []*Silence

	for _, node := range nodes {
		if node.Period > max {
			splitted = append(splitted, node.Split(div)...)
		}
	}
	nodes = append(nodes, splitted...)
}

func createChapters(nodes []*Silence, tMill int) (segments *Segments) {
	segments = new(Segments)
	max := time.Duration(tMill) * time.Millisecond
	chapIndex := 0
	segIndex := 0
	offset := time.Duration(0) * time.Millisecond

	for _, node := range nodes {
		if node.Period > max {
			chapIndex += 1
			segIndex = 1
		} else {
			segIndex += 1
		}

		seg := &Segment{
			Title:    fmt.Sprintf("Chapter %d, part %d", chapIndex, segIndex),
			Offset:   segTime.ToIso8601(offset),
			duration: node.Period,
			offset_:  offset,
		}
		segments.Segments = append(segments.Segments, seg)
		offset += node.Period
	}

	return segments
}

type Segment struct {
	Title    string        `json:"title"`
	Offset   string        `json:"offset"`
	duration time.Duration `json:"-"`
	offset_  time.Duration `json:"-"`
}
