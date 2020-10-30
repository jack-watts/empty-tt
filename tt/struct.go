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
