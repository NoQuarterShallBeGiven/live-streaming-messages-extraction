package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
)

const (
	ansiESC            = "\u001b"
	ansiReset          = "[0m"
	configFileLocation = "config.json"
)

type Config struct {
	Commands map[string]string `json:"commands"`
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func colorString(FgColor int, input string) string {
	return fmt.Sprintf(ansiESC+"[38;5;%dm%s"+ansiESC+ansiReset, FgColor, input)
}

type JsonRequest struct {
	Source  string `json:"Source"`
	User    string `json:"User"`
	Comment string `json:"Comment"`
	Amount  string `json:"Amount"`
}

func speak(input string) {
	// TODO : do not read URLs.
	cmd := exec.Command("espeak", []string{input}...)
	if err := cmd.Run(); err != nil {
		log.Fatalf("%v\n", err)
	}
}

func system(command string) ([]byte, error) {
	ss := strings.Split(command, " ")
	cmd := exec.Command(ss[0], ss[1:]...)
	cmd.Stderr = os.Stderr
	return cmd.Output()
}

func (c *Config) parseCommand(command string, paid bool) {
	rCmd, cmdExist := func(s string) (string, bool) {
		re := regexp.MustCompile("(?i)![A-Za-z]+")
		return re.FindString(s), re.MatchString(s)
	}(command)
	if cmdExist {
		if val, ok := c.Commands[rCmd]; ok {
			fmt.Println(val, ok)
			// TODO :
		}
	}
}

func (c *Config) start() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("%v\n", err)
		}
		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
			j := JsonRequest{}
			if err := json.Unmarshal(msg, &j); err != nil {
				log.Println(err)
			}
			c.parseCommand(j.Comment, false)
			if len(j.Amount) > 1 {
				fmt.Printf("\n%s", colorString(3, "HYPERCHAT"))
				pf, err := strconv.ParseFloat(j.Amount[1:], 64)
				if err != nil {
					log.Panicf("\n%v\n", err)
				}
				if pf > 14.00 {
					fmt.Print(colorString(4, " whale alert lol"))
				}
				c.parseCommand(j.Comment, true)
				//speak(fmt.Sprintf("user %s, gave %s monies, to say %s", j.User, j.Amount[1:], j.Comment))
			}
			ms := fmt.Sprintf("\n%s:%s: %s %s", colorString(5, j.Source), colorString(1, j.User), colorString(3, j.Amount), colorString(2, j.Comment))
			fmt.Print(ms)
			//fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))
			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})
	fmt.Println("starting odysee-livechat server")
	if err := http.ListenAndServe(":8839", nil); err != nil {
		log.Printf("%v\n", err)
	}

}

func (c *Config) writeConfig() error {
	marshal, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return err
	}
	fmt.Println(string(marshal))
	if err := os.WriteFile("config.json", marshal, 0666); err != nil {
		return err
	}
	return nil
}

func exists(filename string) bool {
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func main() {
	c := &Config{}
	c.Commands = make(map[string]string)
	if !exists(configFileLocation) {
		c.Commands["!status"] = "echo status"
		if err := c.writeConfig(); err != nil {
			log.Fatalf("%v\n", err)
		}
		fmt.Printf("created config file at %s\n", configFileLocation)
		return
	}
	bytes, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(bytes, c); err != nil {
		log.Fatalf("%v", err)
	}
	c.start()
}
