package gateways

import (
	"context"
	"log"
	"sp/src/domains/entities"
	fiware "sp/src/drivers/ngsi"
	"sp/src/drivers/ulid"
	"sp/src/usecases/port"

	"gorm.io/gorm"
)

type ContentSQLHandler interface {
	Find(interface{}, ...interface{}) (*entities.Receipt, error)
	First(interface{}, ...interface{}) (*entities.Receipt, error)
	Create(interface{}) (*entities.Receipt, error)
	Save(interface{}) error
	Delete(interface{}) *entities.Content
	Where(interface{}, ...interface{}) *entities.Content
}

type ContentRepository struct {
	Conn *gorm.DB
	ContentSQLHandler
}

func NewContentRepository(conn *gorm.DB) port.ContentRepository {
	return &ContentRepository{
		Conn: conn,
	}
}

func (ur *ContentRepository) Find(id string) (*entities.Content, error) {
	var content = &entities.Content{}
	content.ID = id
	err := ur.Conn.First(&content).Error
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (ur *ContentRepository) Create(c *entities.Content) (receipt *entities.Content, err error) {
	id := ulid.GenerateIdentifier()
	content := fiware.CreateEntityRequest{
		Type_:    "Toet",
		Id:       id.Value(),
	}
	cfg := fiware.NewConfiguration()
	client := fiware.NewAPIClient(cfg)

	ctx := context.Background()
	res, err := client.EntitiesApi.CreateEntity(ctx, content, "application/json", &fiware.EntitiesApiCreateEntityOpts{})
	log.Print(res)
	if err != nil {
		return nil, err
	}
	return &entities.Content{
		ID:       id.Value(),
		Address:  c.Address,
		Content:  entities.SampleData{},
		MetaData: c.MetaData,
		HashData: c.HashData,
	}, nil
}

func (ur *ContentRepository) All() (receipt []*entities.Content, err error) {
	cfg := fiware.NewConfiguration()
	client := fiware.NewAPIClient(cfg)

	ctx := context.Background()
	list, res, err := client.EntitiesApi.ListEntities(ctx, &fiware.EntitiesApiListEntitiesOpts{})
	log.Print(res)
	log.Print(list)
	if err != nil {
		return nil, err
	}
	var l []*entities.Content
	return l, nil
}
