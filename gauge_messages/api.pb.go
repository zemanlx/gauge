// Copyright 2015 ThoughtWorks, Inc.

// This file is part of Gauge.

// Gauge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Gauge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Gauge.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package gauge_messages is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	GetProjectRootRequest
	GetProjectRootResponse
	GetInstallationRootRequest
	GetInstallationRootResponse
	GetAllStepsRequest
	GetAllStepsResponse
	GetAllSpecsRequest
	GetAllSpecsResponse
	GetAllConceptsRequest
	GetAllConceptsResponse
	ConceptInfo
	GetStepValueRequest
	GetStepValueResponse
	GetLanguagePluginLibPathRequest
	GetLanguagePluginLibPathResponse
	ErrorResponse
	PerformRefactoringRequest
	PerformRefactoringResponse
	GetExtractConceptInfoRequest
	GetExtractConceptInfoResponse
	GetFormatConceptHeadingRequest
	GetFormatConceptHeadingResponse
	APIMessage
*/
package gauge_messages

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type APIMessage_APIMessageType int32

const (
	APIMessage_GetProjectRootRequest            APIMessage_APIMessageType = 1
	APIMessage_GetProjectRootResponse           APIMessage_APIMessageType = 2
	APIMessage_GetInstallationRootRequest       APIMessage_APIMessageType = 3
	APIMessage_GetInstallationRootResponse      APIMessage_APIMessageType = 4
	APIMessage_GetAllStepsRequest               APIMessage_APIMessageType = 5
	APIMessage_GetAllStepResponse               APIMessage_APIMessageType = 6
	APIMessage_GetAllSpecsRequest               APIMessage_APIMessageType = 7
	APIMessage_GetAllSpecsResponse              APIMessage_APIMessageType = 8
	APIMessage_GetStepValueRequest              APIMessage_APIMessageType = 9
	APIMessage_GetStepValueResponse             APIMessage_APIMessageType = 10
	APIMessage_GetLanguagePluginLibPathRequest  APIMessage_APIMessageType = 11
	APIMessage_GetLanguagePluginLibPathResponse APIMessage_APIMessageType = 12
	APIMessage_ErrorResponse                    APIMessage_APIMessageType = 13
	APIMessage_GetAllConceptsRequest            APIMessage_APIMessageType = 14
	APIMessage_GetAllConceptsResponse           APIMessage_APIMessageType = 15
	APIMessage_PerformRefactoringRequest        APIMessage_APIMessageType = 16
	APIMessage_PerformRefactoringResponse       APIMessage_APIMessageType = 17
	APIMessage_GetExtractConceptInfoRequest     APIMessage_APIMessageType = 18
	APIMessage_GetExtractConceptInfoResponse    APIMessage_APIMessageType = 19
	APIMessage_GetFormatConceptHeadingRequest   APIMessage_APIMessageType = 20
	APIMessage_GetFormatConceptHeadingResponse  APIMessage_APIMessageType = 21
)

var APIMessage_APIMessageType_name = map[int32]string{
	1:  "GetProjectRootRequest",
	2:  "GetProjectRootResponse",
	3:  "GetInstallationRootRequest",
	4:  "GetInstallationRootResponse",
	5:  "GetAllStepsRequest",
	6:  "GetAllStepResponse",
	7:  "GetAllSpecsRequest",
	8:  "GetAllSpecsResponse",
	9:  "GetStepValueRequest",
	10: "GetStepValueResponse",
	11: "GetLanguagePluginLibPathRequest",
	12: "GetLanguagePluginLibPathResponse",
	13: "ErrorResponse",
	14: "GetAllConceptsRequest",
	15: "GetAllConceptsResponse",
	16: "PerformRefactoringRequest",
	17: "PerformRefactoringResponse",
	18: "GetExtractConceptInfoRequest",
	19: "GetExtractConceptInfoResponse",
	20: "GetFormatConceptHeadingRequest",
	21: "GetFormatConceptHeadingResponse",
}
var APIMessage_APIMessageType_value = map[string]int32{
	"GetProjectRootRequest":            1,
	"GetProjectRootResponse":           2,
	"GetInstallationRootRequest":       3,
	"GetInstallationRootResponse":      4,
	"GetAllStepsRequest":               5,
	"GetAllStepResponse":               6,
	"GetAllSpecsRequest":               7,
	"GetAllSpecsResponse":              8,
	"GetStepValueRequest":              9,
	"GetStepValueResponse":             10,
	"GetLanguagePluginLibPathRequest":  11,
	"GetLanguagePluginLibPathResponse": 12,
	"ErrorResponse":                    13,
	"GetAllConceptsRequest":            14,
	"GetAllConceptsResponse":           15,
	"PerformRefactoringRequest":        16,
	"PerformRefactoringResponse":       17,
	"GetExtractConceptInfoRequest":     18,
	"GetExtractConceptInfoResponse":    19,
	"GetFormatConceptHeadingRequest":   20,
	"GetFormatConceptHeadingResponse":  21,
}

func (x APIMessage_APIMessageType) Enum() *APIMessage_APIMessageType {
	p := new(APIMessage_APIMessageType)
	*p = x
	return p
}
func (x APIMessage_APIMessageType) String() string {
	return proto.EnumName(APIMessage_APIMessageType_name, int32(x))
}
func (x *APIMessage_APIMessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(APIMessage_APIMessageType_value, data, "APIMessage_APIMessageType")
	if err != nil {
		return err
	}
	*x = APIMessage_APIMessageType(value)
	return nil
}

type GetProjectRootRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetProjectRootRequest) Reset()         { *m = GetProjectRootRequest{} }
func (m *GetProjectRootRequest) String() string { return proto.CompactTextString(m) }
func (*GetProjectRootRequest) ProtoMessage()    {}

type GetProjectRootResponse struct {
	ProjectRoot      *string `protobuf:"bytes,1,req,name=projectRoot" json:"projectRoot,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetProjectRootResponse) Reset()         { *m = GetProjectRootResponse{} }
func (m *GetProjectRootResponse) String() string { return proto.CompactTextString(m) }
func (*GetProjectRootResponse) ProtoMessage()    {}

func (m *GetProjectRootResponse) GetProjectRoot() string {
	if m != nil && m.ProjectRoot != nil {
		return *m.ProjectRoot
	}
	return ""
}

type GetInstallationRootRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetInstallationRootRequest) Reset()         { *m = GetInstallationRootRequest{} }
func (m *GetInstallationRootRequest) String() string { return proto.CompactTextString(m) }
func (*GetInstallationRootRequest) ProtoMessage()    {}

type GetInstallationRootResponse struct {
	InstallationRoot *string `protobuf:"bytes,1,req,name=installationRoot" json:"installationRoot,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetInstallationRootResponse) Reset()         { *m = GetInstallationRootResponse{} }
func (m *GetInstallationRootResponse) String() string { return proto.CompactTextString(m) }
func (*GetInstallationRootResponse) ProtoMessage()    {}

func (m *GetInstallationRootResponse) GetInstallationRoot() string {
	if m != nil && m.InstallationRoot != nil {
		return *m.InstallationRoot
	}
	return ""
}

type GetAllStepsRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetAllStepsRequest) Reset()         { *m = GetAllStepsRequest{} }
func (m *GetAllStepsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllStepsRequest) ProtoMessage()    {}

type GetAllStepsResponse struct {
	AllSteps         []*ProtoStepValue `protobuf:"bytes,1,rep,name=allSteps" json:"allSteps,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *GetAllStepsResponse) Reset()         { *m = GetAllStepsResponse{} }
func (m *GetAllStepsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllStepsResponse) ProtoMessage()    {}

func (m *GetAllStepsResponse) GetAllSteps() []*ProtoStepValue {
	if m != nil {
		return m.AllSteps
	}
	return nil
}

type GetAllSpecsRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetAllSpecsRequest) Reset()         { *m = GetAllSpecsRequest{} }
func (m *GetAllSpecsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllSpecsRequest) ProtoMessage()    {}

type GetAllSpecsResponse struct {
	Specs            []*ProtoSpec `protobuf:"bytes,1,rep,name=specs" json:"specs,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *GetAllSpecsResponse) Reset()         { *m = GetAllSpecsResponse{} }
func (m *GetAllSpecsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllSpecsResponse) ProtoMessage()    {}

func (m *GetAllSpecsResponse) GetSpecs() []*ProtoSpec {
	if m != nil {
		return m.Specs
	}
	return nil
}

type GetAllConceptsRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetAllConceptsRequest) Reset()         { *m = GetAllConceptsRequest{} }
func (m *GetAllConceptsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllConceptsRequest) ProtoMessage()    {}

type GetAllConceptsResponse struct {
	Concepts         []*ConceptInfo `protobuf:"bytes,1,rep,name=concepts" json:"concepts,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *GetAllConceptsResponse) Reset()         { *m = GetAllConceptsResponse{} }
func (m *GetAllConceptsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllConceptsResponse) ProtoMessage()    {}

func (m *GetAllConceptsResponse) GetConcepts() []*ConceptInfo {
	if m != nil {
		return m.Concepts
	}
	return nil
}

type ConceptInfo struct {
	StepValue        *ProtoStepValue `protobuf:"bytes,1,req,name=stepValue" json:"stepValue,omitempty"`
	Filepath         *string         `protobuf:"bytes,2,req,name=filepath" json:"filepath,omitempty"`
	LineNumber       *int32          `protobuf:"varint,3,req,name=lineNumber" json:"lineNumber,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *ConceptInfo) Reset()         { *m = ConceptInfo{} }
func (m *ConceptInfo) String() string { return proto.CompactTextString(m) }
func (*ConceptInfo) ProtoMessage()    {}

func (m *ConceptInfo) GetStepValue() *ProtoStepValue {
	if m != nil {
		return m.StepValue
	}
	return nil
}

func (m *ConceptInfo) GetFilepath() string {
	if m != nil && m.Filepath != nil {
		return *m.Filepath
	}
	return ""
}

func (m *ConceptInfo) GetLineNumber() int32 {
	if m != nil && m.LineNumber != nil {
		return *m.LineNumber
	}
	return 0
}

type GetStepValueRequest struct {
	StepText         *string `protobuf:"bytes,1,req,name=stepText" json:"stepText,omitempty"`
	HasInlineTable   *bool   `protobuf:"varint,2,opt,name=hasInlineTable" json:"hasInlineTable,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetStepValueRequest) Reset()         { *m = GetStepValueRequest{} }
func (m *GetStepValueRequest) String() string { return proto.CompactTextString(m) }
func (*GetStepValueRequest) ProtoMessage()    {}

func (m *GetStepValueRequest) GetStepText() string {
	if m != nil && m.StepText != nil {
		return *m.StepText
	}
	return ""
}

func (m *GetStepValueRequest) GetHasInlineTable() bool {
	if m != nil && m.HasInlineTable != nil {
		return *m.HasInlineTable
	}
	return false
}

type GetStepValueResponse struct {
	StepValue        *ProtoStepValue `protobuf:"bytes,1,req,name=stepValue" json:"stepValue,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *GetStepValueResponse) Reset()         { *m = GetStepValueResponse{} }
func (m *GetStepValueResponse) String() string { return proto.CompactTextString(m) }
func (*GetStepValueResponse) ProtoMessage()    {}

func (m *GetStepValueResponse) GetStepValue() *ProtoStepValue {
	if m != nil {
		return m.StepValue
	}
	return nil
}

type GetLanguagePluginLibPathRequest struct {
	Language         *string `protobuf:"bytes,1,req,name=language" json:"language,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetLanguagePluginLibPathRequest) Reset()         { *m = GetLanguagePluginLibPathRequest{} }
func (m *GetLanguagePluginLibPathRequest) String() string { return proto.CompactTextString(m) }
func (*GetLanguagePluginLibPathRequest) ProtoMessage()    {}

func (m *GetLanguagePluginLibPathRequest) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return ""
}

type GetLanguagePluginLibPathResponse struct {
	Path             *string `protobuf:"bytes,1,req,name=path" json:"path,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetLanguagePluginLibPathResponse) Reset()         { *m = GetLanguagePluginLibPathResponse{} }
func (m *GetLanguagePluginLibPathResponse) String() string { return proto.CompactTextString(m) }
func (*GetLanguagePluginLibPathResponse) ProtoMessage()    {}

func (m *GetLanguagePluginLibPathResponse) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

type ErrorResponse struct {
	Error            *string `protobuf:"bytes,1,req,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ErrorResponse) Reset()         { *m = ErrorResponse{} }
func (m *ErrorResponse) String() string { return proto.CompactTextString(m) }
func (*ErrorResponse) ProtoMessage()    {}

func (m *ErrorResponse) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

type PerformRefactoringRequest struct {
	OldStep          *string `protobuf:"bytes,1,req,name=oldStep" json:"oldStep,omitempty"`
	NewStep          *string `protobuf:"bytes,2,req,name=newStep" json:"newStep,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PerformRefactoringRequest) Reset()         { *m = PerformRefactoringRequest{} }
func (m *PerformRefactoringRequest) String() string { return proto.CompactTextString(m) }
func (*PerformRefactoringRequest) ProtoMessage()    {}

func (m *PerformRefactoringRequest) GetOldStep() string {
	if m != nil && m.OldStep != nil {
		return *m.OldStep
	}
	return ""
}

func (m *PerformRefactoringRequest) GetNewStep() string {
	if m != nil && m.NewStep != nil {
		return *m.NewStep
	}
	return ""
}

type PerformRefactoringResponse struct {
	Success          *bool    `protobuf:"varint,1,req,name=success" json:"success,omitempty"`
	Errors           []string `protobuf:"bytes,2,rep,name=errors" json:"errors,omitempty"`
	FilesChanged     []string `protobuf:"bytes,3,rep,name=filesChanged" json:"filesChanged,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *PerformRefactoringResponse) Reset()         { *m = PerformRefactoringResponse{} }
func (m *PerformRefactoringResponse) String() string { return proto.CompactTextString(m) }
func (*PerformRefactoringResponse) ProtoMessage()    {}

func (m *PerformRefactoringResponse) GetSuccess() bool {
	if m != nil && m.Success != nil {
		return *m.Success
	}
	return false
}

func (m *PerformRefactoringResponse) GetErrors() []string {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *PerformRefactoringResponse) GetFilesChanged() []string {
	if m != nil {
		return m.FilesChanged
	}
	return nil
}

type GetExtractConceptInfoRequest struct {
	Text             *string `protobuf:"bytes,1,req,name=text" json:"text,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetExtractConceptInfoRequest) Reset()         { *m = GetExtractConceptInfoRequest{} }
func (m *GetExtractConceptInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetExtractConceptInfoRequest) ProtoMessage()    {}

func (m *GetExtractConceptInfoRequest) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

type GetExtractConceptInfoResponse struct {
	IsValid          *bool   `protobuf:"varint,1,req,name=isValid" json:"isValid,omitempty"`
	HasParam         *bool   `protobuf:"varint,2,opt,name=hasParam" json:"hasParam,omitempty"`
	ConceptHeading   *string `protobuf:"bytes,3,opt,name=conceptHeading" json:"conceptHeading,omitempty"`
	Steps            *string `protobuf:"bytes,4,opt,name=steps" json:"steps,omitempty"`
	ConceptText      *string `protobuf:"bytes,5,opt,name=conceptText" json:"conceptText,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetExtractConceptInfoResponse) Reset()         { *m = GetExtractConceptInfoResponse{} }
func (m *GetExtractConceptInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetExtractConceptInfoResponse) ProtoMessage()    {}

func (m *GetExtractConceptInfoResponse) GetIsValid() bool {
	if m != nil && m.IsValid != nil {
		return *m.IsValid
	}
	return false
}

func (m *GetExtractConceptInfoResponse) GetHasParam() bool {
	if m != nil && m.HasParam != nil {
		return *m.HasParam
	}
	return false
}

func (m *GetExtractConceptInfoResponse) GetConceptHeading() string {
	if m != nil && m.ConceptHeading != nil {
		return *m.ConceptHeading
	}
	return ""
}

func (m *GetExtractConceptInfoResponse) GetSteps() string {
	if m != nil && m.Steps != nil {
		return *m.Steps
	}
	return ""
}

func (m *GetExtractConceptInfoResponse) GetConceptText() string {
	if m != nil && m.ConceptText != nil {
		return *m.ConceptText
	}
	return ""
}

type GetFormatConceptHeadingRequest struct {
	NewConceptHeading *string `protobuf:"bytes,1,req,name=newConceptHeading" json:"newConceptHeading,omitempty"`
	OldConceptHeading *string `protobuf:"bytes,2,req,name=oldConceptHeading" json:"oldConceptHeading,omitempty"`
	OldConceptText    *string `protobuf:"bytes,3,req,name=oldConceptText" json:"oldConceptText,omitempty"`
	XXX_unrecognized  []byte  `json:"-"`
}

func (m *GetFormatConceptHeadingRequest) Reset()         { *m = GetFormatConceptHeadingRequest{} }
func (m *GetFormatConceptHeadingRequest) String() string { return proto.CompactTextString(m) }
func (*GetFormatConceptHeadingRequest) ProtoMessage()    {}

func (m *GetFormatConceptHeadingRequest) GetNewConceptHeading() string {
	if m != nil && m.NewConceptHeading != nil {
		return *m.NewConceptHeading
	}
	return ""
}

func (m *GetFormatConceptHeadingRequest) GetOldConceptHeading() string {
	if m != nil && m.OldConceptHeading != nil {
		return *m.OldConceptHeading
	}
	return ""
}

func (m *GetFormatConceptHeadingRequest) GetOldConceptText() string {
	if m != nil && m.OldConceptText != nil {
		return *m.OldConceptText
	}
	return ""
}

type GetFormatConceptHeadingResponse struct {
	NewConceptText   *string `protobuf:"bytes,1,req,name=newConceptText" json:"newConceptText,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetFormatConceptHeadingResponse) Reset()         { *m = GetFormatConceptHeadingResponse{} }
func (m *GetFormatConceptHeadingResponse) String() string { return proto.CompactTextString(m) }
func (*GetFormatConceptHeadingResponse) ProtoMessage()    {}

func (m *GetFormatConceptHeadingResponse) GetNewConceptText() string {
	if m != nil && m.NewConceptText != nil {
		return *m.NewConceptText
	}
	return ""
}

type APIMessage struct {
	MessageType                  *APIMessage_APIMessageType        `protobuf:"varint,1,req,name=messageType,enum=gauge.messages.APIMessage_APIMessageType" json:"messageType,omitempty"`
	MessageId                    *int64                            `protobuf:"varint,2,req,name=messageId" json:"messageId,omitempty"`
	ProjectRootRequest           *GetProjectRootRequest            `protobuf:"bytes,3,opt,name=projectRootRequest" json:"projectRootRequest,omitempty"`
	ProjectRootResponse          *GetProjectRootResponse           `protobuf:"bytes,4,opt,name=projectRootResponse" json:"projectRootResponse,omitempty"`
	InstallationRootRequest      *GetInstallationRootRequest       `protobuf:"bytes,5,opt,name=installationRootRequest" json:"installationRootRequest,omitempty"`
	InstallationRootResponse     *GetInstallationRootResponse      `protobuf:"bytes,6,opt,name=installationRootResponse" json:"installationRootResponse,omitempty"`
	AllStepsRequest              *GetAllStepsRequest               `protobuf:"bytes,7,opt,name=allStepsRequest" json:"allStepsRequest,omitempty"`
	AllStepsResponse             *GetAllStepsResponse              `protobuf:"bytes,8,opt,name=allStepsResponse" json:"allStepsResponse,omitempty"`
	AllSpecsRequest              *GetAllSpecsRequest               `protobuf:"bytes,9,opt,name=allSpecsRequest" json:"allSpecsRequest,omitempty"`
	AllSpecsResponse             *GetAllSpecsResponse              `protobuf:"bytes,10,opt,name=allSpecsResponse" json:"allSpecsResponse,omitempty"`
	StepValueRequest             *GetStepValueRequest              `protobuf:"bytes,11,opt,name=stepValueRequest" json:"stepValueRequest,omitempty"`
	StepValueResponse            *GetStepValueResponse             `protobuf:"bytes,12,opt,name=stepValueResponse" json:"stepValueResponse,omitempty"`
	LibPathRequest               *GetLanguagePluginLibPathRequest  `protobuf:"bytes,13,opt,name=libPathRequest" json:"libPathRequest,omitempty"`
	LibPathResponse              *GetLanguagePluginLibPathResponse `protobuf:"bytes,14,opt,name=libPathResponse" json:"libPathResponse,omitempty"`
	Error                        *ErrorResponse                    `protobuf:"bytes,15,opt,name=error" json:"error,omitempty"`
	AllConceptsRequest           *GetAllConceptsRequest            `protobuf:"bytes,16,opt,name=allConceptsRequest" json:"allConceptsRequest,omitempty"`
	AllConceptsResponse          *GetAllConceptsResponse           `protobuf:"bytes,17,opt,name=allConceptsResponse" json:"allConceptsResponse,omitempty"`
	PerformRefactoringRequest    *PerformRefactoringRequest        `protobuf:"bytes,18,opt,name=performRefactoringRequest" json:"performRefactoringRequest,omitempty"`
	PerformRefactoringResponse   *PerformRefactoringResponse       `protobuf:"bytes,19,opt,name=performRefactoringResponse" json:"performRefactoringResponse,omitempty"`
	ExtractConceptInfoRequest    *GetExtractConceptInfoRequest     `protobuf:"bytes,20,opt,name=extractConceptInfoRequest" json:"extractConceptInfoRequest,omitempty"`
	ExtractConceptInfoResponse   *GetExtractConceptInfoResponse    `protobuf:"bytes,21,opt,name=extractConceptInfoResponse" json:"extractConceptInfoResponse,omitempty"`
	FormatConceptHeadingRequest  *GetFormatConceptHeadingRequest   `protobuf:"bytes,22,opt,name=formatConceptHeadingRequest" json:"formatConceptHeadingRequest,omitempty"`
	FormatConceptHeadingResponse *GetFormatConceptHeadingResponse  `protobuf:"bytes,23,opt,name=formatConceptHeadingResponse" json:"formatConceptHeadingResponse,omitempty"`
	XXX_unrecognized             []byte                            `json:"-"`
}

func (m *APIMessage) Reset()         { *m = APIMessage{} }
func (m *APIMessage) String() string { return proto.CompactTextString(m) }
func (*APIMessage) ProtoMessage()    {}

func (m *APIMessage) GetMessageType() APIMessage_APIMessageType {
	if m != nil && m.MessageType != nil {
		return *m.MessageType
	}
	return APIMessage_GetProjectRootRequest
}

func (m *APIMessage) GetMessageId() int64 {
	if m != nil && m.MessageId != nil {
		return *m.MessageId
	}
	return 0
}

func (m *APIMessage) GetProjectRootRequest() *GetProjectRootRequest {
	if m != nil {
		return m.ProjectRootRequest
	}
	return nil
}

func (m *APIMessage) GetProjectRootResponse() *GetProjectRootResponse {
	if m != nil {
		return m.ProjectRootResponse
	}
	return nil
}

func (m *APIMessage) GetInstallationRootRequest() *GetInstallationRootRequest {
	if m != nil {
		return m.InstallationRootRequest
	}
	return nil
}

func (m *APIMessage) GetInstallationRootResponse() *GetInstallationRootResponse {
	if m != nil {
		return m.InstallationRootResponse
	}
	return nil
}

func (m *APIMessage) GetAllStepsRequest() *GetAllStepsRequest {
	if m != nil {
		return m.AllStepsRequest
	}
	return nil
}

func (m *APIMessage) GetAllStepsResponse() *GetAllStepsResponse {
	if m != nil {
		return m.AllStepsResponse
	}
	return nil
}

func (m *APIMessage) GetAllSpecsRequest() *GetAllSpecsRequest {
	if m != nil {
		return m.AllSpecsRequest
	}
	return nil
}

func (m *APIMessage) GetAllSpecsResponse() *GetAllSpecsResponse {
	if m != nil {
		return m.AllSpecsResponse
	}
	return nil
}

func (m *APIMessage) GetStepValueRequest() *GetStepValueRequest {
	if m != nil {
		return m.StepValueRequest
	}
	return nil
}

func (m *APIMessage) GetStepValueResponse() *GetStepValueResponse {
	if m != nil {
		return m.StepValueResponse
	}
	return nil
}

func (m *APIMessage) GetLibPathRequest() *GetLanguagePluginLibPathRequest {
	if m != nil {
		return m.LibPathRequest
	}
	return nil
}

func (m *APIMessage) GetLibPathResponse() *GetLanguagePluginLibPathResponse {
	if m != nil {
		return m.LibPathResponse
	}
	return nil
}

func (m *APIMessage) GetError() *ErrorResponse {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *APIMessage) GetAllConceptsRequest() *GetAllConceptsRequest {
	if m != nil {
		return m.AllConceptsRequest
	}
	return nil
}

func (m *APIMessage) GetAllConceptsResponse() *GetAllConceptsResponse {
	if m != nil {
		return m.AllConceptsResponse
	}
	return nil
}

func (m *APIMessage) GetPerformRefactoringRequest() *PerformRefactoringRequest {
	if m != nil {
		return m.PerformRefactoringRequest
	}
	return nil
}

func (m *APIMessage) GetPerformRefactoringResponse() *PerformRefactoringResponse {
	if m != nil {
		return m.PerformRefactoringResponse
	}
	return nil
}

func (m *APIMessage) GetExtractConceptInfoRequest() *GetExtractConceptInfoRequest {
	if m != nil {
		return m.ExtractConceptInfoRequest
	}
	return nil
}

func (m *APIMessage) GetExtractConceptInfoResponse() *GetExtractConceptInfoResponse {
	if m != nil {
		return m.ExtractConceptInfoResponse
	}
	return nil
}

func (m *APIMessage) GetFormatConceptHeadingRequest() *GetFormatConceptHeadingRequest {
	if m != nil {
		return m.FormatConceptHeadingRequest
	}
	return nil
}

func (m *APIMessage) GetFormatConceptHeadingResponse() *GetFormatConceptHeadingResponse {
	if m != nil {
		return m.FormatConceptHeadingResponse
	}
	return nil
}

func init() {
	proto.RegisterEnum("gauge.messages.APIMessage_APIMessageType", APIMessage_APIMessageType_name, APIMessage_APIMessageType_value)
}
