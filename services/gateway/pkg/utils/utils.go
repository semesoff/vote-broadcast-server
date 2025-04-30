package utils

import (
	"encoding/json"
	"gateway-service/pkg/models"
	"gateway-service/proto/vote"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(payload)
	return err
}

func ToProtoCreateVoteData(voteData models.Vote) *vote.CreateVoteRequest {
	request := &vote.CreateVoteRequest{
		PollId:    int64(voteData.PollId),
		UserId:    int64(voteData.UserId),
		OptionsId: make([]int64, 0),
	}

	for _, option := range voteData.OptionsId {
		request.OptionsId = append(request.OptionsId, int64(option))
	}

	return request
}
