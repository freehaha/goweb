package responders

import (
	codecservices "github.com/stretchrcom/codecs/services"
	"github.com/stretchrcom/goweb/context"
	context_test "github.com/stretchrcom/goweb/webcontext/test"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestAPI_Interface(t *testing.T) {

	assert.Implements(t, (*APIResponder)(nil), new(GowebAPIResponder))

}

func TestNewGowebAPIResponder(t *testing.T) {

	http := new(GowebHTTPResponder)
	codecService := new(codecservices.WebCodecService)
	api := NewGowebAPIResponder(codecService, http)

	assert.Equal(t, http, api.httpResponder)
	assert.Equal(t, codecService, api.GetCodecService())

	assert.Equal(t, api.StandardFieldStatusKey, "s")
	assert.Equal(t, api.StandardFieldDataKey, "d")
	assert.Equal(t, api.StandardFieldErrorsKey, "e")

}

func TestRespond(t *testing.T) {

	http := new(GowebHTTPResponder)
	codecService := new(codecservices.WebCodecService)
	API := NewGowebAPIResponder(codecService, http)
	ctx := context_test.MakeTestContext()
	data := map[string]interface{}{"name": "Mat"}

	API.Respond(ctx, 200, data, nil)

	assert.Equal(t, context_test.TestResponseWriter.Output, "{\"d\":{\"name\":\"Mat\"},\"s\":200}")

}

func TestRespondWithCustomFieldnames(t *testing.T) {

	http := new(GowebHTTPResponder)
	codecService := new(codecservices.WebCodecService)
	API := NewGowebAPIResponder(codecService, http)
	ctx := context_test.MakeTestContext()
	data := map[string]interface{}{"name": "Mat"}

	API.StandardFieldDataKey = "data"
	API.StandardFieldStatusKey = "status"

	API.Respond(ctx, 200, data, nil)

	assert.Equal(t, context_test.TestResponseWriter.Output, "{\"data\":{\"name\":\"Mat\"},\"status\":200}")

}

func TestWriteResponseObject(t *testing.T) {

	http := new(GowebHTTPResponder)
	codecService := new(codecservices.WebCodecService)
	API := NewGowebAPIResponder(codecService, http)
	ctx := context_test.MakeTestContext()
	data := map[string]interface{}{"name": "Mat"}

	API.WriteResponseObject(ctx, 200, data)

	assert.Equal(t, context_test.TestResponseWriter.Output, "{\"name\":\"Mat\"}")

}

func TestAPI_StandardResponseObjectTransformer(t *testing.T) {

	http := new(GowebHTTPResponder)
	codecService := new(codecservices.WebCodecService)
	API := NewGowebAPIResponder(codecService, http)
	ctx := context_test.MakeTestContext()
	data := map[string]interface{}{"name": "Mat"}

	API.SetStandardResponseObjectTransformer(func(ctx context.Context, sro map[string]interface{}) (map[string]interface{}, error) {

		return map[string]interface{}{
			"sro":       sro,
			"something": true,
		}, nil

	})

	API.RespondWithData(ctx, data)

	assert.Equal(t, context_test.TestResponseWriter.Output, "{\"something\":true,\"sro\":{\"d\":{\"name\":\"Mat\"},\"s\":200}}")

}

func TestAPI_RespondWithData(t *testing.T) {

	http := new(GowebHTTPResponder)
	codecService := new(codecservices.WebCodecService)
	API := NewGowebAPIResponder(codecService, http)
	ctx := context_test.MakeTestContext()
	data := map[string]interface{}{"name": "Mat"}

	API.RespondWithData(ctx, data)

	assert.Equal(t, context_test.TestResponseWriter.Output, "{\"d\":{\"name\":\"Mat\"},\"s\":200}")

}

func TestAPI_RespondWithError(t *testing.T) {

	http := new(GowebHTTPResponder)
	codecService := new(codecservices.WebCodecService)
	API := NewGowebAPIResponder(codecService, http)
	ctx := context_test.MakeTestContext()
	errObject := "error message"

	API.RespondWithError(ctx, 500, errObject)

	assert.Equal(t, context_test.TestResponseWriter.Output, "{\"e\":[\"error message\"],\"s\":500}")

}
