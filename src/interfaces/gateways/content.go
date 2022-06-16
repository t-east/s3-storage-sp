package gateways

import (
	"context"
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
	var metaStr string
	for i := 0; i < len(c.MetaData); i++ {
		if i != 0 {
			metaStr += ",,,,"
		}
		metaStr += strings.Replace(c.MetaData[i], "=", "", -1)
		log.Print(metaStr)
	}
	content := fiware.CreateEntityRequest{
		Type_: "Azm",
		Id:    id.Value(),
		Point: &fiware.LocationValue{
			Value: fiware.PointValue{
				Type:        "Point",
				Coordinates: []int{c.Content.X, c.Content.Y},
			},
			Type: "geo:json",
		},
		Meta: &fiware.MetaStringValue{
			Type:  "Text",
			Value: metaStr,
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
		// var metas []string
		// for j:=0;i<len(list[i].Metas.Value);i++ {
		// 	metas = append(metas, list[i].Metas.Value[j].Value)
		// }
		metaList := strings.Split(list[i].Meta.Value, ",,,,")
		log.Print(list[i].Meta.Value)
		var metaListReplaced []string
		for j := 0; j < len(metaList); j++ {
			log.Print(metaList[j])
			m := metaList[j]+"="
			metaListReplaced = append(metaListReplaced, m)
		}
		pointList := list[i].Point.Value.Coordinates
		receipt := &entities.Receipt{
			ID: list[i].Id,
			Content: entities.Point{
				X: pointList[0],
				Y: pointList[1],
			},
			MetaData: metaListReplaced,
			HashData: []string{},
		}
		receipts = append(receipts, receipt)
	}
	return receipts, nil
}
