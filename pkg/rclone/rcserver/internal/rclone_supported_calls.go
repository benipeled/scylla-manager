// Code generated by go-swagger; DO NOT EDIT.

// Copyright (C) 2017 ScyllaDB

package internal

import (
	"github.com/scylladb/go-set/strset"
)

var RcloneSupportedCalls = strset.New(
	"core/bwlimit",
	"core/gc",
	"core/stats-delete",
	"core/stats-reset",
	"job/info",
	"job/stop",
	"operations/about",
	"operations/cat",
	"operations/check-permissions",
	"operations/copyfile",
	"operations/deletefile",
	"operations/list",
	"operations/movefile",
	"operations/purge",
	"sync/copy",
)
