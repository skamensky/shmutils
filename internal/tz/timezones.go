package tz

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type TimeZone struct {
	Abbreviation string `json:"Abbreviation"`
	TimeZoneName string `json:"Time zone name"`
	Location     string `json:"Location"`
	Offset       string `json:"Offset"`
}

type Query struct {
	FromTimeZone string
	ToTimeZone   string
	FromTime     string
}

func (q *Query) Execute() (string, error) {
	fromTimezoneEntry := Abbreviations[q.FromTimeZone][0]
	toTimezoneEntry := Abbreviations[q.ToTimeZone][0]

	fromOffset, err := strconv.Atoi(strings.TrimSpace(strings.Replace(fromTimezoneEntry.Offset, "UTC", "", -1)))
	if err != nil {
		return "", err
	}
	toOffset, err := strconv.Atoi(strings.TrimSpace(strings.Replace(toTimezoneEntry.Offset, "UTC", "", -1)))
	if err != nil {
		return "", err
	}

	hour, err := strconv.Atoi(q.FromTime)
	if err != nil {
		return "", err
	}

	fromTimezone := time.FixedZone(fromTimezoneEntry.Abbreviation, fromOffset*3600)
	toTimezone := time.FixedZone(toTimezoneEntry.Abbreviation, toOffset*3600)
	fromTime := time.Date(2022, time.October, 15, hour, 0, 0, 0, fromTimezone)
	toTime := fromTime.In(toTimezone)

	message := ""

	message += fmt.Sprintf("%v in %v is %v in %v", fromTime.Format("15:04"), fromTimezoneEntry.Abbreviation, toTime.Format("15:04"), toTimezoneEntry.Abbreviation)

	if toTime.Day() > fromTime.Day() {
		message += " (next day)"
	}
	if toTime.Day() < fromTime.Day() {
		message += " (previous day)"
	}
	message += fmt.Sprintf(" - %d hour difference", int64(math.Abs(float64(toOffset-fromOffset))))

	return message, nil
}

var Abbreviations = map[string][]TimeZone{}

// from https://www.timeanddate.com/time/zones/
var TimeZones = []TimeZone{
	{
		Abbreviation: "A",
		TimeZoneName: "Alpha Time Zone",
		Location:     "Military",
		Offset:       "UTC +1",
	},
	{
		Abbreviation: "ACDT",
		TimeZoneName: "Australian Central Daylight Time CDT – Central Daylight TimeCDST – Central Daylight Savings Time",
		Location:     "Australia",
		Offset:       "UTC +10:30",
	},
	{
		Abbreviation: "ACST",
		TimeZoneName: "Australian Central Standard Time CST – Central Standard Time",
		Location:     "Australia",
		Offset:       "UTC +9:30",
	},
	{
		Abbreviation: "ACT",
		TimeZoneName: "Acre Time",
		Location:     "South America",
		Offset:       "UTC -5",
	},
	{
		Abbreviation: "ACT",
		TimeZoneName: "Australian Central Time",
		Location:     "Australia",
		Offset:       "UTC +9:30 / +10:30",
	},
	{
		Abbreviation: "ACWST",
		TimeZoneName: "Australian Central Western Standard Time",
		Location:     "Australia",
		Offset:       "UTC +8:45",
	},
	{
		Abbreviation: "ADT",
		TimeZoneName: "Arabia Daylight Time AST – Arabia Summer Time",
		Location:     "Asia",
		Offset:       "UTC +4",
	},
	{
		Abbreviation: "ADT",
		TimeZoneName: "Atlantic Daylight Time ADST – Atlantic Daylight Saving TimeAST – Atlantic Summer Time HAA – Heure Avancée de l'Atlantique (French)",
		Location:     "North AmericaAtlantic",
		Offset:       "UTC -3",
	},
	{
		Abbreviation: "AEDT",
		TimeZoneName: "Australian Eastern Daylight Time EDT – Eastern Daylight TimeEDST – Eastern Daylight Saving Time",
		Location:     "Australia",
		Offset:       "UTC +11",
	},
	{
		Abbreviation: "AEST",
		TimeZoneName: "Australian Eastern Standard Time EST – Eastern Standard TimeAET – Australian Eastern Time",
		Location:     "Australia",
		Offset:       "UTC +10",
	},
	{
		Abbreviation: "AET",
		TimeZoneName: "Australian Eastern Time",
		Location:     "Australia",
		Offset:       "UTC +10:00 / +11:00",
	},
	{
		Abbreviation: "AFT",
		TimeZoneName: "Afghanistan Time",
		Location:     "Asia",
		Offset:       "UTC +4:30",
	},
	{
		Abbreviation: "AKDT",
		TimeZoneName: "Alaska Daylight Time ADST – Alaska Daylight Saving Time",
		Location:     "North America",
		Offset:       "UTC -8",
	},
	{
		Abbreviation: "AKST",
		TimeZoneName: "Alaska Standard Time AT – Alaska Time",
		Location:     "North America",
		Offset:       "UTC -9",
	},
	{
		Abbreviation: "ALMT",
		TimeZoneName: "Alma-Ata Time",
		Location:     "Asia",
		Offset:       "UTC +6",
	},
	{
		Abbreviation: "AMST",
		TimeZoneName: "Amazon Summer Time",
		Location:     "South America",
		Offset:       "UTC -3",
	},
	{
		Abbreviation: "AMST",
		TimeZoneName: "Armenia Summer Time AMDT – Armenia Daylight Time",
		Location:     "Asia",
		Offset:       "UTC +5",
	},
	{
		Abbreviation: "AMT",
		TimeZoneName: "Amazon Time",
		Location:     "South America",
		Offset:       "UTC -4",
	},
	{
		Abbreviation: "AMT",
		TimeZoneName: "Armenia Time",
		Location:     "Asia",
		Offset:       "UTC +4",
	},
	{
		Abbreviation: "ANAST",
		TimeZoneName: "Anadyr Summer Time",
		Location:     "Asia",
		Offset:       "UTC +12",
	},
	{
		Abbreviation: "ANAT",
		TimeZoneName: "Anadyr Time",
		Location:     "Asia",
		Offset:       "UTC +12",
	},
	{
		Abbreviation: "AQTT",
		TimeZoneName: "Aqtobe Time",
		Location:     "Asia",
		Offset:       "UTC +5",
	},
	{
		Abbreviation: "ART",
		TimeZoneName: "Argentina Time",
		Location:     "AntarcticaSouth America",
		Offset:       "UTC -3",
	},
	{
		Abbreviation: "AST",
		TimeZoneName: "Arabia Standard Time AST – Arabic Standard TimeAST – Al Manamah Standard Time",
		Location:     "Asia",
		Offset:       "UTC +3",
	},
	{
		Abbreviation: "AST",
		TimeZoneName: "Atlantic Standard Time AT – Atlantic Time AST – Tiempo Estándar del Atlántico  (Spanish)HNA – Heure Normale de l'Atlantique (French)",
		Location:     "North AmericaAtlanticCaribbean",
		Offset:       "UTC -4",
	},
	{
		Abbreviation: "AT",
		TimeZoneName: "Atlantic Time",
		Location:     "North AmericaAtlantic",
		Offset:       "UTC -4:00 / -3:00",
	},
	{
		Abbreviation: "AWDT",
		TimeZoneName: "Australian Western Daylight Time WDT – Western Daylight TimeWST – Western Summer Time",
		Location:     "Australia",
		Offset:       "UTC +9",
	},
	{
		Abbreviation: "AWST",
		TimeZoneName: "Australian Western Standard Time WST – Western Standard TimeWAT – Western Australia Time",
		Location:     "Australia",
		Offset:       "UTC +8",
	},
	{
		Abbreviation: "AZOST",
		TimeZoneName: "Azores Summer Time AZODT – Azores Daylight Time",
		Location:     "Atlantic",
		Offset:       "UTC +0",
	},
	{
		Abbreviation: "AZOT",
		TimeZoneName: "Azores Time AZOST – Azores Standard Time",
		Location:     "Atlantic",
		Offset:       "UTC -1",
	},
	{
		Abbreviation: "AZST",
		TimeZoneName: "Azerbaijan Summer Time",
		Location:     "Asia",
		Offset:       "UTC +5",
	},
	{
		Abbreviation: "AZT",
		TimeZoneName: "Azerbaijan Time",
		Location:     "Asia",
		Offset:       "UTC +4",
	},
	{
		Abbreviation: "AoE",
		TimeZoneName: "Anywhere on Earth",
		Location:     "Pacific",
		Offset:       "UTC -12",
	},
	{
		Abbreviation: "B",
		TimeZoneName: "Bravo Time Zone",
		Location:     "Military",
		Offset:       "UTC +2",
	},
	{
		Abbreviation: "BNT",
		TimeZoneName: "Brunei Darussalam Time BDT – Brunei Time",
		Location:     "Asia",
		Offset:       "UTC +8",
	},
	{
		Abbreviation: "BOT",
		TimeZoneName: "Bolivia Time",
		Location:     "South America",
		Offset:       "UTC -4",
	},
	{
		Abbreviation: "BRST",
		TimeZoneName: "Brasília Summer Time BST – Brazil Summer TimeBST – Brazilian Summer Time",
		Location:     "South America",
		Offset:       "UTC -2",
	},
	{
		Abbreviation: "BRT",
		TimeZoneName: "Brasília Time BT – Brazil TimeBT – Brazilian Time",
		Location:     "South America",
		Offset:       "UTC -3",
	},
	{
		Abbreviation: "BST",
		TimeZoneName: "Bangladesh Standard Time",
		Location:     "Asia",
		Offset:       "UTC +6",
	},
	{
		Abbreviation: "BST",
		TimeZoneName: "Bougainville Standard Time",
		Location:     "Pacific",
		Offset:       "UTC +11",
	},
	{
		Abbreviation: "BST",
		TimeZoneName: "British Summer Time BDT – British Daylight TimeBDST – British Daylight Saving Time",
		Location:     "Europe",
		Offset:       "UTC +1",
	},
	{
		Abbreviation: "BTT",
		TimeZoneName: "Bhutan Time",
		Location:     "Asia",
		Offset:       "UTC +6",
	},
	{
		Abbreviation: "C",
		TimeZoneName: "Charlie Time Zone",
		Location:     "Military",
		Offset:       "UTC +3",
	},
	{
		Abbreviation: "CAST",
		TimeZoneName: "Casey Time",
		Location:     "Antarctica",
		Offset:       "UTC +8",
	},
	{
		Abbreviation: "CAT",
		TimeZoneName: "Central Africa Time",
		Location:     "Africa",
		Offset:       "UTC +2",
	},
	{
		Abbreviation: "CCT",
		TimeZoneName: "Cocos Islands Time",
		Location:     "Indian Ocean",
		Offset:       "UTC +6:30",
	},
	{
		Abbreviation: "CDT",
		TimeZoneName: "Central Daylight Time CDST – Central Daylight Saving TimeNACDT – North American Central Daylight Time HAC – Heure Avancée du Centre (French)",
		Location:     "North America",
		Offset:       "UTC -5",
	},
	{
		Abbreviation: "CDT",
		TimeZoneName: "Cuba Daylight Time",
		Location:     "Caribbean",
		Offset:       "UTC -4",
	},
	{
		Abbreviation: "CEST",
		TimeZoneName: "Central European Summer Time CEDT – Central European Daylight TimeECST – European Central Summer Time MESZ – Mitteleuropäische Sommerzeit (German)",
		Location:     "EuropeAntarctica",
		Offset:       "UTC +2",
	},
	{
		Abbreviation: "CET",
		TimeZoneName: "Central European Time ECT – European Central TimeCET – Central Europe Time MEZ – Mitteleuropäische Zeit (German)",
		Location:     "EuropeAfrica",
		Offset:       "UTC +1",
	},
	{
		Abbreviation: "CHADT",
		TimeZoneName: "Chatham Island Daylight Time CDT – Chatham Daylight Time",
		Location:     "Pacific",
		Offset:       "UTC +13:45",
	},
	{
		Abbreviation: "CHAST",
		TimeZoneName: "Chatham Island Standard Time",
		Location:     "Pacific",
		Offset:       "UTC +12:45",
	},
	{
		Abbreviation: "CHOST",
		TimeZoneName: "Choibalsan Summer Time CHODT – Choibalsan Daylight TimeCHODST – Choibalsan Daylight Saving Time",
		Location:     "Asia",
		Offset:       "UTC +9",
	},
	{
		Abbreviation: "CHOT",
		TimeZoneName: "Choibalsan Time",
		Location:     "Asia",
		Offset:       "UTC +8",
	},
	{
		Abbreviation: "CHUT",
		TimeZoneName: "Chuuk Time",
		Location:     "Pacific",
		Offset:       "UTC +10",
	},
	{
		Abbreviation: "CIDST",
		TimeZoneName: "Cayman Islands Daylight Saving Time",
		Location:     "Caribbean",
		Offset:       "UTC -4",
	},
	{
		Abbreviation: "CIST",
		TimeZoneName: "Cayman Islands Standard Time CIT – Cayman Islands Time",
		Location:     "Caribbean",
		Offset:       "UTC -5",
	},
	{
		Abbreviation: "CKT",
		TimeZoneName: "Cook Island Time",
		Location:     "Pacific",
		Offset:       "UTC -10",
	},
	{
		Abbreviation: "CLST",
		TimeZoneName: "Chile Summer Time CLDT – Chile Daylight Time",
		Location:     "South AmericaAntarctica",
		Offset:       "UTC -3",
	},
	{
		Abbreviation: "CLT",
		TimeZoneName: "Chile Standard Time CT – Chile TimeCLST – Chile Standard Time",
		Location:     "South AmericaAntarctica",
		Offset:       "UTC -4",
	},
	{
		Abbreviation: "COT",
		TimeZoneName: "Colombia Time",
		Location:     "South America",
		Offset:       "UTC -5",
	},
	{
		Abbreviation: "CST",
		TimeZoneName: "Central Standard Time CT – Central TimeNACST – North American Central Standard Time CST – Tiempo Central Estándar (Spanish)HNC – Heure Normale du Centre (French)",
		Location:     "North AmericaCentral America",
		Offset:       "UTC -6",
	},
	{
		Abbreviation: "CST",
		TimeZoneName: "China Standard Time",
		Location:     "Asia",
		Offset:       "UTC +8",
	},
	{
		Abbreviation: "CST",
		TimeZoneName: "Cuba Standard Time",
		Location:     "Caribbean",
		Offset:       "UTC -5",
	},
	{
		Abbreviation: "CT",
		TimeZoneName: "Central Time",
		Location:     "North America",
		Offset:       "UTC -6:00 / -5:00",
	},
	{
		Abbreviation: "CVT",
		TimeZoneName: "Cape Verde Time",
		Location:     "Africa",
		Offset:       "UTC -1",
	},
	{
		Abbreviation: "CXT",
		TimeZoneName: "Christmas Island Time",
		Location:     "Australia",
		Offset:       "UTC +7",
	},
	{
		Abbreviation: "ChST",
		TimeZoneName: "Chamorro Standard Time GST – Guam Standard Time",
		Location:     "Pacific",
		Offset:       "UTC +10",
	},
	{
		Abbreviation: "D",
		TimeZoneName: "Delta Time Zone",
		Location:     "Military",
		Offset:       "UTC +4",
	},
	{
		Abbreviation: "DAVT",
		TimeZoneName: "Davis Time",
		Location:     "Antarctica",
		Offset:       "UTC +7",
	},
	{
		Abbreviation: "DDUT",
		TimeZoneName: "Dumont-d'Urville Time",
		Location:     "Antarctica",
		Offset:       "UTC +10",
	},
	{
		Abbreviation: "E",
		TimeZoneName: "Echo Time Zone",
		Location:     "Military",
		Offset:       "UTC +5",
	},
	{
		Abbreviation: "EASST",
		TimeZoneName: "Easter Island Summer Time EADT – Easter Island Daylight Time",
		Location:     "Pacific",
		Offset:       "UTC -5",
	},
	{
		Abbreviation: "EAST",
		TimeZoneName: "Easter Island Standard Time",
		Location:     "Pacific",
		Offset:       "UTC -6",
	},
	{
		Abbreviation: "EAT",
		TimeZoneName: "Eastern Africa Time EAT – East Africa Time",
		Location:     "AfricaIndian Ocean",
		Offset:       "UTC +3",
	},
	{
		Abbreviation: "ECT",
		TimeZoneName: "Ecuador Time",
		Location:     "South America",
		Offset:       "UTC -5",
	},
	{
		Abbreviation: "EDT",
		TimeZoneName: "Eastern Daylight Time EDST – Eastern Daylight Savings TimeNAEDT – North American Eastern Daylight Time HAE – Heure Avancée de l'Est  (French)EDT – Tiempo de verano del Este (Spanish)",
		Location:     "North AmericaCaribbean",
		Offset:       "UTC -4",
	},
	{
		Abbreviation: "EEST",
		TimeZoneName: "Eastern European Summer Time EEDT – Eastern European Daylight Time OESZ – Osteuropäische Sommerzeit (German)",
		Location:     "EuropeAsia",
		Offset:       "UTC +3",
	},
	{
		Abbreviation: "EET",
		TimeZoneName: "Eastern European Time  OEZ – Osteuropäische Zeit (German)",
		Location:     "EuropeAsiaAfrica",
		Offset:       "UTC +2",
	},
	{
		Abbreviation: "EGST",
		TimeZoneName: "Eastern Greenland Summer Time EGST – East Greenland Summer Time",
		Location:     "North America",
		Offset:       "UTC +0",
	},
	{
		Abbreviation: "EGT",
		TimeZoneName: "East Greenland Time EGT – Eastern Greenland Time",
		Location:     "North America",
		Offset:       "UTC -1",
	},
	{
		Abbreviation: "EST",
		TimeZoneName: "Eastern Standard Time ET – Eastern Time NAEST – North American Eastern Standard Time ET – Tiempo del Este  (Spanish)HNE – Heure Normale de l'Est (French)",
		Location:     "North AmericaCaribbeanCentral America",
		Offset:       "UTC -5",
	},
	{
		Abbreviation: "ET",
		TimeZoneName: "Eastern Time",
		Location:     "North AmericaCaribbean",
		Offset:       "UTC -5:00 / -4:00",
	},
	{
		Abbreviation: "F",
		TimeZoneName: "Foxtrot Time Zone",
		Location:     "Military",
		Offset:       "UTC +6",
	},
	{
		Abbreviation: "FET",
		TimeZoneName: "Further-Eastern European Time",
		Location:     "Europe",
		Offset:       "UTC +3",
	},
	{
		Abbreviation: "FJST",
		TimeZoneName: "Fiji Summer Time FJDT – Fiji Daylight Time",
		Location:     "Pacific",
		Offset:       "UTC +13",
	},
	{
		Abbreviation: "FJT",
		TimeZoneName: "Fiji Time",
		Location:     "Pacific",
		Offset:       "UTC +12",
	},
	{
		Abbreviation: "FKST",
		TimeZoneName: "Falkland Islands Summer Time FKDT – Falkland Island Daylight Time",
		Location:     "South America",
		Offset:       "UTC -3",
	},
	{
		Abbreviation: "FKT",
		TimeZoneName: "Falkland Island Time FKST – Falkland Island Standard Time",
		Location:     "South America",
		Offset:       "UTC -4",
	},
	{
		Abbreviation: "FNT",
		TimeZoneName: "Fernando de Noronha Time",
		Location:     "South America",
		Offset:       "UTC -2",
	},
	{
		Abbreviation: "G",
		TimeZoneName: "Golf Time Zone",
		Location:     "Military",
		Offset:       "UTC +7",
	},
	{
		Abbreviation: "GALT",
		TimeZoneName: "Galapagos Time",
		Location:     "Pacific",
		Offset:       "UTC -6",
	},
	{
		Abbreviation: "GAMT",
		TimeZoneName: "Gambier Time GAMT – Gambier Islands Time",
		Location:     "Pacific",
		Offset:       "UTC -9",
	},
	{
		Abbreviation: "GET",
		TimeZoneName: "Georgia Standard Time",
		Location:     "AsiaEurope",
		Offset:       "UTC +4",
	},
	{
		Abbreviation: "GFT",
		TimeZoneName: "French Guiana Time",
		Location:     "South America",
		Offset:       "UTC -3",
	},
	{
		Abbreviation: "GILT",
		TimeZoneName: "Gilbert Island Time",
		Location:     "Pacific",
		Offset:       "UTC +12",
	},
	{
		Abbreviation: "GMT",
		TimeZoneName: "Greenwich Mean Time UTC – Coordinated Universal TimeGT – Greenwich Time",
		Location:     "EuropeAfricaNorth AmericaAntarctica",
		Offset:       "UTC +0",
	},
	{
		Abbreviation: "GST",
		TimeZoneName: "Gulf Standard Time",
		Location:     "Asia",
		Offset:       "UTC +4",
	},
	{
		Abbreviation: "GST",
		TimeZoneName: "South Georgia Time",
		Location:     "South America",
		Offset:       "UTC -2",
	},
	{
		Abbreviation: "GYT",
		TimeZoneName: "Guyana Time",
		Location:     "South America",
		Offset:       "UTC -4",
	},
	{
		Abbreviation: "H",
		TimeZoneName: "Hotel Time Zone",
		Location:     "Military",
		Offset:       "UTC +8",
	},
	{
		Abbreviation: "HDT",
		TimeZoneName: "Hawaii-Aleutian Daylight Time HADT – Hawaii Daylight Time",
		Location:     "North America",
		Offset:       "UTC -9",
	},
	{
		Abbreviation: "HKT",
		TimeZoneName: "Hong Kong Time",
		Location:     "Asia",
		Offset:       "UTC +8",
	},
	{
		Abbreviation: "HOVST",
		TimeZoneName: "Hovd Summer Time HOVDT – Hovd Daylight TimeHOVDST – Hovd Daylight Saving Time",
		Location:     "Asia",
		Offset:       "UTC +8",
	},
	{
		Abbreviation: "HOVT",
		TimeZoneName: "Hovd Time",
		Location:     "Asia",
		Offset:       "UTC +7",
	},
	{
		Abbreviation: "HST",
		TimeZoneName: "Hawaii Standard Time HAST – Hawaii-Aleutian Standard Time",
		Location:     "North AmericaPacific",
		Offset:       "UTC -10",
	},
	{
		Abbreviation: "I",
		TimeZoneName: "India Time Zone",
		Location:     "Military",
		Offset:       "UTC +9",
	},
	{
		Abbreviation: "ICT",
		TimeZoneName: "Indochina Time",
		Location:     "Asia",
		Offset:       "UTC +7",
	},
	{
		Abbreviation: "IDT",
		TimeZoneName: "Israel Daylight Time",
		Location:     "Asia",
		Offset:       "UTC +3",
	},
	{
		Abbreviation: "IOT",
		TimeZoneName: "Indian Chagos Time",
		Location:     "Indian Ocean",
		Offset:       "UTC +6",
	},
	{
		Abbreviation: "IRDT",
		TimeZoneName: "Iran Daylight Time IRST – Iran Summer TimeIDT – Iran Daylight Time",
		Location:     "Asia",
		Offset:       "UTC +4:30",
	},
	{
		Abbreviation: "IRKST",
		TimeZoneName: "Irkutsk Summer Time",
		Location:     "Asia",
		Offset:       "UTC +9",
	},
	{
		Abbreviation: "IRKT",
		TimeZoneName: "Irkutsk Time",
		Location:     "Asia",
		Offset:       "UTC +8",
	},
	{
		Abbreviation: "IRST",
		TimeZoneName: "Iran Standard Time IT – Iran Time",
		Location:     "Asia",
		Offset:       "UTC +3:30",
	},
	{
		Abbreviation: "IST",
		TimeZoneName: "India Standard Time IT – India TimeIST – Indian Standard Time",
		Location:     "Asia",
		Offset:       "UTC +5:30",
	},
	{
		Abbreviation: "IST",
		TimeZoneName: "Irish Standard Time IST – Irish Summer Time",
		Location:     "Europe",
		Offset:       "UTC +1",
	},
	{
		Abbreviation: "IST",
		TimeZoneName: "Israel Standard Time",
		Location:     "Asia",
		Offset:       "UTC +2",
	},
	{
		Abbreviation: "JST",
		TimeZoneName: "Japan Standard Time",
		Location:     "Asia",
		Offset:       "UTC +9",
	},
	{
		Abbreviation: "K",
		TimeZoneName: "Kilo Time Zone",
		Location:     "Military",
		Offset:       "UTC +10",
	},
	{
		Abbreviation: "KGT",
		TimeZoneName: "Kyrgyzstan Time",
		Location:     "Asia",
		Offset:       "UTC +6",
	},
	{
		Abbreviation: "KOST",
		TimeZoneName: "Kosrae Time",
		Location:     "Pacific",
		Offset:       "UTC +11",
	},
	{
		Abbreviation: "KRAST",
		TimeZoneName: "Krasnoyarsk Summer Time",
		Location:     "Asia",
		Offset:       "UTC +8",
	},
	{
		Abbreviation: "KRAT",
		TimeZoneName: "Krasnoyarsk Time",
		Location:     "Asia",
		Offset:       "UTC +7",
	},
	{
		Abbreviation: "KST",
		TimeZoneName: "Korea Standard Time KST – Korean Standard TimeKT – Korea Time",
		Location:     "Asia",
		Offset:       "UTC +9",
	},
	{
		Abbreviation: "KUYT",
		TimeZoneName: "Kuybyshev Time SAMST – Samara Summer Time",
		Location:     "Europe",
		Offset:       "UTC +4",
	},
	{
		Abbreviation: "L",
		TimeZoneName: "Lima Time Zone",
		Location:     "Military",
		Offset:       "UTC +11",
	},
	{
		Abbreviation: "LHDT",
		TimeZoneName: "Lord Howe Daylight Time",
		Location:     "Australia",
		Offset:       "UTC +11",
	},
	{
		Abbreviation: "LHST",
		TimeZoneName: "Lord Howe Standard Time",
		Location:     "Australia",
		Offset:       "UTC +10:30",
	},
	{
		Abbreviation: "LINT",
		TimeZoneName: "Line Islands Time",
		Location:     "Pacific",
		Offset:       "UTC +14",
	},
	{
		Abbreviation: "M",
		TimeZoneName: "Mike Time Zone",
		Location:     "Military",
		Offset:       "UTC +12",
	},
	{
		Abbreviation: "MAGST",
		TimeZoneName: "Magadan Summer Time MAGST – Magadan Island Summer Time",
		Location:     "Asia",
		Offset:       "UTC +12",
	},
	{
		Abbreviation: "MAGT",
		TimeZoneName: "Magadan Time MAGT – Magadan Island Time",
		Location:     "Asia",
		Offset:       "UTC +11",
	},
	{
		Abbreviation: "MART",
		TimeZoneName: "Marquesas Time",
		Location:     "Pacific",
		Offset:       "UTC -9:30",
	},
	{
		Abbreviation: "MAWT",
		TimeZoneName: "Mawson Time",
		Location:     "Antarctica",
		Offset:       "UTC +5",
	},
	{
		Abbreviation: "MDT",
		TimeZoneName: "Mountain Daylight Time MDST – Mountain Daylight Saving TimeNAMDT – North American Mountain Daylight Time HAR – Heure Avancée des Rocheuses (French)",
		Location:     "North America",
		Offset:       "UTC -6",
	},
	{
		Abbreviation: "MHT",
		TimeZoneName: "Marshall Islands Time",
		Location:     "Pacific",
		Offset:       "UTC +12",
	},
	{
		Abbreviation: "MMT",
		TimeZoneName: "Myanmar Time",
		Location:     "Asia",
		Offset:       "UTC +6:30",
	},
	{
		Abbreviation: "MSD",
		TimeZoneName: "Moscow Daylight Time Moscow Summer Time",
		Location:     "Europe",
		Offset:       "UTC +4",
	},
	{
		Abbreviation: "MSK",
		TimeZoneName: "Moscow Standard Time MCK – Moscow Time",
		Location:     "EuropeAsia",
		Offset:       "UTC +3",
	},
	{
		Abbreviation: "MST",
		TimeZoneName: "Mountain Standard Time MT – Mountain TimeNAMST – North American Mountain Standard Time HNR – Heure Normale des Rocheuses (French)",
		Location:     "North America",
		Offset:       "UTC -7",
	},
	{
		Abbreviation: "MT",
		TimeZoneName: "Mountain Time",
		Location:     "North America",
		Offset:       "UTC -7:00 / -6:00",
	},
	{
		Abbreviation: "MUT",
		TimeZoneName: "Mauritius Time",
		Location:     "Africa",
		Offset:       "UTC +4",
	},
	{
		Abbreviation: "MVT",
		TimeZoneName: "Maldives Time",
		Location:     "Asia",
		Offset:       "UTC +5",
	},
	{
		Abbreviation: "MYT",
		TimeZoneName: "Malaysia Time MST – Malaysian Standard Time",
		Location:     "Asia",
		Offset:       "UTC +8",
	},
	{
		Abbreviation: "N",
		TimeZoneName: "November Time Zone",
		Location:     "Military",
		Offset:       "UTC -1",
	},
	{
		Abbreviation: "NCT",
		TimeZoneName: "New Caledonia Time",
		Location:     "Pacific",
		Offset:       "UTC +11",
	},
	{
		Abbreviation: "NDT",
		TimeZoneName: "Newfoundland Daylight Time  HAT – Heure Avancée de Terre-Neuve (French)",
		Location:     "North America",
		Offset:       "UTC -2:30",
	},
	{
		Abbreviation: "NFDT",
		TimeZoneName: "Norfolk Daylight Time NFDT – Norfolk Island Daylight Time",
		Location:     "Australia",
		Offset:       "UTC +12",
	},
	{
		Abbreviation: "NFT",
		TimeZoneName: "Norfolk Time NFT – Norfolk Island Time",
		Location:     "Australia",
		Offset:       "UTC +11",
	},
	{
		Abbreviation: "NOVST",
		TimeZoneName: "Novosibirsk Summer Time OMSST – Omsk Summer Time",
		Location:     "Asia",
		Offset:       "UTC +7",
	},
	{
		Abbreviation: "NOVT",
		TimeZoneName: "Novosibirsk Time OMST – Omsk Standard Time",
		Location:     "Asia",
		Offset:       "UTC +7",
	},
	{
		Abbreviation: "NPT",
		TimeZoneName: "Nepal Time",
		Location:     "Asia",
		Offset:       "UTC +5:45",
	},
	{
		Abbreviation: "NRT",
		TimeZoneName: "Nauru Time",
		Location:     "Pacific",
		Offset:       "UTC +12",
	},
	{
		Abbreviation: "NST",
		TimeZoneName: "Newfoundland Standard Time  HNT – Heure Normale de Terre-Neuve (French)",
		Location:     "North America",
		Offset:       "UTC -3:30",
	},
	{
		Abbreviation: "NUT",
		TimeZoneName: "Niue Time",
		Location:     "Pacific",
		Offset:       "UTC -11",
	},
	{
		Abbreviation: "NZDT",
		TimeZoneName: "New Zealand Daylight Time",
		Location:     "PacificAntarctica",
		Offset:       "UTC +13",
	},
	{
		Abbreviation: "NZST",
		TimeZoneName: "New Zealand Standard Time",
		Location:     "PacificAntarctica",
		Offset:       "UTC +12",
	},
	{
		Abbreviation: "O",
		TimeZoneName: "Oscar Time Zone",
		Location:     "Military",
		Offset:       "UTC -2",
	},
	{
		Abbreviation: "OMSST",
		TimeZoneName: "Omsk Summer Time NOVST – Novosibirsk Summer Time",
		Location:     "Asia",
		Offset:       "UTC +7",
	},
	{
		Abbreviation: "OMST",
		TimeZoneName: "Omsk Standard Time OMST – Omsk TimeNOVT – Novosibirsk Time",
		Location:     "Asia",
		Offset:       "UTC +6",
	},
	{
		Abbreviation: "ORAT",
		TimeZoneName: "Oral Time",
		Location:     "Asia",
		Offset:       "UTC +5",
	},
	{
		Abbreviation: "P",
		TimeZoneName: "Papa Time Zone",
		Location:     "Military",
		Offset:       "UTC -3",
	},
	{
		Abbreviation: "PDT",
		TimeZoneName: "Pacific Daylight Time PDST – Pacific Daylight Saving TimeNAPDT – North American Pacific Daylight Time HAP – Heure Avancée du Pacifique (French)",
		Location:     "North America",
		Offset:       "UTC -7",
	},
	{
		Abbreviation: "PET",
		TimeZoneName: "Peru Time",
		Location:     "South America",
		Offset:       "UTC -5",
	},
	{
		Abbreviation: "PETST",
		TimeZoneName: "Kamchatka Summer Time",
		Location:     "Asia",
		Offset:       "UTC +12",
	},
	{
		Abbreviation: "PETT",
		TimeZoneName: "Kamchatka Time PETT – Petropavlovsk-Kamchatski Time",
		Location:     "Asia",
		Offset:       "UTC +12",
	},
	{
		Abbreviation: "PGT",
		TimeZoneName: "Papua New Guinea Time",
		Location:     "Pacific",
		Offset:       "UTC +10",
	},
	{
		Abbreviation: "PHOT",
		TimeZoneName: "Phoenix Island Time",
		Location:     "Pacific",
		Offset:       "UTC +13",
	},
	{
		Abbreviation: "PHT",
		TimeZoneName: "Philippine Time PST – Philippine Standard Time",
		Location:     "Asia",
		Offset:       "UTC +8",
	},
	{
		Abbreviation: "PKT",
		TimeZoneName: "Pakistan Standard Time PKT – Pakistan Time",
		Location:     "Asia",
		Offset:       "UTC +5",
	},
	{
		Abbreviation: "PMDT",
		TimeZoneName: "Pierre & Miquelon Daylight Time",
		Location:     "North America",
		Offset:       "UTC -2",
	},
	{
		Abbreviation: "PMST",
		TimeZoneName: "Pierre & Miquelon Standard Time",
		Location:     "North America",
		Offset:       "UTC -3",
	},
	{
		Abbreviation: "PONT",
		TimeZoneName: "Pohnpei Standard Time",
		Location:     "Pacific",
		Offset:       "UTC +11",
	},
	{
		Abbreviation: "PST",
		TimeZoneName: "Pacific Standard Time PT – Pacific TimeNAPST – North American Pacific Standard Time PT – Tiempo del Pacífico (Spanish)HNP – Heure Normale du Pacifique (French)",
		Location:     "North America",
		Offset:       "UTC -8",
	},
	{
		Abbreviation: "PST",
		TimeZoneName: "Pitcairn Standard Time",
		Location:     "Pacific",
		Offset:       "UTC -8",
	},
	{
		Abbreviation: "PT",
		TimeZoneName: "Pacific Time",
		Location:     "North America",
		Offset:       "UTC -8:00 / -7:00",
	},
	{
		Abbreviation: "PWT",
		TimeZoneName: "Palau Time",
		Location:     "Pacific",
		Offset:       "UTC +9",
	},
	{
		Abbreviation: "PYST",
		TimeZoneName: "Paraguay Summer Time",
		Location:     "South America",
		Offset:       "UTC -3",
	},
	{
		Abbreviation: "PYT",
		TimeZoneName: "Paraguay Time",
		Location:     "South America",
		Offset:       "UTC -4",
	},
	{
		Abbreviation: "PYT",
		TimeZoneName: "Pyongyang Time PYST – Pyongyang Standard Time",
		Location:     "Asia",
		Offset:       "UTC +8:30",
	},
	{
		Abbreviation: "Q",
		TimeZoneName: "Quebec Time Zone",
		Location:     "Military",
		Offset:       "UTC -4",
	},
	{
		Abbreviation: "QYZT",
		TimeZoneName: "Qyzylorda Time",
		Location:     "Asia",
		Offset:       "UTC +6",
	},
	{
		Abbreviation: "R",
		TimeZoneName: "Romeo Time Zone",
		Location:     "Military",
		Offset:       "UTC -5",
	},
	{
		Abbreviation: "RET",
		TimeZoneName: "Reunion Time",
		Location:     "Africa",
		Offset:       "UTC +4",
	},
	{
		Abbreviation: "ROTT",
		TimeZoneName: "Rothera Time",
		Location:     "Antarctica",
		Offset:       "UTC -3",
	},
	{
		Abbreviation: "S",
		TimeZoneName: "Sierra Time Zone",
		Location:     "Military",
		Offset:       "UTC -6",
	},
	{
		Abbreviation: "SAKT",
		TimeZoneName: "Sakhalin Time",
		Location:     "Asia",
		Offset:       "UTC +11",
	},
	{
		Abbreviation: "SAMT",
		TimeZoneName: "Samara Time SAMT – Samara Standard Time",
		Location:     "Europe",
		Offset:       "UTC +4",
	},
	{
		Abbreviation: "SAST",
		TimeZoneName: "South Africa Standard Time SAST – South African Standard Time",
		Location:     "Africa",
		Offset:       "UTC +2",
	},
	{
		Abbreviation: "SBT",
		TimeZoneName: "Solomon Islands Time SBT – Solomon Island Time",
		Location:     "Pacific",
		Offset:       "UTC +11",
	},
	{
		Abbreviation: "SCT",
		TimeZoneName: "Seychelles Time",
		Location:     "Africa",
		Offset:       "UTC +4",
	},
	{
		Abbreviation: "SGT",
		TimeZoneName: "Singapore Time SST – Singapore Standard Time",
		Location:     "Asia",
		Offset:       "UTC +8",
	},
	{
		Abbreviation: "SRET",
		TimeZoneName: "Srednekolymsk Time",
		Location:     "Asia",
		Offset:       "UTC +11",
	},
	{
		Abbreviation: "SRT",
		TimeZoneName: "Suriname Time",
		Location:     "South America",
		Offset:       "UTC -3",
	},
	{
		Abbreviation: "SST",
		TimeZoneName: "Samoa Standard Time",
		Location:     "Pacific",
		Offset:       "UTC -11",
	},
	{
		Abbreviation: "SYOT",
		TimeZoneName: "Syowa Time",
		Location:     "Antarctica",
		Offset:       "UTC +3",
	},
	{
		Abbreviation: "T",
		TimeZoneName: "Tango Time Zone",
		Location:     "Military",
		Offset:       "UTC -7",
	},
	{
		Abbreviation: "TAHT",
		TimeZoneName: "Tahiti Time",
		Location:     "Pacific",
		Offset:       "UTC -10",
	},
	{
		Abbreviation: "TFT",
		TimeZoneName: "French Southern and Antarctic Time KIT – Kerguelen (Islands) Time",
		Location:     "Indian Ocean",
		Offset:       "UTC +5",
	},
	{
		Abbreviation: "TJT",
		TimeZoneName: "Tajikistan Time",
		Location:     "Asia",
		Offset:       "UTC +5",
	},
	{
		Abbreviation: "TKT",
		TimeZoneName: "Tokelau Time",
		Location:     "Pacific",
		Offset:       "UTC +13",
	},
	{
		Abbreviation: "TLT",
		TimeZoneName: "East Timor Time",
		Location:     "Asia",
		Offset:       "UTC +9",
	},
	{
		Abbreviation: "TMT",
		TimeZoneName: "Turkmenistan Time",
		Location:     "Asia",
		Offset:       "UTC +5",
	},
	{
		Abbreviation: "TOST",
		TimeZoneName: "Tonga Summer Time",
		Location:     "Pacific",
		Offset:       "UTC +14",
	},
	{
		Abbreviation: "TOT",
		TimeZoneName: "Tonga Time",
		Location:     "Pacific",
		Offset:       "UTC +13",
	},
	{
		Abbreviation: "TRT",
		TimeZoneName: "Turkey Time",
		Location:     "AsiaEurope",
		Offset:       "UTC +3",
	},
	{
		Abbreviation: "TVT",
		TimeZoneName: "Tuvalu Time",
		Location:     "Pacific",
		Offset:       "UTC +12",
	},
	{
		Abbreviation: "U",
		TimeZoneName: "Uniform Time Zone",
		Location:     "Military",
		Offset:       "UTC -8",
	},
	{
		Abbreviation: "ULAST",
		TimeZoneName: "Ulaanbaatar Summer Time ULAST – Ulan Bator Summer Time",
		Location:     "Asia",
		Offset:       "UTC +9",
	},
	{
		Abbreviation: "ULAT",
		TimeZoneName: "Ulaanbaatar Time ULAT – Ulan Bator Time",
		Location:     "Asia",
		Offset:       "UTC +8",
	},
	{
		Abbreviation: "UTC",
		TimeZoneName: "Coordinated Universal Time",
		Location:     "Worldwide",
		Offset:       "UTC",
	},
	{
		Abbreviation: "UYST",
		TimeZoneName: "Uruguay Summer Time",
		Location:     "South America",
		Offset:       "UTC -2",
	},
	{
		Abbreviation: "UYT",
		TimeZoneName: "Uruguay Time",
		Location:     "South America",
		Offset:       "UTC -3",
	},
	{
		Abbreviation: "UZT",
		TimeZoneName: "Uzbekistan Time",
		Location:     "Asia",
		Offset:       "UTC +5",
	},
	{
		Abbreviation: "V",
		TimeZoneName: "Victor Time Zone",
		Location:     "Military",
		Offset:       "UTC -9",
	},
	{
		Abbreviation: "VET",
		TimeZoneName: "Venezuelan Standard Time  HLV – Hora Legal de Venezuela (Spanish)",
		Location:     "South America",
		Offset:       "UTC -4",
	},
	{
		Abbreviation: "VLAST",
		TimeZoneName: "Vladivostok Summer Time",
		Location:     "Asia",
		Offset:       "UTC +11",
	},
	{
		Abbreviation: "VLAT",
		TimeZoneName: "Vladivostok Time",
		Location:     "Asia",
		Offset:       "UTC +10",
	},
	{
		Abbreviation: "VOST",
		TimeZoneName: "Vostok Time",
		Location:     "Antarctica",
		Offset:       "UTC +6",
	},
	{
		Abbreviation: "VUT",
		TimeZoneName: "Vanuatu Time EFATE – Efate Time",
		Location:     "Pacific",
		Offset:       "UTC +11",
	},
	{
		Abbreviation: "W",
		TimeZoneName: "Whiskey Time Zone",
		Location:     "Military",
		Offset:       "UTC -10",
	},
	{
		Abbreviation: "WAKT",
		TimeZoneName: "Wake Time",
		Location:     "Pacific",
		Offset:       "UTC +12",
	},
	{
		Abbreviation: "WARST",
		TimeZoneName: "Western Argentine Summer Time",
		Location:     "South America",
		Offset:       "UTC -3",
	},
	{
		Abbreviation: "WAST",
		TimeZoneName: "West Africa Summer Time",
		Location:     "Africa",
		Offset:       "UTC +2",
	},
	{
		Abbreviation: "WAT",
		TimeZoneName: "West Africa Time",
		Location:     "Africa",
		Offset:       "UTC +1",
	},
	{
		Abbreviation: "WEST",
		TimeZoneName: "Western European Summer Time WEDT – Western European Daylight Time WESZ – Westeuropäische Sommerzeit (German)",
		Location:     "EuropeAfrica",
		Offset:       "UTC +1",
	},
	{
		Abbreviation: "WET",
		TimeZoneName: "Western European Time GMT – Greenwich Mean Time WEZ – Westeuropäische Zeit (German)",
		Location:     "EuropeAfrica",
		Offset:       "UTC +0",
	},
	{
		Abbreviation: "WFT",
		TimeZoneName: "Wallis and Futuna Time",
		Location:     "Pacific",
		Offset:       "UTC +12",
	},
	{
		Abbreviation: "WGST",
		TimeZoneName: "Western Greenland Summer Time WGST – West Greenland Summer Time",
		Location:     "North America",
		Offset:       "UTC -2",
	},
	{
		Abbreviation: "WGT",
		TimeZoneName: "West Greenland Time WGT – Western Greenland Time",
		Location:     "North America",
		Offset:       "UTC -3",
	},
	{
		Abbreviation: "WIB",
		TimeZoneName: "Western Indonesian Time WIB – Waktu Indonesia Barat",
		Location:     "Asia",
		Offset:       "UTC +7",
	},
	{
		Abbreviation: "WIT",
		TimeZoneName: "Eastern Indonesian Time WIT – Waktu Indonesia Timur",
		Location:     "Asia",
		Offset:       "UTC +9",
	},
	{
		Abbreviation: "WITA",
		TimeZoneName: "Central Indonesian Time WITA – Waktu Indonesia Tengah",
		Location:     "Asia",
		Offset:       "UTC +8",
	},
	{
		Abbreviation: "WST",
		TimeZoneName: "West Samoa Time ST – Samoa Time",
		Location:     "Pacific",
		Offset:       "UTC +13",
	},
	{
		Abbreviation: "WST",
		TimeZoneName: "Western Sahara Summer Time",
		Location:     "Africa",
		Offset:       "UTC +1",
	},
	{
		Abbreviation: "WT",
		TimeZoneName: "Western Sahara Standard Time WT – Western Sahara Time",
		Location:     "Africa",
		Offset:       "UTC +0",
	},
	{
		Abbreviation: "X",
		TimeZoneName: "X-ray Time Zone",
		Location:     "Military",
		Offset:       "UTC -11",
	},
	{
		Abbreviation: "Y",
		TimeZoneName: "Yankee Time Zone",
		Location:     "Military",
		Offset:       "UTC -12",
	},
	{
		Abbreviation: "YAKST",
		TimeZoneName: "Yakutsk Summer Time",
		Location:     "Asia",
		Offset:       "UTC +10",
	},
	{
		Abbreviation: "YAKT",
		TimeZoneName: "Yakutsk Time",
		Location:     "Asia",
		Offset:       "UTC +9",
	},
	{
		Abbreviation: "YAPT",
		TimeZoneName: "Yap Time",
		Location:     "Pacific",
		Offset:       "UTC +10",
	},
	{
		Abbreviation: "YEKST",
		TimeZoneName: "Yekaterinburg Summer Time",
		Location:     "Asia",
		Offset:       "UTC +6",
	},
	{
		Abbreviation: "YEKT",
		TimeZoneName: "Yekaterinburg Time",
		Location:     "Asia",
		Offset:       "UTC +5",
	},
	{
		Abbreviation: "Z",
		TimeZoneName: "Zulu Time Zone",
		Location:     "Military",
		Offset:       "UTC +0",
	},
}

func SetAbbreviations() {
	for _, tz := range TimeZones {
		// make an exception for IST since I use it so often. Exclude others
		if tz.Abbreviation == "IST" && !strings.Contains("Israel", tz.TimeZoneName) {
			continue
		}
		Abbreviations[tz.Abbreviation] = append(Abbreviations[tz.Abbreviation], tz)
	}
}

func RemoveSpaces(list *[]string) {
	cleaned := []string{}
	for _, item := range *list {
		if item != "" {
			cleaned = append(cleaned, item)
		}
	}
	*list = cleaned
}

// validates time part. It can accept 24 hour time or 12 hour time with am/pm.
// if given 12 hour time, it modified the input part to normalize it to 24 hour time
func ValidateTimePart(parts *[]string, position int) error {
	part := (*parts)[position]
	part = strings.ToLower(part)

	if len(part) > 4 {
		return fmt.Errorf("invalid time part at position %d. Expected a maximum of 4 characters and got %d", position, len(part))
	}

	// 24 hour time
	if len(part) <= 2 {

		val, err := strconv.Atoi(part)
		if err != nil {
			return fmt.Errorf("invalid time part at position %d. Expected a number and got %s", position, part)
		}
		if val < 0 || val > 23 {
			return fmt.Errorf("invalid time part at position %d. Expected a number between 0 and 23 and got %d", position, val)
		}
	} else {
		isAM := strings.Contains(part, "am")
		isPM := strings.Contains(part, "pm")
		if !isAM && !isPM {
			return fmt.Errorf("invalid time part at position %d. Expected am or pm and got %s", position, part)
		}

		no_meridian := strings.Replace(part, "am", "", -1)
		no_meridian = strings.Replace(no_meridian, "pm", "", -1)

		// 12 hour time
		val, err := strconv.Atoi(no_meridian)
		if err != nil {
			return fmt.Errorf("invalid time part at position %d. Expected a number and got %s", position, part[:1])
		}
		if val < 0 || val > 12 {
			return fmt.Errorf("invalid time part at position %d. Expected a number between 0 and 12 and got %d", position, val)
		}
		// normalize to 24 hour time
		if isPM {
			(*parts)[position] = fmt.Sprintf("%d", val+12)
		} else {
			(*parts)[position] = fmt.Sprintf("%d", val)
		}
	}
	return nil
}

func ValidateTimeZone(parts *[]string, position int) error {

	part := (*parts)[position]
	part = strings.ToUpper(part)
	(*parts)[position] = part
	zones := Abbreviations[part]
	if len(zones) == 0 {
		return fmt.Errorf("invalid timezone: %s at position %d", (*parts)[position], position)
	}
	if len(zones) > 1 {
		zonesAsString := []string{}
		for _, zone := range zones {
			zonesAsString = append(zonesAsString, zone.TimeZoneName)
		}
		return fmt.Errorf("ambiguous timezone: %s at position %d. Currently there is only support for non-ambigious time zones. Possible time zones: %s", (*parts)[position], position, strings.Join(zonesAsString, ","))
	}
	return nil
}

func ValidatePreposition(parts *[]string, position int) error {
	part := (*parts)[position]
	if part != "to" && part != "in" {
		return fmt.Errorf("invalid preposition: %s at position %d. Expected 'to' or 'in'", part, position)
	}
	return nil
}

func ExecuteQuery(query string) (string, error) {
	SetAbbreviations()
	parts := strings.Split(strings.ToLower(query), " ")
	partsCleaned := make([]string, len(parts))
	copy(partsCleaned, parts)

	RemoveSpaces(&parts)
	for idx, part := range partsCleaned {
		if part == "am" || part == "pm" {
			if idx == 0 {
				// can't set previous value but this is an error and will be caught later
				continue
			}
			partsCleaned[idx-1] = partsCleaned[idx-1] + part
			// scheduled for cleanup
			partsCleaned[idx] = ""
		}
	}
	RemoveSpaces(&partsCleaned)

	if len(partsCleaned) != 4 {
		return "", fmt.Errorf("unable to parse query. Expected 4 items, got %d", len(partsCleaned))
	}

	errors := []error{}
	// structure of query should be "fromTime fromTimeZone [to|in] toTimeZone"
	errors = append(errors, ValidateTimePart(&partsCleaned, 0))
	errors = append(errors, ValidateTimeZone(&partsCleaned, 1))
	errors = append(errors, ValidatePreposition(&partsCleaned, 2))
	errors = append(errors, ValidateTimeZone(&partsCleaned, 3))

	errMessages := []string{}
	for _, err := range errors {
		if err != nil {
			errMessages = append(errMessages, err.Error())
		}
	}
	if len(errMessages) > 0 {
		return "", fmt.Errorf("unable to parse query: Error(s): %s", strings.Join(errMessages, ", "))
	}

	queryObj := Query{
		FromTime:     partsCleaned[0],
		FromTimeZone: partsCleaned[1],
		// ignore preposition
		ToTimeZone: partsCleaned[3],
	}

	return queryObj.Execute()
}
