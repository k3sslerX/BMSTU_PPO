package mongodb

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"sort"
	"strconv"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) ListCars(ctx context.Context, query string) ([]models.Car, error) {
	var docs []carDoc
	if err := r.findAll(ctx, "car", bson.M{}, &docs); err != nil {
		return nil, err
	}

	var raceClasses []raceClassDoc
	if err := r.findAll(ctx, "raceclass", bson.M{}, &raceClasses); err != nil {
		return nil, err
	}
	raceClassByID := make(map[string]raceClassDoc, len(raceClasses))
	for _, raceClass := range raceClasses {
		raceClassByID[raceClass.ID] = raceClass
	}

	cars := make([]models.Car, 0, len(docs))
	for _, doc := range docs {
		raceClass := raceClassByID[doc.RaceClassID].Name
		if query != "" &&
			!containsFold(doc.Model, query) &&
			!containsFold(raceClass, query) &&
			!containsFold(strconv.Itoa(doc.Year), query) {
			continue
		}

		id, err := parseID(doc.ID)
		if err != nil {
			return nil, err
		}
		cars = append(cars, models.Car{
			Id:        id,
			Model:     doc.Model,
			Year:      doc.Year,
			RaceClass: raceClass,
		})
	}

	sort.Slice(cars, func(i, j int) bool {
		if cars[i].Model == cars[j].Model {
			return cars[i].Year > cars[j].Year
		}
		return cars[i].Model < cars[j].Model
	})

	return cars, nil
}

func (r *Repository) ListCarParticipants(ctx context.Context, query string) ([]models.CarParticipant, error) {
	var docs []carParticipantDoc
	if err := r.findAll(ctx, "car_p", bson.M{}, &docs); err != nil {
		return nil, err
	}

	var teamParticipants []teamParticipantDoc
	if err := r.findAll(ctx, "team_p", bson.M{}, &teamParticipants); err != nil {
		return nil, err
	}
	driversByParticipant := make(map[string][]string)
	for _, participant := range teamParticipants {
		driversByParticipant[participant.CarP] = append(driversByParticipant[participant.CarP], participant.Driver)
	}

	var drivers []driverDoc
	if err := r.findAll(ctx, "driver", bson.M{}, &drivers); err != nil {
		return nil, err
	}
	driverByID := make(map[string]driverDoc, len(drivers))
	for _, driver := range drivers {
		driverByID[driver.ID] = driver
	}

	participants := make([]models.CarParticipant, 0, len(docs))
	for _, doc := range docs {
		if query != "" && !containsFold(doc.Number, query) {
			continue
		}

		id, err := parseID(doc.ID)
		if err != nil {
			return nil, err
		}
		carID, err := parseID(doc.Car)
		if err != nil {
			return nil, err
		}
		teamID, err := parseID(doc.Team)
		if err != nil {
			return nil, err
		}

		participant := models.CarParticipant{
			Id:      id,
			CarID:   carID,
			TeamID:  teamID,
			Number:  doc.Number,
			Drivers: make([]models.Driver, 0, len(driversByParticipant[doc.ID])),
		}
		for _, driverID := range driversByParticipant[doc.ID] {
			driver, ok := driverByID[driverID]
			if !ok {
				continue
			}
			model, err := driver.model()
			if err != nil {
				return nil, err
			}
			participant.Drivers = append(participant.Drivers, model)
		}
		sort.Slice(participant.Drivers, func(i, j int) bool {
			return participant.Drivers[i].Name < participant.Drivers[j].Name
		})
		participants = append(participants, participant)
	}

	sort.Slice(participants, func(i, j int) bool {
		return participants[i].Number < participants[j].Number
	})

	return participants, nil
}

func (r *Repository) ListChampionships(ctx context.Context, query string) ([]models.Championship, error) {
	var docs []championshipDoc
	if err := r.findAll(ctx, "championship", bson.M{}, &docs); err != nil {
		return nil, err
	}

	var organizers []organizerDoc
	if err := r.findAll(ctx, "organizer", bson.M{}, &organizers); err != nil {
		return nil, err
	}
	organizerByID := make(map[string]organizerDoc, len(organizers))
	for _, organizer := range organizers {
		organizerByID[organizer.ID] = organizer
	}

	championships := make([]models.Championship, 0, len(docs))
	for _, doc := range docs {
		organizer := organizerByID[doc.Organizer].Name
		if query != "" && !containsFold(strconv.Itoa(doc.Year), query) && !containsFold(organizer, query) {
			continue
		}

		id, err := parseID(doc.ID)
		if err != nil {
			return nil, err
		}
		championships = append(championships, models.Championship{
			Id:        id,
			Year:      doc.Year,
			Organizer: organizer,
		})
	}

	sort.Slice(championships, func(i, j int) bool {
		return championships[i].Year > championships[j].Year
	})

	return championships, nil
}

func (r *Repository) ListRaces(ctx context.Context, query string) ([]models.Race, error) {
	var docs []raceDoc
	if err := r.findAll(ctx, "race", bson.M{}, &docs); err != nil {
		return nil, err
	}

	races := make([]models.Race, 0, len(docs))
	for _, doc := range docs {
		if query != "" && !containsFold(doc.Name, query) {
			continue
		}

		id, err := parseID(doc.ID)
		if err != nil {
			return nil, err
		}
		championshipID, err := parseID(doc.Championship)
		if err != nil {
			return nil, err
		}
		trackID, err := parseID(doc.Track)
		if err != nil {
			return nil, err
		}
		date, err := parseDate(doc.Date)
		if err != nil {
			return nil, err
		}

		races = append(races, models.Race{
			Id:             id,
			Name:           doc.Name,
			Date:           date,
			Type:           doc.Type,
			Duration:       doc.Duration,
			Track:          models.Track{Id: trackID},
			ChampionshipId: championshipID,
		})
	}

	sort.Slice(races, func(i, j int) bool {
		if races[i].Date.Equal(races[j].Date) {
			return races[i].Name < races[j].Name
		}
		return races[i].Date.After(races[j].Date)
	})

	return races, nil
}

func (r *Repository) ListTracks(ctx context.Context, query string) ([]models.Track, error) {
	var docs []trackDoc
	if err := r.findAll(ctx, "track", bson.M{}, &docs); err != nil {
		return nil, err
	}

	tracks := make([]models.Track, 0, len(docs))
	for _, doc := range docs {
		if query != "" && !containsFold(doc.Name, query) {
			continue
		}

		id, err := parseID(doc.ID)
		if err != nil {
			return nil, err
		}
		tracks = append(tracks, models.Track{
			Id:      id,
			Name:    doc.Name,
			Country: doc.Country,
			Length:  doc.Length,
			Turns:   doc.Turns,
		})
	}

	sort.Slice(tracks, func(i, j int) bool {
		return tracks[i].Name < tracks[j].Name
	})

	return tracks, nil
}

func (r *Repository) ListUsers(ctx context.Context, query string) ([]models.User, error) {
	var docs []userDoc
	if err := r.findAll(ctx, "users", bson.M{}, &docs); err != nil {
		return nil, err
	}

	users := make([]models.User, 0, len(docs))
	for _, doc := range docs {
		if query != "" && !containsFold(doc.Name, query) && !containsFold(doc.Email, query) {
			continue
		}

		id, err := parseID(doc.ID)
		if err != nil {
			return nil, err
		}
		users = append(users, models.User{
			Id:    id,
			Name:  doc.Name,
			Email: doc.Email,
			Role:  models.Role(doc.Role),
		})
	}

	sort.Slice(users, func(i, j int) bool {
		return users[i].Email < users[j].Email
	})

	return users, nil
}

func (r *Repository) CreateDriver(ctx context.Context, driver models.Driver) (models.Driver, error) {
	id := newID()
	_, err := r.collection("driver").InsertOne(ctx, bson.M{
		"_id":         id,
		"id":          id,
		"name":        driver.Name,
		"nationality": driver.Nationality,
		"birthday":    driver.Birthday,
	})
	if err != nil {
		return driver, err
	}

	parsedID, err := parseID(id)
	if err != nil {
		return driver, err
	}
	driver.Id = parsedID
	return driver, nil
}

func (r *Repository) CreateTeam(ctx context.Context, team models.Team) (models.Team, error) {
	id := newID()
	_, err := r.collection("team").InsertOne(ctx, bson.M{
		"_id":     id,
		"id":      id,
		"name":    team.Name,
		"country": team.Country,
	})
	if err != nil {
		return team, err
	}

	parsedID, err := parseID(id)
	if err != nil {
		return team, err
	}
	team.Id = parsedID
	return team, nil
}

func (r *Repository) CreateTrack(ctx context.Context, track models.Track) (models.Track, error) {
	id := newID()
	_, err := r.collection("track").InsertOne(ctx, bson.M{
		"_id":        id,
		"id":         id,
		"name":       track.Name,
		"country":    track.Country,
		"lap_length": track.Length,
		"turns":      track.Turns,
	})
	if err != nil {
		return track, err
	}

	parsedID, err := parseID(id)
	if err != nil {
		return track, err
	}
	track.Id = parsedID
	return track, nil
}

func (r *Repository) CreateRace(ctx context.Context, race models.Race) (models.Race, error) {
	id := newID()
	_, err := r.collection("race").InsertOne(ctx, bson.M{
		"_id":          id,
		"id":           id,
		"championship": idString(race.ChampionshipId),
		"track":        idString(race.Track.Id),
		"name":         race.Name,
		"date":         dateString(race.Date),
		"type":         race.Type,
		"duration":     race.Duration,
	})
	if err != nil {
		return race, err
	}

	parsedID, err := parseID(id)
	if err != nil {
		return race, err
	}
	race.Id = parsedID
	return race, nil
}

func (r *Repository) CreateCarParticipant(ctx context.Context, carParticipant models.CarParticipant) (models.CarParticipant, error) {
	id := newID()
	_, err := r.collection("car_p").InsertOne(ctx, bson.M{
		"_id":    id,
		"id":     id,
		"car":    idString(carParticipant.CarID),
		"team":   idString(carParticipant.TeamID),
		"number": carParticipant.Number,
	})
	if err != nil {
		return carParticipant, err
	}

	parsedID, err := parseID(id)
	if err != nil {
		return carParticipant, err
	}
	carParticipant.Id = parsedID

	if err := r.replaceCarParticipantDrivers(ctx, carParticipant.Id, carParticipant.Drivers); err != nil {
		return carParticipant, err
	}

	return carParticipant, nil
}

func (r *Repository) UpsertRaceResult(ctx context.Context, result models.RaceResult) (models.RaceResult, error) {
	if result.RaceID == uuid.Nil || result.CarParticipantID == uuid.Nil ||
		result.FinishPos <= 0 || result.QualifyingPos <= 0 {
		return result, shared.ErrorInvalidData
	}

	raceExists, err := r.exists(ctx, "race", result.RaceID)
	if err != nil {
		return result, err
	}
	carParticipantExists, err := r.exists(ctx, "car_p", result.CarParticipantID)
	if err != nil {
		return result, err
	}
	if !raceExists || !carParticipantExists {
		return result, shared.ErrorNotFound
	}

	raceID := idString(result.RaceID)
	carParticipantID := idString(result.CarParticipantID)
	if err := r.upsertStanding(ctx, "finish", raceID, carParticipantID, result.FinishPos); err != nil {
		return result, err
	}
	if err := r.upsertStanding(ctx, "qualifying", raceID, carParticipantID, result.QualifyingPos); err != nil {
		return result, err
	}

	return result, nil
}

func (r *Repository) UpdateDriver(ctx context.Context, driver models.Driver) (models.Driver, error) {
	if driver.Id == uuid.Nil || (driver.Name == "" && driver.Nationality == "" && driver.Birthday == "") {
		return driver, shared.ErrorInvalidData
	}

	update := bson.M{}
	setString(update, "name", driver.Name)
	setString(update, "nationality", driver.Nationality)
	setString(update, "birthday", driver.Birthday)

	result, err := r.collection("driver").UpdateOne(ctx, bson.M{"id": idString(driver.Id)}, bson.M{"$set": update})
	if err != nil {
		return driver, err
	}
	return driver, updateResultError(result)
}

func (r *Repository) UpdateTeam(ctx context.Context, team models.Team) (models.Team, error) {
	if team.Id == uuid.Nil || (team.Name == "" && team.Country == "") {
		return team, shared.ErrorInvalidData
	}

	update := bson.M{}
	setString(update, "name", team.Name)
	setString(update, "country", team.Country)

	result, err := r.collection("team").UpdateOne(ctx, bson.M{"id": idString(team.Id)}, bson.M{"$set": update})
	if err != nil {
		return team, err
	}
	return team, updateResultError(result)
}

func (r *Repository) UpdateTrack(ctx context.Context, track models.Track) (models.Track, error) {
	if track.Id == uuid.Nil || (track.Name == "" && track.Country == "" && track.Length == 0 && track.Turns == 0) {
		return track, shared.ErrorInvalidData
	}

	update := bson.M{}
	setString(update, "name", track.Name)
	setString(update, "country", track.Country)
	setInt(update, "lap_length", track.Length)
	setInt(update, "turns", track.Turns)

	result, err := r.collection("track").UpdateOne(ctx, bson.M{"id": idString(track.Id)}, bson.M{"$set": update})
	if err != nil {
		return track, err
	}
	return track, updateResultError(result)
}

func (r *Repository) UpdateRace(ctx context.Context, race models.Race) (models.Race, error) {
	if race.Id == uuid.Nil || (race.Name == "" && race.Date.IsZero() && race.Type == 0 &&
		race.Duration == 0 && race.Track.Id == uuid.Nil && race.ChampionshipId == uuid.Nil) {
		return race, shared.ErrorInvalidData
	}

	update := bson.M{}
	setString(update, "name", race.Name)
	if !race.Date.IsZero() {
		update["date"] = dateString(race.Date)
	}
	setInt(update, "type", race.Type)
	setInt(update, "duration", race.Duration)
	setID(update, "track", race.Track.Id)
	setID(update, "championship", race.ChampionshipId)

	result, err := r.collection("race").UpdateOne(ctx, bson.M{"id": idString(race.Id)}, bson.M{"$set": update})
	if err != nil {
		return race, err
	}
	return race, updateResultError(result)
}

func (r *Repository) UpdateCarParticipant(ctx context.Context, carParticipant models.CarParticipant) (models.CarParticipant, error) {
	if carParticipant.Id == uuid.Nil || (carParticipant.CarID == uuid.Nil && carParticipant.TeamID == uuid.Nil && carParticipant.Number == "") {
		return carParticipant, shared.ErrorInvalidData
	}

	update := bson.M{}
	setID(update, "car", carParticipant.CarID)
	setID(update, "team", carParticipant.TeamID)
	setString(update, "number", carParticipant.Number)

	result, err := r.collection("car_p").UpdateOne(ctx, bson.M{"id": idString(carParticipant.Id)}, bson.M{"$set": update})
	if err != nil {
		return carParticipant, err
	}
	return carParticipant, updateResultError(result)
}

func (r *Repository) UpdateCarParticipantDrivers(ctx context.Context, carParticipant models.CarParticipant) (models.CarParticipant, error) {
	if carParticipant.Id == uuid.Nil {
		return carParticipant, shared.ErrorInvalidData
	}

	exists, err := r.exists(ctx, "car_p", carParticipant.Id)
	if err != nil {
		return carParticipant, err
	}
	if !exists {
		return carParticipant, shared.ErrorNotFound
	}

	if err := r.replaceCarParticipantDrivers(ctx, carParticipant.Id, carParticipant.Drivers); err != nil {
		return carParticipant, err
	}

	return carParticipant, nil
}

func (r *Repository) UpdateUserRole(ctx context.Context, user models.User) (models.User, error) {
	if user.Id == uuid.Nil || (user.Role != models.RoleAdmin && user.Role != models.RoleUser) {
		return models.User{}, shared.ErrorInvalidData
	}

	result, err := r.collection("users").UpdateOne(
		ctx,
		bson.M{"id": idString(user.Id)},
		bson.M{"$set": bson.M{"role": string(user.Role)}},
	)
	if err != nil {
		return models.User{}, err
	}
	if err := updateResultError(result); err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *Repository) DeleteDriver(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return shared.ErrorInvalidData
	}
	exists, err := r.exists(ctx, "driver", id)
	if err != nil {
		return err
	}
	if !exists {
		return shared.ErrorNotFound
	}

	driverID := idString(id)
	if err := deleteMany(ctx, r.collection("favourite_drivers"), bson.M{"driver": driverID}); err != nil {
		return err
	}
	if err := deleteMany(ctx, r.collection("team_p"), bson.M{"driver": driverID}); err != nil {
		return err
	}
	return r.deleteOneByID(ctx, "driver", id)
}

func (r *Repository) DeleteTeam(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return shared.ErrorInvalidData
	}
	exists, err := r.exists(ctx, "team", id)
	if err != nil {
		return err
	}
	if !exists {
		return shared.ErrorNotFound
	}

	teamID := idString(id)
	if err := deleteMany(ctx, r.collection("favourite_teams"), bson.M{"team": teamID}); err != nil {
		return err
	}
	var participants []carParticipantDoc
	if err := r.findAll(ctx, "car_p", bson.M{"team": teamID}, &participants); err != nil {
		return err
	}
	for _, participant := range participants {
		parsedID, err := parseID(participant.ID)
		if err != nil {
			return err
		}
		if err := r.deleteCarParticipantCascade(ctx, parsedID); err != nil && !errors.Is(err, shared.ErrorNotFound) {
			return err
		}
	}
	return r.deleteOneByID(ctx, "team", id)
}

func (r *Repository) DeleteTrack(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return shared.ErrorInvalidData
	}
	exists, err := r.exists(ctx, "track", id)
	if err != nil {
		return err
	}
	if !exists {
		return shared.ErrorNotFound
	}

	var races []raceDoc
	if err := r.findAll(ctx, "race", bson.M{"track": idString(id)}, &races); err != nil {
		return err
	}
	for _, race := range races {
		parsedID, err := parseID(race.ID)
		if err != nil {
			return err
		}
		if err := r.deleteRaceCascade(ctx, parsedID); err != nil && !errors.Is(err, shared.ErrorNotFound) {
			return err
		}
	}
	return r.deleteOneByID(ctx, "track", id)
}

func (r *Repository) DeleteRace(ctx context.Context, id uuid.UUID) error {
	return r.deleteRaceCascade(ctx, id)
}

func (r *Repository) DeleteCarParticipant(ctx context.Context, id uuid.UUID) error {
	return r.deleteCarParticipantCascade(ctx, id)
}

func (r *Repository) upsertStanding(ctx context.Context, collection string, raceID string, carParticipantID string, position int) error {
	_, err := r.collection(collection).ReplaceOne(
		ctx,
		bson.M{"_id": rowKey(raceID, carParticipantID)},
		bson.M{
			"_id":   rowKey(raceID, carParticipantID),
			"pos":   position,
			"race":  raceID,
			"car_p": carParticipantID,
		},
		options.Replace().SetUpsert(true),
	)
	return err
}

func (r *Repository) replaceCarParticipantDrivers(ctx context.Context, carParticipantID uuid.UUID, drivers []models.Driver) error {
	carParticipantIDString := idString(carParticipantID)
	if err := deleteMany(ctx, r.collection("team_p"), bson.M{"car_p": carParticipantIDString}); err != nil {
		return err
	}

	if len(drivers) == 0 {
		return nil
	}

	documents := make([]any, 0, len(drivers))
	for _, driver := range drivers {
		if driver.Id == uuid.Nil {
			return shared.ErrorInvalidData
		}
		driverID := idString(driver.Id)
		documents = append(documents, bson.M{
			"_id":    rowKey(carParticipantIDString, driverID),
			"car_p":  carParticipantIDString,
			"driver": driverID,
		})
	}

	_, err := r.collection("team_p").InsertMany(ctx, documents)
	return err
}

func (r *Repository) deleteRaceCascade(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return shared.ErrorInvalidData
	}
	exists, err := r.exists(ctx, "race", id)
	if err != nil {
		return err
	}
	if !exists {
		return shared.ErrorNotFound
	}

	raceID := idString(id)
	if err := deleteMany(ctx, r.collection("finish"), bson.M{"race": raceID}); err != nil {
		return err
	}
	if err := deleteMany(ctx, r.collection("qualifying"), bson.M{"race": raceID}); err != nil {
		return err
	}
	return r.deleteOneByID(ctx, "race", id)
}

func (r *Repository) deleteCarParticipantCascade(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return shared.ErrorInvalidData
	}
	exists, err := r.exists(ctx, "car_p", id)
	if err != nil {
		return err
	}
	if !exists {
		return shared.ErrorNotFound
	}

	carParticipantID := idString(id)
	if err := deleteMany(ctx, r.collection("team_p"), bson.M{"car_p": carParticipantID}); err != nil {
		return err
	}
	if err := deleteMany(ctx, r.collection("finish"), bson.M{"car_p": carParticipantID}); err != nil {
		return err
	}
	if err := deleteMany(ctx, r.collection("qualifying"), bson.M{"car_p": carParticipantID}); err != nil {
		return err
	}
	return r.deleteOneByID(ctx, "car_p", id)
}
