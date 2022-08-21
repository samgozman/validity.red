import { QueryMaker } from "@/services/QueryMaker";

interface AuthCredentials {
  email: string;
  password: string;
}

export class AuthService {
  /**
   * Login user and save auth token to cookie
   * @param credentials
   */
  public static async userLogin(credentials: AuthCredentials): Promise<void> {
    const payload = JSON.stringify({
      action: "UserLogin",
      auth: {
        email: credentials.email,
        password: credentials.password,
      },
    });

    const res = await QueryMaker.post(payload, "/auth/login");
    const { error, message } = res.data;

    if (error) {
      throw new Error(message);
    }
  }

  /**
   * Register user
   * @param credentials
   */
  public static async userRegister(
    // TODO: Use another type with more fields
    credentials: AuthCredentials
  ): Promise<void> {
    const payload = JSON.stringify({
      action: "UserRegister",
      register: {
        email: credentials.email,
        password: credentials.password,
      },
    });

    const res = await QueryMaker.post(payload, "/auth/register");
    const { error, message } = res.data;

    if (error) {
      throw new Error(message);
    }
  }
}
