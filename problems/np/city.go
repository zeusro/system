package np

import (
	"math/rand"
	"time"
)

type City struct {
	Name        string      `yaml:"name"`
	Timezone    string      `yaml:"timezone"`
	Coordinates Coordinates `yaml:"coordinates"`
}

// Coordinates 经纬度
type Coordinates struct {
	Latitude  float64 `yaml:"latitude"`  //纬度
	Longitude float64 `yaml:"longitude"` //经度
}

// usCities 包含美国各州不同区域的至少 50 个城市
var usCities = []City{
	{"New York", "America/New_York", Coordinates{40.7128, -74.0060}},
	{"Los Angeles", "America/Los_Angeles", Coordinates{34.0522, -118.2437}},
	{"Chicago", "America/Chicago", Coordinates{41.8781, -87.6298}},
	{"Houston", "America/Chicago", Coordinates{29.7604, -95.3698}},
	{"Phoenix", "America/Phoenix", Coordinates{33.4484, -112.0740}},
	{"Philadelphia", "America/New_York", Coordinates{39.9526, -75.1652}},
	{"San Antonio", "America/Chicago", Coordinates{29.4241, -98.4936}},
	{"San Diego", "America/Los_Angeles", Coordinates{32.7157, -117.1611}},
	{"Dallas", "America/Chicago", Coordinates{32.7767, -96.7970}},
	{"San Jose", "America/Los_Angeles", Coordinates{37.3382, -121.8863}},
	{"Austin", "America/Chicago", Coordinates{30.2672, -97.7431}},
	{"Jacksonville", "America/New_York", Coordinates{30.3322, -81.6557}},
	{"Fort Worth", "America/Chicago", Coordinates{32.7555, -97.3308}},
	{"Columbus", "America/New_York", Coordinates{39.9612, -82.9988}},
	{"Charlotte", "America/New_York", Coordinates{35.2271, -80.8431}},
	{"San Francisco", "America/Los_Angeles", Coordinates{37.7749, -122.4194}},
	{"Indianapolis", "America/Indiana/Indianapolis", Coordinates{39.7684, -86.1581}},
	{"Seattle", "America/Los_Angeles", Coordinates{47.6062, -122.3321}},
	{"Denver", "America/Denver", Coordinates{39.7392, -104.9903}},
	{"Washington", "America/New_York", Coordinates{38.9072, -77.0369}},
	{"Boston", "America/New_York", Coordinates{42.3601, -71.0589}},
	{"El Paso", "America/Denver", Coordinates{31.7619, -106.4850}},
	{"Nashville", "America/Chicago", Coordinates{36.1627, -86.7816}},
	{"Detroit", "America/Detroit", Coordinates{42.3314, -83.0458}},
	{"Oklahoma City", "America/Chicago", Coordinates{35.4676, -97.5164}},
	{"Portland", "America/Los_Angeles", Coordinates{45.5051, -122.6750}},
	{"Las Vegas", "America/Los_Angeles", Coordinates{36.1699, -115.1398}},
	{"Memphis", "America/Chicago", Coordinates{35.1495, -90.0490}},
	{"Louisville", "America/Kentucky/Louisville", Coordinates{38.2527, -85.7585}},
	{"Baltimore", "America/New_York", Coordinates{39.2904, -76.6122}},
	{"Milwaukee", "America/Chicago", Coordinates{43.0389, -87.9065}},
	{"Albuquerque", "America/Denver", Coordinates{35.0844, -106.6504}},
	{"Tucson", "America/Phoenix", Coordinates{32.2226, -110.9747}},
	{"Fresno", "America/Los_Angeles", Coordinates{36.7378, -119.7871}},
	{"Mesa", "America/Phoenix", Coordinates{33.4152, -111.8315}},
	{"Sacramento", "America/Los_Angeles", Coordinates{38.5816, -121.4944}},
	{"Atlanta", "America/New_York", Coordinates{33.7490, -84.3880}},
	{"Kansas City", "America/Chicago", Coordinates{39.0997, -94.5786}},
	{"Colorado Springs", "America/Denver", Coordinates{38.8339, -104.8214}},
	{"Miami", "America/New_York", Coordinates{25.7617, -80.1918}},
	{"Raleigh", "America/New_York", Coordinates{35.7796, -78.6382}},
	{"Omaha", "America/Chicago", Coordinates{41.2565, -95.9345}},
	{"Long Beach", "America/Los_Angeles", Coordinates{33.7701, -118.1937}},
	{"Virginia Beach", "America/New_York", Coordinates{36.8529, -75.9780}},
	{"Oakland", "America/Los_Angeles", Coordinates{37.8044, -122.2711}},
	{"Minneapolis", "America/Chicago", Coordinates{44.9778, -93.2650}},
	{"Tulsa", "America/Chicago", Coordinates{36.1539, -95.9928}},
	{"Arlington", "America/Chicago", Coordinates{32.7357, -97.1081}},
	{"New Orleans", "America/Chicago", Coordinates{29.9511, -90.0715}},
	{"Wichita", "America/Chicago", Coordinates{37.6872, -97.3301}},
}

// RandomUSCity 生成一个随机的美国城市示例
func RandomUSCity() City {
	// 示例城市列表（可扩展）
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return usCities[r.Intn(len(usCities))]
}

func IsInContinentalUS(lat, lon float64) bool {
	return lat >= 24.5 && lat <= 49.4 && lon >= -124.8 && lon <= -66.9
}
