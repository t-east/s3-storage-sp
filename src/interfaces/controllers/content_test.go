package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"sp/src/core"
	"sp/src/domains/entities"
	"sp/src/usecases/interactor"
	mock_port "sp/src/usecases/port_mock"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
)

func TestContentController_Upload(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		setup    func(mockContentContract *mock_port.MockContentContract, mockContentRepo *mock_port.MockContentRepository, mockContentCrypt *mock_port.MockContentCrypt) *entities.Param
		wantCode int
		wantErr  bool
	}{
		{
			name: "Success",
			setup: func(mockContentContract *mock_port.MockContentContract, mockContentRepo *mock_port.MockContentRepository, mockContentCrypt *mock_port.MockContentCrypt) *entities.Param {
				mockContentCrypt.EXPECT().ContentHashGen(gomock.Any()).AnyTimes().Return(&entities.Content{}, nil)
				mockContentRepo.EXPECT().Create(gomock.Any()).AnyTimes().Return(&entities.Content{}, nil)
				mockContentContract.EXPECT().Set(gomock.Any()).AnyTimes().Return(nil)
				param, _, _ := core.CreateParamMock()
				return param
			},
			wantCode: 201,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/api/chal", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			ctrl := gomock.NewController(t)
			mockContentContract := mock_port.NewMockContentContract(ctrl)
			mockContentCrypt := mock_port.NewMockContentCrypt(ctrl)
			mockContentRepo := mock_port.NewMockContentRepository(ctrl)
			_ = tt.setup(mockContentContract, mockContentRepo, mockContentCrypt)
			cc := interactor.NewContentUseCase(mockContentContract, mockContentRepo, mockContentCrypt)
			ah := NewContentHandler(cc)
			err := ah.Post(c)
			httpErr, ok := err.(*echo.HTTPError)
			if ok && httpErr.Code != tt.wantCode {
				t.Errorf("CreateFAQ() status code = %d, want = %d", httpErr.Code, tt.wantCode)
			}
			if !ok && rec.Code != tt.wantCode {
				t.Errorf("CreateFAQ() status code = %d, want = %d", rec.Code, tt.wantCode)
			}
		})
	}
}
