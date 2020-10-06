package main

import (
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type QueryData struct {
	Page     uint32 `form:"page"`
	PageSize uint16 `form:"page_size"`
	LandingPageDateRange
}

type LandingPageDateRange struct {
	StartDateString string    `form:"start_date"`
	EndDateString   string    `form:"end_date"`
	StartDate       time.Time `form:"-"`
	EndDate         time.Time `form:"-"`
}

var CSTTimeZone = time.FixedZone("CST", 8*3600)

func main() {

	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	ctx, router := gin.CreateTestContext(resp)
	router.GET("/query", Query)

	var err error
	ctx.Request, err = http.NewRequest(http.MethodGet, "/query?start_date=2017-02-01&end_date=2018-01-03&page=2&page_size=20", nil)
	if err != nil {
		logrus.Fatal(err)
	}

	router.ServeHTTP(resp, ctx.Request)
}

func Query(ctx *gin.Context) {

	data := &QueryData{}
	err := ctx.ShouldBindQuery(data)
	if err != nil {
		logrus.Fatal(err)
	}

	data.LandingPageDateRange.Parse()

	logrus.WithField("queryData", data).Info("parsed query data")
}

func (date *LandingPageDateRange) Parse() {

	startDate, startDateError := time.ParseInLocation("2006-01-02", date.StartDateString, CSTTimeZone)
	endDate, endDateError := time.ParseInLocation("2006-01-02", date.EndDateString, CSTTimeZone)
	now := time.Now()

	if startDateError != nil && endDateError != nil {
		endDate = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, CSTTimeZone).AddDate(0, 0, -1)
		startDate = endDate.AddDate(0, 0, -6)
	} else if endDateError != nil {
		// 虽然这里 Jetbrain 编辑器提示可能有错，但是逻辑上并不会出现。
		endDate = startDate.AddDate(0, 0, 6)
	} else if startDateError != nil {
		startDate = endDate.AddDate(0, 0, -6)
	}

	date.StartDate = startDate
	date.EndDate = endDate
}