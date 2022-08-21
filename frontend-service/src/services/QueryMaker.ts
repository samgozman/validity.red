import axios, { type AxiosResponse } from "axios";

export interface IResponse {
  error: boolean;
  message: string;
}

export class QueryMaker {
  /**
   * Send POST query to the handler API
   * @param payload Stringified JSON payload
   * @param route API route, like '/auth/login'
   * @returns
   */
  public static async post<T = IResponse>(
    payload: string,
    route = "/handle"
  ): Promise<AxiosResponse<T>> {
    // TODO: Get URL from ENV
    const url = "http://localhost:8080" + route;
    return axios.post<T>(url, payload, {
      // To pass Set-Cookie header
      withCredentials: true,
      // Handle 401 error like a normal situation
      validateStatus: (status) =>
        (status >= 200 && status < 300) || status === 401,
    });
  }
}
