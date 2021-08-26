package web

import (
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"golang.org/x/net/idna"
	"gorm.io/gorm"
	"log"
	"net/http"
	"net/url"
	"shorty/internal/pkg/shorty"
	"shorty/internal/shorty/config"
	"shorty/pkg/primary"
	"strings"
)

func IndexHandler(c *routing.Context) error {
	tmpl := config.Templates["index"]
	if tmpl == nil {
		return c.WriteWithStatus("Internal Server Error", 500)
	}

	if err := tmpl.Execute(c.Response, nil); err != nil {
		log.Printf("Failed to execute template: %v", err)
		return c.WriteWithStatus("Internal Server Error", 500)
	}

	return nil
}

func CreateHandler(c *routing.Context) error {
	URL := c.Form("url")
	ExpiresIn := c.Form("expiresIn")
	if ExpiresIn == "" {
		ExpiresIn = primary.Period1day
	}

	if len(URL) == 0 {
		tmpl := config.Templates["index"]
		if tmpl == nil {
			return c.WriteWithStatus("Internal Server Error", 500)
		}

		if err := tmpl.Execute(c.Response, nil); err != nil {
			log.Printf("Failed to execute template: %v", err)
			return c.WriteWithStatus("Internal Server Error", 500)
		}
	}

	_, err := url.ParseRequestURI(URL)
	if err != nil {
		tmpl := config.Templates["index"]
		if tmpl == nil {
			return c.WriteWithStatus("Internal Server Error", 500)
		}

		if err := tmpl.Execute(c.Response, nil); err != nil {
			log.Printf("Failed to execute template: %v", err)
			return c.WriteWithStatus("Internal Server Error", 500)
		}
	}

	link, errCode := shorty.CreateLink(URL, ExpiresIn)
	if errCode != 0 && errCode != 302 {
		return c.WriteWithStatus(http.StatusText(errCode), errCode)
	}

	tmpl := config.Templates["found"]
	if tmpl == nil {
		return c.WriteWithStatus("Internal Server Error", 500)
	}

	Host := c.Request.Host
	if strings.HasPrefix(c.Request.Host, "xn--") {
		Host, err = idna.New().ToUnicode(Host)
		if err != nil {
			log.Println("Failed to convert Punycode to Unicode")
			return c.WriteWithStatus("Internal Server Error", 500)
		}
	}

	var Scheme = "https"
	if strings.Contains(Host, "localhost") {
		Scheme = "http"
	}

	viewData := primary.FoundViewData{HostURL: Scheme + "://" + Host, Link: link}
	if err := tmpl.Execute(c.Response, viewData); err != nil {
		log.Printf("Failed to execute template: %v", err)
		return c.WriteWithStatus("Internal Server Error", 500)
	}

	return nil
}

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
	link := primary.Link{}
	err := config.Gorm.Model(&link).Where("uuid = ?", uuid).First(&link).Error
	if err == gorm.ErrRecordNotFound {
		tmpl := config.Templates["not_found"]
		if tmpl == nil {
			return c.WriteWithStatus("Internal Server Error", 500)
		}

		if err := tmpl.Execute(c.Response, nil); err != nil {
			log.Printf("Failed to execute template: %v", err)
			return c.WriteWithStatus("Internal Server Error", 500)
		}
	}

	c.Response.Header().Set("Location", link.FullURL)
	return c.WriteWithStatus("Found", 302)
}
