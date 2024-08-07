import { UserSvcApi } from "@singulatron/client";
import { genericTest } from "./generic_example.js";

async function start() {
  let userService = new UserSvcApi();
  let loginResponse = await userService.login({
    request: {
      email: "singulatron",
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
  if (readTokenResponse.user?.email !== "singulatron") {
    process.exit(1);
  }

  genericTest(token!);
}

start();
