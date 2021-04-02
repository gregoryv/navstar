package navstar

import "testing"

func Test_Pilot(t *testing.T) {
	var usr User
	role := usr.Use(NewSystem(), &Pilot{})

	if _, err := role.ListFlightplans(); err != nil {
		t.Fatal(err)
	}
	if err := role.SubmitFlightplan(Flightplan{}); err != nil {
		t.Fatal(err)
	}
}

func Test_Passenger(t *testing.T) {
	var user User
	role := user.Use(NewSystem(), &Passenger{})

	if _, err := role.ListFlightplans(); err != nil {
		t.Fatal(err)
	}
	if err := role.SubmitFlightplan(Flightplan{}); err == nil {
		t.Fatal("expected unauthorized")
	}
}
