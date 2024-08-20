import { Configuration, DynamicSvcApi } from "@singulatron/client";

export async function dynamicTest(apiKey: string) {
  const dynamicService: DynamicSvcApi = new DynamicSvcApi(
    new Configuration({
      apiKey: apiKey,
    })
  );

  await dynamicService.createObject({
    body: {
      object: {
        table: "uzerz",
        data: {
          fieldA: "valueA",
        },
      },
    },
  });

  await dynamicService.createObject({
    body: {
      object: {
        table: "uzerz",
        data: {
          fieldA: "valueB",
        },
      },
    },
  });

  let rsp = await dynamicService.query({
    body: {
      table: "uzerz",
    },
  });

  if (rsp.objects?.length !== 2) {
    throw "expected result length to be 2";
  }
}
