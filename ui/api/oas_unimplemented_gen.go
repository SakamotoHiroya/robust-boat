// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// AddVote implements AddVote operation.
//
// Add a vote to a poll.
//
// POST /polls/{pollId}/vote
func (UnimplementedHandler) AddVote(ctx context.Context, req VoteCreationParameter, params AddVoteParams) (r AddVoteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreatePoll implements CreatePoll operation.
//
// Create a new poll.
//
// POST /polls/create
func (UnimplementedHandler) CreatePoll(ctx context.Context, req *PollCreationParameter) (r CreatePollRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetPoll implements getPoll operation.
//
// Get a specific poll.
//
// GET /polls/{pollId}
func (UnimplementedHandler) GetPoll(ctx context.Context, params GetPollParams) (r GetPollRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetVote implements getVote operation.
//
// Get a specific vote in a poll.
//
// GET /polls/{pollId}/vote/{voteId}
func (UnimplementedHandler) GetVote(ctx context.Context, params GetVoteParams) (r GetVoteRes, _ error) {
	return r, ht.ErrNotImplemented
}

// Login implements login operation.
//
// User login.
//
// POST /auth/login
func (UnimplementedHandler) Login(ctx context.Context, req *LoginParameter) (r LoginRes, _ error) {
	return r, ht.ErrNotImplemented
}

// Ping implements Ping operation.
//
// Check if the server is running.
//
// GET /ping
func (UnimplementedHandler) Ping(ctx context.Context) (r *PingOK, _ error) {
	return r, ht.ErrNotImplemented
}

// Register implements Register operation.
//
// Register a new user.
//
// POST /auth/register
func (UnimplementedHandler) Register(ctx context.Context, req *UserRegisterParameter) (r RegisterRes, _ error) {
	return r, ht.ErrNotImplemented
}
