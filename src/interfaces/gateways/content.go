package gateways

import (
	"context"
	"encoding/base64"
	"log"
	"sp/src/domains/entities"
	fiware "sp/src/drivers/ngsi"
	"sp/src/drivers/ulid"
	"sp/src/usecases/port"
	"strings"
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

func (ur *ContentRepository) Create(c *entities.Content) (receipt *entities.Content, err error) {
	id := ulid.GenerateIdentifier()
	var meta []string
	for i := 0; i < len(c.MetaData); i++ {
		metaNo := strings.Replace(base64.StdEncoding.EncodeToString(c.MetaData[i]), "=", "", -1)
		meta = append(meta, metaNo)
	}
	content := fiware.CreateEntityRequest{
		Type_: "data",
		Id:    id.Value(),
		Point: &fiware.LocationValue{
			Value: fiware.PointValue{
				Type:        "Point",
				Coordinates: []int{c.Content.X, c.Content.Y},
			},
			Type: "geo:json",
		},
		MetaOne: &fiware.MetaStringValue{
			Type:  "Text",
			Value: meta[0],
		},
		MetaTwo: &fiware.MetaStringValue{
			Type:  "Text",
			Value: meta[1],
		},
		MetaThree: &fiware.MetaStringValue{
			Type:  "Text",
			Value: meta[2],
		},
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
		Content:  c.Content,
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
		var metas [][]byte
		decOne, _ := base64.StdEncoding.DecodeString(list[i].MetaOne.Value + "=")
		decTwo, _ := base64.StdEncoding.DecodeString(list[i].MetaTwo.Value + "=")
		decThree, _ := base64.StdEncoding.DecodeString(list[i].MetaThree.Value + "=")
		metas = append(metas, decOne)
		metas = append(metas, decTwo)
		metas = append(metas, decThree)
		log.Print(metas)
		pointList := list[i].Point.Value.Coordinates
		receipt := &entities.Receipt{
			ID: list[i].Id,
			Content: entities.Point{
				X: pointList[0],
				Y: pointList[1],
			},
			MetaData: metas,
		}
		receipts = append(receipts, receipt)
	}
	return receipts, nil
}
