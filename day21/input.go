package main

import (
	"fmt"
)

var rawStart = ".#./..#/###"

var rawRules = map[string]string{
	"../..":       "###/###/.##",
	"#./..":       "..#/###/##.",
	"##/..":       "..#/##./##.",
	".#/#.":       "#../.#./.##",
	"##/#.":       "#.#/###/.#.",
	"##/##":       "##./.../.#.",
	".../.../...": "...#/.#../#.#./##.#",
	"#../.../...": ".#.#/.#../####/###.",
	".#./.../...": "#.##/#.##/.###/##.#",
	"##./.../...": "..##/#.##/.##./..##",
	"#.#/.../...": ".#.#/#.#./#..#/...#",
	"###/.../...": "#.../.##./.#../.###",
	".#./#../...": "##.#/...#/##.#/.##.",
	"##./#../...": "#.#./###./...#/#.##",
	"..#/#../...": "..##/.###/..../.##.",
	"#.#/#../...": "...#/#..#/#.#./#.#.",
	".##/#../...": "...#/#.##/..##/.###",
	"###/#../...": ".##./..##/##../##.#",
	".../.#./...": "####/.##./##.#/####",
	"#../.#./...": "..../.##./#..#/##.#",
	".#./.#./...": "..../#.##/#.../..#.",
	"##./.#./...": ".###/.#.#/...#/....",
	"#.#/.#./...": "..##/.#../.###/#.##",
	"###/.#./...": "..../..##/##.#/###.",
	".#./##./...": ".###/.#.#/#..#/#.#.",
	"##./##./...": "#..#/#..#/#.##/.##.",
	"..#/##./...": "#.##/...#/..#./.##.",
	"#.#/##./...": "..##/#.../..../...#",
	".##/##./...": "##.#/...#/..##/#..#",
	"###/##./...": "..##/..#./.###/..##",
	".../#.#/...": ".###/..##/.#.#/..##",
	"#../#.#/...": "..##/...#/##../..#.",
	".#./#.#/...": "..##/##.#/#..#/###.",
	"##./#.#/...": "#.../####/..#./#...",
	"#.#/#.#/...": "..../##.#/.##./#..#",
	"###/#.#/...": "..##/#.#./.#.#/.#..",
	".../###/...": "..##/.#../.#.#/#..#",
	"#../###/...": "#.#./.#../.##./....",
	".#./###/...": "##.#/...#/###./#.##",
	"##./###/...": "..../#.../.###/#.#.",
	"#.#/###/...": "####/..../...#/....",
	"###/###/...": "##.#/##../#.##/#...",
	"..#/.../#..": "##.#/..#./#.##/..#.",
	"#.#/.../#..": ".#../...#/..#./.##.",
	".##/.../#..": "...#/#.../#..#/#..#",
	"###/.../#..": ".###/##../.##./.#..",
	".##/#../#..": "..##/#.##/.#.#/...#",
	"###/#../#..": "...#/.###/..../#..#",
	"..#/.#./#..": "#..#/..../..#./..##",
	"#.#/.#./#..": "#..#/..../#.#./.###",
	".##/.#./#..": "..../.##./..##/.#.#",
	"###/.#./#..": "##.#/###./##.#/..##",
	".##/##./#..": "#.#./..../###./####",
	"###/##./#..": "#..#/#.##/#.##/#...",
	"#../..#/#..": "##../#..#/#.../###.",
	".#./..#/#..": "#.#./.#.#/..../.#.#",
	"##./..#/#..": "#.#./#.../#.#./#..#",
	"#.#/..#/#..": "..##/.#.#/.#../.###",
	".##/..#/#..": "##.#/..##/..../.###",
	"###/..#/#..": "..#./.##./...#/.#.#",
	"#../#.#/#..": "#.../.#../#.#./##..",
	".#./#.#/#..": "..../..../##../#...",
	"##./#.#/#..": "..#./..../#.../..#.",
	"..#/#.#/#..": "#.#./.#.#/.#../#.##",
	"#.#/#.#/#..": "...#/##.#/.##./#...",
	".##/#.#/#..": "..#./...#/.##./#...",
	"###/#.#/#..": "..##/#..#/..../..##",
	"#../.##/#..": "##.#/##.#/#.##/.#.#",
	".#./.##/#..": "..##/##../#.#./####",
	"##./.##/#..": "#.#./..../..##/#.##",
	"#.#/.##/#..": "..#./###./##.#/##.#",
	".##/.##/#..": "#..#/...#/..##/....",
	"###/.##/#..": "..##/##../##.#/#.##",
	"#../###/#..": "####/###./.###/....",
	".#./###/#..": "...#/.##./...#/#.##",
	"##./###/#..": "...#/...#/##.#/.##.",
	"..#/###/#..": "..##/.##./#.#./...#",
	"#.#/###/#..": ".###/.##./.###/.#.#",
	".##/###/#..": "##../.#../#.#./##.#",
	"###/###/#..": "..../..../.###/##..",
	".#./#.#/.#.": "##.#/##.#/..##/.##.",
	"##./#.#/.#.": ".#../#.##/#.##/#.#.",
	"#.#/#.#/.#.": "..##/#.#./#.../..##",
	"###/#.#/.#.": "##.#/.#.#/##.#/.###",
	".#./###/.#.": "#.#./..#./..##/.##.",
	"##./###/.#.": "...#/#.##/###./#.##",
	"#.#/###/.#.": "...#/.###/#.#./#.#.",
	"###/###/.#.": ".#.#/#..#/####/#...",
	"#.#/..#/##.": "#.##/#.#./##../####",
	"###/..#/##.": "##.#/...#/..../####",
	".##/#.#/##.": "#.../#..#/..##/....",
	"###/#.#/##.": "##../###./...#/####",
	"#.#/.##/##.": "##.#/..##/..../#...",
	"###/.##/##.": "..#./####/..../#...",
	".##/###/##.": "..##/#.##/..#./####",
	"###/###/##.": "#.##/...#/..../..#.",
	"#.#/.../#.#": "..#./#.##/#..#/#.#.",
	"###/.../#.#": "..#./###./..##/#...",
	"###/#../#.#": ".###/#..#/##../.#..",
	"#.#/.#./#.#": "###./##.#/.#../#..#",
	"###/.#./#.#": "##.#/###./#.../...#",
	"###/##./#.#": "####/##../#.../....",
	"#.#/#.#/#.#": "..#./..##/..#./...#",
	"###/#.#/#.#": "...#/##.#/##.#/#.##",
	"#.#/###/#.#": "..#./####/.#../##.#",
	"###/###/#.#": "..../.#.#/..../...#",
	"###/#.#/###": "#.#./..##/##.#/....",
	"###/###/###": "..#./#.##/####/###.",
}

type sizedKey struct {
	size int
	key  string
}

func (sk sizedKey) String() string {
	return fmt.Sprintf("%1d-%s", sk.size, sk.key)
}

func parseRules(rawRules map[string]string) map[sizedKey]*bitgrid {
	rules := make(map[sizedKey]*bitgrid)

	for rawSource, rawTarget := range rawRules {
		source := parse(rawSource)
		target := parse(rawTarget)

		all := allTransforms(source)
		for key := range all {
			rules[sizedKey{source.size, key}] = target
		}
	}

	return rules
}

func input() (*bitgrid, map[sizedKey]*bitgrid) {
	start := parse(rawStart)
	rules := parseRules(rawRules)

	return start, rules
}
