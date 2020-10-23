// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package api4

import (
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/mattermost/mattermost-server/v5/model"
)

func (api *API) InitImport() {
	api.BaseRoutes.Imports.Handle("", api.ApiSessionRequired(listImports)).Methods("GET")
}

func listImports(c *Context, w http.ResponseWriter, r *http.Request) {
	if !c.IsSystemAdmin() {
		c.SetPermissionError(model.PERMISSION_MANAGE_SYSTEM)
		return
	}

	imports, appErr := c.App.ListDirectory(*c.App.Config().ImportSettings.Directory)
	if appErr != nil {
		c.Err = appErr
		return
	}

	for i := 0; i < len(imports); i++ {
		imports[i] = filepath.Base(imports[i])
	}

	if imports == nil {
		imports = []string{}
	}

	data, err := json.Marshal(imports)
	if err != nil {
		c.Err = model.NewAppError("listImports", "app.import.marshal.app_error", nil, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(data)
}