package mongodb

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"errors"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const rowKeySeparator = "\x1f"

type Repository struct {
	DB *mongo.Database
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) EnsureIndexes(ctx context.Context) error {
	for _, collection := range []string{
		"manufacturer", "raceclass", "team", "organizer", "track", "driver",
		"car", "championship", "car_p", "race", "users",
	} {
		_, err := r.collection(collection).Indexes().CreateOne(ctx, mongo.IndexModel{
			Keys:    bson.D{{Key: "id", Value: 1}},
			Options: options.Index().SetUnique(true),
		})
		if err != nil {
			return err
		}
	}

	_, err := r.collection("users").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{Key: "email", Value: 1}},
	})
	return err
}

func (r *Repository) collection(name string) *mongo.Collection {
	return r.DB.Collection(name)
}

type carDoc struct {
	ID           string `bson:"id"`
	Model        string `bson:"model"`
	Year         int    `bson:"year_of_production"`
	Manufacturer string `bson:"manufacturer,omitempty"`
	RaceClassID  string `bson:"raceclass,omitempty"`
}

type carParticipantDoc struct {
	ID     string `bson:"id"`
	Car    string `bson:"car,omitempty"`
	Team   string `bson:"team,omitempty"`
	Number string `bson:"number"`
}

type championshipDoc struct {
	ID        string `bson:"id"`
	Year      int    `bson:"year"`
	Organizer string `bson:"organizer,omitempty"`
}

type driverDoc struct {
	ID          string `bson:"id"`
	Name        string `bson:"name"`
	Nationality string `bson:"nationality"`
	Birthday    string `bson:"birthday"`
}

type organizerDoc struct {
	ID   string `bson:"id"`
	Name string `bson:"name"`
}

type raceClassDoc struct {
	ID   string `bson:"id"`
	Name string `bson:"name"`
}

type raceDoc struct {
	ID           string `bson:"id"`
	Name         string `bson:"name"`
	Date         string `bson:"date"`
	Type         int    `bson:"type"`
	Duration     int    `bson:"duration"`
	Track        string `bson:"track,omitempty"`
	Championship string `bson:"championship,omitempty"`
}

type standingDoc struct {
	Pos  int    `bson:"pos"`
	Race string `bson:"race,omitempty"`
	CarP string `bson:"car_p,omitempty"`
}

type teamDoc struct {
	ID      string `bson:"id"`
	Name    string `bson:"name"`
	Country string `bson:"country"`
}

type teamParticipantDoc struct {
	CarP   string `bson:"car_p,omitempty"`
	Driver string `bson:"driver,omitempty"`
}

type trackDoc struct {
	ID      string `bson:"id"`
	Name    string `bson:"name"`
	Country string `bson:"country"`
	Length  int    `bson:"lap_length"`
	Turns   int    `bson:"turns"`
}

type userDoc struct {
	ID           string    `bson:"id"`
	Name         string    `bson:"name"`
	Email        string    `bson:"email"`
	PasswordHash string    `bson:"passwordhash"`
	Role         string    `bson:"role"`
	CreatedAt    time.Time `bson:"created_at,omitempty"`
}

type favouriteDriverDoc struct {
	UserID string `bson:"user_id"`
	Driver string `bson:"driver"`
}

type favouriteTeamDoc struct {
	UserID string `bson:"user_id"`
	Team   string `bson:"team"`
}

type sudokuCompletionDoc struct {
	UserID      string `bson:"user_id"`
	MatrixType  string `bson:"matrix_type"`
	CompletedOn string `bson:"completed_on"`
}

func rowKey(parts ...string) string {
	return strings.Join(parts, rowKeySeparator)
}

func idString(id uuid.UUID) string {
	return id.String()
}

func newID() string {
	return uuid.NewString()
}

func parseID(value string) (uuid.UUID, error) {
	if strings.TrimSpace(value) == "" {
		return uuid.Nil, nil
	}
	return uuid.Parse(value)
}

func parseDate(value string) (time.Time, error) {
	if strings.TrimSpace(value) == "" {
		return time.Time{}, nil
	}
	return time.Parse(time.DateOnly, value)
}

func dateString(value time.Time) string {
	if value.IsZero() {
		return ""
	}
	return value.Format(time.DateOnly)
}

func containsFold(value, query string) bool {
	query = strings.TrimSpace(query)
	if query == "" {
		return true
	}
	return strings.Contains(strings.ToLower(value), strings.ToLower(query))
}

func isNoDocuments(err error) bool {
	return errors.Is(err, mongo.ErrNoDocuments)
}

func (r *Repository) findAll(ctx context.Context, collection string, filter any, out any) error {
	cursor, err := r.collection(collection).Find(ctx, filter)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)
	return cursor.All(ctx, out)
}

func (r *Repository) exists(ctx context.Context, collection string, id uuid.UUID) (bool, error) {
	if id == uuid.Nil {
		return false, nil
	}
	err := r.collection(collection).FindOne(ctx, bson.M{"id": idString(id)}).Err()
	if err == nil {
		return true, nil
	}
	if isNoDocuments(err) {
		return false, nil
	}
	return false, err
}

func (r *Repository) deleteOneByID(ctx context.Context, collection string, id uuid.UUID) error {
	if id == uuid.Nil {
		return shared.ErrorInvalidData
	}

	result, err := r.collection(collection).DeleteOne(ctx, bson.M{"id": idString(id)})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return shared.ErrorNotFound
	}

	return nil
}

func setString(update bson.M, name string, value string) {
	if value != "" {
		update[name] = value
	}
}

func setInt(update bson.M, name string, value int) {
	if value != 0 {
		update[name] = value
	}
}

func setID(update bson.M, name string, value uuid.UUID) {
	if value != uuid.Nil {
		update[name] = idString(value)
	}
}

func hasUpdate(update bson.M) bool {
	return len(update) > 0
}

func updateResultError(result *mongo.UpdateResult) error {
	if result.MatchedCount == 0 {
		return shared.ErrorNotFound
	}
	return nil
}

func deleteMany(ctx context.Context, collection *mongo.Collection, filter any) error {
	_, err := collection.DeleteMany(ctx, filter)
	return err
}

func (d driverDoc) model() (models.Driver, error) {
	id, err := parseID(d.ID)
	if err != nil {
		return models.Driver{}, err
	}
	return models.Driver{
		Id:          id,
		Name:        d.Name,
		Birthday:    d.Birthday,
		Nationality: d.Nationality,
	}, nil
}

func (d teamDoc) model() (models.Team, error) {
	id, err := parseID(d.ID)
	if err != nil {
		return models.Team{}, err
	}
	return models.Team{Id: id, Name: d.Name, Country: d.Country}, nil
}

type statsData struct {
	cars            map[string]carDoc
	carParticipants map[string]carParticipantDoc
	drivers         map[string]driverDoc
	raceClasses     map[string]raceClassDoc
	races           map[string]raceDoc
	teams           map[string]teamDoc
	teamP           []teamParticipantDoc
	finishes        []standingDoc
	qualifying      []standingDoc
}

func (r *Repository) loadStatsData(ctx context.Context) (statsData, error) {
	data := statsData{}

	var cars []carDoc
	if err := r.findAll(ctx, "car", bson.M{}, &cars); err != nil {
		return data, err
	}
	data.cars = make(map[string]carDoc, len(cars))
	for _, car := range cars {
		data.cars[car.ID] = car
	}

	var participants []carParticipantDoc
	if err := r.findAll(ctx, "car_p", bson.M{}, &participants); err != nil {
		return data, err
	}
	data.carParticipants = make(map[string]carParticipantDoc, len(participants))
	for _, participant := range participants {
		data.carParticipants[participant.ID] = participant
	}

	var drivers []driverDoc
	if err := r.findAll(ctx, "driver", bson.M{}, &drivers); err != nil {
		return data, err
	}
	data.drivers = make(map[string]driverDoc, len(drivers))
	for _, driver := range drivers {
		data.drivers[driver.ID] = driver
	}

	var raceClasses []raceClassDoc
	if err := r.findAll(ctx, "raceclass", bson.M{}, &raceClasses); err != nil {
		return data, err
	}
	data.raceClasses = make(map[string]raceClassDoc, len(raceClasses))
	for _, raceClass := range raceClasses {
		data.raceClasses[raceClass.ID] = raceClass
	}

	var races []raceDoc
	if err := r.findAll(ctx, "race", bson.M{}, &races); err != nil {
		return data, err
	}
	data.races = make(map[string]raceDoc, len(races))
	for _, race := range races {
		data.races[race.ID] = race
	}

	var teams []teamDoc
	if err := r.findAll(ctx, "team", bson.M{}, &teams); err != nil {
		return data, err
	}
	data.teams = make(map[string]teamDoc, len(teams))
	for _, team := range teams {
		data.teams[team.ID] = team
	}

	if err := r.findAll(ctx, "team_p", bson.M{}, &data.teamP); err != nil {
		return data, err
	}
	if err := r.findAll(ctx, "finish", bson.M{}, &data.finishes); err != nil {
		return data, err
	}
	if err := r.findAll(ctx, "qualifying", bson.M{}, &data.qualifying); err != nil {
		return data, err
	}

	return data, nil
}

func (d statsData) raceClassNameForCarParticipant(carParticipantID string) string {
	participant, ok := d.carParticipants[carParticipantID]
	if !ok {
		return ""
	}
	car, ok := d.cars[participant.Car]
	if !ok {
		return ""
	}
	return d.raceClasses[car.RaceClassID].Name
}

func (d statsData) driverLineups(driverID string) map[string]string {
	lineups := make(map[string]string)
	for _, participant := range d.teamP {
		if participant.Driver != driverID {
			continue
		}
		lineups[participant.CarP] = d.raceClassNameForCarParticipant(participant.CarP)
	}
	return lineups
}

func (d statsData) teamLineups(teamID string) map[string]string {
	lineups := make(map[string]string)
	for _, participant := range d.carParticipants {
		if participant.Team != teamID {
			continue
		}
		lineups[participant.ID] = d.raceClassNameForCarParticipant(participant.ID)
	}
	return lineups
}

type personalPoints struct {
	carParticipantID string
	teamID           string
	raceClass        string
	points           int
}

func (d statsData) calculatePersonalPoints(championshipID string) []personalPoints {
	pointsByParticipant := make(map[string]int)
	for _, finish := range d.finishes {
		race, ok := d.races[finish.Race]
		if !ok || race.Championship != championshipID {
			continue
		}
		pointsByParticipant[finish.CarP] += pointsForFinish(finish.Pos, race.Type, race.Duration)
	}

	points := make([]personalPoints, 0, len(pointsByParticipant))
	for participantID, total := range pointsByParticipant {
		participant, ok := d.carParticipants[participantID]
		if !ok {
			continue
		}
		points = append(points, personalPoints{
			carParticipantID: participantID,
			teamID:           participant.Team,
			raceClass:        d.raceClassNameForCarParticipant(participantID),
			points:           total,
		})
	}

	return points
}

func pointsForFinish(position int, raceType int, duration int) int {
	base := 0
	switch position {
	case 1:
		base = 25
	case 2:
		base = 18
	case 3:
		base = 15
	case 4:
		base = 12
	case 5:
		base = 10
	case 6:
		base = 8
	case 7:
		base = 6
	case 8:
		base = 4
	case 9:
		base = 2
	case 10:
		base = 1
	}

	multiplier := 1.0
	if raceType == 2 && duration > 6 && duration < 24 {
		multiplier = 1.5
	}
	if raceType == 2 && duration >= 24 {
		multiplier = 2
	}

	return int(math.Ceil(float64(base) * multiplier))
}

func minPositive(current int, candidate int) int {
	if candidate <= 0 {
		return current
	}
	if current == 0 || candidate < current {
		return candidate
	}
	return current
}

func statValue(stats models.Stats, field string) (int, error) {
	switch field {
	case "total_races":
		return stats.TotalRaces, nil
	case "total_wins":
		return stats.TotalWins, nil
	case "total_podiums":
		return stats.TotalPodiums, nil
	case "total_points":
		return stats.TotalPoints, nil
	case "total_poles":
		return stats.TotalPoles, nil
	case "best_finish":
		return stats.BestFinish, nil
	case "best_qualifying":
		return stats.BestQualifying, nil
	case "championships_wins":
		return stats.ChampionshipsWins, nil
	case "best_championship_position":
		return stats.BestChampionshipPosition, nil
	default:
		return 0, fmt.Errorf("unsupported sudoku stats field: %s", field)
	}
}

func matchesCondition(stats models.Stats, condition models.SudokuCondition) (bool, error) {
	value, err := statValue(stats, condition.Field)
	if err != nil {
		return false, err
	}
	if isNullableStatsField(condition.Field) && value == 0 {
		return false, nil
	}

	switch condition.Op {
	case models.ConditionOperatorGT:
		return value > condition.Value, nil
	case models.ConditionOperatorLT:
		return value < condition.Value, nil
	case models.ConditionOperatorGTE:
		return value >= condition.Value, nil
	case models.ConditionOperatorLTE:
		return value <= condition.Value, nil
	case models.ConditionOperatorEQ:
		return value == condition.Value, nil
	default:
		return false, fmt.Errorf("unsupported sudoku condition operator: %s", condition.Op)
	}
}

func isNullableStatsField(field string) bool {
	switch field {
	case "best_finish", "best_qualifying", "best_championship_position":
		return true
	default:
		return false
	}
}

func matchesConditions(stats models.Stats, firstCondition models.SudokuCondition, secondCondition models.SudokuCondition) (bool, error) {
	firstOK, err := matchesCondition(stats, firstCondition)
	if err != nil || !firstOK {
		return firstOK, err
	}
	return matchesCondition(stats, secondCondition)
}
