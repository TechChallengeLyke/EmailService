package action

import (
	"fmt"
	"github.com/TechChallengeLyke/EmailService/data"
	"github.com/TechChallengeLyke/EmailService/provider"
)

//limit queue to 1000 emails
var (
	sendQueue   = make(chan data.Email, 1000)
	quitChans   = []chan bool{}
	quitAckChan = make(chan bool)
)

//check mail for completeness of data, persist mail and send it to the workers for processing
func SendMail(email *data.Email) error {

	err := CheckEmailData(email)
	if err != nil {
		fmt.Printf("email data not complete : %v\n", err.Error())
		return err
	}

	err = email.Create()
	if err != nil {
		//since there is no emphasis on logging what happened with the mails:
		//log error, but try to send it anyway
		fmt.Printf("error during email.Create() : ", err.Error())
	}

	sendQueue <- *email
	return nil
}

//start the specified number of workers(goroutines) and initialize their respective quit channels
func StartWorkers(providerList map[string]provider.EmailProvider, numberOfWorkers int) {

	for i := 0; i < numberOfWorkers; i++ {
		quitChans = append(quitChans, make(chan bool))

		go func(providerList map[string]provider.EmailProvider, quitChan chan bool) {

			for {
				select {
				case email := <-sendQueue:
					processEmail(providerList, email)
				case <-quitChan:
					quitAckChan <- true
					return
				}
			}
		}(providerList, quitChans[i])

	}

	fmt.Printf("number of quit channels : %v\n", len(quitChans))
}

//stop all workers
func StopWorkers() {

	//send quit to all quit channels and wait for the ack from all workers to make sure, that they are finished with their current work
	for i, _ := range quitChans {
		fmt.Printf("Quitting worker %v \n", i)
		quitChans[i] <- true
		<-quitAckChan
	}
	quitChans = []chan bool{}
}
