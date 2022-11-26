import { AxiosError } from "axios";
import type { Router } from "vue-router";

export class ErrorDecoder {
  /**
   * Decode error from Axios and return a string.
   * Send all errors with status >= 500 to Sentry && all unknowns.
   * @param error
   * @param router - optional parameter to be able to redirect to error page
   * @returns string error message for user
   */
  public static async decode(error: unknown, router?: Router): Promise<string> {
    let message = "";
    if (error instanceof AxiosError) {
      message = String(error.response?.data?.message || error.message);

      if (error.response?.status === 404 && router) {
        router.push("/404");
      }

      if (error.response?.status && error.response?.status >= 500) {
        console.error(error);
        // TODO: Log error to Sentry
      }
    } else {
      message = "An error occurred, please try again";
      console.error(error);
      // TODO: Log error to Sentry
    }

    return message;
  }
}
