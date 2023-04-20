package ncpdp

import (
	"encoding/xml"
	"time"
)

type Message struct {
	XMLName            xml.Name `xml:"Message" json:"-"`
	DatatypesVersion   string   `xml:"DatatypesVersion,attr" json:"datatypes_version,omitempty"`
	TransportVersion   string   `xml:"TransportVersion,attr" json:"transport_version,omitempty"`
	TransactionDomain  string   `xml:"TransactionDomain,attr" json:"transaction_domain,omitempty"`
	TransactionVersion string   `xml:"TransactionVersion,attr" json:"transaction_version,omitempty"`
	StructuresVersion  string   `xml:"StructuresVersion,attr" json:"structures_version,omitempty"`
	ECLVersion         string   `xml:"ECLVersion,attr" json:"ecl_version,omitempty"`
	Header             Header   `xml:"Header" json:"header,omitempty"`
	Body               Body     `xml:"Body" json:"body,omitempty"`
}

type Header struct {
	XMLName               xml.Name          `xml:"Header" json:"-"`
	To                    QualifierRef      `xml:"To" json:"to,omitempty"`
	From                  QualifierRef      `xml:"From" json:"from,omitempty"`
	MessageID             string            `xml:"MessageID" json:"message_id,omitempty"`
	RelatesToMessageID    string            `xml:"RelatesToMessageID" json:"relates_to_message_id,omitempty"`
	SentTime              time.Time         `xml:"SentTime" json:"sent_time,omitempty"`
	Security              Security          `xml:"Security" json:"security,omitempty"`
	SenderSoftware        SenderSoftware    `xml:"SenderSoftware" json:"sender_software,omitempty"`
	Mailbox               *Mailbox          `xml:"Mailbox" json:"mailbox,omitempty"`
	TestMessage           *bool             `xml:"TestMessage" json:"test_message,omitempty"`
	RxReferenceNumber     *string           `xml:"RxReferenceNumber" json:"rx_reference_number,omitempty"`
	PrescriberOrderNumber string            `xml:"PrescriberOrderNumber" json:"prescriber_order_number,omitempty"`
	DigitalSignature      *DigitalSignature `xml:"DigitalSignature" json:"digital_signature,omitempty"`
}

type QualifierRef struct {
	Value     string `xml:",chardata" json:"value,omitempty"`
	Qualifier string `xml:"Qualifier,attr" json:"qualifier,omitempty"`
}

type Security struct {
	XMLName       xml.Name                `xml:"Security" json:"-"`
	Sender        *TertiaryIdentification `xml:"Sender" json:"sender,omitempty"`
	Receiver      *TertiaryIdentification `xml:"Receiver" json:"receiver,omitempty"`
	UsernameToken *UsernameToken          `xml:"UsernameToken" json:"username_token,omitempty"`
}

type TertiaryIdentification struct {
	TertiaryIdentification string `xml:"TertiaryIdentification" json:"tertiary_identification,omitempty"`
}

type UsernameToken struct {
	Username string    `xml:"Username" json:"username,omitempty"`
	Password Password  `xml:"Password" json:"password,omitempty"`
	Nonce    string    `xml:"Nonce" json:"nonce,omitempty"`
	Created  time.Time `xml:"Created" json:"created,omitempty"`
}

type Password struct {
	Value string `xml:",chardata"`
	Type  string `xml:"Type,attr"`
}

type SenderSoftware struct {
	XMLName                      xml.Name `xml:"SenderSoftware" json:"-"`
	SenderSoftwareDeveloper      string   `xml:"SenderSoftwareDeveloper" json:"sender_software_developer,omitempty"`
	SenderSoftwareProduct        string   `xml:"SenderSoftwareProduct" json:"sender_software_product,omitempty"`
	SenderSoftwareVersionRelease string   `xml:"SenderSoftwareVersionRelease" json:"sender_software_version_release,omitempty"`
}

type Mailbox struct {
	XMLName           xml.Name `xml:"Mailbox" json:"-"`
	DeliveredID       *string  `xml:"DeliveredID" json:"delivered_id,omitempty"`
	AcknowledgementID *string  `xml:"AcknowledgementID" json:"acknowledgement_id,omitempty"`
}

type DigitalSignature struct {
	Version                   string `xml:"Version,attr" json:"version,omitempty"`
	DigitalSignatureIndicator bool   `xml:"DigitalSignatureIndicator" json:"digital_signature_indicator,omitempty"`
}

type Body struct {
	XMLName           xml.Name           `xml:"Body" json:"-"`
	NewRx             *NewRx             `xml:"NewRx" json:"new_rx,omitempty"`
	Status            *Coded             `xml:"Status" json:"status,omitempty"`
	Verify            *Verify            `xml:"Verify" json:"verify,omitempty"`
	RxRenewalRequest  *RxRenewalRequest  `xml:"RxRenewalRequest" json:"rx_renewal_request,omitempty"`
	RxRenewalResponse *RxRenewalResponse `xml:"RxRenewalResponse" json:"rx_renewal_response,omitempty"`
	CancelRx          *CancelRx          `xml:"CancelRx" json:"cancel_rx,omitempty"`
	Error             *Coded             `xml:"Error" json:"error,omitempty"`
}

type NewRx struct {
	XMLName               xml.Name               `xml:"NewRx" json:"-"`
	AllergyOrAdverseEvent *AllergyOrAdverseEvent `xml:"AllergyOrAdverseEvent" json:"allergy_or_adverse_event,omitempty"`
	BenefitsCoordination  *BenefitsCoordination  `xml:"BenefitsCoordination" json:"benefits_coordination,omitempty"`
	Patient               Patient                `xml:"Patient" json:"patient,omitempty"`
	Pharmacy              Pharmacy               `xml:"Pharmacy" json:"pharmacy,omitempty"`
	Prescriber            Prescriber             `xml:"Prescriber" json:"prescriber,omitempty"`
	Observation           *Observation           `xml:"Observation" json:"observation,omitempty"`
	MedicationPrescribed  Medication             `xml:"MedicationPrescribed" json:"medication_prescribed,omitempty"`
}

type Verify struct {
	XMLName      xml.Name `xml:"Verify" json:"-"`
	VerifyStatus *Coded   `xml:"VerifyStatus" json:"verify_status,omitempty"`
}

type RxRenewalRequest struct {
	XMLName                xml.Name   `xml:"RxRenewalRequest" json:"-"`
	RequestReferenceNumber *string    `xml:"RequestReferenceNumber" json:"request_reference_number,omitempty"`
	Patient                Patient    `xml:"Patient" json:"patient,omitempty"`
	Pharmacy               Pharmacy   `xml:"Pharmacy" json:"pharmacy,omitempty"`
	Prescriber             Prescriber `xml:"Prescriber" json:"prescriber,omitempty"`
	MedicationDispensed    Medication `xml:"MedicationDispensed" json:"medication_dispensed,omitempty"`
	MedicationPrescribed   Medication `xml:"MedicationPrescribed" json:"medication_prescribed,omitempty"`
}

type RxRenewalResponse struct {
	XMLName                xml.Name               `xml:"RxRenewalResponse" json:"-"`
	RequestReferenceNumber *string                `xml:"RequestReferenceNumber" json:"request_reference_number,omitempty"`
	Response               *Response              `xml:"Response" json:"response,omitempty"`
	AllergyOrAdverseEvent  *AllergyOrAdverseEvent `xml:"AllergyOrAdverseEvent" json:"allergy_or_adverse_event,omitempty"`
	Facility               *Facility              `xml:"Facility" json:"facility,omitempty"`
	Patient                Patient                `xml:"Patient" json:"patient,omitempty"`
	Pharmacy               Pharmacy               `xml:"Pharmacy" json:"pharmacy,omitempty"`
	Prescriber             Prescriber             `xml:"Prescriber" json:"prescriber,omitempty"`
	Supervisor             *Supervisor            `xml:"Supervisor" json:"supervisor,omitempty"`
	Observation            *Observation           `xml:"Observation" json:"observation,omitempty"`
	MedicationResponse     Medication             `xml:"MedicationResponse" json:"medication_response,omitempty"`
}

type CancelRx struct {
	XMLName              xml.Name   `xml:"CancelRx" json:"-"`
	Patient              Patient    `xml:"Patient" json:"patient,omitempty"`
	Pharmacy             Pharmacy   `xml:"Pharmacy" json:"pharmacy,omitempty"`
	Prescriber           Prescriber `xml:"Prescriber" json:"prescriber,omitempty"`
	MedicationPrescribed Medication `xml:"MedicationPrescribed" json:"medication_prescribed,omitempty"`
}

type Response struct {
	Approved            *Reason   `xml:"Approved" json:"approved,omitempty"`
	Replace             *struct{} `xml:"Replace" json:"replace,omitempty"`
	ApprovedWithChanges *struct{} `xml:"ApprovedWithChanges" json:"approved_with_changes,omitempty"`
	Denied              *Reason   `xml:"Denied" json:"denied,omitempty"`
}

type Reason struct {
	ReasonCode      *string `xml:"ReasonCode" json:"reason_code,omitempty"`
	ReferenceNumber *string `xml:"ReferenceNumber" json:"reference_number,omitempty"`
	DenialReason    *string `xml:"DenialReason" json:"denial_reason,omitempty"`
}

type Facility struct {
	Identification       PayerIdentification  `xml:"Identification" json:"identification,omitempty"`
	FacilityName         string               `xml:"FacilityName" json:"facility_name,omitempty"`
	Address              Address              `xml:"Address" json:"address,omitempty"`
	CommunicationNumbers CommunicationNumbers `xml:"CommunicationNumbers" json:"communication_numbers,omitempty"`
}

type Patient struct {
	XMLName      xml.Name     `xml:"Patient" json:"-"`
	HumanPatient HumanPatient `xml:"HumanPatient" json:"human_patient,omitempty"`
}

type HumanPatient struct {
	XMLName              xml.Name               `xml:"HumanPatient" json:"-"`
	Identification       *PatientIdentification `xml:"Identification" json:"identification,omitempty"`
	Name                 Name                   `xml:"Name" json:"name,omitempty"`
	Gender               string                 `xml:"Gender" json:"gender,omitempty"`
	DateOfBirth          DateOfBirth            `xml:"DateOfBirth" json:"date_of_birth,omitempty"`
	Address              Address                `xml:"Address" json:"address,omitempty"`
	CommunicationNumbers CommunicationNumbers   `xml:"CommunicationNumbers" json:"communication_numbers,omitempty"`
	LanguageNameCode     string                 `xml:"LanguageNameCode" json:"language_name_code,omitempty"`
}

type PatientIdentification struct {
	MedicalRecordIdentificationNumberEHR *string `xml:"MedicalRecordIdentificationNumberEHR" json:"medical_record_identification_number_ehr,omitempty"`
	SocialSecurity                       *string `xml:"SocialSecurity" json:"social_security,omitempty"`
}

type Name struct {
	FirstName  string  `xml:"FirstName" json:"first_name,omitempty"`
	MiddleName *string `xml:"MiddleName" json:"middle_name,omitempty"`
	LastName   string  `xml:"LastName" json:"last_name,omitempty"`
	Prefix     string  `xml:"Prefix" json:"prefix,omitempty"`
	Suffix     string  `xml:"Suffix" json:"suffix,omitempty"`
}

type DateOfBirth struct {
	XMLName xml.Name `xml:"DateOfBirth" json:"-"`
	Date    Date     `xml:"Date" json:"date,omitempty"`
}

type Date struct {
	time.Time
}

type DateTime struct {
	DateTime *time.Time `xml:"DateTime" json:"date_time,omitempty"`
}

type Address struct {
	XMLName       xml.Name `xml:"Address" json:"-"`
	AddressLine1  string   `xml:"AddressLine1" json:"address_line_1,omitempty"`
	AddressLine2  string   `xml:"AddressLine2" json:"address_line_2,omitempty"`
	City          string   `xml:"City" json:"city,omitempty"`
	StateProvince string   `xml:"StateProvince" json:"state_province,omitempty"`
	PostalCode    string   `xml:"PostalCode" json:"postal_code,omitempty"`
	CountryCode   string   `xml:"CountryCode" json:"country_code,omitempty"`
}

type PrescriberAgent struct {
	Name Name `xml:"Name" json:"name,omitempty"`
}

type CommunicationNumbers struct {
	XMLName          xml.Name   `xml:"CommunicationNumbers" json:"-"`
	PrimaryTelephone *Telephone `xml:"PrimaryTelephone" json:"primary_telephone,omitempty"`
	HomeTelephone    *Telephone `xml:"HomeTelephone" json:"home_telephone,omitempty"`
	OtherTelephone   *Telephone `xml:"OtherTelephone" json:"other_telephone,omitempty"`
	Fax              *Fax       `xml:"Fax" json:"fax,omitempty"`
	WorkTelephone    *Telephone `xml:"WorkTelephone" json:"work_telephone,omitempty"`
	ElectronicMail   string     `xml:"ElectronicMail" json:"electronic_mail,omitempty"`
}

type Telephone struct {
	Number      string `xml:"Number" json:"number,omitempty"`
	SupportsSMS string `xml:"SupportsSMS" json:"supports_sms,omitempty"`
}

type Fax struct {
	XMLName xml.Name `xml:"Fax" json:"-"`
	Number  string   `xml:"Number" json:"number,omitempty"`
}

type Pharmacy struct {
	XMLName              xml.Name               `xml:"Pharmacy" json:"-"`
	Identification       ProviderIdentification `xml:"Identification" json:"identification,omitempty"`
	Pharmacist           *Pharmacist            `xml:"Pharmacist" json:"pharmacist,omitempty"`
	BusinessName         string                 `xml:"BusinessName" json:"business_name,omitempty"`
	Address              Address                `xml:"Address" json:"address,omitempty"`
	CommunicationNumbers CommunicationNumbers   `xml:"CommunicationNumbers" json:"communication_numbers,omitempty"`
}

type ProviderIdentification struct {
	XMLName            xml.Name `xml:"Identification" json:"-"`
	NCPDPID            string   `xml:"NCPDPID" json:"ncpdpid,omitempty"`
	NPI                string   `xml:"NPI" json:"npi,omitempty"`
	DEANumber          string   `xml:"DEANumber" json:"dea_number,omitempty"`
	StateLicenseNumber string   `xml:"StateLicenseNumber" json:"state_license_number,omitempty"`
}

type Pharmacist struct {
	Name Name `xml:"Name" json:"name,omitempty"`
}

type Prescriber struct {
	XMLName         xml.Name        `xml:"Prescriber" json:"-"`
	NonVeterinarian NonVeterinarian `xml:"NonVeterinarian" json:"non_veterinarian,omitempty"`
}

type Supervisor struct {
	XMLName         xml.Name        `xml:"Supervisor" json:"-"`
	NonVeterinarian NonVeterinarian `xml:"NonVeterinarian" json:"non_veterinarian,omitempty"`
}

type NonVeterinarian struct {
	XMLName              xml.Name               `xml:"NonVeterinarian" json:"-"`
	Identification       ProviderIdentification `xml:"Identification" json:"identification,omitempty"`
	Specialty            string                 `xml:"Specialty" json:"specialty,omitempty"`
	PracticeLocation     *PracticeLocation      `xml:"PracticeLocation" json:"practice_location,omitempty"`
	Name                 Name                   `xml:"Name" json:"name,omitempty"`
	Address              Address                `xml:"Address" json:"address,omitempty"`
	PrescriberAgent      *PrescriberAgent       `xml:"PrescriberAgent" json:"prescriber_agent,omitempty"`
	CommunicationNumbers CommunicationNumbers   `xml:"CommunicationNumbers" json:"communication_numbers,omitempty"`
}

type PracticeLocation struct {
	XMLName      xml.Name `xml:"PracticeLocation" json:"-"`
	BusinessName string   `xml:"BusinessName" json:"business_name,omitempty"`
}

type Observation struct {
	Measurement []Measurement `xml:"Measurement" json:"measurement,omitempty"`
}

type Measurement struct {
	VitalSign       string           `xml:"VitalSign" json:"vital_sign,omitempty"`
	LOINCVersion    string           `xml:"LOINCVersion" json:"loinc_version,omitempty"`
	Value           string           `xml:"Value" json:"value,omitempty"`
	UnitOfMeasure   string           `xml:"UnitOfMeasure" json:"unit_of_measure,omitempty"`
	UCUMVersion     string           `xml:"UCUMVersion" json:"ucum_version,omitempty"`
	ObservationDate *ObservationDate `xml:"ObservationDate" json:"observation_date,omitempty"`
}

type ObservationDate struct {
	*DateTime
	Date *Date `xml:"Date" json:"date,omitempty"`
}

type Medication struct {
	DrugDescription           string               `xml:"DrugDescription" json:"drug_description,omitempty"`
	DrugCoded                 DrugCoded            `xml:"DrugCoded" json:"drug_coded,omitempty"`
	Quantity                  Quantity             `xml:"Quantity" json:"quantity,omitempty"`
	DaysSupply                float64              `xml:"DaysSupply" json:"days_supply,omitempty"`
	WrittenDate               WrittenDate          `xml:"WrittenDate" json:"written_date,omitempty"`
	LastFillDate              *LastFillDate        `xml:"LastFillDate" json:"last_fill_date,omitempty"`
	Substitutions             *int                 `xml:"Substitutions" json:"substitutions,omitempty"`
	NumberOfRefills           *int                 `xml:"NumberOfRefills" json:"number_of_refills,omitempty"`
	Diagnosis                 []Diagnosis          `xml:"Diagnosis" json:"diagnosis,omitempty"`
	Note                      string               `xml:"Note" json:"note,omitempty"`
	Sig                       Sig                  `xml:"Sig" json:"sig,omitempty"`
	RxFillIndicator           string               `xml:"RxFillIndicator" json:"rx_fill_indicator,omitempty"`
	PriorAuthorizationStatus  string               `xml:"PriorAuthorizationStatus" json:"prior_authorization_status,omitempty"`
	PrescriberCheckedREMS     string               `xml:"PrescriberCheckedREMS" json:"prescriber_checked_rems,omitempty"`
	OfficeOfPharmacyAffairsID string               `xml:"OfficeOfPharmacyAffairsID" json:"office_of_pharmacy_affairs_id,omitempty"`
	OtherMedicationDate       *OtherMedicationDate `xml:"OtherMedicationDate" json:"other_medication_date,omitempty"`
	PharmacyRequestedRefills  int                  `xml:"PharmacyRequestedRefills" json:"pharmacy_requested_refills,omitempty"`
}

type AllergyOrAdverseEvent struct {
	NoKnownAllergies string      `xml:"NoKnownAllergies" json:"no_known_allergies,omitempty"`
	Allergies        []Allergies `xml:"Allergies" json:"allergies,omitempty"`
}

type Allergies struct {
	SourceOfInformation string        `xml:"SourceOfInformation" json:"source_of_information,omitempty"`
	EffectiveDate       EffectiveDate `xml:"EffectiveDate" json:"effective_date,omitempty"`
	AdverseEvent        UnitOfMeasure `xml:"AdverseEvent" json:"adverse_event,omitempty"`
	DrugProductCoded    UnitOfMeasure `xml:"DrugProductCoded" json:"drug_product_coded,omitempty"`
}

type EffectiveDate struct {
	*DateTime
	Date *Date `xml:"Date" json:"date,omitempty"`
}

type BenefitsCoordination struct {
	PayerIdentification PayerIdentification `xml:"PayerIdentification" json:"payer_identification,omitempty"`
	PayerName           string              `xml:"PayerName" json:"payer_name,omitempty"`
	CardholderID        string              `xml:"CardholderID" json:"cardholder_id,omitempty"`
	CardHolderName      *Name               `xml:"CardHolderName" json:"card_holder_name,omitempty"`
	GroupID             string              `xml:"GroupID" json:"group_id,omitempty"`
	GroupName           string              `xml:"GroupName" json:"group_name,omitempty"`
	PBMMemberID         string              `xml:"PBMMemberID" json:"pbm_member_id,omitempty"`
}

type PayerIdentification struct {
	MutuallyDefined               string `xml:"MutuallyDefined" json:"mutually_defined,omitempty"`
	IINNumber                     string `xml:"IINNumber" json:"iin_number,omitempty"`
	PayerID                       string `xml:"PayerID" json:"payer_id,omitempty"`
	ProcessorIdentificationNumber string `xml:"ProcessorIdentificationNumber" json:"processor_identification_number,omitempty"`
}

type DrugCoded struct {
	XMLName     xml.Name     `xml:"DrugCoded" json:"-"`
	ProductCode Coded        `xml:"ProductCode" json:"product_code,omitempty"`
	Strength    *Strength    `xml:"Strength" json:"strength,omitempty"`
	DrugDBCode  *Coded       `xml:"DrugDBCode" json:"drug_db_code,omitempty"`
	DEASchedule *DEASchedule `xml:"DEASchedule" json:"dea_schedule,omitempty"`
}

type Coded struct {
	Code        string  `xml:"Code" json:"code,omitempty"`
	Qualifier   string  `xml:"Qualifier" json:"qualifier,omitempty"`
	Description *string `xml:"Description" json:"description,omitempty"`
}

type Strength struct {
	StrengthValue         string         `xml:"StrengthValue" json:"strength_value,omitempty"`
	StrengthForm          *UnitOfMeasure `xml:"StrengthForm" json:"strength_form,omitempty"`
	StrengthUnitOfMeasure *UnitOfMeasure `xml:"StrengthUnitOfMeasure" json:"strength_unit_of_measure,omitempty"`
}

type UnitOfMeasure struct {
	Text      *string `xml:"Text" json:"text,omitempty"`
	Qualifier *string `xml:"Qualifier" json:"qualifier,omitempty"`
	Code      *string `xml:"Code" json:"code,omitempty"`
}

type DEASchedule struct {
	Code string `xml:"Code" json:"code,omitempty"`
}

type Quantity struct {
	Value                 float64       `xml:"Value" json:"value,omitempty"`
	CodeListQualifier     string        `xml:"CodeListQualifier" json:"code_list_qualifier,omitempty"`
	QuantityUnitOfMeasure UnitOfMeasure `xml:"QuantityUnitOfMeasure" json:"quantity_unit_of_measure,omitempty"`
}

type WrittenDate struct {
	*DateTime
	Date *Date `xml:"Date" json:"date,omitempty"`
}

type LastFillDate struct {
	*DateTime
	Date *Date `xml:"Date" json:"date,omitempty"`
}

type Diagnosis struct {
	ClinicalInformationQualifier string `xml:"ClinicalInformationQualifier" json:"clinical_information_qualifier,omitempty"`
	Primary                      Coded  `xml:"Primary" json:"primary,omitempty"`
	Secondary                    *Coded `xml:"Secondary" json:"secondary,omitempty"`
}

type Sig struct {
	SigText     string       `xml:"SigText" json:"sig_text,omitempty"`
	CodeSystem  *CodeSystem  `xml:"CodeSystem" json:"code_system,omitempty"`
	Instruction *Instruction `xml:"Instruction" json:"instruction,omitempty"`
}

type CodeSystem struct {
	SNOMEDVersion string `xml:"SNOMEDVersion" json:"snomed_version,omitempty"`
	FMTVersion    string `xml:"FMTVersion" json:"fmt_version,omitempty"`
}

type Instruction struct {
	DoseAdministration DoseAdministration  `xml:"DoseAdministration" json:"dose_administration,omitempty"`
	TimingAndDuration  []TimingAndDuration `xml:"TimingAndDuration" json:"timing_and_duration,omitempty"`
}

type DoseAdministration struct {
	DoseDeliveryMethod    UnitOfMeasure `xml:"DoseDeliveryMethod" json:"dose_delivery_method,omitempty"`
	Dosage                Dosage        `xml:"Dosage" json:"dosage,omitempty"`
	RouteOfAdministration UnitOfMeasure `xml:"RouteOfAdministration" json:"route_of_administration,omitempty"`
}

type Dosage struct {
	DoseQuantity      int           `xml:"DoseQuantity" json:"dose_quantity,omitempty"`
	DoseUnitOfMeasure UnitOfMeasure `xml:"DoseUnitOfMeasure" json:"dose_unit_of_measure,omitempty"`
}

type TimingAndDuration struct {
	Frequency            *Frequency            `xml:"Frequency" json:"frequency,omitempty"`
	AdministrationTiming *AdministrationTiming `xml:"AdministrationTiming" json:"administration_timing,omitempty"`
}

type Frequency struct {
	FrequencyNumericValue int           `xml:"FrequencyNumericValue" json:"frequency_numeric_value,omitempty"`
	FrequencyUnits        UnitOfMeasure `xml:"FrequencyUnits" json:"frequency_units,omitempty"`
}

type AdministrationTiming struct {
	AdministrationTimingEvent UnitOfMeasure `xml:"AdministrationTimingEvent" json:"administration_timing_event,omitempty"`
}

type OtherMedicationDate struct {
	OtherMedicationDate2         *OtherMedicationDate2 `xml:"OtherMedicationDate" json:"other_medication_date,omitempty"`
	OtherMedicationDateQualifier string                `xml:"OtherMedicationDateQualifier" json:"other_medication_date_qualifier,omitempty"`
}

type OtherMedicationDate2 struct {
	Date Date `xml:"Date" json:"date,omitempty"`
}

func (t *Date) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var p struct {
		Date string `xml:",chardata"`
	}

	err := d.DecodeElement(&p, &start)
	if err != nil {
		return err
	}

	const format = "2006-01-02"
	parsed, err := time.Parse(format, p.Date)
	if err != nil {
		return err
	}

	t.Time = parsed
	return nil
}
