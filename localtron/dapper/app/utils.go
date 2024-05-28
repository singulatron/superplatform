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
	"fmt"
	"strings"

	dt "github.com/singulatron/singulatron/localtron/dapper/types"
)

func (cm ConfigurationManager) flatDefinitionsOfInvokedFunctions(featureInvocations []any) ([]dt.Feature, error) {
	ret := []dt.Feature{}

	for _, feature := range featureInvocations {
		var (
			feat   dt.Feature
			exists bool
		)
		if featureID, ok := feature.(string); ok {
			feat, exists = cm.Features[featureID]
			if !exists {
				return nil, fmt.Errorf("feature \"%s\" not found", featureID)
			}
		} else if featureInvocation, ok := feature.(dt.FeatureInvocation); ok {
			feat, exists = cm.Features[featureInvocation.FeatureID]
			if !exists {
				return nil, fmt.Errorf("feature \"%s\" not found", featureInvocation.FeatureID)
			}
		}

		ret = append(ret, feat)
		feats, err := cm.flatDefinitionsOfInvokedFunctions(feat.Features)
		if err != nil {
			return nil, err
		}
		ret = append(ret, feats...)
		feats, err = cm.flatDefinitionsOfInvokedFunctions(feat.PlatformFeatures[cm.CurrentPlatform])
		if err != nil {
			return nil, err
		}
		ret = append(ret, feats...)
	}

	return ret, nil
}

func indentEachLine(indentString string, str string) string {
	ss := strings.FieldsFunc(str, func(c rune) bool {
		return c == '\n' || c == '\r'
	})
	for i := range ss {
		// skip first line
		if i == 0 {
			continue
		}
		ss[i] = indentString + ss[i]
	}
	return strings.Join(ss, "\n")
}

func indentEachLineExceptFirst(indentString string, str string) string {
	ss := strings.FieldsFunc(str, func(c rune) bool {
		return c == '\n' || c == '\r'
	})
	for i := range ss {
		// skip first line
		if i == 0 {
			continue
		}
		ss[i] = indentString + ss[i]
	}
	return strings.Join(ss, "\n")
}
