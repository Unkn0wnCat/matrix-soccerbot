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
	" (Overtime)": 24,
	" (Own goal)": 25,
	" (Penalty)":  23,
	"# Game <font color=\"%[5]s\">%[1]s</font> vs. <font color=\"%[6]s\">%[2]s</font> ***(<font color=\"%[5]s\">%[3]d</font>:<font color=\"%[6]s\">%[4]d</font>)***\n\n": 16,
	"## **Goals**\n\n": 20,
	"## Results\n\n":   27,
	"* %[1]s - <font color=\"%[2]s\">%[4]d</font>:<font color=\"%[3]s\">%[5]d</font> - Goal for %[6]s by %[7]s%[8]s\n": 26,
	"* **Halftime result":                                28,
	"* **Result":                                         33,
	"* **Result after 90 minutes":                        29,
	"* **Result after extended Time":                     30,
	"* **Result after overtime":                          31,
	"* **Result after penalty shots":                     32,
	"**Game Completed** - ":                              17,
	"*Game began %s*\n\n\n\n":                            18,
	"*Game beginning %s*\n\n\n\n":                        19,
	"*No goals yet*\n\n":                                 21,
	"Configuration error: Could not save configuration!": 15,
	"Corrupted configuration: Could not load room configurations!\n%v":                                                           14,
	"Could not accept invite to %s due to internal error":                                                                        0,
	"Could not accept invite to %s due to join error":                                                                            1,
	"Data provided by [OpenLigaDB.de](https://www.openligadb.de) | [Sourcecode](https://github.com/Unkn0wnCat/matrix-soccerbot)": 34,
	"Goodbye!": 13,
	"Please provide either an access-key or password": 5,
	"Shutting down...":            12,
	"Successfully joined room %s": 2,
	"matrix-soccerbot could not read joined rooms, something is horribly wrong":    11,
	"matrix-soccerbot could not save the accessKey to config":                      9,
	"matrix-soccerbot couldn't initialize matrix client, please check credentials": 7,
	"matrix-soccerbot couldn't sign in, please check credentials":                  8,
	"matrix-soccerbot has encountered a fatal error whilst syncing":                10,
	"matrix-soccerbot has started.":                                                6,
	"matrix-soccerbot is missing user credentials (access-key / password)":         4,
	"matrix-soccerbot is missing user identification (homeserver / username)":      3,
	"unknown": 22,
}

var deIndex = []uint32{ // 36 elements
	// Entry 0 - 1F
	0x00000000, 0x00000046, 0x0000008c, 0x000000ac,
	0x000000f6, 0x00000137, 0x00000170, 0x00000192,
	0x000001f2, 0x00000240, 0x0000028b, 0x000002d3,
	0x00000334, 0x00000347, 0x00000358, 0x000003a0,
	0x000003dc, 0x0000047b, 0x00000494, 0x000004b1,
	0x000004cf, 0x000004e1, 0x000004f9, 0x00000503,
	0x00000516, 0x0000052b, 0x0000053b, 0x000005ae,
	0x000005c2, 0x000005d4, 0x000005ee, 0x0000060b,
	// Entry 20 - 3F
	0x00000628, 0x00000649, 0x00000656, 0x000006da,
} // Size: 168 bytes

const deData string = "" + // Size: 1754 bytes
	"\x02Konnte Einladung zu %[1]s wegen eines internen Fehlers nicht annehme" +
	"n\x02Konnte Einladung zu %[1]s wegen eines Beitrittsfehlers nicht annehm" +
	"en\x02Raum %[1]s erfolgreich betreten\x02matrix-soccerbot fehlt die Benu" +
	"tzeridentifikation (homeserver / username)\x02matrix-soccerbot fehlen di" +
	"e Zugangsdaten (access-key / password)\x02Bitte gib entweder einen acces" +
	"s-key oder ein Passwort an\x02matrix-soccerbot wurde gestartet.\x02matri" +
	"x-soccerbot konnte den Matrix-Client nicht initialisieren, bitte Zugangs" +
	"daten ??berpr??fen\x02matrix-soccerbot konnte sich nicht einloggen, bitte " +
	"Zugangsdaten ??berpr??fen\x02matrix-soccerbot konnte den accessKey nicht i" +
	"n der Konfiguration speichern\x02matrix-soccerbot hat einen fatalen Fehl" +
	"er beim Synchronisieren erfahren\x02matrix-soccerbot konnte nicht die be" +
	"igetretenen R??ume abrufen, irgendwas geht grauenvoll schief\x02Fahre her" +
	"runter...\x02Auf Wiedersehen!\x02Korrumpierte Konfiguration: Konnte Raum" +
	"konfiguration nicht laden!\x0a%[1]v\x02Konfigurationsfehler: Konnte Konf" +
	"iguration nicht speichern!\x04\x00\x02\x0a\x0a\x98\x01\x02# Spiel <font " +
	"color=\x22%[5]s\x22>%[1]s</font> vs. <font color=\x22%[6]s\x22>%[2]s</fo" +
	"nt> ***(<font color=\x22%[5]s\x22>%[3]d</font>:<font color=\x22%[6]s\x22" +
	">%[4]d</font>)***\x04\x00\x01 \x14\x02**Spiel beendet** -\x04\x00\x04" +
	"\x0a\x0a\x0a\x0a\x15\x02*Spiel begann %[1]s*\x04\x00\x04\x0a\x0a\x0a\x0a" +
	"\x16\x02*Spiel beginnt %[1]s*\x04\x00\x02\x0a\x0a\x0c\x02## **Tore**\x04" +
	"\x00\x02\x0a\x0a\x12\x02*Noch keine Tore*\x02unbekannt\x04\x01 \x00\x0e" +
	"\x02(Strafschuss)\x04\x01 \x00\x10\x02(Nachspielzeit)\x04\x01 \x00\x0b" +
	"\x02(Eigentor)\x04\x00\x01\x0an\x02* %[1]s - <font color=\x22%[2]s\x22>%" +
	"[4]d</font>:<font color=\x22%[3]s\x22>%[5]d</font> - Tor f??r %[6]s durch" +
	" %[7]s%[8]s\x04\x00\x02\x0a\x0a\x0e\x02## Ergebnisse\x02* **Halbzeitstan" +
	"d\x02* **Stand nach 90 Minuten\x02* **Stand nach Nachspielzeit\x02* **St" +
	"and nach Verl??ngerung\x02* **Stand nach Elfmeterschie??en\x02* **Ergebnis" +
	"\x02Daten bereitgestellt durch [OpenLigaDB.de](https://www.openligadb.de" +
	") | [Quellcode](https://github.com/Unkn0wnCat/matrix-soccerbot)"

var enIndex = []uint32{ // 36 elements
	// Entry 0 - 1F
	0x00000000, 0x00000037, 0x0000006a, 0x00000089,
	0x000000d1, 0x00000116, 0x00000146, 0x00000164,
	0x000001b1, 0x000001ed, 0x00000225, 0x00000263,
	0x000002ad, 0x000002be, 0x000002c7, 0x0000030a,
	0x0000033d, 0x000003db, 0x000003f5, 0x00000410,
	0x0000042f, 0x00000442, 0x00000457, 0x0000045f,
	0x0000046e, 0x0000047e, 0x0000048e, 0x000004fe,
	0x0000050f, 0x00000523, 0x0000053f, 0x0000055e,
	// Entry 20 - 3F
	0x00000578, 0x00000597, 0x000005a2, 0x0000061d,
} // Size: 168 bytes

const enData string = "" + // Size: 1565 bytes
	"\x02Could not accept invite to %[1]s due to internal error\x02Could not " +
	"accept invite to %[1]s due to join error\x02Successfully joined room %[1" +
	"]s\x02matrix-soccerbot is missing user identification (homeserver / user" +
	"name)\x02matrix-soccerbot is missing user credentials (access-key / pass" +
	"word)\x02Please provide either an access-key or password\x02matrix-socce" +
	"rbot has started.\x02matrix-soccerbot couldn't initialize matrix client," +
	" please check credentials\x02matrix-soccerbot couldn't sign in, please c" +
	"heck credentials\x02matrix-soccerbot could not save the accessKey to con" +
	"fig\x02matrix-soccerbot has encountered a fatal error whilst syncing\x02" +
	"matrix-soccerbot could not read joined rooms, something is horribly wron" +
	"g\x02Shutting down...\x02Goodbye!\x02Corrupted configuration: Could not " +
	"load room configurations!\x0a%[1]v\x02Configuration error: Could not sav" +
	"e configuration!\x04\x00\x02\x0a\x0a\x97\x01\x02# Game <font color=\x22%" +
	"[5]s\x22>%[1]s</font> vs. <font color=\x22%[6]s\x22>%[2]s</font> ***(<fo" +
	"nt color=\x22%[5]s\x22>%[3]d</font>:<font color=\x22%[6]s\x22>%[4]d</fon" +
	"t>)***\x04\x00\x01 \x15\x02**Game Completed** -\x04\x00\x04\x0a\x0a\x0a" +
	"\x0a\x13\x02*Game began %[1]s*\x04\x00\x04\x0a\x0a\x0a\x0a\x17\x02*Game " +
	"beginning %[1]s*\x04\x00\x02\x0a\x0a\x0d\x02## **Goals**\x04\x00\x02\x0a" +
	"\x0a\x0f\x02*No goals yet*\x02unknown\x04\x01 \x00\x0a\x02(Penalty)\x04" +
	"\x01 \x00\x0b\x02(Overtime)\x04\x01 \x00\x0b\x02(Own goal)\x04\x00\x01" +
	"\x0ak\x02* %[1]s - <font color=\x22%[2]s\x22>%[4]d</font>:<font color=" +
	"\x22%[3]s\x22>%[5]d</font> - Goal for %[6]s by %[7]s%[8]s\x04\x00\x02" +
	"\x0a\x0a\x0b\x02## Results\x02* **Halftime result\x02* **Result after 90" +
	" minutes\x02* **Result after extended Time\x02* **Result after overtime" +
	"\x02* **Result after penalty shots\x02* **Result\x02Data provided by [Op" +
	"enLigaDB.de](https://www.openligadb.de) | [Sourcecode](https://github.co" +
	"m/Unkn0wnCat/matrix-soccerbot)"

	// Total table size 3655 bytes (3KiB); checksum: 4BEE325
