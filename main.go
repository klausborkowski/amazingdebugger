package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type RuleSet struct {
	Id     string
	Expect string
}
type Signal struct {
	Id    string
	Name  string
	Value string
}

var RuleSets []RuleSet = []RuleSet{
	{Id: "1", Expect: "green"},
	{Id: "2", Expect: "red"},
	{Id: "3", Expect: "5"},
}

func main() {
	fmt.Println("Started on localhost:8000")

	//hander function vs resource
	http.HandleFunc("/", root)
	http.HandleFunc("/rule-sets/", matchSignalToRuleSets)

	log.Fatal(http.ListenAndServe(":8000", nil))

}

func root(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	initRuleSets := []RuleSet{}
	rulesetInit := map[string][]RuleSet{
		"RuleSets": initRuleSets,
	}
	tmpl.Execute(w, rulesetInit)
}

func matchSignalToRuleSets(w http.ResponseWriter, r *http.Request) {
	signalId := r.PostFormValue("signalid")

	matchedRS := []RuleSet{}
	for _, rs := range RuleSets {
		if strings.Contains(signalId, rs.Expect) {
			matchedRS = append(matchedRS, rs)
		}
	}
	fmt.Println("Matched RS:", matchedRS)
	matchedRuleSetsBlock := map[string][]RuleSet{
		"RuleSets": matchedRS,
	}
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.ExecuteTemplate(w, "rulesets-list-element", matchedRuleSetsBlock)
}
