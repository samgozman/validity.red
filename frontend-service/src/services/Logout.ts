import { setIsAuthenticated, setCalendarId } from "@/state";

/**
 * Logout from current user: clear state, cookie and local storage
 */
export const logout = () => {
  setIsAuthenticated(false);
  setCalendarId("");
  localStorage.removeItem("user");
  document.cookie = "token=;expires=Thu, 01 Jan 1970 00:00:01 GMT;";
};
