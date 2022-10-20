import type { IResponse } from "@/services/QueryMaker";
import { QueryMaker } from "@/services/QueryMaker";

interface IAuthCredentials {
  email: string;
  password: string;
}

interface IAuthResponse extends IResponse {
  calendarId: string;
}

export class AuthService {
  /**
   * Login user and save auth token to cookie
   * @param credentials
   */
  public static async userLogin(credentials: IAuthCredentials): Promise<void> {
    const payload = JSON.stringify({
      email: credentials.email,
      password: credentials.password,
    });

    const res = await new QueryMaker({
      route: "/auth/login",
      payload,
    }).post<IAuthResponse>();
    const { error, message, calendarId } = res.data;

    // Save calendar to local storage
    localStorage.setItem("calendarId", calendarId);

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
    credentials: IAuthCredentials
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
