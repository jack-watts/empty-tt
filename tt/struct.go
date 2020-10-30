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
along with this program.If not, see http://www.gnu.org/licenses.*/
package tt

import (
	"encoding/xml"
)

// BEGIN ST 428-7 SUBTITLE STRUCT //

// SubtitleReel as per http://www.smpte-ra.org/schemas/428-7/2014/DCST
type SubtitleReel struct {
	XMLName          xml.Name  `xml:"SubtitleReel"`
	Xmlns            string    `xml:"xmlns,attr,omitempty"`
	ID               string    `xml:"Id"`
	ContentTitleText string    `xml:"ContentTitleText,omitempty"`
	IssueDate        string    `xml:"IssueDate"`
	ReelNumber       int       `xml:"ReelNumber"`
	Language         string    `xml:"Language"`
	EditRate         string    `xml:"EditRate"`
	TimeCodeRate     string    `xml:"TimeCodeRate"`
	StartTime        string    `xml:"StartTime"`
	DisplayType      string    `xml:"DisplayType"`
	LoadFont         *LoadFont `xml:"LoadFont,omitempty"`
	SubtitleList     *Font     `xml:"SubtitleList>Font"`
	Filename         string    `xml:",omitempty"`
}

// LoadFont as per http://www.smpte-ra.org/schemas/428-7/2014/DCST#LoadFont
type LoadFont struct {
	ID   string `xml:"ID,attr"`
	Font string `xml:",chardata"`
}

// Font as per http://www.smpte-ra.org/schemas/428-7/2014/DCST#Font
type Font struct {
	ID           string      `xml:"ID,attr,omitempty"`
	Weight       string      `xml:"Weight,attr,omitempty"`
	Size         string      `xml:"Size,attr,omitempty"`
	Color        string      `xml:"Color,attr,omitempty"`
	Effect       string      `xml:"Effect,attr,omitempty"`
	EffectColor  string      `xml:"EffectColor,attr,omitempty"`
	EffectSize   string      `xml:"EffectSize,attr,omitempty"`
	Italic       string      `xml:"Italic,attr,omitempty"`
	Underline    string      `xml:"Underline,attr,omitempty"`
	AspectAdjust string      `xml:"AspectAdjust,attr,omitempty"`
	Spacing      string      `xml:"Spacing,attr,omitempty"`
	Feather      string      `xml:"Feather,attr,omitempty"`
	Subtitle     []*Subtitle `xml:"Subtitle"`
}

// Subtitle as per http://www.smpte-ra.org/schemas/428-7/2014/DCST#Subtitle
type Subtitle struct {
	SpotNumber   string      `xml:"SpotNumber,attr,omitempty"`
	TimeIn       string      `xml:"TimeIn,attr"`
	TimeOut      string      `xml:"TimeOut,attr"`
	FadeUpTime   string      `xml:"FadeUpTime,attr,omitempty"`
	FadeDownTime string      `xml:"FadeDownTime,attr,omitempty"`
	Text         []*Text     `xml:"Text,allowempty"`
	Image        []*Image    `xml:"Image,omitempty"`
	Font         *NestedFont `xml:"Font,omitempty"`
}

// NestedFont as per http://www.smpte-ra.org/schemas/428-7/2014/DCST#NestedFont
type NestedFont struct {
	ID           string `xml:"ID,attr,omitempty"`
	Weight       string `xml:"Weight,attr,omitempty"`
	Size         string `xml:"Size,attr,omitempty"`
	Color        string `xml:"Color,attr,omitempty"`
	Effect       string `xml:"Effect,attr,omitempty"`
	EffectColor  string `xml:"EffectColor,attr,omitempty"`
	EffectSize   string `xml:"EffectSize,attr,omitempty"`
	Italic       string `xml:"Italic,attr,omitempty"`
	Underline    string `xml:"Underline,attr,omitempty"`
	AspectAdjust string `xml:"AspectAdjust,attr,omitempty"`
	Spacing      string `xml:"Spacing,attr,omitempty"`
	Feather      string `xml:"Feather,attr,omitempty"`
	Text         string `xml:",chardata"`
}

// Text as per http://www.smpte-ra.org/schemas/428-7/2014/DCST#Text
type Text struct {
	Text      string      `xml:",chardata"`
	Halign    string      `xml:"Halign,attr,omitempty"`
	Hposition string      `xml:"Hposition,attr,omitempty"`
	Valign    string      `xml:"Valign,attr,omitempty"`
	Vposition string      `xml:"Vposition,attr,omitempty"`
	Direction string      `xml:"Direction,attr,omitempty"`
	Zposition string      `xml:"Zposition,attr,omitempty"`
	VariableZ string      `xml:"VariableZ,attr,omitempty"`
	Font      *NestedFont `xml:"Font,omitempty"`
	Ruby      []*Ruby     `xml:"Ruby,omitempty"`
}

// Image as per http://www.smpte-ra.org/schemas/428-7/2014/DCST#Image
type Image struct {
	XMLName   xml.Name `xml:"Image"`
	Image     string   `xml:",chardata"`
	ImageType string   `xml:"ImageType,omitempty"`
	Halign    string   `xml:"Halign,attr,omitempty"`
	Hposition string   `xml:"Hposition,attr,omitempty"`
	Valign    string   `xml:"Valign,attr,omitempty"`
	Vposition string   `xml:"Vposition,attr,omitempty"`
	Zposition string   `xml:"Zposition,attr,omitempty"`
	VariableZ string   `xml:"VariableZ,attr,omitempty"`
}

// Ruby as per http://www.smpte-ra.org/schemas/428-7/2014/DCST#Ruby
type Ruby struct {
	Rb string `xml:"Ruby>Rb,omitempty"`
	Rt *Rt    `xml:"Ruby>Rt,omitempty"`
}

// Rt as per http://www.smpte-ra.org/schemas/428-7/2014/DCST#Rt
type Rt struct {
	Size         string `xml:"Size,attr"`
	Position     string `xml:"Position,attr"`
	Offset       string `xml:"Offset,attr"`
	Spacing      string `xml:"Spacing,attr"`
	AspectAdjust string `xml:"AspectAdjust,attr"`
}

// END ST 428-7 SUBTITLE STRUCT //
