package commonModel_test

import (
	"encoding/json"
	"testing"

	commonModel "github.com/enki-polvo/polvo-logger/model"
	eventModel "github.com/enki-polvo/polvo-logger/model/event"
)

func TestDecodeMetadata(t *testing.T) {
	// Example of how to test the decoding of metadata
	// This is a placeholder for actual test implementation
	// You would typically use a library like "encoding/json" or "github.com/mitchellh/mapstructure"
	// to decode the metadata into the appropriate struct.
	data := `{"EventName":"ProcessCreate","Source":"eBPF","Timestamp":"2025-06-09T16:54:26.270720921+09:00","Metadata":{"PID":97469,"PPID":97467,"UID":1000,"Username":"shhong","TGID":97469,"Commandline":"sh sed s/-//","ENV":"GJS_DEBUG_TOPICS=JS ERROR;JS LOG,XDG_ACTIVATION_TOKEN=gnome-shell/System Monitor/3726-8-shhong-server_TIME15324537,LANGUAGE=en,LC_TIME=ko_KR.UTF-8,USER=shhong,XDG_SESSION_TYPE=x11,SHLVL=0,CLUTTER_DISABLE_MIPMAPPED_TEXT=1,HOME=/home/shhong,DESKTOP_SESSION=ubuntu","Image":"/usr/bin/sed"}}`

	// decode data into CommonModelWrapper
	cmw := &commonModel.CommonModelWrapper{}
	err := json.Unmarshal([]byte(data), cmw)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	// t.Logf("Decoded CommonModelWrapper: %+v", cmw)
	// Create a destination struct for the metadata
	dest := &eventModel.ProcessCreateEvent{
		CommonHeader: commonModel.CommonHeader{
			EventName: cmw.EventName,
			Source:    cmw.Source,
			Timestamp: cmw.Timestamp,
		},
		Metadata: eventModel.ProcessCreateMetadata{},
	}
	err = eventModel.DecodeMetadataAs[eventModel.ProcessCreateMetadata](cmw.Metadata, &dest.Metadata)
	if err != nil {
		t.Fatalf("Failed to decode metadata: %v", err)
	}
	// Check if the metadata was decoded correctly
	t.Logf("Decoded Metadata: %+v", dest.Metadata)
}

func TestDecodeMetadataWithInvalidData(t *testing.T) {
	// Example of how to test the decoding of metadata with an invalid type
	data := `{"EventName":"ProcessCreate","Source":"eBPF","Timestamp":"2025-06-09T16:54:26.270720921+09:00","Metadata":{"PID":"invalid_type"}}`

	// decode data into CommonModelWrapper
	cmw := &commonModel.CommonModelWrapper{}
	err := json.Unmarshal([]byte(data), cmw)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Attempt to decode into a specific type that expects a different structure
	dest := &eventModel.ProcessCreateEvent{
		CommonHeader: commonModel.CommonHeader{
			EventName: cmw.EventName,
			Source:    cmw.Source,
			Timestamp: cmw.Timestamp,
		},
		Metadata: eventModel.ProcessCreateMetadata{},
	}
	err = eventModel.DecodeMetadataAs[eventModel.ProcessCreateMetadata](cmw.Metadata, &dest.Metadata)
	if err == nil {
		t.Fatal("Expected an error due to invalid type, but got none")
	}
	t.Logf("Expected error occurred: %v", err)
}

func TestDecodeMetadataWithDifferentDataType(t *testing.T) {
	// Example of how to test the decoding of metadata with a different data type
	data := `{"EventName":"ProcessCreate","Source":"eBPF","Timestamp":"2025-06-09T16:54:26.270720921+09:00","Metadata":{"PID":97469,"PPID":97467,"UID":1000,"Username":"shhong","TGID":97469,"Commandline":"sh sed s/-//","ENV":"GJS_DEBUG_TOPICS=JS ERROR;JS LOG,XDG_ACTIVATION_TOKEN=gnome-shell/System Monitor/3726-8-shhong-server_TIME15324537,LANGUAGE=en,LC_TIME=ko_KR.UTF-8,USER=shhong,XDG_SESSION_TYPE=x11,SHLVL=0,CLUTTER_DISABLE_MIPMAPPED_TEXT=1,HOME=/home/shhong,DESKTOP_SESSION=ubuntu","Image":"/usr/bin/sed"}}`

	// decode data into CommonModelWrapper
	cmw := &commonModel.CommonModelWrapper{}
	err := json.Unmarshal([]byte(data), cmw)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Create a destination struct for the metadata with a different type
	dest := &eventModel.ProcessTerminateEvent{
		CommonHeader: commonModel.CommonHeader{
			EventName: cmw.EventName,
			Source:    cmw.Source,
			Timestamp: cmw.Timestamp,
		},
		Metadata: eventModel.ProcessTerminateMetadata{},
	}
	err = eventModel.DecodeMetadataAs[eventModel.ProcessTerminateMetadata](cmw.Metadata, &dest.Metadata)
	if err != nil {
		t.Fatalf("Failed to decode metadata: %v", err)
	}
	t.Logf("Decoded Metadata with different type: %+v", dest.Metadata)
}
