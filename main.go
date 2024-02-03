package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"
	"github.com/gin-gonic/gin"
)

type LogEntry struct {
	Msg   string    `json:"msg"`
	Level string    `json:"level"`
	Time  time.Time `json:"time"`
	Op    string    `json:"op"`
	Step  string    `json:"step"`
	Err   string    `json:"err,omitempty"`
	IP    string    `json:"ip"`
	UA    string    `json:"ua"`
	Sta   string    `json:"sta"`
	Asid  string    `json:"asid"`
	Spi   string    `json:"spi"`
	Spn   string    `json:"spn"`
	Svi   string    `json:"svi"`
	Svn   string    `json:"svn"`
	Ct    string    `json:"ct,omitempty"`
	Ca    bool      `json:"ca,omitempty"`
	Pa    bool      `json:"pa,omitempty"`
	Acr   string    `json:"acr,omitempty"`
}

// Possible values for each field

// for idc-login         {"AUTHERR004","AUTHERR005","USRCRM001","CARDERR004","CARDERR006","CARDERR007","COMPERR001","COMPERR002", "COMPERR004","REQERR000", "REQERR001", "REQERR000", "REQERR005" ,"REQERR008" } 

var possibleValues = map[string][]string{
	"Msg":   {"for_icp"},
	"Level": {"info"},
	"Op":    {"auth"},
	"Step":  {"begin", "end", "attempt", "ok", "ko", "otp_email", "otp_sms", "otp_pushNotif", "lock", "end"},
	"Err":   {"cancel", "E_AUTH_AF_LOCKED", "E_AUTH_AF_BLOCKED", "E_AUTH_AF_PWD_WRONGCREDENTIALS", "E_AUTH_AF_WRONG_OTP", "E_AUTH_USER_NOT_FOUND"},
	"Ua":    {"Chrome v 119.0.0.0 - Windows v 10.0", "Chrome v 120.0.0.0 - Windows v 10.0", "Safari v 21.0.0.0 - Macos v 12.0"},
	"Sta":   {"a39487c32eae", "59ab3a08d9ee", "59ab3a08d9ee","59ab3a08dare","13lb3a08dare"},
	"Asid":  {"gzer15Gujtcrgg", " "},
	"Spn":   {"e-iDarati", "My Bank", "Portail citoyen CNIE"},
	"Spi":   {"5541902c-9dac-47e1-a780-dd962c6abd98", "d9649157-bde3-48f7-8c6f-09cc5ae95ca4", "d9649157-bde3-48f7-8c6f-09cdaae95ca4"},
	"Svi":   {"29b31af2-3027-4c3b-925d-8f542e638956", "29b31af2-3027-4c3b-925d-8f542e638956", "d9649157-bde3-48f7-8c6f-09cc5ae95ca4"},
	"Svn":   {"Authentification", ""},
	"Ct":    {"ma-cnie-v1", "ma-cnie-v2"},
	"Acr":   {"pwd", "otp", "pwdNotp", "scNpin", "scNbio", "scNpinNbio", "scNpinNface", "scNpinNbioNface", "scNface", "phNpin", "phNpinNface", "phNface"},
}

const (
	LogLevelInfo = "info"
)

func getRandomValue(options []string) string {
	return options[rand.Intn(len(options))]
}

func generateRandomLogEntry() LogEntry {
	logEntry := LogEntry{
		Msg:   getRandomValue(possibleValues["Msg"]),
		Level: LogLevelInfo,
		Time:  time.Now(),
		Op:    getRandomValue(possibleValues["Op"]),
		Step:  getRandomValue(possibleValues["Step"]),
		IP:    generateRandomIP(),
		UA:    getRandomValue(possibleValues["Ua"]),
		Sta:   getRandomValue(possibleValues["Sta"]),
		Asid:  getRandomValue(possibleValues["Asid"]),
		Spi:   getRandomValue(possibleValues["Spi"]),
		Spn:   getRandomValue(possibleValues["Spn"]),
		Svi:   getRandomValue(possibleValues["Svi"]),
		Svn:   getRandomValue(possibleValues["Svn"]),
		Ct:    getRandomValue(possibleValues["Ct"]),
		Ca:    true,
		Pa:    false,
		Acr:   getRandomValue(possibleValues["Acr"]),
	}

	if rand.Float32() < 0.2 {
		logEntry.Err = getRandomValue(possibleValues["Err"])
	}

	return logEntry
}

func generateRandomIP() string {
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
}

func main() {
	rand.Seed(time.Now().UnixNano())

	router := gin.Default()

	router.GET("/simulate-logs", func(c *gin.Context) {
		for i := 0; i < 50; i++ {
			logEntry := generateRandomLogEntry()
			logJSON, err := json.Marshal(logEntry)
			if err != nil {
				log.Println("Error marshalling JSON:", err)
				continue
			}
			log.Println(string(logJSON))
			fmt.Println(string(logJSON))
		}

		c.JSON(200, gin.H{
			"message": "Logs simulated successfully",
		})
	})

	router.Run(":8080")
}
