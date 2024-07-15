/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package configservice

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func (cs *ConfigService) GetClientId() (string, error) {
	cs.clientIdFileMutex.Lock()
	defer cs.clientIdFileMutex.Unlock()

	if cs.clientId != "" {
		return cs.clientId, nil
	}

	clientIdFilePath := filepath.Join(cs.ConfigDirectory, "clientId.txt")

	if _, err := os.Stat(clientIdFilePath); os.IsNotExist(err) {
		newUUID, err := uuid.NewRandom()
		if err != nil {
			return "", err
		}

		clientId := newUUID.String()

		err = ioutil.WriteFile(clientIdFilePath, []byte(clientId), 0644)
		if err != nil {
			return "", err
		}

		return clientId, nil
	}

	clientIdBytes, err := ioutil.ReadFile(clientIdFilePath)
	if err != nil {
		return "", err
	}
	cs.clientId = string(clientIdBytes)

	return cs.clientId, nil
}
