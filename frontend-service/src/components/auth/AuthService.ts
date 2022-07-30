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

    const res = await QueryMaker.post(payload);
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
      auth: {
        email: credentials.email,
        password: credentials.password,
      },
    });

    const res = await QueryMaker.post(payload);
    const { error, message } = res.data;

    if (error) {
      throw new Error(message);
    }
  }
}
