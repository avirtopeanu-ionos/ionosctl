package config

import (
	"github.com/ionos-cloud/ionosctl/pkg/constants"
	"os"
	"path/filepath"
	"testing"

	sdk "github.com/ionos-cloud/sdk-go/v6"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestUsingJustTokenEnvVar(t *testing.T) {
	os.Clearenv()
	viper.Reset()

	viper.SetConfigFile(os.DevNull)
	viper.Set(constants.ArgConfig, os.DevNull)

	assert.NoError(t, os.Setenv(sdk.IonosTokenEnvVar, "tok"))
	assert.NoError(t, Load())
	assert.Equal(t, "tok", viper.GetString(Token))
	assert.Equal(t, "", viper.GetString(Username))
	assert.Equal(t, "", viper.GetString(Password))
	assert.Equal(t, "", viper.GetString(ServerUrl))
}

func TestTokEnvWithUserPassConfigBackup(t *testing.T) {
	// Useful for API routes which don't accept bearer tokens, or custom IonosCTL commands (Image Upload)
	os.Clearenv()
	viper.Reset()

	assert.NoError(t, os.Setenv(sdk.IonosTokenEnvVar, "tok"))
	path := filepath.Join("..", "testdata", "config_user_pass.json") // TODO: These files should be created and then destroyed by the tests
	viper.SetConfigFile(path)
	viper.Set(constants.ArgConfig, path)
	assert.NoError(t, os.Chmod(path, 0600))
	assert.NoError(t, Load())

	assert.Equal(t, "tok", viper.GetString(Token))
	assert.Equal(t, "test@ionos.com", viper.GetString(Username))
	assert.Equal(t, "test", viper.GetString(Password))
	assert.Equal(t, "", viper.GetString(ServerUrl))
}

func TestTokEnvWithFullConfig(t *testing.T) {
	// Config token should not override env var token
	os.Clearenv()
	viper.Reset()

	assert.NoError(t, os.Setenv(sdk.IonosTokenEnvVar, "tok"))
	path := filepath.Join("..", "testdata", "config.json") // TODO: These files should be created and then destroyed by the tests
	viper.SetConfigFile(path)
	viper.Set(constants.ArgConfig, path)
	assert.NoError(t, os.Chmod(path, 0600))
	assert.NoError(t, Load())

	assert.Equal(t, "tok", viper.GetString(Token))
	assert.Equal(t, "test@ionos.com", viper.GetString(Username))
	assert.Equal(t, "test", viper.GetString(Password))
	assert.Equal(t, "https://api.ionos.com", viper.GetString(ServerUrl))
}

func TestEnvVarsHavePriority(t *testing.T) {
	// Make sure env vars not overriden by config file
	os.Clearenv()
	viper.Reset()

	assert.NoError(t, os.Setenv(sdk.IonosTokenEnvVar, "env_tok"))
	assert.NoError(t, os.Setenv(sdk.IonosUsernameEnvVar, "env_user"))
	assert.NoError(t, os.Setenv(sdk.IonosPasswordEnvVar, "env_pass"))
	assert.NoError(t, os.Setenv(sdk.IonosApiUrlEnvVar, "env_url"))
	path := filepath.Join("..", "testdata", "config.json") // TODO: These files should be created and then destroyed by the tests
	viper.SetConfigFile(path)
	viper.Set(constants.ArgConfig, path)
	assert.NoError(t, os.Chmod(path, 0600))
	assert.NoError(t, Load())

	assert.Equal(t, "env_tok", viper.GetString(Token))
	assert.Equal(t, "env_user", viper.GetString(Username))
	assert.Equal(t, "env_pass", viper.GetString(Password))
	assert.Equal(t, "env_url", viper.GetString(ServerUrl))
}

func TestAuthErr(t *testing.T) {
	os.Clearenv()
	viper.Reset()

	viper.SetConfigFile(os.DevNull)
	viper.Set(constants.ArgConfig, os.DevNull)

	assert.NoError(t, os.Setenv(sdk.IonosUsernameEnvVar, "env_user"))
	assert.NoError(t, os.Setenv(sdk.IonosApiUrlEnvVar, "env_url"))

	assert.Error(t, Load()) // Need password or token

	assert.Equal(t, "env_user", viper.GetString(Username))
	assert.Equal(t, "env_url", viper.GetString(ServerUrl))
}

func TestUsingJustUsernameAndPasswordEnvVar(t *testing.T) {
	os.Clearenv()
	viper.Reset()

	viper.SetConfigFile(os.DevNull)
	viper.Set(constants.ArgConfig, os.DevNull)

	assert.NoError(t, os.Setenv(sdk.IonosUsernameEnvVar, "user"))
	assert.NoError(t, os.Setenv(sdk.IonosPasswordEnvVar, "pass"))
	assert.NoError(t, Load())
	assert.Equal(t, "", viper.GetString(Token))
	assert.Equal(t, "user", viper.GetString(Username))
	assert.Equal(t, "pass", viper.GetString(Password))
	assert.Equal(t, "", viper.GetString(ServerUrl))
}

func TestBadConfigPerms(t *testing.T) {
	os.Clearenv()
	viper.Reset()

	path := filepath.Join("..", "testdata", "config.json") // TODO: These files should be created and then destroyed by the tests
	viper.SetConfigFile(path)
	viper.Set(constants.ArgConfig, path)
	assert.NoError(t, os.Chmod(path, 0000)) // no read perms
	assert.Error(t, Load())

	assert.Equal(t, "", viper.GetString(Token))
	assert.Equal(t, "", viper.GetString(Username))
	assert.Equal(t, "", viper.GetString(Password))
	assert.Equal(t, "", viper.GetString(ServerUrl))
}

func TestUsingJustTokenConfig(t *testing.T) {
	os.Clearenv()
	viper.Reset()

	path := filepath.Join("..", "testdata", "config_just_token.json") // TODO: These files should be created and then destroyed by the tests
	viper.SetConfigFile(path)
	viper.Set(constants.ArgConfig, path)
	assert.NoError(t, os.Chmod(path, 0600))
	assert.NoError(t, Load())

	assert.Equal(t, "tok", viper.GetString(Token))
	assert.Equal(t, "", viper.GetString(Username))
	assert.Equal(t, "", viper.GetString(Password))
	assert.Equal(t, "", viper.GetString(ServerUrl))
}

func TestUsingJustUsernameAndPasswordConfig(t *testing.T) {
	os.Clearenv()
	viper.Reset()

	path := filepath.Join("..", "testdata", "config_user_pass.json") // TODO: These files should be created and then destroyed by the tests
	viper.SetConfigFile(path)
	viper.Set(constants.ArgConfig, path)
	assert.NoError(t, os.Chmod(path, 0600))
	assert.NoError(t, Load())

	assert.Equal(t, "", viper.GetString(Token))
	assert.Equal(t, "test@ionos.com", viper.GetString(Username))
	assert.Equal(t, "test", viper.GetString(Password))
	assert.Equal(t, "", viper.GetString(ServerUrl))
}

func TestGetServerUrl(t *testing.T) {
	os.Clearenv()
	viper.Reset()

	// use env
	assert.NoError(t, os.Setenv(sdk.IonosApiUrlEnvVar, "url"))
	err := Load()
	assert.Error(t, err) // Error because neither token nor user & pass set
	assert.Equal(t, "url", viper.GetString(ServerUrl))

	// from config
	os.Clearenv()
	viper.Reset()
	viper.SetConfigFile(filepath.Join("..", "testdata", "config.json")) // TODO: These files should be created and then destroyed by the tests
	viper.Set(constants.ArgConfig, filepath.Join("..", "testdata", "config.json"))
	assert.NoError(t, os.Chmod(filepath.Join("..", "testdata", "config.json"), 0600))
	assert.NoError(t, Load())
	assert.Equal(t, "https://api.ionos.com", GetServerUrl())

	viper.Reset()
	fs := pflag.NewFlagSet(constants.ArgServerUrl, pflag.ContinueOnError)
	_ = fs.String(constants.ArgServerUrl, "default", "test flag")
	viper.BindPFlags(fs)
	assert.Equal(t, "default", GetServerUrl())

	assert.NoError(t, fs.Parse([]string{"--" + constants.ArgServerUrl, "explicit"}))
	assert.Equal(t, "explicit", GetServerUrl())
}

func TestLoadFile(t *testing.T) {
	os.Clearenv()
	viper.Reset()

	viper.SetConfigFile(filepath.Join("..", "testdata", "config.json")) // TODO: These files should be created and then destroyed by the tests
	viper.Set(constants.ArgConfig, filepath.Join("..", "testdata", "config.json"))
	assert.NoError(t, os.Chmod(filepath.Join("..", "testdata", "config.json"), 0600))
	assert.NoError(t, LoadFile())
	assert.Equal(t, "test@ionos.com", viper.GetString(Username))
	assert.Equal(t, "test", viper.GetString(Password))
	assert.Equal(t, "jwt-token", viper.GetString(Token))
	assert.Equal(t, "https://api.ionos.com", viper.GetString(ServerUrl))
}

func TestLoadEnvFallback(t *testing.T) {
	os.Clearenv()
	viper.Reset()

	viper.SetConfigFile(filepath.Join("..", "testdata", "config.json")) // TODO: These files should be created and then destroyed by the tests
	viper.Set(constants.ArgConfig, filepath.Join("..", "testdata", "config.json"))
	assert.NoError(t, os.Setenv(sdk.IonosUsernameEnvVar, "user"))
	assert.NoError(t, os.Setenv(sdk.IonosPasswordEnvVar, "pass"))
	assert.NoError(t, os.Setenv(sdk.IonosTokenEnvVar, "token"))
	assert.NoError(t, os.Setenv(sdk.IonosApiUrlEnvVar, "url"))
	assert.NoError(t, Load())
	assert.Equal(t, "user", viper.GetString(Username))
	assert.Equal(t, "pass", viper.GetString(Password))
	assert.Equal(t, "token", viper.GetString(Token))
	assert.Equal(t, "url", viper.GetString(ServerUrl))
}

func TestGetClient(t *testing.T) {
	type args struct {
		name    string
		pwd     string
		token   string
		hostUrl string
	}
	tests := []struct {
		name    string
		runs    int
		args    args
		want    *Client
		wantErr bool
	}{
		{"MissingCredentials", 1, args{"", "", "", ""}, nil, true},
		{"MultipleGetClients", 4, args{"user", "pass", "token", "url"}, instance, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Clearenv()
			viper.Reset()

			assert.NoError(t, os.Setenv(sdk.IonosUsernameEnvVar, tt.args.name))
			assert.NoError(t, os.Setenv(sdk.IonosPasswordEnvVar, tt.args.pwd))
			assert.NoError(t, os.Setenv(sdk.IonosTokenEnvVar, tt.args.token))
			assert.NoError(t, os.Setenv(sdk.IonosApiUrlEnvVar, tt.args.hostUrl))

			for i := 0; i < tt.runs; i++ {
				client, err := GetClient()
				if !tt.wantErr && err != nil {
					t.Errorf("Did not expect error: %v", err)
				}
				assert.Equalf(t, tt.want, client, "newClient(%v, %v, %v, %v)", tt.args.name, tt.args.pwd, tt.args.token, tt.args.hostUrl)
			}
		})
	}
}
