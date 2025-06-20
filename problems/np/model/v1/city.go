package v1

import (
	"math/rand"
	"time"
)

type City struct {
	Coordinates
	Name     string
	Distance float64
}

// Coordinates 经纬度
type Coordinates struct {
	Latitude  float64 `yaml:"latitude"`  //纬度
	Longitude float64 `yaml:"longitude"` //经度
}

// USCities 包含美国各州不同区域的至少 50 个城市
var USCities = []City{
	City{
		Name: "New York",

		Coordinates: Coordinates{
			Latitude:  40.7128,
			Longitude: -74.0060,
		},
		Distance: 0,
	},
	City{
		Name: "Los Angeles",

		Coordinates: Coordinates{
			Latitude:  34.0522,
			Longitude: -118.2437,
		},
		Distance: 0,
	},
	City{
		Name: "Chicago",

		Coordinates: Coordinates{
			Latitude:  41.8781,
			Longitude: -87.6298,
		},
		Distance: 0,
	},
	City{
		Name: "Houston",

		Coordinates: Coordinates{
			Latitude:  29.7604,
			Longitude: -95.3698,
		},
		Distance: 0,
	},
	City{
		Name: "Phoenix",

		Coordinates: Coordinates{
			Latitude:  33.4484,
			Longitude: -112.0740,
		},
		Distance: 0,
	},
	City{
		Name: "Philadelphia",

		Coordinates: Coordinates{
			Latitude:  39.9526,
			Longitude: -75.1652,
		},
		Distance: 0,
	},
	City{
		Name: "San Antonio",

		Coordinates: Coordinates{
			Latitude:  29.4241,
			Longitude: -98.4936,
		},
		Distance: 0,
	},
	City{
		Name: "San Diego",

		Coordinates: Coordinates{
			Latitude:  32.7157,
			Longitude: -117.1611,
		},
		Distance: 0,
	},
	City{
		Name: "Dallas",

		Coordinates: Coordinates{
			Latitude:  32.7767,
			Longitude: -96.7970,
		},
		Distance: 0,
	},
	City{
		Name: "San Jose",

		Coordinates: Coordinates{
			Latitude:  37.3382,
			Longitude: -121.8863,
		},
		Distance: 0,
	},
	City{
		Name: "Austin",

		Coordinates: Coordinates{
			Latitude:  30.2672,
			Longitude: -97.7431,
		},
		Distance: 0,
	},
	City{
		Name: "Jacksonville",

		Coordinates: Coordinates{
			Latitude:  30.3322,
			Longitude: -81.6557,
		},
		Distance: 0,
	},
	City{
		Name: "Fort Worth",

		Coordinates: Coordinates{
			Latitude:  32.7555,
			Longitude: -97.3308,
		},
		Distance: 0,
	},
	City{
		Name: "Columbus",

		Coordinates: Coordinates{
			Latitude:  39.9612,
			Longitude: -82.9988,
		},
		Distance: 0,
	},
	City{
		Name: "Charlotte",

		Coordinates: Coordinates{
			Latitude:  35.2271,
			Longitude: -80.8431,
		},
		Distance: 0,
	},
	City{
		Name: "San Francisco",

		Coordinates: Coordinates{
			Latitude:  37.7749,
			Longitude: -122.4194,
		},
		Distance: 0,
	},
	City{
		Name: "Indianapolis",

		Coordinates: Coordinates{
			Latitude:  39.7684,
			Longitude: -86.1581,
		},
		Distance: 0,
	},
	City{
		Name: "Seattle",

		Coordinates: Coordinates{
			Latitude:  47.6062,
			Longitude: -122.3321,
		},
		Distance: 0,
	},
	City{
		Name: "Denver",

		Coordinates: Coordinates{
			Latitude:  39.7392,
			Longitude: -104.9903,
		},
		Distance: 0,
	},
	City{
		Name: "Washington",

		Coordinates: Coordinates{
			Latitude:  38.9072,
			Longitude: -77.0369,
		},
		Distance: 0,
	},
	City{
		Name: "Boston",

		Coordinates: Coordinates{
			Latitude:  42.3601,
			Longitude: -71.0589,
		},
		Distance: 0,
	},
	City{
		Name: "El Paso",

		Coordinates: Coordinates{
			Latitude:  31.7619,
			Longitude: -106.4850,
		},
		Distance: 0,
	},
	City{
		Name: "Nashville",

		Coordinates: Coordinates{
			Latitude:  36.1627,
			Longitude: -86.7816,
		},
		Distance: 0,
	},
	City{
		Name: "Detroit",

		Coordinates: Coordinates{
			Latitude:  42.3314,
			Longitude: -83.0458,
		},
		Distance: 0,
	},
	City{
		Name: "Oklahoma City",

		Coordinates: Coordinates{
			Latitude:  35.4676,
			Longitude: -97.5164,
		},
		Distance: 0,
	},
	City{
		Name: "Portland",

		Coordinates: Coordinates{
			Latitude:  45.5051,
			Longitude: -122.6750,
		},
		Distance: 0,
	},
	City{
		Name: "Las Vegas",

		Coordinates: Coordinates{
			Latitude:  36.1699,
			Longitude: -115.1398,
		},
		Distance: 0,
	},
	City{
		Name: "Memphis",

		Coordinates: Coordinates{
			Latitude:  35.1495,
			Longitude: -90.0490,
		},
		Distance: 0,
	},
	City{
		Name: "Louisville",

		Coordinates: Coordinates{
			Latitude:  38.2527,
			Longitude: -85.7585,
		},
		Distance: 0,
	},
	City{
		Name: "Baltimore",

		Coordinates: Coordinates{
			Latitude:  39.2904,
			Longitude: -76.6122,
		},
		Distance: 0,
	},
	City{
		Name: "Milwaukee",

		Coordinates: Coordinates{
			Latitude:  43.0389,
			Longitude: -87.9065,
		},
		Distance: 0,
	},
	City{
		Name: "Albuquerque",

		Coordinates: Coordinates{
			Latitude:  35.0844,
			Longitude: -106.6504,
		},
		Distance: 0,
	},
	City{
		Name: "Tucson",

		Coordinates: Coordinates{
			Latitude:  32.2226,
			Longitude: -110.9747,
		},
		Distance: 0,
	},
	City{
		Name: "Fresno",

		Coordinates: Coordinates{
			Latitude:  36.7378,
			Longitude: -119.7871,
		},
		Distance: 0,
	},
	City{
		Name: "Mesa",

		Coordinates: Coordinates{
			Latitude:  33.4152,
			Longitude: -111.8315,
		},
		Distance: 0,
	},
	City{
		Name: "Sacramento",

		Coordinates: Coordinates{
			Latitude:  38.5816,
			Longitude: -121.4944,
		},
		Distance: 0,
	},
	City{
		Name: "Atlanta",

		Coordinates: Coordinates{
			Latitude:  33.7490,
			Longitude: -84.3880,
		},
		Distance: 0,
	},
	City{
		Name: "Kansas City",

		Coordinates: Coordinates{
			Latitude:  39.0997,
			Longitude: -94.5786,
		},
		Distance: 0,
	},
	City{
		Name: "Colorado Springs",

		Coordinates: Coordinates{
			Latitude:  38.8339,
			Longitude: -104.8214,
		},
		Distance: 0,
	},
	City{
		Name: "Miami",

		Coordinates: Coordinates{
			Latitude:  25.7617,
			Longitude: -80.1918,
		},
		Distance: 0,
	},
	City{
		Name: "Raleigh",

		Coordinates: Coordinates{
			Latitude:  35.7796,
			Longitude: -78.6382,
		},
		Distance: 0,
	},
	City{
		Name: "Omaha",

		Coordinates: Coordinates{
			Latitude:  41.2565,
			Longitude: -95.9345,
		},
		Distance: 0,
	},
	City{
		Name: "Long Beach",

		Coordinates: Coordinates{
			Latitude:  33.7701,
			Longitude: -118.1937,
		},
		Distance: 0,
	},
	City{
		Name: "Virginia Beach",

		Coordinates: Coordinates{
			Latitude:  36.8529,
			Longitude: -75.9780,
		},
		Distance: 0,
	},
	City{
		Name: "Oakland",

		Coordinates: Coordinates{
			Latitude:  37.8044,
			Longitude: -122.2711,
		},
		Distance: 0,
	},
	City{
		Name: "Minneapolis",

		Coordinates: Coordinates{
			Latitude:  44.9778,
			Longitude: -93.2650,
		},
		Distance: 0,
	},
	City{
		Name: "Tulsa",

		Coordinates: Coordinates{
			Latitude:  36.1539,
			Longitude: -95.9928,
		},
		Distance: 0,
	},
	City{
		Name: "Arlington",

		Coordinates: Coordinates{
			Latitude:  32.7357,
			Longitude: -97.1081,
		},
		Distance: 0,
	},
	City{
		Name: "New Orleans",

		Coordinates: Coordinates{
			Latitude:  29.9511,
			Longitude: -90.0715,
		},
		Distance: 0,
	},
	City{
		Name: "Wichita",

		Coordinates: Coordinates{
			Latitude:  37.6872,
			Longitude: -97.3301,
		},
		Distance: 0,
	},
}

// RandomUSCity 生成一个随机的美国城市示例
func RandomUSCity() City {
	// 示例城市列表（可扩展）
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return USCities[r.Intn(len(USCities))]
}

func IsInContinentalUS(lat, lon float64) bool {
	return lat >= 24.5 && lat <= 49.4 && lon >= -124.8 && lon <= -66.9
}
