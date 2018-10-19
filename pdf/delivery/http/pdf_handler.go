package http

import (
	"context"
	"net/http"

	models "github.com/ExPreman/go-svg2pdf/models"
	usecase "github.com/ExPreman/go-svg2pdf/pdf"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type ResponseError struct {
	Message string `json:"message"`
}
type HttpArticleHandler struct {
	Usecase usecase.PDFUsecase
}

func (a *HttpArticleHandler) GeneratePDF(c echo.Context) error {
	template := c.Param("template")
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := a.Usecase.GeneratePDF(ctx, template)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, "OK")
	// return c.File("OK")
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case models.INTERNAL_SERVER_ERROR:
		return http.StatusInternalServerError
	case models.NOT_FOUND_ERROR:
		return http.StatusNotFound
	case models.CONFLIT_ERROR:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func NewPDFHttpHandler(e *echo.Echo, us usecase.PDFUsecase) {
	handler := &HttpArticleHandler{
		Usecase: us,
	}
	e.GET("/pdf/:template", handler.GeneratePDF)
}
