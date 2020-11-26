package controlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/raihan/echorest/models"
)

func FetchAllMahasiswa(c echo.Context) error {
	result, err := models.FetchAllMahasiswa()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func StoreMahasiswa(c echo.Context) error {
	Nama := c.FormValue("Nama")
	Jurusan := c.FormValue("Jurusan")
	Konsentrasi := c.FormValue("Konsentrasi")

	result, err := models.StoreMahasiswa(Nama, Jurusan, Konsentrasi)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())

	}
	return c.JSON(http.StatusOK, result)
}
