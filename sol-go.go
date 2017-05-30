package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/pebbe/novas"

	"fmt"
	"time"
)

const (
	sundip = float64(-0.8)
)

// Sol data struct.
type Sol struct {
	SunData struct {
		Time                  novas.Time `json:"time"`
		Latitude              float64    `json:"latitude"`
		Longitude             float64    `json:"longitude"`
		Altitude              float64    `json:"altitude"`
		Azimuth               float64    `json:"azimuth"`
		SunriseAzimuth        float64    `json:"sunriseAzimuth"`
		SunriseTime           novas.Time `json:"sunriseTime"`
		SolarNoonAltitude     float64    `json:"solarNoonAltitude"`
		SolarNoonAzimuth      float64    `json:"solarNoonAzimuth"`
		SolarNoonTime         novas.Time `json:"solarNoonTime"`
		SunsetAzimuth         float64    `json:"sunsetAzimuth"`
		SunsetTime            novas.Time `json:"sunsetTime"`
		SolarMidnightAltitude float64    `json:"solarMidnightAltitude"`
		SolarMidnightAzimuth  float64    `json:"solarMidnightAzimuth"`
		SolarMidnightTime     novas.Time `json:"solarMidnightTime"`
	} `json:"sunData"`
}

// Retrieve data for specified location in decimal degrees.
func getSolData(latitude float64, longitude float64) *Sol {

	// Create new instance of Sol.
	sol := &Sol{}

	sol.SunData.Latitude = latitude
	sol.SunData.Longitude = longitude

	// Begin solar calculations
	now := novas.Now()
	// lat, long, elevation in meters, temperature C, pressure.
	geo := novas.NewPlace(latitude, longitude, 0, 20, 1010)

	sun := novas.Sun()
	data := sun.Topo(now, geo, novas.REFR_NONE)
	t0 := novas.Date(now.Year(), int(now.Month()), now.Day(), 0, 0, 0, 0, now.Location())
	sol.SunData.Time = novas.Now()

	// Sunrise
	sunriseTime, sunriseTopo, err := sun.Rise(t0, geo, sundip, time.Second, novas.REFR_NONE)
	if err != nil {
		fmt.Println(err)
	} else {
		t0 = sunriseTime
	}

	// Solar noon
	solarNoonTime, solarNoonTopo, err := sun.High(t0, geo, time.Second, novas.REFR_NONE)
	if err != nil {
		fmt.Println(err)
	} else {
		t0 = solarNoonTime
	}

	// Sunset
	sunsetTime, sunsetTopo, err := sun.Set(t0, geo, sundip, time.Second, novas.REFR_NONE)
	if err != nil {
		fmt.Println(err)
	} else {
		t0 = sunsetTime
	}

	// Solar midnight
	solarMidnightTime, solarMidnightTopo, err := sun.Low(t0, geo, time.Second, novas.REFR_NONE)
	if err != nil {
		fmt.Println(err)
	} else {
		t0 = solarMidnightTime
	}

	// Populate sol struct.
	sol.SunData.Altitude = data.Alt
	sol.SunData.Azimuth = data.Az
	sol.SunData.SunsetAzimuth = sunsetTopo.Az
	sol.SunData.SunsetTime = sunsetTime
	sol.SunData.SolarNoonAltitude = solarNoonTopo.Alt
	sol.SunData.SolarNoonAzimuth = solarNoonTopo.Az
	sol.SunData.SolarNoonTime = solarNoonTime
	sol.SunData.SunriseAzimuth = sunriseTopo.Az
	sol.SunData.SunriseTime = sunriseTime
	sol.SunData.SolarMidnightAltitude = solarMidnightTopo.Alt
	sol.SunData.SolarMidnightAzimuth = solarMidnightTopo.Az
	sol.SunData.SolarMidnightTime = solarMidnightTime

	return sol

}

// Encode handler for web requests
func solHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	latitude, latErr := strconv.ParseFloat(r.FormValue("lat"), 64)
	longitude, lonErr := strconv.ParseFloat(r.FormValue("lon"), 64)

	if latErr != nil || lonErr != nil {

	}

	sol := getSolData(latitude, longitude)

	// Convert Sol struct to JSON for output.
	buffer, _ := json.MarshalIndent(sol, "", "    ")

	// Send proper JSON reponse to ResponseWriter.
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(buffer)

}

func main() {

	// Instantiate a new router.
	router := httprouter.New()

	router.GET("/sol", solHandler)
	router.POST("/sol", solHandler)
	http.ListenAndServe(":3001", router)

}
