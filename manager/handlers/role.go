/*
 *     Copyright 2020 The Dragonfly Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package handlers

import (
	"net/http"

	"d7y.io/dragonfly/v2/manager/types"
	"github.com/gin-gonic/gin"
)

// @Summary Create Role
// @Description Create Role by json config
// @Tags Role
// @Accept json
// @Produce json
// @Param Role body types.CreateRoleRequest true "Role"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /roles [post]
func (h *Handlers) CreateRole(ctx *gin.Context) {
	var json types.CreateRoleRequest
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	if err := h.Service.CreateRole(json); err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}

// @Summary Destroy Role
// @Description Destroy role by json config
// @Tags Role
// @Accept json
// @Produce json
// @Param role path string true "role"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /roles/:role [delete]
func (h *Handlers) DestroyRole(ctx *gin.Context) {
	var params types.RoleParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	if ok, err := h.Service.DestroyRole(params.Role); err != nil {
		ctx.Error(err)
		return
	} else if !ok {
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.Status(http.StatusOK)
}

// @Summary Get Role
// @Description Get Role
// @Tags Role
// @Accept json
// @Produce json
// @Param role path string true "role"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /roles/:role [get]
func (h *Handlers) GetRole(ctx *gin.Context) {
	var params types.RoleParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, h.Service.GetRole(params.Role))
}

// @Summary Get Roles
// @Description Get roles
// @Tags Role
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Failure 500
// @Router /roles [get]
func (h *Handlers) GetRoles(ctx *gin.Context) {
	roles := h.Service.GetRoles()
	ctx.JSON(http.StatusOK, roles)
}

// @Summary Add Permission For Role
// @Description Add Permission by json config
// @Tags Role
// @Accept json
// @Produce json
// @Param Permission body types.AddPermissionForRoleRequest true "Permission"
// @Param role path string true "role"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /roles/:role/permissions [post]
func (h *Handlers) AddPermissionForRole(ctx *gin.Context) {
	var params types.RoleParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	var json types.AddPermissionForRoleRequest
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	if ok, err := h.Service.AddPermissionForRole(params.Role, json); err != nil {
		ctx.Error(err)
		return
	} else if !ok {
		ctx.Status(http.StatusConflict)
		return
	}

	ctx.Status(http.StatusOK)
}

// @Summary Update Role
// @Description Remove Role Permission by json config
// @Tags Role
// @Accept json
// @Produce json
// @Param Permission body types.DeletePermissionForRoleRequest true "Permission"
// @Param role path string true "role"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /roles/:role/permissions [delete]
func (h *Handlers) DeletePermissionForRole(ctx *gin.Context) {
	var params types.RoleParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	var json types.DeletePermissionForRoleRequest
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	if ok, err := h.Service.DeletePermissionForRole(params.Role, json); err != nil {
		ctx.Error(err)
		return
	} else if !ok {
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.Status(http.StatusOK)
}