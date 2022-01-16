package models

import (
	"context"
	"enttest/ent"
	"enttest/ent/aircraft"
	"fmt"
	"github.com/google/uuid"
)

type Aircraft struct {
	AircraftId uuid.UUID
}

func GetAircrafts(ctx context.Context, client *ent.Client) ([]*Aircraft, error) {
	a, err := client.Aircraft.
		Query().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	// log.Println("user returned: ", a)

	aircraftArray := make([]*Aircraft, len(a))
	for i, aircraftInfo := range a {
		aircraftArray[i] = &Aircraft{
			AircraftId: aircraftInfo.ID,
		}
	}

	return aircraftArray, nil
}

/*func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New User must not include id or it must be set to zero")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}*/

func GetAircraftByID(ctx context.Context, client *ent.Client, uuid uuid.UUID) (*Aircraft, error) {
	/*for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}

	return Aircraft{}, fmt.Errorf("User with ID '%v' not found", id)*/
	a, err := client.Aircraft.
		Query().
		Where(aircraft.IDEQ(uuid)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	// log.Println("user returned: ", a)

	return &Aircraft{
		AircraftId: a.ID,
	}, nil
}

/*
func UpdateUser(u Aircraft) (Aircraft, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

func RemoveUserById(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("User with ID '%v' not found", id)
}*/
