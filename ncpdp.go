package ncpdp

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"

	"github.com/gocarina/gocsv"
)

var (
	//go:embed NCPDPTerminology.txt
	NCPDPTerminologyDataFile []byte

	//go:embed Loinc.csv
	LoincDataFile []byte
)

type Decoder struct {
	msg *Message
	r   io.Reader
	buf []byte
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

func (d *Decoder) Decode() (*Message, error) {
	return d.msg, d.decode()
}

func (d *Decoder) decode() error {
	if d.buf == nil && d.r != nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(d.r)
		d.buf = buf.Bytes()
	}

	return xml.Unmarshal(d.buf, &d.msg)
}

func (d *Decoder) ToJson() ([]byte, error) {
	err := d.decode()
	if err != nil {
		return nil, err
	}

	return json.Marshal(d.msg)
}

type Terminologies []*Terminology

type Terminology struct {
	NCItSubsetCode      string
	SubsetPreferredTerm string
	NCItCode            string
	PreferredTerm       string
	Synonym             string
	NCItPreferredTerm   string
	NCItDefinition      string
}

func LoadTerminology(r io.Reader) (*Terminologies, error) {
	if r == nil {
		r = bytes.NewReader(NCPDPTerminologyDataFile)
	}

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = '\t'
		r.FieldsPerRecord = 7

		return r
	})

	var data Terminologies
	if err := gocsv.UnmarshalWithoutHeaders(r, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (t *Terminologies) Len() int {
	if t == nil {
		return 0
	}

	return len(*t)
}

func (t Terminologies) FindTermByQuantityUnitOfMeasureCode(s string) string {
	if len(t) == 0 {
		fmt.Println("no terminology data loaded")
		return ""
	}

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

func LoadLoinc(r io.Reader) (*LoincData, error) {
	if r == nil {
		r = bytes.NewReader(LoincDataFile)
	}

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		return gocsv.LazyCSVReader(in)
	})

	var data LoincData
	if err := gocsv.Unmarshal(r, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (l *LoincData) Len() int {
	if l == nil {
		return 0
	}

	return len(*l)
}
