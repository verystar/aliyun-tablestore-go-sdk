package search

import (
	"github.com/golang/protobuf/proto"
	"github.com/verystar/aliyun-tablestore-go-sdk/tablestore/otsprotocol"
)

type GeoBoundingBoxQuery struct {
	FieldName   string
	TopLeft     string
	BottomRight string
}

func (q *GeoBoundingBoxQuery) Type() QueryType {
	return QueryType_GeoBoundingBoxQuery
}

func (q *GeoBoundingBoxQuery) Serialize() ([]byte, error) {
	query := &otsprotocol.GeoBoundingBoxQuery{}
	query.FieldName = &q.FieldName
	query.TopLeft = &q.TopLeft
	query.BottomRight = &q.BottomRight
	data, err := proto.Marshal(query)
	return data, err
}

func (q *GeoBoundingBoxQuery) ProtoBuffer() (*otsprotocol.Query, error) {
	return BuildPBForQuery(q)
}
