//go:build windows
// +build windows

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
	data := `{"EventID":0,"EventName":"process_creation","Source":"sysmon","TimeStamp":"2025-07-16T17:25:26.6982545+09:00","Metadata":{"CommandLine":"\"C:\\Program Files\\WindowsApps\\Microsoft.WindowsTerminal_1.22.11141.0_x64__8wekyb3d8bbwe\\WindowsTerminal.exe\" ","Company":"-","CurrentDirectory":"C:\\WINDOWS\\System32\\","FileVersion":"-","Hashes":"SHA1=3FE7F5EEBFA945B993760235BA4A1CD9DEDD5319,MD5=08480D610277845A42200424044B910F,SHA256=9FB7D27FF43617EC9BADAD2FDABFEB1DD76C2EBA31D990B81B6CEF4DC8CCA4B8,IMPHASH=CE008010C3B058594523090A70F63DF4","Image":"C:\\Program Files\\WindowsApps\\Microsoft.WindowsTerminal_1.22.11141.0_x64__8wekyb3d8bbwe\\WindowsTerminal.exe","IntegrityLevel":"Medium","LogonId":"140592926","OriginalFileName":"-","ParentCommandLine":"C:\\WINDOWS\\Explorer.EXE","ParentImage":"C:\\Windows\\explorer.exe","ParentProcessId":24044,"ProcessId":21964,"Product":"-","User":"DESKTOP-8S5JU8M\\user01"}}`

	// decode data into CommonModelWrapper
	cmw := &commonModel.CommonModelWrapper{}
	err := json.Unmarshal([]byte(data), cmw)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	// t.Logf("Decoded CommonModelWrapper: %+v", cmw)
	// Create a destination struct for the metadata
	dest := &eventModel.ProcessCreationEvent{
		CommonHeader: commonModel.CommonHeader{
			EventName: cmw.EventName,
			Source:    cmw.Source,
			Timestamp: cmw.Timestamp,
		},
		Metadata: eventModel.ProcessCreationMetadata{},
	}
	err = eventModel.DecodeMetadataAs[eventModel.ProcessCreationMetadata](cmw.Metadata, &dest.Metadata)
	if err != nil {
		t.Fatalf("Failed to decode metadata: %v", err)
	}
	// Check if the metadata was decoded correctly
	t.Logf("Decoded Metadata: %+v", dest.Metadata)
}

func TestDecodeMetadataWithInvalidData(t *testing.T) {
	// Example of how to test the decoding of metadata with an invalid type
	data := `{"EventName":"ProcessCreate","Source":"eBPF","Timestamp":"2025-06-09T16:54:26.270720921+09:00","Metadata":{"ProcessId":"invalid_type"}}`

	// decode data into CommonModelWrapper
	cmw := &commonModel.CommonModelWrapper{}
	err := json.Unmarshal([]byte(data), cmw)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Attempt to decode into a specific type that expects a different structure
	dest := &eventModel.ProcessTerminatedEvent{
		CommonHeader: commonModel.CommonHeader{
			EventName: cmw.EventName,
			Source:    cmw.Source,
			Timestamp: cmw.Timestamp,
		},
		Metadata: eventModel.ProcessTerminatedMetadata{},
	}
	err = eventModel.DecodeMetadataAs[eventModel.ProcessTerminatedMetadata](cmw.Metadata, &dest.Metadata)
	if err == nil {
		t.Fatal("Expected an error due to invalid type, but got none")
	}
	t.Logf("Expected error occurred: %v", err)
}

func TestDecodeMetadataWithDifferentDataType(t *testing.T) {
	// Example of how to test the decoding of metadata with a different data type
	data := `{"EventID":0,"EventName":"process_creation","Source":"sysmon","TimeStamp":"2025-07-16T17:25:26.6982545+09:00","Metadata":{"CommandLine":"\"C:\\Program Files\\WindowsApps\\Microsoft.WindowsTerminal_1.22.11141.0_x64__8wekyb3d8bbwe\\WindowsTerminal.exe\" ","Company":"-","CurrentDirectory":"C:\\WINDOWS\\System32\\","FileVersion":"-","Hashes":"SHA1=3FE7F5EEBFA945B993760235BA4A1CD9DEDD5319,MD5=08480D610277845A42200424044B910F,SHA256=9FB7D27FF43617EC9BADAD2FDABFEB1DD76C2EBA31D990B81B6CEF4DC8CCA4B8,IMPHASH=CE008010C3B058594523090A70F63DF4","Image":"C:\\Program Files\\WindowsApps\\Microsoft.WindowsTerminal_1.22.11141.0_x64__8wekyb3d8bbwe\\WindowsTerminal.exe","IntegrityLevel":"Medium","LogonId":"140592926","OriginalFileName":"-","ParentCommandLine":"C:\\WINDOWS\\Explorer.EXE","ParentImage":"C:\\Windows\\explorer.exe","ParentProcessId":24044,"ProcessId":21964,"Product":"-","User":"DESKTOP-8S5JU8M\\user01"}}`
	
	// decode data into CommonModelWrapper
	cmw := &commonModel.CommonModelWrapper{}
	err := json.Unmarshal([]byte(data), cmw)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Create a destination struct for the metadata with a different type
	dest := &eventModel.ProcessTerminatedEvent{
		CommonHeader: commonModel.CommonHeader{
			EventName: cmw.EventName,
			Source:    cmw.Source,
			Timestamp: cmw.Timestamp,
		},
		Metadata: eventModel.ProcessTerminatedMetadata{},
	}
	err = eventModel.DecodeMetadataAs[eventModel.ProcessTerminatedMetadata](cmw.Metadata, &dest.Metadata)
	if err != nil {
		t.Fatalf("Failed to decode metadata: %v", err)
	}
	t.Logf("Decoded Metadata with different type: %+v", dest.Metadata)
}
