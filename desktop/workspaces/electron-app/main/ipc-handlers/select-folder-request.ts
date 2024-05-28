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
import { IpcMainEvent, dialog } from 'electron';
import { WindowApiConst } from 'shared-lib';
import { SelectFolderRequest } from 'shared-lib/models/event-request-response';

export function selectFolderRequest(): (
	event: IpcMainEvent,
	args: SelectFolderRequest
) => void {
	return (event: IpcMainEvent, args: SelectFolderRequest) => {
		dialog
			.showOpenDialog({
				properties: ['openDirectory'],
				buttonLabel: 'Select Folder',
			})
			.then((result) => {
				if (result) {
					let location = result.filePaths[0];
					if (!location) {
						return;
					}
					let data = {
						location: result.filePaths[0],
					};
					// sender will receive this and save the selected folder to localtron
					event?.sender.send(WindowApiConst.ON_FOLDER_SELECT, data);
				}
			});
	};
}
