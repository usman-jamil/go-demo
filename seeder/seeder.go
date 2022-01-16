package seeder

import (
	"context"
	"enttest/ent"
	"fmt"
	"github.com/google/uuid"
)

func stringToUUID(newUUID string) (*uuid.UUID, error) {
	id, err := uuid.Parse(newUUID)

	if err != nil {
		return nil, fmt.Errorf("failed to convert string to []byte: %w", err)
	}

	return &id, nil
}

func getUUID(newUUID string) uuid.UUID {
	id, _ := stringToUUID(newUUID)

	return *id
}

func GetSeedData(ctx context.Context,
	client *ent.Client) []*ent.AircraftCreate {
	return []*ent.AircraftCreate{
		client.Aircraft.
			Create().
			SetID(getUUID("35b33702-53b4-40d1-85d0-191df3dd742e")).
			SetCompanyID(getUUID("b144a41a-e5de-4032-b776-24ca87536706")).
			SetCurrentFlightHours(8653.7).
			SetCurrentCycles(220).
			SetAircraftRegistration("C-FSLR").
			SetBaseAirportCode("CEZ4").
			SetManufacturer("CESSNA").
			SetManufacturerDesignator("U206G").
			SetCommonDesignation("C206").
			SetCommonName("C206").
			SetDefaultValues("{\"ClimbRate\": 0, \"FlightType\": \"charter\", \"FlightRules\": \"VFR/703\", \"FlightNumber\": \"CWA604\", \"TrueAirSpeed\": 110, \"ApproachSpeed\": 0, \"TaxiFuelWeight\": 10, \"ConfigurationID\": \"6f5cfa2b-461a-4a49-bc59-0748128e756a\", \"AscentFuelFlowRate\": 105, \"CruiseFuelFlowRate\": 100}").
			SetMaximumValues("{\"Altitude\": 14800, \"RampWeight\": 3610, \"GroundSpeed\": 174, \"LandingWeight\": 3600, \"TakeoffWeight\": 3600, \"ZeroFuelWeight\": 3600, \"TotalFuelWeight\": 0}").
			SetCurrentLandings(211).
			SetFuelDetails("{\"Type\": \"100LL\", \"CapacityLitres\": null, \"CapacityGallons\": 61}").
			SetOilDetails("{\"Grade\": \"AEROSHELL W-80\", \"CapacityLitres\": null, \"CapacityGallons\": 3}"),
		client.Aircraft.
			Create().
			SetID(getUUID("7461c815-2591-4fb5-bb7d-182daa12b05c")).
			SetCompanyID(getUUID("b144a41a-e5de-4032-b776-24ca87536707")).
			SetCurrentFlightHours(16065.6).
			SetCurrentCycles(1359).
			SetAircraftRegistration("C-GWVT").
			SetBaseAirportCode("CEZ4").
			SetManufacturer("CESSNA").
			SetManufacturerDesignator("C206").
			SetCommonDesignation("C206").
			SetCommonName("C206").
			SetDefaultValues("{\"ClimbRate\": 0, \"FlightType\": \"charter\", \"FlightRules\": \"VFR/703\", \"FlightNumber\": \"CWA602\", \"TrueAirSpeed\": 110, \"ApproachSpeed\": 0, \"TaxiFuelWeight\": 10, \"ConfigurationID\": \"6aa7fa21-4b71-477d-84fb-1b90aac8116a\", \"AscentFuelFlowRate\": 105, \"CruiseFuelFlowRate\": 100}").
			SetMaximumValues("{\"Altitude\": 14800, \"RampWeight\": 3610, \"GroundSpeed\": 174, \"LandingWeight\": 3600, \"TakeoffWeight\": 3600, \"ZeroFuelWeight\": 3600, \"TotalFuelWeight\": 0}").
			SetCurrentLandings(1359).
			SetFuelDetails("{\"Type\": \"100LL\", \"CapacityLitres\": null, \"CapacityGallons\": 61}").
			SetOilDetails("{\"Grade\": \"AEROSHELL W-80\", \"CapacityLitres\": null, \"CapacityGallons\": 3}"),
		client.Aircraft.
			Create().
			SetID(getUUID("004564aa-8a4d-4a7b-ba10-d328aa04cf85")).
			SetCompanyID(getUUID("b144a41a-e5de-4032-b776-24ca87536708")).
			SetCurrentFlightHours(6580.4).
			SetCurrentCycles(14169).
			SetAircraftRegistration("C-GLRT").
			SetBaseAirportCode("CEZ4").
			SetManufacturer("CESSNA").
			SetManufacturerDesignator("C208B").
			SetCommonDesignation("208B").
			SetCommonName("C 208B").
			SetDefaultValues("{\"ClimbRate\": 0, \"FlightType\": \"charter\", \"FlightRules\": \"VFR/703\", \"FlightNumber\": \"CWA802\", \"TrueAirSpeed\": 145, \"ApproachSpeed\": 0, \"TaxiFuelWeight\": 35, \"ConfigurationID\": \"04385841-3954-4d57-9ef6-496c4f388912\", \"AscentFuelFlowRate\": 400, \"CruiseFuelFlowRate\": 360}").
			SetMaximumValues("{\"Altitude\": 25000, \"RampWeight\": 9097, \"GroundSpeed\": 186, \"LandingWeight\": 9000, \"TakeoffWeight\": 9062, \"ZeroFuelWeight\": 9062, \"TotalFuelWeight\": 0}").
			SetCurrentLandings(14169).
			SetFuelDetails("{\"Type\": \"Jet A/A1 - See AFM for alternates\", \"CapacityLitres\": null, \"CapacityGallons\": 335}").
			SetOilDetails("{\"Grade\": \"AEROSHELL 500 Turbine Oil\", \"CapacityLitres\": null, \"CapacityGallons\": 3.5}"),
		client.Aircraft.
			Create().
			SetID(getUUID("4957c4b7-329a-4b13-a003-be930c6ebe24")).
			SetCompanyID(getUUID("b144a41a-e5de-4032-b776-24ca87536709")).
			SetCurrentFlightHours(15873.8).
			SetCurrentCycles(1409).
			SetAircraftRegistration("C-GICJ").
			SetBaseAirportCode("CEZ4").
			SetManufacturer("CESSNA").
			SetManufacturerDesignator("U206F").
			SetCommonDesignation("C206").
			SetCommonName("C206").
			SetDefaultValues("{\"ClimbRate\": 0, \"FlightType\": \"charter\", \"FlightRules\": \"VFR/703\", \"FlightNumber\": \"CWA603\", \"TrueAirSpeed\": 110, \"ApproachSpeed\": 0, \"TaxiFuelWeight\": 10, \"ConfigurationID\": \"41e355bd-c1b2-4627-95f1-0f76c8a17dc8\", \"AscentFuelFlowRate\": 105, \"CruiseFuelFlowRate\": 100}").
			SetMaximumValues("{\"Altitude\": 14800, \"RampWeight\": 3610, \"GroundSpeed\": 174, \"LandingWeight\": 3600, \"TakeoffWeight\": 3600, \"ZeroFuelWeight\": 3600, \"TotalFuelWeight\": 0}").
			SetCurrentLandings(1409).
			SetFuelDetails("{\"Type\": \"100LL\", \"CapacityLitres\": null, \"CapacityGallons\": 80}").
			SetOilDetails("{\"Grade\": \"AEROSHELL W-80\", \"CapacityLitres\": null, \"CapacityGallons\": 3}"),
		client.Aircraft.
			Create().
			SetID(getUUID("17097f1f-cfdd-46bf-a1dd-50194da25c86")).
			SetCompanyID(getUUID("b144a41a-e5de-4032-b776-24ca87536710")).
			SetCurrentFlightHours(22018.6).
			SetCurrentCycles(1073).
			SetAircraftRegistration("C-GGBZ").
			SetBaseAirportCode("CEZ4").
			SetManufacturer("CESSNA").
			SetManufacturerDesignator("U206G").
			SetCommonDesignation("C206").
			SetCommonName("C206").
			SetDefaultValues("{\"ClimbRate\": 0, \"FlightType\": \"charter\", \"FlightRules\": \"VFR/703\", \"FlightNumber\": \"CWA605\", \"TrueAirSpeed\": 110, \"ApproachSpeed\": 0, \"TaxiFuelWeight\": 10, \"ConfigurationID\": \"74f3e4fd-9e94-4a9d-be8f-fc1c4931efa4\", \"AscentFuelFlowRate\": 105, \"CruiseFuelFlowRate\": 100}").
			SetMaximumValues("{\"Altitude\": 14800, \"RampWeight\": 3610, \"GroundSpeed\": 174, \"LandingWeight\": 3600, \"TakeoffWeight\": 3600, \"ZeroFuelWeight\": 3600, \"TotalFuelWeight\": 0}").
			SetCurrentLandings(902).
			SetFuelDetails("{\"Type\": \"100LL\", \"CapacityLitres\": null, \"CapacityGallons\": 92}").
			SetOilDetails("{\"Grade\": \"AEROSHELL W-80\", \"CapacityLitres\": null, \"CapacityGallons\": 3}"),
		client.Aircraft.
			Create().
			SetID(getUUID("efc72293-55a2-497b-b36d-a9c93ba590e9")).
			SetCompanyID(getUUID("b144a41a-e5de-4032-b776-24ca87536711")).
			SetAircraftRegistration("C-GPNB").
			SetBaseAirportCode("CYZH").
			SetManufacturer("Beechcraft").
			SetManufacturerDesignator("BE20").
			SetCommonName("King Air 200").
			SetDefaultValues("{\"ClimbRate\": 0, \"FlightType\": \"charter\", \"FlightRules\": \"IFR/703\", \"FlightNumber\": \"CWA911\", \"TrueAirSpeed\": 245, \"ApproachSpeed\": 0, \"TaxiFuelWeight\": 90, \"ConfigurationID\": \"7e3457cd-6e19-43e0-9688-d29c659fedcb\", \"AscentFuelFlowRate\": 750, \"CruiseFuelFlowRate\": 700}").
			SetMaximumValues("{\"Altitude\": 35000, \"RampWeight\": 12590, \"GroundSpeed\": 270, \"LandingWeight\": 0, \"TakeoffWeight\": 12500, \"ZeroFuelWeight\": 11000, \"TotalFuelWeight\": 0}"),
		client.Aircraft.
			Create().
			SetID(getUUID("961a486d-937c-4434-b4c2-41a9390db818")).
			SetCompanyID(getUUID("b144a41a-e5de-4032-b776-24ca87536712")).
			SetAircraftRegistration("C-FKCW").
			SetBaseAirportCode("CYEG").
			SetManufacturer("Beechcraft").
			SetManufacturerDesignator("BE20").
			SetCommonName("King Air 200").
			SetDefaultValues("{\"ClimbRate\": 0, \"FlightType\": \"charter\", \"FlightRules\": \"IFR/703\", \"FlightNumber\": \"CWA917\", \"TrueAirSpeed\": 245, \"ApproachSpeed\": 0, \"TaxiFuelWeight\": 90, \"ConfigurationID\": \"852c3b05-7f24-45de-80b8-28dd21b83190\", \"AscentFuelFlowRate\": 750, \"CruiseFuelFlowRate\": 700}").
			SetMaximumValues("{\"Altitude\": 35000, \"RampWeight\": 12590, \"GroundSpeed\": 270, \"LandingWeight\": 0, \"TakeoffWeight\": 12500, \"ZeroFuelWeight\": 11000, \"TotalFuelWeight\": 0}"),
		client.Aircraft.
			Create().
			SetID(getUUID("5a77d53b-7c2d-4704-8c9a-4d75b9b00bd7")).
			SetCompanyID(getUUID("b144a41a-e5de-4032-b776-24ca87536713")).
			SetCurrentFlightHours(12343.5).
			SetCurrentCycles(25372).
			SetAircraftRegistration("GGUH-2").
			SetBaseAirportCode("CEZ4").
			SetManufacturer("CESSNA").
			SetManufacturerDesignator("C208B").
			SetCommonDesignation("208B").
			SetCommonName("C 208B").
			SetDefaultValues("{\"ClimbRate\": 0, \"FlightType\": \"charter\", \"FlightRules\": \"VFR/703\", \"FlightNumber\": \"CWA801\", \"TrueAirSpeed\": 145, \"ApproachSpeed\": 0, \"TaxiFuelWeight\": 35, \"ConfigurationID\": \"1c9dfd76-1cc2-4f89-ae89-3d5d753f867c\", \"AscentFuelFlowRate\": 400, \"CruiseFuelFlowRate\": 360}").
			SetMaximumValues("{\"Altitude\": 25000, \"RampWeight\": 9097, \"GroundSpeed\": 186, \"LandingWeight\": 9000, \"TakeoffWeight\": 9062, \"ZeroFuelWeight\": 9062, \"TotalFuelWeight\": 0}").
			SetCurrentLandings(25372).
			SetFuelDetails("{\"Type\": \"Jet A/A1 - See AFM for alternates\", \"CapacityLitres\": null, \"CapacityGallons\": 335}").
			SetOilDetails("{\"Grade\": \"AEROSHELL 500 Turbine Oil\", \"CapacityLitres\": null, \"CapacityGallons\": 3.5}"),
		client.Aircraft.
			Create().
			SetID(getUUID("a0d5cec9-fbf4-4c06-9346-dc488172b851")).
			SetCompanyID(getUUID("b144a41a-e5de-4032-b776-24ca87536714")).
			SetCurrentFlightHours(14596.5).
			SetCurrentCycles(558).
			SetAircraftRegistration("C-GLGD").
			SetBaseAirportCode("CEZ4").
			SetManufacturer("CESSNA").
			SetManufacturerDesignator("U206G").
			SetCommonDesignation("C206").
			SetCommonName("C206").
			SetPilotsRequiredToFly(1).
			SetDefaultValues("{\"ClimbRate\": 0, \"FlightType\": \"charter\", \"FlightRules\": \"VFR/703\", \"FlightNumber\": \"CWA606\", \"TrueAirSpeed\": 110, \"ApproachSpeed\": 0, \"TaxiFuelWeight\": 10, \"ConfigurationID\": \"4f781166-e73d-4b7b-ba10-f89a51001502\", \"AscentFuelFlowRate\": 105, \"CruiseFuelFlowRate\": 100}").
			SetMaximumValues("{\"Altitude\": 14800, \"RampWeight\": 3610, \"GroundSpeed\": 174, \"LandingWeight\": 3600, \"TakeoffWeight\": 3600, \"ZeroFuelWeight\": 3600, \"TotalFuelWeight\": 0}").
			SetCurrentLandings(525).
			SetFuelDetails("{\"Type\": \"100LL\", \"CapacityLitres\": null, \"CapacityGallons\": 92}").
			SetOilDetails("{\"Grade\": \"AEROSHELL W-80\", \"CapacityLitres\": null, \"CapacityGallons\": 3}"),
		client.Aircraft.
			Create().
			SetID(getUUID("310f4cb2-1757-47e5-bda0-792908ef1411")).
			SetCompanyID(getUUID("b144a41a-e5de-4032-b776-24ca87536715")).
			SetCurrentFlightHours(3170.6).
			SetCurrentCycles(4215).
			SetAircraftRegistration("C-GLUH").
			SetBaseAirportCode("CYQU").
			SetManufacturer("TEXTRON AVIATION").
			SetManufacturerDesignator("B200GT").
			SetCommonDesignation("B200GT").
			SetCommonName("Super King Air 250").
			SetDefaultValues("{\"ClimbRate\": 0, \"FlightType\": \"medevac\", \"FlightRules\": \"IFR/703\", \"FlightNumber\": \"CWA924\", \"TrueAirSpeed\": 245, \"ApproachSpeed\": 0, \"TaxiFuelWeight\": 90, \"ConfigurationID\": \"0397dc84-6ae5-4e48-aaf7-9c5764585c95\", \"AscentFuelFlowRate\": 750, \"CruiseFuelFlowRate\": 700}").
			SetMaximumValues("{\"Altitude\": 35000, \"RampWeight\": 13510, \"GroundSpeed\": 310, \"LandingWeight\": 12500, \"TakeoffWeight\": 13420, \"ZeroFuelWeight\": 11500, \"TotalFuelWeight\": 0}").
			SetCurrentLandings(4215).
			SetFuelDetails("{\"Type\": \"Jet A/A1 - See AFM for alternates\", \"CapacityLitres\": null, \"CapacityGallons\": 549}").
			SetOilDetails("{\"Grade\": \"AEROSHELL 500 Turbine Oil\", \"CapacityLitres\": null, \"CapacityGallons\": 2.5}"),
		client.Aircraft.
			Create().
			SetID(getUUID("81a54951-5aa7-45d2-8a06-29c1c3a7825e")).
			SetCompanyID(getUUID("b144a41a-e5de-4032-b776-24ca87536716")).
			SetCurrentFlightHours(3793.3).
			SetCurrentCycles(3770).
			SetAircraftRegistration("C-GLUQ").
			SetBaseAirportCode("CYOJ").
			SetManufacturer("TEXTRON AVIATION").
			SetManufacturerDesignator("B200GT").
			SetCommonDesignation("B200GT").
			SetCommonName("Super King Air 250 Halo").
			SetDefaultValues("{\"ClimbRate\": 0, \"FlightType\": \"medevac\", \"FlightRules\": \"IFR/703\", \"FlightNumber\": \"CWA926\", \"TrueAirSpeed\": 245, \"ApproachSpeed\": 0, \"TaxiFuelWeight\": 90, \"ConfigurationID\": \"59137f23-c519-4a01-8945-be28947c842a\", \"AscentFuelFlowRate\": 750, \"CruiseFuelFlowRate\": 650}").
			SetMaximumValues("{\"Altitude\": 35000, \"RampWeight\": 13510, \"GroundSpeed\": 310, \"LandingWeight\": 12500, \"TakeoffWeight\": 13420, \"ZeroFuelWeight\": 11500, \"TotalFuelWeight\": 0}").
			SetCurrentLandings(3768).
			SetFuelDetails("{\"Type\": \"Jet A/A1 - See AFM for alternates\", \"CapacityLitres\": null, \"CapacityGallons\": 549}").
			SetOilDetails("{\"Grade\": \"AEROSHELL 500 Turbine Oil\", \"CapacityLitres\": null, \"CapacityGallons\": 2.5}"),
		client.Aircraft.
			Create().
			SetID(getUUID("b9053dd0-7ccd-43e2-9ac6-30066e1f529b")).
			SetCompanyID(getUUID("b144a41a-e5de-4032-b776-24ca87536717")).
			SetCurrentFlightHours(4075.2).
			SetCurrentCycles(4850).
			SetAircraftRegistration("C-GLUD").
			SetBaseAirportCode("CYEG").
			SetManufacturer("TEXTRON AVIATION").
			SetManufacturerDesignator("B300C").
			SetCommonDesignation("B300").
			SetCommonName("Super King Air 350C").
			SetDefaultValues("{\"ClimbRate\": 0, \"FlightType\": \"medevac\", \"FlightRules\": \"IFR/703\", \"FlightNumber\": \"CWA935\", \"TrueAirSpeed\": 245, \"ApproachSpeed\": 0, \"TaxiFuelWeight\": 100, \"ConfigurationID\": \"2e70ac6f-ef67-43e7-bb8e-fad5a753535f\", \"AscentFuelFlowRate\": 800, \"CruiseFuelFlowRate\": 750}").
			SetMaximumValues("{\"Altitude\": 35000, \"RampWeight\": 16600, \"GroundSpeed\": 312, \"LandingWeight\": 15675, \"TakeoffWeight\": 16500, \"ZeroFuelWeight\": 13000, \"TotalFuelWeight\": 0}").
			SetCurrentLandings(1883).
			SetFuelDetails("{\"Type\": \"Jet A/A1 - See AFM for alternates\", \"CapacityLitres\": null, \"CapacityGallons\": 546}").
			SetOilDetails("{\"Grade\": \"AEROSHELL 500 Turbine Oil\", \"CapacityLitres\": null, \"CapacityGallons\": 2.5}"),
		client.Aircraft.
			Create().
			SetID(getUUID("eb07d86a-1b89-42fb-80a2-52101a988735")).
			SetCompanyID(getUUID("b144a41a-e5de-4032-b776-24ca87536718")).
			SetCurrentFlightHours(3797.5).
			SetCurrentCycles(4728).
			SetAircraftRegistration("C-GLUF").
			SetBaseAirportCode("CYYC").
			SetManufacturer("TEXTRON AVIATION").
			SetManufacturerDesignator("B200CGT").
			SetCommonDesignation("B200GT").
			SetCommonName("Super King Air 250 Halo").
			SetDefaultValues("{\"ClimbRate\": 0, \"FlightType\": \"medevac\", \"FlightRules\": \"IFR/703\", \"FlightNumber\": \"CWA922\", \"TrueAirSpeed\": 245, \"ApproachSpeed\": 0, \"TaxiFuelWeight\": 90, \"ConfigurationID\": \"c2154e54-1947-44a7-bc5a-ed1d92571706\", \"AscentFuelFlowRate\": 750, \"CruiseFuelFlowRate\": 700}").
			SetMaximumValues("{\"Altitude\": 35000, \"RampWeight\": 14090, \"GroundSpeed\": 310, \"LandingWeight\": 13500, \"TakeoffWeight\": 14000, \"ZeroFuelWeight\": 11500, \"TotalFuelWeight\": 0}").
			SetCurrentLandings(4728).
			SetFuelDetails("{\"Type\": \"Jet A/A1 - See AFM for alternates\", \"CapacityLitres\": null, \"CapacityGallons\": 549}").
			SetOilDetails("{\"Grade\": \"AEROSHELL 500 Turbine Oil\", \"CapacityLitres\": null, \"CapacityGallons\": 2.5}"),
		client.Aircraft.
			Create().
			SetID(getUUID("86cf375d-3edf-43df-b16f-cb3155d959fc")).
			SetCompanyID(getUUID("b144a41a-e5de-4032-b776-24ca87536719")).
			SetCurrentFlightHours(3591.1).
			SetCurrentCycles(5212).
			SetAircraftRegistration("C-GLUL").
			SetBaseAirportCode("CYXH").
			SetManufacturer("TEXTRON AVIATION").
			SetManufacturerDesignator("B200GT").
			SetCommonDesignation("B200GT").
			SetCommonName("Super King Air 250 Halo").
			SetDefaultValues("{\"ClimbRate\": 0, \"FlightType\": \"medevac\", \"FlightRules\": \"IFR/703\", \"FlightNumber\": \"CWA927\", \"TrueAirSpeed\": 245, \"ApproachSpeed\": 0, \"TaxiFuelWeight\": 90, \"ConfigurationID\": \"eac94031-36d1-48c1-b81a-d503c3baf220\", \"AscentFuelFlowRate\": 750, \"CruiseFuelFlowRate\": 700}").
			SetMaximumValues("{\"Altitude\": 35000, \"RampWeight\": 13510, \"GroundSpeed\": 310, \"LandingWeight\": 12500, \"TakeoffWeight\": 13420, \"ZeroFuelWeight\": 11500, \"TotalFuelWeight\": 0}").
			SetCurrentLandings(5212).
			SetFuelDetails("{\"Type\": \"Jet A/A1 - See AFM for alternates\", \"CapacityLitres\": null, \"CapacityGallons\": 549}").
			SetOilDetails("{\"Grade\": \"AEROSHELL 500 Turbine Oil\", \"CapacityLitres\": null, \"CapacityGallons\": 2.5}"),
		client.Aircraft.
			Create().
			SetID(getUUID("707c4d8b-2a47-4d06-8cf6-22810c0420c4")).
			SetCompanyID(getUUID("b144a41a-e5de-4032-b776-24ca87536720")).
			SetCurrentFlightHours(4225.5).
			SetCurrentCycles(5441).
			SetAircraftRegistration("C-GLUN").
			SetBaseAirportCode("CYPE").
			SetManufacturer("TEXTRON AVIATION").
			SetManufacturerDesignator("B200GT").
			SetCommonDesignation("B200GT").
			SetCommonName("Super King Air 250 Halo").
			SetDefaultValues("{\"ClimbRate\": 0, \"FlightType\": \"medevac\", \"FlightRules\": \"IFR/703\", \"FlightNumber\": \"CWA928\", \"TrueAirSpeed\": 245, \"ApproachSpeed\": 0, \"TaxiFuelWeight\": 90, \"ConfigurationID\": \"af88f31b-ff34-4d8f-82a5-6e47388f1881\", \"AscentFuelFlowRate\": 750, \"CruiseFuelFlowRate\": 650}").
			SetMaximumValues("{\"Altitude\": 35000, \"RampWeight\": 13510, \"GroundSpeed\": 310, \"LandingWeight\": 12500, \"TakeoffWeight\": 13420, \"ZeroFuelWeight\": 11500, \"TotalFuelWeight\": 0}").
			SetCurrentLandings(5441).
			SetFuelDetails("{\"Type\": \"Jet A/A1 - See AFM for alternates\", \"CapacityLitres\": null, \"CapacityGallons\": 549}").
			SetOilDetails("{\"Grade\": \"AEROSHELL 500 Turbine Oil\", \"CapacityLitres\": null, \"CapacityGallons\": 2.5}"),
		client.Aircraft.
			Create().
			SetID(getUUID("c4b439fd-21a3-4383-87a3-43a07a6c94a8")).
			SetCompanyID(getUUID("b144a41a-e5de-4032-b776-24ca87536721")).
			SetCurrentFlightHours(5031.3).
			SetCurrentCycles(6222).
			SetAircraftRegistration("C-GLUP").
			SetBaseAirportCode("CYPE").
			SetManufacturer("TEXTRON AVIATION").
			SetManufacturerDesignator("B200GT").
			SetCommonDesignation("B200GT").
			SetCommonName("King Air 250 Halo").
			SetDefaultValues("{\"ClimbRate\": 0, \"FlightType\": \"medevac\", \"FlightRules\": \"IFR/703\", \"FlightNumber\": \"CWA929\", \"TrueAirSpeed\": 245, \"ApproachSpeed\": 0, \"TaxiFuelWeight\": 90, \"ConfigurationID\": \"41f1bf50-33b1-4788-91c8-e10fa47c73ba\", \"AscentFuelFlowRate\": 750, \"CruiseFuelFlowRate\": 650}").
			SetMaximumValues("{\"Altitude\": 35000, \"RampWeight\": 13510, \"GroundSpeed\": 310, \"LandingWeight\": 12500, \"TakeoffWeight\": 13420, \"ZeroFuelWeight\": 11500, \"TotalFuelWeight\": 0}").
			SetCurrentLandings(6222).
			SetFuelDetails("{\"Type\": \"Jet A/A1 - See AFM for alternates\", \"CapacityLitres\": null, \"CapacityGallons\": 549}").
			SetOilDetails("{\"Grade\": \"AEROSHELL 500 Turbine Oil\", \"CapacityLitres\": null, \"CapacityGallons\": 2.5}"),
		client.Aircraft.
			Create().
			SetID(getUUID("e3b88d74-b657-43c8-acc9-8336d3349ffb")).
			SetCompanyID(getUUID("b144a41a-e5de-4032-b776-24ca87536722")).
			SetCurrentFlightHours(3589.4).
			SetCurrentCycles(3688).
			SetAircraftRegistration("C-GLUT").
			SetBaseAirportCode("CEZ4").
			SetManufacturer("TEXTRON AVIATION").
			SetManufacturerDesignator("B200GT").
			SetCommonDesignation("B200GT").
			SetCommonName("King Air 250 Halo").
			SetDefaultValues("{\"ClimbRate\": 0, \"FlightType\": \"medevac\", \"FlightRules\": \"IFR/703\", \"FlightNumber\": \"CWA925\", \"TrueAirSpeed\": 245, \"ApproachSpeed\": 0, \"TaxiFuelWeight\": 90, \"ConfigurationID\": \"0ae5d3d0-adec-42d6-9df6-40bd6f52b898\", \"AscentFuelFlowRate\": 750, \"CruiseFuelFlowRate\": 650}").
			SetMaximumValues("{\"Altitude\": 35000, \"RampWeight\": 13510, \"GroundSpeed\": 310, \"LandingWeight\": 12500, \"TakeoffWeight\": 13420, \"ZeroFuelWeight\": 11500, \"TotalFuelWeight\": 0}").
			SetCurrentLandings(3688).
			SetFuelDetails("{\"Type\": \"Jet A/A1 - See AFM for alternates\", \"CapacityLitres\": null, \"CapacityGallons\": 549}").
			SetOilDetails("{\"Grade\": \"AEROSHELL 500 Turbine Oil\", \"CapacityLitres\": null, \"CapacityGallons\": 2.5}"),
		client.Aircraft.
			Create().
			SetID(getUUID("336f432e-4307-44da-a20b-6384637c92fc")).
			SetCompanyID(getUUID("b144a41a-e5de-4032-b776-24ca87536723")).
			SetCurrentFlightHours(5000.1).
			SetCurrentCycles(7154).
			SetAircraftRegistration("C-GLUR").
			SetBaseAirportCode("CYZH").
			SetManufacturer("TEXTRON AVIATION").
			SetManufacturerDesignator("B200GT").
			SetCommonDesignation("B200GT").
			SetCommonName("King Air 250 Halo").
			SetDefaultValues("{\"ClimbRate\": 0, \"FlightType\": \"medevac\", \"FlightRules\": \"IFR/703\", \"FlightNumber\": \"CWA921\", \"TrueAirSpeed\": 245, \"ApproachSpeed\": 0, \"TaxiFuelWeight\": 90, \"ConfigurationID\": \"f8876076-373b-4ade-8156-34dc5ae4d90d\", \"AscentFuelFlowRate\": 750, \"CruiseFuelFlowRate\": 650}").
			SetMaximumValues("{\"Altitude\": 35000, \"RampWeight\": 13510, \"GroundSpeed\": 310, \"LandingWeight\": 12500, \"TakeoffWeight\": 13420, \"ZeroFuelWeight\": 11500, \"TotalFuelWeight\": 0}").
			SetCurrentLandings(7154).
			SetFuelDetails("{\"Type\": \"Jet A/A1 - See AFM for alternates\", \"CapacityLitres\": null, \"CapacityGallons\": 549}").
			SetOilDetails("{\"Grade\": \"AEROSHELL 500 Turbine Oil\", \"CapacityLitres\": null, \"CapacityGallons\": 2.5}"),
		client.Aircraft.
			Create().
			SetID(getUUID("e1204199-cb8e-4956-a110-ba12ef8e8033")).
			SetCompanyID(getUUID("b144a41a-e5de-4032-b776-24ca87536724")).
			SetCurrentFlightHours(15636.8).
			SetCurrentCycles(1082).
			SetAircraftRegistration("C-FOOS").
			SetBaseAirportCode("CEZ4").
			SetManufacturer("CESSNA").
			SetManufacturerDesignator("C206").
			SetCommonDesignation("C206").
			SetCommonName("C206").
			SetDefaultValues("{\"ClimbRate\": 0, \"FlightType\": \"charter\", \"FlightRules\": \"VFR/703\", \"FlightNumber\": \"CWA601\", \"TrueAirSpeed\": 110, \"ApproachSpeed\": 0, \"TaxiFuelWeight\": 10, \"ConfigurationID\": \"28aa8209-64f0-46bf-bc67-ef1670b129bf\", \"AscentFuelFlowRate\": 105, \"CruiseFuelFlowRate\": 100}").
			SetMaximumValues("{\"Altitude\": 14800, \"RampWeight\": 3610, \"GroundSpeed\": 174, \"LandingWeight\": 3600, \"TakeoffWeight\": 3600, \"ZeroFuelWeight\": 3600, \"TotalFuelWeight\": 0}").
			SetCurrentLandings(926).
			SetFuelDetails("{\"Type\": \"100LL\", \"CapacityLitres\": null, \"CapacityGallons\": 65}").
			SetOilDetails("{\"Grade\": \"AEROSHELL W-80\", \"CapacityLitres\": null, \"CapacityGallons\": 3}"),
		client.Aircraft.
			Create().
			SetID(getUUID("4a635a22-6ea8-43a4-b833-87d5366d93f1")).
			SetCompanyID(getUUID("b144a41a-e5de-4032-b776-24ca87536725")).
			SetCurrentFlightHours(7012.6).
			SetAircraftRegistration("C-GETG").
			SetBaseAirportCode("CYEG").
			SetManufacturerDesignator("C337C").
			SetCommonDesignation("337C").
			SetCommonName("Cessna Skymaster").
			SetDefaultValues("{\"ClimbRate\": 0, \"FlightType\": \"charter\", \"FlightRules\": \"IFR/703\", \"FlightNumber\": \"CWA301\", \"TrueAirSpeed\": 0, \"ApproachSpeed\": 0, \"TaxiFuelWeight\": 20, \"ConfigurationID\": \"dadfebd3-d232-4854-bb45-85354f45a36d\", \"AscentFuelFlowRate\": 110, \"CruiseFuelFlowRate\": 105}").
			SetMaximumValues("{\"Altitude\": 19500, \"RampWeight\": 4650, \"GroundSpeed\": 173, \"LandingWeight\": 4400, \"TakeoffWeight\": 4630, \"ZeroFuelWeight\": 0, \"TotalFuelWeight\": 0}").
			SetFuelDetails("{\"Type\": \"100LL\", \"CapacityLitres\": null, \"CapacityGallons\": 92.8}").
			SetOilDetails("{\"Grade\": \"AEROSHELL W-80\", \"CapacityLitres\": null, \"CapacityGallons\": 2.5}"),
		client.Aircraft.
			Create().
			SetID(getUUID("c8e50686-ee80-4563-8f42-03820ea45fa1")).
			SetCompanyID(getUUID("b144a41a-e5de-4032-b776-24ca87536726")).
			SetAircraftRegistration("C-FZGO").
			SetBaseAirportCode("CYEG").
			SetManufacturer("TEXTRON AVIATION").
			SetManufacturerDesignator("C337F").
			SetCommonDesignation("C337").
			SetCommonName("Cessna Skymaster").
			SetDefaultValues("{\"ClimbRate\": 0, \"FlightType\": \"charter\", \"FlightRules\": \"IFR/703\", \"FlightNumber\": \"CWA303\", \"TrueAirSpeed\": 0, \"ApproachSpeed\": 0, \"TaxiFuelWeight\": 20, \"ConfigurationID\": \"b809a827-8cee-4f8c-afc3-cc290d7a86c2\", \"AscentFuelFlowRate\": 110, \"CruiseFuelFlowRate\": 105}").
			SetMaximumValues("{\"Altitude\": 19500, \"RampWeight\": 4650, \"GroundSpeed\": 173, \"LandingWeight\": 4400, \"TakeoffWeight\": 4630, \"ZeroFuelWeight\": 0, \"TotalFuelWeight\": 0}").
			SetFuelDetails("{\"Type\": \"100LL\", \"CapacityLitres\": null, \"CapacityGallons\": 92.8}").
			SetOilDetails("{\"Grade\": \"AEROSHELL W-80\", \"CapacityLitres\": null, \"CapacityGallons\": 2.5}"),
		client.Aircraft.
			Create().
			SetID(getUUID("7649ace6-f137-4b98-b47a-da6cdce786d2")).
			SetCompanyID(getUUID("b144a41a-e5de-4032-b776-24ca87536727")).
			SetCurrentFlightHours(6586.1).
			SetAircraftRegistration("C-FVAS").
			SetBaseAirportCode("CYEG").
			SetManufacturerDesignator("C337G").
			SetCommonDesignation("337G").
			SetCommonName("Cessna Skymaster").
			SetDefaultValues("{\"ClimbRate\": 0, \"FlightType\": \"charter\", \"FlightRules\": \"IFR/703\", \"FlightNumber\": \"CWA305\", \"TrueAirSpeed\": 0, \"ApproachSpeed\": 0, \"TaxiFuelWeight\": 20, \"ConfigurationID\": \"a2c7ef2c-bf48-47d1-b975-1ebac35a965d\", \"AscentFuelFlowRate\": 110, \"CruiseFuelFlowRate\": 105}").
			SetMaximumValues("{\"Altitude\": 19500, \"RampWeight\": 4650, \"GroundSpeed\": 173, \"LandingWeight\": 4400, \"TakeoffWeight\": 4630, \"ZeroFuelWeight\": 0, \"TotalFuelWeight\": 0}").
			SetFuelDetails("{\"Type\": \"100LL\", \"CapacityLitres\": null, \"CapacityGallons\": 150.6}").
			SetOilDetails("{\"Grade\": \"AEROSHELL W-80\", \"CapacityLitres\": null, \"CapacityGallons\": 2}"),
		client.Aircraft.
			Create().
			SetID(getUUID("7331d92d-7320-40ca-b6a4-3337a7b1ddcd")).
			SetCompanyID(getUUID("b144a41a-e5de-4032-b776-24ca87536728")).
			SetCurrentFlightHours(7277.2).
			SetAircraftRegistration("C-GGFA").
			SetBaseAirportCode("CYEG").
			SetManufacturerDesignator("U206E").
			SetCommonDesignation("C337").
			SetCommonName("Cessna Skymaster").
			SetDefaultValues("{\"ClimbRate\": 0, \"FlightType\": \"charter\", \"FlightRules\": \"IFR/703\", \"FlightNumber\": \"CWA302\", \"TrueAirSpeed\": 0, \"ApproachSpeed\": 0, \"TaxiFuelWeight\": 20, \"ConfigurationID\": \"387940b8-0f35-4de0-8a89-5389b7453704\", \"AscentFuelFlowRate\": 110, \"CruiseFuelFlowRate\": 105}").
			SetMaximumValues("{\"Altitude\": 19500, \"RampWeight\": 4650, \"GroundSpeed\": 173, \"LandingWeight\": 4400, \"TakeoffWeight\": 4630, \"ZeroFuelWeight\": 0, \"TotalFuelWeight\": 0}").
			SetFuelDetails("{\"Type\": \"100LL\", \"CapacityLitres\": null, \"CapacityGallons\": 92.8}").
			SetOilDetails("{\"Grade\": \"AEROSHELL W-80\", \"CapacityLitres\": null, \"CapacityGallons\": 2.5}"),
	}
}
