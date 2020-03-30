package models_test

import (
	"errors"
	"testing"

	"github.com/ONSdigital/dp-code-list-api/models"
	dbmodels "github.com/ONSdigital/dp-graph/v2/models"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewCode(t *testing.T) {

	Convey("NewCode function called with a nil argument results in an empty API Code model", t, func() {
		code := models.NewCode(nil)
		So(code, ShouldResemble, &models.Code{})
	})

	Convey("Given a valid database Code model", t, func() {
		dbCode := &dbmodels.Code{
			ID:    "testID",
			Label: "testLabel",
			Code:  "testCode",
		}

		Convey("NewCode function returns the corresponding API Code model", func() {
			code := models.NewCode(dbCode)
			So(code, ShouldResemble, &models.Code{
				ID:    "testID",
				Label: "testLabel",
				Code:  "testCode",
			})
		})
	})
}

func TestNewCodeResults(t *testing.T) {

	Convey("NewCodeResults function called with a nil argument results in an empty API CodeResults model", t, func() {
		codeResults := models.NewCodeResults(nil)
		So(codeResults, ShouldResemble, &models.CodeResults{})
	})

	Convey("Given a CodeResults model with nil Items", t, func() {
		dbCodeResults := &dbmodels.CodeResults{Items: nil}

		Convey("NewCodeResults function returns the corresponding API Code model", func() {
			codeResults := models.NewCodeResults(dbCodeResults)
			So(codeResults, ShouldResemble, &models.CodeResults{Items: nil})
		})
	})

	Convey("Given a valid database CodeResults model with items", t, func() {
		dbCodeResults := &dbmodels.CodeResults{
			Items: []dbmodels.Code{
				dbmodels.Code{
					ID:    "testID",
					Label: "testLabel",
					Code:  "testCode",
				},
			},
		}

		Convey("NewCodeResults function returns the corresponding API Code model", func() {
			codeResults := models.NewCodeResults(dbCodeResults)
			So(codeResults, ShouldResemble, &models.CodeResults{
				Items: []models.Code{
					models.Code{
						ID:    "testID",
						Label: "testLabel",
						Code:  "testCode",
					},
				},
			})
		})
	})
}

func TestCodeUpdateLinks(t *testing.T) {

	Convey("Given a Code struct without ID", t, func() {
		code := models.Code{}

		Convey("UpdateLinks fails with the expected error", func() {
			err := code.UpdateLinks("host1", "codelist1", "edition1")
			So(err, ShouldResemble, errors.New("unable to create links - code id not provided"))
		})
	})

	Convey("Given a valid Code struct", t, func() {

		code := models.Code{
			ID: "testID",
		}

		expectedCodeWithLinks := models.Code{
			ID: "testID",
			Links: &models.CodeLinks{
				Self: &models.Link{
					ID:   "testID",
					Href: "host1/code-lists/codelist1/editions/edition1/codes/testID",
				},
				CodeList: &models.Link{
					Href: "host1/code-lists/codelist1",
				},
				Datasets: &models.Link{
					Href: "host1/code-lists/codelist1/editions/edition1/codes/testID/datasets",
				},
			},
		}

		Convey("UpdateLinks generates the expected links", func() {
			err := code.UpdateLinks("host1", "codelist1", "edition1")
			So(err, ShouldBeNil)
			So(code, ShouldResemble, expectedCodeWithLinks)
		})
	})
}
