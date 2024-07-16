export interface DockerInfo {
    hasDocker: boolean;
    dockerDaemonAddress?: string;
    error?: string;
}
export interface DockerInfoResponse {
    info: DockerInfo;
}
