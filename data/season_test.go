package data

import "testing"

func TestNextSeason(t *testing.T) {
	var thisSeason Season = "2014-15"

	nextSeason := thisSeason.Next()
	if nextSeason != "2015-16" {
		t.Errorf("Expected next season to be 2015-16, but got `%s`", nextSeason)
	}
}
