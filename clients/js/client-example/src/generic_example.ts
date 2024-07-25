import { GenericService } from "@singulatron/client";

export async function genericTest(apiKey: string) {
  const genericService: GenericService = new GenericService({
    apiKey: apiKey,
  });

  await genericService.create({
    object: {
      table: "uzerz",
      data: {
        fieldA: "valueA",
      },
    },
  });

  await genericService.create({
    object: {
      table: "uzerz",
      data: {
        fieldA: "valueB",
      },
    },
  });

  let rsp = await genericService.find({});

  if (rsp.objects.length !== 2) {
    throw "expected find length to be 2";
  }
}
