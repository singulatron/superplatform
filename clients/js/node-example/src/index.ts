import { UserService } from "@singulatron/node";

async function start() {
  let userService = new UserService({});
  let loginResponse = await userService.login("singulatron", "changeme");

  const token = loginResponse.token.token;
  userService = new UserService({
    apiKey: token,
  });

  const readTokenResponse = await userService.readUserByToken(token!);
  if (readTokenResponse.user.email !== "singulatron") {
    process.exit(1);
  }
}

start();
