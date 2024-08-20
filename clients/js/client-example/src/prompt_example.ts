import {
  Configuration,
  PromptSvcApi,
  ModelSvcApi,
  DownloadSvcApi,
} from "@singulatron/client";

const tinyLamaModelId = `huggingface/TheBloke/tinyllama-1.1b-chat-v1.0.Q4_K_S.gguf`;
const tinyLamaAssetURL = `https://huggingface.co/TheBloke/TinyLlama-1.1B-Chat-v1.0-GGUF/resolve/main/tinyllama-1.1b-chat-v1.0.Q4_K_S.gguf?download=true`;

export async function promptTest(apiKey: string) {
  const promptSvc: PromptSvcApi = new PromptSvcApi(
    new Configuration({
      apiKey: apiKey,
    })
  );

  const modelSvc: ModelSvcApi = new ModelSvcApi(
    new Configuration({
      apiKey: apiKey,
    })
  );

  const downloadSvc: DownloadSvcApi = new DownloadSvcApi(
    new Configuration({
      apiKey: apiKey,
    })
  );

  await downloadSvc.download({
    request: {
      url: tinyLamaAssetURL,
    },
  });

  let exists = false;
  while (!exists) {
    const dlResponse = await downloadSvc.getDownload({
      downloadId: tinyLamaAssetURL,
    });

    console.log(dlResponse);
    const exists = dlResponse._exists;

    if (!exists || dlResponse.download?.status !== "completed") {
      await new Promise((resolve) => setTimeout(resolve, 2000));
    } else {
      break;
    }
  }

  console.log(`Making model with ID ${tinyLamaModelId} default`);

  await modelSvc.makeDefault({
    modelId: tinyLamaModelId,
  });

  console.log("Starting the default model");

  await modelSvc.startDefaultModel();

  console.log("Prompting");

  const promptRsp = await Promise.race([
    promptSvc.addPrompt({
      request: {
        sync: true,
        prompt: "Is a cat an animal? Just answer with yes or no please.",
      },
    }),
    timeout(10000),
  ]);

  console.log(promptRsp);
}

const timeout = (ms: number) =>
  new Promise((_, reject) =>
    setTimeout(() => reject(new Error("Request timed out")), ms)
  );
