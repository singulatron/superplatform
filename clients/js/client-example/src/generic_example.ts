import { Configuration, GenericSvcApi } from "@singulatron/client";

export async function genericTest(apiKey: string) {
  const genericService: GenericSvcApi = new GenericSvcApi(
    new Configuration({
      apiKey: apiKey,
      // basePath: "https://demo-api.singulatron.com",
    })
  );

  await genericService.createObject({
    body: {
      object: {
        table: "uzerz",
        data: {
          fieldA: "valueA",
        },
      },
    },
  });

  await genericService.createObject({
    body: {
      object: {
        table: "uzerz",
        data: {
          fieldA: "valueB",
        },
      },
    },
  });

  let rsp = await genericService.query({
    body: {
      table: "uzerz",
    },
  });

  if (rsp.objects?.length !== 2) {
    throw "expected find length to be 2";
  }
}
