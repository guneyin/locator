package location

import (
	"context"
	"encoding/json"
	"github.com/AvraamMavridis/randomcolor"
	database "github.com/guneyin/locator/db"
	"github.com/guneyin/locator/dto"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

var (
	svc *Service
	td  *testData
)

func init() {
	db, err := database.NewTestDB()
	if err != nil {
		panic(err)
	}

	svc = New(db)
	td = newLocationTestData()
}

type testData struct {
	DataView struct {
		DataTable struct {
			TotalRows    int `json:"totalRows"`
			RowsReturned int `json:"rowsReturned"`
			Rows         []struct {
				Text []string `json:"text"`
			} `json:"rows"`
		} `json:"dataTable"`
	} `json:"dataView"`
}

func newLocationTestData() *testData {
	file, err := os.ReadFile("../../testdata/locations.json")
	if err != nil {
		panic(err)
	}

	data := new(testData)
	err = json.Unmarshal(file, &data)
	if err != nil {
		panic(err)
	}

	return data
}

func TestLocationService(t *testing.T) {
	ctx := context.Background()

	Convey("Add", t, func() {
		for _, row := range td.DataView.DataTable.Rows {
			loc := &dto.LocationDto{
				Latitude:    row.Text[8],
				Longitude:   row.Text[9],
				Name:        row.Text[1],
				MarkerColor: randomcolor.GetRandomColorInHex(),
			}

			added, err := svc.Add(ctx, loc)
			So(err, ShouldBeNil)
			So(added, ShouldNotBeNil)
		}
	})

	Convey("List", t, func() {
		locList, err := svc.List(ctx)
		So(err, ShouldBeNil)
		So(locList, ShouldNotBeNil)
		size := len(locList.Items)
		So(size, ShouldEqual, td.DataView.DataTable.TotalRows)
	})

	Convey("Detail", t, func() {
		loc, err := svc.Detail(ctx, 10)
		So(err, ShouldBeNil)
		So(loc, ShouldNotBeNil)
		So(loc.Name, ShouldEqual, loc.Name)
	})

	Convey("Edit", t, func() {
		locBesiktas, err := svc.Detail(ctx, 10)
		So(err, ShouldBeNil)
		So(locBesiktas, ShouldNotBeNil)
		So(locBesiktas.Name, ShouldEqual, "Beşiktaş")

		locEdited := new(dto.LocationDto)
		locEdited.Name = "Besiktas"

		edited, err := svc.Edit(ctx, 10, locEdited)
		So(err, ShouldBeNil)
		So(edited, ShouldNotBeNil)
		So(edited.Id, ShouldEqual, 10)
		So(edited.Name, ShouldEqual, locEdited.Name)
		So(edited.Latitude, ShouldEqual, locBesiktas.Latitude)
		So(edited.Longitude, ShouldEqual, locBesiktas.Longitude)
		So(edited.MarkerColor, ShouldEqual, locBesiktas.MarkerColor)
	})
}
