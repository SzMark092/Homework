package db

import (
	"time"
)

/*
https://github.com/go-pg/pg/wiki/Model-Definition
*/

type DataPointDescription struct {
	ID    int64 `pg:",pk"`
	Name  string
	Title string
	Max   int
	Min   int
}

type Module struct {
	ID    int64 `pg:",pk"`
	Name  string
	Title string
	Max   int
}

type DataPoint struct {
	ID                   int64 `pg:",pk"`
	Module               *Module
	DataPointDescription *DataPointDescription
	Value                float64
	Timestamp            time.Time
}
