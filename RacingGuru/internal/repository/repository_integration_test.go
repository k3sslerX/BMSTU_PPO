package repository

import (
	"RacingGuru/internal/models"
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type integrationFixture struct {
	repo *Repository
	pool *pgxpool.Pool

	organizerID       uuid.UUID
	championshipID    uuid.UUID
	manufacturerID    uuid.UUID
	raceclassID       uuid.UUID
	carID             uuid.UUID
	trackID           uuid.UUID
	teamID            uuid.UUID
	driverID          uuid.UUID
	secondDriverID    uuid.UUID
	raceID            uuid.UUID
	userID            uuid.UUID
	carParticipantIDs []uuid.UUID
}

func TestRepositoryAdminLifecycleIntegration(t *testing.T) {
	ctx := context.Background()
	fixture := newIntegrationFixture(t)

	driver, err := fixture.repo.CreateDriver(ctx, models.Driver{
		Name:        integrationName("driver"),
		Birthday:    "1990-01-02",
		Nationality: "Testland",
	})
	if err != nil {
		t.Fatalf("CreateDriver() error = %v", err)
	}
	fixture.driverID = driver.Id

	updatedDriver, err := fixture.repo.UpdateDriver(ctx, models.Driver{
		Id:          driver.Id,
		Nationality: "Updatedland",
	})
	if err != nil {
		t.Fatalf("UpdateDriver() error = %v", err)
	}
	if updatedDriver.Id != driver.Id {
		t.Fatalf("UpdateDriver() id = %v, want %v", updatedDriver.Id, driver.Id)
	}

	team, err := fixture.repo.CreateTeam(ctx, models.Team{
		Name:    integrationName("team"),
		Country: "Testland",
	})
	if err != nil {
		t.Fatalf("CreateTeam() error = %v", err)
	}
	fixture.teamID = team.Id

	updatedTeam, err := fixture.repo.UpdateTeam(ctx, models.Team{
		Id:      team.Id,
		Country: "Updatedland",
	})
	if err != nil {
		t.Fatalf("UpdateTeam() error = %v", err)
	}
	if updatedTeam.Id != team.Id {
		t.Fatalf("UpdateTeam() id = %v, want %v", updatedTeam.Id, team.Id)
	}

	track, err := fixture.repo.CreateTrack(ctx, models.Track{
		Name:    integrationName("track"),
		Country: "Testland",
		Length:  5000,
		Turns:   12,
	})
	if err != nil {
		t.Fatalf("CreateTrack() error = %v", err)
	}
	fixture.trackID = track.Id

	updatedTrack, err := fixture.repo.UpdateTrack(ctx, models.Track{
		Id:     track.Id,
		Turns:  14,
		Length: 5100,
	})
	if err != nil {
		t.Fatalf("UpdateTrack() error = %v", err)
	}
	if updatedTrack.Id != track.Id {
		t.Fatalf("UpdateTrack() id = %v, want %v", updatedTrack.Id, track.Id)
	}

	fixture.insertOrganizer(t, integrationName("organizer"))
	fixture.insertChampionship(t, 2026)
	fixture.insertManufacturer(t, integrationName("manufacturer"))
	fixture.insertRaceclass(t, integrationName("raceclass"))
	fixture.insertCar(t, integrationName("car-model"), 2024)

	secondDriverID := uuid.New()
	fixture.secondDriverID = secondDriverID
	fixture.insertDriver(t, secondDriverID, integrationName("second-driver"), "1995-03-04", "Secondland")

	carParticipant, err := fixture.repo.CreateCarParticipant(ctx, models.CarParticipant{
		CarID:   fixture.carID,
		TeamID:  team.Id,
		Number:  "77",
		Drivers: []models.Driver{{Id: driver.Id}},
	})
	if err != nil {
		t.Fatalf("CreateCarParticipant() error = %v", err)
	}
	fixture.carParticipantIDs = append(fixture.carParticipantIDs, carParticipant.Id)
	assertTeamPDriverCount(t, fixture.pool, carParticipant.Id, 1)

	updatedCarParticipant, err := fixture.repo.UpdateCarParticipant(ctx, models.CarParticipant{
		Id:     carParticipant.Id,
		Number: "99",
	})
	if err != nil {
		t.Fatalf("UpdateCarParticipant() error = %v", err)
	}
	if updatedCarParticipant.Id != carParticipant.Id {
		t.Fatalf("UpdateCarParticipant() id = %v, want %v", updatedCarParticipant.Id, carParticipant.Id)
	}

	_, err = fixture.repo.UpdateCarParticipantDrivers(ctx, models.CarParticipant{
		Id: carParticipant.Id,
		Drivers: []models.Driver{
			{Id: driver.Id},
			{Id: secondDriverID},
		},
	})
	if err != nil {
		t.Fatalf("UpdateCarParticipantDrivers() error = %v", err)
	}
	assertTeamPDriverCount(t, fixture.pool, carParticipant.Id, 2)

	race, err := fixture.repo.CreateRace(ctx, models.Race{
		Name:           integrationName("race"),
		Date:           time.Date(2026, time.April, 8, 12, 0, 0, 0, time.UTC),
		Type:           1,
		Duration:       90,
		Track:          models.Track{Id: track.Id},
		ChampionshipId: fixture.championshipID,
	})
	if err != nil {
		t.Fatalf("CreateRace() error = %v", err)
	}
	fixture.raceID = race.Id

	updatedRace, err := fixture.repo.UpdateRace(ctx, models.Race{
		Id:       race.Id,
		Name:     integrationName("race-updated"),
		Duration: 95,
	})
	if err != nil {
		t.Fatalf("UpdateRace() error = %v", err)
	}
	if updatedRace.Id != race.Id {
		t.Fatalf("UpdateRace() id = %v, want %v", updatedRace.Id, race.Id)
	}

	_, err = fixture.repo.UpsertRaceResult(ctx, models.RaceResult{
		RaceID:           race.Id,
		CarParticipantID: carParticipant.Id,
		FinishPos:        2,
		QualifyingPos:    1,
	})
	if err != nil {
		t.Fatalf("UpsertRaceResult(insert) error = %v", err)
	}
	assertStandingPosition(t, fixture.pool, "finish", race.Id, carParticipant.Id, 2)
	assertStandingPosition(t, fixture.pool, "qualifying", race.Id, carParticipant.Id, 1)

	_, err = fixture.repo.UpsertRaceResult(ctx, models.RaceResult{
		RaceID:           race.Id,
		CarParticipantID: carParticipant.Id,
		FinishPos:        1,
		QualifyingPos:    2,
	})
	if err != nil {
		t.Fatalf("UpsertRaceResult(update) error = %v", err)
	}
	assertStandingPosition(t, fixture.pool, "finish", race.Id, carParticipant.Id, 1)
	assertStandingPosition(t, fixture.pool, "qualifying", race.Id, carParticipant.Id, 2)
}

func TestRepositoryUsersLifecycleIntegration(t *testing.T) {
	ctx := context.Background()
	fixture := newIntegrationFixture(t)

	driverID := uuid.New()
	fixture.driverID = driverID
	fixture.insertDriver(t, driverID, integrationName("driver"), "1991-05-06", "Favland")

	teamID := uuid.New()
	fixture.teamID = teamID
	fixture.insertTeam(t, teamID, integrationName("team"), "Favland")

	email := integrationEmail("user")
	user := models.User{
		Name:     integrationName("user"),
		Email:    email,
		Password: "secret",
		Role:     models.RoleUser,
	}

	registeredUser, err := fixture.repo.UserRegister(ctx, user)
	if err != nil {
		t.Fatalf("UserRegister() error = %v", err)
	}

	err = fixture.pool.QueryRow(ctx, "SELECT id FROM users WHERE email = $1", email).Scan(&fixture.userID)
	if err != nil {
		t.Fatalf("select registered user id: %v", err)
	}

	loggedInUser, err := fixture.repo.UserLogin(ctx, models.User{
		Email:    email,
		Password: "secret",
	})
	if err != nil {
		t.Fatalf("UserLogin() error = %v", err)
	}
	if loggedInUser.Email != email || loggedInUser.Name != registeredUser.Name {
		t.Fatalf("UserLogin() returned unexpected user: %+v", loggedInUser)
	}

	err = fixture.repo.UserChangePassword(ctx, loggedInUser, "new-secret")
	if err != nil {
		t.Fatalf("UserChangePassword() error = %v", err)
	}

	_, err = fixture.repo.UserLogin(ctx, models.User{
		Email:    email,
		Password: "new-secret",
	})
	if err != nil {
		t.Fatalf("UserLogin() after password change error = %v", err)
	}

	err = fixture.repo.ToggleFavouriteDriver(ctx, loggedInUser, models.Driver{Id: driverID})
	if err != nil {
		t.Fatalf("ToggleFavouriteDriver(add) error = %v", err)
	}
	assertFavouriteCount(t, fixture.pool, "favourite_drivers", "driver", fixture.userID, driverID, 1)

	err = fixture.repo.ToggleFavouriteDriver(ctx, loggedInUser, models.Driver{Id: driverID})
	if err != nil {
		t.Fatalf("ToggleFavouriteDriver(remove) error = %v", err)
	}
	assertFavouriteCount(t, fixture.pool, "favourite_drivers", "driver", fixture.userID, driverID, 0)

	err = fixture.repo.ToggleFavouriteTeam(ctx, loggedInUser, models.Team{Id: teamID})
	if err != nil {
		t.Fatalf("ToggleFavouriteTeam(add) error = %v", err)
	}
	assertFavouriteCount(t, fixture.pool, "favourite_teams", "team", fixture.userID, teamID, 1)

	err = fixture.repo.ToggleFavouriteTeam(ctx, loggedInUser, models.Team{Id: teamID})
	if err != nil {
		t.Fatalf("ToggleFavouriteTeam(remove) error = %v", err)
	}
	assertFavouriteCount(t, fixture.pool, "favourite_teams", "team", fixture.userID, teamID, 0)
}

func TestRepositoryStatsIntegration(t *testing.T) {
	ctx := context.Background()
	fixture := newIntegrationFixture(t)

	driverID := uuid.New()
	fixture.driverID = driverID
	driverName := integrationName("stats-driver")
	fixture.insertDriver(t, driverID, driverName, "1988-02-03", "Statsland")

	teamID := uuid.New()
	fixture.teamID = teamID
	teamName := integrationName("stats-team")
	fixture.insertTeam(t, teamID, teamName, "Statsland")

	trackID := uuid.New()
	fixture.trackID = trackID
	fixture.insertTrack(t, trackID, integrationName("stats-track"), "Statsland", 4800, 10)

	fixture.insertOrganizer(t, integrationName("stats-organizer"))
	fixture.insertChampionship(t, 2026)
	fixture.insertManufacturer(t, integrationName("stats-manufacturer"))
	fixture.insertRaceclass(t, integrationName("stats-raceclass"))
	fixture.insertCar(t, integrationName("stats-car-model"), 2025)

	carParticipantID := uuid.New()
	fixture.carParticipantIDs = append(fixture.carParticipantIDs, carParticipantID)
	fixture.insertCarParticipant(t, carParticipantID, fixture.carID, teamID, "21")
	fixture.linkDriverToCarParticipant(t, carParticipantID, driverID)

	raceID := uuid.New()
	fixture.raceID = raceID
	fixture.insertRace(t, raceID, fixture.championshipID, trackID, integrationName("stats-race"), time.Date(2026, time.June, 10, 11, 0, 0, 0, time.UTC), 1, 100)
	fixture.insertFinish(t, raceID, carParticipantID, 1)
	fixture.insertQualifying(t, raceID, carParticipantID, 1)

	foundDriver, err := fixture.repo.GetDriverByName(ctx, driverName)
	if err != nil {
		t.Fatalf("GetDriverByName() error = %v", err)
	}
	if foundDriver.Id != driverID {
		t.Fatalf("GetDriverByName() id = %v, want %v", foundDriver.Id, driverID)
	}

	foundTeam, err := fixture.repo.GetTeamByName(ctx, teamName)
	if err != nil {
		t.Fatalf("GetTeamByName() error = %v", err)
	}
	if foundTeam.Id != teamID {
		t.Fatalf("GetTeamByName() id = %v, want %v", foundTeam.Id, teamID)
	}

	driverStats, err := fixture.repo.GetDriverStats(ctx, models.Driver{Id: driverID, Name: driverName})
	if err != nil {
		t.Fatalf("GetDriverStats() error = %v", err)
	}
	if driverStats.Driver.Id != driverID {
		t.Fatalf("GetDriverStats() driver id = %v, want %v", driverStats.Driver.Id, driverID)
	}

	teamStats, err := fixture.repo.GetTeamStats(ctx, models.Team{Id: teamID, Name: teamName})
	if err != nil {
		t.Fatalf("GetTeamStats() error = %v", err)
	}
	if teamStats.Team.Id != teamID {
		t.Fatalf("GetTeamStats() team id = %v, want %v", teamStats.Team.Id, teamID)
	}
}

func newIntegrationFixture(t *testing.T) *integrationFixture {
	t.Helper()

	databaseURL := "postgresql://postgres:1337@localhost:5432/races"
	//databaseURL := os.Getenv("TEST_DATABASE_URL")
	//if databaseURL == "" {
	//	databaseURL = os.Getenv("DATABASE_URL")
	//}
	//if databaseURL == "" {
	//	t.Skip("TEST_DATABASE_URL or DATABASE_URL is not set")
	//}

	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx, databaseURL)
	if err != nil {
		t.Fatalf("connect test database: %v", err)
	}
	t.Cleanup(func() {
		pool.Close()
	})

	fixture := &integrationFixture{
		repo: NewRepository(pool),
		pool: pool,
	}
	t.Cleanup(func() {
		fixture.cleanup(t)
	})

	return fixture
}

func (f *integrationFixture) cleanup(t *testing.T) {
	t.Helper()

	ctx := context.Background()
	if f.userID != uuid.Nil {
		f.execCleanup(t, ctx, "DELETE FROM favourite_drivers WHERE user_id = $1", f.userID)
		f.execCleanup(t, ctx, "DELETE FROM favourite_teams WHERE user_id = $1", f.userID)
		f.execCleanup(t, ctx, "DELETE FROM users WHERE id = $1", f.userID)
	}
	if f.raceID != uuid.Nil {
		f.execCleanup(t, ctx, "DELETE FROM finish WHERE race = $1", f.raceID)
		f.execCleanup(t, ctx, "DELETE FROM qualifying WHERE race = $1", f.raceID)
		f.execCleanup(t, ctx, "DELETE FROM race WHERE id = $1", f.raceID)
	}
	for _, carParticipantID := range f.carParticipantIDs {
		f.execCleanup(t, ctx, "DELETE FROM team_p WHERE car_p = $1", carParticipantID)
		f.execCleanup(t, ctx, "DELETE FROM car_p WHERE id = $1", carParticipantID)
	}
	if f.driverID != uuid.Nil {
		f.execCleanup(t, ctx, "DELETE FROM team_p WHERE driver = $1", f.driverID)
		f.execCleanup(t, ctx, "DELETE FROM driver WHERE id = $1", f.driverID)
	}
	if f.secondDriverID != uuid.Nil {
		f.execCleanup(t, ctx, "DELETE FROM team_p WHERE driver = $1", f.secondDriverID)
		f.execCleanup(t, ctx, "DELETE FROM driver WHERE id = $1", f.secondDriverID)
	}
	if f.carID != uuid.Nil {
		f.execCleanup(t, ctx, "DELETE FROM car WHERE id = $1", f.carID)
	}
	if f.teamID != uuid.Nil {
		f.execCleanup(t, ctx, "DELETE FROM team WHERE id = $1", f.teamID)
	}
	if f.trackID != uuid.Nil {
		f.execCleanup(t, ctx, "DELETE FROM track WHERE id = $1", f.trackID)
	}
	if f.championshipID != uuid.Nil {
		f.execCleanup(t, ctx, "DELETE FROM championship WHERE id = $1", f.championshipID)
	}
	if f.organizerID != uuid.Nil {
		f.execCleanup(t, ctx, "DELETE FROM organizer WHERE id = $1", f.organizerID)
	}
	if f.manufacturerID != uuid.Nil {
		f.execCleanup(t, ctx, "DELETE FROM manufacturer WHERE id = $1", f.manufacturerID)
	}
	if f.raceclassID != uuid.Nil {
		f.execCleanup(t, ctx, "DELETE FROM raceclass WHERE id = $1", f.raceclassID)
	}
}

func (f *integrationFixture) execCleanup(t *testing.T, ctx context.Context, sql string, arg uuid.UUID) {
	t.Helper()
	if _, err := f.pool.Exec(ctx, sql, arg); err != nil {
		t.Logf("cleanup failed for %q with %v: %v", sql, arg, err)
	}
}

func (f *integrationFixture) insertOrganizer(t *testing.T, name string) {
	t.Helper()
	f.organizerID = uuid.New()
	_, err := f.pool.Exec(context.Background(),
		"INSERT INTO organizer (id, name) VALUES ($1, $2)",
		f.organizerID, name)
	if err != nil {
		t.Fatalf("insert organizer: %v", err)
	}
}

func (f *integrationFixture) insertChampionship(t *testing.T, year int) {
	t.Helper()
	f.championshipID = uuid.New()
	_, err := f.pool.Exec(context.Background(),
		"INSERT INTO championship (id, organizer, year) VALUES ($1, $2, $3)",
		f.championshipID, f.organizerID, year)
	if err != nil {
		t.Fatalf("insert championship: %v", err)
	}
}

func (f *integrationFixture) insertManufacturer(t *testing.T, name string) {
	t.Helper()
	f.manufacturerID = uuid.New()
	_, err := f.pool.Exec(context.Background(),
		"INSERT INTO manufacturer (id, name, country, year_of_foundation) VALUES ($1, $2, $3, $4)",
		f.manufacturerID, name, "Testland", 2000)
	if err != nil {
		t.Fatalf("insert manufacturer: %v", err)
	}
}

func (f *integrationFixture) insertRaceclass(t *testing.T, name string) {
	t.Helper()
	f.raceclassID = uuid.New()
	_, err := f.pool.Exec(context.Background(),
		"INSERT INTO raceclass (id, name) VALUES ($1, $2)",
		f.raceclassID, name)
	if err != nil {
		t.Fatalf("insert raceclass: %v", err)
	}
}

func (f *integrationFixture) insertCar(t *testing.T, model string, year int) {
	t.Helper()
	f.carID = uuid.New()
	_, err := f.pool.Exec(context.Background(),
		"INSERT INTO car (id, manufacturer, model, year_of_production, raceclass) VALUES ($1, $2, $3, $4, $5)",
		f.carID, f.manufacturerID, model, year, f.raceclassID)
	if err != nil {
		t.Fatalf("insert car: %v", err)
	}
}

func (f *integrationFixture) insertDriver(t *testing.T, id uuid.UUID, name, birthday, nationality string) {
	t.Helper()
	_, err := f.pool.Exec(context.Background(),
		"INSERT INTO driver (id, name, birthday, nationality) VALUES ($1, $2, $3, $4)",
		id, name, birthday, nationality)
	if err != nil {
		t.Fatalf("insert driver: %v", err)
	}
}

func (f *integrationFixture) insertTeam(t *testing.T, id uuid.UUID, name, country string) {
	t.Helper()
	_, err := f.pool.Exec(context.Background(),
		"INSERT INTO team (id, name, country) VALUES ($1, $2, $3)",
		id, name, country)
	if err != nil {
		t.Fatalf("insert team: %v", err)
	}
}

func (f *integrationFixture) insertTrack(t *testing.T, id uuid.UUID, name, country string, length, turns int) {
	t.Helper()
	_, err := f.pool.Exec(context.Background(),
		"INSERT INTO track (id, name, country, lap_length, turns) VALUES ($1, $2, $3, $4, $5)",
		id, name, country, length, turns)
	if err != nil {
		t.Fatalf("insert track: %v", err)
	}
}

func (f *integrationFixture) insertCarParticipant(t *testing.T, id, carID, teamID uuid.UUID, number string) {
	t.Helper()
	_, err := f.pool.Exec(context.Background(),
		"INSERT INTO car_p (id, car, team, number) VALUES ($1, $2, $3, $4)",
		id, carID, teamID, number)
	if err != nil {
		t.Fatalf("insert car_p: %v", err)
	}
}

func (f *integrationFixture) linkDriverToCarParticipant(t *testing.T, carParticipantID, driverID uuid.UUID) {
	t.Helper()
	_, err := f.pool.Exec(context.Background(),
		"INSERT INTO team_p (car_p, driver) VALUES ($1, $2)",
		carParticipantID, driverID)
	if err != nil {
		t.Fatalf("insert team_p: %v", err)
	}
}

func (f *integrationFixture) insertRace(t *testing.T, id, championshipID, trackID uuid.UUID, name string, date time.Time, raceType, duration int) {
	t.Helper()
	_, err := f.pool.Exec(context.Background(),
		"INSERT INTO race (id, championship, track, name, date, type, duration) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		id, championshipID, trackID, name, date, raceType, duration)
	if err != nil {
		t.Fatalf("insert race: %v", err)
	}
}

func (f *integrationFixture) insertFinish(t *testing.T, raceID, carParticipantID uuid.UUID, pos int) {
	t.Helper()
	_, err := f.pool.Exec(context.Background(),
		"INSERT INTO finish (race, car_p, pos) VALUES ($1, $2, $3)",
		raceID, carParticipantID, pos)
	if err != nil {
		t.Fatalf("insert finish: %v", err)
	}
}

func (f *integrationFixture) insertQualifying(t *testing.T, raceID, carParticipantID uuid.UUID, pos int) {
	t.Helper()
	_, err := f.pool.Exec(context.Background(),
		"INSERT INTO qualifying (race, car_p, pos) VALUES ($1, $2, $3)",
		raceID, carParticipantID, pos)
	if err != nil {
		t.Fatalf("insert qualifying: %v", err)
	}
}

func assertTeamPDriverCount(t *testing.T, pool *pgxpool.Pool, carParticipantID uuid.UUID, want int) {
	t.Helper()
	var got int
	err := pool.QueryRow(context.Background(),
		"SELECT COUNT(*) FROM team_p WHERE car_p = $1", carParticipantID,
	).Scan(&got)
	if err != nil {
		t.Fatalf("count team_p: %v", err)
	}
	if got != want {
		t.Fatalf("team_p count = %d, want %d", got, want)
	}
}

func assertStandingPosition(t *testing.T, pool *pgxpool.Pool, table string, raceID, carParticipantID uuid.UUID, want int) {
	t.Helper()
	var got int
	err := pool.QueryRow(context.Background(),
		"SELECT pos FROM "+table+" WHERE race = $1 AND car_p = $2", raceID, carParticipantID,
	).Scan(&got)
	if err != nil {
		t.Fatalf("select %s pos: %v", table, err)
	}
	if got != want {
		t.Fatalf("%s pos = %d, want %d", table, got, want)
	}
}

func assertFavouriteCount(t *testing.T, pool *pgxpool.Pool, table, column string, userID, entityID uuid.UUID, want int) {
	t.Helper()
	var got int
	err := pool.QueryRow(context.Background(),
		"SELECT COUNT(*) FROM "+table+" WHERE user_id = $1 AND "+column+" = $2", userID, entityID,
	).Scan(&got)
	if err != nil {
		t.Fatalf("count %s: %v", table, err)
	}
	if got != want {
		t.Fatalf("%s count = %d, want %d", table, got, want)
	}
}

func integrationName(prefix string) string {
	return "itest_" + prefix + "_" + uuid.New().String()
}

func integrationEmail(prefix string) string {
	return integrationName(prefix) + "@example.com"
}
