export const environment = {
	production: false,
	brandName: 'Singulatron',
	shortBrandName: 'S',
	/** Model list is loaded from a central server.
	 * @todo Consider making this configurable so users
	 * can supply their own servers.
	 */
	backendAddress: 'https://api.commonagi.com',
	// Don't forget to change this when trying to USB debug/accessing this from local network.
	// Find the address of your laptop with `ifconfig` or the respective tool and change:
	// serverAddress: 'http://192.168.176.163:58231'
	serverAddress: 'http://127.0.0.1:58231',
};
