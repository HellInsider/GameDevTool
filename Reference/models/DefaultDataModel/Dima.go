package DefaultDataModel

type T struct {
	LatInterval []int `json:"latInterval"`
	LonInterval []int `json:"lonInterval"`
	SW          struct {
		LonInterval []int `json:"lonInterval"`
		LatInterval []int `json:"latInterval"`
		SW          struct {
			LonInterval []int `json:"lonInterval"`
			LatInterval []int `json:"latInterval"`
			IsEnd       bool  `json:"isEnd"`
			Locs        []struct {
				Id     int     `json:"id"`
				X      float64 `json:"x"`
				Y      float64 `json:"y"`
				Radius int     `json:"radius"`
			} `json:"locs"`
		} `json:"SW"`
		NW struct {
			LonInterval []int `json:"lonInterval"`
			LatInterval []int `json:"latInterval"`
			IsEnd       bool  `json:"isEnd"`
			Locs        []struct {
				Id     int     `json:"id"`
				X      float64 `json:"x"`
				Y      float64 `json:"y"`
				Radius int     `json:"radius"`
			} `json:"locs"`
		} `json:"NW"`
		NE struct {
			LonInterval []int `json:"lonInterval"`
			LatInterval []int `json:"latInterval"`
			IsEnd       bool  `json:"isEnd"`
			Locs        []struct {
				Id     int     `json:"id"`
				X      float64 `json:"x"`
				Y      float64 `json:"y"`
				Radius int     `json:"radius"`
			} `json:"locs"`
		} `json:"NE"`
		Center []int `json:"center"`
	} `json:"SW"`
	SE struct {
		LonInterval []int `json:"lonInterval"`
		LatInterval []int `json:"latInterval"`
		NE          struct {
			LonInterval []int `json:"lonInterval"`
			LatInterval []int `json:"latInterval"`
			IsEnd       bool  `json:"isEnd"`
			Locs        []struct {
				Id     int     `json:"id"`
				X      float64 `json:"x"`
				Y      float64 `json:"y"`
				Radius int     `json:"radius"`
			} `json:"locs"`
		} `json:"NE"`
		SE struct {
			LonInterval []int `json:"lonInterval"`
			LatInterval []int `json:"latInterval"`
			IsEnd       bool  `json:"isEnd"`
			Locs        []struct {
				Id     int     `json:"id"`
				X      float64 `json:"x"`
				Y      float64 `json:"y"`
				Radius int     `json:"radius"`
			} `json:"locs"`
		} `json:"SE"`
		NW struct {
			LonInterval []int `json:"lonInterval"`
			LatInterval []int `json:"latInterval"`
			IsEnd       bool  `json:"isEnd"`
			Locs        []struct {
				Id     int     `json:"id"`
				X      float64 `json:"x"`
				Y      float64 `json:"y"`
				Radius int     `json:"radius"`
			} `json:"locs"`
		} `json:"NW"`
		Center []int `json:"center"`
	} `json:"SE"`
	NW struct {
		LonInterval []int `json:"lonInterval"`
		LatInterval []int `json:"latInterval"`
		IsEnd       bool  `json:"isEnd"`
		Locs        []struct {
			Id     int     `json:"id"`
			X      float64 `json:"x"`
			Y      float64 `json:"y"`
			Radius int     `json:"radius"`
		} `json:"locs"`
	} `json:"NW"`
	Center []int `json:"center"`
	NE     struct {
		LonInterval []int `json:"lonInterval"`
		LatInterval []int `json:"latInterval"`
		SE          struct {
			LonInterval []int `json:"lonInterval"`
			LatInterval []int `json:"latInterval"`
			IsEnd       bool  `json:"isEnd"`
			Locs        []struct {
				Id     int     `json:"id"`
				X      float64 `json:"x"`
				Y      float64 `json:"y"`
				Radius int     `json:"radius"`
			} `json:"locs"`
		} `json:"SE"`
		SW struct {
			LonInterval []int `json:"lonInterval"`
			LatInterval []int `json:"latInterval"`
			IsEnd       bool  `json:"isEnd"`
			Locs        []struct {
				Id     int     `json:"id"`
				X      float64 `json:"x"`
				Y      float64 `json:"y"`
				Radius int     `json:"radius"`
			} `json:"locs"`
		} `json:"SW"`
		Center []int `json:"center"`
	} `json:"NE"`
}
