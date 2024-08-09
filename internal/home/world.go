package home

import (
	"fmt"
	"path/filepath"
	"github.com/AdguardTeam/AdGuardHome/internal/world"
)

func initWorld() (err error){

	worldDir, err := checkWorldDirs(&Context, config)
	if err != nil {
		return err
	}

	worldConfig := world.Config{
		Filename:          filepath.Join(worldDir, "world.db"),
	}



}

// checkWorldDirs checks and returns directory paths to store
// world data ip.
func checkWorldDirs(
	ctx *homeContext,
	conf *configuration,
) (statsDir, querylogDir string, err error) {
	baseDir := ctx.getDataDir()

	statsDir = conf.Stats.DirPath
	if statsDir == "" {
		statsDir = baseDir
	} else {
		err = checkDir(statsDir)
		if err != nil {
			return "", "", fmt.Errorf("world: custom directory: %w", err)
		}
	}

	return statsDir, nil
}