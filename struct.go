package main

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
