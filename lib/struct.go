package lib

import (
	"encoding/xml"
)

var (
	dcst2010      string = "dcst2010"
	dcst2014      string = "dcst2014"
	XmlNsSubtitle        = map[string]string{
		"http://www.smpte-ra.org/schemas/428-7/2010/DCST": dcst2010,
		"http://www.smpte-ra.org/schemas/428-7/2014/DCST": dcst2014,
	}
)

// BEGIN SUBTITLE STRUCT //

type SubtitleReel struct {
	XMLName          xml.Name  `xml:"SubtitleReel"`
	Xmlns            string    `xml:"xmlns,attr" json:"xmlns,attr,omitempty"`
	Dcst             string    `xml:"xmlns:dcst,attr" json:"dcst,attr,omitempty"`
	ID               string    `xml:"Id"`
	ContentTitleText string    `xml:"ContentTitleText"`
	IssueDate        string    `xml:"IssueDate"`
	ReelNumber       int       `xml:"ReelNumber"`
	Language         string    `xml:"Language"`
	EditRate         string    `xml:"EditRate"`
	TimeCodeRate     string    `xml:"TimeCodeRate"`
	StartTime        string    `xml:"StartTime"`
	DisplayType      string    `xml:"DisplayType"`
	LoadFont         *LoadFont `xml:"LoadFont"`
	SubtitleList     *Font     `xml:"SubtitleList>Font"`
	Filename         string
}

type LoadFont struct {
	ID       string `xml:"ID,attr"`
	Font     string `xml:",chardata"`
	Filename string
	Fontpath string
	Size     int64
}

type Font struct {
	ID           string      `xml:"ID,attr"`
	Weight       string      `xml:"Weight,attr,omitempty"`
	Size         string      `xml:"Size,attr"`
	Color        string      `xml:"Color,attr"`
	Effect       string      `xml:"Effect,attr,omitempty"`
	EffectColor  string      `xml:"EffectColor,attr,omitempty"`
	EffectSize   string      `xml:"EffectSize,attr"`
	Italic       string      `xml:"Italic,attr"`
	Underline    string      `xml:"Underline,attr"`
	AspectAdjust string      `xml:"AspectAdjust,attr"`
	Spacing      string      `xml:"Spacing,attr"`
	Feather      string      `xml:"Feather,attr"`
	Subtitle     []*Subtitle `xml:"Subtitle"`
}

type Subtitle struct {
	SpotNumber   string      `xml:"SpotNumber,attr"`
	TimeIn       string      `xml:"TimeIn,attr"`
	TimeOut      string      `xml:"TimeOut,attr"`
	FadeUpTime   string      `xml:"FadeUpTime,attr" json:"FadeUpTime,omitempty"`
	FadeDownTime string      `xml:"FadeDownTime,attr" json:"FadeDownTime,omitempty"`
	Text         []*Text     `xml:"Text" json:"Text,omitempty"`
	Image        []*Image    `xml:"Image" json:"Image,omitempty"`
	Font         *NestedFont `xml:"Font" json:"Font,omitempty"`
}

type NestedFont struct {
	ID           string `xml:"ID,attr"`
	Weight       string `xml:"Weight,attr,omitempty"`
	Size         string `xml:"Size,attr"`
	Color        string `xml:"Color,attr"`
	Effect       string `xml:"Effect,attr,omitempty"`
	EffectColor  string `xml:"EffectColor,attr,omitempty"`
	EffectSize   string `xml:"EffectSize,attr"`
	Italic       string `xml:"Italic,attr"`
	Underline    string `xml:"Underline,attr"`
	AspectAdjust string `xml:"AspectAdjust,attr"`
	Spacing      string `xml:"Spacing,attr"`
	Feather      string `xml:"Feather,attr"`
	Text         string `xml:",chardata"`
}

type Text struct {
	Text      string      `xml:",chardata"`
	Halign    string      `xml:"Halign,attr"`
	Hposition string      `xml:"Hposition,attr"`
	Valign    string      `xml:"Valign,attr"`
	Vposition string      `xml:"Vposition,attr"`
	Direction string      `xml:"Direction,attr,omitempty" json:"Direction,omitempty"`
	Zposition string      `xml:"Zposition,attr,omitempty" json:"Zposition,omitempty"`
	VariableZ string      `xml:"VariableZ,attr,omitempty" json:"VariableZ,omitempty"`
	Font      *NestedFont `xml:"Font,omitempty" json:"Font,omitempty"`
	Ruby      []*Ruby     `xml:"Ruby,omitempty" json:"Ruby,omitempty"`
}

type Image struct {
	XMLName   xml.Name `xml:"Image"`
	ImageType string   `xml:"ImageType"`
	Halign    string   `xml:"Halign,attr"`
	Hposition string   `xml:"Hposition,attr"`
	Valign    string   `xml:"Valign,attr"`
	Vposition string   `xml:"Vposition,attr"`
	Zposition string   `xml:"Zposition,attr,omitempty"`
	VariableZ string   `xml:"VariableZ,attr,omitempty"`
}

type Ruby struct {
	Rb string `xml:"Ruby>Rb,omitempty"`
	Rt *Rt    `xml:"Ruby>Rt,omitempty"`
}

type Rt struct {
	Size         string `xml:"Size,attr"`
	Position     string `xml:"Position,attr"`
	Offset       string `xml:"Offset,attr"`
	Spacing      string `xml:"Spacing,attr"`
	AspectAdjust string `xml:"AspectAdjust,attr"`
}

// END SUBTITLE STRUCT //
