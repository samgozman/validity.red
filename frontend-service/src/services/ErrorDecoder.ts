import { AxiosError } from "axios";

export class ErrorDecoder {
  /**
   * Decode error from Axios and return a string.
   * Send all errors with status >= 500 to Sentry && all unknowns.
   * @param error
   * @returns string error message for user
   */
  public static async decode(error: unknown): Promise<string> {
    let message = "";
    if (error instanceof AxiosError) {
      message = String(error.response?.data?.message || error.message);

      if (error.response?.status === 404) {
        // TODO: navigate user to 404 page
      }

      if (error.response?.status && error.response?.status >= 500) {
        // TODO: Log error to Sentry
      }
    } else {
      message = "An error occurred, please try again";
      // TODO: Log error to Sentry
    }

    return message;
  }
}
