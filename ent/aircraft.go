// Code generated by entc, DO NOT EDIT.

package ent

import (
	"enttest/ent/aircraft"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Aircraft is the model entity for the Aircraft schema.
type Aircraft struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CompanyID holds the value of the "company_id" field.
	CompanyID uuid.UUID `json:"company_id,omitempty"`
	// CurrentFlightHours holds the value of the "current_flight_hours" field.
	CurrentFlightHours *float32 `json:"current_flight_hours,omitempty"`
	// CurrentCycles holds the value of the "current_cycles" field.
	CurrentCycles *int `json:"current_cycles,omitempty"`
	// AircraftRegistration holds the value of the "aircraft_registration" field.
	AircraftRegistration string `json:"aircraft_registration,omitempty"`
	// BaseAirportCode holds the value of the "base_airport_code" field.
	BaseAirportCode string `json:"base_airport_code,omitempty"`
	// Manufacturer holds the value of the "manufacturer" field.
	Manufacturer *string `json:"manufacturer,omitempty"`
	// ManufacturerDesignator holds the value of the "manufacturer_designator" field.
	ManufacturerDesignator string `json:"manufacturer_designator,omitempty"`
	// CommonDesignation holds the value of the "common_designation" field.
	CommonDesignation *string `json:"common_designation,omitempty"`
	// CommonName holds the value of the "common_name" field.
	CommonName string `json:"common_name,omitempty"`
	// PilotsRequiredToFly holds the value of the "pilots_required_to_fly" field.
	PilotsRequiredToFly *int `json:"pilots_required_to_fly,omitempty"`
	// DefaultValues holds the value of the "default_values" field.
	DefaultValues string `json:"default_values,omitempty"`
	// MaximumValues holds the value of the "maximum_values" field.
	MaximumValues string `json:"maximum_values,omitempty"`
	// CurrentLandings holds the value of the "current_landings" field.
	CurrentLandings *int `json:"current_landings,omitempty"`
	// FuelDetails holds the value of the "fuel_details" field.
	FuelDetails string `json:"fuel_details,omitempty"`
	// OilDetails holds the value of the "oil_details" field.
	OilDetails string `json:"oil_details,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Aircraft) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case aircraft.FieldCurrentFlightHours:
			values[i] = new(sql.NullFloat64)
		case aircraft.FieldCurrentCycles, aircraft.FieldPilotsRequiredToFly, aircraft.FieldCurrentLandings:
			values[i] = new(sql.NullInt64)
		case aircraft.FieldAircraftRegistration, aircraft.FieldBaseAirportCode, aircraft.FieldManufacturer, aircraft.FieldManufacturerDesignator, aircraft.FieldCommonDesignation, aircraft.FieldCommonName, aircraft.FieldDefaultValues, aircraft.FieldMaximumValues, aircraft.FieldFuelDetails, aircraft.FieldOilDetails:
			values[i] = new(sql.NullString)
		case aircraft.FieldID, aircraft.FieldCompanyID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Aircraft", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Aircraft fields.
func (a *Aircraft) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case aircraft.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case aircraft.FieldCompanyID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field company_id", values[i])
			} else if value != nil {
				a.CompanyID = *value
			}
		case aircraft.FieldCurrentFlightHours:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field current_flight_hours", values[i])
			} else if value.Valid {
				a.CurrentFlightHours = new(float32)
				*a.CurrentFlightHours = float32(value.Float64)
			}
		case aircraft.FieldCurrentCycles:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field current_cycles", values[i])
			} else if value.Valid {
				a.CurrentCycles = new(int)
				*a.CurrentCycles = int(value.Int64)
			}
		case aircraft.FieldAircraftRegistration:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field aircraft_registration", values[i])
			} else if value.Valid {
				a.AircraftRegistration = value.String
			}
		case aircraft.FieldBaseAirportCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field base_airport_code", values[i])
			} else if value.Valid {
				a.BaseAirportCode = value.String
			}
		case aircraft.FieldManufacturer:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field manufacturer", values[i])
			} else if value.Valid {
				a.Manufacturer = new(string)
				*a.Manufacturer = value.String
			}
		case aircraft.FieldManufacturerDesignator:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field manufacturer_designator", values[i])
			} else if value.Valid {
				a.ManufacturerDesignator = value.String
			}
		case aircraft.FieldCommonDesignation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field common_designation", values[i])
			} else if value.Valid {
				a.CommonDesignation = new(string)
				*a.CommonDesignation = value.String
			}
		case aircraft.FieldCommonName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field common_name", values[i])
			} else if value.Valid {
				a.CommonName = value.String
			}
		case aircraft.FieldPilotsRequiredToFly:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field pilots_required_to_fly", values[i])
			} else if value.Valid {
				a.PilotsRequiredToFly = new(int)
				*a.PilotsRequiredToFly = int(value.Int64)
			}
		case aircraft.FieldDefaultValues:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field default_values", values[i])
			} else if value.Valid {
				a.DefaultValues = value.String
			}
		case aircraft.FieldMaximumValues:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field maximum_values", values[i])
			} else if value.Valid {
				a.MaximumValues = value.String
			}
		case aircraft.FieldCurrentLandings:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field current_landings", values[i])
			} else if value.Valid {
				a.CurrentLandings = new(int)
				*a.CurrentLandings = int(value.Int64)
			}
		case aircraft.FieldFuelDetails:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field fuel_details", values[i])
			} else if value.Valid {
				a.FuelDetails = value.String
			}
		case aircraft.FieldOilDetails:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field oil_details", values[i])
			} else if value.Valid {
				a.OilDetails = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Aircraft.
// Note that you need to call Aircraft.Unwrap() before calling this method if this Aircraft
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Aircraft) Update() *AircraftUpdateOne {
	return (&AircraftClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Aircraft entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Aircraft) Unwrap() *Aircraft {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Aircraft is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Aircraft) String() string {
	var builder strings.Builder
	builder.WriteString("Aircraft(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", company_id=")
	builder.WriteString(fmt.Sprintf("%v", a.CompanyID))
	if v := a.CurrentFlightHours; v != nil {
		builder.WriteString(", current_flight_hours=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := a.CurrentCycles; v != nil {
		builder.WriteString(", current_cycles=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", aircraft_registration=")
	builder.WriteString(a.AircraftRegistration)
	builder.WriteString(", base_airport_code=")
	builder.WriteString(a.BaseAirportCode)
	if v := a.Manufacturer; v != nil {
		builder.WriteString(", manufacturer=")
		builder.WriteString(*v)
	}
	builder.WriteString(", manufacturer_designator=")
	builder.WriteString(a.ManufacturerDesignator)
	if v := a.CommonDesignation; v != nil {
		builder.WriteString(", common_designation=")
		builder.WriteString(*v)
	}
	builder.WriteString(", common_name=")
	builder.WriteString(a.CommonName)
	if v := a.PilotsRequiredToFly; v != nil {
		builder.WriteString(", pilots_required_to_fly=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", default_values=")
	builder.WriteString(a.DefaultValues)
	builder.WriteString(", maximum_values=")
	builder.WriteString(a.MaximumValues)
	if v := a.CurrentLandings; v != nil {
		builder.WriteString(", current_landings=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", fuel_details=")
	builder.WriteString(a.FuelDetails)
	builder.WriteString(", oil_details=")
	builder.WriteString(a.OilDetails)
	builder.WriteByte(')')
	return builder.String()
}

// Aircrafts is a parsable slice of Aircraft.
type Aircrafts []*Aircraft

func (a Aircrafts) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}