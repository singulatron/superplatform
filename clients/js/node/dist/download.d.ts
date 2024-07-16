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
export type DownloadsResponse = {
    downloads: DownloadDetails[];
};
export interface DownloadStatusChangeEvent {
    allDownloads: DownloadDetails[];
}
