package graphql

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExploreBuilder(t *testing.T) {

	t.Run("Simple Explore", func(t *testing.T) {
		conMock := &MockRunREST{}

		builder := Explore{
			connection: conMock,
		}

		query := builder.WithFields([]ExploreFields{Certainty, Beacon}).WithConcepts([]string{"Cheese", "pineapple"}).build()

		expected := `{Explore(concepts: ["Cheese","pineapple"] ){certainty beacon }}`
		assert.Equal(t, expected, query)
	})

	t.Run("Explore limit and certainty", func(t *testing.T) {
		conMock := &MockRunREST{}

		builder := Explore{
			connection: conMock,
		}

		query := builder.WithFields([]ExploreFields{Beacon}).WithConcepts([]string{"Cheese"}).WithLimit(5).WithCertainty(0.71).build()

		expected := `{Explore(concepts: ["Cheese"] limit: 5 certainty: 0.71 ){beacon }}`
		assert.Equal(t, expected, query)
	})

	t.Run("Explore with move", func(t *testing.T) {
		conMock := &MockRunREST{}

		builder := Explore{
			connection: conMock,
		}

		fields := []ExploreFields{Beacon}
		concepts := []string{"Cheese"}
		moveTo := &MoveParameters{
			Concepts: []string{"pizza", "pineapple"},
			Force:    0.2,
		}
		moveAwayFrom := &MoveParameters{
			Concepts: []string{"fish"},
			Force:    0.1,
		}
		query := builder.WithFields(fields).WithConcepts(concepts).WithMoveTo(moveTo).WithMoveAwayFrom(moveAwayFrom).build()

		expected := `{Explore(concepts: ["Cheese"] moveTo: {concepts: ["pizza","pineapple"] force: 0.2} moveAwayFrom: {concepts: ["fish"] force: 0.1} ){beacon }}`
		assert.Equal(t, expected, query)
	})

	t.Run("Missuse", func(t *testing.T) {
		conMock := &MockRunREST{}
		builder := Explore{
			connection: conMock,
		}
		query := builder.build()
		assert.NotEmpty(t, query, "Check that there is no panic if query is not validly build")
	})


}