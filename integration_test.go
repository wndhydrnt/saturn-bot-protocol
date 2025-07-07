package main

import (
	"bytes"
	"context"
	"embed"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/require"
)

var pluginPath = flag.String("plugin-path", "", "Path to the plugin to test")
var saturnBotPath = flag.String("saturn-bot-path", "saturn-bot", "Path to the binary of saturn-bot")

type callResultStderrKey struct{}
type callResultStdoutKey struct{}
type pluginFlagsKey struct{}

func pluginFlags(ctx context.Context) map[string][]string {
	flags, ok := ctx.Value(pluginFlagsKey{}).(map[string][]string)
	if !ok {
		flags = make(map[string][]string)
	}

	return flags
}

func executeSaturnBot(t godog.TestingT, cmdName string, flags map[string][]string) (string, string) {
	flags["path"] = []string{*pluginPath}
	args := []string{"plugin", cmdName, "--log-level", "debug", "--log-format", "json"}
	for flagKey, flagValues := range flags {
		for _, flagValue := range flagValues {
			args = append(args, "--"+flagKey)
			args = append(args, flagValue)
		}
	}

	cmd := exec.Command(*saturnBotPath, args...)
	stdoutBuf := &bytes.Buffer{}
	cmd.Stdout = stdoutBuf
	stderrBuf := &bytes.Buffer{}
	cmd.Stderr = stderrBuf
	err := cmd.Start()
	require.NoError(t, err, "Command starts successfully")
	err = cmd.Wait()
	if err != nil {
		t.Fatalf("Command failed\n  Stderr: %s\n  Call: %s %s", stderrBuf.String(), *saturnBotPath, strings.Join(args, " "))
	}
	return stdoutBuf.String(), stderrBuf.String()
}

func applyIsCalled(ctx context.Context) (context.Context, error) {
	workDir, err := os.MkdirTemp("", "")
	require.NoErrorf(godog.T(ctx), err, "Should create temporary working directory at %s", workDir)
	flags := pluginFlags(ctx)
	flags["workdir"] = []string{workDir}
	resultStdout, resultStderr := executeSaturnBot(godog.T(ctx), "apply", flags)
	ctx = context.WithValue(ctx, callResultStdoutKey{}, resultStdout)
	ctx = context.WithValue(ctx, callResultStderrKey{}, resultStderr)
	return ctx, nil
}

func filterIsCalled(ctx context.Context) (context.Context, error) {
	result, _ := executeSaturnBot(godog.T(ctx), "filter", pluginFlags(ctx))
	return context.WithValue(ctx, callResultStdoutKey{}, result), nil
}

func onPrClosedIsCalled(ctx context.Context) error {
	executeSaturnBot(godog.T(ctx), "onPrClosed", pluginFlags(ctx))
	return nil
}

func onPrCreatedIsCalled(ctx context.Context) error {
	executeSaturnBot(godog.T(ctx), "onPrCreated", pluginFlags(ctx))
	return nil
}

func onPrMergedIsCalled(ctx context.Context) error {
	executeSaturnBot(godog.T(ctx), "onPrMerged", pluginFlags(ctx))
	return nil
}

func theContextContainsTheRepository(ctx context.Context, repoName string) (context.Context, error) {
	flags := pluginFlags(ctx)

	pluginCtx := make(map[string]any)
	if len(flags["context"]) == 1 {
		dec := json.NewDecoder(strings.NewReader(flags["context"][0]))
		err := dec.Decode(&pluginCtx)
		require.NoError(godog.T(ctx), err, "Decode plugin context from JSON")
	}

	repository := map[string]string{
		"full_name":      repoName,
		"clone_url_http": fmt.Sprintf("http://%s.git", repoName),
		"clone_url_ssh":  fmt.Sprintf("git@%s.git", repoName),
		"web_url":        fmt.Sprintf("https://%s", repoName),
	}
	pluginCtx["repository"] = repository
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	err := enc.Encode(pluginCtx)
	require.NoError(godog.T(ctx), err, "Encode plugin context to JSON")
	flags["context"] = []string{buf.String()}
	return context.WithValue(ctx, pluginFlagsKey{}, flags), nil
}

func theContentOfTemporaryFileMatches(ctx context.Context, fileName string, content *godog.DocString) error {
	path := filepath.Join(os.TempDir(), fileName)
	b, err := os.ReadFile(path)
	require.NoErrorf(godog.T(ctx), err, "Read file %s", path)
	require.Equal(godog.T(ctx), content.Content, string(b))
	return nil
}

func theContextContainsRunData(ctx context.Context, runDataRaw *godog.DocString) (context.Context, error) {
	flags := pluginFlags(ctx)

	pluginCtx := make(map[string]any)
	if len(flags["context"]) == 1 {
		dec := json.NewDecoder(strings.NewReader(flags["context"][0]))
		err := dec.Decode(&pluginCtx)
		require.NoError(godog.T(ctx), err, "Decode plugin context from JSON")
	}

	runData := make(map[string]string)
	dec := json.NewDecoder(strings.NewReader(runDataRaw.Content))
	err := dec.Decode(&runData)
	require.NoError(godog.T(ctx), err, "Successfully decodes run data from input")
	pluginCtx["run_data"] = runData
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	err = enc.Encode(pluginCtx)
	require.NoError(godog.T(ctx), err, "Encode plugin context to JSON")
	flags["context"] = []string{buf.String()}
	return context.WithValue(ctx, pluginFlagsKey{}, flags), nil
}

func theFileExistsWithContent(ctx context.Context, fileName string, fileContent *godog.DocString) error {
	flags := ctx.Value(pluginFlagsKey{}).(map[string][]string)
	path := filepath.Join(flags["workdir"][0], fileName)
	b, err := os.ReadFile(path)
	require.NoError(godog.T(ctx), err, "Reads file in repository checkout")
	require.Equal(godog.T(ctx), fileContent.Content, string(b), "Creates file in repository checkout")
	return nil
}

func theMessageIsWrittenToTheLog(ctx context.Context, msg string) error {
	type logMessage struct {
		Msg string `json:"msg"`
	}

	callResult := ctx.Value(callResultStderrKey{}).(string)
	lines := strings.Split(callResult, "\n")
	for _, line := range lines {
		lmsg := logMessage{}
		_ = json.Unmarshal([]byte(line), &lmsg)
		if lmsg.Msg == msg {
			return nil
		}
	}

	return fmt.Errorf("Log message \"%s\" not found. Log output:\n%s", msg, callResult)
}

func theTemporaryFileIsDeleted(ctx context.Context, fileName string) error {
	path := filepath.Join(os.TempDir(), fileName)
	err := os.Remove(path)
	if !errors.Is(err, os.ErrNotExist) {
		require.NoErrorf(godog.T(ctx), err, "Deletes the file %s", path)
	}
	return nil
}

func thePluginConfiguration(ctx context.Context, configurationJSON *godog.DocString) (context.Context, error) {
	flags, ok := ctx.Value(pluginFlagsKey{}).(map[string][]string)
	if !ok {
		flags = make(map[string][]string)
	}

	config := make(map[string]string)
	dec := json.NewDecoder(strings.NewReader(configurationJSON.Content))
	err := dec.Decode(&config)
	require.NoError(godog.T(ctx), err, "Successfully decodes plugin configuration")

	var configArgs []string
	for k, v := range config {
		configArgs = append(configArgs, fmt.Sprintf("%s=%s", k, v))
	}

	flags["config"] = configArgs
	return context.WithValue(ctx, pluginFlagsKey{}, flags), nil
}

func theResponseShouldMatchJSON(ctx context.Context, payload *godog.DocString) error {
	callResult := ctx.Value(callResultStdoutKey{}).(string)
	require.JSONEq(godog.T(ctx), payload.Content, callResult, "Response is equal")
	return nil
}

func shutdownIsCalled(ctx context.Context) (context.Context, error) {
	resultStdout, resultStderr := executeSaturnBot(godog.T(ctx), "shutdown", pluginFlags(ctx))
	ctx = context.WithValue(ctx, callResultStdoutKey{}, resultStdout)
	ctx = context.WithValue(ctx, callResultStderrKey{}, resultStderr)
	return ctx, nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^Apply is called$`, applyIsCalled)
	ctx.Step(`^Filter is called$`, filterIsCalled)
	ctx.Step(`^OnPrClosed is called$`, onPrClosedIsCalled)
	ctx.Step(`^OnPrCreated is called$`, onPrCreatedIsCalled)
	ctx.Step(`^OnPrMerged is called$`, onPrMergedIsCalled)
	ctx.Step(`^the content of temporary file "([^"]*)" matches:$`, theContentOfTemporaryFileMatches)
	ctx.Step(`^the context contains run data:$`, theContextContainsRunData)
	ctx.Step(`^the context contains the repository "([^"]*)"$`, theContextContainsTheRepository)
	ctx.Step(`^the file "([^"]*)" exists with content:$`, theFileExistsWithContent)
	ctx.Step(`^the temporary file "([^"]*)" is deleted$`, theTemporaryFileIsDeleted)
	ctx.Step(`^the plugin configuration:$`, thePluginConfiguration)
	ctx.Step(`^the response should match JSON:$`, theResponseShouldMatchJSON)
	ctx.Step(`^the message "([^"]*)" is written to the log$`, theMessageIsWrittenToTheLog)
	ctx.Step(`^Shutdown is called$`, shutdownIsCalled)
}

//go:embed features/*
var features embed.FS

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			FS:       features,
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
