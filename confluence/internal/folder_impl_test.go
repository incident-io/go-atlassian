package internal

import (
	"context"
	"errors"
	"net/http"
	"testing"

	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_internalFolderImpl_Get(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx      context.Context
		folderID int
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name: "when the parameters are correct",
			args: args{
				ctx:      context.Background(),
				folderID: 20001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/folders/20001",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.FolderScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the folder id is not provided",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoFolderIDError,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:      context.Background(),
				folderID: 20001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/folders/20001",
					"", nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client
			},

			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the call fails",
			args: args{
				ctx:      context.Background(),
				folderID: 20001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/folders/20001",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.FolderScheme{}).
					Return(&model.ResponseScheme{}, errors.New("error, request failed"))

				fields.c = client
			},

			wantErr: true,
			Err:     errors.New("error, request failed"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewFolderService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Get(testCase.args.ctx, testCase.args.folderID)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}

func Test_internalFolderImpl_Descendants(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx      context.Context
		folderID int
		depth    int
		cursor   string
		limit    int
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name: "when the parameters are correct",
			args: args{
				ctx:      context.Background(),
				folderID: 20001,
				depth:    10,
				cursor:   "cursor-sample",
				limit:    200,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/folders/20001/descendants?cursor=cursor-sample&depth=10&limit=200",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.FolderDescendantChunkScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the depth is not provided it is omitted",
			args: args{
				ctx:      context.Background(),
				folderID: 20001,
				cursor:   "cursor-sample",
				limit:    200,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/folders/20001/descendants?cursor=cursor-sample&limit=200",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.FolderDescendantChunkScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the folder id is not provided",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoFolderIDError,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:      context.Background(),
				folderID: 20001,
				depth:    10,
				cursor:   "cursor-sample",
				limit:    200,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/folders/20001/descendants?cursor=cursor-sample&depth=10&limit=200",
					"", nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client
			},

			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the call fails",
			args: args{
				ctx:      context.Background(),
				folderID: 20001,
				depth:    10,
				cursor:   "cursor-sample",
				limit:    200,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/api/v2/folders/20001/descendants?cursor=cursor-sample&depth=10&limit=200",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.FolderDescendantChunkScheme{}).
					Return(&model.ResponseScheme{}, errors.New("error, request failed"))

				fields.c = client
			},

			wantErr: true,
			Err:     errors.New("error, request failed"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewFolderService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Descendants(testCase.args.ctx, testCase.args.folderID,
				testCase.args.depth, testCase.args.cursor, testCase.args.limit)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}
