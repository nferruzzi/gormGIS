package gormGIS_test

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/nferruzzi/gormGIS"

	"testing"
)

var (
	DB gorm.DB
)

func init() {
	var err error
	fmt.Println("testing postgres...")
	DB, err = gorm.Open("postgres", "user=gorm dbname=gormGIS sslmode=disable")
	DB.LogMode(true)

	DB.Exec("CREATE EXTENSION postgis")
	DB.Exec("CREATE EXTENSION postgis_topology")

	//DB.LogMode(false)

	if err != nil {
		panic(fmt.Sprintf("No error should happen when connect database, but got %+v", err))
	}

	DB.DB().SetMaxIdleConns(10)
}

type TestPoint struct {
	Location gormGIS.GeoPoint `sql:"type:geometry(Geometry,4326)"`
}

func TestGeoPoint(t *testing.T) {
	if DB.CreateTable(&TestPoint{}) == nil {
		t.Errorf("Can't create table")
	}

	p := TestPoint{
		Location: gormGIS.GeoPoint{
			Lat: 43.76857094631136,
			Lng: 11.292383687705296,
		},
	}

	if DB.Create(&p) == nil {
		t.Errorf("Can't create row")
	}

	var res TestPoint
	DB.First(&res)

	if res.Location.Lat != 43.76857094631136 {
		t.Errorf("Latitude not correct")
	}

	if res.Location.Lng != 11.292383687705296 {
		t.Errorf("Longitude not correct")
	}
}
