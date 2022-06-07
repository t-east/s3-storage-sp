package http

// import (
// 	"bytes"
// 	"io"
// 	"mime/multipart"
// 	"net/http"
// 	"net/http/httptest"
// 	"sp/src/asserts"
// 	"sp/src/core"
// 	"sp/src/domains/entities"
// 	"sp/src/drivers/router"
// 	"sp/src/interfaces/controllers"
// 	mock_port "sp/src/mocks"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// )

// func check(t *testing.T, err error) {
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestContentHandler_UploadContent(t *testing.T) { // nolint:gocognit
// 	t.Parallel()

// 	tests := []struct {
// 		name     string
// 		setup    func(mc mock_port.MockContentContract, ms mock_port.MockContentStorage, mr mock_port.MockContentRepository) *entities.Receipt
// 		body     func() (*bytes.Buffer, string, error)
// 		wantCode int
// 		wantErr  bool
// 	}{
// 		{
// 			name: "Success",
// 			setup: func(mcc mock_port.MockContentContract, mcs mock_port.MockContentStorage, mcr mock_port.MockContentRepository) *entities.Receipt {
// 				receipt := &entities.Receipt{
// 					ID:       "",
// 					Content:  entities.SampleData{},
// 					MetaData: [][]byte{},
// 					HashData: [][]byte{},
// 				}
// 				mcs.EXPECT().Create(gomock.Any()).AnyTimes().Return(nil, nil)
// 				mcr.EXPECT().Create(gomock.Any()).AnyTimes().Return(receipt, nil)
// 				mcc.EXPECT().Register(gomock.Any()).AnyTimes().Return(nil)
// 				return receipt
// 			},
// 			body: func() (*bytes.Buffer, string, error) {

// 				// testByte := []byte{233}

// 				body := &bytes.Buffer{}
// 				writer := multipart.NewWriter(body)

// 				f, err := core.UseFileRead("./linux_logo.jpg")
// 				check(t, err)
// 				part, err := writer.CreateFormFile("content", "./linux_logo.jpg")
// 				check(t, err)
// 				if _, err := io.Copy(part, f); err != nil {
// 					t.Fatal(err)
// 				}

// 				err = writer.WriteField("user_id", "ssssssss")
// 				if err != nil {
// 					t.Fatal(err)
// 				}

// 				// m1, err := os.Create("/tmp/dat2")
// 				// check(t, err)
// 				// _, err = m1.Write(testByte)
// 				// check(t, err)
// 				// meta1, err := writer.CreateFormFile("meta_1", "meta1")
// 				// check(t, err)
// 				// if _, err := io.Copy(meta1, m1); err != nil {
// 				// 	t.Fatal(err)
// 				// }

// 				// m2, err := os.Create("/tmp/dat2")
// 				// check(t, err)
// 				// _, err = m1.Write(testByte)
// 				// check(t, err)
// 				// meta2, err := writer.CreateFormFile("meta_2", "meta2")
// 				// check(t, err)
// 				// if _, err := io.Copy(meta2, m2); err != nil {
// 				// 	t.Fatal(err)
// 				// }

// 				// m3, err := os.Create("/tmp/dat2")
// 				// check(t, err)
// 				// _, err = m3.Write(testByte)
// 				// check(t, err)
// 				// meta3, err := writer.CreateFormFile("meta_3", "meta3")
// 				// check(t, err)
// 				// if _, err := io.Copy(meta3, m3); err != nil {
// 				// 	t.Fatal(err)
// 				// }

// 				// m4, err := os.Create("/tmp/dat2")
// 				// check(t, err)
// 				// _, err = m4.Write(testByte)
// 				// check(t, err)
// 				// meta4, err := writer.CreateFormFile("meta_4", "meta4")
// 				// check(t, err)
// 				// if _, err := io.Copy(meta4, m4); err != nil {
// 				// 	t.Fatal(err)
// 				// }

// 				// m5, err := os.Create("/tmp/dat2")
// 				// check(t, err)
// 				// _, err = m5.Write(testByte)
// 				// check(t, err)
// 				// meta5, err := writer.CreateFormFile("meta_5", "meta5")
// 				// check(t, err)
// 				// if _, err := io.Copy(meta5, m5); err != nil {
// 				// 	t.Fatal(err)
// 				// }

// 				err = writer.Close()
// 				check(t, err)
// 				header := writer.FormDataContentType()
// 				return body, header, nil
// 			},
// 			wantCode: 201,
// 			wantErr:  false,
// 		},
// 	}

// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
// 			body, header, err := tt.body()
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			db, err := router.LoadTestDB()
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			req := httptest.NewRequest(http.MethodPost, "/api/content", body)
// 			req.Header.Set("Content-Type", header)

// 			rec := httptest.NewRecorder()
// 			cc := controllers.LoadContentController(db)
// 			cc.Post(rec, req)
// 			asserts.AssertEqual(t, http.StatusCreated, rec.Code, rec.Body.String()+rec.Result().Status)
// 		})
// 	}
// }
