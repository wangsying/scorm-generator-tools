package sxml

import (
	"encoding/xml"
)

// XMLManifestNode ...
type XMLManifestNode struct {
	XMLName           xml.Name             `xml:"manifest"`
	Identifier        xml.Attr             `xml:"identifier,attr"`
	Version           xml.Attr             `xml:"version,attr"`
	Xmlns             xml.Attr             `xml:"xmlns,attr,omitempty"`
	XmlnsXsi          xml.Attr             `xml:"xmlns:xsi,attr"`
	XmlnsAdlcp        xml.Attr             `xml:"xmlns adlcp,attr"`
	XmlnsAdlseq       xml.Attr             `xml:"xmlns adlseq,attr"`
	XmlnsAdlnav       xml.Attr             `xml:"xmlns adlnav,attr"`
	XmlnsImsss        xml.Attr             `xml:"xmlns imsss,attr"`
	XsiSchemaLocation xml.Attr             `xml:"schemaLocation,attr"`
	MetadataNode      XMLMetadataType      `xml:"metadata"`
	OrganizationNode  XMLOrganizationsNode `xml:"organizations"`
	ResourceNode      XMLResourcesNode     `xml:"resources"`
}

// XMLMetadataType ...
type XMLMetadataType struct {
	XMLName       xml.Name `xml:"metadata"`
	Schema        string   `xml:"schema"`
	Schemaversion string   `xml:"schemaversion"`
}

// XMLOrganizationsNode ...
type XMLOrganizationsNode struct {
	XMLName       xml.Name          `xml:"organizations"`
	Default       xml.Attr          `xml:"default,attr"`
	Organizations []XMLOrganization `xml:"organization"`
}

// XMLOrganization ...
type XMLOrganization struct {
	XMLName    xml.Name      `xml:"organization"`
	Identifier xml.Attr      `xml:"identifier,attr"`
	Title      string        `xml:"title"`
	Items      []XMLItemType `xml:"item"`
}

// XMLItemType ...
type XMLItemType struct {
	XMLName       xml.Name      `xml:"item"`
	Identifier    xml.Attr      `xml:"identifier,attr"`
	Identifierref xml.Attr      `xml:"identifierref,attr"`
	Title         string        `xml:"title"`
	Items         []XMLItemType `xml:"item"`
}

// XMLResourcesNode ...
type XMLResourcesNode struct {
	XMLName  xml.Name          `xml:"resources"`
	Resource []XMLResourceType `xml:"resource"`
}

// XMLResourceType ...
type XMLResourceType struct {
	XMLName    xml.Name            `xml:"resource"`
	Identifier xml.Attr            `xml:"identifier,attr"`
	Type       xml.Attr            `xml:"type,attr"`
	ScormType  xml.Attr            `xml:"scormType,attr"`
	Href       xml.Attr            `xml:"href,attr"`
	Files      []XMLFileType       `xml:"file"`
	Dependency []XMLDependencyType `xml:"dependency"`
}

// XMLFileType ...
type XMLFileType struct {
	XMLName xml.Name `xml:"file"`
	Href    xml.Attr `xml:"href,attr"`
}

// XMLDependencyType ...
type XMLDependencyType struct {
	XMLName       xml.Name `xml:"dependency"`
	Identifierref xml.Attr `xml:"identifierref,attr"`
}
