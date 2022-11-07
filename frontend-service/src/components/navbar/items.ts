import type { NavbarItem } from "@/components/navbar/interfaces/NavbarItem";

/** Main navbar items for upper menu */
export const navbarItems: NavbarItem[] = [
  {
    name: "Home",
    url: { name: "home" },
    iconClass: "planet-outline",
  },
  {
    name: "Dashboard",
    url: { name: "dashboard" },
    iconClass: "speedometer-outline",
  },
  {
    name: "Documents",
    url: { name: "documents" },
    iconClass: "documents-outline",
  },
  {
    name: "About",
    url: { name: "about" },
    iconClass: "cube-outline",
  },
];

/** Navbar items for user drop-down menu */
export const navbarItemsUser: NavbarItem[] = [
  {
    name: "Login",
    url: { name: "login" },
    iconClass: "log-in-outline",
  },
  {
    name: "Register",
    url: { name: "register" },
    iconClass: "person-add-outline",
  },
];

/** Navbar items for authenticated user drop-down menu */
export const navbarItemsUserAuth: NavbarItem[] = [
  {
    name: "Logout",
    url: { name: "logout" },
    iconClass: "log-out-outline",
  },
];
