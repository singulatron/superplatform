/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 *
 * For commercial use, a separate license must be obtained by purchasing from The Authors.
 * For commercial licensing inquiries, please contact The Authors listed in the AUTHORS file.
 */
package dapper

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"

	dt "github.com/singulatron/singulatron/localtron/dapper/types"
)

func (cm *ConfigurationManager) Run(app *dt.App, tlv map[string]string, anon bool) (*dt.RunContext, error) {
	univ := &dt.RunContext{
		TopLevelVariables: tlv,
		Level:             0,
		Anon:              anon,
		CommandCounter:    0,
	}

	var err error
	for _, feature := range append(app.Features, app.PlatformFeatures[cm.CurrentPlatform]...) {
		if featureID, ok := feature.(string); ok {
			feature, exists := cm.Features[featureID]
			if !exists {
				cm.Printf("Feature \"%s\" not found.\n", featureID)
				continue
			}
			featureInvocation := &dt.FeatureInvocation{
				FeatureID: featureID,
				Args:      []any{},
			}
			if err = cm.recurse(&feature, featureInvocation, univ); err == nil {
				continue
			}
		} else if featureInvocation, ok := feature.(map[string]any); ok {
			for i, v := range featureInvocation["args"].([]any) {
				switch val := v.(type) {
				case string:
					featureInvocation["args"].([]any)[i] = applySubs(tlv, val)
				}

			}

			featureID := featureInvocation["featureId"].(string)
			feature, exists := cm.Features[featureID]
			if !exists {
				cm.Printf("Feature \"%s\" not found.\n", featureID)
				continue
			}
			featureInvocation := &dt.FeatureInvocation{
				FeatureID: featureID,
				Args:      featureInvocation["args"].([]any),
			}

			if err = cm.recurse(&feature, featureInvocation, univ); err == nil {
				continue
			}
		} else {
			log.Fatal("Unkown feature invocation type in app definition", feature)
		}
	}

	return univ, err
}

func (cm *ConfigurationManager) recurse(feature *dt.Feature, featureInvocation *dt.FeatureInvocation, universe *dt.RunContext) error {
	cfp := &dt.FeatureProcessed{
		Feature: feature,
		Level:   universe.Level,
	}
	universe.CommandCounter++

	userFacingFeatureName := feature.Name
	if universe.Anon {
		userFacingFeatureName = fmt.Sprintf("Step %v", universe.CommandCounter)
	}
	universe.FeaturesProcessed = append(universe.FeaturesProcessed, cfp)

	indent := strings.Repeat("  ", universe.Level)
	indentPlusTwo := strings.Repeat("  ", universe.Level+2)

	substitutionMap := map[string]string{}
	subLogStrs := []string{}
	for i := range feature.Arguments {
		if len(featureInvocation.Args) <= i {
			return fmt.Errorf("Feature '%v' has arg specified in position %v (%v), but the invocation does not", feature.Name, i, feature.Arguments[i].Name)
		}
		subLogStrs = append(subLogStrs, fmt.Sprintf("%v=%v", feature.Arguments[i].Name, featureInvocation.Args[i]))
		substitutionMap[feature.Arguments[i].Name] = featureInvocation.Args[i].(string)
	}
	subLog := ""
	if !universe.Anon && len(subLogStrs) > 0 {
		subLog = fmt.Sprintf(" (%v)", strings.Join(subLogStrs, " "))
	}
	cm.Printf("%s- Checking \"%s\"%v: ", indent, userFacingFeatureName, subLog)

	if cm.PreCheckCallback != nil {
		cm.PreCheckCallback(feature, universe)
	}
	result, err := cm.checkFeature(*feature, *featureInvocation)
	if err != nil && strings.Contains(err.Error(), "no check script") {
		return err
	}
	cfp.CheckResult = &result
	cfp.CheckError = err
	if cm.PostCheckCallback != nil {
		cm.PostCheckCallback(feature, universe)
	}
	universe.Level++
	defer func() {
		universe.Level--
	}()

	if err != nil || !result.Success {
		cm.Printf("FAILED. Fixing...\n")

		for _, featureInv := range append(feature.Features, feature.PlatformFeatures[cm.CurrentPlatform]...) {

			if featureID, ok := featureInv.(string); ok {
				nextFeature, exists := cm.Features[featureID]
				if !exists {
					cm.Printf("Feature \"%s\" not found.\n", featureID)
					continue
				}
				nextFeatureInvocation := &dt.FeatureInvocation{
					FeatureID: featureID,
					Args:      []any{},
				}
				if err = cm.recurse(&nextFeature, nextFeatureInvocation, universe); err == nil {
					continue
				}
			} else if nextFeatureInvocationMap, ok := featureInv.(map[string]any); ok {
				featureID := nextFeatureInvocationMap["featureId"].(string)
				nextFeature, exists := cm.Features[featureID]
				if !exists {
					cm.Printf("Feature \"%s\" not found.\n", featureID)
					continue
				}
				nextFeatureInvocation := &dt.FeatureInvocation{
					FeatureID: featureID,
				}
				if v, ok := nextFeatureInvocationMap["args"].([]any); ok {
					nextFeatureInvocation.Args = v
				}

				for i, v := range nextFeatureInvocation.Args {
					switch val := v.(type) {
					case string:

						nextFeatureInvocation.Args[i] = applySubs(substitutionMap, val)
					}

				}

				if err = cm.recurse(&nextFeature, nextFeatureInvocation, universe); err == nil {
					continue
				}
			} else {
				log.Fatal("Unkown feature invocation type ", feature)
			}
		}

		if cm.PreExeccutionCallback != nil {
			cm.PreExeccutionCallback(feature, universe)
		}
		if _, ok := feature.PlatformScripts[cm.CurrentPlatform]; ok &&
			feature.PlatformScripts[cm.CurrentPlatform].Execute != nil &&
			feature.PlatformScripts[cm.CurrentPlatform].Execute.Reboot {
			cm.Printf(" RESTART REQUIRED ")
			universe.RebootRequired = true
		}
		execResult, err := cm.executeFeature(*feature, *featureInvocation)
		cfp.ExecutionResult = &execResult
		cfp.ExecutionError = err
		if cm.PostExecutionCallback != nil {
			cm.PostExecutionCallback(feature, universe)
		}
		if err != nil {
			cm.Printf("%s- Error fixing \"%s\" (EXITING): '%v'.\n", indent, userFacingFeatureName, indentEachLine(indentPlusTwo, execResult.Stdout+execResult.Stdout+err.Error()))
			return err
		}

		if !execResult.Success {
			cm.Printf("%s- Error fixing \"%s\" (EXITING): '%v'.\n", indent, userFacingFeatureName, indentEachLine(indentPlusTwo, execResult.Stdout+execResult.Stderr))
			return errors.New("failed to recover")
		}

		againCheckResult, err := cm.checkFeature(*feature, *featureInvocation)

		if err != nil {
			cm.Printf("%s- Error checking \"%s\" after fixing (EXITING): '%v'.\n", indent, userFacingFeatureName, indentEachLine(indentPlusTwo, againCheckResult.Stdout+againCheckResult.Stdout+err.Error()))
			return err
		}
		if !againCheckResult.Success {
			cm.Printf("%s- Error checking \"%s\" after fixing (EXITING): '%v'.\n", indent, userFacingFeatureName, indentEachLine(indentPlusTwo, againCheckResult.Stdout+againCheckResult.Stderr))
			return errors.New("failed to recover")
		}

		cm.Printf("%s- Fixed \"%s\". OK.\n", indent, userFacingFeatureName)
		return nil
	}

	cm.Printf("OK.\n")

	return nil
}

func (cm ConfigurationManager) substitutions(feature dt.Feature, featureInvocation dt.FeatureInvocation) map[string]string {
	substitutionMap := map[string]string{}
	for i, _ := range feature.Arguments {

		substitutionMap[feature.Arguments[i].Name] = featureInvocation.Args[i].(string)
	}
	return substitutionMap
}

func applySubs(subs map[string]string, script string) string {
	for k, v := range subs {
		// fmt.Println("k", k, "v", v)
		script = strings.Replace(script, fmt.Sprintf("{{.%v}}", k), v, -1)
	}
	return script
}

// CheckFeature runs the check script for a feature and returns the result.
func (cm ConfigurationManager) checkFeature(feature dt.Feature, featureInvocation dt.FeatureInvocation) (dt.FeatureCheckResult, error) {
	if feature.PlatformScripts[cm.CurrentPlatform] == nil ||
		feature.PlatformScripts[cm.CurrentPlatform].Check == nil {
		return dt.FeatureCheckResult{
			ID:        feature.ID,
			CheckedAt: time.Now(),
			Success:   false,
		}, fmt.Errorf("Feature '%v' has no check script for platform '%v'", feature.Name, cm.CurrentPlatform)
	}
	subs := cm.substitutions(feature, featureInvocation)
	success, stdout, stderr, dur, err := executeScript(subs, feature.PlatformScripts[cm.CurrentPlatform].Check)
	return dt.FeatureCheckResult{
		ID:        feature.ID,
		CheckedAt: time.Now(),
		Success:   success,
		Error:     err,
		Stdout:    stdout,
		Stderr:    stderr,
		Duration:  dur,
	}, err
}

// ExecuteFeature runs the execution script for a feature and returns the result.
func (cm ConfigurationManager) executeFeature(feature dt.Feature, featureInvocation dt.FeatureInvocation) (dt.FeatureExecuteResult, error) {
	if feature.PlatformScripts[cm.CurrentPlatform] == nil ||
		feature.PlatformScripts[cm.CurrentPlatform].Execute == nil {
		return dt.FeatureExecuteResult{
			ID:         feature.ID,
			ExecutedAt: time.Now(),
			Success:    false,
		}, fmt.Errorf("Feature '%v' has no execute script for platform '%v'", feature.Name, cm.CurrentPlatform)
	}
	subs := cm.substitutions(feature, featureInvocation)
	success, stdout, stderr, dur, err := executeScript(subs, feature.PlatformScripts[cm.CurrentPlatform].Execute)
	return dt.FeatureExecuteResult{
		ID:         feature.ID,
		ExecutedAt: time.Now(),
		Success:    success,
		Error:      err,
		Stdout:     stdout,
		Stderr:     stderr,
		Duration:   dur,
	}, err
}

// executeScript runs the given script based on its defined runtime environment.
func executeScript(subs map[string]string, script *dt.Script) (bool, string, string, time.Duration, error) {
	source := applySubs(subs, script.Source)
	// fmt.Println("source", source)

	start := time.Now()
	var cmd *exec.Cmd
	switch script.Runtime {
	case "cmd":
		cmd = exec.Command("cmd", "/C", source)
	case "powershell":
		cmd = exec.Command("powershell", "-Command", `$env:WSL_UTF8=1\n`+source)
	case "bash":
		cmd = exec.Command("bash", "-c", source)
	default:
		return false, "", "", 0, fmt.Errorf("unsupported runtime: '%s'", script.Runtime)
	}

	if script.Detach {
		cmd.Start()
		return true, "", "", 0, nil
	}

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	success := err == nil
	return success, stdout.String(), stderr.String(), time.Since(start), err
}
