package domain

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("FormationEntity Builder Tests", Ordered, func() {

	Describe("WithRowsFromJSON", func() {
		var builder *FormationEntityBuilder

		BeforeEach(func() {
			builder = NewFormationEntity()
		})

		It("should handle empty JSON", func() {
			result := builder.WithRowsFromJSON([]byte{}).Build()
			Expect(result.Rows).To(BeEmpty())
		})

		It("should handle nil JSON", func() {
			result := builder.WithRowsFromJSON(nil).Build()
			Expect(result.Rows).To(BeEmpty())
		})

		It("should handle invalid JSON", func() {
			invalidJSON := []byte(`invalid json`)
			result := builder.WithRowsFromJSON(invalidJSON).Build()
			Expect(result.Rows).To(BeEmpty())
		})

		It("should handle empty rows array", func() {
			emptyRowsJSON := []byte(`[]`)
			result := builder.WithRowsFromJSON(emptyRowsJSON).Build()
			Expect(result.Rows).To(BeEmpty())
		})

		It("should parse valid formation rows JSON", func() {
			validJSON := []byte(`[
				{
					"tiles": [
						{"position_x": 0, "position_y": 0, "unit_id": "unit_1"},
						{"position_x": 1, "position_y": 0, "unit_id": "unit_2"}
					]
				},
				{
					"tiles": [
						{"position_x": 0, "position_y": 1, "unit_id": "unit_3"}
					]
				}
			]`)

			result := builder.WithRowsFromJSON(validJSON).Build()

			Expect(result.Rows).To(HaveLen(2))

			firstRow := result.Rows[0]
			Expect(firstRow.Tiles).To(HaveLen(2))
			Expect(firstRow.Tiles[0].PositionX).To(Equal(int32(0)))
			Expect(firstRow.Tiles[0].PositionY).To(Equal(int32(0)))
			Expect(firstRow.Tiles[0].UnitId).To(Equal("unit_1"))
			Expect(firstRow.Tiles[1].PositionX).To(Equal(int32(1)))
			Expect(firstRow.Tiles[1].PositionY).To(Equal(int32(0)))
			Expect(firstRow.Tiles[1].UnitId).To(Equal("unit_2"))

			secondRow := result.Rows[1]
			Expect(secondRow.Tiles).To(HaveLen(1))
			Expect(secondRow.Tiles[0].PositionX).To(Equal(int32(0)))
			Expect(secondRow.Tiles[0].PositionY).To(Equal(int32(1)))
			Expect(secondRow.Tiles[0].UnitId).To(Equal("unit_3"))
		})
	})
})
