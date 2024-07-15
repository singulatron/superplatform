export interface CreateUserRequest {
  password: string;
  roleIds: string[];
}

export function sayHello(name: string) {
  console.log("hi user" + name);
}
