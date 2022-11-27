import type { NavbarItem } from "@/components/navbar/interfaces/NavbarItem";
import {
  PlanetOutline,
  SpeedometerOutline,
  DocumentsOutline,
  CubeOutline,
  LogInOutline,
  PersonAddOutline,
  LogOutOutline,
} from "@vicons/ionicons5";

/** Main navbar items for upper menu */
export const navbarItems: NavbarItem[] = [
  {
    name: "Home",
    url: { name: "home" },
    icon: PlanetOutline,
  },
  {
    name: "Dashboard",
    url: { name: "dashboard" },
    icon: SpeedometerOutline,
  },
  {
    name: "Documents",
    url: { name: "documents" },
    icon: DocumentsOutline,
  },
  {
    name: "About",
    url: { name: "about" },
    icon: CubeOutline,
  },
];

/** Navbar items for user drop-down menu */
export const navbarItemsUser: NavbarItem[] = [
  {
    name: "Login",
    url: { name: "login" },
    icon: LogInOutline,
  },
  {
    name: "Register",
    url: { name: "register" },
    icon: PersonAddOutline,
  },
];

/** Navbar items for authenticated user drop-down menu */
export const navbarItemsUserAuth: NavbarItem[] = [
  {
    name: "Logout",
    url: { name: "logout" },
    icon: LogOutOutline,
  },
];
