import type { IResponse } from "@/services/QueryMaker";
import { QueryMaker } from "@/services/QueryMaker";
import { setCalendarId, setUsersTimezone } from "@/state";
import { ResponseError } from "@/services/ErrorDecoder";

interface IAuthCredentials {
  email: string;
  password: string;
}

interface IRegisterCredentials extends IAuthCredentials {
  timezone: string;
  // hCaptcha response token
  hcaptcha: string;
}

interface IAuthResponse extends IResponse {
  calendarId: string;
  timezone: string;
}

export class AuthService {
  /**
   * Login user and save auth token to cookie
   * @param credentials
   */
  public static async userLogin(credentials: IAuthCredentials): Promise<void> {
    const payload = JSON.stringify(credentials);

    const res = await new QueryMaker({
      route: "/auth/login",
      payload,
    }).post<IAuthResponse>();
    const { error, message, calendarId, timezone } = res.data;

    setCalendarId(calendarId);
    setUsersTimezone(timezone);

    if (error) {
      throw new ResponseError(message);
    }
  }

  /**
   * Register user
   * @param credentials
   */
  public static async userRegister(
    credentials: IRegisterCredentials
  ): Promise<void> {
    const payload = JSON.stringify(credentials);

    const res = await new QueryMaker({
      route: "/auth/register",
      payload,
    }).post();
    const { error, message } = res.data;

    if (error) {
      throw new ResponseError(message);
    }
  }

  /**
   * Verify user's email
   * @param token verification JWT token
   */
  public static async userVerifyEmail(token: string): Promise<void> {
    const payload = JSON.stringify({ token });

    const res = await new QueryMaker({
      route: "/auth/verify",
      payload,
    }).post();
    const { error, message } = res.data;

    if (error) {
      throw new ResponseError(message);
    }
  }
}
