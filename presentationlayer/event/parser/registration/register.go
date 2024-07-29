// Package registration parses the registration viewmodel.
//
// It's based on the following Lingo code (source: Habbo Origins)
//
//	on construct me
//	pValidPartProps = [:]
//	pValidPartGroups = [:]
//	pFigurePartListLoadedFlag = 0
//	pAvailableSetListLoadedFlag = 0
//	pState = 0
//	pAgeCheckFlag = VOID
//	pParentEmailNeededFlag = VOID
//	pParentEmailAddress = EMPTY
//	pRegMsgStruct = [:]
//	pRegMsgStruct["parentagree"] = [#id: 1, "type": #boolean]
//	pRegMsgStruct["name"] = [#id: 2, "type": #string]
//	pRegMsgStruct["password"] = [#id: 3, "type": #string]
//	pRegMsgStruct["figure"] = [#id: 4, "type": #string]
//	pRegMsgStruct["sex"] = [#id: 5, "type": #string]
//	pRegMsgStruct["customData"] = [#id: 6, "type": #string]
//	pRegMsgStruct["email"] = [#id: 7, "type": #string]
//	pRegMsgStruct["birthday"] = [#id: 8, "type": #string]
//	pRegMsgStruct["directMail"] = [#id: 9, "type": #boolean]
//	pRegMsgStruct["has_read_agreement"] = [#id: 10, "type": #boolean]
//	pRegMsgStruct["isp_id"] = [#id: 11, "type": #string]
//	pRegMsgStruct["partnersite"] = [#id: 12, "type": #string]
//	pRegMsgStruct["oldpassword"] = [#id: 13, "type": #string]
//	registerMessage(#enterRoom, me.getID(), #closeFigureCreator)
//	registerMessage(#changeRoom, me.getID(), #closeFigureCreator)
//	registerMessage(#leaveRoom, me.getID(), #closeFigureCreator)
//	registerMessage(#show_registration, me.getID(), #openFigureCreator)
//	registerMessage(#hide_registration, me.getID(), #closeFigureCreator)
//	registerMessage(#figure_ready, me.getID(), #figureSystemReady)
//	end
//
// There's a pattern here. Each struct item contains a map with an id and type.
//
// The type gives us the ability to read the packet accordingly (e.g. reader.String() for #string; reader.Bool()
// (or Byte()) for #boolean.
//
// #id is the header. This is a short - a base64-encoded integer value.
package registration

import (
	"hc/pkg/packet"
)

type (
	RegisterID   int
	RegisterType string
)

var registerMessages = map[RegisterID]RegisterType{
	ParentAgree:      "boolean",
	Name:             "string",
	Password:         "string",
	Figure:           "string",
	Sex:              "string",
	CustomData:       "string",
	Email:            "string",
	Birthday:         "string",
	DirectMail:       "boolean",
	HasReadAgreement: "boolean",
	IspID:            "string",
	PartnerSite:      "string",
	OldPassword:      "string",
}

const (
	ParentAgree = iota + 1
	Name
	Password
	Figure
	Sex
	CustomData
	Email
	Birthday
	DirectMail
	HasReadAgreement
	IspID
	PartnerSite
	OldPassword
)

type Register struct {
	Username   string
	Password   string
	Email      string
	Figure     string
	Sex        string
	CustomData string
}

func ParseRegister(body []byte) (register Register, err error) {
	packetReader := packet.AcquireReader(body)
	defer packet.ReleaseReader(packetReader)

	messages := make(map[RegisterID]any)

	for packetReader.Buffer.Len() > 0 {
		id, err := packetReader.Short()
		if err != nil {
			return Register{}, err
		}

		registerMessage, ok := registerMessages[RegisterID(id)]
		if !ok {
			continue
		}

		if registerMessage == "string" {
			str, err := packetReader.String()
			if err != nil {
				return Register{}, err
			}

			messages[RegisterID(id)] = str
		} else {
			b, err := packetReader.Buffer.ReadByte()
			if err != nil {
				return Register{}, err
			}

			messages[RegisterID(id)] = string(b) == "A"
		}
	}

	username, ok := messages[Name]
	if ok {
		register.Username = username.(string)
	}

	password, ok := messages[Password]
	if ok {
		register.Password = password.(string)
	}

	email, ok := messages[Email]
	if ok {
		register.Email = email.(string)
	}

	figure, ok := messages[Figure]
	if ok {
		register.Figure = figure.(string)
	}

	sex, ok := messages[Sex]
	if ok {
		register.Sex = sex.(string)
	}

	customData, ok := messages[CustomData]
	if ok {
		register.CustomData = customData.(string)
	}

	return
}
