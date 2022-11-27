import axios, { type AxiosRequestConfig, type AxiosResponse } from "axios";

export interface IResponse {
  error: boolean;
  message: string;
}

interface IQueryMakerParams {
  /** API route, like `/auth/login` */
  route: string;
  /** Stringified JSON payload for POST/PATCH */
  payload?: string;
}

export class QueryMaker {
  public readonly routeUrl: string;
  private readonly baseUrl = import.meta.env.VITE_API_URL;
  private readonly axiosConfig: AxiosRequestConfig<unknown> = {
    // To pass Set-Cookie header
    withCredentials: true,
    // Handle 401 error like a normal situation
    validateStatus: (status) =>
      (status >= 200 && status < 300) || status === 401,
  };

  constructor(public readonly params: IQueryMakerParams) {
    this.routeUrl = this.baseUrl + this.params.route;
    // TODO: set query params
  }

  public async post<T = IResponse>(): Promise<AxiosResponse<T>> {
    return axios.post<T>(this.routeUrl, this.params.payload, this.axiosConfig);
  }

  public async patch<T = IResponse>(): Promise<AxiosResponse<T>> {
    return axios.patch<T>(this.routeUrl, this.params.payload, this.axiosConfig);
  }

  public async delete<T = IResponse>(): Promise<AxiosResponse<T>> {
    return axios.delete<T>(this.routeUrl, this.axiosConfig);
  }

  public async get<T = IResponse>(): Promise<AxiosResponse<T>> {
    return axios.get<T>(this.routeUrl, this.axiosConfig);
  }
}
