//Armaly Albert
package main

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type FeedStruct struct {
	Title string
	Location string
	Company string
	Date string
	Description string
	Url string
}

/*func main(){

	var list = populateStruct();

	fmt.Println(list[1].company)


	fmt.Print("NOW TESTING")
}*/
//Preconditions: Is called by main
//Postconditions: Outputs a populated list struct array with the strings parsed from the url
func populateStruct()[]*FeedStruct{

	//Creates a new feed and sets it to the stackoverflow url with the set location and distance parameters
	//https://github.com/mmcdole/gofeed#basic-usage
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://stackoverflow.com/jobs/feed?l=Bridgewater%2c+MA%2c+USA&u=Miles&d=100")

	//Creates Feedstruct array
	list := []*FeedStruct{}

	for i := 0; i<= len(feed.Items) -1; i++{
		temp := new(FeedStruct)
		//Title
		temp.Title = getTitle(feed.Items[i].Title)
		//Job Description
		temp.Description = feed.Items[i].Description
		//Job URL
		temp.Url = feed.Items[i].Link
		// Location of job
		temp.Location = getLocation(feed.Items[i].Title)
		// Job posting date
		temp.Date = feed.Items[i].Published
		//Company
		temp.Company = getCompany(feed.Items[i].Title)

		//Places parsed strings into array
		list = append(list, temp)

		fmt.Println(temp.Title," ", temp.Url," ", temp.Location," ", temp.Date," ", temp.Company," ",temp.Description)
	}
	return list
}
//Preconditions: Is called
//Postconditions: Parses string and returns the location pattern it separates everything between the (, and )
func getLocation(title string)string{

	r, _ := regexp.Compile(`(.*)\((.*[A-z]*), ([A-z]*)\)`)

	match := r.FindAllStringSubmatch(title, -1)

	return match[0][2]
}
//Preconditions: Is called
//Postconditions: Parses string and returns the job title pattern it separates everything between the (, and )
func getTitle(title string)string{
	r, _ := regexp.Compile(`(.*)\((.*[A-z]*), ([A-z]*)\)`)

	match := r.FindAllStringSubmatch(title, -1)

	return match[0][1]
}
//Preconditions: Is Called
//Postconditions: Parses the string and returns the company pattern is looks for "at" in string then takes everything up to the ( which is the beginning of the location
func getCompany(title string)string{

	r, _ := regexp.Compile(`at(.*)\(`)

	match := r.FindAllStringSubmatch(title, -1)

	return match[0][1]
}
//Preconditions: Is called by keyWordMatch
//Postconditions: Returns boolean depending on if the keyWord is in the string
func parseDescription(keyWord string, text string)bool{

	return strings.Contains(strings.ToUpper(keyWord), strings.ToUpper(text))
}
//Preconditions: User enters keywords to either search by or limit by (bool tells me which)
//Postconditions: Returns a feed with matching filter options set by user
func keyWordMatch(keyWord1 string, key1 bool, keyWord2 string, key2 bool, list []*FeedStruct)[]*FeedStruct{
	tempArray := []*FeedStruct{} //empty slice of struct potiners
	tempItem := new(FeedStruct)
	count := 0
	fmt.Println(len(list))
	fmt.Println(list[0].Title)
	fmt.Println("KEYWORD MATCH REACHED")
	//For bool: True = INCLUDE False = EXCLUDE
	for i, _ := range list{
		fmt.Println("FOR LOOP REACHED")
		if(key1 == false && key2 == false){
			if(parseDescription(keyWord1, list[i].Description) && parseDescription(keyWord2, list[i].Description)){
				fmt.Println("DOOUBLE FALSE HIT")
				tempItem = list[i]
				//fmt.Println(tempItem.Title)
				tempArray= append(tempArray, tempItem)
				count++
				fmt.Println("APPENDED")
			}
		} else if(key1 == true && key2 == false){
			if(parseDescription(keyWord1, list[i].Description) && !parseDescription(keyWord2, list[i].Description)){
				fmt.Println("TRUE FALSE HIT")
				tempItem = list[i]
				tempArray= append(tempArray, tempItem)
				count++
			}
		}else if(key1 == false && key2 == true){
			if(!parseDescription(keyWord1, list[i].Description) && parseDescription(keyWord2, list[i].Description)){
				fmt.Println("FALSE TRUE HIT")
				tempItem = list[i]
				tempArray= append(tempArray, tempItem)
				count++
			}
		}else if(key1 == true && key2 == true){
			if(parseDescription(keyWord1, list[i].Description) && parseDescription(keyWord2, list[i].Description)){
				fmt.Println("DOOUBLE TRUE HIT")
				tempItem = list[i]
				tempArray= append(tempArray, tempItem)
				count++
			}
		}else{
			fmt.Println("Somehow you reached here...")
		}
		fmt.Println(tempArray[count-1].Title)
	}

	fmt.Println(tempArray[0].Title)
	fmt.Println("KEYWORD MATCH ENDED")

	return tempArray
}

//Skeleton taken from  https://astaxie.gitbooks.io/build-web-application-with-golang/en/04.1.html

func redirectUser(w http.ResponseWriter, r *http.Request){
	http.Redirect(w ,r,"/scraper",301)
}
//Preconditions: Users has entered search paremeters
//Postconditions: Updates feed with entered search paremeters
func search(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	list := populateStruct()
	//fmt.Println("path", r.URL.Path)
	//fmt.Println("scheme", r.URL.Scheme)

	fmt.Println(r.Form)
	fmt.Println( "KEY WORD 1 ",r.FormValue("keyWord1")," KEYWORD 2 ",   r.FormValue("keyWord2"))

	fmt.Println("First button", r.Form.Get("filter1"),"2nd button", r.Form.Get("filter2"))

	button1 := true
	button2 := false

	if(r.Form.Get("filter1") == "1"){
		fmt.Println("1 TRUE HIT")
		button1 = true
	} else{
		fmt.Println("1 FALSE HIT")
		button1 = false
	}
	if(r.Form.Get("filter2") == "1"){
		fmt.Println("1 TRUE HIT")
		button2 = true
	}else{
		fmt.Println("1 FALSE HIT")
		button2 = false
	}
	fmt.Println("BEFORE SEARCH", button1, button2)

	searchBy := keyWordMatch(r.FormValue("keyWord1"), button1,r.FormValue("keyWord2"), button2, list)

	fmt.Println("SUCCESFULLY RETURNED FROM SEARCH")
	//fmt.Println(searchBy[10].Description)

	tmpl := template.Must(template.ParseFiles("search.gtpl"))

	//http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
	//fmt.Printf("List: %T: %+v\n", list, list)
	tmpl.Execute(w,searchBy)
	//http.Redirect(w,r,"/search",301)
	//})
	fmt.Println("TEMPLATE MADE")
}
func moreInfo(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("MORE INFO REACHED")


	//	list := populateStruct()

	fmt.Println("MoreInfo HIT:",r.FormValue("MoreInfo"))


}
func main() {

	list := populateStruct()

	fmt.Println(list[5].Title)
	//http.HandleFunc("/", sayhelloName) // setting router rule

	http.HandleFunc("/search", search)
	http.HandleFunc("/moreInfo", moreInfo)

	tmpl := template.Must(template.ParseFiles("scraper.gtpl"))

	fmt.Println(list[10].Title)
	http.HandleFunc("/scraper", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Println(r.Form)
		fmt.Println(r.Form)
		//fmt.Printf("List: %T: %+v\n", list, list)
		tmpl.Execute(w,list)
	})

	//http.HandleFunc("/scraper", scraper)

	err := http.ListenAndServe(":9090", nil) // setting listening port
	//log.Panic(http.ListenAndServe(":9090", nil))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	fmt.Println("ENNNNND")
	//http.ListenAndServe(":80", nil)
}
//http://127.0.0.1:9090/scraper