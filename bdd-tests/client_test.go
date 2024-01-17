package bdd_tests

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/go-resty/resty/v2"
	. "github.com/onsi/gomega"
	"testing"
)

type apiFeature struct {
	resp *resty.Response
}

func (a *apiFeature) resetResponse(*godog.Scenario) {
	a.resp = nil
}

func (a *apiFeature) iReceiveAResponseWithHTTPStatus(code int) (err error) {
	Expect(a.resp.StatusCode()).To(Equal(code))
	return nil
}

func (a *apiFeature) iRequestThePortFromTheService(portId string) (err error) {
	client := resty.New()
	a.resp, err = client.R().Get(fmt.Sprintf("http://localhost:8080/ports/%v", portId))
	if err != nil {
		return
	}
	defer func() {
		switch t := recover().(type) {
		case string:
			err = fmt.Errorf(t)
		case error:
			err = t
		}
	}()
	return
}

func (a *apiFeature) theInitialDatabaseOfThePortsWasUploadedSuccessfully() (err error) {
	return
}

func (a *apiFeature) theResponseBodyShouldBeEqualTo(contentType string, body string) (err error) {

	var expected, actual interface{}
	Expect(a.resp.Header().Get("Content-Type")).To(Equal(contentType))
	if contentType == "application/json" {
		// re-encode expected response
		if err = json.Unmarshal([]byte(body), &expected); err != nil {
			return
		}

		// re-encode actual response too
		if err = json.Unmarshal(a.resp.Body(), &actual); err != nil {
			return
		}

		Expect(actual).To(Equal(expected))
	} else {
		Expect(string(a.resp.Body())).To(Equal(body))
	}
	return
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	RegisterFailHandler(func(message string, _ ...int) {
		panic(message)
	})
	api := &apiFeature{}
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		api.resetResponse(sc)
		return ctx, nil
	})
	ctx.Step(`^I receive a response with HTTP status (\d+)$`, api.iReceiveAResponseWithHTTPStatus)
	ctx.Step(`^I request the port "([^"]*)" from the service$`, api.iRequestThePortFromTheService)
	ctx.Step(`^the initial database of the ports was uploaded successfully$`, api.theInitialDatabaseOfThePortsWasUploadedSuccessfully)
	ctx.Step(`^the response content should be "([^"]*)" and the body should should match (.*)$`, api.theResponseBodyShouldBeEqualTo)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}