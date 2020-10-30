package time

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"time"
)

// ToIso8601 returns a string from duration.
// It is not fully compatible with the standard, however
//  it will work with the protocol we agreed on.
func ToIso8601(d time.Duration) (s string) {
	sec := toIso8601Seconds(d)

	if d > time.Hour {
		min := math.Mod(d.Minutes(), 60)
		s = fmt.Sprintf("PT%.0fH%.0fM%s", d.Hours(), min, sec)
		return
	}
	if d > time.Minute {
		s = fmt.Sprintf("PT%.0fM%s", d.Minutes(), sec)
		return
	}
	if d > time.Second {
		s = fmt.Sprintf("PT%s", sec)
		return
	}

	// Can implement for milliseconds
	s = fmt.Sprintf("PT%s", sec)

	return s
}

func toIso8601Seconds(d time.Duration) (s string) {
	sec := math.Mod(d.Seconds(), 60)
	_, frac := math.Modf(sec)
	if frac < 1e-7 {
		s = fmt.Sprintf("%.0fS", sec)
	} else {
		s = fmt.Sprintf("%.3fS", sec)
	}

	return s
}

// Compile once
var reDuration = regexp.MustCompile(`P(\.+)?T((?P<hours>\d{1,2})H)?((?P<minutes>\d{1,2})M)?(((?P<seconds>\d{1,2})(\.(?P<milliseconds>\d{1,3}))?)S)?`)

func FromIso8601(s string) (t time.Duration, err error) {
	var num int
	var matches []string = reDuration.FindStringSubmatch(s)
	var keys []string = reDuration.SubexpNames()

	for i, item := range matches {
		if keys[i] == "" || item == "" {
			continue
		}

		num, err = namedMatchToInt(keys[i], item)
		if err != nil {
			return
		}
		t += time.Duration(num) * time.Millisecond
	}

	return t, err
}

func namedMatchToInt(name, match string) (n int, err error) {
	n, err = strconv.Atoi(match)
	if err != nil {
		return
	}

	switch name {
	case "hours":
		n *= 60 * 60 * 1000
		break
	case "minutes":
		n *= 60 * 1000
		break
	case "seconds":
		n *= 1000
		break
	case "milliseconds":
		break
	default:
		err = errors.New("Invalid named submatch key")
		return
	}

	return n, err
}
