package shorty

import (
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/rs/xid"
	"gorm.io/gorm"
	"net/url"
	"shorty/pkg/shorty"
	"strconv"
)

// LinkByUUIDHandler godoc
// @Security ApiKeyAuth
// @Summary Fetch link by unique identifier
// @Description Fetches link by unique identifier
// @ID get-link-by-uuid
// @Tags Links
// @Param id path string true "Link unique identifier"
// @Router /{id} [get]
func LinkByUUIDHandler(c *routing.Context) error {
	uuid := c.Param("id")
	link := shorty.Link{}
	err := Gorm.Model(&shorty.Link{}).Where("short_url = ?", uuid).First(&link).Error
	if err == gorm.ErrRecordNotFound {
		return c.WriteWithStatus("Not Found", 404)
	}

	c.Response.Header().Set("Location", link.FullURL)
	return c.WriteWithStatus("Found", 302)
}

// LinkByIDHandler godoc
// @Security ApiKeyAuth
// @Summary Fetch link by unique identifier
// @Description Fetches link by unique identifier
// @ID get-link-by-id
// @Tags Links
// @Produce json
// @Param id path int true "Link identifier"
// @Success 200 {object} models.Link
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
// @Success 200 {array} models.Link
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /links [get]
func LinksListHandler(c *routing.Context) error {
	return nil
}

// LinkUpdateHandler godoc
// @Security ApiKeyAuth
// @Summary Update link by identifier
// @Description Updates link by identifier
// @ID update-link-by-id
// @Tags Links
// @Accept json
// @Produce json
// @Param id path int true "Link identifier"
// @Param params body models.Link true "Link data"
// @Success 200 {string} string ""
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Router /links/{id} [put]
func LinkUpdateHandler(c *routing.Context) error {
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
// @Param params body models.Link true "Create a link request"
// @Success 200 {integer} string "1"
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /links [post]
func LinkCreateHandler(c *routing.Context) error {
	URL := c.Form("url")
	u, err := url.Parse(URL)
	if err != nil {
		return c.WriteWithStatus("Bad Request", 400)
	}

	link := shorty.Link{}
	UUID := ""
	Found := true
	for index := 0; index < 10; index++ {
		UUID = xid.New().String()
		err := Gorm.Where("uuid = ?", xid.New().String()).Find(&shorty.Link{}).Error
		if err == nil || err == gorm.ErrRecordNotFound {
			Found = false
			break
		} else {
			Found = true
		}
	}

	if Found {
		return c.WriteWithStatus("Internal Server Error", 500)
	}

	link.UUID = UUID
	link.Host = u.Host

	if len(u.Port()) != 0 {
		link.Port, err = strconv.Atoi(u.Port())
		if err != nil {

		}
	} else {
		link.Port = 80
	}
	link.Scheme = u.Scheme
	link.Path = u.Path
	link.Query = u.Query().Encode()

	existing := shorty.Link{}
	err = Gorm.Debug().Where("scheme = ? AND host = ? AND path = ? AND query = ?", link.Scheme, link.Host, link.Path, link.Query).First(&existing).Error
	if existing.ID != 0 {
		return c.WriteWithStatus(existing, 302)
	}

	err = Gorm.Save(&link).Error
	if err != nil {
		return c.WriteWithStatus("Bad Request", 400)
	}

	c.Response.Header().Set("Content-Type", "application/json")
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
