package mongodb

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"sort"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) ListDrivers(ctx context.Context, query string) ([]models.Driver, error) {
	var docs []driverDoc
	if err := r.findAll(ctx, "driver", bson.M{}, &docs); err != nil {
		return nil, err
	}

	drivers := make([]models.Driver, 0, len(docs))
	for _, doc := range docs {
		if query != "" && !containsFold(doc.Name, query) {
			continue
		}
		driver, err := doc.model()
		if err != nil {
			return nil, err
		}
		drivers = append(drivers, driver)
	}

	sort.Slice(drivers, func(i, j int) bool {
		return drivers[i].Name < drivers[j].Name
	})

	return drivers, nil
}

func (r *Repository) ListTeams(ctx context.Context, query string) ([]models.Team, error) {
	var docs []teamDoc
	if err := r.findAll(ctx, "team", bson.M{}, &docs); err != nil {
		return nil, err
	}

	teams := make([]models.Team, 0, len(docs))
	for _, doc := range docs {
		if query != "" && !containsFold(doc.Name, query) {
			continue
		}
		team, err := doc.model()
		if err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}

	sort.Slice(teams, func(i, j int) bool {
		return teams[i].Name < teams[j].Name
	})

	return teams, nil
}

func (r *Repository) GetDriverStats(ctx context.Context, driver models.Driver) (models.DriverStats, error) {
	data, err := r.loadStatsData(ctx)
	if err != nil {
		return models.DriverStats{}, err
	}

	return models.DriverStats{
		Driver: driver,
		Stats:  data.driverStats(driver.Id.String()),
	}, nil
}

func (r *Repository) GetTeamStats(ctx context.Context, team models.Team) (models.TeamStats, error) {
	data, err := r.loadStatsData(ctx)
	if err != nil {
		return models.TeamStats{}, err
	}

	return models.TeamStats{
		Team:  team,
		Stats: data.teamStats(team.Id.String()),
	}, nil
}

func (r *Repository) GetDriverByName(ctx context.Context, name string) (models.Driver, error) {
	var doc driverDoc
	err := r.collection("driver").FindOne(ctx, bson.M{"name": name}).Decode(&doc)
	if err != nil {
		if isNoDocuments(err) {
			return models.Driver{}, shared.ErrorNotFound
		}
		return models.Driver{}, err
	}
	return doc.model()
}

func (r *Repository) GetTeamByName(ctx context.Context, name string) (models.Team, error) {
	var doc teamDoc
	err := r.collection("team").FindOne(ctx, bson.M{"name": name}).Decode(&doc)
	if err != nil {
		if isNoDocuments(err) {
			return models.Team{}, shared.ErrorNotFound
		}
		return models.Team{}, err
	}
	return doc.model()
}

func (d statsData) driverStats(driverID string) models.Stats {
	lineups := d.driverLineups(driverID)
	raceIDs := make(map[string]struct{})
	championships := make(map[string]struct{})
	stats := models.Stats{}

	for _, finish := range d.finishes {
		if _, ok := lineups[finish.CarP]; !ok {
			continue
		}
		race, ok := d.races[finish.Race]
		if !ok {
			continue
		}
		raceIDs[finish.Race] = struct{}{}
		championships[race.Championship] = struct{}{}
		if finish.Pos == 1 {
			stats.TotalWins++
		}
		if finish.Pos <= 3 {
			stats.TotalPodiums++
		}
		stats.TotalPoints += pointsForFinish(finish.Pos, race.Type, race.Duration)
		stats.BestFinish = minPositive(stats.BestFinish, finish.Pos)
	}
	stats.TotalRaces = len(raceIDs)

	for _, qualifying := range d.qualifying {
		if _, ok := lineups[qualifying.CarP]; !ok {
			continue
		}
		if qualifying.Pos == 1 {
			stats.TotalPoles++
		}
		stats.BestQualifying = minPositive(stats.BestQualifying, qualifying.Pos)
	}

	positionByChampionshipRaceClass := make(map[string]int)
	for championshipID := range d.allChampionships() {
		if _, ok := championships[championshipID]; !ok {
			continue
		}
		for raceClass, points := range groupPersonalPointsByRaceClass(d.calculatePersonalPoints(championshipID)) {
			sort.Slice(points, func(i, j int) bool {
				if points[i].points == points[j].points {
					return points[i].carParticipantID < points[j].carParticipantID
				}
				return points[i].points > points[j].points
			})
			for idx, point := range points {
				if lineups[point.carParticipantID] != raceClass {
					continue
				}
				key := rowKey(championshipID, raceClass)
				position := idx + 1
				current := positionByChampionshipRaceClass[key]
				if current == 0 || position < current {
					positionByChampionshipRaceClass[key] = position
				}
			}
		}
	}

	for _, position := range positionByChampionshipRaceClass {
		if position == 1 {
			stats.ChampionshipsWins++
		}
		stats.BestChampionshipPosition = minPositive(stats.BestChampionshipPosition, position)
	}

	return stats
}

func (d statsData) teamStats(teamID string) models.Stats {
	lineups := d.teamLineups(teamID)
	raceIDs := make(map[string]struct{})
	championships := make(map[string]struct{})
	stats := models.Stats{}

	for _, finish := range d.finishes {
		if _, ok := lineups[finish.CarP]; !ok {
			continue
		}
		race, ok := d.races[finish.Race]
		if !ok {
			continue
		}
		raceIDs[finish.Race] = struct{}{}
		championships[race.Championship] = struct{}{}
		if finish.Pos == 1 {
			stats.TotalWins++
		}
		if finish.Pos <= 3 {
			stats.TotalPodiums++
		}
		stats.TotalPoints += pointsForFinish(finish.Pos, race.Type, race.Duration)
		stats.BestFinish = minPositive(stats.BestFinish, finish.Pos)
	}
	stats.TotalRaces = len(raceIDs)

	for _, qualifying := range d.qualifying {
		if _, ok := lineups[qualifying.CarP]; !ok {
			continue
		}
		if qualifying.Pos == 1 {
			stats.TotalPoles++
		}
		stats.BestQualifying = minPositive(stats.BestQualifying, qualifying.Pos)
	}

	positions := make([]int, 0)
	for championshipID := range d.allChampionships() {
		if _, ok := championships[championshipID]; !ok {
			continue
		}
		for raceClass, teamPoints := range d.teamPointsByRaceClass(championshipID) {
			sort.Slice(teamPoints, func(i, j int) bool {
				if teamPoints[i].points == teamPoints[j].points {
					return teamPoints[i].teamID < teamPoints[j].teamID
				}
				return teamPoints[i].points > teamPoints[j].points
			})
			for idx, point := range teamPoints {
				if point.teamID != teamID {
					continue
				}
				if lineupsForRaceClass(lineups, raceClass) {
					positions = append(positions, idx+1)
				}
			}
		}
	}

	for _, position := range positions {
		if position == 1 {
			stats.ChampionshipsWins++
		}
		stats.BestChampionshipPosition = minPositive(stats.BestChampionshipPosition, position)
	}

	return stats
}

func (d statsData) allChampionships() map[string]struct{} {
	championships := make(map[string]struct{})
	for _, race := range d.races {
		if race.Championship != "" {
			championships[race.Championship] = struct{}{}
		}
	}
	return championships
}

func groupPersonalPointsByRaceClass(points []personalPoints) map[string][]personalPoints {
	grouped := make(map[string][]personalPoints)
	for _, point := range points {
		grouped[point.raceClass] = append(grouped[point.raceClass], point)
	}
	return grouped
}

func (d statsData) teamPointsByRaceClass(championshipID string) map[string][]personalPoints {
	grouped := make(map[string]map[string]personalPoints)
	for _, point := range d.calculatePersonalPoints(championshipID) {
		if _, ok := grouped[point.raceClass]; !ok {
			grouped[point.raceClass] = make(map[string]personalPoints)
		}
		teamPoint := grouped[point.raceClass][point.teamID]
		teamPoint.teamID = point.teamID
		teamPoint.raceClass = point.raceClass
		teamPoint.points += point.points
		grouped[point.raceClass][point.teamID] = teamPoint
	}

	result := make(map[string][]personalPoints, len(grouped))
	for raceClass, teamPoints := range grouped {
		for _, point := range teamPoints {
			result[raceClass] = append(result[raceClass], point)
		}
	}
	return result
}

func lineupsForRaceClass(lineups map[string]string, raceClass string) bool {
	for _, lineupRaceClass := range lineups {
		if lineupRaceClass == raceClass {
			return true
		}
	}
	return false
}
