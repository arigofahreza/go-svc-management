package services

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"go-svc-management/src/models"
	"go-svc-management/src/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AccountServices struct {
	Data []*models.AccountView `json:"data,omitempty"`
}

func InitUserService() *AccountServices {
	return &AccountServices{}
}

func (service AccountServices) CreateAccountService(ctx context.Context, mongoCollection *mongo.Collection, accountModel models.AccountModel) ([]*models.AccountView, error) {
	hashId, err := bson.Marshal(accountModel)
	if err != nil {
		return nil, err
	}
	hash := md5.Sum(hashId)
	accountModel.Id = hex.EncodeToString(hash[:])
	accountModel.Password, err = utils.HashPassword(accountModel.Password)
	if err != nil {
		return nil, err
	}
	_, err = mongoCollection.InsertOne(ctx, accountModel)
	if err != nil {
		return nil, err
	}
	accountView := models.AccountView{}
	jsonAccount, err := json.Marshal(accountModel)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(jsonAccount, &accountView)
	service.Data = append(service.Data, &accountView)
	return service.Data, nil
}

func (service AccountServices) GetAllAccountService(ctx context.Context, mongoCollection *mongo.Collection, page int, size int) ([]*models.AccountView, error) {
	option := options.FindOptions{}
	if page == 1 {
		option.SetSkip(0)
		option.SetLimit(int64(size))
	}
	option.SetSkip(int64((page - 1) * size))
	option.SetLimit(int64(size))
	cursor, err := mongoCollection.Find(ctx, bson.D{{}}, &option)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var account models.AccountView
		err := cursor.Decode(&account)
		if err != nil {
			return nil, err
		}
		service.Data = append(service.Data, &account)
	}
	return service.Data, nil
}

func (service AccountServices) GetAccountByIdService(ctx context.Context, mongoCollection *mongo.Collection, id string) ([]*models.AccountView, error) {
	filter := bson.M{"_id": id}
	var account models.AccountView
	err := mongoCollection.FindOne(ctx, filter).Decode(&account)
	if err == mongo.ErrNoDocuments {
		return service.Data, nil
	} else if err != nil {
		return nil, err
	}
	service.Data = append(service.Data, &account)
	return service.Data, nil
}

func (service AccountServices) UpdateAccountService(ctx context.Context, mongoCollection *mongo.Collection, accountView models.AccountView) ([]*models.AccountView, error) {
	filter := bson.M{"_id": accountView.Id}
	_, err := mongoCollection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: accountView}})
	if err == mongo.ErrNoDocuments {
		return service.Data, nil
	} else if err != nil {
		return nil, err
	}
	service.Data = append(service.Data, &accountView)
	return service.Data, nil
}

func (service AccountServices) DeleteAccountService(ctx context.Context, mongoCollection *mongo.Collection, id string) ([]*models.AccountView, error) {
	filter := bson.M{"_id": id}
	var account models.AccountView
	err := mongoCollection.FindOneAndDelete(ctx, filter).Decode(&account)
	if err == mongo.ErrNoDocuments {
		return service.Data, nil
	} else if err != nil {
		return nil, err
	}
	service.Data = append(service.Data, &account)
	return service.Data, nil
}
