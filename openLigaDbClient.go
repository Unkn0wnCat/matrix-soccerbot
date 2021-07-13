package main

import (
	"net/http"
	"log"
	"encoding/json"
	"io"
	"strconv"
	"time"
)

type Group struct {
	GroupName 		string 	`json:"groupName"`
	GroupOrderID 	int 	`json:"groupOderID"`
	GroupID 		int 	`json:"groupID"`
}

type Team struct {
	TeamID			int		`json:"teamId"`
	TeamName		string	`json:"teamName"`
	ShortName		string	`json:"shortName"`
	TeamIconURL		string	`json:"teamIconUrl"`
	teamGroupName	string	`json:"teamGroupName"`
}

const (
	ResultHalftime 	= 1
	ResultEnd 		= 2
	ResultExtended	= 3
	ResultOvertime 	= 4
	ResultEleven 	= 5
)

type MatchResult struct {
	ResultID			int		`json:"resultID"`
	ResultName			string	`json:"resultName"`
	PointsTeam1			int		`json:"pointsTeam1"`
	PointsTeam2			int		`json:"pointsTeam2"`
	ResultOrderID		int		`json:"resultOrderID"`
	ResultTypeID		int		`json:"resultTypeID"`
	ResultDescription	string	`json: "resultDescription"`
}

type Goal struct {
	GoalID			int		`json:"goalID"`
	ScoreTeam1		int		`json:"scoreTeam1"`
	ScoreTeam2		int		`json:"scoreTeam2"`
	MatchMinute		int		`json:"matchMinute"`
	GoalGetterID	int		`json:"goalGetterID"`
	GoalGetterName	string	`json:"goalGetterName"`
	IsPenalty		bool	`json:"isPenalty"`
	IsOwnGoal		bool	`json:"isOwnGoal"`
	IsOvertime		bool	`json:"isOvertime"`
	Comment			string	`json:"comment"`
}

type Match struct {
	MatchID 			int 			`json:"matchID"`
	MatchDateTime 		string 			`json:"matchDateTime"`
	TimeZoneID 			string 			`json:"timeZoneID"`
	LeagueID 			int 			`json:"leagueId"`
	LeagueName 			string 			`json:"leagueName"`
	LeagueSeason 		int 			`json:"leagueSeason"`
	LeagueShortcut 		string 			`json:"leagueShortcut"`
	MatchDateTimeUTC 	string 			`json:"matchDateTimeUTC"`
	Group 				Group 			`json:"group"`
	Team1 				Team 			`json:"team1"`
	Team2 				Team 			`json:"team2"`
	LastUpdateDatetime 	string 			`json:"lastUpdateDateTime"`
	MatchIsFinished 	bool 			`json:"matchIsFinished"`
	MatchResults		[]MatchResult	`json:"matchResults"`
	Goals				[]Goal			`json:"goals"`

	Location			interface{}		`json:"location"`
	NumberOfViewers		interface{}		`json:"numberOfViewers"`
}

const baseURL = "https://api.openligadb.de"

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

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func GetMatchByID(id int) (Match) {
	body, err := sendGETRequest("/getmatchdata/"+strconv.Itoa(id))

	if err != nil {
		log.Fatal(err)
	}

	var match Match

	json.Unmarshal([]byte(body), &match)
	
	return match
}

func GetMatchesByLeague(league string) ([]Match) {
	body, err := sendGETRequest("/getmatchdata/"+league)

	if err != nil {
		log.Fatal(err)
	}

	var matches []Match

	json.Unmarshal([]byte(body), &matches)
	
	return matches
}

func ParseTime(timestr string) (time.Time, error) {
	time, err := time.Parse(time.RFC3339, timestr)
	return time, err
}