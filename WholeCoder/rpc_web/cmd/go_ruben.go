package cmd

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"sort"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Person struct {
	Name       string `json:"Name"`
	Age        int    `json:"Age"`
	Occupation string `json:"Occupation"`
}

type hackerLink struct {
	Name string
	Path string
	Type string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	hackerLinks := []hackerLink{
		{"Snow White", "Snow White", grpcCall("Snow White").Name},
		{"Focal Recursion", "FocalRecursion", grpcCall("FocalRecursion").Name},
		{"Jsldub", "Jsldub", grpcCall("Jsldub").Name},
		{"Sandwich Man", "SandwichMan", grpcCall("SandwichMan").Name},
		{"Rube5", "Rube5", grpcCall("Rube5").Name},
		{"Paper Naffin", "PaperNaffin", grpcCall("PaperNaffin").Name},
		{"Boarder Mimi", "BoarderMimi", grpcCall("BoarderMimi").Name},
		{"Alpha Cipher", "Alpha Cipher", grpcCall("Alpha Cipher").Name},
		{"Injection", "Injection", grpcCall("Injection").Name},
		{"Queen", "Queen", grpcCall("Queen").Name},
	}

	// Read the contents of a file into a byte slice
	data, err := ioutil.ReadFile("template.html")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert the byte slice to a string and print it
	homePageTemplate := string(data)
	tmpl, err :=
		template.New("home").Parse(homePageTemplate)
	if err != nil {
		log.Printf("Error parsing home page template: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, hackerLinks)
	if err != nil {
		log.Printf("Error executing home page template: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	fmt.Println("Endpoint: /")
	PrintIt("Rubes is awesome!")
}

func printToStringAnAnchor(name, nameForPath string) string {
	return "<a href=\"\\hacker\\" + nameForPath + "\">" + name + "</a>"
}

func returnAllhackers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in returnAllhackers =", persons)
	json.NewEncoder(w).Encode(persons)
}

func PrintIt[T any](a T) {
	fmt.Println(a)
}
func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/hackers", returnAllhackers)
	router.HandleFunc("/hacker/{name}", returnHacker)
	router.HandleFunc("/computer/{name}", returnComputer)
	log.Fatal(http.ListenAndServe(":10001", router))
}

func returnComputer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["name"]

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":10000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect:  %s", err)
	}

	defer conn.Close()

	c := NewComputerServiceClient(conn)

	response, err := c.SayHello(context.Background(), &Computer{Body: key})
	if err != nil {
		log.Fatalf("Error when calling SayHello:  %s", err)
	}

	log.Printf("Response from gRPC server:  %s", response.Body)

	fmt.Fprintf(w, response.Body)
}

func grpcCall(name string) *Hacker {
	conn, err := grpc.Dial(":10000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := NewHackerServiceClient(conn)

	req := &Hacker{Name: name}
	resp, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func returnHacker(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["name"]

	response := grpcCall(key)

	fmt.Fprintf(w, "<html><body>")
	fmt.Fprintf(w, response.Name)
	fmt.Fprintf(w, "<br /><br /><a href=\"/\">Back</a>")
	fmt.Fprintf(w, "</body></html>")
}

func (p *Person) changeAgeTo(age int) {
	p.Age = age
}

type Employee interface {
	printName()
	getName() string
}

type DeservesLove interface {
	printDeservesLoveName()
}

func (p *Person) printDeservesLoveName() {
	fmt.Println("Deserves Love =", p.Name)
}

func (p *Person) printName() {
	fmt.Println(p.Name)
}

func (p *Person) getName() string {
	return p.Name
}

var persons []Person

func mainRunWebServer() {
	persons = []Person{}

	persons = append(persons, Person{Name: "Ruben", Age: 44, Occupation: "Software Engineer"})
	persons = append(persons, Person{"Brian", 51, "SSIS Developer"})
	persons = append(persons, Person{"Snow White", 43, "Hacker"})
	persons = append(persons, Person{"Focal Recursion", 43, "Hacker"})
	persons = append(persons, Person{"Jsldub", 43, "Hacker"})
	persons = append(persons, Person{"Rube5", 44, "Hacker"})
	persons = append(persons, Person{"SandwichMan", 44, "Hacker"})
	persons = append(persons, Person{"PaperNaffin", 43, "Hacker"})

	sort.Slice(persons, func(i, j int) bool {
		return persons[i].Name < persons[j].Name
	})
	fmt.Println("persons =", persons)

	persons[0].changeAgeTo(50)

	employees := []Employee{}

	for _, p := range persons {
		pTemp := p
		employees = append(employees, Employee(&pTemp))
	}

	for _, p := range employees {
		p.printName()
		fmt.Println("Name again is", p.getName())
		fmt.Println("p = ", p)
	}

	fmt.Println("*** beginning ***")
	fmt.Println(employees[0].(*Person))
	fmt.Println(employees[0].(*Person).Name)
	fmt.Println(employees[0].(*Person).Age)
	fmt.Println("*****************")

	fmt.Printf("Employees = %#v\n", employees)

	var dLove DeservesLove = &persons[1]
	dLove.printDeservesLoveName()

	handleRequests()
}
