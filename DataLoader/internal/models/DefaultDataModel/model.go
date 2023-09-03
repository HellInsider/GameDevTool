package DefaultDataModel

type Game struct {
	appId               int      `json:"appid"`
	name                string   `json:"name"`
	appType             string   `json:"type"`
	requiredAge         int      `json:"required_age"`
	isFree              bool     `json:"is_free"`
	dlc                 []int    `json:"dlc"`
	aboutTheGame        string   `json:"about_the_game"`
	detailedDescription string   `json:"detailed_description"`
	shortDescription    string   `json:"short_description"`
	developers          []string `json:"developers"`
	publishers          []string `json:"publishers"`
	packages            []int    `json:"packages"`
	recommendations     int      `json:"recommendations"`
}

type PlayedGame struct {
	appId                    int  `json:"appid"`
	playtimeForever          int  `json:"playtime_forever"`
	playtimeWinForever       int  `json:"playtime_windows_forever"`
	playtimeMacForever       int  `json:"playtime_mac_forever"`
	playtimeLinuxForever     int  `json:"playtime_linux_forever"`
	rtimeLastPlayed          int  `json:"rtime_last_played"`
	playtime2Weeks           int  `json:"playtime_2weeks"`
	hasCommunityVisibleStats bool `json:"has_community_visible_stats"`
}

type User struct {
	userID     int    `json:"appid"`
	name       string `json:"name"`
	gamesCount int    `json:"games_count"`
}

type Genre struct {
	genreId     int    `json:"genre_id"`
	description string `json:"description"`
}

type Category struct {
	categoryId  int    `json:"category_id"`
	description string `json:"description"`
}

type Prices struct {
	appId           int    `json:"appid"`
	currency        string `json:"currency"`
	initial         int    `json:"initial"`
	final           int    `json:"final"`
	discountPercent int    `json:"discount_percent"`
	country         string `json:"country"`
}

type ReleaseDate struct {
	appId       int  `json:"appid"`
	commingSoon bool `json:"comming_soon"`
	date        int  `json:"date"`
}

type Requirements struct {
	appId   int    `json:"appid"`
	minimum string `json:"minimum"`
}

type Metacritic struct {
	appId int    `json:"appid"`
	score int    `json:"score"`
	url   string `json:"url"`
}

type Platforms struct {
	appId         int  `json:"appid"`
	platformWin   bool `json:"platform_win"`
	platformMac   bool `json:"platform_mac"`
	platformLinux bool `json:"platform_linux"`
}

type AllGamesResponse struct {
}
