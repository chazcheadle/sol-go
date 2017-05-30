# sol-go
Get Sunrise and Sunset data for current day based on Lat and Lon. Written in Golang

This program uses the NOVAS software libary, http://aa.usno.navy.mil/software/novas/novas_info.php, as implemented here: https://github.com/pebbe/novas

Example usage:
localhost:3001/sol?lat=40.758923&lon=-73.979541

Sample output:

```
{
    "sunData": {
        "time": "2017-05-29T12:54:19.380049199-04:00",
        "latitude": 40.758923,
        "longitude": -73.979541,
        "altitude": 71.104909612734,
        "azimuth": 180.5742233922943,
        "sunriseAzimuth": 59.81787301662578,
        "sunriseTime": "2017-05-30T05:27:52.55859375-04:00",
        "solarNoonAltitude": 71.10559179969913,
        "solarNoonAzimuth": 180.02379076253195,
        "solarNoonTime": "2017-05-30T12:53:33.28125-04:00",
        "sunsetAzimuth": 300.3091077486129,
        "sunsetTime": "2017-05-30T20:19:33.33984375-04:00",
        "solarMidnightAltitude": -27.30660995423206,
        "solarMidnightAzimuth": 359.9668887335456,
        "solarMidnightTime": "2017-05-31T00:53:28.0078125-04:00"
    }
}
```
