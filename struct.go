package main

import (
	"database/sql/driver"
	"encoding/json"
)

type goals struct {
	For     int `json:"for"`
	Against int `json:"against"`
}

type standingData struct {
	LeagueId int
	Season   int
	Type     string
	TeamId   int
	Played   int   `json:"played"`
	Win      int   `json:"win"`
	Draw     int   `json:"draw"`
	Lose     int   `json:"lose"`
	Goals    goals `json:"goals" gorm:"embedded"`
}

type standingTeam struct {
	TeamId int    `json:"id"`
	Name   string `json:"name"`
	Logo   string `json:"logo"`
}

type standing struct {
	LeagueId    int
	League      string
	Season      int
	Logo        string
	Country     string
	Rank        int          `json:"rank"`
	Team        standingTeam `json:"team" gorm:"embedded"`
	Points      int          `json:"points"`
	GoalsDiff   int          `json:"goalsDiff"`
	Group       string       `json:"group"`
	Form        string       `json:"form"`
	Status      string       `json:"status"`
	Description string       `json:"description"`
	Update      string       `json:"update"`
	All         standingData `json:"all" gorm:"-"`
	Home        standingData `json:"home" gorm:"-"`
	Away        standingData `json:"away" gorm:"-"`
}

type standingLeague struct {
	Id        int          `json:"id"`
	Name      string       `json:"name"`
	Country   string       `json:"country"`
	Logo      string       `json:"logo"`
	Flag      string       `json:"flag"`
	Season    int          `json:"season"`
	Standings [][]standing `json:"standings"`
}

type oneStandingResponse struct {
	League standingLeague `json:"league"`
}

type standingRsp struct {
	Response []oneStandingResponse `json:"response"`
}

type fixtureLeague struct {
	LeagueId   int    `json:"id"`
	LeagueName string `json:"name"`
	Season     int    `json:"season"`
	Round      string `json:"round"`
}

type fixtureBase struct {
	FixtureId int    `json:"id"`
	Date      string `json:"date"`
}

type fixtureGoals struct {
	GoalHome int `json:"home"`
	GoalAway int `json:"away"`
}

type fixtureTeamHome struct {
	HomeTeamId   int    `json:"id"`
	HomeTeamName string `json:"name"`
}

type fixtureTeamAway struct {
	AwayTeamId   int    `json:"id"`
	AwayTeamName string `json:"name"`
}

type fixtureTeams struct {
	Home fixtureTeamHome `json:"home" gorm:"embedded"`
	Away fixtureTeamAway `json:"away" gorm:"embedded"`
}

type Fixture struct {
	Fixture fixtureBase   `json:"fixture" gorm:"embedded"`
	League  fixtureLeague `json:"league" gorm:"embedded"`
	Teams   fixtureTeams  `json:"teams" gorm:"embedded"`
	Goals   fixtureGoals  `json:"goals" gorm:"embedded"`
}

type fixturesRsp struct {
	Response []Fixture `json:"response"`
}

type birth struct {
	Date    string `json:"date"`
	Place   string `json:"place"`
	Country string `json:"country"`
}

type player struct {
	TeamId      int
	TeamName    string
	LeagueId    int
	LeagueName  string
	Season      int
	PlayerId    int    `json:"id"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Height      string `json:"height"`
	Wight       string `json:"wight"`
	Nationality string `json:"nationality"`
	Injured     bool   `json:"injured"`
	Photo       string `json:"photo"`
	Birth       birth  `json:"birth" gorm:"embedded"`
}

type team struct {
	TeamId   int    `json:"id"`
	TeamName string `json:"name"`
}

type league struct {
	LeagueId   int    `json:"id"`
	LeagueName string `json:"name"`
	Season     int    `json:"season"`
}

type games struct {
	Appearences int    `json:"appearences"`
	Lineups     int    `json:"lineups"`
	Minutes     int    `json:"minutes"`
	Number      int    `json:"number"`
	Position    string `json:"position"`
	Rating      string `json:"rating"`
	Captain     bool   `json:"captain"`
}

func (g games) Value() (driver.Value, error) {
	data, _ := json.Marshal(g)
	return data, nil
}

type substitutes struct {
	In    int `json:"in"`
	Out   int `json:"out"`
	Bench int `json:"bench"`
}

func (s substitutes) Value() (driver.Value, error) {
	data, _ := json.Marshal(s)
	return data, nil
}

type shots struct {
	Total int `json:"total"`
	On    int `json:"on"`
}

func (s shots) Value() (driver.Value, error) {
	data, _ := json.Marshal(s)
	return data, nil
}

type playerGoals struct {
	Total    int `json:"total"`
	Conceded int `json:"conceded"`
	Assists  int `json:"assists"`
	Saves    int `json:"saves"`
}

func (p playerGoals) Value() (driver.Value, error) {
	data, _ := json.Marshal(p)
	return data, nil
}

type passes struct {
	Total    int `json:"total"`
	Key      int `json:"key"`
	Accuracy int `json:"accuracy"`
}

func (p passes) Value() (driver.Value, error) {
	data, _ := json.Marshal(p)
	return data, nil
}

type tackles struct {
	Total         int `json:"total"`
	Blocks        int `json:"blocks"`
	Interceptions int `json:"interceptions"`
}

func (t tackles) Value() (driver.Value, error) {
	data, _ := json.Marshal(t)
	return data, nil
}

type duels struct {
	Total int `json:"total"`
	Won   int `json:"won"`
}

func (d duels) Value() (driver.Value, error) {
	data, _ := json.Marshal(d)
	return data, nil
}

type dribbles struct {
	Attempts int `json:"attempts"`
	Success  int `json:"success"`
	Past     int `json:"past"`
}

func (d dribbles) Value() (driver.Value, error) {
	data, _ := json.Marshal(d)
	return data, nil
}

type fouls struct {
	Drawn     int `json:"drawn"`
	Committed int `json:"committed"`
}

func (f fouls) Value() (driver.Value, error) {
	data, _ := json.Marshal(f)
	return data, nil
}

type cards struct {
	Yellow    int `json:"yellow"`
	Yellowred int `json:"yellowred"`
	Red       int `json:"red"`
}

func (c cards) Value() (driver.Value, error) {
	data, _ := json.Marshal(c)
	return data, nil
}

type penalty struct {
	Won      int `json:"won"`
	Commited int `json:"commited"`
	Scored   int `json:"scored"`
	Missed   int `json:"missed"`
	Saved    int `json:"saved"`
}

func (p penalty) Value() (driver.Value, error) {
	data, _ := json.Marshal(p)
	return data, nil
}

type statistic struct {
	PlayerId    int
	Team        team        `json:"team" gorm:"embedded"`
	League      league      `json:"league" gorm:"embedded"`
	Games       games       `json:"games" gorm:"type:json"`
	Substitutes substitutes `json:"substitutes" gorm:"type:json"`
	Shots       shots       `json:"shots" gorm:"type:json"`
	Goals       playerGoals `json:"goals" gorm:"type:json"`
	Passes      passes      `json:"passes" gorm:"type:json"`
	Tackles     tackles     `json:"tackles" gorm:"type:json"`
	Duels       duels       `json:"duels" gorm:"type:json"`
	Dribbles    dribbles    `json:"dribbles" gorm:"type:json"`
	Fouls       fouls       `json:"fouls" gorm:"type:json"`
	Cards       cards       `json:"cards" gorm:"type:json"`
	Penalty     penalty     `json:"penalty" gorm:"type:json"`
}

type playerStatistics struct {
	Player     player      `json:"player"`
	Statistics []statistic `json:"statistics"`
}

type playerRsp struct {
	Response []playerStatistics `json:"response"`
}
