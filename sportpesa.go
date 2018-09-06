package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"io/ioutil"
	"net/url"
	"encoding/json"
)

type Game struct {
	Id          int64
	League      string
	DateTime    string
	Teams       []string
	HomeWin     float64
	Draw        float64
	VisitingWin float64
}

type Result struct {
	Game_id int
	Sms_id int
	Start_date int
	Finish_date int
	Team1 string
	Team2 string
	Result string
	Team1_id int
	Team2_id int
	League string
	Country string
	League_id int
	Country_id int
	Sport_id int
	Sport_name string
	Ison_name string
	League_pos int
	Ttl int
	Datdate int
	Page_order int
}

var (
	games      []Game
	totalGames int
)

func main() {
	url := "https://www.sportpesa.co.ke/sportgames?sportId=1"
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer req.Body.Close()

	doc, err := goquery.NewDocumentFromReader(req.Body)

	if err != nil {
		log.Fatal(err)
	}

	parseDoc(doc)

	for _, game := range games {
		fmt.Println(game)
	}

	getResults()
}

func parseDoc(doc *goquery.Document) {
	doc.Find(".bp").Each(func(i int, section *goquery.Selection) {
		section.Find(".bp-events").Each(func(j int, innerSection *goquery.Selection) {
			innerSection.Find(".match.football.FOOTBALL").Each(func(index int, item *goquery.Selection) {
				// Extract league name
				leagueName := item.Find(".leaguename")
				league := leagueName.Find(".name").Text()
				//fmt.Printf("League :: %s\n", strings.TrimSpace(league))

				// Metadata
				metaData := item.Find(".meta")
				date, _ := metaData.Find(".date").Find("timeComponent").Attr("datetime")
				gameId := metaData.Find(".game-id").Text()
				gameId = gameId[8:]
				id, _ := strconv.ParseInt(gameId, 10, 32)
				//fmt.Printf("Date ::%s\n", date)
				//fmt.Printf("%d\n", id)

				// Teams
				teams := make([]string, 2)
				item.Find(".teams").Find("li").Each(func(idx int, team *goquery.Selection) {
					teams[idx] = team.Text()
				})
				//fmt.Println(teams)

				// Odds
				bets := item.Find(".bet-selector")
				pick01 := bets.Find(".pick01")
				pick0x := bets.Find(".pick0X")
				pick02 := bets.Find(".pick02")

				hw, _ := strconv.ParseFloat(pick01.Find(".odd").Text(), 64)
				draw, _ := strconv.ParseFloat(pick0x.Find(".odd").Text(), 64)
				vw, _ := strconv.ParseFloat(pick02.Find(".odd").Text(), 64)

				//fmt.Printf("HW :: %v D :: %v VW :: %v\n", hw, draw, vw)

				game := Game{
					id,
					strings.TrimSpace(league),
					strings.Trim(date, "'"),
					teams,
					hw,
					draw,
					vw,
				}

				games = append(games, game)
			})
		})
	})
}

func getResults()  {
	api := "https://www.sportpesa.co.ke/api/results/search"
	res,  err := http.PostForm(api, url.Values{
		"date": { "1532206800" },
	})
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	results, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var gamesResults []Result
	err = json.Unmarshal([]byte(string(results)), &gamesResults)
	if err != nil {
		log.Fatal(err)
	}

	for _, result := range gamesResults {
		fmt.Printf("%v\n", result)
	}
	fmt.Println(len(gamesResults))

}