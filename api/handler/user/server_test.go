package user

import (
	"encoding/json"
	"github.com/Seiya-Tagami/Recollect-Service/api/domain/entity"
	mock_user "github.com/Seiya-Tagami/Recollect-Service/api/mock/user"
	"github.com/Seiya-Tagami/Recollect-Service/api/response"
	"github.com/Seiya-Tagami/Recollect-Service/api/usecase/user"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type handlerTest struct {
	h  Handler
	hs *handler
	mi *mock_user.MockInteractor
}

func newHandlerTest(t *testing.T) (*handlerTest, func()) {
	t.Helper()
	ctrl := gomock.NewController(t)
	mi := mock_user.NewMockInteractor(ctrl)
	h := New(mi)
	hs := h.(*handler)

	return &handlerTest{
			h:  h,
			hs: hs,
			mi: mi,
		}, func() {
			ctrl.Finish()
		}
}

func Test_handler_CreateUser(t *testing.T) {
	var (
		userReq = entity.User{
			Sub:    "testSub",
			UserID: "testUserID",
			Email:  "test@example.com",
		}
		user = entity.User{
			Sub:       "testSub",
			UserID:    "testUserID",
			UserName:  "testUserName",
			Email:     "test@example.com",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: time.Time{},
		}
		userResponse = response.UserResponse{
			UserID:    "testUserID",
			UserName:  "testUserName",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}
		ht, finish = newHandlerTest(t)
	)
	defer finish()

	tests := []struct {
		name         string
		expects      func(*handlerTest)
		wantStatus   int
		wantResponse response.UserResponse
	}{
		{
			name: "正常系",
			expects: func(ht *handlerTest) {
				ht.mi.EXPECT().CreateUser(userReq).Return(user, nil)
			},
			wantStatus:   http.StatusOK,
			wantResponse: userResponse,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			ht.hs.CreateUser(c)

			assert.Equal(t, tt.wantStatus, w.Code)

			var gotRes response.UserResponse
			if err := json.Unmarshal(w.Body.Bytes(), &gotRes); err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.wantResponse, gotRes)
		})
	}
}

func Test_handler_CheckEmailDuplication(t *testing.T) {
	type fields struct {
		userInteractor user.Interactor
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handler{
				userInteractor: tt.fields.userInteractor,
			}
			h.CheckEmailDuplication(tt.args.c)
		})
	}
}

func Test_handler_CheckUserIDDuplication(t *testing.T) {
	type fields struct {
		userInteractor user.Interactor
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handler{
				userInteractor: tt.fields.userInteractor,
			}
			h.CheckUserIDDuplication(tt.args.c)
		})
	}
}

func Test_handler_DeleteUser(t *testing.T) {
	type fields struct {
		userInteractor user.Interactor
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handler{
				userInteractor: tt.fields.userInteractor,
			}
			h.DeleteUser(tt.args.c)
		})
	}
}

func Test_handler_UpdateUser(t *testing.T) {
	type fields struct {
		userInteractor user.Interactor
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handler{
				userInteractor: tt.fields.userInteractor,
			}
			h.UpdateUser(tt.args.c)
		})
	}
}
