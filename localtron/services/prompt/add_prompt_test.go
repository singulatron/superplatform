/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package promptservice_test

import (
	"testing"

	"github.com/singulatron/singulatron/localtron/di"
	"github.com/stretchr/testify/require"
)

func TestAddPrompt(t *testing.T) {
	universe, err := di.BigBang(di.UniverseOptions{
		Test: true,
	})
	require.NoError(t, err)
	require.NotNil(t, universe)
}
