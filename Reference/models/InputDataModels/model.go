package InputDataModel

type AllGamesRequest struct {
	Apps []App `json:"apps"`
}

type AppDetailsRequest struct {
	AppId   uint
	Success bool       `json:"success"`
	Data    AppDetails `json:"data"`
}

type App struct {
	AppId uint   `json:"appid"`
	Name  string `json:"name"`
}

type AppDetails struct {
	AppType             string         `json:"type"`
	RequiredAge         string         `json:"required_age"`
	IsFree              bool           `json:"is_free"`
	AboutTheGame        string         `json:"about_the_game"`
	DetailedDescription string         `json:"detailed_description"`
	ShortDescription    string         `json:"short_description"`
	SupportedLanguages  string         `json:"supported_languages"`
	Developers          []string       `json:"developers"`
	Publishers          []string       `json:"publishers"`
	Packages            []int          `json:"packages"`
	Price               Prices         `json:"price_overview"`
	ReleaseDate         ReleaseDates   `json:"release_date"`
	Platform            Platforms      `json:"platforms"`
	Categories          []Category     `json:"categories"`
	Genres              []Genre        `json:"genres"`
	Recommendations     Recommendation `json:"recommendations"`
	PCRequirements      Requirements   `json:"pc_requirements"`
	Critic              CriticScore    `json:"metacritic"`
	DLC                 []int          `json:"dlc"`
}

type PlayedGame struct {
	AppId                    int  `json:"appid"`
	PlaytimeForever          int  `json:"playtime_forever"`
	PlaytimeWinForever       int  `json:"playtime_windows_forever"`
	PlaytimeMacForever       int  `json:"playtime_mac_forever"`
	PlaytimeLinuxForever     int  `json:"playtime_linux_forever"`
	RtimeLastPlayed          int  `json:"rtime_last_played"`
	Playtime2Weeks           int  `json:"playtime_2weeks"`
	HasCommunityVisibleStats bool `json:"has_community_visible_stats"`
}

type User struct {
	UserID     int    `json:"appid"`
	Name       string `json:"name"`
	GamesCount int    `json:"games_count"`
}

type Genre struct {
	GenreId     string `json:"id"`
	Description string `json:"description"`
}

type Category struct {
	CategoryId  int    `json:"id"`
	Description string `json:"description"`
}

type Prices struct {
	AppId           uint   `json:"appid"`
	Currency        string `json:"currency"`
	Initial         int    `json:"initial"`
	Final           int    `json:"final"`
	DiscountPercent int    `json:"discount_percent"`
	Country         string `json:"country"`
}

type ReleaseDates struct {
	AppId       uint   `json:"appid"`
	CommingSoon bool   `json:"comming_soon"`
	Date        string `json:"date"`
}

type Requirements struct {
	AppId       uint   `json:"appid"`
	Minimum     string `json:"minimum"`
	Recommended string `json:"recommended"`
}

type CriticScore struct {
	AppId       uint   `json:"appid"`
	CriticScore int    `json:"critic_score"`
	UsersScore  int    `json:"user_score"`
	Url         string `json:"url"`
}

type Platforms struct {
	AppId         uint `json:"appid"`
	PlatformWin   bool `json:"windows"`
	PlatformMac   bool `json:"mac"`
	PlatformLinux bool `json:"linux"`
}

type Recommendation struct {
	Total int `json:"total"`
}
