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

package openLigaDbClient

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

// Group is a OpenLigaDB-Group (aka. League)
type Group struct {
	GroupName    string `json:"groupName"`
	GroupOrderID int    `json:"groupOderID"`
	GroupID      int    `json:"groupID"`
}

// Team is a OpenLigaDB-Team
type Team struct {
	TeamID        int    `json:"teamId"`
	TeamName      string `json:"teamName"`
	ShortName     string `json:"shortName"`
	TeamIconURL   string `json:"teamIconUrl"`
	TeamGroupName string `json:"teamGroupName"`
}

// Result types
const (
	ResultHalftime = 1 // ResultHalftime is the points as of half-time-break
	ResultEnd      = 2 // ResultEnd is the result at the end of the regular 90 minutes
	ResultExtended = 3 // ResultExtended is the result after time extensions (this happens if the match had delays)
	ResultOvertime = 4 // ResultOvertime is the result after overtime (which happens when there is no winner after extended time)
	ResultEleven   = 5 // ResultEleven is the result after 11-meter-shooting (which happens when there is no winner after overtime)
)

// MatchResult is a result entry
type MatchResult struct {
	ResultID          int    `json:"resultID"`
	ResultName        string `json:"resultName"`
	PointsTeam1       int    `json:"pointsTeam1"`
	PointsTeam2       int    `json:"pointsTeam2"`
	ResultOrderID     int    `json:"resultOrderID"`
	ResultTypeID      int    `json:"resultTypeID"` // See ResultHalftime, ResultEnd, ResultExtended, ResultOvertime and ResultEleven
	ResultDescription string `json:"resultDescription"`
}

// Goal is a goal entry
type Goal struct {
	GoalID         int    `json:"goalID"`
	ScoreTeam1     int    `json:"scoreTeam1"` // ScoreTeam1 is team1's score after the goal
	ScoreTeam2     int    `json:"scoreTeam2"` // ScoreTeam2 is team2's score after the goal
	MatchMinute    int    `json:"matchMinute"`
	GoalGetterID   int    `json:"goalGetterID"`
	GoalGetterName string `json:"goalGetterName"`
	IsPenalty      bool   `json:"isPenalty"`  // IsPenalty is true if the goal was scored on a penalty shot
	IsOwnGoal      bool   `json:"isOwnGoal"`  // IsOwnGoal is true if the goal was shot by a player on their own team's goal
	IsOvertime     bool   `json:"isOvertime"` // IsOvertime is true if this goal was scored during overtime
	Comment        string `json:"comment"`
}

// Match combines all data on a match. The match may be planned, ongoing or finished
type Match struct {
	MatchID            int           `json:"matchID"`
	MatchDateTime      string        `json:"matchDateTime"`
	TimeZoneID         string        `json:"timeZoneID"`
	LeagueID           int           `json:"leagueId"`
	LeagueName         string        `json:"leagueName"`
	LeagueSeason       int           `json:"leagueSeason"`
	LeagueShortcut     string        `json:"leagueShortcut"`
	MatchDateTimeUTC   string        `json:"matchDateTimeUTC"`
	Group              Group         `json:"group"`
	Team1              Team          `json:"team1"`
	Team2              Team          `json:"team2"`
	LastUpdateDatetime string        `json:"lastUpdateDateTime"`
	MatchIsFinished    bool          `json:"matchIsFinished"`
	MatchResults       []MatchResult `json:"matchResults"`
	Goals              []Goal        `json:"goals"`

	Location        interface{} `json:"location"`
	NumberOfViewers interface{} `json:"numberOfViewers"`
}

// baseURL for OpenLigaDB API
const baseURL = "https://api.openligadb.de"

// sendGETRequest is a helper function for sending a GET-request to an endpoint and returning the response
func sendGETRequest(endpoint string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", baseURL+endpoint, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", `application/json`)
	req.Header.Add("User-Agent", `Matrix-Soccerbot/1.0 (+https://github.com/Unkn0wnCat/matrix-soccerbot)`)

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

// GetMatchByID gets a specific match
func GetMatchByID(id int) (*Match, error) {
	body, err := sendGETRequest("/getmatchdata/" + strconv.Itoa(id))

	if err != nil {
		return nil, err
	}

	var match Match

	err = json.Unmarshal([]byte(body), &match)
	if err != nil {
		return nil, err
	}

	return &match, err
}

// GetMatchesByLeague finds all matches in a given league
func GetMatchesByLeague(league string) ([]Match, error) {
	body, err := sendGETRequest("/getmatchdata/" + league)

	if err != nil {
		return nil, err
	}

	var matches []Match

	err = json.Unmarshal([]byte(body), &matches)
	if err != nil {
		return nil, err
	}

	return matches, nil
}

// ParseTime parses the timestamps from the API into a time.Time
func ParseTime(timeStr string) (time.Time, error) {
	myTime, err := time.Parse(time.RFC3339, timeStr)
	return myTime, err
}
