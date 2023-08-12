package test

import (
	"github.com/mixedmachine/exoplanet-data-pipeline/data-ingestion-service/src/pkg/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueryBuilder(t *testing.T) {
	q := api.NewQueryBuilder()

	// Test AddSelect method
	t.Run("Test AddSelect", func(t *testing.T) {
		q.AddSelect("name, age")
		assert.Equal(t, "?query=select+name, age", q.GetQuery())
		q.ResetQuery()
	})

	// Test AddFrom method
	t.Run("Test AddFrom", func(t *testing.T) {
		q.AddSelect("name, age").AddFrom("users")
		assert.Equal(t, "?query=select+name, age+from+users", q.GetQuery())
		q.ResetQuery()
	})

	// Test AddWhere method
	t.Run("Test AddWhere", func(t *testing.T) {
		q.AddSelect("name, age").AddFrom("users").AddWhere()
		assert.Equal(t, "?query=select+name, age+from+users+where", q.GetQuery())
		q.ResetQuery()
	})

	// Test AddWhereParameter method
	t.Run("Test AddWhereParameter", func(t *testing.T) {
		q.AddSelect("name, age").AddFrom("users").AddWhere().AddWhereParameter("age", ">", "30")
		assert.Equal(t, "?query=select+name, age+from+users+where+age+>+'30'", q.GetQuery())
		q.ResetQuery()
	})

	// Test AddAndWhereParameter method
	t.Run("Test AddAndWhereParameter", func(t *testing.T) {
		q.AddSelect("name, age").AddFrom("users").AddWhere().AddWhereParameter("age", ">", "30").AddAndWhereParameter("name", "=", "John")
		assert.Equal(t, "?query=select+name, age+from+users+where+age+>+'30'+and+name+=+'John'", q.GetQuery())
		q.ResetQuery()
	})

	// Test AddOrWhereParameter method
	t.Run("Test AddOrWhereParameter", func(t *testing.T) {
		q.AddSelect("name, age").AddFrom("users").AddWhere().AddWhereParameter("age", ">", "30").AddOrWhereParameter("name", "=", "John")
		assert.Equal(t, "?query=select+name, age+from+users+where+age+>+'30'+or+name+=+'John'", q.GetQuery())
		q.ResetQuery()
	})

	// Test AddFormat method
	t.Run("Test AddFormat", func(t *testing.T) {
		q.AddSelect("name, age").AddFrom("users").AddFormat("json")
		assert.Equal(t, "?query=select+name, age+from+users&format=json", q.GetQuery())
		q.ResetQuery()
	})

	// Test Build method
	t.Run("Test Build", func(t *testing.T) {
		query := q.AddSelect("name, age").AddFrom("users").AddWhere().AddWhereParameter("age", ">", "30").AddFormat("json").Build()
		assert.Equal(t, "?query=select+name, age+from+users+where+age+>+'30'&format=json", query)
	})

	// Test ResetQuery method
	t.Run("Test ResetQuery", func(t *testing.T) {
		q.ResetQuery()
		assert.Equal(t, "?query=", q.GetQuery())
	})
}
