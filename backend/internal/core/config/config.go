package config

import legacy "safeguild/backend/internal/config"

type Config = legacy.Config

func Load() Config {
	return legacy.Load()
}
