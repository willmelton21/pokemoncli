package main


import (
    "fmt"
     "strings"
     "unicode"
     "os"
     "bufio"
     "log"
     "internal/pokeapi"

 )

 type cliCommand struct {
        name    string
        description string
        callback func(cfg *config) error
    }

 type config struct {
	 Next   	string
	 Previous	string
 }

var commands map[string]cliCommand
 
func cleanInput(text string) []string {

        var retSlice []string

        lowerText := strings.ToLower(text)
        
        var app string

        for _, ch := range lowerText {
        
        if unicode.IsSpace(ch) {
            retSlice = append(retSlice,app)
            app = ""
        } else{
            app += string(ch)
        }
    }
    retSlice = append(retSlice,app) //append last one

        return retSlice
}

func commandExit(cfg *config) error {

    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
    }

func help(cfg *config) error {

    fmt.Println("Welcome to the Pokedex!")
    fmt.Printf("Usage:\n\n")
    for _, value := range commands {
        fmt.Printf("%s: %s\n",value.name,value.description)
    }
    return nil
}

func displayMap(cfg *config) error {

	str := "https://pokeapi.co/api/v2/location-area"
	if cfg.Next != "" {
		str = cfg.Next
	}
	res,err := pokeapi.GetLocationAreas(&str)
	if err != nil {
		log.Fatal(err)
	}
	cfg.Next = *res.Next
	if(res.Previous != nil) {

	cfg.Previous = *res.Previous 
	}
	for i := 0; i < 19; i++ {
		fmt.Println(res.Results[i].Name)	
	}

	return nil
}

func displayMapB(cfg *config) error {

	if(cfg.Previous == "") {
		fmt.Println("you're on the first page")
		return nil
	} else {

		res,err := pokeapi.GetLocationAreas(&cfg.Previous)
	if err != nil {
		log.Fatal(err)
	}
	cfg.Next = *res.Next
	if (res.Previous != nil) {

	cfg.Previous = *res.Previous 
	}
	for i := 0; i < 19; i++ {
		fmt.Println(res.Results[i].Name)	
	}

	return nil


	}

}

func main() {
     commands = map[string]cliCommand{
        "exit": {
            name:   "exit",
            description: "Exit the Pokedex",
            callback: commandExit,
           },
        "help": {
            name: "help",
            description: "Displays a help message",
            callback: help,
        },
        "map": {
            name: "map",
            description: "Displays 20 location areas in Pokemon world",
            callback: displayMap,
        },
	"mapb":{
		name: "mapb",
		description: "Display the previous 20 location areas in a Pokemon world",
		callback: displayMapB,
	},
    }

    cfg := config{Next: "",Previous: ""}
    scanner := bufio.NewScanner(os.Stdin)
    for {
        if fileInfo, _ := os.Stdin.Stat(); (fileInfo.Mode() & os.ModeCharDevice) != 0 {
        fmt.Print("Pokedex > ")
        }
        scanner.Scan()
        userInput := scanner.Text()
        userInput = strings.ToLower(userInput)
        fields := strings.Fields(userInput)
        if len(fields) == 0 {
            continue;
        }
       if val, ok := commands[fields[0]]; ok {
            err := val.callback(&cfg)
            if err != nil {
                fmt.Println(err)
            }
        } else {
            fmt.Println("Unkown command")
        }
 }

 }

     

