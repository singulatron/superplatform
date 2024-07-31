export interface DownloadDetails {
    id: string;
    url: string;
    fileName: string;
    dir?: string;
    progress?: number;
    downloadedBytes: number;
    fullFileSize?: number;
    status: "inProgress" | "completed" | "paused" | "cancelled" | "failed";
    filePath?: string;
    paused?: boolean;
    cancelled?: boolean;
    error?: string;
}
export type ListResponse = {
    downloads: DownloadDetails[];
};
export interface DownloadStatusChangeEvent {
    allDownloads: DownloadDetails[];
}
export type GetRequest = {
    download: DownloadDetails;
};
export type GetResponse = {
    download: DownloadDetails;
};
