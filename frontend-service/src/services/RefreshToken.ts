import { QueryMaker } from "@/services/QueryMaker";

/**
 * Refresh JWT token
 */
export class RefreshToken {
  public static async call() {
    const token = this.getCookie("token");
    if (!token) return;

    try {
      await new QueryMaker({ route: "/user/token/refresh" }).get();
    } catch (error) {
      console.error("Token refresh failed!");
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
