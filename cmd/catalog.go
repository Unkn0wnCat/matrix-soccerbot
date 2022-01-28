// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package cmd

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p, ok := messageKeyToIndex[key]
	if !ok {
		return "", false
	}
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"de": &dictionary{index: deIndex, data: deData},
		"en": &dictionary{index: enIndex, data: enData},
	}
	fallback := language.MustParse("en")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	" (Overtime)": 15,
	" (Own goal)": 16,
	" (Penalty)":  14,
	"# Game <font color=\"%[5]s\">%[1]s</font> vs. <font color=\"%[6]s\">%[2]s</font> ***(<font color=\"%[5]s\">%[3]d</font>:<font color=\"%[6]s\">%[4]d</font>)***\n\n": 7,
	"## **Goals**\n\n": 11,
	"## Results\n\n":   18,
	"* %[1]s - <font color=\"%[2]s\">%[4]d</font>:<font color=\"%[3]s\">%[5]d</font> - Goal for %[6]s by %[7]s%[8]s\n": 17,
	"* **Halftime result":            19,
	"* **Result":                     24,
	"* **Result after 90 minutes":    20,
	"* **Result after extended Time": 21,
	"* **Result after overtime":      22,
	"* **Result after penalty shots": 23,
	"**Game Completed** - ":          8,
	"*Game began %s*\n\n\n\n":        9,
	"*Game beginning %s*\n\n\n\n":    10,
	"*No goals yet*\n\n":             12,
	"Data provided by [OpenLigaDB.de](https://www.openligadb.de) | [Sourcecode](https://github.com/Unkn0wnCat/matrix-soccerbot)": 25,
	"Please provide either an access-key or password":                              2,
	"matrix-soccerbot could not save the accessKey to config":                      6,
	"matrix-soccerbot couldn't initialize matrix client, please check credentials": 4,
	"matrix-soccerbot couldn't sign in, please check credentials":                  5,
	"matrix-soccerbot has started.":                                                3,
	"matrix-soccerbot is missing user credentials (access-key / password)":         1,
	"matrix-soccerbot is missing user identification (homeserver / username)":      0,
	"unknown": 13,
}

var deIndex = []uint32{ // 27 elements
	0x00000000, 0x0000004a, 0x0000008b, 0x000000c4,
	0x000000e6, 0x00000146, 0x00000194, 0x000001df,
	0x0000027e, 0x00000297, 0x000002b4, 0x000002d2,
	0x000002e4, 0x000002fc, 0x00000306, 0x00000319,
	0x0000032e, 0x0000033e, 0x000003b1, 0x000003c5,
	0x000003d7, 0x000003f1, 0x0000040e, 0x0000042b,
	0x0000044c, 0x00000459, 0x000004dd,
} // Size: 132 bytes

const deData string = "" + // Size: 1245 bytes
	"\x02matrix-soccerbot fehlt die Benutzeridentifikation (homeserver / user" +
	"name)\x02matrix-soccerbot fehlen die Zugangsdaten (access-key / password" +
	")\x02Bitte gib entweder einen access-key oder ein Passwort an\x02matrix-" +
	"soccerbot wurde gestartet.\x02matrix-soccerbot konnte den Matrix-Client " +
	"nicht initialisieren, bitte Zugangsdaten überprüfen\x02matrix-soccerbot " +
	"konnte sich nicht einloggen, bitte Zugangsdaten überprüfen\x02matrix-soc" +
	"cerbot konnte den accessKey nicht in der Konfiguration speichern\x04\x00" +
	"\x02\x0a\x0a\x98\x01\x02# Spiel <font color=\x22%[5]s\x22>%[1]s</font> v" +
	"s. <font color=\x22%[6]s\x22>%[2]s</font> ***(<font color=\x22%[5]s\x22>" +
	"%[3]d</font>:<font color=\x22%[6]s\x22>%[4]d</font>)***\x04\x00\x01 \x14" +
	"\x02**Spiel beendet** -\x04\x00\x04\x0a\x0a\x0a\x0a\x15\x02*Spiel begann" +
	" %[1]s*\x04\x00\x04\x0a\x0a\x0a\x0a\x16\x02*Spiel beginnt %[1]s*\x04\x00" +
	"\x02\x0a\x0a\x0c\x02## **Tore**\x04\x00\x02\x0a\x0a\x12\x02*Noch keine T" +
	"ore*\x02unbekannt\x04\x01 \x00\x0e\x02(Strafschuss)\x04\x01 \x00\x10\x02" +
	"(Nachspielzeit)\x04\x01 \x00\x0b\x02(Eigentor)\x04\x00\x01\x0an\x02* %[1" +
	"]s - <font color=\x22%[2]s\x22>%[4]d</font>:<font color=\x22%[3]s\x22>%[" +
	"5]d</font> - Tor für %[6]s durch %[7]s%[8]s\x04\x00\x02\x0a\x0a\x0e\x02#" +
	"# Ergebnisse\x02* **Halbzeitstand\x02* **Stand nach 90 Minuten\x02* **St" +
	"and nach Nachspielzeit\x02* **Stand nach Verlängerung\x02* **Stand nach " +
	"Elfmeterschießen\x02* **Ergebnis\x02Daten bereitgestellt durch [OpenLiga" +
	"DB.de](https://www.openligadb.de) | [Quellcode](https://github.com/Unkn0" +
	"wnCat/matrix-soccerbot)"

var enIndex = []uint32{ // 27 elements
	0x00000000, 0x00000048, 0x0000008d, 0x000000bd,
	0x000000db, 0x00000128, 0x00000164, 0x0000019c,
	0x0000023a, 0x00000254, 0x0000026f, 0x0000028e,
	0x000002a1, 0x000002b6, 0x000002be, 0x000002cd,
	0x000002dd, 0x000002ed, 0x0000035d, 0x0000036e,
	0x00000382, 0x0000039e, 0x000003bd, 0x000003d7,
	0x000003f6, 0x00000401, 0x0000047c,
} // Size: 132 bytes

const enData string = "" + // Size: 1148 bytes
	"\x02matrix-soccerbot is missing user identification (homeserver / userna" +
	"me)\x02matrix-soccerbot is missing user credentials (access-key / passwo" +
	"rd)\x02Please provide either an access-key or password\x02matrix-soccerb" +
	"ot has started.\x02matrix-soccerbot couldn't initialize matrix client, p" +
	"lease check credentials\x02matrix-soccerbot couldn't sign in, please che" +
	"ck credentials\x02matrix-soccerbot could not save the accessKey to confi" +
	"g\x04\x00\x02\x0a\x0a\x97\x01\x02# Game <font color=\x22%[5]s\x22>%[1]s<" +
	"/font> vs. <font color=\x22%[6]s\x22>%[2]s</font> ***(<font color=\x22%[" +
	"5]s\x22>%[3]d</font>:<font color=\x22%[6]s\x22>%[4]d</font>)***\x04\x00" +
	"\x01 \x15\x02**Game Completed** -\x04\x00\x04\x0a\x0a\x0a\x0a\x13\x02*Ga" +
	"me began %[1]s*\x04\x00\x04\x0a\x0a\x0a\x0a\x17\x02*Game beginning %[1]s" +
	"*\x04\x00\x02\x0a\x0a\x0d\x02## **Goals**\x04\x00\x02\x0a\x0a\x0f\x02*No" +
	" goals yet*\x02unknown\x04\x01 \x00\x0a\x02(Penalty)\x04\x01 \x00\x0b" +
	"\x02(Overtime)\x04\x01 \x00\x0b\x02(Own goal)\x04\x00\x01\x0ak\x02* %[1]" +
	"s - <font color=\x22%[2]s\x22>%[4]d</font>:<font color=\x22%[3]s\x22>%[5" +
	"]d</font> - Goal for %[6]s by %[7]s%[8]s\x04\x00\x02\x0a\x0a\x0b\x02## R" +
	"esults\x02* **Halftime result\x02* **Result after 90 minutes\x02* **Resu" +
	"lt after extended Time\x02* **Result after overtime\x02* **Result after " +
	"penalty shots\x02* **Result\x02Data provided by [OpenLigaDB.de](https://" +
	"www.openligadb.de) | [Sourcecode](https://github.com/Unkn0wnCat/matrix-s" +
	"occerbot)"

	// Total table size 2657 bytes (2KiB); checksum: EBBB58E5