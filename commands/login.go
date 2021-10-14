package commands

import (
	"bufio"
	"context"
	"errors"
	"os"
	"strings"

	"github.com/ionos-cloud/ionosctl/internal/config"
	"github.com/ionos-cloud/ionosctl/internal/core"
	"github.com/ionos-cloud/ionosctl/internal/printer"
	"github.com/ionos-cloud/ionosctl/services/cloudapi-v5/resources"
	sdk "github.com/ionos-cloud/sdk-go/v5"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

const (
	loginExamples = `ionosctl login --user $IONOS_USERNAME --password $IONOS_PASSWORD
Status: Authentication successful!

ionosctl login --token $IONOS_TOKEN
Status: Authentication successful!

ionosctl login 
Enter your username:
USERNAME
Enter your password:

Status: Authentication successful!`
)

func LoginCmd() *core.Command {
	ctx := context.TODO()
	loginCmd := core.NewCommand(ctx, nil, core.CommandBuilder{
		Namespace: "login",
		Resource:  "login",
		Verb:      "login",
		Aliases:   []string{"log", "auth"},
		ShortDesc: "Authentication command for SDK",
		LongDesc: `Use this command to authenticate. You can use  ` + "`" + `--user` + "`" + ` and ` + "`" + `--password` + "`" + ` flags or you can use  ` + "`" + `--token` + "`" + ` flag to set the credentials.

By default, the user data after running this command will be saved in:

* macOS: ` + "`" + `${HOME}/Library/Application Support/ionosctl/config.json` + "`" + `
* Linux: ` + "`" + `${XDG_CONFIG_HOME}/ionosctl/config.json` + "`" + `
* Windows: ` + "`" + `%APPDATA%\ionosctl\config.json` + "`" + `.

You can use another configuration file for authentication with the ` + "`" + `--config` + "`" + ` global option.

Note: The IONOS Cloud CLI supports also authentication with environment variables: $IONOS_USERNAME, $IONOS_PASSWORD or $IONOS_TOKEN.`,
		Example:    loginExamples,
		PreCmdRun:  PreRunLoginCmd,
		CmdRun:     RunLoginUser,
		InitClient: false,
	})
	loginCmd.AddStringFlag(config.ArgUser, "", "", "Username to authenticate")
	loginCmd.AddStringFlag(config.ArgPassword, config.ArgPasswordShort, "", "Password to authenticate")
	loginCmd.AddStringFlag(config.ArgToken, config.ArgTokenShort, "", "Token to authenticate")

	return loginCmd
}

func PreRunLoginCmd(c *core.PreCommandConfig) error {
	if viper.IsSet(core.GetFlagName(c.NS, config.ArgUser)) && viper.IsSet(core.GetFlagName(c.NS, config.ArgPassword)) && viper.IsSet(core.GetFlagName(c.NS, config.ArgToken)) {
		return errors.New("it is recommended to use either username + password, either token")
	}
	return nil
}

func RunLoginUser(c *core.CommandConfig) error {
	c.Printer.Verbose("Note: The login command will save the credentials in a configuration file after the authentication is successful!")
	c.Printer.Verbose("Note: As an alternative to this, ionosctl offers support for environment variables: $%s, $%s or $%s.",
		sdk.IonosUsernameEnvVar, sdk.IonosPasswordEnvVar, sdk.IonosTokenEnvVar)
	username := viper.GetString(core.GetFlagName(c.NS, config.ArgUser))
	pwd := viper.GetString(core.GetFlagName(c.NS, config.ArgPassword))
	token := viper.GetString(core.GetFlagName(c.NS, config.ArgToken))

	if token != "" {
		// If token is set, use only token
		viper.Set(config.Token, token)
		c.Printer.Verbose("Token is set.")
	} else {
		// If token and username are not set, display messages
		if username == "" {
			err := c.Printer.Print("Enter your username:")
			if err != nil {
				return err
			}
			in := bufio.NewReader(c.Stdin)
			username, err = in.ReadString('\n')
			if err != nil {
				return err
			}
			username = strings.TrimRight(username, "\r\n")
		}
		if pwd == "" {
			err := c.Printer.Print("Enter your password:")
			if err != nil {
				return err
			}
			bytesPwd, err := term.ReadPassword(int(os.Stdin.Fd()))
			if err != nil {
				return err
			}
			pwd = string(bytesPwd)
		}
		viper.Set(config.Username, username)
		c.Printer.Verbose("Username is set %s", viper.GetString(config.Username))
		viper.Set(config.Password, pwd)
		c.Printer.Verbose("Password is set.")
	}
	c.Printer.Verbose("ServerUrl: %s", config.GetServerUrl())
	viper.Set(config.ServerUrl, viper.GetString(config.ArgServerUrl))
	clientSvc, err := resources.NewClientService(
		viper.GetString(config.Username),
		viper.GetString(config.Password),
		viper.GetString(config.Token),
		config.GetServerUrl(),
	)
	if err != nil {
		return err
	}
	// Check the auth is correct
	c.Printer.Verbose("Checking authentication...")
	dcsSvc := resources.NewDataCenterService(clientSvc.Get(), context.TODO())
	_, _, err = dcsSvc.List()
	if err != nil {
		return err
	}

	// Store credentials
	c.Printer.Verbose("Storing credentials to the configuration file: %v", viper.GetString(config.ArgConfig))
	err = config.WriteFile()
	if err != nil {
		return err
	}

	return c.Printer.Print(printer.Result{
		Message: "Authentication successful!",
	})
}
