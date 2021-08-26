package shorty

import (
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"shorty/internal/shorty/config"
	"shorty/pkg/primary"
	"strconv"
	"time"
)

func CreateLink(URL string, ExpiresIn string) (link primary.Link, error int) {
	_, err := url.ParseRequestURI(URL)
	if err != nil {
		return link, http.StatusBadRequest
	}

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
