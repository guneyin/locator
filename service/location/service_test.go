package location

import (
	"context"
	"encoding/json"
	"github.com/AvraamMavridis/randomcolor"
	database "github.com/guneyin/locator/db"
	"github.com/guneyin/locator/dto"
	"github.com/guneyin/locator/repository/location"
	"github.com/guneyin/locator/util"
	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/gorm"
	"os"
	"testing"
)

func newDB(withTestData bool) *gorm.DB {
	db, err := database.NewTestDB()
	if err != nil {
		panic(err)
	}

	if withTestData {
		if err = db.AutoMigrate(location.Location{}); err != nil {
			panic(err)
		}

		td := newLocationTestData()
		for _, row := range td.DataView.DataTable.Rows {
			rec := location.Location{
				Latitude:    util.StrToFloat(row.Text[8]),
				Longitude:   util.StrToFloat(row.Text[9]),
				Name:        row.Text[1],
				MarkerColor: randomcolor.GetRandomColorInHex(),
			}
			db.Create(&rec)
		}
	}

	return db
}

func newSvc(withTestData bool) *Service {
	db := newDB(withTestData)
	return New(db)
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

func TestLocationService_Add(t *testing.T) {
	ctx := context.Background()
	svc := newSvc(false)

	Convey("Add", t, func() {

		td := newLocationTestData()
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
}

func TestLocationService_List(t *testing.T) {
	ctx := context.Background()
	svc := newSvc(true)

	Convey("List", t, func() {
		locList, err := svc.List(ctx)
		So(err, ShouldBeNil)
		So(locList, ShouldNotBeNil)
	})
}

func TestLocationService_Detail(t *testing.T) {
	ctx := context.Background()
	svc := newSvc(true)

	Convey("Detail", t, func() {
		loc, err := svc.Detail(ctx, 10)
		So(err, ShouldBeNil)
		So(loc, ShouldNotBeNil)
		So(loc.Name, ShouldEqual, loc.Name)
	})
}

func TestLocationService_Edit(t *testing.T) {
	ctx := context.Background()
	svc := newSvc(true)

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

func TestLocationService_Route(t *testing.T) {
	ctx := context.Background()
	svc := newSvc(true)

	Convey("Route", t, func() {
		locMecidiyekoy := &dto.LocationDto{
			Latitude:    "41.06322081539445",
			Longitude:   "28.992461802104373",
			Name:        "Mecidiyeköy",
			MarkerColor: "#FFF000",
		}

		route, err := svc.Route(ctx, locMecidiyekoy)
		So(err, ShouldBeNil)
		So(route, ShouldNotBeNil)
	})
}
