package main

import (
    "fmt"
    "log"
    "io/ioutil"
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
)

// Get only the Headlines available in the Properties
type Properties struct {
    Headline    string `json:"headline"`
}

// Get all properties in each features
type Features struct{
    Properties Properties `json:"properties"`
}

// Get all the features available in Json response body
type APIResponse struct{
    Features []Features `json:"features"`
}

// Struct to create list of headlines as alerts
type Message struct {
    Alert []string
}

// All the 50 state abbreviations
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

// Default Home page Rest Call
func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "This is Weather API client's Home Page!")
    fmt.Println("Endpoint Hit: homePage")
}

/* 
Generic API Call which differentiates either GET method or POST method 
If GET method returns alerts of State of New York 
else POST takes two letter Abbreviations of all the states in USA and 
returns the alert for that state given in parameter.
If the abbreviations isn't among the 50 , returns 400 Bad request
*/
func genericAPIcall(w http.ResponseWriter, r *http.Request){
    // Creating variables for http.response and error variable
    var response *http.Response
    var reponseErr error
    // Distinguish if method is GET or POST or any other call return 405 error code
    if r.Method == "GET" {
        response, reponseErr = http.Get("https://api.weather.gov/alerts/active?area=NY")
    } else if r.Method == "POST"{
        r.ParseForm()                     
        stateAbbr := r.Form.Get("state") 
        _, found := Find(stateAbbreviations, stateAbbr)
        if !found {
            // If state is not one of the 50 states return 400 error code
            w.WriteHeader(http.StatusBadRequest)
            fmt.Fprintf(w, "Please enter valid two letter state abbreviation")
            return
        } else{
            response, reponseErr = http.Get("https://api.weather.gov/alerts/active?area="+stateAbbr)
        }
    } else{
        // If it's not GET or POST return 405 error code stating method is not allowed
        w.WriteHeader(http.StatusMethodNotAllowed)
        fmt.Fprintf(w, "Method %v is not allowed", r.Method)
        return
    }

    if reponseErr != nil {
        // If the downstream API does not return json return 503 service unable
        w.WriteHeader(http.StatusServiceUnavailable)
        fmt.Fprintf(w, "The Weather API service Unavailable at this time")
        return
    }


    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        // If the response bodu isnt able to be parsed return status not available with error code of 404
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintf(w, "The Weather API service Response isn't available at this time, Please try again later")
        return
    }

    // Parse the API response
    var resp APIResponse
    err = json.Unmarshal(responseData, &resp)
    if err != nil {
        // If the response body is unable to be marshalled return 404 error code stating status not available
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintf(w, "UnMarshalling json data from weather API had following err: %v", err)
        return
    }

    // Get all the headlines in the slice
    alertsHeadline := []string{}
    for _,s := range resp.Features{
        alertsHeadline = append(alertsHeadline, s.Properties.Headline)
    }

    // create a json response
    out, err := json.Marshal(&Message{alertsHeadline})
    if err != nil {
        // If there's an error in Marshall give a 404 error code stating status not available
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintf(w, "Marshall json to create headline had following err: %v", err)
        return
    }
    // If everything seems fine give a success code of 200 and write the alerts as json object
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, string(out))
    fmt.Printf("Endpoint Hit: %v Call\n", r.Method)
}

/*
This was initial code 
Modification done where 
1. Making it as just one function handling the GET and POST
2. Comment all the code
3. Error code returned
*/

// func getAPIcall(w http.ResponseWriter, r *http.Request){
//     // call the REST API with New york state abbreivations
//     response, err := http.Get("https://api.weather.gov/alerts/active?area=NY")
    
//     // if err exit
//     if err != nil {
//         fmt.Print(err.Error())
//         os.Exit(1)
//     }

//     // Read the response body
//     responseData, err := ioutil.ReadAll(response.Body)
//     if err != nil {
//         log.Fatal(err)
//     }
//     // Unmarshall the json response
//     var resp APIResponse
//     err = json.Unmarshal(responseData, &resp)
//     if err != nil {
//         fmt.Printf("err was %v", err)
//     }

//     // Get all the headline as Slice
//     alertsHeadline := []string{}
//     for _,s := range resp.Features{
//         alertsHeadline = append(alertsHeadline, s.Properties.Headline)
//     }
    
//     // create alert json
//     out, err := json.Marshal(&Message{alertsHeadline})
//     if err != nil {
//         panic (err)
//     }
//     // return the alert json
//     fmt.Fprintf(w, string(out))
//     fmt.Println("Endpoint Hit: Get Call")
// }

// func postAPIcall(w http.ResponseWriter, r *http.Request){
//     //Get the state parameter
//     r.ParseForm()                     
//     stateAbbr := r.Form.Get("state") 
//     _, found := Find(stateAbbreviations, stateAbbr)
//     // If state is not found write data saying please enter two letter abbreviations
//     if !found {
//         fmt.Fprintf(w, "Please enter valid two letter state abbreviation")
//     } else{
//         // Get the response to the state
//         response, err := http.Get("https://api.weather.gov/alerts/active?area="+stateAbbr)

//         if err != nil {
//             fmt.Print(err.Error())
//             os.Exit(1)
//         }

//         // Unmarshall the response body to get all headline
//         responseData, err := ioutil.ReadAll(response.Body)
//         if err != nil {
//             log.Fatal(err)
//         }
//         var resp APIResponse
//         err = json.Unmarshal(responseData, &resp)
//         if err != nil {
//             fmt.Printf("err was %v", err)
//         }

//         // put the alerts information in an array
//         alertsHeadline := []string{}
//         for _,s := range resp.Features{
//             alertsHeadline = append(alertsHeadline, s.Properties.Headline)
//         }
        
//         out, err := json.Marshal(&Message{alertsHeadline})
//         if err != nil {
//             panic (err)
//         }
//         // return the json alerts object
//         fmt.Fprintf(w, string(out))
//     }
//     fmt.Println("Endpoint Hit: Post Call")
// }


func handleRequests() {
    /*
    Have three End points, 
    one GET which defaults to alerts of New york 
    and other POST which takes two letter Abbreviations
    */
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/api/weather",genericAPIcall).Methods("POST")
    myRouter.HandleFunc("/api/weather",genericAPIcall)
    
    log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
    /*
    call the handle request which creates various 
    end points available when program is run
    */
    handleRequests()
}