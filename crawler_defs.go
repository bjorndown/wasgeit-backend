package wasgeit

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var dateTimeRe = regexp.MustCompile(`(\d{1,2}.\d{1,2} \d{4}) - Doors: (\d{2}:\d{2})`)
var timeRe = regexp.MustCompile(`\d{2}:\d{2}`)

var kairoCrawler = HTMLCrawler{
	venue:         GetVenueOrPanic("kairo"),
	EventSelector: "article[id]",
	TimeFormat:    "02.01.200615:04",
	GetDateTimeString: func(eventSelection *goquery.Selection) string {
		rawDateTimeString := eventSelection.Find(".concerts_date").Parent().Text()
		timeString := timeRe.FindString(rawDateTimeString)
		return rawDateTimeString[3:13] + timeString
	},
	TitleSelector: "h1",
	LinkBuilder: func(crawler *HTMLCrawler, eventSelection *goquery.Selection) string {
		return eventSelection.AttrOr("id", crawler.venue.URL)
	}}

var dachstockCrawler = HTMLCrawler{
	venue:         GetVenueOrPanic("dachstock"),
	EventSelector: ".event.event-list",
	TimeFormat:    "2.1 200615:04",
	GetDateTimeString: func(eventSelection *goquery.Selection) string {
		rawDateTimeString := eventSelection.Find(".event-date").Text()
		captures := dateTimeRe.FindStringSubmatch(rawDateTimeString)
		return captures[1] + captures[2]
	},
	TitleSelector: "h3",
	LinkBuilder: func(crawler *HTMLCrawler, eventSelection *goquery.Selection) string {
		return eventSelection.AttrOr("data-url", crawler.venue.URL)
	}}

var turnhalleCrawler = HTMLCrawler{
	venue:         GetVenueOrPanic("turnhalle"),
	EventSelector: ".event",
	TimeFormat:    "02. 01. 0615:04",
	GetDateTimeString: func(eventSelection *goquery.Selection) string {
		rawDateTimeString := eventSelection.Find("h4").Text()
		dateString := rawDateTimeString[4:14]
		matches := timeRe.FindAllStringSubmatch(rawDateTimeString, 2)
		var timeString string
		if len(matches) > 0 && len(matches[0]) == 1 {
			timeString = matches[0][0]
		}
		return dateString + timeString
	},
	TitleSelector: "h2",
	LinkBuilder: func(crawler *HTMLCrawler, eventSelection *goquery.Selection) string {
		return eventSelection.Find("a").AttrOr("href", crawler.venue.URL)
	}}

var brasserieLorraineCrawler = HTMLCrawler{
	venue:         GetVenueOrPanic("brasserie-lorraine"),
	EventSelector: ".type-tribe_events",
	TimeFormat:    "January 2",
	GetDateTimeString: func(eventSelection *goquery.Selection) string {
		rawDateTimeString := eventSelection.Find(".tribe-event-date-start").Text()
		dateString := rawDateTimeString[0:11]
		return strings.TrimSpace(dateString)
	},
	TitleSelector: ".tribe-events-list-event-title",
	LinkBuilder: func(crawler *HTMLCrawler, eventSelection *goquery.Selection) string {
		return eventSelection.Find("h2 > a").AttrOr("href", crawler.venue.URL)
	}}

var kofmehlCrawler = HTMLCrawler{
	venue:         GetVenueOrPanic("kofmehl"),
	EventSelector: ".events__element",
	TimeFormat:    "02.01",
	GetDateTimeString: func(eventSelection *goquery.Selection) string {
		rawDateTimeString := eventSelection.Find("time").Text()
		dateString := rawDateTimeString[3:8]
		return dateString
	},
	TitleSelector: ".events__title",
	LinkBuilder: func(crawler *HTMLCrawler, eventSelection *goquery.Selection) string {
		return eventSelection.Find("a.events__link").AttrOr("href", crawler.venue.URL)
	}}

var kiffCrawler = HTMLCrawler{
	venue:         GetVenueOrPanic("kiff"),
	EventSelector: ".programm-grid a:not(.teaserlink)",
	TimeFormat:    "2 Jan",
	GetDateTimeString: func(eventSelection *goquery.Selection) string {
		return eventSelection.Find(".event-date").Text()[3:]
	},
	TitleSelector: ".event-title-wrapper > h2",
	LinkBuilder: func(crawler *HTMLCrawler, eventSelection *goquery.Selection) string {
		if href, exists := eventSelection.Attr("href"); exists {
			return fmt.Sprintf("%s%s", crawler.venue.URL, href)
		}
		return crawler.venue.URL // TODO set as default in Crawl if this function returns ""
	}}

var coqDorCrawler = HTMLCrawler{
	venue:         GetVenueOrPanic("coq-d-or"),
	EventSelector: "#main table:not(.shows)",
	TimeFormat:    "02.01.0615:04",
	GetDateTimeString: func(eventSelection *goquery.Selection) string {
		rawDateTimeString := eventSelection.Find("td.list_first a").Text()
		dateString := strings.Split(rawDateTimeString, ", ")[1]
		rawTimeString := eventSelection.Find("div.entry").Text()
		timeString := timeRe.FindString(rawTimeString)
		return dateString + timeString
	},
	TitleSelector: "td.list_second h2",
	LinkBuilder: func(crawler *HTMLCrawler, eventSelection *goquery.Selection) string {
		return eventSelection.Find("td.list_second h2 a").AttrOr("href", crawler.venue.URL)
	}}

var iscCrawler = HTMLCrawler{
	venue:         GetVenueOrPanic("isc"),
	EventSelector: ".page_programm a.event_preview",
	TimeFormat:    "02.01.",
	GetDateTimeString: func(eventSelection *goquery.Selection) string {
		return eventSelection.Find(".event_title_date").Text()
	},
	TitleSelector: ".event_title_title",
	LinkBuilder: func(crawler *HTMLCrawler, eventSelection *goquery.Selection) string {
		return eventSelection.AttrOr("href", crawler.venue.URL)
	}}

var mahoganyHallCrawler = HTMLCrawler{
	venue:         GetVenueOrPanic("mahogany-hall"),
	EventSelector: ".view-konzerte .views-row",
	TimeFormat:    "02. January 2006|15.04",
	GetDateTimeString: func(eventSelection *goquery.Selection) string {
		dateTimeString := eventSelection.Find(".concert-tueroeffnung").Text()
		dateTimeString = StripSomeWhiteSpaces(dateTimeString)
		dateTimeString = strings.Split(dateTimeString, ", ")[1]
		dateTimeString = strings.Split(dateTimeString, "Uhr")[0]
		if dateTimeString[1] == '.' {
			return "0" + dateTimeString
		}
		return dateTimeString
	},
	TitleSelector: ".views-field-title h2",
	LinkBuilder: func(crawler *HTMLCrawler, eventSelection *goquery.Selection) string {
		href := eventSelection.Find(".views-field-title h2 a").AttrOr("href", "")
		return fmt.Sprint(crawler.venue.URL, href)
	}}

var heitereFahneCrawler = HTMLCrawler{
	venue:         GetVenueOrPanic("heitere-fahne"),
	EventSelector: ".events .event",
	TimeFormat:    "02.01.200615:04",
	GetDateTimeString: func(eventSelection *goquery.Selection) string {
		rawDateTimeString := eventSelection.Find(".date + .time").Parent().Text()
		rawDateTimeString = StripSomeWhiteSpaces(rawDateTimeString)
		rawDateTimeString = strings.TrimSpace(rawDateTimeString)
		return rawDateTimeString[3:13] + rawDateTimeString[33:38]
	},
	TitleSelector: ".alpha.omega.text .inner h2 a",
	LinkBuilder: func(crawler *HTMLCrawler, eventSelection *goquery.Selection) string {
		href := eventSelection.Find(".alpha.omega.text .inner h2 a").AttrOr("href", "")
		return fmt.Sprint(crawler.venue.URL, href)
	}}

var onoCrawler = HTMLCrawler{
	venue:         GetVenueOrPanic("ono"),
	EventSelector: ".EventItem",
	TimeFormat:    "02.01.0615:04",
	GetDateTimeString: func(eventSelection *goquery.Selection) string {
		rawDateTimeString := eventSelection.Find(".EventInfo.subnav").Text()
		rawDateTimeString = wrp.Replace(rawDateTimeString)
		dateString := rawDateTimeString[3:11]
		timeString := timeRe.FindString(rawDateTimeString)
		return dateString + timeString
	},
	TitleSelector: ".EventTextTitle",
	LinkBuilder: func(crawler *HTMLCrawler, eventSelection *goquery.Selection) string {
		href := eventSelection.Find(".EventImage a").AttrOr("href", "")
		return fmt.Sprint(crawler.venue.URL, href)
	}}

var martaCrawler = HTMLCrawler{
	venue:         GetVenueOrPanic("marta"),
	EventSelector: "table.music tbody tr",
	TimeFormat:    "02.01.200615:04",
	GetDateTimeString: func(eventSelection *goquery.Selection) string {
		dateString := eventSelection.Find("td:nth-child(1)").Text()
		rawTimeString := eventSelection.Find("td:nth-child(4)").Text()
		timeString := timeRe.FindString(rawTimeString)
		return dateString + timeString
	},
	TitleSelector: "td:nth-child(3) p",
	LinkBuilder: func(crawler *HTMLCrawler, eventSelection *goquery.Selection) string {
		href := eventSelection.Find(".EventImage a").AttrOr("href", "")
		return fmt.Sprint(crawler.venue.URL, href)
	}}

var bierhuebeliCrawler = HTMLCrawler{
	venue:         GetVenueOrPanic("bierhuebeli"),
	EventSelector: "ul.bh-event-list.all-events li",
	TimeFormat:    "02.01.06",
	GetDateTimeString: func(eventSelection *goquery.Selection) string {
		rawTimeString := eventSelection.Find(".evendates").Text()
		return rawTimeString[8:16]
	},
	TitleSelector: ".eventlink a",
	LinkBuilder: func(crawler *HTMLCrawler, eventSelection *goquery.Selection) string {
		return eventSelection.Find(".eventlink a").AttrOr("href", crawler.venue.URL)
	}}

var dampfzentraleCrawler = HTMLCrawler{
	venue:         GetVenueOrPanic("dampfzentrale"),
	EventSelector: "article .agenda-container",
	TimeFormat:    "2.1.15:04",
	GetDateTimeString: func(eventSelection *goquery.Selection) string {
		article := eventSelection.Parent().Parent()
		month, _ := article.Attr("data-month")
		day, _ := article.Attr("data-date")
		dateString := fmt.Sprintf("%s.%s.", day, month)
		timeString := strings.TrimSpace(eventSelection.Find(".agenda-details .span1").Text())
		return dateString + timeString
	},
	TitleSelector: "h1.agenda-title",
	LinkBuilder: func(crawler *HTMLCrawler, eventSelection *goquery.Selection) string {
		id := eventSelection.Parent().AttrOr("id", "")
		return fmt.Sprintf("%s#%s", crawler.venue.URL, id)
	}}

var HTMLCrawlers = map[string]HTMLCrawler{
	iscCrawler.venue.ShortName:               iscCrawler,
	kiffCrawler.venue.ShortName:              kiffCrawler,
	kofmehlCrawler.venue.ShortName:           kofmehlCrawler,
	kairoCrawler.venue.ShortName:             kairoCrawler,
	dachstockCrawler.venue.ShortName:         dachstockCrawler,
	coqDorCrawler.venue.ShortName:            coqDorCrawler,
	turnhalleCrawler.venue.ShortName:         turnhalleCrawler,
	brasserieLorraineCrawler.venue.ShortName: brasserieLorraineCrawler,
	mahoganyHallCrawler.venue.ShortName:      mahoganyHallCrawler,
	heitereFahneCrawler.venue.ShortName:      heitereFahneCrawler,
	onoCrawler.venue.ShortName:               onoCrawler,
	martaCrawler.venue.ShortName:             martaCrawler,
	bierhuebeliCrawler.venue.ShortName:       bierhuebeliCrawler,
	dampfzentraleCrawler.venue.ShortName:     dampfzentraleCrawler}

var Crawlers = []Crawler{
	&iscCrawler,
	&kiffCrawler,
	&kofmehlCrawler,
	&kairoCrawler,
	&coqDorCrawler,
	&dachstockCrawler,
	&turnhalleCrawler,
	&brasserieLorraineCrawler,
	&mahoganyHallCrawler,
	&heitereFahneCrawler,
	&onoCrawler,
	&martaCrawler,
	&bierhuebeliCrawler,
	&dampfzentraleCrawler}

// http://wartsaal-kaffee.ch/veranstaltungen/
// https://www.facebook.com/pg/CaffeBarSattler/events/t
// http://www.cafete.ch/
// http://www.schlachthaus.ch/spielplan/index.php
// https://www.effinger.ch/events/
// https://www.facebook.com/pg/loescherbern/events/?ref=page_internal
// https://www.facebook.com/peterflamingobern/
// roessli, sous-le-pont,
