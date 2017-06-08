package action

import (
	"testing"
	"time"
)

func TestStartWorkers(t *testing.T) {

	providerList, err := InitializeImplementations()
	if err != nil {
		t.Fatalf("provider could not be initialized : %v", err.Error())
	}

	numberOfWorkers := 17
	StartWorkers(providerList, numberOfWorkers)
	if len(quitChans) != numberOfWorkers {
		t.Errorf("invalid number of channels created : %v, expected : %v", len(quitChans), numberOfWorkers)
	}
	//clean up
	StopWorkers()
}

func TestReturnAck(t *testing.T) {

	providerList, err := InitializeImplementations()
	if err != nil {
		t.Fatalf("provider could not be initialized : %v", err.Error())
	}

	numberOfWorkers := 5
	StartWorkers(providerList, numberOfWorkers)

	for i, _ := range quitChans {
		quitChans[i] <- true
		select {
		case <-quitAckChan:
		case <-time.After(3 * time.Second):
			t.Errorf("timeout while waiting for acknowledge of close worker")
		}
	}
	quitChans = []chan bool{}
}
