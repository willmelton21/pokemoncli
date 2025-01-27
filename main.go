package main


import (
    "fmt"
     "strings"
     "unicode"
     "os"
     "bufio"

 )

 type cliCommand struct {
        name    string
        description string
        callback func() error
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

func commandExit() error {

    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
    }

func help() error {

    fmt.Println("Welcome to the Pokedex!")
    fmt.Printf("Usage:\n\n")
    for _, value := range commands {
        fmt.Printf("%s: %s\n",value.name,value.description)
    }
    return nil
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
    }



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
            err := val.callback()
            if err != nil {
                fmt.Println(err)
            }
        } else {
            fmt.Println("Unkown command")
        }
 }

 }

     

