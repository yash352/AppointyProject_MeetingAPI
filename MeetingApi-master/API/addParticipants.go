package API

import (
	"encoding/json"
	"fmt"
	"net/http"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/manzar97/reastapi/models"
)


func (client Client) AddMeetingParticipants(meetingId int,
	email string,
	firstName string,
	lastName string,
	"rsvp":"yes"
	
	addMeetingParticipantsRequest := AddMeetingParticipantsRequest{
		Email:                 email,
		FirstName:             firstName,
		LastName:              lastName,
		
	}

	endpoint := fmt.Sprintf("/meetings/%d/participants", meetingId)
	httpMethod := http.MethodPost

	var reqBody []byte
	reqBody, err = json.Marshal(addMeetingParticipantsRequest)
	if err != nil {
		return
	}

	var b []byte
	b, err = client.executeRequestWithBody(endpoint, httpMethod, reqBody)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &addMeetingParticipantsResponse)
	if err != nil {
		return
	}

	return
}