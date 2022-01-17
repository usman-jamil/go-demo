package models

import (
	"context"
	"enttest/ent"
	"enttest/ent/aircraft"
	"fmt"
	"github.com/google/uuid"
)

type Aircraft struct {
	AircraftId             uuid.UUID
	CompanyId              uuid.UUID
	CurrentFlightHours     float32
	CurrentCycles          int
	AircraftRegistration   string
	BaseAirportCode        string
	Manufacturer           string
	ManufacturerDesignator string
	CommonDesignation      string
	CommonName             string
	PilotsRequiredToFly    int
	DefaultValues          string
	MaximumValues          string
	CurrentLandings        int
	FuelDetails            string
	OilDetails             string
}

func GetAircrafts(ctx context.Context, client *ent.Client) ([]*Aircraft, error) {
	a, err := client.Aircraft.
		Query().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	aircraftArray := make([]*Aircraft, len(a))
	for i, aircraftInfo := range a {
		aircraftArray[i] = &Aircraft{
			AircraftId:             aircraftInfo.ID,
			CompanyId:              aircraftInfo.CompanyID,
			CurrentFlightHours:     aircraftInfo.CurrentFlightHours,
			CurrentCycles:          aircraftInfo.CurrentCycles,
			AircraftRegistration:   aircraftInfo.AircraftRegistration,
			BaseAirportCode:        aircraftInfo.BaseAirportCode,
			Manufacturer:           aircraftInfo.Manufacturer,
			ManufacturerDesignator: aircraftInfo.ManufacturerDesignator,
			CommonDesignation:      aircraftInfo.CommonDesignation,
			CommonName:             aircraftInfo.CommonName,
			PilotsRequiredToFly:    aircraftInfo.PilotsRequiredToFly,
			DefaultValues:          aircraftInfo.DefaultValues,
			MaximumValues:          aircraftInfo.MaximumValues,
			CurrentLandings:        aircraftInfo.CurrentLandings,
			OilDetails:             aircraftInfo.OilDetails,
			FuelDetails:            aircraftInfo.FuelDetails,
		}
	}

	return aircraftArray, nil
}

func GetAircraftsPaged(ctx context.Context, client *ent.Client, limit int, offset int) ([]*Aircraft, error) {
	a, err := client.Aircraft.
		Query().
		Offset(offset).
		Limit(limit).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	aircraftArray := make([]*Aircraft, len(a))
	for i, aircraftInfo := range a {
		aircraftArray[i] = &Aircraft{
			AircraftId:             aircraftInfo.ID,
			CompanyId:              aircraftInfo.CompanyID,
			CurrentFlightHours:     aircraftInfo.CurrentFlightHours,
			CurrentCycles:          aircraftInfo.CurrentCycles,
			AircraftRegistration:   aircraftInfo.AircraftRegistration,
			BaseAirportCode:        aircraftInfo.BaseAirportCode,
			Manufacturer:           aircraftInfo.Manufacturer,
			ManufacturerDesignator: aircraftInfo.ManufacturerDesignator,
			CommonDesignation:      aircraftInfo.CommonDesignation,
			CommonName:             aircraftInfo.CommonName,
			PilotsRequiredToFly:    aircraftInfo.PilotsRequiredToFly,
			DefaultValues:          aircraftInfo.DefaultValues,
			MaximumValues:          aircraftInfo.MaximumValues,
			CurrentLandings:        aircraftInfo.CurrentLandings,
			OilDetails:             aircraftInfo.OilDetails,
			FuelDetails:            aircraftInfo.FuelDetails,
		}
	}

	return aircraftArray, nil
}

func GetAircraftsByFlightHoursAscending(ctx context.Context, client *ent.Client) ([]*Aircraft, error) {
	a, err := client.Aircraft.
		Query().
		Where(aircraft.CurrentFlightHoursNotNil()).
		Order(ent.Asc(aircraft.FieldCurrentFlightHours)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	aircraftArray := make([]*Aircraft, len(a))
	for i, aircraftInfo := range a {
		aircraftArray[i] = &Aircraft{
			AircraftId:             aircraftInfo.ID,
			CompanyId:              aircraftInfo.CompanyID,
			CurrentFlightHours:     aircraftInfo.CurrentFlightHours,
			CurrentCycles:          aircraftInfo.CurrentCycles,
			AircraftRegistration:   aircraftInfo.AircraftRegistration,
			BaseAirportCode:        aircraftInfo.BaseAirportCode,
			Manufacturer:           aircraftInfo.Manufacturer,
			ManufacturerDesignator: aircraftInfo.ManufacturerDesignator,
			CommonDesignation:      aircraftInfo.CommonDesignation,
			CommonName:             aircraftInfo.CommonName,
			PilotsRequiredToFly:    aircraftInfo.PilotsRequiredToFly,
			DefaultValues:          aircraftInfo.DefaultValues,
			MaximumValues:          aircraftInfo.MaximumValues,
			CurrentLandings:        aircraftInfo.CurrentLandings,
			OilDetails:             aircraftInfo.OilDetails,
			FuelDetails:            aircraftInfo.FuelDetails,
		}
	}

	return aircraftArray, nil
}

func GetAircraftsByFlightHoursDescending(ctx context.Context, client *ent.Client) ([]*Aircraft, error) {
	a, err := client.Aircraft.
		Query().
		Where(aircraft.CurrentFlightHoursNotNil()).
		Order(ent.Desc(aircraft.FieldCurrentFlightHours)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	aircraftArray := make([]*Aircraft, len(a))
	for i, aircraftInfo := range a {
		aircraftArray[i] = &Aircraft{
			AircraftId:             aircraftInfo.ID,
			CompanyId:              aircraftInfo.CompanyID,
			CurrentFlightHours:     aircraftInfo.CurrentFlightHours,
			CurrentCycles:          aircraftInfo.CurrentCycles,
			AircraftRegistration:   aircraftInfo.AircraftRegistration,
			BaseAirportCode:        aircraftInfo.BaseAirportCode,
			Manufacturer:           aircraftInfo.Manufacturer,
			ManufacturerDesignator: aircraftInfo.ManufacturerDesignator,
			CommonDesignation:      aircraftInfo.CommonDesignation,
			CommonName:             aircraftInfo.CommonName,
			PilotsRequiredToFly:    aircraftInfo.PilotsRequiredToFly,
			DefaultValues:          aircraftInfo.DefaultValues,
			MaximumValues:          aircraftInfo.MaximumValues,
			CurrentLandings:        aircraftInfo.CurrentLandings,
			OilDetails:             aircraftInfo.OilDetails,
			FuelDetails:            aircraftInfo.FuelDetails,
		}
	}

	return aircraftArray, nil
}

func GetAircraftsByType(ctx context.Context, client *ent.Client, designation string) ([]*Aircraft, error) {
	a, err := client.Aircraft.
		Query().
		Where(aircraft.CommonDesignationEqualFold(designation)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	aircraftArray := make([]*Aircraft, len(a))
	for i, aircraftInfo := range a {
		aircraftArray[i] = &Aircraft{
			AircraftId:             aircraftInfo.ID,
			CompanyId:              aircraftInfo.CompanyID,
			CurrentFlightHours:     aircraftInfo.CurrentFlightHours,
			CurrentCycles:          aircraftInfo.CurrentCycles,
			AircraftRegistration:   aircraftInfo.AircraftRegistration,
			BaseAirportCode:        aircraftInfo.BaseAirportCode,
			Manufacturer:           aircraftInfo.Manufacturer,
			ManufacturerDesignator: aircraftInfo.ManufacturerDesignator,
			CommonDesignation:      aircraftInfo.CommonDesignation,
			CommonName:             aircraftInfo.CommonName,
			PilotsRequiredToFly:    aircraftInfo.PilotsRequiredToFly,
			DefaultValues:          aircraftInfo.DefaultValues,
			MaximumValues:          aircraftInfo.MaximumValues,
			CurrentLandings:        aircraftInfo.CurrentLandings,
			OilDetails:             aircraftInfo.OilDetails,
			FuelDetails:            aircraftInfo.FuelDetails,
		}
	}

	return aircraftArray, nil
}

func GetAircraftsByRegistration(ctx context.Context, client *ent.Client, registration string) ([]*Aircraft, error) {
	a, err := client.Aircraft.
		Query().
		Where(aircraft.AircraftRegistrationEqualFold(registration)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	aircraftArray := make([]*Aircraft, len(a))
	for i, aircraftInfo := range a {
		aircraftArray[i] = &Aircraft{
			AircraftId:             aircraftInfo.ID,
			CompanyId:              aircraftInfo.CompanyID,
			CurrentFlightHours:     aircraftInfo.CurrentFlightHours,
			CurrentCycles:          aircraftInfo.CurrentCycles,
			AircraftRegistration:   aircraftInfo.AircraftRegistration,
			BaseAirportCode:        aircraftInfo.BaseAirportCode,
			Manufacturer:           aircraftInfo.Manufacturer,
			ManufacturerDesignator: aircraftInfo.ManufacturerDesignator,
			CommonDesignation:      aircraftInfo.CommonDesignation,
			CommonName:             aircraftInfo.CommonName,
			PilotsRequiredToFly:    aircraftInfo.PilotsRequiredToFly,
			DefaultValues:          aircraftInfo.DefaultValues,
			MaximumValues:          aircraftInfo.MaximumValues,
			CurrentLandings:        aircraftInfo.CurrentLandings,
			OilDetails:             aircraftInfo.OilDetails,
			FuelDetails:            aircraftInfo.FuelDetails,
		}
	}

	return aircraftArray, nil
}

func GetAircraftByID(ctx context.Context, client *ent.Client, uuid uuid.UUID) (*Aircraft, error) {
	a, err := client.Aircraft.
		Query().
		Where(aircraft.IDEQ(uuid)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	return &Aircraft{
		AircraftId:             a.ID,
		CompanyId:              a.CompanyID,
		CurrentFlightHours:     a.CurrentFlightHours,
		CurrentCycles:          a.CurrentCycles,
		AircraftRegistration:   a.AircraftRegistration,
		BaseAirportCode:        a.BaseAirportCode,
		Manufacturer:           a.Manufacturer,
		ManufacturerDesignator: a.ManufacturerDesignator,
		CommonDesignation:      a.CommonDesignation,
		CommonName:             a.CommonName,
		PilotsRequiredToFly:    a.PilotsRequiredToFly,
		DefaultValues:          a.DefaultValues,
		MaximumValues:          a.MaximumValues,
		CurrentLandings:        a.CurrentLandings,
		OilDetails:             a.OilDetails,
		FuelDetails:            a.FuelDetails,
	}, nil
}
