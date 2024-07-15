
interface DockerInfo {
	hasDocker: boolean;
	dockerDaemonAddress?: string;
	error?: string;
}

// {
//   "info": {
//     "hasDocker": true
//   }
// }
interface DockerInfoResponse {
	info: DockerInfo;
}
