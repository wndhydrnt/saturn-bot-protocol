package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/require"
	protoV1 "github.com/wndhydrnt/saturn-bot-go/protocol/v1"
	"github.com/wndhydrnt/saturn-bot/pkg/command"
)

var pluginPath = flag.String("path", "", "Path to the plugin to test")

type callResultKey struct{}
type pluginOptsKey struct{}

func applyIsCalled(ctx context.Context) (context.Context, error) {
	opts, ok := ctx.Value(pluginOptsKey{}).(command.ExecPluginOptions)
	if !ok {
		opts = command.ExecPluginOptions{}
	}

	opts.LogLevel = "error"
	opts.Path = *pluginPath

	outBuf := &bytes.Buffer{}
	opts.Out = outBuf

	workDir, err := os.MkdirTemp("", "")
	require.NoErrorf(godog.T(ctx), err, "Should create temporary working directory at %s", workDir)
	opts.WorkDir = workDir

	err = command.ExecPlugin("apply", opts)
	require.NoError(godog.T(ctx), err, "Call to apply succeeds")
	ctx = context.WithValue(ctx, pluginOptsKey{}, opts)
	return context.WithValue(ctx, callResultKey{}, outBuf.String()), nil
}

func filterIsCalled(ctx context.Context) (context.Context, error) {
	opts, ok := ctx.Value(pluginOptsKey{}).(command.ExecPluginOptions)
	if !ok {
		opts = command.ExecPluginOptions{}
	}

	opts.LogLevel = "error"
	outBuf := &bytes.Buffer{}
	opts.Out = outBuf
	opts.Path = *pluginPath
	err := command.ExecPlugin("filter", opts)
	require.NoError(godog.T(ctx), err, "Call to filter succeeds")
	return context.WithValue(ctx, callResultKey{}, outBuf.String()), nil
}

func onPrClosedIsCalled(ctx context.Context) error {
	opts, ok := ctx.Value(pluginOptsKey{}).(command.ExecPluginOptions)
	if !ok {
		opts = command.ExecPluginOptions{}
	}

	opts.LogLevel = "error"
	opts.Out = io.Discard
	opts.Path = *pluginPath
	err := command.ExecPlugin("onPrClosed", opts)
	require.NoError(godog.T(ctx), err, "Call to OnPrClosed succeeds")
	return nil
}

func onPrCreatedIsCalled(ctx context.Context) error {
	opts, ok := ctx.Value(pluginOptsKey{}).(command.ExecPluginOptions)
	if !ok {
		opts = command.ExecPluginOptions{}
	}

	opts.LogLevel = "error"
	opts.Out = io.Discard
	opts.Path = *pluginPath
	err := command.ExecPlugin("onPrCreated", opts)
	require.NoError(godog.T(ctx), err, "Call to OnPrCreated succeeds")
	return nil
}

func onPrMergedIsCalled(ctx context.Context) error {
	opts, ok := ctx.Value(pluginOptsKey{}).(command.ExecPluginOptions)
	if !ok {
		opts = command.ExecPluginOptions{}
	}

	opts.LogLevel = "error"
	opts.Out = io.Discard
	opts.Path = *pluginPath
	err := command.ExecPlugin("onPrMerged", opts)
	require.NoError(godog.T(ctx), err, "Call to OnPrMerged succeeds")
	return nil
}

func theContextContainsTheRepository(ctx context.Context, repoName string) (context.Context, error) {
	opts, ok := ctx.Value(pluginOptsKey{}).(command.ExecPluginOptions)
	if !ok {
		opts = command.ExecPluginOptions{}
	}

	if opts.Context == nil {
		opts.Context = &protoV1.Context{}
	}

	opts.Context.Repository = &protoV1.Repository{
		FullName:     repoName,
		CloneUrlHttp: fmt.Sprintf("http://%s.git", repoName),
		CloneUrlSsh:  fmt.Sprintf("git@%s.git", repoName),
		WebUrl:       fmt.Sprintf("https://%s", repoName),
	}
	return context.WithValue(ctx, pluginOptsKey{}, opts), nil
}

func theContentOfTemporaryFileMatches(ctx context.Context, fileName string, content *godog.DocString) error {
	path := filepath.Join(os.TempDir(), fileName)
	b, err := os.ReadFile(path)
	require.NoErrorf(godog.T(ctx), err, "Read file %s", path)
	require.Equal(godog.T(ctx), content.Content, string(b))
	return nil
}

func theContextContainsRunData(ctx context.Context, runDataRaw *godog.DocString) (context.Context, error) {
	opts, ok := ctx.Value(pluginOptsKey{}).(command.ExecPluginOptions)
	if !ok {
		opts = command.ExecPluginOptions{}
	}

	if opts.Context == nil {
		opts.Context = &protoV1.Context{}
	}

	runData := make(map[string]string)
	dec := json.NewDecoder(strings.NewReader(runDataRaw.Content))
	err := dec.Decode(&runData)
	require.NoError(godog.T(ctx), err, "Successfully decodes run data")
	opts.Context.RunData = runData

	return context.WithValue(ctx, pluginOptsKey{}, opts), nil
}

func theFileExistsWithContent(ctx context.Context, fileName string, fileContent *godog.DocString) error {
	opts := ctx.Value(pluginOptsKey{}).(command.ExecPluginOptions)
	path := filepath.Join(opts.WorkDir, fileName)
	b, err := os.ReadFile(path)
	require.NoError(godog.T(ctx), err, "Reads file in repository checkout")
	require.Equal(godog.T(ctx), fileContent.Content, string(b), "Creates file in repository checkout")
	return nil
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
	opts, ok := ctx.Value(pluginOptsKey{}).(command.ExecPluginOptions)
	if !ok {
		opts = command.ExecPluginOptions{}
	}

	opts.Config = make(map[string]string)
	dec := json.NewDecoder(strings.NewReader(configurationJSON.Content))
	err := dec.Decode(&opts.Config)
	require.NoError(godog.T(ctx), err, "Successfully decodes plugin configuration")
	return context.WithValue(ctx, pluginOptsKey{}, opts), nil
}

func theResponseShouldMatchJSON(ctx context.Context, payload *godog.DocString) error {
	callResult := ctx.Value(callResultKey{}).(string)
	require.JSONEq(godog.T(ctx), payload.Content, callResult, "Response is equal")
	return nil
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
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
