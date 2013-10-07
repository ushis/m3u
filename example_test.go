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

package m3u_test

import (
  "fmt"
  "github.com/ushis/m3u"
  "os"
)

func ExampleParse() {
  f, err := os.Open("/path/to/playlist.m3u")

  if err != nil {
    panic(err)
  }
  defer f.Close()

  pl, err := m3u.Parse(f)

  if err != nil {
    panic(err)
  }

  for _, track := range pl {
    fmt.Printf("%s (%s) (%d)\n", track.Path, track.Title, track.Time)
  }
}

func ExamplePlaylist_WriteTo() {
  f, err := os.Create("/path/to/playlist.m3u")

  if err != nil {
    panic(err)
  }
  defer f.Close()

  pl := m3u.Playlist{
    m3u.Track{
      Path:  "Kraftwerk/Kraftwerk/Megaherz.mp3",
      Title: "Megaherz",
      Time:  573,
    },
    m3u.Track{
      Path:  "https://music.com/stream.ogg",
      Title: "Music Stream",
      Time:  -1,
    },
  }

  if _, err := pl.WriteTo(f); err != nil {
    panic(err)
  }
}
