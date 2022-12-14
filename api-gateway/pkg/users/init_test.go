package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/maslow123/api-gateway/pkg/config"
	"github.com/stretchr/testify/require"
)

func NewServer(t *testing.T) *ServiceClient {
	r := gin.Default()
	config, err := config.LoadConfig("../config/envs", "test")

	require.NoError(t, err)
	svc := RegisterRoutes(r, &config)
	return svc
}

func addAuthorization(
	t *testing.T,
	server *ServiceClient,
	isCompany bool,
) string {
	type LoginResponse struct {
		Status string
		Token  string
	}
	// Login User
	log.Println("=========== LOGIN USER! ===========")
	recorder := httptest.NewRecorder()

	username := "company1@gmail.com"
	if !isCompany {
		username = "user1@gmail.com"
	}

	body := gin.H{
		"username": username,
		"password": "111111",
	}

	data, err := json.Marshal(body)
	require.NoError(t, err)

	url := "/users/login"
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	require.NoError(t, err)

	server.Router.ServeHTTP(recorder, request)

	x, err := ioutil.ReadAll(recorder.Body)
	require.NoError(t, err)

	var resp LoginResponse
	err = json.Unmarshal(x, &resp)
	require.NoError(t, err)

	authorizationHeader := fmt.Sprintf("Bearer %s", resp.Token)

	return authorizationHeader
}
