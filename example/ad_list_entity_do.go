package example

import (
	"context"

	"github.com/suifengpiao14/apihandler"
	"github.com/suifengpiao14/stream"
)

func AdListDoFn(ctx context.Context, input *AdListInput) (apihandler.OutputI, error) {
	output := input.Output
	output.Items = make([]AdListOutputItem, 0)
	output.Pagination.Index = input.Index
	output.Pagination.Size = input.Size
	output.Pagination.Total = 100
	item := AdListOutputItem{
		ID: "1",
	}
	output.Items = append(output.Items, item)
	return output, nil
}

func (i *AdListInput) GetStream() (s stream.StreamInterface) {
	s, err := apihandler.LineschemaPacketStream(i, i)
	if err != nil {
		panic(err)
	}
	return s
}
