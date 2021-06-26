package shorty

import (
	"errors"
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/rs/xid"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"shorty/internal/app/config"
	"shorty/pkg/primary"
	"strconv"
	"time"
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
		c.Response.Header().Set("Location", c.Request.URL.String())
		return c.WriteWithStatus("No Content", 204)
	}

	link, errCode := CreateLink(URL, ExpiresIn)
	if errCode != 0 && errCode != 302 {
		return c.WriteWithStatus(http.StatusText(errCode), errCode)
	}

	tmpl := config.Templates["found"]
	if tmpl == nil {
		return c.WriteWithStatus("Internal Server Error", 500)
	}

	viewData := primary.FoundViewData{HostURL: "http://" + c.Request.Host, Link: link}
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

	link, errCode := CreateLink(URL, ExpiresIn)
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

func CreateLink(URL string, ExpiresIn string) (link primary.Link, error int) {
	u, err := url.Parse(URL)
	if err != nil {
		return link, http.StatusBadRequest
	}

	UUID := ""
	Found := true
	for index := 0; index < 10; index++ {
		UUID = StringWithCharset(config.Settings.Link.Length, primary.DefaultUUIDCharset)
		err := config.Gorm.Where("uuid = ?", xid.New().String()).Find(&primary.Link{}).Error
		if err == nil || err == gorm.ErrRecordNotFound {
			Found = false
			break
		} else {
			Found = true
		}
	}

	if Found {
		return link, http.StatusInternalServerError
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
	link.FullURL = URL

	duration, err := GetDuration(ExpiresIn)
	if err != nil {
		log.Println("Cannot parse duration from string value", ExpiresIn)
	} else {
		link.ExpiresAt = time.Now().Add(duration)
		link.ExpiresIn = ExpiresIn
	}
	existing := primary.Link{}
	err = config.Gorm.Where("scheme = ? AND host = ? AND path = ? AND query = ?", link.Scheme, link.Host, link.Path, link.Query).First(&existing).Error
	if existing.ID != 0 {
		return existing, http.StatusFound
	}

	err = config.Gorm.Save(&link).Error
	if err != nil {
		return primary.Link{}, http.StatusInternalServerError
	}

	return link, 0
}

func GetDuration(ExpiresIn string) (Duration time.Duration, err error) {
	switch ExpiresIn {
	case primary.Period5min:
		fallthrough
	case primary.Period30min:
		fallthrough
	case primary.Period1hour:
		fallthrough
	case primary.Period1day:
		fallthrough
	case primary.Period1week:
		fallthrough
	case primary.Period1month:
		fallthrough
	case primary.Period1year:
		return ParseDuration(ExpiresIn), nil
	}

	return time.Duration(0), errors.New("Invalid duration period")
}

var durationRegex = regexp.MustCompile(`P([\d\.]+Y)?([\d\.]+M)?([\d\.]+D)?T?([\d\.]+H)?([\d\.]+M)?([\d\.]+?S)?`)

// ParseDuration converts a ISO8601 duration into a time.Duration
func ParseDuration(str string) time.Duration {
	matches := durationRegex.FindStringSubmatch(str)

	years := parseDurationPart(matches[1], time.Hour*24*365)
	months := parseDurationPart(matches[2], time.Hour*24*30)
	days := parseDurationPart(matches[3], time.Hour*24)
	hours := parseDurationPart(matches[4], time.Hour)
	minutes := parseDurationPart(matches[5], time.Second*60)
	seconds := parseDurationPart(matches[6], time.Second)

	return years + months + days + hours + minutes + seconds
}

func parseDurationPart(value string, unit time.Duration) time.Duration {
	if len(value) != 0 {
		if parsed, err := strconv.ParseFloat(value[:len(value)-1], 64); err == nil {
			return time.Duration(float64(unit) * parsed)
		}
	}
	return 0
}

func StringWithCharset(length uint, charset string) string {
	var seededRand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
