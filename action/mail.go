package action

import (
	"fmt"
	"github.com/TechChallengeLyke/EmailService/data"
)

//limit queue to 1000 emails
var (
	sendQueue   = make(chan data.Email, 1000)
	quitChans   = []chan bool{}
	quitAckChan = make(chan bool)
)

//persist mail and send it to the workers for processing
func SendMail(email data.Email) {

	err := email.Save()
	if err != nil {
		//log error, but try to send it anyway
	}

	sendQueue <- email
}

//start the specified number of workers(goroutines) and initialize their respective quit channels
func StartWorkers(numberOfWorkers int) {

	for i := 0; i < numberOfWorkers; i++ {
		quitChans = append(quitChans, make(chan bool))

		go func(quitChan chan bool) {
			//go func(quitChan chan bool) {
			fmt.Printf("Started anonymous goroutine\n")

			for {
				select {
				case email := <-sendQueue:
					processEmail(email)
				case <-quitChan:
					fmt.Printf("Stopping anonymous goroutine \n")
					quitAckChan <- true
					return
				}
			}
			fmt.Printf("anonymous goroutine stopped normally")
		}(quitChans[i])

	}

	fmt.Printf("number of quit channels : %v\n", len(quitChans))
}

//stop all workers
func StopWorkers() {

	//send quit to all quit channels and wait for the ack from all workers to make sure, that they are finished with their current work
	for i, _ := range quitChans {
		fmt.Printf("Quitting channel %v \n", i)
		quitChans[i] <- true
		<-quitAckChan
	}
}

func processEmail(email data.Email) {

	fmt.Printf("Processing email: %v \n", email.Subject)
}
