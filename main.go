package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	baseUrl      = "https://api-football-v1.p.rapidapi.com/v3/standings"
	apiKeyKey    = "X-RapidAPI-Key"
	apiKeyValue  = "59b585f5a0mshe6fac969f3271a0p16fe19jsnf484b32978bd"
	apiHostKey   = "X-RapidAPI-Host"
	apiHostValue = "api-football-v1.p.rapidapi.com"
)

func main() {
	dsn := "root:123456@tcp(localhost)/rapid_football?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	pullStanding(db, 39, 2020)
}

func pullStanding(db *gorm.DB, league int, season int) error {
	if db.Migrator().HasTable(&standing{}) == false {
		db.AutoMigrate(&standing{})
	}
	if db.Migrator().HasTable(&standingData{}) == false {
		db.AutoMigrate(&standingData{})
	}
	kvp := make(map[string]string)
	kvp["season"] = fmt.Sprintf("%d", season)
	kvp["league"] = fmt.Sprintf("%d", league)
	kvh := make(map[string]string)
	kvh[apiKeyKey] = apiKeyValue
	kvh[apiHostKey] = apiHostValue
	data, err := httpGet(baseUrl, kvp, kvh)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var rsp standingRsp
	err = json.Unmarshal(data, &rsp)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if len(rsp.Response) > 0 {
		r := rsp.Response[0]
		if len(r.League.Standings) > 0 {
			for _, t := range r.League.Standings[0] {
				t.LeagueId = r.League.Id
				t.League = r.League.Name
				t.Season = r.League.Season
				t.Country = r.League.Country
				t.Logo = r.League.Logo
				fmt.Printf("%+v", t)
				db.Create(t)
				t.All.LeagueId = r.League.Id
				t.All.Season = r.League.Season
				t.All.TeamId = t.Team.TeamId
				t.All.Type = "all"
				db.Create(t.All)
				t.Home.LeagueId = r.League.Id
				t.Home.Season = r.League.Season
				t.Home.TeamId = t.Team.TeamId
				t.Home.Type = "home"
				db.Create(t.Home)
				t.Away.LeagueId = r.League.Id
				t.Away.Season = r.League.Season
				t.Away.TeamId = t.Team.TeamId
				t.Away.Type = "away"
				db.Create(t.Away)
			}
		}
	}
	return nil
}
