import { AxiosError } from "axios";
import * as Sentry from "@sentry/vue";
import type { Router } from "vue-router";

/**
 * Custom error class for handling expected errors
 */
export class ResponseError extends Error {
  constructor(message: string) {
    super(message);
    Object.setPrototypeOf(this, ResponseError.prototype);
  }
}

export class ErrorDecoder {
  /**
   * Decode error from Axios and return a string.
   * Send all errors with status >= 500 to Sentry && all unknowns.
   * @param error
   * @param router - optional parameter to be able to redirect to error page
   * @returns string error message for user
   */
  public static async decode(error: unknown, router?: Router): Promise<string> {
    if (error instanceof ResponseError) {
      return error.message;
    }
    if (error instanceof AxiosError && error.response?.status) {
      if (error.response?.status === 404 && router) {
        router.push("/404");
      }

      if (error.response?.status >= 500) {
        console.error(error);
        Sentry.captureException(error);
      }

      return String(error.response?.data?.message || error.message);
    }

    console.error(error);
    Sentry.captureException(error);

    return "An error occurred, please try again";
  }
}
