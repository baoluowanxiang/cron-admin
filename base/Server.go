package base

import "cron-admin/option"

type Server interface {
	Start(option.Options) error
}