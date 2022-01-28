/*
 * Copyright © 2022 Kevin Kandlbinder.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package messageCreator

import (
	"fmt"
	"github.com/Unkn0wnCat/matrix-soccerbot/internal/openLigaDbClient"
	"github.com/gomarkdown/markdown"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"log"
	"strconv"
	"time"
)

func GenerateMessageForMatch(targetLang string, match openLigaDbClient.Match) string {
	lang := language.MustParse(targetLang)

	p := message.NewPrinter(lang)

	out := ""

	matchTime, err := openLigaDbClient.ParseTime(match.MatchDateTimeUTC)

	if err != nil {
		log.Fatal(err)
		return ""
	}

	latestGoal := openLigaDbClient.Goal{}

	if len(match.Goals) != 0 {
		latestGoal = match.Goals[len(match.Goals)-1]
	}

	/*isTeam1Winning := false
	isTeam2Winning := false*/

	team1Color := "white"
	team2Color := "white"

	if latestGoal.ScoreTeam1 > latestGoal.ScoreTeam2 {
		//isTeam1Winning = true
		team1Color = "green"
		team2Color = "red"
	}

	if latestGoal.ScoreTeam1 < latestGoal.ScoreTeam2 {
		//isTeam2Winning = true
		team2Color = "green"
		team1Color = "red"
	}

	out += p.Sprintf("# Game <font color=\"%[5]s\">%[1]s</font> vs. <font color=\"%[6]s\">%[2]s</font> ***(<font color=\"%[5]s\">%[3]d</font>:<font color=\"%[6]s\">%[4]d</font>)***\n\n", match.Team1.TeamName, match.Team2.TeamName, latestGoal.ScoreTeam1, latestGoal.ScoreTeam2, team1Color, team2Color)

	if match.MatchIsFinished {
		out += p.Sprintf("**Game Completed** - ")
	}

	if matchTime.Before(time.Now()) {
		out += p.Sprintf("*Game began %s*\n\n\n\n", matchTime.Format(time.RFC850))
	} else {
		out += p.Sprintf("*Game beginning %s*\n\n\n\n", matchTime.Format(time.RFC850))
	}

	out += p.Sprintf("## **Goals**\n\n")

	if len(match.Goals) == 0 {
		out += p.Sprintf("*No goals yet*\n\n")
	}

	prevScore1 := 0
	prevScore2 := 0
	lastTs := 0

	for _, goal := range match.Goals {
		scoringTeam := p.Sprintf("unknown")

		if prevScore1 < goal.ScoreTeam1 {
			scoringTeam = "<font color=\"" + team1Color + "\">" + match.Team1.TeamName + "</font>"
		}
		if prevScore2 < goal.ScoreTeam2 {
			scoringTeam = "<font color=\"" + team2Color + "\">" + match.Team2.TeamName + "</font>"
		}

		scoringTs := strconv.Itoa(goal.MatchMinute) + "m"

		if goal.MatchMinute == 0 && lastTs > 0 {
			scoringTs = "???"
		} else {
			lastTs = goal.MatchMinute
		}

		bonusInfos := ""

		if goal.IsPenalty {
			bonusInfos += p.Sprintf(" (Penalty)")
		}

		if goal.IsOvertime {
			bonusInfos += p.Sprintf(" (Overtime)")
		}

		if goal.IsOwnGoal {
			bonusInfos += p.Sprintf(" (Own goal)")
		}

		if goal.Comment != "" {
			bonusInfos += " \"" + goal.Comment + "\""
		}

		out += p.Sprintf("* %[1]s - <font color=\"%[2]s\">%[4]d</font>:<font color=\"%[3]s\">%[5]d</font> - Goal for %[6]s by %[7]s%[8]s\n", scoringTs, team1Color, team2Color, goal.ScoreTeam1, goal.ScoreTeam2, scoringTeam, goal.GoalGetterName, bonusInfos)

		prevScore1 = goal.ScoreTeam1
		prevScore2 = goal.ScoreTeam2
	}

	out += "\n\n\n"

	if len(match.MatchResults) > 0 {
		out += p.Sprintf("## Results\n\n")

		for _, result := range match.MatchResults {
			switch result.ResultTypeID {
			case openLigaDbClient.ResultHalftime:
				out += p.Sprintf("* **Halftime result")
				break
			case openLigaDbClient.ResultEnd:
				out += p.Sprintf("* **Result after 90 minutes")
				break
			case openLigaDbClient.ResultExtended:
				out += p.Sprintf("* **Result after extended Time")
				break
			case openLigaDbClient.ResultOvertime:
				out += p.Sprintf("* **Result after overtime")
				break
			case openLigaDbClient.ResultEleven:
				out += p.Sprintf("* **Result after penalty shots")
				break
			default:
				out += p.Sprintf("* **Result")
				break
			}

			out += " <font color=\"" + team1Color + "\">" + strconv.Itoa(result.PointsTeam1) + "</font>:<font color=\"" + team2Color + "\">" + strconv.Itoa(result.PointsTeam2) + "</font>**\n\n"
		}

		out += "---\n\n"
		out += p.Sprintf("Data provided by [OpenLigaDB.de](https://www.openligadb.de) | [Sourcecode](https://github.com/Unkn0wnCat/matrix-soccerbot)")
	}

	fmt.Println(out)
	return out
}

func RenderMarkdown(md string) string {
	html := markdown.ToHTML([]byte(md), nil, nil)
	return string(html)
}
