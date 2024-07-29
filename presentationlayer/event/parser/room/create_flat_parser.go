package room

import (
	"fmt"
	"hc/presentationlayer/event/viewmodel"
	"strings"
)

const expectedSepSize = 6

func ParseCreateFlat(body []byte) (viewmodel.CreateFlat, error) {
	flatData := strings.Split(string(body), "/")
	if len(flatData) != expectedSepSize {
		return viewmodel.CreateFlat{}, fmt.Errorf("expected len(flatData) to be %d, got %d", expectedSepSize, len(flatData))
	}

	model := viewmodel.CreateFlat{
		RoomName:   flatData[2],
		RoomModel:  flatData[3],
		RoomAccess: flatData[4],
		ShowName:   flatData[5] == "1",
	}

	return model, nil
}
