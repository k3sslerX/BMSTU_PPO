package mongodb

import (
	"RacingGuru/internal/models"
	"RacingGuru/internal/shared"
	"context"
	"sort"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) UserRegister(ctx context.Context, user models.User) (models.User, error) {
	err := r.collection("users").FindOne(ctx, bson.M{"email": user.Email}).Err()
	if err == nil {
		return models.User{}, shared.ErrorUserAlreadyExists
	}
	if !isNoDocuments(err) {
		return models.User{}, err
	}

	id := newID()
	_, err = r.collection("users").InsertOne(ctx, bson.M{
		"_id":          id,
		"id":           id,
		"name":         user.Name,
		"email":        user.Email,
		"passwordhash": user.Password,
		"role":         string(user.Role),
		"created_at":   time.Now().UTC().Truncate(time.Second),
	})
	if err != nil {
		return models.User{}, err
	}

	parsedID, err := parseID(id)
	if err != nil {
		return models.User{}, err
	}
	user.Id = parsedID
	return user, nil
}

func (r *Repository) HasAdmin(ctx context.Context) (bool, error) {
	err := r.collection("users").FindOne(ctx, bson.M{"role": string(models.RoleAdmin)}).Err()
	if err == nil {
		return true, nil
	}
	if isNoDocuments(err) {
		return false, nil
	}
	return false, err
}

func (r *Repository) UserLogin(ctx context.Context, user models.User) (models.User, error) {
	var doc userDoc
	err := r.collection("users").FindOne(ctx, bson.M{"email": user.Email}).Decode(&doc)
	if err != nil {
		if isNoDocuments(err) {
			return user, shared.ErrorIncorrectPassword
		}
		return user, err
	}
	if doc.PasswordHash != user.Password {
		return user, shared.ErrorIncorrectPassword
	}

	id, err := parseID(doc.ID)
	if err != nil {
		return user, err
	}
	user.Id = id
	user.Name = doc.Name
	user.Role = models.Role(doc.Role)
	return user, nil
}

func (r *Repository) UserChangePassword(ctx context.Context, user models.User, pwd string) error {
	result, err := r.collection("users").UpdateOne(
		ctx,
		bson.M{"id": idString(user.Id)},
		bson.M{"$set": bson.M{"passwordhash": pwd}},
	)
	if err != nil {
		return err
	}
	return updateResultError(result)
}

func (r *Repository) GetUserByID(ctx context.Context, user models.User) (models.User, error) {
	var doc userDoc
	err := r.collection("users").FindOne(ctx, bson.M{"id": idString(user.Id)}).Decode(&doc)
	if err != nil {
		if isNoDocuments(err) {
			return models.User{}, shared.ErrorNotFound
		}
		return models.User{}, err
	}

	id, err := parseID(doc.ID)
	if err != nil {
		return models.User{}, err
	}
	return models.User{
		Id:    id,
		Name:  doc.Name,
		Email: doc.Email,
		Role:  models.Role(doc.Role),
	}, nil
}

func (r *Repository) ListFavouriteDrivers(ctx context.Context, user models.User) ([]models.Driver, error) {
	var favourites []favouriteDriverDoc
	if err := r.findAll(ctx, "favourite_drivers", bson.M{"user_id": idString(user.Id)}, &favourites); err != nil {
		return nil, err
	}

	drivers := make([]models.Driver, 0, len(favourites))
	for _, favourite := range favourites {
		var doc driverDoc
		err := r.collection("driver").FindOne(ctx, bson.M{"id": favourite.Driver}).Decode(&doc)
		if err != nil {
			if isNoDocuments(err) {
				continue
			}
			return nil, err
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

func (r *Repository) ListFavouriteTeams(ctx context.Context, user models.User) ([]models.Team, error) {
	var favourites []favouriteTeamDoc
	if err := r.findAll(ctx, "favourite_teams", bson.M{"user_id": idString(user.Id)}, &favourites); err != nil {
		return nil, err
	}

	teams := make([]models.Team, 0, len(favourites))
	for _, favourite := range favourites {
		var doc teamDoc
		err := r.collection("team").FindOne(ctx, bson.M{"id": favourite.Team}).Decode(&doc)
		if err != nil {
			if isNoDocuments(err) {
				continue
			}
			return nil, err
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

func (r *Repository) ToggleFavouriteDriver(ctx context.Context, user models.User, driver models.Driver) error {
	userID := idString(user.Id)
	driverID := idString(driver.Id)
	key := rowKey(userID, driverID)

	result, err := r.collection("favourite_drivers").DeleteOne(ctx, bson.M{"_id": key})
	if err != nil {
		return err
	}
	if result.DeletedCount > 0 {
		return nil
	}

	_, err = r.collection("favourite_drivers").UpdateOne(
		ctx,
		bson.M{"_id": key},
		bson.M{"$setOnInsert": bson.M{
			"_id":     key,
			"user_id": userID,
			"driver":  driverID,
		}},
		options.Update().SetUpsert(true),
	)
	return err
}

func (r *Repository) ToggleFavouriteTeam(ctx context.Context, user models.User, team models.Team) error {
	userID := idString(user.Id)
	teamID := idString(team.Id)
	key := rowKey(userID, teamID)

	result, err := r.collection("favourite_teams").DeleteOne(ctx, bson.M{"_id": key})
	if err != nil {
		return err
	}
	if result.DeletedCount > 0 {
		return nil
	}

	_, err = r.collection("favourite_teams").UpdateOne(
		ctx,
		bson.M{"_id": key},
		bson.M{"$setOnInsert": bson.M{
			"_id":     key,
			"user_id": userID,
			"team":    teamID,
		}},
		options.Update().SetUpsert(true),
	)
	return err
}
