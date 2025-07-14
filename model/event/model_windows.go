//go:build windows
// +build windows

// event/model.go
package eventModel

import (
	commonModel "github.com/enki-polvo/polvo-logger/model"
	"github.com/mitchellh/mapstructure"
)

// Event defines the interface for all event types.
type Event any

type Metadata interface {
		ProcessCreationMetadata |
		ProcessAccessMetadata |
		ProcessTerminatedMetadata |
		NetworkConnectionMetadata |
		DnsQueryMetadata |
		DriverLoadMetadata |
		ImageLoadMetadata |
		FileEventMetadata |
		FileDeleteMetadata |
		RegistryEventMetadata |
		RegistryAddMetadata |
		RegistryDeleteMetadata |
		RegistrySetMetadata |
		CreateStreamHashMetadata
}

// --------------------------------------------------
// Event Metadata
// --------------------------------------------------

type ProcessCreationMetadata struct {
	CommandLine       string `json:"CommandLine" mapstructure:"CommandLine"`
	Company           string `json:"Company" mapstructure:"Company"`
	CurrentDirectory  string `json:"CurrentDirectory" mapstructure:"CurrentDirectory"`
	FileVersion       string `json:"FileVersion" mapstructure:"FileVersion"`
	Hashes            string `json:"Hashes" mapstructure:"Hashes"`
	Image             string `json:"Image" mapstructure:"Image"`
	IntegrityLevel    string `json:"IntegrityLevel" mapstructure:"IntegrityLevel"`
	LogonId           string `json:"LogonId" mapstructure:"LogonId"`
	OriginalFileName  string `json:"OriginalFileName" mapstructure:"OriginalFileName"`
	ParentCommandLine string `json:"ParentCommandLine" mapstructure:"ParentCommandLine"`
	ParentImage       string `json:"ParentImage" mapstructure:"ParentImage"`
	ParentProcessId   int64  `json:"ParentProcessId" mapstructure:"ParentProcessId"`
	ProcessId         int64  `json:"ProcessId" mapstructure:"ProcessId"`
	Product           string `json:"Product" mapstructure:"Product"`
	User              string `json:"User" mapstructure:"User"`
}

type ProcessAccessMetadata struct {
	SourceImage     string `json:"SourceImage" mapstructure:"SourceImage"`
	SourceProcessId int64  `json:"SourceProcessId" mapstructure:"SourceProcessId"`
	SourceUser      string `json:"SourceUser" mapstructure:"SourceUser"`
	TargetImage     string `json:"TargetImage" mapstructure:"TargetImage"`
	TargetProcessId int64  `json:"TargetProcessId" mapstructure:"TargetProcessId"`
	TargetUser      string `json:"TargetUser" mapstructure:"TargetUser"`
}

type ProcessTerminatedMetadata struct {
	ProcessId         int64  `json:"ProcessId" mapstructure:"ProcessId"`
	Image             string `json:"Image" mapstructure:"Image"`
	User              string `json:"User" mapstructure:"User"`
}

type NetworkConnectionMetadata struct {
	DestinationHostname string `json:"DestinationHostname" mapstructure:"DestinationHostname"`
	DestinationIp       string `json:"DestinationIp" mapstructure:"DestinationIp"`
	DestinationIsIpv6   bool   `json:"DestinationIsIpv6" mapstructure:"DestinationIsIpv6"`
	DestinationPort     int64  `json:"DestinationPort" mapstructure:"DestinationPort"`
	Image               string `json:"Image" mapstructure:"Image"`
	Initiated           bool   `json:"Initiated" mapstructure:"Initiated"`
	ProcessGuid         string `json:"ProcessGuid" mapstructure:"ProcessGuid"`
	ProcessId           int64  `json:"ProcessId" mapstructure:"ProcessId"`
	Protocol            string `json:"Protocol" mapstructure:"Protocol"`
	SourceHostname      string `json:"SourceHostname" mapstructure:"SourceHostname"`
	SourceIp            string `json:"SourceIp" mapstructure:"SourceIp"`
	SourceIsIpv6        bool   `json:"SourceIsIpv6" mapstructure:"SourceIsIpv6"`
	SourcePort          int64  `json:"SourcePort" mapstructure:"SourcePort"`
	User                string `json:"User" mapstructure:"User"`
	ParentImage         string `json:"ParentImage" mapstructure:"ParentImage"`
}

type DnsQueryMetadata struct {
	Image     string `json:"Image" mapstructure:"Image"`
	ProcessId int64  `json:"ProcessId" mapstructure:"ProcessId"`
	QueryName string `json:"QueryName" mapstructure:"QueryName"`
	User      string `json:"User" mapstructure:"User"`
}

type DriverLoadMetadata struct {
	Hashes    string `json:"Hashes" mapstructure:"Hashes"`
	Signature string `json:"Signature" mapstructure:"Signature"`
}

type ImageLoadMetadata struct {
	Company          string `json:"Company" mapstructure:"Company"`
	FileVersion      string `json:"FileVersion" mapstructure:"FileVersion"`
	Hashes           string `json:"Hashes" mapstructure:"Hashes"`
	Image            string `json:"Image" mapstructure:"Image"`
	OriginalFileName string `json:"OriginalFileName" mapstructure:"OriginalFileName"`
	ProcessId        int64  `json:"ProcessId" mapstructure:"ProcessId"`
	Product          string `json:"Product" mapstructure:"Product"`
	User             string `json:"User" mapstructure:"User"`
}

type FileEventMetadata struct {
	ProcessId       int64  `json:"ProcessId" mapstructure:"ProcessId"`
	Image           string `json:"Image" mapstructure:"Image"`
	TargetFilename  string `json:"TargetFilename" mapstructure:"TargetFilename"`
	CreationUtcTime string `json:"CreationUtcTime" mapstructure:"CreationUtcTime"`
	User            string `json:"User" mapstructure:"User"`
}

type FileDeleteMetadata struct {
	Image          string `json:"Image" mapstructure:"Image"`
	ProcessId      int64  `json:"ProcessId" mapstructure:"ProcessId"`
	TargetFilename string `json:"TargetFilename" mapstructure:"TargetFilename"`
	User           string `json:"User" mapstructure:"User"`
}

type RegistryAddMetadata struct {
	ProcessId    int64  `json:"ProcessId" mapstructure:"ProcessId"`
	Image        string `json:"Image" mapstructure:"Image"`
	TargetObject string `json:"TargetObject" mapstructure:"TargetObject"`
	User         string `json:"User" mapstructure:"User"`
}

type RegistryDeleteMetadata struct {
	Details      string `json:"Details" mapstructure:"Details"`
	Image        string `json:"Image" mapstructure:"Image"`
	ProcessId    int64  `json:"ProcessId" mapstructure:"ProcessId"`
	TargetObject string `json:"TargetObject" mapstructure:"TargetObject"`
}

type RegistrySetMetadata struct {
	Details      string `json:"Details" mapstructure:"Details"`
	Image        string `json:"Image" mapstructure:"Image"`
	ProcessId    int64  `json:"ProcessId" mapstructure:"ProcessId"`
	TargetObject string `json:"TargetObject" mapstructure:"TargetObject"`
	User         string `json:"User" mapstructure:"User"`
}

type RegistryEventMetadata struct {
	Details      string `json:"Details" mapstructure:"Details"`
	Image        string `json:"Image" mapstructure:"Image"`
	ProcessId    int64  `json:"ProcessId" mapstructure:"ProcessId"`
	TargetObject string `json:"TargetObject" mapstructure:"TargetObject"`
	User         string `json:"User" mapstructure:"User"`
}

type CreateStreamHashMetadata struct {
	CreationUtcTime string `json:"CreationUtcTime" mapstructure:"CreationUtcTime"`
	Hash            string `json:"Hash" mapstructure:"Hash"`
	Image           string `json:"Image" mapstructure:"Image"`
	ProcessId       int64  `json:"ProcessId" mapstructure:"ProcessId"`
	TargetFilename  string `json:"TargetFilename" mapstructure:"TargetFilename"`
	User            string `json:"User" mapstructure:"User"`
}



// DecodeMetadataAs decodes the map into a Metadata.
// It uses mapstructure to decode the Metadata field into the appropriate structure.
// Warning: This function does not return an error when attempting to decode with the wrong type due to limitations in mapstructure.
func DecodeMetadataAs[T Metadata](origin map[string]any, dest *T) (err error) {
	err = mapstructure.Decode(origin, dest)
	return err
}

// --------------------------------------------------
// System events Metadata
//
// They define the event structures for each type of event.
// --------------------------------------------------

type ProcessCreationEvent struct {
	commonModel.CommonHeader
	Metadata ProcessCreationMetadata `json:"Metadata"`
}

type ProcessTerminatedEvent struct {
	commonModel.CommonHeader
	Metadata ProcessTerminatedMetadata `json:"Metadata"`
}

type NetworkConnectionEvent struct {
	commonModel.CommonHeader
	Metadata NetworkConnectionMetadata `json:"Metadata"`
}

type DriverLoadEvent struct {
	commonModel.CommonHeader
	Metadata DriverLoadMetadata `json:"Metadata"`
}

type ImageLoadEvent struct {
	commonModel.CommonHeader
	Metadata ImageLoadMetadata `json:"Metadata"`
}

type ProcessAccessEvent struct {
	commonModel.CommonHeader
	Metadata ProcessAccessMetadata `json:"Metadata"`
}

type FileEvent struct {
	commonModel.CommonHeader
	Metadata FileEventMetadata `json:"Metadata"`
}

type RegistryAddEvent struct {
	commonModel.CommonHeader
	Metadata RegistryAddMetadata `json:"Metadata"`
}

type RegistryDeleteEvent struct {
	commonModel.CommonHeader
	Metadata RegistryDeleteMetadata `json:"Metadata"`
}

type RegistrySetEvent struct {
	commonModel.CommonHeader
	Metadata RegistrySetMetadata `json:"Metadata"`
}

type RegistryEvent struct {
	commonModel.CommonHeader
	Metadata RegistryEventMetadata `json:"Metadata"`
}

type CreateStreamHashEvent struct {
	commonModel.CommonHeader
	Metadata CreateStreamHashMetadata `json:"Metadata"`
}

type DnsQueryEvent struct {
	commonModel.CommonHeader
	Metadata DnsQueryMetadata `json:"Metadata"`
}

type FileDeleteEvent struct {
	commonModel.CommonHeader
	Metadata FileDeleteMetadata `json:"Metadata"`
}

