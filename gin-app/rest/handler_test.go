package rest

import (
	"encoding/json"
	"errors"
	"go-music/gin-app/db"
	"go-music/gin-app/models"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHandler_GetProducts(t *testing.T) {
	gin.SetMode(gin.TestMode)                               // Switch to test mode so you don't get such noisy output.
	mockdbLayer := db.NewMockDBLayerWithData()              // Initialize mock DBLayer with data.
	h := NewHandlerWithDB(mockdbLayer)                      // Initialize handler.
	const productsURL string = "/products"                  // Define the URL for testing.

	type errMSG struct {                                    // Represent the error message.
		Error string `json:"error"`
	}

	tests := []struct {                                     // tests array hold a list of test cases.
		name             string                             //     Test case name.
		inErr            error                              //     Input error.
		outStatusCode    int                                //     Expected HTTP status code.
		expectedRespBody interface{}                        //     Expected HTTP response body.
	}{
		{
			"getproductsnoerrors",                    // Test case 1: happy path (no error)
			nil,
			http.StatusOK,
			mockdbLayer.GetMockProductsData(),
		},
		{
			"getproductswitherror",                          // Test case 2: test error situation
			errors.New("get products error"),           // Inject an error called "get products error" into the mock db layer type
			http.StatusInternalServerError,
			errMSG{Error: "get products error"},  // The input error will be expected to be returned as the HTTP response body
		},
	}
	for _, tt := range tests {                                                 // Loop through each test case.
		t.Run(tt.name, func(t *testing.T) {
			mockdbLayer.SetError(tt.inErr)                                     // Inject the input error.
			req := httptest.NewRequest(http.MethodGet, productsURL, nil) // Create a test request.
			w := httptest.NewRecorder()                                        // Create a http response recorder.
			_, engine := gin.CreateTestContext(w)                              // Create a fresh gin context and gin engine object from the response recorder.

			engine.GET(productsURL, h.GetProducts)                             // Map handler to the URL.
			engine.ServeHTTP(w, req)                                           // Serve the request.

			response := w.Result()                                             // Get the response.

			// Verify HTTP status code.
			if response.StatusCode != tt.outStatusCode {
				t.Errorf("Received Status code %d does not match expected status code %d", response.StatusCode, tt.outStatusCode)
			}

			// Decode HTTP response body.
			var respBody interface{}
			if tt.inErr != nil {
				// If an error was injected, then the response should be decoded to an error message type.
				var errmsg errMSG
				json.NewDecoder(response.Body).Decode(&errmsg)
				respBody = errmsg
			} else {
				// If an error was not injected, the response should be decoded to a slice of products data types.
				products := []models.Product{}
				json.NewDecoder(response.Body).Decode(&products)
				respBody = products
			}

			// Verify HTTP response body.
			if !reflect.DeepEqual(respBody, tt.expectedRespBody) {
				t.Logf("%+v , %+v", respBody, tt.expectedRespBody)
				t.Errorf("Received HTTP response body %+v does not match expected HTTP response Body %+v", respBody, tt.expectedRespBody)
			}
		})
	}
}
