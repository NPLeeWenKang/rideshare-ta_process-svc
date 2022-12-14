package main

type _TripStatus struct {
	PENDING  string
	REJECTED string
	ACCEPTED string
	DRIVING  string
	DONE     string
}

var TripStatus = _TripStatus{
	PENDING:  "PENDING",
	REJECTED: "REJECTED",
	ACCEPTED: "ACCEPTED",
	DRIVING:  "DRIVING",
	DONE:     "DONE",
}
