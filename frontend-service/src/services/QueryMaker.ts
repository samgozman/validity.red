import axios, { type AxiosResponse } from "axios";

export interface IResponse {
  error: boolean;
  message: string;
}

export class QueryMaker {
  /**
   * Send POST query to the handler API
   * @param payload Stringified JSON payload
   * @returns
   */
  public static async post<T = IResponse>(
    payload: string
  ): Promise<AxiosResponse<T>> {
    return axios.post<T>("http://localhost:8080/handle", payload, {
      // To pass Set-Cookie header
      withCredentials: true,
      // Handle 401 error like a normal situation
      validateStatus: (status) =>
        (status >= 200 && status < 300) || status === 401,
    });
  }
}
