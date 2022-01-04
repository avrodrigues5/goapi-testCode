package main

import (
    "fmt"
    "log"
    "io/ioutil"
    "net/http"
    "github.com/gorilla/mux"
    "os"
    "encoding/json"
)

type Properties struct {
    Headline    string `json:"headline"`
}

type Features struct{
    Properties Properties `json:"properties"`
}

type APIResponse struct{
    Features []Features `json:"features"`
}

type Message struct {
    Alert []string
}


var stateAbbreviations = []string{"AL","AK","AS","AZ","AR","CA","CO","CT",
"DE","DC","FM","FL","GA","GU","HI","ID","IL","IN","IA","KS","KY","LA",
"ME","MH","MD","MA","MI","MN","MS","MO","MT","NE","NV","NH","NJ","NM",
"NY","NC","ND","MP","OH","OK","OR","PW","PA","PR","RI","SC","SD","TN",
"TX","UT","VT","VI","VA","WA","WV","WI","WY"}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []string, val string) (int, bool) {
    for i, item := range slice {
        if item == val {
            return i, true
        }
    }
    return -1, false
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "This is Weather API client's Home Page!")
    fmt.Println("Endpoint Hit: homePage")
}

func getAPIcall(w http.ResponseWriter, r *http.Request){
    response, err := http.Get("https://api.weather.gov/alerts/active?area=NY")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    var resp APIResponse
    err = json.Unmarshal(responseData, &resp)
    if err != nil {
        fmt.Printf("err was %v", err)
    }

    alertsHeadline := []string{}
    for _,s := range resp.Features{
        alertsHeadline = append(alertsHeadline, s.Properties.Headline)
    }
    
    out, err := json.Marshal(&Message{alertsHeadline})
    if err != nil {
        panic (err)
    }
    fmt.Fprintf(w, string(out))
    fmt.Println("Endpoint Hit: Get Call")
}

func postAPIcall(w http.ResponseWriter, r *http.Request){
    r.ParseForm()                     
    stateAbbr := r.Form.Get("state") 
    _, found := Find(stateAbbreviations, stateAbbr)
    if !found {
        fmt.Fprintf(w, "Please enter valid two letter state abbreviation")
    } else{
        response, err := http.Get("https://api.weather.gov/alerts/active?area="+stateAbbr)

        if err != nil {
            fmt.Print(err.Error())
            os.Exit(1)
        }


        responseData, err := ioutil.ReadAll(response.Body)
        if err != nil {
            log.Fatal(err)
        }
        var resp APIResponse
        err = json.Unmarshal(responseData, &resp)
        if err != nil {
            fmt.Printf("err was %v", err)
        }

        alertsHeadline := []string{}
        for _,s := range resp.Features{
            alertsHeadline = append(alertsHeadline, s.Properties.Headline)
        }
        
        out, err := json.Marshal(&Message{alertsHeadline})
        if err != nil {
            panic (err)
        }
        fmt.Fprintf(w, string(out))
    }
    fmt.Println("Endpoint Hit: Post Call")
}


func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/api/weather",postAPIcall).Methods("POST")
    myRouter.HandleFunc("/api/weather",getAPIcall)
    
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
    handleRequests()
}