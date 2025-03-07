package main

//It's obvious that strings, ints, and other basic types should be available as map keys, but perhaps unexpected are struct keys. Struct can be used to key data by multiple dimensions. For example, this map of maps could be used to tally web page hits by country:
//This is map of string to (map of string to int). Each key of the outer map is the path to a web page with its own inner map. Each inner map key is a two-letter country code. This expression retrieves the number of times an Australian has loaded the documentation page:
hits := make(map[string]map[string]int)
n := hits["/doc/"]["au"]

func add(m map[string]map[string]int, path, country string) {
	mm, ok := m[path]
	if !ok {
		mm = make(map[string]int)
		m[path] = mm
	}
	mm[country]++
}
add(hits, "/doc/", "au")

//Second approach is to use a struct as a key. This is useful when you need to key by multiple dimensions. For example, this map of ints could be used to tally web page hits by path and country:
type Key struct {
	Path, Country string
}
hits := make(map[Key]int)

hits[Key{"/", "vn"}]++
n := hits[Key{"/ref/spec", "ch"}]
