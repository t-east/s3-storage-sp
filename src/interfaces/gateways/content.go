package gateways

import (
	"context"
	"log"
	"sp/src/domains/entities"
	fiware "sp/src/drivers/ngsi"
	"sp/src/drivers/ulid"
	"sp/src/usecases/port"
)

type ContentRepository struct {
}

func NewContentRepository() port.ContentRepository {
	return &ContentRepository{}
}

func (ur *ContentRepository) Find(id string) (*entities.Content, error) {
	var content = &entities.Content{}
	return content, nil
}

type InterValue struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

func (ur *ContentRepository) Create(c *entities.Content) (receipt *entities.Content, err error) {
	id := ulid.GenerateIdentifier()
	h := fiware.StringValue{
		Value: "12",
		Type:  "string",
	}
	content := fiware.CreateEntityRequest{
		Type_:       "Tos",
		Id:          id.Value(),
		Temperature: &h,
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

func (ur *ContentRepository) All() (receipt []*entities.Receipt, err error) {
	cfg := fiware.NewConfiguration()
	client := fiware.NewAPIClient(cfg)

	ctx := context.Background()
	list, _, err := client.EntitiesApi.ListEntities(ctx, &fiware.EntitiesApiListEntitiesOpts{})
	if err != nil {
		return nil, err
	}
	var receipts []*entities.Receipt
	for i := 0; i < len(list); i++ {
		receipt := &entities.Receipt{
			ID:       list[i].Id,
			Content:  entities.SampleData{},
			MetaData: [][]byte{},
			HashData: []string{},
			Str:      list[i].Temperature.Value,
		}
		receipts = append(receipts, receipt)
	}
	return receipts, nil
}
