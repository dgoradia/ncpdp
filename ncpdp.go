package ncpdp

import (
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"io"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gocarina/gocsv"
)

func NewMessage(message string) (*Message, error) {
	dec, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return nil, err
	}

	var msg Message
	return &msg, xml.Unmarshal(dec, &msg)
}

func (m *Message) ToJson() ([]byte, error) {
	return json.Marshal(m)
}

func baseModulePath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Dir(b)
}

type Terminologies []Terminology

type Terminology struct {
	NCItSubsetCode      string
	SubsetPreferredTerm string
	NCItCode            string
	PreferredTerm       string
	Synonym             string
	NCItPreferredTerm   string
	NCItDefinition      string
}

func LoadTerminology() (*Terminologies, error) {
	file, err := os.Open(baseModulePath() + "/NCPDPTerminology.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data Terminologies
	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.FieldsPerRecord = 7

	// Skip the first row
	if _, err := reader.Read(); err != nil {
		return nil, err
	}

	for {
		rec, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		term := Terminology{
			rec[0], rec[1], rec[2], rec[3], rec[4], rec[5], rec[6],
		}

		data = append(data, term)
	}

	return &data, nil
}

func (t Terminologies) FindTermByQuantityUnitOfMeasureCode(s string) string {
	for i := range t {
		if t[i].NCItCode == s {
			return t[i].PreferredTerm
		}
	}

	return ""
}

type LoincData []*Loinc

type Loinc struct {
	LoincNum                  string `csv:"LOINC_NUM"`
	Component                 string `csv:"COMPONENT"`
	Property                  string `csv:"PROPERTY"`
	TimeAspect                string `csv:"TIME_ASPCT"`
	System                    string `csv:"SYSTEM"`
	ScaleTyp                  string `csv:"SCALE_TYP"`
	MethodTyp                 string `csv:"METHOD_TYP"`
	Class                     string `csv:"CLASS"`
	VersionLastChanged        string `csv:"VersionLastChanged"`
	ChngType                  string `csv:"CHNG_TYPE"`
	DefinitionDescription     string `csv:"DefinitionDescription"`
	Status                    string `csv:"STATUS"`
	ConsumerName              string `csv:"CONSUMER_NAME"`
	ClassType                 string `csv:"CLASSTYPE"`
	Formula                   string `csv:"FORMULA"`
	ExmplAnswers              string `csv:"EXMPL_ANSWERS"`
	SurveryQuestText          string `csv:"SURVEY_QUEST_TEXT"`
	SurveyQuestSrc            string `csv:"SURVEY_QUEST_SRC"`
	UnitsRequired             string `csv:"UNITSREQUIRED"`
	SubmittedUnits            string `csv:"SUBMITTED_UNITS"`
	RelatedNames2             string `csv:"RELATEDNAMES2"`
	ShortName                 string `csv:"SHORTNAME"`
	OrderObs                  string `csv:"ORDER_OBS"`
	CDISCCommonTests          string `csv:"CDISC_COMMON_TESTS"`
	HL7FieldSubfieldID        string `csv:"HL7_FIELD_SUBFIELD_ID"`
	ExternalCopyrightNotice   string `csv:"EXTERNAL_COPYRIGHT_NOTICE"`
	ExampleUnits              string `csv:"EXAMPLE_UNITS"`
	LongCommonName            string `csv:"LONG_COMMON_NAME"`
	UnitsAndRange             string `csv:"UnitsAndRange"`
	ExampleUCUMUnits          string `csv:"EXAMPLE_UCUM_UNITS"`
	ExampleSiUCUMUnits        string `csv:"EXAMPLE_SI_UCUM_UNITS"`
	StatusReason              string `csv:"STATUS_REASON"`
	StatusText                string `csv:"STATUS_TEXT"`
	ChangeReasonPublic        string `csv:"CHANGE_REASON_PUBLIC"`
	CommonTestRank            string `csv:"COMMON_TEST_RANK"`
	CommonOrderRank           string `csv:"COMMON_ORDER_RANK"`
	CommonSiTestRank          string `csv:"COMMON_SI_TEST_RANK"`
	HL7AttachmentStructure    string `csv:"HL7_ATTACHMENT_STRUCTURE"`
	ExternalCopyrightLink     string `csv:"EXTERNAL_COPYRIGHT_LINK"`
	PanelType                 string `csv:"PanelType"`
	AskAtOrderEntry           string `csv:"AskAtOrderEntry"`
	AssociatedObservations    string `csv:"AssociatedObservations"`
	VersionFirstReleased      string `csv:"VersionFirstReleased"`
	ValidHL7AttachmentRequest string `csv:"ValidHL7AttachmentRequest"`
	DisplayName               string `csv:"DisplayName"`
}

func LoadLoinc() (LoincData, error) {
	file, err := os.Open(baseModulePath() + "/Loinc.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []*Loinc
	if err := gocsv.UnmarshalFile(file, &data); err != nil {
		return nil, err
	}

	return data, nil
}
