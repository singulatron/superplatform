interface DockerInfo {
    hasDocker: boolean;
    dockerDaemonAddress?: string;
    error?: string;
}
interface DockerInfoResponse {
    info: DockerInfo;
}
