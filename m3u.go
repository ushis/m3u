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
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Represents a list of tracks.
type Playlist []Track

// Represents a single track.
type Track struct {
	Path  string // path to the file
	Title string // title of the track
	Time  int64  // duration of the track
}

// Parses simple and extended m3u files. Returns the playlist.
func Parse(r io.Reader) (Playlist, error) {
	scanner := bufio.NewScanner(r)
	pl := Playlist{}

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 && line[0] != '#' {
			pl = append(pl, Track{Path: line, Title: "", Time: -1})
			continue
		}

		if len(line) > 8 && line[:8] == "#EXTINF:" {
			i := strings.Index(line[8:], ",")

			if i < 0 {
				return pl, fmt.Errorf("unexpected line: %q", line)
			}
			ftime, err := strconv.ParseFloat(line[8:i+8], 64)

			if err != nil {
				return pl, err
			}
			time := int64(ftime)
			scanner.Scan()
			path := scanner.Text()

			if err := scanner.Err(); err != nil {
				return pl, err
			}
			pl = append(pl, Track{Path: path, Title: line[i+9:], Time: time})
		}
	}
	if err := scanner.Err(); err != nil {
		return pl, err
	}
	return pl, nil
}

// Writes the playlist to a writer in the extended m3u format. Returns the
// number of written bytes.
func (pl Playlist) WriteTo(w io.Writer) (int64, error) {
	var buf bytes.Buffer
	buf.WriteString("#EXTM3U\n")

	for _, t := range pl {
		time := t.Time

		if time < 1 {
			time = -1
		}
		var b bytes.Buffer
		strconv.AppendInt(b.Bytes(), time, 10)
		buf.WriteString("#EXTINF:")
		buf.Write(b.Bytes())
		buf.WriteString(",")
		buf.WriteString(t.Title)
		buf.WriteString("\n")
		buf.WriteString(t.Path)
		buf.WriteString("\n")
	}

	// Write the buffer to the output.
	n, err := buf.WriteTo(w)
	if err != nil {
		return 0, err
	}

	return n, nil
}

// Writes the playlist to a writer in the simple m3u format. Returns the number
// of written bytes.
func (pl Playlist) WriteSimpleTo(w io.Writer) (int64, error) {
	var buf bytes.Buffer

	for _, t := range pl {
		buf.WriteString(t.Path)
		buf.WriteString("\n")
	}

	// Write the buffer to the output.
	n, err := buf.WriteTo(w)
	if err != nil {
		return 0, err
	}

	return n, nil
}
