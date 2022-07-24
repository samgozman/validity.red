import axios from "axios";

interface RefreshResponse {
  error: boolean;
  message: string;
}

/**
 * Refresh JWT token
 */
export class RefreshToken {
  public static async call() {
    const token = this.getCookie("token");
    if (!token) return;

    const payload = JSON.stringify({
      action: "UserRefreshToken",
    });

    const res = await axios.post<RefreshResponse>(
      `http://localhost:8080/handle`,
      payload,
      {
        // To pass Set-Cookie header
        withCredentials: true,
        // Handle 401 error like a normal situation
        validateStatus: (status) =>
          (status >= 200 && status < 300) || status === 401,
      }
    );

    const { error } = res.data;

    if (error) {
      // TODO: Clear token cookie
      return;
    }
  }

  /**
   * Get cookie value by name
   * @param cookieName Name of the cookie
   * @returns
   */
  private static getCookie(cookieName: string): string | undefined {
    const cookie: { [name: string]: string } = {};
    document.cookie.split(";").forEach((el) => {
      const [key, value] = el.split("=");
      cookie[key.trim()] = value;
    });
    return cookie[cookieName];
  }
}
