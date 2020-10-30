package time

import (
	"testing"
	"time"
)

func TestToIso8601MatchPlatformStandards(t *testing.T) {
	tests := []struct {
		input    time.Duration
		expected string
	}{
		{time.Duration(11234928) * time.Millisecond, "PT3H7M14.928S"},
		{time.Duration(11234000) * time.Millisecond, "PT3H7M14S"},
		{time.Duration(189) * time.Second, "PT3M9S"},
		{time.Duration(191) * time.Second, "PT3M11S"},
		{time.Duration(191000) * time.Millisecond, "PT3M11S"},
		{time.Duration(191001) * time.Millisecond, "PT3M11.001S"},
		{time.Duration(191010) * time.Millisecond, "PT3M11.010S"},
		{time.Duration(191123) * time.Millisecond, "PT3M11.123S"},
		{time.Duration(0) * time.Millisecond, "PT0S"},
		{time.Duration(1) * time.Millisecond, "PT0.001S"},
		{time.Duration(10) * time.Millisecond, "PT0.010S"},
		{time.Duration(100) * time.Millisecond, "PT0.100S"},
		{time.Duration(1000) * time.Millisecond, "PT1S"},
	}

	for i, item := range tests {
		if ToIso8601(item.input) != item.expected {
			t.Errorf("%d Found %s, while expecting %s", i, ToIso8601(item.input), item.expected)
		}
	}
}

func TestFromIso8601MatchPlatformStandards(t *testing.T) {
	tests := []struct {
		input    string
		expected time.Duration
	}{
		{"PT3H7M14.928S", time.Duration(11234928) * time.Millisecond},
		{"PT3H7M14S", time.Duration(11234000) * time.Millisecond},
		{"PT1H3M9S", time.Duration(3789000) * time.Millisecond},
		{"PT3M9S", time.Duration(189000) * time.Millisecond},
		{"PT03M9S", time.Duration(189000) * time.Millisecond},
		{"PT03M9.0S", time.Duration(189000) * time.Millisecond},
		{"PT03M9.01S", time.Duration(189001) * time.Millisecond},
		{"PT03M9.10S", time.Duration(189010) * time.Millisecond},
		{"PT3M11S", time.Duration(191000) * time.Millisecond},
		{"PT3M11.001S", time.Duration(191001) * time.Millisecond},
		{"PT3M11.010S", time.Duration(191010) * time.Millisecond},
		{"PT3M11.123S", time.Duration(191123) * time.Millisecond},
		{"PT17H23M39.82S", time.Duration(62619082) * time.Millisecond},
		{"PT52M30S", time.Duration(3150000) * time.Millisecond},
		{"PT52M30.9S", time.Duration(3150009) * time.Millisecond},
		{"PT52M30.96S", time.Duration(3150096) * time.Millisecond},
		{"PT52M30.961S", time.Duration(3150961) * time.Millisecond},
	}

	for i, item := range tests {
		res, err := FromIso8601(item.input)
		if err != nil {
			t.Errorf("%d: %s", i, err.Error())
		} else if res != item.expected {
			t.Errorf("%d Found %d, while expecting %d", i, res.Milliseconds(), item.expected)

		}
	}

}
