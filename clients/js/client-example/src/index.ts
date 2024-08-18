import { UserSvcApi } from "@singulatron/client";
import { dynamicTest } from "./dynamic_example";

async function start() {
  let userService = new UserSvcApi();
  let loginResponse = await userService.login({
    request: {
      slug: "singulatron",
      password: "changeme",
    },
  });

  const token = loginResponse.token?.token;
  userService = new UserSvcApi();
  ({
    apiKey: token,
  });

  const readTokenResponse = await userService.readUserByToken({
    body: { token: token! },
  });
  if (readTokenResponse.user?.slug !== "singulatron") {
    process.exit(1);
  }

  dynamicTest(token!);
}

start();
