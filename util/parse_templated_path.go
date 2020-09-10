package util

import (
	"regexp"
	"strings"
)

// ParseTemplatedPath parses a url-like path using an input template and outputs a map of matched values.
// template should place the named params to be extracted in single braces, e.g.
// ParseTemplatedPath("{my_param1}/literal/path/{my_param2}/parts", "foo/literal/path/bar/parts")
// returns {"my_param1": "foo", "my_param2": "bar"}
func ParseTemplatedPath(template string, path string) (map[string]string, error) {
	// escape any needed strings for regex
	template = regexp.QuoteMeta(template)
	// convert braces to capture groups
	template = strings.Replace(template, `\{`, `(?P<`, -1)
	template = "^" + strings.Replace(template, `\}`, `>[a-zA-z0-9_\-]+)`, -1)
	rex, err := regexp.Compile(template)
	if err != nil {
		return nil, err
	}
	match := rex.FindStringSubmatch(path)
	names := rex.SubexpNames()
	params := map[string]string{}
	for i, name := range names {
		if i != 0 && len(match) > i {
			params[name] = match[i]
		}
	}
	return params, nil
}
