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

	id := ulid.GenerateIdentifier()
	if err != nil {
		return nil, err
	}
	return contents, nil
}

func (ur *ContentRepository) Create(c *entities.Content) (receipt *entities.Content, err error) {
	content := fiware.CreateEntityRequest{
		Type_: "string",
		Id:    "urn:ngsi-ld:Store:001",
	}
	cfg := fiware.NewConfiguration()
	client := fiware.NewAPIClient(cfg)

	ctx := context.Background()
	res, err := client.EntitiesApi.CreateEntity(ctx, content, "application/json", &fiware.EntitiesApiCreateEntityOpts{})
	if err != nil {
		return nil, err
	}
	log.Print(res.Request.Response)
	return c, nil
}
