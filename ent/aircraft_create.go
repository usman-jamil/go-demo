// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"enttest/ent/aircraft"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AircraftCreate is the builder for creating a Aircraft entity.
type AircraftCreate struct {
	config
	mutation *AircraftMutation
	hooks    []Hook
}

// SetCompanyID sets the "company_id" field.
func (ac *AircraftCreate) SetCompanyID(u uuid.UUID) *AircraftCreate {
	ac.mutation.SetCompanyID(u)
	return ac
}

// SetCurrentFlightHours sets the "current_flight_hours" field.
func (ac *AircraftCreate) SetCurrentFlightHours(f float32) *AircraftCreate {
	ac.mutation.SetCurrentFlightHours(f)
	return ac
}

// SetNillableCurrentFlightHours sets the "current_flight_hours" field if the given value is not nil.
func (ac *AircraftCreate) SetNillableCurrentFlightHours(f *float32) *AircraftCreate {
	if f != nil {
		ac.SetCurrentFlightHours(*f)
	}
	return ac
}

// SetCurrentCycles sets the "current_cycles" field.
func (ac *AircraftCreate) SetCurrentCycles(i int) *AircraftCreate {
	ac.mutation.SetCurrentCycles(i)
	return ac
}

// SetNillableCurrentCycles sets the "current_cycles" field if the given value is not nil.
func (ac *AircraftCreate) SetNillableCurrentCycles(i *int) *AircraftCreate {
	if i != nil {
		ac.SetCurrentCycles(*i)
	}
	return ac
}

// SetAircraftRegistration sets the "aircraft_registration" field.
func (ac *AircraftCreate) SetAircraftRegistration(s string) *AircraftCreate {
	ac.mutation.SetAircraftRegistration(s)
	return ac
}

// SetNillableAircraftRegistration sets the "aircraft_registration" field if the given value is not nil.
func (ac *AircraftCreate) SetNillableAircraftRegistration(s *string) *AircraftCreate {
	if s != nil {
		ac.SetAircraftRegistration(*s)
	}
	return ac
}

// SetBaseAirportCode sets the "base_airport_code" field.
func (ac *AircraftCreate) SetBaseAirportCode(s string) *AircraftCreate {
	ac.mutation.SetBaseAirportCode(s)
	return ac
}

// SetNillableBaseAirportCode sets the "base_airport_code" field if the given value is not nil.
func (ac *AircraftCreate) SetNillableBaseAirportCode(s *string) *AircraftCreate {
	if s != nil {
		ac.SetBaseAirportCode(*s)
	}
	return ac
}

// SetManufacturer sets the "manufacturer" field.
func (ac *AircraftCreate) SetManufacturer(s string) *AircraftCreate {
	ac.mutation.SetManufacturer(s)
	return ac
}

// SetNillableManufacturer sets the "manufacturer" field if the given value is not nil.
func (ac *AircraftCreate) SetNillableManufacturer(s *string) *AircraftCreate {
	if s != nil {
		ac.SetManufacturer(*s)
	}
	return ac
}

// SetManufacturerDesignator sets the "manufacturer_designator" field.
func (ac *AircraftCreate) SetManufacturerDesignator(s string) *AircraftCreate {
	ac.mutation.SetManufacturerDesignator(s)
	return ac
}

// SetNillableManufacturerDesignator sets the "manufacturer_designator" field if the given value is not nil.
func (ac *AircraftCreate) SetNillableManufacturerDesignator(s *string) *AircraftCreate {
	if s != nil {
		ac.SetManufacturerDesignator(*s)
	}
	return ac
}

// SetCommonDesignation sets the "common_designation" field.
func (ac *AircraftCreate) SetCommonDesignation(s string) *AircraftCreate {
	ac.mutation.SetCommonDesignation(s)
	return ac
}

// SetNillableCommonDesignation sets the "common_designation" field if the given value is not nil.
func (ac *AircraftCreate) SetNillableCommonDesignation(s *string) *AircraftCreate {
	if s != nil {
		ac.SetCommonDesignation(*s)
	}
	return ac
}

// SetCommonName sets the "common_name" field.
func (ac *AircraftCreate) SetCommonName(s string) *AircraftCreate {
	ac.mutation.SetCommonName(s)
	return ac
}

// SetNillableCommonName sets the "common_name" field if the given value is not nil.
func (ac *AircraftCreate) SetNillableCommonName(s *string) *AircraftCreate {
	if s != nil {
		ac.SetCommonName(*s)
	}
	return ac
}

// SetPilotsRequiredToFly sets the "pilots_required_to_fly" field.
func (ac *AircraftCreate) SetPilotsRequiredToFly(i int) *AircraftCreate {
	ac.mutation.SetPilotsRequiredToFly(i)
	return ac
}

// SetNillablePilotsRequiredToFly sets the "pilots_required_to_fly" field if the given value is not nil.
func (ac *AircraftCreate) SetNillablePilotsRequiredToFly(i *int) *AircraftCreate {
	if i != nil {
		ac.SetPilotsRequiredToFly(*i)
	}
	return ac
}

// SetDefaultValues sets the "default_values" field.
func (ac *AircraftCreate) SetDefaultValues(s string) *AircraftCreate {
	ac.mutation.SetDefaultValues(s)
	return ac
}

// SetNillableDefaultValues sets the "default_values" field if the given value is not nil.
func (ac *AircraftCreate) SetNillableDefaultValues(s *string) *AircraftCreate {
	if s != nil {
		ac.SetDefaultValues(*s)
	}
	return ac
}

// SetMaximumValues sets the "maximum_values" field.
func (ac *AircraftCreate) SetMaximumValues(s string) *AircraftCreate {
	ac.mutation.SetMaximumValues(s)
	return ac
}

// SetNillableMaximumValues sets the "maximum_values" field if the given value is not nil.
func (ac *AircraftCreate) SetNillableMaximumValues(s *string) *AircraftCreate {
	if s != nil {
		ac.SetMaximumValues(*s)
	}
	return ac
}

// SetCurrentLandings sets the "current_landings" field.
func (ac *AircraftCreate) SetCurrentLandings(i int) *AircraftCreate {
	ac.mutation.SetCurrentLandings(i)
	return ac
}

// SetNillableCurrentLandings sets the "current_landings" field if the given value is not nil.
func (ac *AircraftCreate) SetNillableCurrentLandings(i *int) *AircraftCreate {
	if i != nil {
		ac.SetCurrentLandings(*i)
	}
	return ac
}

// SetFuelDetails sets the "fuel_details" field.
func (ac *AircraftCreate) SetFuelDetails(s string) *AircraftCreate {
	ac.mutation.SetFuelDetails(s)
	return ac
}

// SetNillableFuelDetails sets the "fuel_details" field if the given value is not nil.
func (ac *AircraftCreate) SetNillableFuelDetails(s *string) *AircraftCreate {
	if s != nil {
		ac.SetFuelDetails(*s)
	}
	return ac
}

// SetOilDetails sets the "oil_details" field.
func (ac *AircraftCreate) SetOilDetails(s string) *AircraftCreate {
	ac.mutation.SetOilDetails(s)
	return ac
}

// SetNillableOilDetails sets the "oil_details" field if the given value is not nil.
func (ac *AircraftCreate) SetNillableOilDetails(s *string) *AircraftCreate {
	if s != nil {
		ac.SetOilDetails(*s)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AircraftCreate) SetID(u uuid.UUID) *AircraftCreate {
	ac.mutation.SetID(u)
	return ac
}

// Mutation returns the AircraftMutation object of the builder.
func (ac *AircraftCreate) Mutation() *AircraftMutation {
	return ac.mutation
}

// Save creates the Aircraft in the database.
func (ac *AircraftCreate) Save(ctx context.Context) (*Aircraft, error) {
	var (
		err  error
		node *Aircraft
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AircraftMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AircraftCreate) SaveX(ctx context.Context) *Aircraft {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AircraftCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AircraftCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AircraftCreate) defaults() {
	if _, ok := ac.mutation.CompanyID(); !ok {
		v := aircraft.DefaultCompanyID()
		ac.mutation.SetCompanyID(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := aircraft.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AircraftCreate) check() error {
	if _, ok := ac.mutation.CompanyID(); !ok {
		return &ValidationError{Name: "company_id", err: errors.New(`ent: missing required field "company_id"`)}
	}
	return nil
}

func (ac *AircraftCreate) sqlSave(ctx context.Context) (*Aircraft, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (ac *AircraftCreate) createSpec() (*Aircraft, *sqlgraph.CreateSpec) {
	var (
		_node = &Aircraft{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: aircraft.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: aircraft.FieldID,
			},
		}
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.CompanyID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: aircraft.FieldCompanyID,
		})
		_node.CompanyID = value
	}
	if value, ok := ac.mutation.CurrentFlightHours(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: aircraft.FieldCurrentFlightHours,
		})
		_node.CurrentFlightHours = value
	}
	if value, ok := ac.mutation.CurrentCycles(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: aircraft.FieldCurrentCycles,
		})
		_node.CurrentCycles = value
	}
	if value, ok := ac.mutation.AircraftRegistration(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: aircraft.FieldAircraftRegistration,
		})
		_node.AircraftRegistration = value
	}
	if value, ok := ac.mutation.BaseAirportCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: aircraft.FieldBaseAirportCode,
		})
		_node.BaseAirportCode = value
	}
	if value, ok := ac.mutation.Manufacturer(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: aircraft.FieldManufacturer,
		})
		_node.Manufacturer = value
	}
	if value, ok := ac.mutation.ManufacturerDesignator(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: aircraft.FieldManufacturerDesignator,
		})
		_node.ManufacturerDesignator = value
	}
	if value, ok := ac.mutation.CommonDesignation(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: aircraft.FieldCommonDesignation,
		})
		_node.CommonDesignation = value
	}
	if value, ok := ac.mutation.CommonName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: aircraft.FieldCommonName,
		})
		_node.CommonName = value
	}
	if value, ok := ac.mutation.PilotsRequiredToFly(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: aircraft.FieldPilotsRequiredToFly,
		})
		_node.PilotsRequiredToFly = value
	}
	if value, ok := ac.mutation.DefaultValues(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: aircraft.FieldDefaultValues,
		})
		_node.DefaultValues = value
	}
	if value, ok := ac.mutation.MaximumValues(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: aircraft.FieldMaximumValues,
		})
		_node.MaximumValues = value
	}
	if value, ok := ac.mutation.CurrentLandings(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: aircraft.FieldCurrentLandings,
		})
		_node.CurrentLandings = value
	}
	if value, ok := ac.mutation.FuelDetails(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: aircraft.FieldFuelDetails,
		})
		_node.FuelDetails = value
	}
	if value, ok := ac.mutation.OilDetails(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: aircraft.FieldOilDetails,
		})
		_node.OilDetails = value
	}
	return _node, _spec
}

// AircraftCreateBulk is the builder for creating many Aircraft entities in bulk.
type AircraftCreateBulk struct {
	config
	builders []*AircraftCreate
}

// Save creates the Aircraft entities in the database.
func (acb *AircraftCreateBulk) Save(ctx context.Context) ([]*Aircraft, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Aircraft, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AircraftMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AircraftCreateBulk) SaveX(ctx context.Context) []*Aircraft {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AircraftCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AircraftCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
