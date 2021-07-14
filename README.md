# sol-go
Get Sunrise and Sunset data for current day based on Lat and Lon. Written in Golang

This program uses the NOVAS software libary, http://aa.usno.navy.mil/software/novas/novas_info.php, as implemented here: https://github.com/pebbe/novas

Example usage:
localhost:3000/sol?lat=40.75&lon=-73.98

Sample output:

```
{
    "sunData": {
        "time": "2017-05-29T12:54:19.380049199-04:00",
        "latitude": 40.75,
        "longitude": -73.98,
        "altitude": 71.104909612734,
        "azimuth": 180.42233922943,
        "sunriseAzimuth": 59.17301662578,
        "sunriseTime": "2017-05-30T05:27:52.55859375-04:00",
        "solarNoonAltitude": 71.5591769913,
        "solarNoonAzimuth": 180.23790753195,
        "solarNoonTime": "2017-05-30T12:53:33.28125-04:00",
        "sunsetAzimuth": 300.3091077486129,
        "sunsetTime": "2017-05-30T20:19:33.33984375-04:00",
        "solarMidnightAltitude": -27.300995423206,
        "solarMidnightAzimuth": 359.96687335456,
        "solarMidnightTime": "2017-05-31T00:53:28.0078125-04:00"
    }
}
```
