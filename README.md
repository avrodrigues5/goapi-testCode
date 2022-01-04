# goapi-testCode

Simple API application to get weather data

## Getting started

Run the application with 

```
 go run main.go
```

### Test case 1

Open PostMan or execute GET call

```
curl --location --request GET 'http://localhost:10000/api/weather'
```

You will get all the alerts to defaulted NY state in GET call

### Test case 2
Open PostMan or execute POST call

```
curl --location --request POST 'http://localhost:10000/api/weather?state=CA'
```
You will get all the alert highlight for state of California

```
{"Alert":["Winter Storm Warning issued January 3 at 3:54PM PST until January 4 at 12:00AM PST by NWS Sacramento CA","Winter Weather Advisory issued January 3 at 3:54PM PST until January 4 at 7:00AM PST by NWS Sacramento CA","Winter Storm Warning issued January 3 at 3:54PM PST until January 4 at 7:00AM PST by NWS Sacramento CA","Wind Advisory issued January 3 at 2:54PM PST until January 3 at 7:00PM PST by NWS Eureka CA","Wind Advisory issued January 3 at 2:54PM PST until January 3 at 7:00PM PST by NWS Eureka CA","Coastal Flood Advisory issued January 3 at 2:24PM PST until January 4 at 1:00PM PST by NWS Eureka CA","Wind Advisory issued January 3 at 2:14PM PST until January 4 at 4:00AM PST by NWS Sacramento CA","Coastal Flood Advisory issued January 3 at 1:57PM PST until January 4 at 3:00PM PST by NWS San Francisco CA","Coastal Flood Advisory issued January 3 at 1:57PM PST until January 4 at 3:00PM PST by NWS San Francisco CA","Special Weather Statement issued January 3 at 1:45PM PST by NWS Reno NV","Winter Weather Advisory issued January 3 at 1:16PM PST until January 4 at 4:00AM PST by NWS Medford OR","Winter Storm Warning issued January 3 at 1:16PM PST until January 4 at 4:00AM PST by NWS Medford OR","Winter Weather Advisory issued January 3 at 1:16PM PST until January 4 at 4:00AM PST by NWS Medford OR","Winter Storm Warning issued January 3 at 1:16PM PST until January 4 at 4:00AM PST by NWS Medford OR","Winter Weather Advisory issued January 3 at 1:16PM PST until January 4 at 4:00AM PST by NWS Medford OR","Winter Storm Warning issued January 3 at 1:16PM PST until January 4 at 4:00AM PST by NWS Medford OR","Flood Watch issued January 3 at 11:54AM PST until January 3 at 8:00PM PST by NWS Eureka CA","Winter Storm Warning issued January 3 at 11:50AM PST until January 4 at 4:00AM PST by NWS Eureka CA","High Wind Warning issued January 3 at 11:27AM PST until January 4 at 7:00AM PST by NWS Reno NV","High Wind Warning issued January 3 at 11:27AM PST until January 4 at 7:00AM PST by NWS Reno NV","Wind Advisory issued January 3 at 11:27AM PST until January 4 at 7:00AM PST by NWS Reno NV","Flood Watch issued January 3 at 10:52AM PST until January 5 at 7:00AM PST by NWS Medford OR","Wind Advisory issued January 2 at 11:08AM PST until January 4 at 4:00AM PST by NWS Sacramento CA","Winter Storm Warning issued January 2 at 3:53AM PST until January 4 at 4:00AM PST by NWS Eureka CA"]}
```


### Test case 3

Open PostMan or execute POST call

```
curl --location --request POST 'http://localhost:10000/api/weather?state=CAA'
```

You will get following response 

```
Please enter valid two letter state abbreviation
```