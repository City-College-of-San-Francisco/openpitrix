// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package task

import (
	"google.golang.org/grpc"

	"openpitrix.io/openpitrix/pkg/config"
	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/db"
	"openpitrix.io/openpitrix/pkg/manager"
	"openpitrix.io/openpitrix/pkg/pb"
)

type Server struct {
	db *db.Database
}

func Serve(cfg *config.Config) {
	m := manager.GrpcServer{
		ServiceName: "task",
		Port:        constants.TaskManagerPort,
		MysqlConfig: cfg.Mysql,
	}
	m.Serve(func(server *grpc.Server, db *db.Database) {
		pb.RegisterTaskManagerServer(server, &Server{db})
	})
}
