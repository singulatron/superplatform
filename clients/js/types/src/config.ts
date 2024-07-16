export interface Config {
  download?: {
    downloadFolder?: string;
  };

  model?: {
    currentModelId?: string;
  };

  /** This flag drives a minor UX feature:
   * if the user has not installed the runtime we show an INSTALL
   * button, but if the user has already installed the runtime we show
   * we show a START runtime button.
   * */
  isRuntimeInstalled?: boolean;
}

// eslint-disable-next-line
export interface ConfigGetRequest {}

export interface ConfigGetResponse {
  config: Config;
}
