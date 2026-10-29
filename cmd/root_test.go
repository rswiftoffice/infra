package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestRootVersionFlag(t *testing.T) {
	origVersion, origCommit, origDate := version, commit, date
	origCmdVersion := rootCmd.Version
	t.Cleanup(func() {
		version, commit, date = origVersion, origCommit, origDate
		rootCmd.Version = origCmdVersion
		rootCmd.SetArgs(nil)
		rootCmd.SetOut(nil)
	})

	version, commit, date = "9.9.9", "abc1234", "2026-01-02T03:04:05Z"
	rootCmd.Version = "9.9.9 (commit abc1234, built 2026-01-02T03:04:05Z)"

	var out bytes.Buffer
	rootCmd.SetOut(&out)
	rootCmd.SetArgs([]string{"--version"})

	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("--version returned error: %v", err)
	}

	got := out.String()
	for _, want := range []string{"9.9.9", "abc1234", "2026-01-02T03:04:05Z"} {
		if !strings.Contains(got, want) {
			t.Errorf("--version output %q missing %q", got, want)
		}
	}
}

func TestVersionDefaultsIdentifyUnstampedBuild(t *testing.T) {
	if version != "dev" {
		t.Skipf("version stamped by ldflags: %q", version)
	}
	if !strings.HasPrefix(rootCmd.Version, "dev ") {
		t.Errorf("unstamped build should report dev version, got %q", rootCmd.Version)
	}
}
