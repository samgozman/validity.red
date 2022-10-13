import type { IResponse } from "@/services/QueryMaker";
import { QueryMaker } from "@/services/QueryMaker";

interface AuthCredentials {
  email: string;
  password: string;
}

interface AuthResponse extends IResponse {
  data: {
    calendarId: string;
  };
}

export class AuthService {
  /**
   * Login user and save auth token to cookie
   * @param credentials
   */
  public static async userLogin(credentials: AuthCredentials): Promise<void> {
    const payload = JSON.stringify({
      email: credentials.email,
      password: credentials.password,
    });

    const res = await new QueryMaker({
      route: "/auth/login",
      payload,
    }).post<AuthResponse>();
    const { error, message, data } = res.data;

    // Save calendar to local storage
    localStorage.setItem("calendarId", data.calendarId);

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
      email: credentials.email,
      password: credentials.password,
    });

    const res = await new QueryMaker({
      route: "/auth/register",
      payload,
    }).post();
    const { error, message } = res.data;

    if (error) {
      throw new Error(message);
    }
  }
}
