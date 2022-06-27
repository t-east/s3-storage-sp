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

func TestAuditController_Proof(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		setup    func(mockAuditContract *mock_port.MockAuditContract, mockAuditCrypt *mock_port.MockAuditCrypt, mockContentRepo *mock_port.MockContentRepository) *entities.Param
		wantCode int
		wantErr  bool
	}{
		{
			name: "Success",
			setup: func(mockAuditContract *mock_port.MockAuditContract, mockAuditCrypt *mock_port.MockAuditCrypt, mockContentRepo *mock_port.MockContentRepository) *entities.Param {
				var contents []*entities.Receipt
				contents = append(contents, &entities.Receipt{
					ID:       "",
					Content:  entities.Point{},
					MetaData: []string{},
					HashData: []string{},
				})
				proof := &entities.Proof{
					Myu:       "a",
					Gamma:     "a",
					ContentId: "a",
				}
				mockContentRepo.EXPECT().All().AnyTimes().Return(contents, nil)
				mockAuditCrypt.EXPECT().AuditProofGen(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(proof, nil)
				mockAuditContract.EXPECT().GetContentLog(gomock.Any()).AnyTimes().Return(nil, nil)
				mockAuditContract.EXPECT().GetChallen(gomock.Any()).AnyTimes().Return(nil, nil)
				mockAuditContract.EXPECT().RegisterProof(gomock.Any()).AnyTimes().Return(nil)
				param, _, _ := core.CreateParamMock()
				return param
			},
			wantCode: 200,
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
			mockAuditContract := mock_port.NewMockAuditContract(ctrl)
			mockAuditCrypt := mock_port.NewMockAuditCrypt(ctrl)
			mockContentRepo := mock_port.NewMockContentRepository(ctrl)
			_ = tt.setup(mockAuditContract, mockAuditCrypt, mockContentRepo)
			ac := interactor.NewAuditUseCase(mockAuditContract, mockAuditCrypt, mockContentRepo)
			ah := NewAuditHandler(ac)
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
