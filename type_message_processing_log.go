package main

import (
	"encoding/json"
	"io/ioutil"
)

func ParseMessageProcessingLogFromFile(p string) *MessageProcessingLog {

	dat, _ := ioutil.ReadFile(p)

	return ParseMessageProcessingLogFromString(dat)

}

// ParseMessageProcessingLogFromString
func ParseMessageProcessingLogFromString(s []byte) *MessageProcessingLog {
	rt := &MessageProcessingLog{}
	json.Unmarshal(s, rt)
	return rt
}

type MessageProcessingLog struct {
	D D `json:"d"`
}

type D struct {
	Count   *string  `json:"__count,omitempty"`
	Results []Result `json:"results"`
}

type Result struct {
	Metadata               *ResultMetadata        `json:"__metadata,omitempty"`
	MessageGUID            *string                `json:"MessageGuid,omitempty"`
	CorrelationID          *string                `json:"CorrelationId,omitempty"`
	ApplicationMessageID   interface{}            `json:"ApplicationMessageId"`
	ApplicationMessageType interface{}            `json:"ApplicationMessageType"`
	LogStart               *string                `json:"LogStart,omitempty"`
	LogEnd                 *string                `json:"LogEnd,omitempty"`
	Sender                 interface{}            `json:"Sender"`
	Receiver               interface{}            `json:"Receiver"`
	IntegrationFlowName    *IntegrationFlowName   `json:"IntegrationFlowName,omitempty"`
	Status                 *Status                `json:"Status,omitempty"`
	AlternateWebLink       *string                `json:"AlternateWebLink,omitempty"`
	IntegrationArtifact    *IntegrationArtifact   `json:"IntegrationArtifact,omitempty"`
	LogLevel               *LogLevel              `json:"LogLevel,omitempty"`
	CustomStatus           *Status                `json:"CustomStatus,omitempty"`
	TransactionID          *string                `json:"TransactionId,omitempty"`
	PreviousComponentName  *PreviousComponentName `json:"PreviousComponentName,omitempty"`
	CustomHeaderProperties *AdapterAttributes     `json:"CustomHeaderProperties,omitempty"`
	MessageStoreEntries    *AdapterAttributes     `json:"MessageStoreEntries,omitempty"`
	ErrorInformation       *AdapterAttributes     `json:"ErrorInformation,omitempty"`
	AdapterAttributes      *AdapterAttributes     `json:"AdapterAttributes,omitempty"`
	Attachments            *AdapterAttributes     `json:"Attachments,omitempty"`
	Runs                   *AdapterAttributes     `json:"Runs,omitempty"`
}
type AdapterAttributes struct {
	Deferred Deferred `json:"__deferred"`
}

type Deferred struct {
	URI string `json:"uri"`
}

type IntegrationArtifact struct {
	Metadata IntegrationArtifactMetadata `json:"__metadata"`
	ID       IntegrationFlowName         `json:"Id"`
	Name     Name                        `json:"Name"`
	Type     Type                        `json:"Type"`
}

type IntegrationArtifactMetadata struct {
	Type PurpleType `json:"type"`
}

type ResultMetadata struct {
	ID   string     `json:"id"`
	URI  string     `json:"uri"`
	Type FluffyType `json:"type"`
}

type Status string

const (
	Failed Status = "FAILED"
)

type IntegrationFlowName string

const (
	C4CtoPropertyManageSystemCustomerObject                       IntegrationFlowName = "C4CtoPropertyManageSystem_CustomerObject"
	IntegrationFlowNameC4CtoPropertyManageSystemBuilding          IntegrationFlowName = "C4CtoPropertyManageSystem_Building"
	IntegrationFlowNameC4CtoPropertyManageSystemIndividualProduct IntegrationFlowName = "C4CtoPropertyManageSystem_individualProduct"
	IntegrationFlowNameCustomerUpdate                             IntegrationFlowName = "CustomerUpdate"
	IntegrationFlowNameQRoomNotRTBuilding                         IntegrationFlowName = "Q_RoomNotRTBuilding"
	TestSurveyResultCreate                                        IntegrationFlowName = "Test_Survey_Result_Create"
)

type PurpleType string

const (
	COMSapHCIAPIIntegrationArtifact PurpleType = "com.sap.hci.api.IntegrationArtifact"
)

type Name string

const (
	CustomerObjectUpdateAndQuery                   Name = "CustomerObjectUpdateAndQuery"
	NameC4CtoPropertyManageSystemBuilding          Name = "C4CtoPropertyManageSystem_Building"
	NameC4CtoPropertyManageSystemIndividualProduct Name = "C4CtoPropertyManageSystem_individualProduct"
	NameCustomerUpdate                             Name = "CustomerUpdate"
	NameQRoomNotRTBuilding                         Name = "Q_RoomNotRTBuilding"
	NameTestSurveyResultCreate                     Name = "Test Survey Result Create"
)

type Type string

const (
	IntegrationFlow Type = "INTEGRATION_FLOW"
)

type LogLevel string

const (
	Debug LogLevel = "DEBUG"
	Info  LogLevel = "INFO"
	Trace LogLevel = "TRACE"
)

type FluffyType string

const (
	COMSapHCIAPIMessageProcessingLog FluffyType = "com.sap.hci.api.MessageProcessingLog"
)

type PreviousComponentName string

const (
	CPIE600033 PreviousComponentName = "CPI_e600033"
)
