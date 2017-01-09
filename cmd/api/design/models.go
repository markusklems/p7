package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var _ = StorageGroup("p7", func() {
	Description("This is the global storage group")
	Store("mysql", gorma.MySQL, func() {
		Description("This is the mysql relational store")
		Model("Lambda", func() {
			BuildsFrom(func() {
				Payload("lambda", "create")
				Payload("lambda", "update")
			})
			RendersTo(Lambda)
			Description("Lambda Model")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the Lambda Model PK field")
			})
			Field("name", gorma.String, func() {})
			Field("code", gorma.String, func() {})
			Field("created_at", gorma.Timestamp, func() {})
			Field("updated_at", gorma.Timestamp, func() {})
			Field("deleted_at", gorma.NullableTimestamp, func() {})
		})
	})
})
