package tpl

import (
	"encoding/xml"
)

type schema struct {
	XMLName xml.Name `xml:"template"`
	Content string   `xml:",chardata"`
	Embed   *embed   `xml:"embed,omitempty"`
}

type embed struct {
	XMLName     xml.Name `xml:"embed"`
	Color       *string  `xml:"color,attr,omitempty"`
	Title       *string  `xml:"title,omitempty"`
	Footer      *string  `xml:"footer,omitempty"`
	Description string   `xml:",chardata"`
	Fields      *[]field `xml:"fields>field,omitempty"`
}

type field struct {
	XMLName xml.Name `xml:"field"`
	Name    string   `xml:"name,attr"`
	Inline  bool     `xml:"inline,attr"`
	Value   string   `xml:",chardata"`
}
