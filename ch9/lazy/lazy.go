//go:build ignore
// +build ignore

package main

import "image"

var icons map[string]image.Image

func loadIcons() {
	icons = map[string]image.Image{
		"spades.png":   loadIcon("spades.png"),
		"hearts.png":   loadIcon("hearts.png"),
		"diamonds.png": loadIcon("diamonds.png"),
		"clubs.png":    loadIcon("clubs.png"),
	}
}

func Icon(name string) image.Image {
	if icons == nil {
		loadIcons() // initialize only one time
	}
	return icons[name]
}
