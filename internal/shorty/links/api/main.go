package api

import (
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"net/http"
	"shorty/internal/pkg/shorty"
	"shorty/pkg/primary"
)

// LinkByIDHandler godoc
// @Security ApiKeyAuth
// @Summary Fetch link by unique identifier
// @Description Fetches link by unique identifier
// @ID get-link-by-id
// @Tags Links
// @Produce json
// @Param id path int true "Link identifier"
// @Success 200 {object} primary.Link
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /links/{id} [get]
func LinkByIDHandler(c *routing.Context) error {
	return nil
}

// LinksListHandler godoc
// @Security ApiKeyAuth
// @Summary Fetch a list of links
// @Description Fetches a list of links
// @ID get-links-list
// @Tags Links
// @Produce json
// @Param count query int false "Links per page"
// @Param page query int false "Page offset"
// @Success 200 {array} primary.Link
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /links [get]
func LinksListHandler(c *routing.Context) error {
	return nil
}

// LinkCreateHandler godoc
// @Security ApiKeyAuth
// @Summary Create a link
// @Description Creates a link
// @ID create-link
// @Tags Links
// @Accept json
// @Produce json
// @Param params body primary.Link true "Create a link request"
// @Success 200 {integer} string "1"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /links [post]
func LinkCreateHandler(c *routing.Context) error {
	URL := c.Form("url")
	ExpiresIn := c.Form("expiresIn")
	if ExpiresIn == "" {
		ExpiresIn = primary.Period1day
	}

	link, errCode := shorty.CreateLink(URL, ExpiresIn)
	if errCode != 0 {
		return c.WriteWithStatus(http.StatusText(errCode), errCode)
	}

	return c.Write(link)
}

// LinkDeleteHandler godoc
// @Security ApiKeyAuth
// @Summary Delete a link by identifier
// @Description Deletes link by identifier
// @ID delete-link
// @Tags Links
// @Accept json
// @Produce json
// @Param id path int true "Link identifier"
// @Success 200 {string} string ""
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /link/{id} [delete]
func LinkDeleteHandler(c *routing.Context) error {
	return nil
}
