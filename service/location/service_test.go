package location

import (
	"context"
	database "github.com/guneyin/locator/db"
	"github.com/guneyin/locator/dto"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var svc *Service

func init() {
	db, err := database.NewTestDB()
	if err != nil {
		panic(err)
	}

	svc = New(db)
}

func TestLocationService(t *testing.T) {
	ctx := context.Background()

	locAdd := &dto.LocationDto{
		Latitude:    41.00641613914117,
		Longitude:   28.975732826139218,
		Name:        "Sultanahmet Meydanı",
		MarkerColor: "#3cb371",
	}

	var locId uint
	Convey("Add", t, func() {
		added, err := svc.Add(ctx, locAdd)
		So(err, ShouldBeNil)
		So(added, ShouldNotBeNil)

		locId = added.Id
	})

	Convey("List", t, func() {
		for range 10 {
			_, _ = svc.Add(ctx, locAdd)
		}

		locList, err := svc.List(ctx)
		So(err, ShouldBeNil)
		So(locList, ShouldNotBeNil)
		size := len(locList.Items)
		So(size, ShouldEqual, 11)
	})

	Convey("Detail", t, func() {
		loc, err := svc.Detail(ctx, locId)
		So(err, ShouldBeNil)
		So(loc, ShouldNotBeNil)
		So(loc.Name, ShouldEqual, locAdd.Name)
	})

	Convey("Edit", t, func() {
		loc := new(dto.LocationDto)
		loc.Name = "Sultanahmet Square"

		edited, err := svc.Edit(ctx, locId, loc)
		So(err, ShouldBeNil)
		So(edited, ShouldNotBeNil)
		So(edited.Id, ShouldEqual, locId)
		So(edited.Name, ShouldEqual, loc.Name)
		So(edited.Latitude, ShouldEqual, locAdd.Latitude)
		So(edited.Longitude, ShouldEqual, locAdd.Longitude)
		So(edited.MarkerColor, ShouldEqual, locAdd.MarkerColor)
	})
}
