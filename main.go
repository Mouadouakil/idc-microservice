package main

import (
	"fmt"
	"math/rand"
	"time"
    "encoding/json"
)

// LogEntry represents a log entry structure
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
var possibleValues = map[string][]string{
	"Msg":   {"for_icp"},
    "Level": {"info"},
	"Op":    {"auth"},
	"Step":  {"begin", "end", "attempt", "ok", "ko", "otp_email", "otp_sms", "otp_pushNotif", "lock", "end"},
	"Err":   {"cancel", "E_AUTH_AF_LOCKED", "E_AUTH_AF_BLOCKED", "E_AUTH_AF_PWD_WRONGCREDENTIALS", "E_AUTH_AF_WRONG_OTP", "E_AUTH_USER_NOT_FOUND"},
    "Ua" : {"Chrome v 119.0.0.0 - Windows v 10.0","Chrome v 120.0.0.0 - Windows v 10.0"},
    "Sta" : {"a39487c32eae","59ab3a08d9ee", "59ab3a08d9ee"},
    "Asid" : {"gzer15Gujtcrgg", " "},
    "Spn" : {"e-iDarati", "My Bank", "Portail citoyen CNIE"},
    "Spi" : {"5541902c-9dac-47e1-a780-dd962c6abd98","d9649157-bde3-48f7-8c6f-09cc5ae95ca4"},
    "Svi" : {"29b31af2-3027-4c3b-925d-8f542e638956", "29b31af2-3027-4c3b-925d-8f542e638956", "d9649157-bde3-48f7-8c6f-09cc5ae95ca4"},
    "Svn" : {"Authentification", ""},
    "Ct" : { "ma-cnie-v1", "ma-cnie-v2"},
    "Acr" : { "pwd", "otp", "pwdNotp", "scNpin", "scNbio", "scNpinNbio", "scNpinNface", "scNpinNbioNface", "scNface", "phNpin", "phNpinNface", "phNface" },
}
 

func generateRandomLogEntry() LogEntry {
	rand.Seed(time.Now().UnixNano())

	return LogEntry{
		Msg:    possibleValues["Msg"][0],
		Level:  possibleValues["Level"][0],
		Time:   time.Now(),
		Op:     possibleValues["Op"][0],
		Step:   possibleValues["Step"][rand.Intn(len(possibleValues["Step"]))],
		Err:    possibleValues["Err"][rand.Intn(len(possibleValues["Err"]))],
		IP:     generateRandomIP(),
		UA:     possibleValues["Ua"][rand.Intn(len(possibleValues["Ua"]))],
		Sta:    possibleValues["Sta"][rand.Intn(len(possibleValues["Sta"]))],
		Asid:   possibleValues["Asid"][rand.Intn(len(possibleValues["Asid"]))],
		Spi:    possibleValues["Spi"][rand.Intn(len(possibleValues["Spi"]))],
		Spn:    possibleValues["Spn"][rand.Intn(len(possibleValues["Spn"]))],
		Svi:    possibleValues["Svi"][rand.Intn(len(possibleValues["Svi"]))],
		Svn:    possibleValues["Svn"][rand.Intn(len(possibleValues["Svn"]))],
		Ct:     possibleValues["Ct"][rand.Intn(len(possibleValues["Ct"]))],
		Ca:     true,
		Pa:     false,
		Acr:    possibleValues["Acr"][rand.Intn(len(possibleValues["Acr"]))],
	}
}

// generateRandomIP generates a random IP address for simulation purposes
func generateRandomIP() string {
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
}

func main() {
	// Generate and print 5 random log entries
	for i := 0; i < 5; i++ {
		logEntry := generateRandomLogEntry()
        logJSON, err := json.Marshal(logEntry);
        if err != nil {
            fmt.Println("Error marshalling JSON:", err)
            return
        };
        fmt.Println(string(logJSON))
	}


}
