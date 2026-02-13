package config

import legacy "oris/backend/internal/config"

type Config = legacy.Config

func Load() Config {
	return legacy.Load()
}
