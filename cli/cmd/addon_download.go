package cmd

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"strings"
	"syscall"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	cqapiauth "github.com/cloudquery/cloudquery-api-go/auth"
	"github.com/cloudquery/cloudquery-api-go/config"
	"github.com/cloudquery/cloudquery/cli/internal/team"
	"github.com/cloudquery/plugin-pb-go/managedplugin"
	"github.com/spf13/cobra"
)

const (
	addonDownloadShort = "Download addon from CloudQuery Hub."
	addonDownloadLong  = `Download addon from CloudQuery Hub.

This downloads an addon from CloudQuery Hub to local disk.
`
	addonDownloadExample = `
# Download an addon to local disk
cloudquery addon download addon-team/addon-type/addon-name@v1.0.0`
)

func newCmdAddonDownload() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "download addon-team/addon-type/addon-name@v1.0.0 [-t directory]",
		Short:   addonDownloadShort,
		Long:    addonDownloadLong,
		Example: addonDownloadExample,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Set up a channel to listen for OS signals for graceful shutdown.
			ctx, cancel := context.WithCancel(cmd.Context())

			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, syscall.SIGTERM)

			go func() {
				<-sigChan
				cancel()
			}()

			return runAddonDownload(ctx, cmd, args)
		},
	}
	cmd.Flags().StringP("target", "t", ".", `Download to specified directory. Use - for stdout`)

	return cmd
}

func runAddonDownload(ctx context.Context, cmd *cobra.Command, args []string) error {
	tc := cqapiauth.NewTokenClient()
	token, err := tc.GetToken()
	if err != nil {
		return fmt.Errorf("failed to get auth token: %w", err)
	}

	addonParts := strings.Split(args[0], "/")
	if len(addonParts) != 3 {
		return fmt.Errorf("invalid addon ref: %s", args[0])
	}
	addonVer := strings.Split(addonParts[2], "@")
	if len(addonVer) != 2 {
		return fmt.Errorf("invalid addon ref %q: no version specified", args[0])
	}
	if !strings.HasPrefix(addonVer[1], "v") {
		return fmt.Errorf("invalid addon ref %q: version must start with 'v'", args[0])
	}

	c, err := cloudquery_api.NewClientWithResponses(managedplugin.APIBaseURL(),
		cloudquery_api.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
			return nil
		}))
	if err != nil {
		return fmt.Errorf("failed to create hub client: %w", err)
	}

	targetDir, err := cmd.Flags().GetString("target")
	if err != nil {
		return err
	}

	currentTeam, err := getTeamForAnyToken(ctx, c, token.Type)
	if err != nil {
		return err
	}

	location, checksum, err := getAddonMetadata(ctx, c, currentTeam, addonParts[0], addonParts[1], addonVer[0], addonVer[1])
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, location, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make download request: %w", err)
	}
	if res.StatusCode > 399 {
		return fmt.Errorf("addon download failed: %d %s", res.StatusCode, location)
	}

	return downloadAddonFromResponse(res, checksum, targetDir)
}

func getTeamForAnyToken(ctx context.Context, c *cloudquery_api.ClientWithResponses, tokenType cqapiauth.TokenType) (string, error) {
	switch tokenType {
	case cqapiauth.BearerToken:
		currentTeam, err := config.GetValue("team")
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return "", fmt.Errorf("failed to get team from config: %w", err)
		}
		if currentTeam == "" {
			return "", fmt.Errorf("team is required. Hint: use `cloudquery switch` to set a team")
		}
		return currentTeam, nil
	case cqapiauth.APIKey:
		teams, err := team.NewClientFromAPI(c).ListAllTeams(ctx)
		if err != nil {
			return "", err
		}
		switch l := len(teams); l {
		case 0:
			return "", errors.New("api key has no assigned team")
		case 1:
			return teams[0], nil
		default:
			return "", fmt.Errorf("api key has more than one team: %s", strings.Join(teams, ", "))
		}
	default:
		return "", fmt.Errorf("unknown token type %v", tokenType)
	}
}

func getAddonMetadata(ctx context.Context, c *cloudquery_api.ClientWithResponses, currentTeam, addonTeam, addonType, addonName, addonVersion string) (location, checksum string, retErr error) {
	aj := "application/json"
	resp, err := c.DownloadAddonAssetByTeamWithResponse(ctx, currentTeam, addonTeam, cloudquery_api.AddonType(addonType), addonName, addonVersion, &cloudquery_api.DownloadAddonAssetByTeamParams{Accept: &aj})
	if err != nil {
		return "", "", fmt.Errorf("failed to get team addon metadata: %w", err)
	}
	if resp.StatusCode() > 299 || resp.JSON200 == nil {
		return "", "", fmt.Errorf("failed to read team addon metadata: %w", errorFromHTTPResponse(resp.HTTPResponse, resp))
	}
	return resp.JSON200.Location, resp.JSON200.Checksum, nil
}

func downloadAddonFromResponse(res *http.Response, expectedChecksum, targetDir string) (retErr error) {
	var (
		fileWriter io.WriteCloser
		size       int64
		err        error
	)

	switch targetDir {
	case "-":
		fileWriter = os.Stdout
	default:
		zipPath := filepath.Join(targetDir, path.Base(res.Request.URL.Path))
		if st, err := os.Stat(zipPath); err == nil {
			if st.IsDir() {
				return fmt.Errorf("file %s already exists: is a directory", zipPath)
			}
			return fmt.Errorf("file %s already exists", zipPath)
		}

		f, err := os.Create(zipPath)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
		fileWriter = f

		defer func() {
			if retErr != nil {
				_ = os.Remove(zipPath)
				return
			}
			fmt.Fprintf(os.Stderr, "Wrote %d bytes to %s\n", size, zipPath)
		}()
	}

	shaWriter := sha256.New()
	w := io.MultiWriter(fileWriter, shaWriter)
	if size, err = io.Copy(w, res.Body); err != nil {
		_ = fileWriter.Close()
		return fmt.Errorf("failed to write: %w", err)
	}
	if err := fileWriter.Close(); err != nil {
		return fmt.Errorf("failed to close: %w", err)
	}
	if err := res.Body.Close(); err != nil {
		return fmt.Errorf("failed to close response body: %w", err)
	}

	writtenChecksum := fmt.Sprintf("%x", shaWriter.Sum(nil))
	if writtenChecksum != expectedChecksum {
		return fmt.Errorf("checksum mismatch: expected %s, got %s", expectedChecksum, writtenChecksum)
	}

	return nil
}
