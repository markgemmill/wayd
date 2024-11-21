package services

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/pelletier/go-toml/v2"

	// "github.com/BurntSushi/toml"
	"github.com/markgemmill/appdirs"
	"github.com/markgemmill/pathlib"
	"github.com/wailsapp/wails/v3/pkg/application"
)

const version = "0.1.7"

var positions = map[string]bool{"center": true, "ul": true, "ur": true, "ll": true, "lr": true}

var defaultSettings string = `prompt_cyle = 1800
sync_cycle_to = "HFH"
day_starts = 07:00:00
day_ends = 15:00:00
dock_position = "center"

`

type Settings struct {
	PromptEvery  int            `toml:"prompt_every"`
	DockPosition string         `toml:"dock_position"`
	PromptCycle  int            `toml:"prompt_cycle"`
	SyncCycleTo  string         `toml:"sync_cycle_to"`
	DayStartsAt  toml.LocalTime `toml:"day_starts"`
	DayEndsAt    toml.LocalTime `toml:"day_ends"`
}

type SettingsService struct {
	logger       *slog.Logger
	configDir    pathlib.Path
	dataDir      pathlib.Path
	settingsFile pathlib.Path
	Settings     *Settings
}

func (s *SettingsService) GetSettings() *Settings {
	return s.Settings
}

func (s *SettingsService) SetSettings(settings *Settings) {
	s.Settings.DayEndsAt = settings.DayEndsAt
	s.Settings.DayStartsAt = settings.DayStartsAt
	s.Settings.DockPosition = settings.DockPosition
	s.Settings.PromptCycle = settings.PromptCycle
	s.Settings.SyncCycleTo = settings.SyncCycleTo
	s.Write()
}

func (s *SettingsService) Initialize(appDirs appdirs.AppDirs) error {
	// make sure we have our application user directories
	s.logger.Debug("Initializing wayd settings...")
	// appDirs := appdirs.NewAppDirs("wayd", "")
	// create config directory
	configDir := pathlib.NewPath(appDirs.UserConfigDir(), 0777)
	dataDir := pathlib.NewPath(appDirs.UserDataDir(), 0777)
	if !configDir.Exists() {
		err := configDir.MkDirs()
		if err != nil {
			return err
		}
	}
	s.logger.Debug(fmt.Sprintf(`wayd config director: %s`, dataDir.String()))

	// create data directory
	if !dataDir.Exists() {
		err := dataDir.MkDirs()
		if err != nil {
			return err
		}
	}
	s.logger.Debug(fmt.Sprintf(`wayd data directory: %s`, dataDir.String()))

	s.configDir = configDir
	s.dataDir = dataDir

	return nil
}

func (s *SettingsService) Load() error {

	settingsFile := s.configDir.Join("wayd.toml")
	s.logger.Debug(fmt.Sprintf(`wayd settings file: %s`, settingsFile.String()))

	if !settingsFile.Exists() {
		s.logger.Debug(fmt.Sprintf(`Creating default wayd settings file: %s`, settingsFile.String()))
		settingsFile.Write([]byte(defaultSettings))
	}

	s.settingsFile = settingsFile

	s.logger.Debug("Loadig wayd settings...")

	cfgText, err := settingsFile.Read()
	if err != nil {
		return err
	}

	err = toml.Unmarshal(cfgText, s.Settings)
	if err != nil {
		return err
	}

	return nil
}

func (s *SettingsService) Write() error {

	cfgUpdate, err := toml.Marshal(s.Settings)
	if err != nil {
		return nil
	}
	err = s.settingsFile.Write(cfgUpdate)
	if err != nil {
		return nil
	}
	return nil
}

func (s *SettingsService) DatabasePath() string {
	return s.dataDir.Join("wayd.db").String()
}

func (s *SettingsService) Name() string {
	return "SettingsService"
}

func (s *SettingsService) OnStartup(ctx context.Context, options application.ServiceOptions) error {
	s.logger.Debug(fmt.Sprintf("SettingsService.OnStartup... %s", options.Name))
	return nil
}

func (s *SettingsService) OnShutdown() error {
	s.logger.Debug("SettingsService.OnShutdown...")
	return nil
}

func NewSettings(appDirs appdirs.AppDirs, logger *slog.Logger) (*SettingsService, error) {
	settings := Settings{}
	service := SettingsService{
		logger:   logger,
		Settings: &settings,
	}

	err := service.Initialize(appDirs)

	if err != nil {
		return &service, err
	}

	err = service.Load()

	if err != nil {
		return &service, err
	}

	return &service, nil
}
