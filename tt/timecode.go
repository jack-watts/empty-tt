// Package tt enables a user to create a minimal ST 428-7 Subtitle
// XML Document in accordance with ISDCF Technicial Doc 16 - SMPTE
// ST 428-7 D-Cinema Distribution Master Subtitle - Minimal Empty Document Requirements
// as per said requirements stipulated in RDD 52 - SMPTE DCP Bv2.1 Application Profile
// available at https://doi.org/10.5594/SMPTE.RDD52.2020.
//
/* Copyright (c) 2020, Jack Watts. All rights reserved.

This program is free software : you can redistribute it and / or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

THIS SOFTWARE IS PROVIDED BY THE AUTHOR ``AS IS'' AND ANY EXPRESS OR
IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES
OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT,
INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF
THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE

You should have received a copy of the GNU General Public License
along with this program.If not, see <http://www.gnu.org/licenses/>.*/
package tt

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

var tcRegexp = regexp.MustCompile(`^(\d\d)[:;](\d\d)[:;](\d\d)[:.](\d+)$`)

// divMod returns the floating-point remainder of a/b
func divMod(a float64, b float64) (float64, float64) {
	q := a / b
	r := math.Mod(a, b)
	return q, float64(r)
}

// Timecode struct.
type Timecode struct {
	frameRate   float64
	hours       int
	minutes     int
	seconds     int
	frames      int
	totalFrames int
}

// NewTimecode initialises a new timecode type from a given framerate.
func NewTimecode(frameRate float64) (*Timecode, error) {
	if frameRate >= 0 {
		tc := &Timecode{
			frameRate:   frameRate,
			hours:       0,
			minutes:     0,
			seconds:     0,
			frames:      0,
			totalFrames: 0,
		}
		return tc, nil
	}
	return nil, fmt.Errorf("unsupported framerate")
}

// GetTimeCode method generates a SMPTE timcode when called against type Timecode.
func (tc *Timecode) GetTimeCode() string {
	h := tc.hours
	m := tc.minutes
	s := tc.seconds
	f := tc.frames
	// set 23Pulldown 1
	fr := tc.frameRate
	a, b := math.Modf(fr)
	if b != 0.0 {
		fr = a + 1
	}
	totalFrames := float64(h*3600+m*60+s)*fr + float64(f)
	seconds, frames := divMod(totalFrames, fr)
	minutes, seconds := divMod(seconds, 60)
	hours, minutes := divMod(minutes, 60)
	return fmt.Sprintf("%02d:%02d:%02d:%02d", int(hours), int(minutes), int(seconds), int(frames))
}

// SetFrames sets the frame count in type Timecode.
func (tc *Timecode) SetFrames(frameCount int) {
	tc.totalFrames = frameCount
	seconds, frames := divMod(float64(tc.totalFrames), tc.frameRate)
	minutes, seconds := divMod(float64(seconds), 60)
	hours, minutes := divMod(float64(minutes), 60)
	tc.frames = int(frames)
	tc.seconds = int(seconds)
	tc.minutes = int(minutes)
	tc.hours = int(hours)
}

// getFloat returns a value of type float64 from a given value of type string that contains a whole number.
func getFloat(s string) float64 {
	if s == "" {
		return 0.0
	}
	f, _ := strconv.ParseFloat(s, 64)
	return float64(f)
}
