// The MIT License (MIT)
//
// Copyright (c) 2013 ushi <ushi@honkgong.info>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package m3u

import (
	"bytes"
	"io"
	"os"
	"testing"
)

var extended = Playlist{
	Track{
		Path:  "Alternative\\everclear_SMFTA.mp3",
		Title: "Everclear - So Much For The Afterglow",
		Time:  233,
	},
	Track{
		Path:  "Comedy/Weird_Al_Everything_You_Know_Is_Wrong.mp3",
		Title: "",
		Time:  227,
	},
	Track{
		Path:  "Weird_Al_This_Is_The_Life.mp3",
		Title: "Weird Al Yankovic - This is the Life",
		Time:  187,
	},
	Track{
		Path:  "http://www.site.com/~user/gump.mp3",
		Title: "Weird Al: Bad Hair Day - Gump",
		Time:  129,
	},
	Track{
		Path:  "http://www.site.com:8000/listen.pls",
		Title: "My Cool Stream",
		Time:  -1,
	},
}

var simple = Playlist{
	Track{Time: -1, Title: "", Path: "Alternative\\everclear_SMFTA.mp3"},
	Track{Time: -1, Title: "", Path: "Comedy/Weird_Al_Everything_You_Know_Is_Wrong.mp3"},
	Track{Time: -1, Title: "", Path: "Weird_Al_This_Is_The_Life.mp3"},
	Track{Time: -1, Title: "", Path: "http://www.site.com/~user/gump.mp3"},
	Track{Time: -1, Title: "", Path: "http://www.site.com:8000/listen.pls"},
}

func assertPlaylist(t *testing.T, a, b Playlist) {
	if len(a) != len(b) {
		t.Fatalf("Result:   %v\nExpected: %v\n", a, b)
	}

	for i := range a {
		if a[i].Path != b[i].Path || a[i].Title != b[i].Title || a[i].Time != b[i].Time {
			t.Fatalf("\nResult:   %v\nExpected: %v\n", a, b)
		}
	}
}

func parse(t *testing.T, path string) Playlist {
	f, err := os.Open(path)

	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	pl, err := Parse(f)

	if err != nil {
		t.Fatal(err)
	}
	return pl
}

func writeAndParse(t *testing.T, w func(io.Writer) (int64, error)) Playlist {
	var buf bytes.Buffer

	if _, err := w(&buf); err != nil {
		t.Fatal(err)
	}
	pl, err := Parse(&buf)

	if err != nil {
		t.Fatal(err)
	}
	return pl
}

func TestParse(t *testing.T) {
	assertPlaylist(t, parse(t, "testdata/extended.m3u"), extended)
}

func TestParseSimple(t *testing.T) {
	assertPlaylist(t, parse(t, "testdata/simple.m3u"), simple)
}

func TestWriteTo(t *testing.T) {
	assertPlaylist(t, writeAndParse(t, extended.WriteTo), extended)
}

func TestWriteSimpleTo(t *testing.T) {
	assertPlaylist(t, writeAndParse(t, extended.WriteSimpleTo), simple)
}
