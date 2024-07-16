export interface DockerInfo {
  hasDocker: boolean;
  dockerDaemonAddress?: string;
  error?: string;
}

// {
//   "info": {
//     "hasDocker": true
//   }
// }
export interface DockerInfoResponse {
  info: DockerInfo;
}
