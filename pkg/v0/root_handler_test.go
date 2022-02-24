package v0_test

import (
	"testing"
)

func TestRoot(t *testing.T) {
	// s := httptest.NewServer(api.New())
	// t.Cleanup(s.Close)

	// c := &http.Client{}

	// tests := []struct {
	// 	name       string
	// 	method     string
	// 	header     http.Header
	// 	wantStatus int
	// 	wantHeader http.Header
	// 	wantBody   []byte
	// }{
	// 	{
	// 		name:       "POST not allowed",
	// 		method:     http.MethodPost,
	// 		wantStatus: http.StatusMethodNotAllowed,
	// 	},
	// 	{
	// 		name:       "DELETE not allowed",
	// 		method:     http.MethodDelete,
	// 		wantStatus: http.StatusMethodNotAllowed,
	// 	},
	// 	{
	// 		name:       "PUT not allowed",
	// 		method:     http.MethodPut,
	// 		wantStatus: http.StatusMethodNotAllowed,
	// 	},
	// 	{
	// 		name:       "happy path",
	// 		method:     http.MethodGet,
	// 		wantStatus: http.StatusOK,
	// 		wantHeader: map[string][]string{
	// 			"Content-Type": {"text/plain; charset=utf-8"},
	// 		},
	// 		wantBody: []byte("hello great spirits"),
	// 	},
	// }
	// for _, test := range tests {
	// 	test := test
	// 	t.Run(test.name, func(t *testing.T) {
	// 		req, err := http.NewRequest(test.method, s.URL+"/", nil)
	// 		require.NoError(t, err)

	// 		rsp, err := c.Do(req)
	// 		require.NoError(t, err)
	// 		require.Equal(t, test.wantStatus, rsp.StatusCode)
	// 		if test.wantHeader != nil {
	// 			test.wantHeader.Add("Content-Length", fmt.Sprintf("%d", len(test.wantBody)))
	// 			rsp.Header.Del("Date")
	// 			require.Equal(t, test.wantHeader, rsp.Header)
	// 		}
	// 		if test.wantBody == nil {
	// 			test.wantBody = []byte{}
	// 		}
	// 		require.Equal(t, test.wantBody, readBody(t, rsp.Body))
	// 	})
	// }
}
