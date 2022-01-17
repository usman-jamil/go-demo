package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Aircraft holds the schema definition for the Aircraft entity.
type Aircraft struct {
	ent.Schema
}

// Fields of the Aircraft.
func (Aircraft) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique().
			Default(uuid.New).
			Immutable().
			StorageKey("aircraft_id"),
		field.UUID("company_id", uuid.UUID{}).
			Default(uuid.New),
		field.Float32("current_flight_hours").
			Optional(),
		field.Int("current_cycles").
			Optional(),
		field.String("aircraft_registration").
			Optional(),
		field.String("base_airport_code").
			Optional(),
		field.String("manufacturer").
			Optional(),
		field.String("manufacturer_designator").
			Optional(),
		field.String("common_designation").
			Optional(),
		field.String("common_name").
			Optional(),
		field.Int("pilots_required_to_fly").
			Optional(),
		field.String("default_values").
			Optional(),
		field.String("maximum_values").
			Optional(),
		field.Int("current_landings").
			Optional(),
		field.String("fuel_details").
			Optional(),
		field.String("oil_details").
			Optional(),
	}
}

// Edges of the Aircraft.
func (Aircraft) Edges() []ent.Edge {
	return nil
}

func (Aircraft) Indexes() []ent.Index {
	return []ent.Index{
		// non-unique index.
		index.Fields("current_flight_hours"),
		index.Fields("aircraft_registration"),
		index.Fields("common_designation"),
	}
}
